package linked_slice

import (
	"encoding/json"

	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/vector"
)

// LinkedSliceMem implements linking to previous slices through pointers.
// Pointers are to be resolved through the data manager that handles persistence and memory limitation
type LinkedSliceMem struct {
	data     []any  // actual payload
	id, prev uint64 // identifiers of linked slices
	db       state.DbLinked
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
	l.db.SaveNode(l.id, l.prev, l.Serialize())
}

func (l *LinkedSliceMem) Spawn() *LinkedSliceMem {
	return &LinkedSliceMem{
		make([]any, 0),
		l.db.NextId(),
		l.prev,
		l.db,
	}
}

// GetIterator allows to navigate linked slices effectively through the data db management
// Iteration starts from the end and navigates through the linked lists until the initial element found
func (l *LinkedSliceMem) GetIterator() vector.Iterator {
	curSlice := l
	curIndex := len(l.data) - 1

	return func() any {
		for curIndex < 0 { // this loop skips empty slices in the chain
			if curSlice.prev == 0 {
				return nil // end of the iteration
			}
			parent, state := l.db.FindNode(curSlice.prev)
			curSlice = Deserialize(curSlice.prev, parent, state, l.db)
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

func MakeLinkedSliceMem(prev uint64, db state.DbLinked) *LinkedSliceMem {
	s := &LinkedSliceMem{
		id:   db.NextId(),
		prev: prev,
		db:   db,
	}
	db.SaveNode(s.id, s.prev, s.Serialize())
	return s
}

func Deserialize(id, prev uint64, state []byte, db state.DbLinked) *LinkedSliceMem {
	var decodedData []any
	err := json.Unmarshal(state, &decodedData)
	if err != nil {
		panic(err)
	}

	return &LinkedSliceMem{
		decodedData,
		id,
		prev,
		db,
	}
}