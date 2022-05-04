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
			`event signed_up group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
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
						ses.SetWindow{},
					),
				},
				"user_id",
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{,10} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, 10, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{1,} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{,} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*qty plus*/
		{
			`event signed_up+ group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*qty asterisk*/
		{
			`event signed_up* group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, make([]*ses.Condition, 0)),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*Condition EQ*/
		{
			`event signed_up{1,2} where name="Jar Jar Binks" group by session`, // expands the missing event name to the current event
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 2, []*ses.Condition{
							ses.MakeCondition("=", ses.EventAttributeOperand{"signed_up", "name", true, ses.SelectAll}, "Jar Jar Binks"),
						}),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*Condition Multiple*/
		{
			`event signed_up where signed_up.a <= 2 and 1000 != signed_up.amount group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{
							ses.MakeCondition("<=", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, 2.0),
							ses.MakeCondition("!=", 1000.0, ses.EventAttributeOperand{"signed_up", "amount", true, ses.SelectAll}),
						}),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/*Condition With Modified Event Attribute*/
		{
			`event signed_up where signed_up.a < prev(signed_up.a) group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{
							ses.MakeCondition("<", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, ses.EventAttributeOperand{"signed_up", "a", false, ses.SelectLast}),
						}),
					},
						ses.SetWindow{},
					),
				},
				"session",
				ses.SesWindow{},
			),
		},
		/* SetWindow */
		{
			`within 1 minute event signed_up group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, []*ses.Condition{}),
					},
						ses.SetWindow{0, time.Minute},
					),
				},
				"session",
				ses.SesWindow{},
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
			`window from last day to last day within 60 days event a group by b`,
			`don't mix FROM and WITHIN in %s`,
		},
		{
			`event signed_up where 2>1 group by session`,
			`at least one operand must refer to an event attribute`,
		},
		{
			`event signed_up
			 event logged_in where signed_up.a<logged_in.b
			 group by session`,
			`operands should not be both related to this event set as events are not ordered within a set`,
		},
		{
			`event signed_up where other.a<another.b group by session`,
			`operands are not related to signed_up event`,
		},
		{
			`event signed_up where signed_up.a<another.b group by session`,
			`event name [another] is not recognized`,
		},
		{
			`event signed_up where signed_up.a>="name" group by session`,
			`string operand can not be used with this operator`,
		},
		{
			`event signed_up`,
			`group-by must be specified`,
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
