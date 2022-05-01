package ses

import "fmt"

// SES is an ordered Sequence of Event Sets: (0:[event1,event2], 1:[event3], 2:[event4,event5,event6])
// note: within a set the order is irrelevant
type SES struct {
	sets      [][]*Event
	groupAttr string // group all events by this attribute value
}

func (s *SES) GetSets() [][]*Event { return s.sets }
func (s *SES) GetGroupBy() string  { return s.groupAttr }

func (s *SES) SetGroupBy(groupBy string) { s.groupAttr = groupBy }
func (s *SES) AddEvent(set int, event *Event) {
	s.sets[set] = append(s.sets[set], event)
	s.validate()
}

// AddSet pushes a new empty set to the sequence
func (s *SES) AddSet() { s.sets = append(s.sets, make([]*Event, 0)) }

// Validate checks the contents of the SES and ensures that SES has correct semantics, otherwise panics
func (s *SES) validate() {

	// ---------------------------------------------------------------------------------------------------------
	// within one set events should not be dependent on each other as the events are not ordered within the set
	for _, eventSets := range s.sets {
		// Check that conditions do not refer to the same set events in both operands
		for _, event := range eventSets {
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
				for _, event2 := range eventSets {
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
	}
	// ---------------------------------------------------------------------------------------------------------

	// ---------------------------------------------------------------------------------------------------------
	// Check that an operand refers to the unknown event
	declaredEvents := make(map[string]bool)

	for _, setEvents := range s.sets {
		// Check that conditions do not refer to the same set events in both operands
		for _, setEvent := range setEvents {
			declaredEvents[setEvent.GetName()] = true

			for _, c := range setEvent.GetConditions() {
				leftEventAttribute, lOk := c.GetLeftOperand().(EventAttributeOperand)
				rightEventAttribute, rOk := c.GetRightOperand().(EventAttributeOperand)

				if lOk && !declaredEvents[leftEventAttribute.EventName] {
					panic(fmt.Sprintf("event name %s is not recognized", leftEventAttribute.EventName))
				}
				if rOk && !declaredEvents[rightEventAttribute.EventName] {
					panic(fmt.Sprintf("event name %s is not recognized", rightEventAttribute.EventName))
				}
			}
		}
	}
	// ---------------------------------------------------------------------------------------------------------
}

// MakeSES make a correct new SES with given default sets, otherwise panics
func MakeSES(sets [][]*Event, groupAttr string) *SES {
	ses := &SES{sets, groupAttr}
	ses.validate()
	return ses
}
