package parser

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/pkg/errors"
	"ses_pm_antlr/antlr_parser"
	"ses_pm_antlr/ses"
)

type SESListener struct {
	*antlr_parser.BaseSESParserListener
	curEventName, curEventNameOriginal string

	*ses.SES
	groupBy    *string
	windowSpec *string
}

// ParseSESQuery read input text and makes an IR of it (ses)
// time is needed to resolve relative time frames to absolute time values
func ParseSESQuery(query string, t time.Time) *ses.SES {
	listener := &SESListener{
		SES: ses.MakeSES(make([]*ses.Set, 0), nil, ses.SesWindow{}),
	}

	// Configure ANTLR
	input := antlr.NewInputStream(query)
	lexer := antlr_parser.NewSESParserLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parse and Walk
	p := antlr_parser.NewSESParserParser(tokenStream)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Parse())

	// If ses window is specified, resolve it against the given time
	if listener.windowSpec != nil {
		listener.SES.SetWindow(ses.MakeSesWindowFromText(*listener.windowSpec, t))
	}

	//after all validate the whole SES
	listener.SES.Validate()

	// Listener contains the parsed structure
	return listener.SES
}

func (s *SESListener) EnterSes_window(ctx *antlr_parser.Ses_windowContext) {
	windowRe := regexp.MustCompile("^(window)?(.*)[\n;]$")
	matches := windowRe.FindStringSubmatch(ctx.GetText())
	windowText := strings.TrimSpace(matches[2])
	s.windowSpec = &windowText
}

func (s *SESListener) EnterEvent(ctx *antlr_parser.EventContext) {
	// 1. Extract data to make a new event spec

	// 1.1 Event name
	s.curEventName = ctx.GetName().GetText()
	if alias := ctx.GetAlias(); alias != nil { // alias check
		s.curEventNameOriginal = s.curEventName
		s.curEventName = alias.GetText()
	}

	// 1.2 Quantity
	from, to := 1, 1 // default
	if qtyNode := ctx.GetQty(); qtyNode != nil {
		switch qty := qtyNode.(type) {
		case *antlr_parser.Qty_asteriscContext:
			from = 0
			to = math.MaxInt64
		case *antlr_parser.Qty_plusContext:
			from = 1
			to = math.MaxInt64
		case *antlr_parser.Qty_precise_altContext:
			if exact, err := strconv.Atoi(qty.GetExact().GetText()); err == nil {
				from = exact
				to = exact
			}
		case *antlr_parser.Qty_preciseContext:
			from = 0
			to = math.MaxInt64
			if fromNode := qty.GetFrom(); fromNode != nil {
				if fromParsed, err := strconv.Atoi(fromNode.GetText()); err == nil {
					from = fromParsed
				}
			}
			if toNode := qty.GetTo(); toNode != nil {
				if toParsed, err := strconv.Atoi(toNode.GetText()); err == nil {
					to = toParsed
				}
			}
		}
	}

	// 1.3 Condition
	var condition *ses.Condition
	if ctx.Where_expression() != nil {
		condition = s.parseWhereExpr(ctx.Where_expression())
	}

	// 2. Push to the current event set
	e := ses.MakeEvent(s.curEventName, s.curEventNameOriginal, from, to, condition)
	s.SES.AddEvent(len(s.SES.GetSets())-1, e)
}

func (s *SESListener) mapAttr(op antlr_parser.IExpr_operandContext) any {
	switch o := op.(type) {
	case *antlr_parser.AttrModifiedContext:
		operandEventName := s.curEventName // if no name specified, use the one from the current event statement
		if o.EventAttr().GetEventName() != nil {
			operandEventName = o.EventAttr().GetEventName().GetText()
		}

		modifier := o.GetModifier().GetText()
		operandSelectMode := ses.SelectAll // if we reference an event attribute, by default it means the vector of all events ('event a where a.x<b.x')
		switch modifier {
		case "prev":
			operandSelectMode = ses.SelectLast
		case "any":
			operandSelectMode = ses.SelectAny
		case "first":
			operandSelectMode = ses.SelectFirst
		}

		return ses.EventAttributeOperand{
			operandEventName,
			o.EventAttr().GetAttrName().GetText(),
			false,
			int8(operandSelectMode),
		}
	case *antlr_parser.AttrContext:
		operandEventName := s.curEventName // if no name specified, use the one from the current event statement
		if o.EventAttr().GetEventName() != nil {
			operandEventName = o.EventAttr().GetEventName().GetText()
		}

		operandIsCurrent := false // if we reference an event attribute, by default it means the vector of all events ('event a where a.x<b.x')
		if operandEventName == s.curEventName {
			operandIsCurrent = true // if the event name is the current name then it means only the value of the current event ('event a where a.x>2')
		}

		return ses.EventAttributeOperand{
			operandEventName,
			o.EventAttr().GetAttrName().GetText(),
			operandIsCurrent,
			int8(ses.SelectAll), // this one does not matter as this operand references the current event
		}
	case *antlr_parser.LitContext:
		switch o.GetChildren()[0].(type) {
		case *antlr_parser.Literal_numberContext:
			number, err := strconv.ParseFloat(o.GetText(), 64) // cast any number to float64
			if err != nil {
				panic(fmt.Sprintf("Unable to parse the number %s", o.GetText()))
			}
			return number
		case *antlr_parser.Literal_stringContext:
			cleanStr := strings.Trim(o.GetText(), "'\"")
			return cleanStr
		}
		panic(fmt.Sprintf("Unsupported literal %s", o.GetText()))
	default:
		panic(fmt.Sprintf("Unsupported operand %s", o.GetText()))
	}
}

