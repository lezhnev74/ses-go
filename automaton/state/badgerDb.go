package state

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
)

type treeNode struct {
	Id, Parent  uint64
	ChildrenIds []uint64
	Linked      bool // outside pointer exists, so data must remain
	Data        []byte
}

type BadgerDb struct {
	scope    string // unique scope for Data isolation
	db       *badger.DB
	sequence *badger.Sequence
}

// NextId generates monotonic integers (scoped to this db instance)
func (b *BadgerDb) NextId() uint64 {
	num, err := b.sequence.Next()
	if err != nil {
		panic(err)
	}
	if num == 0 {
		return b.NextId() // skip 0 key
	}
	return num
}

func (b *BadgerDb) Save(id uint64, state []byte) {
	err := b.db.Update(func(txn *badger.Txn) error {
		key := b.makeKey(id)
		return txn.Set(key, state)
	})
	if err != nil {
		panic(err)
	}
}

func (b *BadgerDb) Find(id uint64) (data []byte) {
	err := b.db.View(func(txn *badger.Txn) error {
		key := b.makeKey(id)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		data, err = item.ValueCopy(nil)
		return err
	})
	if err != nil && err.Error() != "Key not found" {
		panic(errors.WithMessagef(err, "key: %d", id))
	}

	return data
}

func (b *BadgerDb) Delete(id uint64) {
	err := b.db.Update(func(txn *badger.Txn) error {
		key := b.makeKey(id)
		return txn.Delete(key)
	})
	if err != nil {
		panic(err)
	}

	// Update parent children
}

func (b *BadgerDb) makeKey(id uint64) []byte {
	key := []byte(fmt.Sprintf("%s-%d", b.scope, id))
	return key
}

// makeChildrenKey return a key for id->children index
func (b *BadgerDb) makeChildrenKey(id uint64) []byte {
	key := []byte(fmt.Sprintf("%s-children", b.makeKey(id)))
	return key
}

func (b *BadgerDb) Close() {
	b.sequence.Release()

	err := b.db.Close()
	if err != nil {
		panic(err)
	}
}

func (b *BadgerDb) SaveNode(id, parent uint64, state []byte) {
	isNew := false
	n := b.findNode(id)
	if n == nil {
		n = &treeNode{Id: id, Parent: parent, Linked: true}
		isNew = true
	}
	n.Data = state
	b.saveNode(n)

	// if the node is new, update parent
	if !isNew {
		return
	}
	np := b.findNode(parent)
	if np == nil {
		return
	}
	np.ChildrenIds = append(np.ChildrenIds, id)
	b.saveNode(np)
}

func (b *BadgerDb) saveNode(node *treeNode) {
	s, err := json.Marshal(*node)
	if err != nil {
		panic(err)
	}
	b.Save(node.Id, s)
}

func (b *BadgerDb) deleteNode(n *treeNode) {
	b.Delete(n.Id)
	// Update parent's children
	p := b.findNode(n.Parent)
	if p == nil {
		return
	}
	for i, child := range p.ChildrenIds {
		if child == n.Id {
			p.ChildrenIds = append(p.ChildrenIds[:i], p.ChildrenIds[i+1:]...)
			break
		}
	}
	b.saveNode(p)
}

func (b *BadgerDb) FindNode(id uint64) (parent uint64, state []byte) {
	d := b.findNode(id)
	if d == nil {
		return
	}
	parent = d.Parent
	state = d.Data
	return
}

func (b *BadgerDb) Unlink(id uint64) {
	n := b.findNode(id)
	n.Linked = false
	b.saveNode(n)
	b.freeNode(n)
}

func (b *BadgerDb) findNode(id uint64) *treeNode {
	defer func() {
		if err := recover(); err != nil {
			return // just return
		}
	}()
	data := b.Find(id)
	if data == nil {
		return nil
	}

	var d treeNode
	err := json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	return &d
}

// freeNode is an iterative function
// it check the node and if it has no children/nor linked then it removes it, then checks its parent
func (b *BadgerDb) freeNode(d *treeNode) {
	if d == nil {
		return
	}

	if d.Linked {
		return
	}

	if len(d.ChildrenIds) != 0 {
		return
	}

	b.deleteNode(d)
	b.freeNode(b.findNode(d.Parent))
}

func MakeBadgerDb(scope string, inMemoryOnly bool) *BadgerDb {
	opts := badger.DefaultOptions("")
	if inMemoryOnly {
		opts = opts.WithInMemory(inMemoryOnly)
	} else {
		opts = opts.WithDir(fmt.Sprintf("db.badger.%sequence", scope)) // todo use config
	}

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	seq, err := db.GetSequence([]byte(scope), 1000)

	return &BadgerDb{
		scope:    scope,
		db:       db,
		sequence: seq,
	}
}
