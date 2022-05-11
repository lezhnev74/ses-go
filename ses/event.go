package ses

import "fmt"

// Event of a single event in a Sequenced Event Set
// a tuple (name, quantity, condition)
type Event struct {
	eventName      string
	qtyMin, qtyMax int
	condition      *Condition
}

// validate is used during the parsing phase to make sure the event data is consistent
func (e *Event) validate() {
	// At least one operand in every condition must refer to this event
	if e.condition != nil {
		e.validateCurEventAttributes()
	}
}

// At least one operand in every condition must refer to this event
func (e *Event) validateCurEventAttributes() {
	leftEventAttribute, lOk := e.condition.GetLeftOperand().(EventAttributeOperand)
	rightEventAttribute, rOk := e.condition.GetRightOperand().(EventAttributeOperand)

	if !lOk && !rOk {
		return // not the case
	}

	// Ok, at least one must refer the current event
	if lOk && leftEventAttribute.EventName == e.eventName {
		return
	}

	if rOk && rightEventAttribute.EventName == e.eventName {
		return
	}

	panic(fmt.Sprintf("operands are not related to %s event", e.eventName))
}

func (e *Event) GetName() string          { return e.eventName }
func (e *Event) GetQty() (int, int)       { return e.qtyMin, e.qtyMax }
func (e *Event) GetCondition() *Condition { return e.condition }

// MakeEvent creates a valid event otherwise panics
func MakeEvent(eventName string, qtyMin, qtyMax int, condition *Condition) *Event {
	ev := &Event{eventName, qtyMin, qtyMax, condition}
	ev.validate()
	return ev
}
