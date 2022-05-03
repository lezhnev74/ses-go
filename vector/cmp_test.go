package vector

import (
	"fmt"
	"testing"
)

func TestMakeCmpFunc(t *testing.T) {
	type testInput struct {
		dynamicOperand any
		operator       string
		constOperand   any
		expectedResult bool
	}
	tests := []testInput{
		// EQ and NEQ:
		{"true", "=", 2.0, false},
		{2.0, "=", 2.0, true},
		{true, "=", true, true},
		// Ordered CMP:
		{1., ">", .0, true},
		{1., "<=", 1.0, true},
		{false, "<", 1.0, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			cmpFunc := MakeCmpFunc(tt.operator, tt.constOperand)
			if cmpFunc(tt.dynamicOperand) != tt.expectedResult {
				t.Errorf("MakeCmpFunc(%v) is not %v\n", tt.dynamicOperand, tt.expectedResult)
			}
		})
	}
}

func TestMakeCmpVectorFunc(t *testing.T) {
	type testInput struct {
		dynamicOperand       any
		operator             string
		dynamicVectorOperand *LinkedSlice
		strategy             int
		expectedResult       bool
	}
	tests := []testInput{
		// EQ and NEQ:
		{"true", "=", MakeLinkedSliceInitialized(nil, []any{2.0}), OP_ALL, false},
		{true, "=", MakeLinkedSliceInitialized(nil, []any{true}), OP_ALL, true},
		// Ordered CMP:
		//{1.0, ">", linked_slice.MakeLinkedSliceInitialized([]any{-5.0},
		//	&linked_slice.cappedLinkedSlice{&linked_slice.LinkedSlice{[]any{-6.0}, nil}, 1}}, OP_ALL, true},
		{1.0, ">", MakeLinkedSliceInitialized(nil, []any{-1.0, 2.0}), OP_ALL, false},
		{1.0, ">", MakeLinkedSliceInitialized(nil, []any{-1.0, 2.0}), OP_ANY, true},
		{5.0, "<=", MakeLinkedSliceInitialized(nil, []any{6.0, 5.0, 10.0}), OP_ALL, true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			cmpFunc := MakeCmpVectorFunc(tt.operator, tt.strategy)

			if cmpFunc(tt.dynamicOperand, tt.dynamicVectorOperand.GetIterator()) != tt.expectedResult {
				t.Errorf("MakeCmpVectorFunc(%v, %v) is not %v\n", tt.dynamicOperand, tt.dynamicVectorOperand, tt.expectedResult)
			}
		})
	}
}
