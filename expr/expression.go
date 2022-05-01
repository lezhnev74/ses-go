package expr

import (
	"fmt"
	"regexp"

	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
)

// Expressions are tuples of an operation and operands
// operands can be direct literals("name") or symbolic names resolved at run-time (a.x)

// Expression is anything that can be evaluated in a given Environment
type Expression func(env Environment) any
type BoolExpression func(env Environment) bool

func MakeAndBoolExpression(exprs ...Expression) BoolExpression {
	return func(env Environment) bool {
		for _, e := range exprs {
			exprResult, isBool := e(env).(bool)
			if !isBool {
				panic(fmt.Sprintf("One of the expressions returns not boolean result: %T", exprResult))
			}
			if !exprResult {
				return false // all expressions must be true for AND-union
			}
		}

		return true
	}
}

func MakeExpressionFromCondition(cond *ses.Condition) Expression {

	// 1. Prepare operands
	operand1, operand2 := MakeExpressionFromOperand(cond.GetLeftOperand()), MakeExpressionFromOperand(cond.GetRightOperand())

	// 2. Make operation expr
	// note, the condition is always normalized, so the current event attr is on the left,
	// the right operand controls if the operation is vectorized (event attr) or not (literal)

	rightEventAttrOperand, yay := cond.GetRightOperand().(ses.EventAttributeOperand)
	if yay == true {
		// vectorized op b/w the current event and captured ones
		vectorStrategy := vector.OP_ALL
		if rightEventAttrOperand.SelectMode == ses.SelectAny {
			vectorStrategy = vector.OP_ANY
		}
		vectorOp := vector.MakeCmpVectorFunc(cond.GetOperator(), vectorStrategy)

		return func(env Environment) any {
			op1 := operand1(env)
			op2, isIterator := operand2(env).(vector.Iterator)
			if !isIterator {
				return false // runtime problem, but we just return false
			}
			return vectorOp(op1, op2)
		}
	}

	// otherwise operands are floats (should be)
	switch cond.GetOperator() {
	case "=":
		return func(env Environment) any { return operand1(env) == operand2(env) }
	case "~=":
		return func(env Environment) any {
			s1, isString1 := operand1(env).(string)
			s2, isString2 := operand2(env).(string)
			if !isString1 || !isString2 {
				return false
			}
			re := regexp.MustCompile(s2)
			return re.MatchString(s1)
		}
	case "!=":
		return func(env Environment) any { return operand1(env) != operand2(env) }
	case ">":
		return func(env Environment) any {
			// make sure the types are compatible, otherwise return false
			op1, isFloat := operand1(env).(float64)
			if !isFloat {
				return false // the operand is not float and thus can not be used ion comparison with a float
			}
			op2, _ := cond.GetRightOperand().(float64) // should already be float
			return op1 > op2
		}
	case ">=":
		return func(env Environment) any {
			// make sure the types are compatible, otherwise return false
			op1, isFloat := operand1(env).(float64)
			if !isFloat {
				return false // the operand is not float and thus can not be used ion comparison with a float
			}
			op2, _ := cond.GetRightOperand().(float64) // should already be float
			return op1 >= op2
		}
	case "<":
		return func(env Environment) any {
			// make sure the types are compatible, otherwise return false
			op1, isFloat := operand1(env).(float64)
			if !isFloat {
				return false // the operand is not float and thus can not be used ion comparison with a float
			}
			op2, _ := cond.GetRightOperand().(float64) // should already be float
			return op1 < op2
		}
	case "<=":
		return func(env Environment) any {
			// make sure the types are compatible, otherwise return false
			op1, isFloat := operand1(env).(float64)
			if !isFloat {
				return false // the operand is not float and thus can not be used ion comparison with a float
			}
			op2, _ := cond.GetRightOperand().(float64) // should already be float
			return op1 <= op2
		}
	}

	panic(fmt.Sprintf("Failed to make an expression"))
}

func MakeExpressionFromOperand(operand any) Expression {
	switch o := operand.(type) {
	case bool:
		return func(env Environment) any { return operand }
	case string:
		return func(env Environment) any { return operand }
	case float64:
		return func(env Environment) any { return operand }
	case ses.EventAttributeOperand:
		return func(env Environment) any {
			return env.Resolve(o)
		}
	default:
		panic(fmt.Sprintf("Unexpected operand type %T", operand))
	}
}
