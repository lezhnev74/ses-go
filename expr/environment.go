package expr

import (
	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector/linked_slice"
)

// Environment is an area with visible captured/data
type Environment interface {
	Resolve(operand ses.EventAttributeOperand) any
}

// SimpleEnvironment used for testing only
type SimpleEnvironment struct {
	captured map[string]*linked_slice.LinkedSlice
	current  map[string]any
}

func (env *SimpleEnvironment) Resolve(operand ses.EventAttributeOperand) any {
	if operand.IsCurrent { // scalar value
		val, exists := env.current[operand.GetQualifiedAttributeName()]
		if !exists {
			return nil
		}
		return val
	}

	// captured vector value
	val, exists := env.captured[operand.GetQualifiedAttributeName()]
	if !exists {
		return linked_slice.MakeLinkedSlice(nil).GetIterator()
	}
	return val.GetIterator()
}
