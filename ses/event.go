package ses

import "fmt"

// Event of a single event in a Sequenced Event Set
// a tuple (name, quantity, condition)
type Event struct {
	eventName      string
	qtyMin, qtyMax int
	conditions     []*Condition
}

// Validate is used during the parsing phase to make sure the event data is consistent
func (e *Event) validate() {

	// At least one operand in every condition must refer to this event
	for _, c := range e.conditions {
		leftEventAttribute, lOk := c.GetLeftOperand().(EventAttributeOperand)
		rightEventAttribute, rOk := c.GetRightOperand().(EventAttributeOperand)

		if !lOk && !rOk {
			break // not the case
		}

		// Ok, at least one must refer the current event
		if lOk && leftEventAttribute.EventName == e.eventName {
			continue
		}

		if rOk && rightEventAttribute.EventName == e.eventName {
			continue
		}

		panic(fmt.Sprintf("operands are not related to %s event", e.eventName))
	}
}

func (e *Event) GetName() string             { return e.eventName }
func (e *Event) GetQty() (int, int)          { return e.qtyMin, e.qtyMax }
func (e *Event) GetConditions() []*Condition { return e.conditions }

// MakeEvent creates a valid event otherwise panics
func MakeEvent(eventName string, qtyMin, qtyMax int, conditions []*Condition) *Event {
	ev := &Event{eventName, qtyMin, qtyMax, conditions}
	ev.validate()
	return ev
}
