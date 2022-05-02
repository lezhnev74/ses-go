package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"ses_pm_antlr/antlr_parser"
	"ses_pm_antlr/ses"
)

type SESListener struct {
	*antlr_parser.BaseSESParserListener

	*ses.SES
	curSet  int
	groupBy *string
}

// ParseSESQuery read input text and makes an IR of it (ses)
func ParseSESQuery(query string) *ses.SES {
	listener := &SESListener{
		SES: ses.MakeSES([][]*ses.Event{}, ""),
	}
	listener.SES.AddSet()

	// Configure ANTLR
	input := antlr.NewInputStream(query)
	lexer := antlr_parser.NewSESParserLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parse and Walk
	p := antlr_parser.NewSESParserParser(tokenStream)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Parse())

	// Listener contains the parsed structure
	return listener.SES
}

func (s *SESListener) EnterEvent(ctx *antlr_parser.EventContext) {
	// 1. Extract data to make a new event spec

	// 1.1 Event name
	eventName := ctx.GetName().GetText()

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
	conditions := make([]*ses.Condition, len(ctx.AllEvent_expression()))
	mapAttr := func(op antlr_parser.IExpr_operandContext) any {
		switch o := op.(type) {
		case *antlr_parser.AttrModifiedContext:
			operandEventName := eventName // if no name specified, use the one from the current event statement
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
			operandEventName := eventName // if no name specified, use the one from the current event statement
			if o.EventAttr().GetEventName() != nil {
				operandEventName = o.EventAttr().GetEventName().GetText()
			}

			operandIsCurrent := false // if we reference an event attribute, by default it means the vector of all events ('event a where a.x<b.x')
			if operandEventName == eventName {
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
	for i, expr := range ctx.AllEvent_expression() {
		// operator
		op := ""
		switch expr.GetOp().GetText() {
		case "=", "!=", "~=", ">", ">=", "<", "<=":
			op = expr.GetOp().GetText()
		default:
			panic(fmt.Sprintf("Unsupported operator %s", expr.GetOp().GetText()))
		}
		// operands
		left, right := mapAttr(expr.GetLeft()), mapAttr(expr.GetRight())
		conditions[i] = ses.MakeCondition(op, left, right)
	}

	// 2. Push to the current event set
	s.AddEvent(s.curSet, ses.MakeEvent(eventName, from, to, conditions))
}

func (s *SESListener) EnterSes_window(ctx *antlr_parser.Ses_windowContext) {
	s.curSet++
	s.AddSet()
}

func (s *SESListener) EnterGroup(ctx *antlr_parser.GroupContext) {
	s.SES.SetGroupBy(ctx.EventAttr().GetText())
}
