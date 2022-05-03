package ses

import (
	"fmt"
	"time"
)

type Window struct {
	Skip, Within time.Duration
}

type Set struct {
	events []*Event
	window Window // note: first set's window serves as global window as of now
}

func (s *Set) GetWindow() Window     { return s.window }
func (s *Set) GetEvents() []*Event   { return s.events }
func (s *Set) AddEvent(event *Event) { s.events = append(s.events, event) }

func (s *Set) validate() {
	// ---------------------------------------------------------------------------------------------------------
	// within one set events should not be dependent on each other as the events are not ordered within the set
	// Check that conditions do not refer to the same set events in both operands
	for _, event := range s.events {
		for _, c := range event.GetConditions() {
			leftEventAttribute, lOk := c.GetLeftOperand().(EventAttributeOperand)
			rightEventAttribute, rOk := c.GetRightOperand().(EventAttributeOperand)

			if !lOk || !rOk {
				continue // not the case
			}

			// edge-case: event can refer to itself
			if leftEventAttribute.EventName == rightEventAttribute.EventName {
				continue
			}

			var leftEventSameSet, rightEventSameSet bool
			for _, event2 := range s.events {
				if leftEventAttribute.EventName == event2.GetName() {
					leftEventSameSet = true
				}
				if rightEventAttribute.EventName == event2.GetName() {
					rightEventSameSet = true
				}
			}

			if leftEventSameSet && rightEventSameSet {
				panic(fmt.Sprintf("operands should not be both related to this event set as events are not ordered within a set"))
			}
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

func MakeSet(events []*Event, w Window) *Set {
	return &Set{
		events: events,
		window: w,
	}
}

// SES is an ordered Sequence of Event Sets: (0:[event1,event2], 1:[event3], 2:[event4,event5,event6])
// note: within a set the order is irrelevant
type SES struct {
	sets      []*Set
	groupAttr string // group all events by this attribute value
}

func (s *SES) GetSets() []*Set                { return s.sets }
func (s *SES) GetGroupBy() string             { return s.groupAttr }
func (s *SES) SetGroupBy(groupBy string)      { s.groupAttr = groupBy }
func (s *SES) AddEvent(set int, event *Event) { s.sets[set].AddEvent(event) }

func (s *SES) GetWindow() Window {
	if len(s.sets) == 0 {
		panic("Unable to find a widow as there are no sets")
	}
	return s.sets[0].GetWindow()
}

// AddSet pushes a new empty set to the sequence
func (s *SES) AddSet(w Window) {
	s.sets = append(s.sets, &Set{
		make([]*Event, 0),
		w,
	})
}

// Validate checks the contents of the SES and ensures that SES has correct semantics, otherwise panics
func (s *SES) Validate() {
	if len(s.sets) == 0 {
		panic("no sets available for matching")
	}

	if len(s.groupAttr) == 0 {
		panic("group-by must be specified")
	}

	// ---------------------------------------------------------------------------------------------------------
	// Check that an operand refers to the unknown event
	declaredEvents := make(map[string]bool)
	for _, set := range s.sets {
		for _, setEvent := range set.events {
			declaredEvents[setEvent.GetName()] = true

			for _, c := range setEvent.GetConditions() {
				leftEventAttribute, lOk := c.GetLeftOperand().(EventAttributeOperand)
				rightEventAttribute, rOk := c.GetRightOperand().(EventAttributeOperand)

				if lOk && !declaredEvents[leftEventAttribute.EventName] {
					panic(fmt.Sprintf("event name [%s] is not recognized", leftEventAttribute.EventName))
				}
				if rOk && !declaredEvents[rightEventAttribute.EventName] {
					panic(fmt.Sprintf("event name [%s] is not recognized", rightEventAttribute.EventName))
				}
			}
		}
	}

	// ---------------------------------------------------------------------------------------------------------

	for _, set := range s.sets {
		set.validate()
	}
}

// MakeSES make a correct new SES with given default sets, otherwise panics
func MakeSES(sets []*Set, groupBy string) *SES {
	ses := &SES{
		sets:      sets,
		groupAttr: groupBy,
	}
	return ses
}
