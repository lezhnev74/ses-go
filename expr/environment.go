package expr

import (
	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
)

// Environment is an area with visible captured/data
type Environment interface {
	Resolve(operand ses.EventAttributeOperand) any
}

// SimpleEnvironment used for testing only
type SimpleEnvironment struct {
	captured map[string]*vector.LinkedSlice
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
		return vector.MakeLinkedSlice(nil).GetIterator()
	}
	return val.GetIterator()
}
