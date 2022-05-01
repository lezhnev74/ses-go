package automaton

import (
	"fmt"
	"reflect"
	"testing"
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
				buf = MakeLastEventAttrBuffer(buf) // link with the previous one
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

			var buf *FirstEventAttrBuffer
			for _, bufferContent := range tt.prepareBufferContents { // every slice is a  buffer with values
				buf = MakeFirstEventAttrBuffer(buf) // link with the previous one
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
			[][]any{{1, 2, 1, 2, 3}},
			[]any{1, 2, 3},
		},
		{
			[][]any{{1, 2, 1}, {1, 3, 2}},
			[]any{1, 2, 3},
		},
		{
			[][]any{{"a", 2}, {true, "a"}},
			[]any{"a", 2, true},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			var buf *UniqueEventAttrBuffer
			for _, bufferContent := range tt.prepareBufferContents { // every slice is a  buffer with values
				buf = MakeUniqueEventAttrBuffer(buf) // link with the previous one
				for _, val := range bufferContent {
					buf.Accept(val)
				}
			}

			actual := buf.GetIterator().ToSlice()
			expected := tt.expectedContent
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("got = %v, want %v", actual, expected)
			}
		})
	}
}
