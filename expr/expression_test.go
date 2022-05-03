package expr

import (
	"fmt"
	"reflect"
	"testing"

	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
)

func TestMakeExpressionFromCondition(t *testing.T) {
	tests := []struct {
		cond               *ses.Condition
		env                Environment
		expectedEvalResult any
	}{
		// current event and a scalar value EQ
		{
			ses.MakeCondition("=", ses.EventAttributeOperand{"a", "x", true, ses.SelectAll}, 2.0),
			&SimpleEnvironment{
				current: map[string]any{"a.x": 2.0},
			},
			true,
		},
		// current event and a scalar value GT
		{
			ses.MakeCondition("<", ses.EventAttributeOperand{"a", "x", true, ses.SelectAll}, 2.0),
			&SimpleEnvironment{
				current: map[string]any{"a.x": 2.0},
			},
			false,
		},
		// current event and a vector EQ ALL
		{
			ses.MakeCondition("=", ses.EventAttributeOperand{"a", "x", true, ses.SelectAll}, ses.EventAttributeOperand{"b", "x", false, ses.SelectAll}),
			&SimpleEnvironment{
				current:  map[string]any{"a.x": 2.0},
				captured: map[string]*vector.LinkedSlice{"b.x": vector.MakeLinkedSliceInitialized(nil, []any{2.0, 2.0})},
			},
			true,
		},
		// current event and a vector NEQ ALL
		{
			ses.MakeCondition("!=", ses.EventAttributeOperand{"a", "x", true, ses.SelectAll}, ses.EventAttributeOperand{"b", "x", false, ses.SelectAll}),
			&SimpleEnvironment{
				current:  map[string]any{"a.x": 2.0},
				captured: map[string]*vector.LinkedSlice{"b.x": vector.MakeLinkedSliceInitialized(nil, []any{1.0, 3.0})},
			},
			true,
		},
		// current event and a vector NEQ ANY
		{
			ses.MakeCondition("!=", ses.EventAttributeOperand{"a", "x", true, ses.SelectAll}, ses.EventAttributeOperand{"b", "x", false, ses.SelectAny}),
			&SimpleEnvironment{
				current:  map[string]any{"a.x": 2.0},
				captured: map[string]*vector.LinkedSlice{"b.x": vector.MakeLinkedSliceInitialized(nil, []any{2.0, 1.0, 0.0})},
			},
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			expr := MakeExpressionFromCondition(tt.cond)
			if !reflect.DeepEqual(expr(tt.env), tt.expectedEvalResult) {
				t.Errorf("MakeExpressionFromCondition() = %v, want %v", expr(tt.env), tt.expectedEvalResult)
			}
		})
	}
}
