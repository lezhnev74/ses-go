package vector

import (
	"fmt"
	"reflect"
	"testing"

	"ses_pm_antlr/automaton/state"
)

func TestSliceIteration(t *testing.T) {
	// every slice in a save is linked to the previous one
	tests := [][][]any{
		{{}},
		{{1.0}, {}},
		{{1.0}, {}, {2.0}},
		{{1.0, 2.0, 3.0}}, // ints are converted to floats upon serialization to JSON
		{{1.0}, {2.0, 3.0}},
		{{1.0}, {2.0, 3.0}, {}, {4.0, "hello"}},
	}

	db := state.MakeBadgerDb("scope", true)

	for i, input := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var lastSlice *LinkedSliceMem
			expectedRevertedSlice := make([]any, 0)
			for _, s := range input {
				if lastSlice == nil {
					lastSlice = MakeLinkedSliceMem(db)
				} else {
					lastSlice = lastSlice.Spawn()
				}

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
			actualSlice := lastSlice.GetIterator().ToSlice()

			if !reflect.DeepEqual(expectedSlice, actualSlice) {
				t.Errorf("%v/%[1]T not equals to %v/%[2]T", actualSlice, expectedSlice)
			}
		})
	}
}
