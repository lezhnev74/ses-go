package linked_slice

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
	"ses_pm_antlr/vector"
)

// DataPool a pool of linked slices managed to fit into the given memory
// The idea is to be able to work with sequences of slices of arbitrary length within limited memory
type DataPool struct {
	badger   *badger.DB
	sequence *badger.Sequence
	name     string // unique name for scoping the pool
}

func (p *DataPool) MakeSlice(prev uint64) *LinkedSliceMem {
	s := &LinkedSliceMem{
		make([]any, 0),
		p.nextId(),
		prev,
		p,
	}
	s.pool.save(s)
	return s
}

func (p *DataPool) nextId() uint64 {
	num, err := p.sequence.Next()
	if err != nil {
		panic(err)
	}
	if num == 0 {
		return p.nextId() // skip 0 key
	}
	return num
}

func (p *DataPool) save(l *LinkedSliceMem) {

	// serialize
	encodedSlice, err := json.Marshal(l.data)
	if err != nil {
		panic(err)
	}
	serialized := fmt.Sprintf(`%d %s`, l.prev, encodedSlice)

	// save
	err = p.badger.Update(func(txn *badger.Txn) error {
		key := []byte(fmt.Sprintf("%s-%d", p.name, l.id)) // id name forging
		return txn.Set(key, []byte(serialized))
	})
	if err != nil {
		panic(err)
	}
}

func (p *DataPool) get(id uint64) *LinkedSliceMem {
	// read
	var data []byte
	err := p.badger.View(func(txn *badger.Txn) error {
		key := []byte(fmt.Sprintf("%s-%d", p.name, id)) // id name forging
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		data, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		panic(errors.WithMessagef(err, "key: %d", id))
	}

	if data == nil {
		return nil
	}

	// deserialize
	var prev uint64
	var payload string
	_, err = fmt.Sscanf(string(data), `%d %s`, &prev, &payload) // serialization format
	if err != nil {
		panic(err)
	}

	var decodedData []any
	err = json.Unmarshal([]byte(payload), &decodedData)
	if err != nil {
		panic(err)
	}

	return &LinkedSliceMem{
		decodedData,
		id,
		prev,
		p,
	}
}

func (p *DataPool) Close() {
	p.sequence.Release()

	err := p.badger.Close()
	if err != nil {
		panic(err)
	}
}

func MakeDataPool(name string, inMemoryOnly bool) *DataPool {
	opts := badger.DefaultOptions("")
	if inMemoryOnly {
		opts = opts.WithInMemory(inMemoryOnly)
	} else {
		opts = opts.WithDir(fmt.Sprintf("db.badger.%s", name))
	}

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	seq, err := db.GetSequence([]byte(name), 1000)

	return &DataPool{
		badger:   db,
		sequence: seq,
		name:     name,
	}
}

// LinkedSliceMem implements linking to previous slices through pointers.
// Pointers are to be resolved through the data manager that handles persistence and memory limitation
type LinkedSliceMem struct {
	data     []any  // actual payload
	id, prev uint64 // identifiers of linked slices
	pool     *DataPool
}

func (l *LinkedSliceMem) Append(v any) {
	l.data = append(l.data, v)
	l.pool.save(l)
}

func (l *LinkedSliceMem) Spawn() *LinkedSliceMem {
	return l.pool.MakeSlice(l.id)
}

// GetIterator allows to navigate linked slices effectively through the data pool management
// Iteration starts from the end and navigates through the linked lists until the initial element found
func (l *LinkedSliceMem) GetIterator() vector.Iterator {
	curSlice := l
	curIndex := len(l.data) - 1

	return func() any {
		for curIndex < 0 { // this loop skips empty slices in the chain
			if curSlice.prev == 0 {
				return nil // end of the iteration
			}
			curSlice = l.pool.get(curSlice.prev)
			curIndex = len(curSlice.data) - 1
		}

		if curSlice == nil {
			return nil
		}

		val := curSlice.data[curIndex]
		curIndex = curIndex - 1
		return val
	}
}
