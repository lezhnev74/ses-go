package automaton

import (
	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
)

// EventAttrBuffer is a container with public access methods, but hidden Data persisting strategies
// it implements different Data strategies (like keep unique values, min/max, calculate average, keep only the last/first one etc.)
type EventAttrBuffer interface {
	Accept(value any)             // store a new matched value to keep for later (store accordingly to the buffer impl)
	GetIterator() vector.Iterator // return an iterator (reads the next value on each call)
	Spawn() EventAttrBuffer       // start a new buffer with the same type (may link to the previous buffer for Data integrity)
}

// SimpleEventAttrBuffer keeps all appending values, no check for duplication or anything
type SimpleEventAttrBuffer struct {
	values *vector.LinkedSliceMem
}

func MakeSimpleEventAttrBuffer(db state.Db) *SimpleEventAttrBuffer {
	slice := vector.MakeLinkedSliceMem(db)
	return &SimpleEventAttrBuffer{slice}
}

func (b *SimpleEventAttrBuffer) Accept(value any) {
	b.values.Append(value)
}

func (b *SimpleEventAttrBuffer) GetIterator() vector.Iterator {
	return b.values.GetIterator()
}

func (b *SimpleEventAttrBuffer) Spawn() EventAttrBuffer {
	// Whenever we fork the buffer, we want to for the slice twice to freeze the current state of the buffer
	// new and current buffer will get new underlying slices that point to the same frozen slice
	// and both will be able to grow independently sharing the same initial slice
	s1, s2 := b.values.Spawn(), b.values.Spawn()
	b.values = s1
	return &SimpleEventAttrBuffer{s2}
}

// UniqueEventAttrBuffer keeps only unique values (with respect to previous linked buffers)
// useful for all sorts of comparison like "a.x<b.x" or "a.x=any(b.x)"
type UniqueEventAttrBuffer struct {
	values *vector.LinkedSliceMem
}

func MakeUniqueEventAttrBuffer(db state.Db) *UniqueEventAttrBuffer {
	slice := vector.MakeLinkedSliceMem(db)
	return &UniqueEventAttrBuffer{slice}
}

func (b *UniqueEventAttrBuffer) Accept(value any) {
	for next := b.values.GetIterator(); ; {
		v := next()
		if v == nil {
			break // finish iteration
		}

		if v == value {
			return // avoid duplications as the same value exists
		}
	}

	b.values.Append(value)
}

func (b *UniqueEventAttrBuffer) GetIterator() vector.Iterator {
	return b.values.GetIterator()
}

func (b *UniqueEventAttrBuffer) Spawn() EventAttrBuffer {
	// Whenever we fork the buffer, we want to for the slice twice to freeze the current state of the buffer
	// new and current buffer will get new underlying slices that point to the same frozen slice
	// and both will be able to grow independently sharing the same initial slice
	s1, s2 := b.values.Spawn(), b.values.Spawn()
	b.values = s1
	return &UniqueEventAttrBuffer{s2}
}

// FirstEventAttrBuffer only ever memorizes the first accepted value
type FirstEventAttrBuffer struct {
	// no link to previous buffers as we are only interested in the first value
	// upon spawning the first value is copied to the new buffer
	value any
}

func MakeFirstEventAttrBuffer() *FirstEventAttrBuffer {
	return &FirstEventAttrBuffer{}
}

func (b *FirstEventAttrBuffer) Accept(value any) {
	if b.value != nil {
		return // value already present
	}

	b.value = value // save the first value
}

func (b *FirstEventAttrBuffer) GetIterator() vector.Iterator {
	i := 0
	return func() any { // return the most recent one (this one)
		if i > 0 {
			return nil
		}
		i++
		return b.value
	}
}

func (b *FirstEventAttrBuffer) Spawn() EventAttrBuffer {
	return &FirstEventAttrBuffer{b.value}
}

// LastEventAttrBuffer only memorizes the last value and nothing else
type LastEventAttrBuffer struct {
	// no link to previous buffers as we are only interested in the latest value
	// upon spawning the last value is copied to the new buffer
	value any
}

func MakeLastEventAttrBuffer() *LastEventAttrBuffer {
	return &LastEventAttrBuffer{}
}

func (b *LastEventAttrBuffer) Accept(value any) {
	b.value = value // remember the last one
}

func (b *LastEventAttrBuffer) GetIterator() vector.Iterator {
	i := 0
	return func() any { // return the most recent one (this one)
		if i > 0 {
			return nil
		}
		i++
		return b.value
	}
}

func (b *LastEventAttrBuffer) Spawn() EventAttrBuffer {
	b2 := MakeLastEventAttrBuffer()
	b2.value = b.value // copy the current value and don't make a chain
	return b2
}

func MakeBufferFromSpec(spec attrBufferSpec, db state.Db) EventAttrBuffer {
	switch spec.operand.SelectMode {
	case ses.SelectFirst:
		return MakeFirstEventAttrBuffer()
	case ses.SelectLast:
		return MakeLastEventAttrBuffer()
	default:
		return MakeUniqueEventAttrBuffer(db)
	}
}