func (s *SESListener) parseWhereExpr(ctx antlr_parser.IWhere_expressionContext) (c *ses.Condition) {
	switch e := ctx.(type) {
	case *antlr_parser.Expr_binContext:
		return s.parseBinaryExpr(e.Binary_expr())
	case *antlr_parser.Expr_bracketedContext:
		return s.parseWhereExpr(e.Where_expression())
	case *antlr_parser.Expr_andContext:
		return ses.MakeCondition("and", s.parseWhereExpr(e.GetLeft()), s.parseWhereExpr(e.GetRight()))
	case *antlr_parser.Expr_orContext:
		return ses.MakeCondition("or", s.parseWhereExpr(e.GetLeft()), s.parseWhereExpr(e.GetRight()))
	}
	panic(fmt.Sprintf("WHERE expression not recognized"))
}

func (s *SESListener) parseBinaryExpr(expr antlr_parser.IBinary_exprContext) *ses.Condition {
	// operator
	op := ""
	switch expr.GetOp().GetText() {
	case "=", "!=", "~=", ">", ">=", "<", "<=":
		op = expr.GetOp().GetText()
	default:
		panic(fmt.Sprintf("Unsupported operator %s", expr.GetOp().GetText()))
	}
	// operands
	left, right := s.mapAttr(expr.GetLeft()), s.mapAttr(expr.GetRight())
	return ses.MakeCondition(op, left, right)
}

func (s *SESListener) EnterSes(ctx *antlr_parser.SesContext) {
	// Parse window
	var w ses.SetWindow
	if ctx.Set_window() != nil {
		w = ses.SetWindow{
			Skip:   extractDuration(ctx.Set_window().GetSkip()),
			Within: extractDuration(ctx.Set_window().GetWithin()),
		}
	}

	s.AddSet(w)
}

func (s *SESListener) EnterGroup(ctx *antlr_parser.GroupContext) {
	var recAttrs func(ctx antlr_parser.IGroupAttrContext) (s []string)
	recAttrs = func(ctx antlr_parser.IGroupAttrContext) (s []string) {
		s = append(s, ctx.GetAttr1().GetText())
		if extra := ctx.GetExtraAttr(); extra != nil {
			for _, attr := range recAttrs(extra) {
				s = append(s, attr)
			}
		}
		return
	}

	attrs := recAttrs(ctx.GroupAttr())
	s.SES.SetGroupBy(attrs)
}

func extractDuration(ctx antlr_parser.IDateIntervalContext) (d time.Duration) {
	if ctx == nil {
		return // default zero duration returned
	}

	i, err := strconv.ParseInt(ctx.GetNum().GetText(), 10, 64)
	if err != nil {
		panic(errors.WithMessagef(err, "unable to parse duration number %s %s", ctx.GetNum().GetText(), ctx.GetUnit().GetText()))
	}
	number := time.Duration(i)
	unit := ctx.GetUnit().GetText()

	switch unit {
	default:
	case "nanosecond", "nanoseconds":
		d = time.Nanosecond * number
	case "microsecond", "microseconds":
		d = time.Microsecond * number
	case "second", "seconds":
		d = time.Second * number
	case "minute", "minutes":
		d = time.Minute * number
	case "hour", "hours":
		d = time.Hour * number
	case "day", "days":
		d = time.Hour * 24 * number
	case "week", "weeks":
		d = time.Hour * 24 * 7 * number
		panic(fmt.Errorf("unable to parse duration unit %s %s", ctx.GetNum().GetText(), ctx.GetUnit().GetText()))
	}

	if ctx.GetExtra() != nil {
		d = d + extractDuration(ctx.GetExtra())
	}

	return d
}
