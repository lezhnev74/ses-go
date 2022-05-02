package linked_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceIteration(t *testing.T) {
	// every slice in a save is linked to the previous one
	tests := [][][]any{
		{{}},
		{{1.0, 2.0, 3.0}},
		{{1.0}, {2.0, 3.0}},
		{{1.0}, {2.0, 3.0}, {}, {4.0, "hello"}},
	}

	for i, input := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			pool := MakeDataPool(t.Name(), true)
			var lastSlice *LinkedSliceMem
			expectedRevertedSlice := make([]any, 0)
			for i, s := range input {
				lastSlice = pool.MakeSlice(uint64(i))
				for _, val := range s {
					lastSlice.Append(val)
					expectedRevertedSlice = append(expectedRevertedSlice, val)
				}
			}

			// revert the slice back (as iteration is done from end-to-start)
			expectedSlice := make([]any, len(expectedRevertedSlice))
			for i, v := range expectedRevertedSlice {
				expectedSlice[len(expectedRevertedSlice)-i-1] = v
			}

			// Iterate to a slice
			actualSlice := pool.get(lastSlice.id).GetIterator().ToSlice()

			if !reflect.DeepEqual(expectedSlice, actualSlice) {
				t.Errorf("%v/%[1]T not equals to %v/%[2]T", actualSlice, expectedSlice)
			}
		})
	}
}
