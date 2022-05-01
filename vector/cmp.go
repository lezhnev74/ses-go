package vector

import (
	"fmt"
)

// MakeCmpFunc compiles a CMP function that is used to test an attribute of the current event
// This function accepts one literal value
func MakeCmpFunc(operator string, constOperand any) func(any) bool {
	// Two simple cases, "any" can be used directly
	if operator == OP_EQ {
		return func(eventAttribute any) bool { return eventAttribute == constOperand }
	}
	if operator == OP_NEQ {
		return func(eventAttribute any) bool { return eventAttribute != constOperand }
	}

	// Comparable operators like >,<, ... must be type checked
	switch constOperandTyped := constOperand.(type) {
	case float64:
		// ok the two values are comparable, now switch the operator again
		switch operator {
		case OP_GT:
			return func(eventAttribute any) bool {
				eventAttributeFloat, ok := castFloat64(eventAttribute)
				if !ok {
					return false // the input is not a float
				}
				return eventAttributeFloat > constOperandTyped
			}
		case OP_GTE:
			return func(eventAttribute any) bool {
				eventAttributeFloat, ok := castFloat64(eventAttribute)
				if !ok {
					return false // the input is not a float
				}
				return eventAttributeFloat >= constOperandTyped
			}
		case OP_LT:
			return func(eventAttribute any) bool {
				eventAttributeFloat, ok := castFloat64(eventAttribute)
				if !ok {
					return false // the input is not a float
				}
				return eventAttributeFloat < constOperandTyped
			}
		case OP_LTE:
			return func(eventAttribute any) bool {
				eventAttributeFloat, ok := castFloat64(eventAttribute)
				if !ok {
					return false // the input is not a float
				}
				return eventAttributeFloat <= constOperandTyped
			}
		default:
			panic(fmt.Sprintf("operator %s is not supported", operator))
		}

	default:
		panic(fmt.Sprintf("const operand %v(%[1]T) is not supported", constOperandTyped))
	}
}

// MakeCmpVectorFunc works with both dynamic values (not known at compile time)
// So the resulting func should contain type checks for both arguments at run-time
// strategy: ALL, ANY specifies which vector values must be tested against the given eventAttribute
func MakeCmpVectorFunc(operator string, strategy int) func(eventAttribute any, vector Iterator) bool {
	// Two simple cases, cmp operator can be used directly
	if operator == OP_EQ {
		if strategy == OP_ALL {
			return func(eventAttribute any, vector Iterator) bool {
				return allMatch(eventAttribute, vector, func(eventAttribute, vectorValue any) bool { return eventAttribute == vectorValue })
			}
		}
		return func(eventAttribute any, vector Iterator) bool {
			return anyMatch(eventAttribute, vector, func(eventAttribute, vectorValue any) bool { return eventAttribute == vectorValue })
		}
	}
	if operator == OP_NEQ {
		if strategy == OP_ALL {
			return func(eventAttribute any, vector Iterator) bool {
				return allMatch(eventAttribute, vector, func(eventAttribute, vectorValue any) bool { return eventAttribute != vectorValue })
			}
		}
		return func(eventAttribute any, vector Iterator) bool {
			return anyMatch(eventAttribute, vector, func(eventAttribute, vectorValue any) bool { return eventAttribute != vectorValue })
		}
	}

	// Ordered CMP on floats only
	return func(eventAttribute any, vector Iterator) bool {
		// test types of both operands for types compatibility
		eventAttributeFloat, floatEventAttributeOk := castFloat64(eventAttribute)

		if !floatEventAttributeOk {
			return false // if at least one is non-comparable then the whole operation is false
		}

		switch operator {
		case OP_GT: // --------------------------------------------------------------------------------------
			if strategy == OP_ALL {
				return allMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a > b })
			}
			return anyMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a > b })
		case OP_GTE: // --------------------------------------------------------------------------------------
			if strategy == OP_ALL {
				return allMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a >= b })
			}
			return anyMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a >= b })
		case OP_LT: // --------------------------------------------------------------------------------------
			if strategy == OP_ALL {
				return allMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a < b })
			}
			return anyMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a < b })
		case OP_LTE: // --------------------------------------------------------------------------------------
			if strategy == OP_ALL {
				return allMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a <= b })
			}
			return anyMatchFloat64(eventAttributeFloat, vector, func(a, b float64) bool { return a <= b })
		default:
			panic(fmt.Sprintf("operator %s is not supported", operator))
		}

		return false
	}
}
