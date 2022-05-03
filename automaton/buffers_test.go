package automaton

import (
	"fmt"
	"reflect"
	"testing"

	"ses_pm_antlr/automaton/state"
)

func TestLastBuffer(t *testing.T) {
	tests := []struct {
		prepareBufferContents [][]any
		expectedContent       []any
	}{
		{
			[][]any{{1, 2, 1, 2, 3}},
			[]any{3},
		},
		{
			[][]any{{1, 2, 1}, {1, 3, 2}},
			[]any{2},
		},
		{
			[][]any{{"a", 2}, {true, "a"}},
			[]any{"a"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			var buf *LastEventAttrBuffer
			for _, bufferContent := range tt.prepareBufferContents { // every slice is a  buffer with values
				buf = MakeLastEventAttrBuffer() // link with the previous one
				for _, val := range bufferContent {
					buf.Accept(val)
				}
			}

			actual := buf.GetIterator().ToSlice()
			expected := tt.expectedContent
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got %v, want %v", actual, expected)
			}
		})
	}
}

func TestFirstBuffer(t *testing.T) {
	tests := []struct {
		prepareBufferContents [][]any
		expectedContent       []any
	}{
		{
			[][]any{{1, 2, 1, 2, 3}},
			[]any{1},
		},
		{
			[][]any{{1, 2, 1}, {1, 3, 2}},
			[]any{1},
		},
		{
			[][]any{{"a", 2}, {true, "a"}},
			[]any{"a"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			var buf EventAttrBuffer
			buf = MakeFirstEventAttrBuffer()                         // link with the previous one
			for _, bufferContent := range tt.prepareBufferContents { // every slice is a  buffer with values
				buf = buf.Spawn()
				for _, val := range bufferContent {
					buf.Accept(val)
				}
			}

			actual := buf.GetIterator().ToSlice()
			expected := tt.expectedContent
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got %v, want %v", actual, expected)
			}
		})
	}
}

func TestUniqueBuffer(t *testing.T) {
	tests := []struct {
		prepareBufferContents [][]any
		expectedContent       []any
	}{
		{
			[][]any{{1.0, 2.0, 1.0, 2.0, 3.0}},
			[]any{1.0, 2.0, 3.0},
		},
		{
			[][]any{{1., 2., 1.}, {1., 3., 2.}},
			[]any{1., 2., 3.},
		},
		{
			[][]any{{"a", 2.}, {true, "a"}},
			[]any{"a", 2., true},
		},
	}

	db := state.MakeBadgerDb("scope", "")
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			var buf EventAttrBuffer
			buf = MakeUniqueEventAttrBuffer(db)
			for _, bufferContent := range tt.prepareBufferContents { // every slice is a  buffer with values
				buf = buf.Spawn()
				for _, val := range bufferContent {
					buf.Accept(val)
				}
			}

			actual := buf.GetIterator().ToInvertedSlice()
			expected := tt.expectedContent
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got = %v, want %v", actual, expected)
			}
		})
	}
}
