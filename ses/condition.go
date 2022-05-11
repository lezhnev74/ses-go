package ses

import "fmt"

// EventAttributeOperand attribute values:
const (
	SelectAll   = iota // "a.x > b.y" equivalent to  "a.x > all(b.y)", all matched
	SelectAny          // "a.x = any(b.y)"
	SelectFirst        // "a.x < first(a.x)", first matched (not current)
	SelectLast         // "a.x < last(a.x)", last matched (not current)
)

// MakeCondition creates valid normalized Condition otherwise panics
func MakeCondition(operator string, leftOperand, rightOperand any) *Condition {
	// Normalize: current event attr should always come first
	rightEventAttribute, rOk := rightOperand.(EventAttributeOperand)
	if rOk && rightEventAttribute.IsCurrent == true {
		leftOperand, rightOperand = rightOperand, leftOperand // swap
		// update the operator after the operands swap
		switch operator {
		case ">":
			operator = "<"
		case ">=":
			operator = "<="
		case "<":
			operator = ">"
		case "<=":
			operator = ">="
		}
	}
	cond := &Condition{operator, leftOperand, rightOperand}
	cond.validate()
	return cond
}

// Condition is a result of parsing expressions in WHERE part of the event query
// usually a binary expression, can be compound with "and"/"or" operator and sub-conditions
type Condition struct {
	operation                 string
	operandLeft, operandRight any
}

func (c *Condition) GetOperator() string  { return c.operation }
func (c *Condition) GetLeftOperand() any  { return c.operandLeft }
func (c *Condition) GetRightOperand() any { return c.operandRight }

func (c *Condition) String() string {
	return fmt.Sprintf("%s %s %s", c.operandLeft, c.operation, c.operandRight)
}

// validate is used during the parsing phase to make sure the condition data is consistent
func (c *Condition) validate() {
	// At least one operand must be EventAttributeOperand
	_, operandLeftEventAttribute := c.operandLeft.(EventAttributeOperand)
	_, operandRightEventAttribute := c.operandRight.(EventAttributeOperand)
	if c.operation != "and" && c.operation != "or" && !operandLeftEventAttribute && !operandRightEventAttribute {
		panic(fmt.Sprintf("at least one operand must refer to an event attribute"))
	}
	//note: IsCurrent is checked in Event.validate()

	// Check that only floats can be compared
	if c.operation == ">" || c.operation == ">=" || c.operation == "<" || c.operation == "<=" {
		_, isFloat := c.operandRight.(float64)
		_, isEventAttr := c.operandRight.(EventAttributeOperand)
		if !isFloat && !isEventAttr {
			panic(fmt.Sprintf("%T operand can not be used with this operator", c.operandRight))
		}
	}

	// RE_EQ only works with a string
	if _, isString := c.operandRight.(string); c.operation == "~=" && !isString {
		panic(fmt.Sprintf("%T operand can not be used with ~= operator", c.operandRight))
	}
}

// EventAttributeOperand allows specifying which events are to be selected for a vectorized operation
type EventAttributeOperand struct {
	EventName, EventAttribute string
	// IsCurrent specifies if the operand references the current event or the captured events
	IsCurrent bool
	// SelectMode specifies the amount of captured data to be selected for matching
	SelectMode int8
}

func (op EventAttributeOperand) String() string {
	return op.GetQualifiedAttributeName()
}

func (op *EventAttributeOperand) GetQualifiedAttributeName() string {
	return fmt.Sprintf("%s.%s", op.EventName, op.EventAttribute)
}
