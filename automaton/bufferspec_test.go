package automaton

import (
	"fmt"
	"reflect"
	"testing"

	"ses_pm_antlr/parser"
	"ses_pm_antlr/ses"
)

func TestMakeBufferSpecs(t *testing.T) {
	tests := []struct {
		input string // query
		want  []attrBufferSpec
	}{
		{
			`event a`,
			[]attrBufferSpec{},
		},
		{
			`event a 
			 then event b where a.x=b.x`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectAll}, 1},
			},
		},
		{
			`event a where a.x<prev(a.x)`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectLast}, 0},
			},
		},
		{
			`event a then event b where first(a.x)<b.y`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectFirst}, 1},
			},
		},
		{
			`event a then event b where any(a.x)=b.y`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectAny}, 1},
			},
		},
		{
			`event a 
			 then event b where any(a.x)=b.y // any value for a.x
			 then event c where a.x<c.y      // all value for a.x
			`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectAll}, 2},
			},
		},
		{
			`event a where a.x<prev(a.x)
			 then event b where b.x>first(a.x)`,
			[]attrBufferSpec{
				{ses.EventAttributeOperand{"a", "x", false, ses.SelectAny}, 1},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			s := parser.ParseSESQuery(tt.input)
			got := makeBufferSpecs(s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeBufferSpecs() = %v, want %v", got, tt.want)
			}
		})
	}
}
