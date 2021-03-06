// Code generated from /home/dmitry/Code/go/src/ses_query/grammar/SESParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr_parser // SESParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SESParserListener is a complete listener for a parse tree produced by SESParserParser.
type SESParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterSes_window is called when entering the ses_window production.
	EnterSes_window(c *Ses_windowContext)

	// EnterSes is called when entering the ses production.
	EnterSes(c *SesContext)

	// EnterEvent is called when entering the event production.
	EnterEvent(c *EventContext)

	// EnterQty_plus is called when entering the qty_plus production.
	EnterQty_plus(c *Qty_plusContext)

	// EnterQty_asterisc is called when entering the qty_asterisc production.
	EnterQty_asterisc(c *Qty_asteriscContext)

	// EnterQty_precise is called when entering the qty_precise production.
	EnterQty_precise(c *Qty_preciseContext)

	// EnterQty_precise_alt is called when entering the qty_precise_alt production.
	EnterQty_precise_alt(c *Qty_precise_altContext)

	// EnterSet_window is called when entering the set_window production.
	EnterSet_window(c *Set_windowContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterAbsoluteDate is called when entering the absoluteDate production.
	EnterAbsoluteDate(c *AbsoluteDateContext)

	// EnterRelativeDate is called when entering the relativeDate production.
	EnterRelativeDate(c *RelativeDateContext)

	// EnterDateInterval is called when entering the dateInterval production.
	EnterDateInterval(c *DateIntervalContext)

	// EnterExpr_or is called when entering the expr_or production.
	EnterExpr_or(c *Expr_orContext)

	// EnterExpr_bin is called when entering the expr_bin production.
	EnterExpr_bin(c *Expr_binContext)

	// EnterExpr_bracketed is called when entering the expr_bracketed production.
	EnterExpr_bracketed(c *Expr_bracketedContext)

	// EnterExpr_and is called when entering the expr_and production.
	EnterExpr_and(c *Expr_andContext)

	// EnterBinary_expr is called when entering the binary_expr production.
	EnterBinary_expr(c *Binary_exprContext)

	// EnterAttrModified is called when entering the attrModified production.
	EnterAttrModified(c *AttrModifiedContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterLit is called when entering the lit production.
	EnterLit(c *LitContext)

	// EnterEventAttr is called when entering the eventAttr production.
	EnterEventAttr(c *EventAttrContext)

	// EnterLiteral_number is called when entering the literal_number production.
	EnterLiteral_number(c *Literal_numberContext)

	// EnterLiteral_string is called when entering the literal_string production.
	EnterLiteral_string(c *Literal_stringContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterGroupAttr is called when entering the groupAttr production.
	EnterGroupAttr(c *GroupAttrContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitSes_window is called when exiting the ses_window production.
	ExitSes_window(c *Ses_windowContext)

	// ExitSes is called when exiting the ses production.
	ExitSes(c *SesContext)

	// ExitEvent is called when exiting the event production.
	ExitEvent(c *EventContext)

	// ExitQty_plus is called when exiting the qty_plus production.
	ExitQty_plus(c *Qty_plusContext)

	// ExitQty_asterisc is called when exiting the qty_asterisc production.
	ExitQty_asterisc(c *Qty_asteriscContext)

	// ExitQty_precise is called when exiting the qty_precise production.
	ExitQty_precise(c *Qty_preciseContext)

	// ExitQty_precise_alt is called when exiting the qty_precise_alt production.
	ExitQty_precise_alt(c *Qty_precise_altContext)

	// ExitSet_window is called when exiting the set_window production.
	ExitSet_window(c *Set_windowContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitAbsoluteDate is called when exiting the absoluteDate production.
	ExitAbsoluteDate(c *AbsoluteDateContext)

	// ExitRelativeDate is called when exiting the relativeDate production.
	ExitRelativeDate(c *RelativeDateContext)

	// ExitDateInterval is called when exiting the dateInterval production.
	ExitDateInterval(c *DateIntervalContext)

	// ExitExpr_or is called when exiting the expr_or production.
	ExitExpr_or(c *Expr_orContext)

	// ExitExpr_bin is called when exiting the expr_bin production.
	ExitExpr_bin(c *Expr_binContext)

	// ExitExpr_bracketed is called when exiting the expr_bracketed production.
	ExitExpr_bracketed(c *Expr_bracketedContext)

	// ExitExpr_and is called when exiting the expr_and production.
	ExitExpr_and(c *Expr_andContext)

	// ExitBinary_expr is called when exiting the binary_expr production.
	ExitBinary_expr(c *Binary_exprContext)

	// ExitAttrModified is called when exiting the attrModified production.
	ExitAttrModified(c *AttrModifiedContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitLit is called when exiting the lit production.
	ExitLit(c *LitContext)

	// ExitEventAttr is called when exiting the eventAttr production.
	ExitEventAttr(c *EventAttrContext)

	// ExitLiteral_number is called when exiting the literal_number production.
	ExitLiteral_number(c *Literal_numberContext)

	// ExitLiteral_string is called when exiting the literal_string production.
	ExitLiteral_string(c *Literal_stringContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitGroupAttr is called when exiting the groupAttr production.
	ExitGroupAttr(c *GroupAttrContext)
}
