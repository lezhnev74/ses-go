// Package vector implements operations on vectors.
// It accepts arbitrary values, checks their types and decides what operator version to apply.
package vector

// Operator codes
const (
	OP_EQ  = "="
	OP_NEQ = "!="
	OP_GT  = ">"
	OP_GTE = ">="
	OP_LT  = "<"
	OP_LTE = "<="
)

// Matching Strategy
const (
	OP_ALL = iota
	OP_ANY
)
