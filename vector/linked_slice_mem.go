package vector

import (
	"encoding/json"

	"ses_pm_antlr/automaton/state"
)

// LinkedSliceMem implements linking to previous slices through pointers.
// Pointers are to be resolved through the data manager that handles persistence and memory limitation
type LinkedSliceMem struct {
	data     []any  // actual payload
	id, prev uint64 // identifiers of linked slices
	db       state.Db
}

func (l *LinkedSliceMem) Serialize() []byte {
	encodedSlice, err := json.Marshal(l.data)
	if err != nil {
		panic(err)
	}
	return encodedSlice
}

func (l *LinkedSliceMem) Append(v any) {
	l.data = append(l.data, v)
	l.save()
}

func (l *LinkedSliceMem) Spawn() *LinkedSliceMem {
	slice := &LinkedSliceMem{
		data: make([]any, 0),
		id:   l.db.NextId(),
		prev: l.id, // <-- link a new slice to the current one
		db:   l.db,
	}
	slice.save()
	return slice
}

// GetIterator allows to navigate linked slices effectively through the data db management
// Iteration starts from the end and navigates through the linked lists until the initial element found
func (l *LinkedSliceMem) GetIterator() Iterator {
	curSlice := l
	curIndex := len(l.data) - 1

	return func() any {
		for curIndex < 0 { // this loop skips empty slices in the chain
			if curSlice.prev == 0 {
				return nil // end of the iteration
			}
			parent, serialized := l.db.FindNode(curSlice.prev)
			curSlice = Deserialize(curSlice.prev, parent, serialized, l.db)
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

func (l *LinkedSliceMem) save() {
	l.db.SaveNode(l.id, l.prev, l.Serialize())
}

// MakeLinkedSliceMem starts a new chain of slices
func MakeLinkedSliceMem(db state.Db) *LinkedSliceMem {
	s := &LinkedSliceMem{
		id:   db.NextId(),
		prev: 0, // new chain
		db:   db,
	}
	db.SaveNode(s.id, s.prev, s.Serialize())
	return s
}

func Deserialize(id, prev uint64, state []byte, db state.Db) *LinkedSliceMem {
	var decodedData []any

	if state != nil {
		err := json.Unmarshal(state, &decodedData)
		if err != nil {
			panic(err)
		}
	}

	return &LinkedSliceMem{
		decodedData,
		id,
		prev,
		db,
	}
}
