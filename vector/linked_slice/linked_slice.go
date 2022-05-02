package linked_slice

import "ses_pm_antlr/vector"

// LinkedSlice allows reuse of existing data and append only new data on top of that
// most of the methods are working with pointers
type LinkedSlice struct {
	data []any // current data
	prev *cappedLinkedSlice
}
type cappedLinkedSlice struct {
	ptr *LinkedSlice // a link to the previous data (a pointer makes sure data is not copied but remains the same)
	len int          // the length of the slice that are visible to this list (as it may grow after the fork)
}

func MakeLinkedSlice(prev *LinkedSlice) *LinkedSlice {
	return MakeLinkedSliceInitialized(prev, make([]any, 0))
}

func MakeLinkedSliceInitialized(prev *LinkedSlice, data []any) *LinkedSlice {
	if prev == nil {
		return &LinkedSlice{data, nil}
	}
	return &LinkedSlice{data, &cappedLinkedSlice{prev, len(prev.data)}} // remember the size of the slice upon forking
}

// Len calculates total size of the whole linked slice
func (ls *LinkedSlice) Len() (result int) {
	result += len(ls.data)
	for prev := ls.prev; prev != nil; prev = prev.ptr.prev {
		result += prev.len
	}
	return result
}

func (ls *LinkedSlice) Append(value any) {
	ls.data = append(ls.data, value)
}

/*
GetIterator returns a function that returns the next value from the LinkedSlice or nil if iteration is over
Usage:
	for next := GetIterator(buf);; {
		val := next()
		if val == nil {
			break
		}
		// use val
	}
*/
func (ls *LinkedSlice) GetIterator() vector.Iterator {
	// collect all linked buffers with the length values
	buffers := make([]*cappedLinkedSlice, 1)
	buffers[0] = &cappedLinkedSlice{ls, len(ls.data)}        // put the current slice
	for prev := ls.prev; prev != nil; prev = prev.ptr.prev { // put previous slices
		buffers = append(buffers, prev)
	}

	// local variables are shared with the returned closure:
	curBuffer := len(buffers) - 1 // start from the last one (the deepest buffer in the chain)
	curBufferItem := -1

	// findNextPosition find the next buffer and item indexes or returns -1, -1 if none can be found
	var findNextPosition func() (int, int)
	findNextPosition = func() (int, int) {

		if curBuffer < 0 {
			return -1, -1 // out of buffers
		}

		curBufferItem++

		if curBufferItem == buffers[curBuffer].len { // the end of the current buffer reached, move to the next one
			curBuffer--
			curBufferItem = -1

			return findNextPosition() // try the luck with the next buffer
		}

		return -1, -1
	}

	return func() any {
		findNextPosition() // updates local variables
		if curBuffer == -1 {
			return nil
		}
		return buffers[curBuffer].ptr.data[curBufferItem]
	}
}
