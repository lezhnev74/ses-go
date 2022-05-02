package automaton

import (
	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
	"ses_pm_antlr/vector/linked_slice"
)

// EventAttrBuffer is a container with public access methods, but hidden Data persisting strategies
// it implements different Data strategies (like keep unique values, min/max, calculate average, keep only the last/first one etc.)
type EventAttrBuffer interface {
	Accept(value any)             // accept a new matched value to keep for later
	GetIterator() vector.Iterator // return an iterator (reads the next value on each call)
	Spawn() EventAttrBuffer       // Start a new buffer with the same type (may link to the previous buffer for Data integrity)
}

// SimpleEventAttrBuffer keeps all appending values, no check for duplication or anything
type SimpleEventAttrBuffer struct {
	values *linked_slice.LinkedSlice
}

func MakeSimpleEventAttrBufferInitialized(prev *SimpleEventAttrBuffer, data []any) *SimpleEventAttrBuffer {
	var prevLinkedSlice *linked_slice.LinkedSlice
	if prev != nil {
		prevLinkedSlice = prev.values
	}
	slice := linked_slice.MakeLinkedSlice(prevLinkedSlice)
	for _, d := range data {
		slice.Append(d)
	}
	return &SimpleEventAttrBuffer{slice}
}

func MakeSimpleEventAttrBuffer(prev *SimpleEventAttrBuffer) *SimpleEventAttrBuffer {
	return MakeSimpleEventAttrBufferInitialized(prev, nil)
}

func (b *SimpleEventAttrBuffer) Accept(value any) {
	b.values.Append(value)
}

func (b *SimpleEventAttrBuffer) GetIterator() vector.Iterator {
	return b.values.GetIterator()
}

func (b *SimpleEventAttrBuffer) Spawn() EventAttrBuffer {
	return MakeSimpleEventAttrBuffer(b)
}

// UniqueEventAttrBuffer keeps only unique values (with respect to previous linked buffers)
// useful for all sorts of comparison like "a.x<b.x" or "a.x=any(b.x)"
type UniqueEventAttrBuffer struct {
	values *linked_slice.LinkedSlice
}

func MakeUniqueEventAttrBuffer(prev *UniqueEventAttrBuffer) *UniqueEventAttrBuffer {
	var prevLinkedSlice *linked_slice.LinkedSlice
	if prev != nil {
		prevLinkedSlice = prev.values
	}
	return &UniqueEventAttrBuffer{
		linked_slice.MakeLinkedSlice(prevLinkedSlice),
	}
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
	return MakeUniqueEventAttrBuffer(b)
}

// FirstEventAttrBuffer only ever memorizes the first accepted value
type FirstEventAttrBuffer struct {
	values *linked_slice.LinkedSlice
}

func MakeFirstEventAttrBuffer(prev *FirstEventAttrBuffer) *FirstEventAttrBuffer {
	var prevLinkedSlice *linked_slice.LinkedSlice
	if prev != nil {
		prevLinkedSlice = prev.values
	}
	return &FirstEventAttrBuffer{
		linked_slice.MakeLinkedSlice(prevLinkedSlice),
	}
}

func (b *FirstEventAttrBuffer) Accept(value any) {
	if b.GetIterator()() != nil {
		return // value already present
	}

	b.values.Append(value) // save the first value
}

func (b *FirstEventAttrBuffer) GetIterator() vector.Iterator {
	return b.values.GetIterator()
}

func (b *FirstEventAttrBuffer) Spawn() EventAttrBuffer {
	return MakeFirstEventAttrBuffer(b)
}

// LastEventAttrBuffer only memorizes the last value and nothing else
type LastEventAttrBuffer struct {
	// no link to previous buffers as we are only interested in the latest value
	// upon spawning the last value is copied to the new buffer
	value any
}

func MakeLastEventAttrBuffer(prev *LastEventAttrBuffer) *LastEventAttrBuffer {
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
	return MakeLastEventAttrBuffer(b)
}

func MakeBufferFromSpec(spec attrBufferSpec) EventAttrBuffer {
	switch spec.operand.SelectMode {
	case ses.SelectFirst:
		return MakeFirstEventAttrBuffer(nil)
	case ses.SelectLast:
		return MakeLastEventAttrBuffer(nil)
	default:
		return MakeUniqueEventAttrBuffer(nil)
	}
}
