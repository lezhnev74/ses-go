package linked_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestForking(t *testing.T) {
	// test that forked slice only sees partial initial slice, because the initial slice can keep growing independently
	// and the forked slice must only see data up until the fork

	slice1 := MakeLinkedSliceInitialized(nil, []any{1})
	slice2 := MakeLinkedSlice(slice1) // fork
	slice1.Append(2)                  // modify initial slice

	if len(slice2.GetIterator().ToSlice()) == len(slice1.GetIterator().ToSlice()) {
		t.Errorf("slice2 has access to new data in slice1")
	}
}

func TestDataNotDuplicated(t *testing.T) {
	// [1 2] <- slice 1 keeps growing
	//  |
	//  [ 3 ] <- slice 2 links to a part of slice 1
	// Idea is that one slice's part is shared with others while the initial slice keeps growing
	// this test is to prove that data is not duplicated (copied) upon forking a new slice

	slice1 := MakeLinkedSliceInitialized(nil, []any{1})
	slice2 := MakeLinkedSlice(slice1)

	// append a few values that will make a new underlying array (extended size)
	slice1.Append(2)
	slice1.Append(3)
	slice1.Append(4)
	slice1.Append(5)

	// now modify the internal value
	slice1.data[0] = -1

	// now make sure that this modification is visible to slice2
	// this confirms that data is not copied to slice2
	slice2All := slice2.GetIterator().ToSlice()
	if slice2All[0] != slice1.data[0] {
		t.Errorf("data is copied: %v and %v", slice1.data, slice2All)
	}
}

func TestMakeIterator(t *testing.T) {
	tests := []struct {
		lb            LinkedSlice // make a chain of buffers manually
		expectedSlice []any       //iterate over all buffers and make a new slice with accumulated values
	}{
		// single buffer
		{
			LinkedSlice{[]any{1, 2}, nil},
			[]any{1, 2},
		},
		// nested buffers
		{
			LinkedSlice{
				[]any{1, 2},
				&cappedLinkedSlice{
					&LinkedSlice{
						[]any{true},
						&cappedLinkedSlice{
							&LinkedSlice{
								[]any{"hello"},
								nil,
							},
							1,
						},
					},
					1,
				},
			},
			[]any{"hello", true, 1, 2},
		},
		// nested buffers with empty slices
		{
			LinkedSlice{
				[]any{}, // empty
				&cappedLinkedSlice{
					&LinkedSlice{
						[]any{true},
						&cappedLinkedSlice{
							&LinkedSlice{
								nil, // empty
								nil,
							},
							0,
						},
					},
					1,
				},
			},
			[]any{true},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			accumulatedSlice := make([]any, 0)
			for next := tt.lb.GetIterator(); ; {
				val := next()
				if val == nil {
					break
				}
				accumulatedSlice = append(accumulatedSlice, val)
			}

			if !reflect.DeepEqual(accumulatedSlice, tt.expectedSlice) {
				t.Errorf("%v\nwant\n%v", accumulatedSlice, tt.expectedSlice)
			}
		})
	}
}

func TestLinkedSlice_Len(t *testing.T) {
	tests := []struct {
		lb           LinkedSlice
		expectedSize int
	}{
		// single buffer
		{
			LinkedSlice{[]any{1, 2}, nil},
			2,
		},
		// nested buffers
		{
			LinkedSlice{
				[]any{1, 2},
				&cappedLinkedSlice{
					&LinkedSlice{
						[]any{true},
						&cappedLinkedSlice{
							&LinkedSlice{
								[]any{"hello"},
								nil,
							},
							1,
						},
					},
					1,
				},
			},
			4,
		},
		// nested buffers with empty slices
		{
			LinkedSlice{
				[]any{}, // empty
				&cappedLinkedSlice{
					&LinkedSlice{
						[]any{true},
						&cappedLinkedSlice{
							&LinkedSlice{
								nil, // empty
								nil,
							},
							0,
						},
					},
					1,
				},
			},
			1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(tt.lb.Len(), tt.expectedSize) {
				t.Errorf("%v\nwant\n%v", tt.lb.Len(), tt.expectedSize)
			}
		})
	}
}
