package parser

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"ses_pm_antlr/ses"
)

func TestSimpleQueries(t *testing.T) {
	now := dateparse.MustParse("1 May 2022") // Freeze the Now

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
						ses.MakeEvent("signed_up", 1, 1, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*qty full*/
		{
			`event signed_up{1,10} group by user_id`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 10, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"user_id"},
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{,10} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, 10, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{1,} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*qty partial*/
		{
			`event signed_up{,} group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*qty plus*/
		{
			`event signed_up+ group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, math.MaxInt64, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*qty asterisk*/
		{
			`event signed_up* group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 0, math.MaxInt64, nil),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*Condition EQ*/
		{
			`event signed_up{1,2} where name="Jar Jar Binks" group by session`, // expands the missing event name to the current event
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 2, ses.MakeCondition("=", ses.EventAttributeOperand{"signed_up", "name", true, ses.SelectAll}, "Jar Jar Binks")),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*Condition Multiple*/
		{
			`event signed_up where signed_up.a <= 2 and 1000 != signed_up.amount group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, ses.MakeCondition("and",
							ses.MakeCondition("<=", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, 2.0),
							ses.MakeCondition("!=", 1000.0, ses.EventAttributeOperand{"signed_up", "amount", true, ses.SelectAll}),
						)),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/*Condition With Modified Event Attribute*/
		{
			`event signed_up where signed_up.a < prev(signed_up.a) group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, ses.MakeCondition("<", ses.EventAttributeOperand{"signed_up", "a", true, ses.SelectAll}, ses.EventAttributeOperand{"signed_up", "a", false, ses.SelectLast})),
					},
						ses.SetWindow{},
					),
				},
				[]string{"session"},
				ses.SesWindow{},
			),
		},
		/* SetWindow + multiple group by */
		{
			`within 1 second event a and then event b group by session,country,language.iso`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{ses.MakeEvent("a", 1, 1, nil)}, ses.SetWindow{0, time.Second}),
					ses.MakeSet([]*ses.Event{ses.MakeEvent("b", 1, 1, nil)}, ses.SetWindow{}),
				},
				[]string{"session", "country", "language.iso"},
				ses.SesWindow{},
			),
		},
		/* SesWindow */
		{
			`window within 1 minute; event signed_up group by session`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("signed_up", 1, 1, nil),
					},
						ses.SetWindow{0, 0},
					),
				},
				[]string{"session"},
				ses.MakeSesWindowFromText("within 1 minute", now),
			),
		},
		/* SesWindow + SetWindow */
		{
			`window 1 day; within 60 days event a group by b`,
			ses.MakeSES(
				[]*ses.Set{
					ses.MakeSet([]*ses.Event{
						ses.MakeEvent("a", 1, 1, nil),
					},
						ses.SetWindow{0, 60 * 24 * time.Hour},
					),
				},
				[]string{"b"},
				ses.MakeSesWindowFromText("within 1 day", now),
			),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := ParseSESQuery(tt.query, now)
			if !reflect.DeepEqual(*got, *tt.expected) {
				t.Errorf("\nactual:\n%v \n expected:\n%v\n", *got, *tt.expected)
			}
		})
	}
}

func TestValidation(t *testing.T) {
	now := dateparse.MustParse("1 May 2022") // Freeze the Now
	tests := []struct {
		query        string
		panicMessage string
	}{
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
			ParseSESQuery(tt.query, now)
			t.Errorf("\nshould have panicked with: %s\n", tt.panicMessage)
		})
	}
}
