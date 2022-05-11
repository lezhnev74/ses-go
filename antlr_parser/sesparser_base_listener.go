// Code generated from /home/dmitry/Code/go/src/ses_query/grammar/SESParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr_parser // SESParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSESParserListener is a complete listener for a parse tree produced by SESParserParser.
type BaseSESParserListener struct{}

var _ SESParserListener = &BaseSESParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSESParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSESParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSESParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSESParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseSESParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseSESParserListener) ExitParse(ctx *ParseContext) {}

// EnterSes_window is called when production ses_window is entered.
func (s *BaseSESParserListener) EnterSes_window(ctx *Ses_windowContext) {}

// ExitSes_window is called when production ses_window is exited.
func (s *BaseSESParserListener) ExitSes_window(ctx *Ses_windowContext) {}

// EnterSes is called when production ses is entered.
func (s *BaseSESParserListener) EnterSes(ctx *SesContext) {}

// ExitSes is called when production ses is exited.
func (s *BaseSESParserListener) ExitSes(ctx *SesContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseSESParserListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseSESParserListener) ExitEvent(ctx *EventContext) {}

// EnterQty_plus is called when production qty_plus is entered.
func (s *BaseSESParserListener) EnterQty_plus(ctx *Qty_plusContext) {}

// ExitQty_plus is called when production qty_plus is exited.
func (s *BaseSESParserListener) ExitQty_plus(ctx *Qty_plusContext) {}

// EnterQty_asterisc is called when production qty_asterisc is entered.
func (s *BaseSESParserListener) EnterQty_asterisc(ctx *Qty_asteriscContext) {}

// ExitQty_asterisc is called when production qty_asterisc is exited.
func (s *BaseSESParserListener) ExitQty_asterisc(ctx *Qty_asteriscContext) {}

// EnterQty_precise is called when production qty_precise is entered.
func (s *BaseSESParserListener) EnterQty_precise(ctx *Qty_preciseContext) {}

// ExitQty_precise is called when production qty_precise is exited.
func (s *BaseSESParserListener) ExitQty_precise(ctx *Qty_preciseContext) {}

// EnterQty_precise_alt is called when production qty_precise_alt is entered.
func (s *BaseSESParserListener) EnterQty_precise_alt(ctx *Qty_precise_altContext) {}

// ExitQty_precise_alt is called when production qty_precise_alt is exited.
func (s *BaseSESParserListener) ExitQty_precise_alt(ctx *Qty_precise_altContext) {}

// EnterSet_window is called when production set_window is entered.
func (s *BaseSESParserListener) EnterSet_window(ctx *Set_windowContext) {}

// ExitSet_window is called when production set_window is exited.
func (s *BaseSESParserListener) ExitSet_window(ctx *Set_windowContext) {}

// EnterDate is called when production date is entered.
func (s *BaseSESParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseSESParserListener) ExitDate(ctx *DateContext) {}

// EnterAbsoluteDate is called when production absoluteDate is entered.
func (s *BaseSESParserListener) EnterAbsoluteDate(ctx *AbsoluteDateContext) {}

// ExitAbsoluteDate is called when production absoluteDate is exited.
func (s *BaseSESParserListener) ExitAbsoluteDate(ctx *AbsoluteDateContext) {}

// EnterRelativeDate is called when production relativeDate is entered.
func (s *BaseSESParserListener) EnterRelativeDate(ctx *RelativeDateContext) {}

// ExitRelativeDate is called when production relativeDate is exited.
func (s *BaseSESParserListener) ExitRelativeDate(ctx *RelativeDateContext) {}

// EnterDateInterval is called when production dateInterval is entered.
func (s *BaseSESParserListener) EnterDateInterval(ctx *DateIntervalContext) {}

// ExitDateInterval is called when production dateInterval is exited.
func (s *BaseSESParserListener) ExitDateInterval(ctx *DateIntervalContext) {}

// EnterExpr_or is called when production expr_or is entered.
func (s *BaseSESParserListener) EnterExpr_or(ctx *Expr_orContext) {}

// ExitExpr_or is called when production expr_or is exited.
func (s *BaseSESParserListener) ExitExpr_or(ctx *Expr_orContext) {}

// EnterExpr_bin is called when production expr_bin is entered.
func (s *BaseSESParserListener) EnterExpr_bin(ctx *Expr_binContext) {}

// ExitExpr_bin is called when production expr_bin is exited.
func (s *BaseSESParserListener) ExitExpr_bin(ctx *Expr_binContext) {}

// EnterExpr_bracketed is called when production expr_bracketed is entered.
func (s *BaseSESParserListener) EnterExpr_bracketed(ctx *Expr_bracketedContext) {}

// ExitExpr_bracketed is called when production expr_bracketed is exited.
func (s *BaseSESParserListener) ExitExpr_bracketed(ctx *Expr_bracketedContext) {}

// EnterExpr_and is called when production expr_and is entered.
func (s *BaseSESParserListener) EnterExpr_and(ctx *Expr_andContext) {}

// ExitExpr_and is called when production expr_and is exited.
func (s *BaseSESParserListener) ExitExpr_and(ctx *Expr_andContext) {}

// EnterBinary_expr is called when production binary_expr is entered.
func (s *BaseSESParserListener) EnterBinary_expr(ctx *Binary_exprContext) {}

// ExitBinary_expr is called when production binary_expr is exited.
func (s *BaseSESParserListener) ExitBinary_expr(ctx *Binary_exprContext) {}

// EnterAttrModified is called when production attrModified is entered.
func (s *BaseSESParserListener) EnterAttrModified(ctx *AttrModifiedContext) {}

// ExitAttrModified is called when production attrModified is exited.
func (s *BaseSESParserListener) ExitAttrModified(ctx *AttrModifiedContext) {}

// EnterAttr is called when production attr is entered.
func (s *BaseSESParserListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BaseSESParserListener) ExitAttr(ctx *AttrContext) {}

// EnterLit is called when production lit is entered.
func (s *BaseSESParserListener) EnterLit(ctx *LitContext) {}

// ExitLit is called when production lit is exited.
func (s *BaseSESParserListener) ExitLit(ctx *LitContext) {}

// EnterEventAttr is called when production eventAttr is entered.
func (s *BaseSESParserListener) EnterEventAttr(ctx *EventAttrContext) {}

// ExitEventAttr is called when production eventAttr is exited.
func (s *BaseSESParserListener) ExitEventAttr(ctx *EventAttrContext) {}

// EnterLiteral_number is called when production literal_number is entered.
func (s *BaseSESParserListener) EnterLiteral_number(ctx *Literal_numberContext) {}

// ExitLiteral_number is called when production literal_number is exited.
func (s *BaseSESParserListener) ExitLiteral_number(ctx *Literal_numberContext) {}

// EnterLiteral_string is called when production literal_string is entered.
func (s *BaseSESParserListener) EnterLiteral_string(ctx *Literal_stringContext) {}

// ExitLiteral_string is called when production literal_string is exited.
func (s *BaseSESParserListener) ExitLiteral_string(ctx *Literal_stringContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseSESParserListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseSESParserListener) ExitGroup(ctx *GroupContext) {}

// EnterGroupAttr is called when production groupAttr is entered.
func (s *BaseSESParserListener) EnterGroupAttr(ctx *GroupAttrContext) {}

// ExitGroupAttr is called when production groupAttr is exited.
func (s *BaseSESParserListener) ExitGroupAttr(ctx *GroupAttrContext) {}
