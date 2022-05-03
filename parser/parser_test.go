package parser

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"

	"ses_pm_antlr/ses"
)

func TestSimpleQueries(t *testing.T) {
	tests := []struct {
		query    string
		expected *ses.SES
	}{
		/*empty*/
		{
			`event signed_up`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*qty full*/
		{
			`event signed_up{1,10} group by user_id`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 10, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"user_id",
			),
		},
		/*qty partial*/
		{
			`event signed_up{,10}`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, 10, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*qty partial*/
		{
			`event signed_up{1,}`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*qty partial*/
		{
			`event signed_up{,}`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*qty plus*/
		{
			`event signed_up+`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*qty asterisk*/
		{
			`event signed_up*`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*Condition EQ*/
		{
			`event signed_up{1,2} where name="Jar Jar Binks"`, // expands the missing event name to the current event
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 2, []*ses.Condition{
							ses.MakeCondition("=", ses.EventAttributeOperand{"signed_up", "name", true, ses.SelectAll}, "Jar Jar Binks"),
						}),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*Condition Multiple*/
		{
			`event signed_up where signed_up.a <= 2 and 1000 != signed_up.amount`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{
							ses.MakeCondition("<=", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, 2.0),
							ses.MakeCondition("!=", 1000.0, ses.EventAttributeOperand{"signed_up", "amount", true, ses.SelectAll}),
						}),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/*Condition With Modified Event Attribute*/
		{
			`event signed_up where signed_up.a < prev(signed_up.a)`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{
							ses.MakeCondition("<", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, ses.EventAttributeOperand{"signed_up", "a", false, ses.SelectLast}),
						}),
					},
						ses.Window{},
					),
				},
				"",
			),
		},
		/* Window */
		{
			`within 1 minute event signed_up`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{}),
					},
						ses.Window{0, time.Minute},
					),
				},
				"",
			),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := ParseSESQuery(tt.query)
			if !reflect.DeepEqual(*got, *tt.expected) {
				t.Errorf("\nactual:\n%v \n expected:\n%v\n", *got, *tt.expected)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	tests := []struct {
		query        string
		panicMessage string
	}{
		{
			`event signed_up where 2>1`,
			`at least one operand must refer to an event attribute`,
		},
		{
			`event signed_up
			 event logged_in where signed_up.a<logged_in.b`,
			`operands should not be both related to this event set as events are not ordered within a set`,
		},
		{
			`event signed_up where other.a<another.b`,
			`operands are not related to signed_up event`,
		},
		{
			`event signed_up where signed_up.a<another.b`,
			`event name [another] is not recognized`,
		},
		{
			`event signed_up where signed_up.a>="name"`,
			`string operand can not be used with this operator`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					panicString, _ := r.(string)
					if !reflect.DeepEqual(panicString, tt.panicMessage) {
						t.Errorf("\nactual:\t%v\nexpect:\t%v\n", r, tt.panicMessage)
					}
				}
			}()
			ParseSESQuery(tt.query)
			t.Errorf("\nshould have panicked with: %s\n", tt.panicMessage)
		})
	}
}
