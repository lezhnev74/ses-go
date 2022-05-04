// Code generated from /home/dmitry/Code/go/src/ses_query/grammar/SESParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr_parser // SESParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SESParserParser struct {
	*antlr.BaseParser
}

var sesparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sesparserParserInit() {
	staticData := &sesparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'.'", "'('", "')'", "'{'", "'}'", "','", "';'", "'+'", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "LINE_COMMENT", "SPACE", "EVENT", "WINDOW", "FROM", "TO",
		"LAST", "SKIP_", "WITHIN", "THEN", "WHERE", "AND", "OR", "GROUP", "DATE_UNIT",
		"OP_LOGICAL", "DOT", "L_BRACKET", "R_BRACKET", "L_CURLY_BRACKET", "R_CURLY_BRACKET",
		"COMMA", "SEMI", "PLUS", "ASTERISK", "NUMBER", "STRING", "ID",
	}
	staticData.ruleNames = []string{
		"parse", "window", "ses", "event", "event_qty", "set_window", "dateInterval",
		"date", "absoluteDate", "relativeDate", "event_expression", "expr_operand",
		"eventAttr", "literal", "group",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 158, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 3, 0,
		32, 8, 0, 1, 0, 4, 0, 35, 8, 0, 11, 0, 12, 0, 36, 1, 0, 3, 0, 40, 8, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 49, 8, 1, 3, 1, 51, 8,
		1, 1, 1, 3, 1, 54, 8, 1, 1, 1, 3, 1, 57, 8, 1, 1, 2, 3, 2, 60, 8, 2, 1,
		2, 4, 2, 63, 8, 2, 11, 2, 12, 2, 64, 1, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 76, 8, 3, 10, 3, 12, 3, 79, 9, 3, 3, 3, 81,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 87, 8, 4, 1, 4, 1, 4, 3, 4, 91, 8,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 97, 8, 4, 1, 5, 3, 5, 100, 8, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 105, 8, 5, 1, 5, 3, 5, 108, 8, 5, 1, 5, 1, 5, 3, 5, 112,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 118, 8, 6, 10, 6, 12, 6, 121, 9, 6,
		1, 7, 1, 7, 3, 7, 125, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 143,
		8, 11, 1, 12, 1, 12, 3, 12, 147, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 3,
		13, 153, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 0, 169, 0, 31, 1, 0, 0, 0, 2, 43,
		1, 0, 0, 0, 4, 59, 1, 0, 0, 0, 6, 66, 1, 0, 0, 0, 8, 96, 1, 0, 0, 0, 10,
		99, 1, 0, 0, 0, 12, 113, 1, 0, 0, 0, 14, 124, 1, 0, 0, 0, 16, 126, 1, 0,
		0, 0, 18, 128, 1, 0, 0, 0, 20, 131, 1, 0, 0, 0, 22, 142, 1, 0, 0, 0, 24,
		146, 1, 0, 0, 0, 26, 152, 1, 0, 0, 0, 28, 154, 1, 0, 0, 0, 30, 32, 3, 2,
		1, 0, 31, 30, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 34, 1, 0, 0, 0, 33, 35,
		3, 4, 2, 0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0,
		36, 37, 1, 0, 0, 0, 37, 39, 1, 0, 0, 0, 38, 40, 3, 28, 14, 0, 39, 38, 1,
		0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 5, 0, 0, 1, 42,
		1, 1, 0, 0, 0, 43, 50, 5, 5, 0, 0, 44, 45, 5, 6, 0, 0, 45, 48, 3, 14, 7,
		0, 46, 47, 5, 7, 0, 0, 47, 49, 3, 14, 7, 0, 48, 46, 1, 0, 0, 0, 48, 49,
		1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 44, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0,
		51, 56, 1, 0, 0, 0, 52, 54, 5, 10, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1,
		0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 3, 12, 6, 0, 56, 53, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 60, 3, 10, 5, 0, 59, 58, 1, 0, 0,
		0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 63, 3, 6, 3, 0, 62, 61,
		1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 5, 1, 0, 0, 0, 66, 67, 5, 4, 0, 0, 67, 69, 5, 29, 0, 0, 68, 70, 3,
		8, 4, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 80, 1, 0, 0, 0, 71,
		72, 5, 12, 0, 0, 72, 77, 3, 20, 10, 0, 73, 74, 5, 13, 0, 0, 74, 76, 3,
		20, 10, 0, 75, 73, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		77, 78, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 71, 1,
		0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 7, 1, 0, 0, 0, 82, 97, 5, 25, 0, 0, 83,
		97, 5, 26, 0, 0, 84, 86, 5, 21, 0, 0, 85, 87, 5, 27, 0, 0, 86, 85, 1, 0,
		0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 5, 23, 0, 0, 89,
		91, 5, 27, 0, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0,
		0, 0, 92, 97, 5, 22, 0, 0, 93, 94, 5, 21, 0, 0, 94, 95, 5, 27, 0, 0, 95,
		97, 5, 22, 0, 0, 96, 82, 1, 0, 0, 0, 96, 83, 1, 0, 0, 0, 96, 84, 1, 0,
		0, 0, 96, 93, 1, 0, 0, 0, 97, 9, 1, 0, 0, 0, 98, 100, 5, 13, 0, 0, 99,
		98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 104, 5,
		11, 0, 0, 102, 103, 5, 9, 0, 0, 103, 105, 3, 12, 6, 0, 104, 102, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 111, 1, 0, 0, 0, 106, 108, 5, 13, 0, 0,
		107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		110, 5, 10, 0, 0, 110, 112, 3, 12, 6, 0, 111, 107, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 11, 1, 0, 0, 0, 113, 114, 5, 27, 0, 0, 114, 119, 5, 16,
		0, 0, 115, 116, 5, 13, 0, 0, 116, 118, 3, 12, 6, 0, 117, 115, 1, 0, 0,
		0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		13, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 125, 3, 16, 8, 0, 123, 125,
		3, 18, 9, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 15, 1, 0,
		0, 0, 126, 127, 5, 28, 0, 0, 127, 17, 1, 0, 0, 0, 128, 129, 5, 8, 0, 0,
		129, 130, 3, 12, 6, 0, 130, 19, 1, 0, 0, 0, 131, 132, 3, 22, 11, 0, 132,
		133, 5, 17, 0, 0, 133, 134, 3, 22, 11, 0, 134, 21, 1, 0, 0, 0, 135, 136,
		5, 29, 0, 0, 136, 137, 5, 19, 0, 0, 137, 138, 3, 24, 12, 0, 138, 139, 5,
		20, 0, 0, 139, 143, 1, 0, 0, 0, 140, 143, 3, 24, 12, 0, 141, 143, 3, 26,
		13, 0, 142, 135, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0,
		143, 23, 1, 0, 0, 0, 144, 145, 5, 29, 0, 0, 145, 147, 5, 18, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149,
		5, 29, 0, 0, 149, 25, 1, 0, 0, 0, 150, 153, 5, 27, 0, 0, 151, 153, 5, 28,
		0, 0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 27, 1, 0, 0, 0,
		154, 155, 5, 15, 0, 0, 155, 156, 3, 24, 12, 0, 156, 29, 1, 0, 0, 0, 24,
		31, 36, 39, 48, 50, 53, 56, 59, 64, 69, 77, 80, 86, 90, 96, 99, 104, 107,
		111, 119, 124, 142, 146, 152,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SESParserParserInit initializes any static state used to implement SESParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSESParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SESParserParserInit() {
	staticData := &sesparserParserStaticData
	staticData.once.Do(sesparserParserInit)
}

// NewSESParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSESParserParser(input antlr.TokenStream) *SESParserParser {
	SESParserParserInit()
	this := new(SESParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sesparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SESParser.g4"

	return this
}

// SESParserParser tokens.
const (
	SESParserParserEOF             = antlr.TokenEOF
	SESParserParserWS              = 1
	SESParserParserLINE_COMMENT    = 2
	SESParserParserSPACE           = 3
	SESParserParserEVENT           = 4
	SESParserParserWINDOW          = 5
	SESParserParserFROM            = 6
	SESParserParserTO              = 7
	SESParserParserLAST            = 8
	SESParserParserSKIP_           = 9
	SESParserParserWITHIN          = 10
	SESParserParserTHEN            = 11
	SESParserParserWHERE           = 12
	SESParserParserAND             = 13
	SESParserParserOR              = 14
	SESParserParserGROUP           = 15
	SESParserParserDATE_UNIT       = 16
	SESParserParserOP_LOGICAL      = 17
	SESParserParserDOT             = 18
	SESParserParserL_BRACKET       = 19
	SESParserParserR_BRACKET       = 20
	SESParserParserL_CURLY_BRACKET = 21
	SESParserParserR_CURLY_BRACKET = 22
	SESParserParserCOMMA           = 23
	SESParserParserSEMI            = 24
	SESParserParserPLUS            = 25
	SESParserParserASTERISK        = 26
	SESParserParserNUMBER          = 27
	SESParserParserSTRING          = 28
	SESParserParserID              = 29
)

// SESParserParser rules.
const (
	SESParserParserRULE_parse            = 0
	SESParserParserRULE_window           = 1
	SESParserParserRULE_ses              = 2
	SESParserParserRULE_event            = 3
	SESParserParserRULE_event_qty        = 4
	SESParserParserRULE_set_window       = 5
	SESParserParserRULE_dateInterval     = 6
	SESParserParserRULE_date             = 7
	SESParserParserRULE_absoluteDate     = 8
	SESParserParserRULE_relativeDate     = 9
	SESParserParserRULE_event_expression = 10
	SESParserParserRULE_expr_operand     = 11
	SESParserParserRULE_eventAttr        = 12
	SESParserParserRULE_literal          = 13
	SESParserParserRULE_group            = 14
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(SESParserParserEOF, 0)
}

func (s *ParseContext) Window() IWindowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowContext)
}

func (s *ParseContext) AllSes() []ISesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISesContext); ok {
			len++
		}
	}

	tst := make([]ISesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISesContext); ok {
			tst[i] = t.(ISesContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Ses(i int) ISesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISesContext)
}

func (s *ParseContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *SESParserParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SESParserParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWINDOW {
		{
			p.SetState(30)
			p.Window()
		}

	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserEVENT)|(1<<SESParserParserTHEN)|(1<<SESParserParserAND))) != 0) {
		{
			p.SetState(33)
			p.Ses()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserGROUP {
		{
			p.SetState(38)
			p.Group()
		}

	}
	{
		p.SetState(41)
		p.Match(SESParserParserEOF)
	}

	return localctx
}

// IWindowContext is an interface to support dynamic dispatch.
type IWindowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from rule contexts.
	GetFrom() IDateContext

	// GetTo returns the to rule contexts.
	GetTo() IDateContext

	// GetWithin returns the within rule contexts.
	GetWithin() IDateIntervalContext

	// SetFrom sets the from rule contexts.
	SetFrom(IDateContext)

	// SetTo sets the to rule contexts.
	SetTo(IDateContext)

	// SetWithin sets the within rule contexts.
	SetWithin(IDateIntervalContext)

	// IsWindowContext differentiates from other interfaces.
	IsWindowContext()
}

type WindowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	from   IDateContext
	to     IDateContext
	within IDateIntervalContext
}

func NewEmptyWindowContext() *WindowContext {
	var p = new(WindowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_window
	return p
}

func (*WindowContext) IsWindowContext() {}

func NewWindowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowContext {
	var p = new(WindowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_window

	return p
}

func (s *WindowContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowContext) GetFrom() IDateContext { return s.from }

func (s *WindowContext) GetTo() IDateContext { return s.to }

func (s *WindowContext) GetWithin() IDateIntervalContext { return s.within }

func (s *WindowContext) SetFrom(v IDateContext) { s.from = v }

func (s *WindowContext) SetTo(v IDateContext) { s.to = v }

func (s *WindowContext) SetWithin(v IDateIntervalContext) { s.within = v }

func (s *WindowContext) WINDOW() antlr.TerminalNode {
	return s.GetToken(SESParserParserWINDOW, 0)
}

func (s *WindowContext) FROM() antlr.TerminalNode {
	return s.GetToken(SESParserParserFROM, 0)
}

func (s *WindowContext) AllDate() []IDateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateContext); ok {
			len++
		}
	}

	tst := make([]IDateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateContext); ok {
			tst[i] = t.(IDateContext)
			i++
		}
	}

	return tst
}

func (s *WindowContext) Date(i int) IDateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *WindowContext) DateInterval() IDateIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalContext)
}

func (s *WindowContext) TO() antlr.TerminalNode {
	return s.GetToken(SESParserParserTO, 0)
}

func (s *WindowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
}

func (s *WindowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterWindow(s)
	}
}

func (s *WindowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitWindow(s)
	}
}

func (p *SESParserParser) Window() (localctx IWindowContext) {
	this := p
	_ = this

	localctx = NewWindowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SESParserParserRULE_window)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(SESParserParserWINDOW)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserFROM {
		{
			p.SetState(44)
			p.Match(SESParserParserFROM)
		}
		{
			p.SetState(45)

			var _x = p.Date()

			localctx.(*WindowContext).from = _x
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserTO {
			{
				p.SetState(46)
				p.Match(SESParserParserTO)
			}
			{
				p.SetState(47)

				var _x = p.Date()

				localctx.(*WindowContext).to = _x
			}

		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWITHIN || _la == SESParserParserNUMBER {
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserWITHIN {
			{
				p.SetState(52)
				p.Match(SESParserParserWITHIN)
			}

		}
		{
			p.SetState(55)

			var _x = p.DateInterval()

			localctx.(*WindowContext).within = _x
		}

	}

	return localctx
}

// ISesContext is an interface to support dynamic dispatch.
type ISesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSesContext differentiates from other interfaces.
	IsSesContext()
}

type SesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySesContext() *SesContext {
	var p = new(SesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_ses
	return p
}

func (*SesContext) IsSesContext() {}

func NewSesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SesContext {
	var p = new(SesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_ses

	return p
}

func (s *SesContext) GetParser() antlr.Parser { return s.parser }

func (s *SesContext) Set_window() ISet_windowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_windowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_windowContext)
}

func (s *SesContext) AllEvent() []IEventContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEventContext); ok {
			len++
		}
	}

	tst := make([]IEventContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEventContext); ok {
			tst[i] = t.(IEventContext)
			i++
		}
	}

	return tst
}

func (s *SesContext) Event(i int) IEventContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventContext)
}

func (s *SesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterSes(s)
	}
}

func (s *SesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitSes(s)
	}
}

func (p *SESParserParser) Ses() (localctx ISesContext) {
	this := p
	_ = this

	localctx = NewSesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SESParserParserRULE_ses)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserTHEN || _la == SESParserParserAND {
		{
			p.SetState(58)
			p.Set_window()
		}

	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(61)
				p.Event()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IEventContext is an interface to support dynamic dispatch.
type IEventContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetQty returns the qty rule contexts.
	GetQty() IEvent_qtyContext

	// SetQty sets the qty rule contexts.
	SetQty(IEvent_qtyContext)

	// IsEventContext differentiates from other interfaces.
	IsEventContext()
}

type EventContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	qty    IEvent_qtyContext
}

func NewEmptyEventContext() *EventContext {
	var p = new(EventContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_event
	return p
}

func (*EventContext) IsEventContext() {}

func NewEventContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventContext {
	var p = new(EventContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_event

	return p
}

func (s *EventContext) GetParser() antlr.Parser { return s.parser }

func (s *EventContext) GetName() antlr.Token { return s.name }

func (s *EventContext) SetName(v antlr.Token) { s.name = v }

func (s *EventContext) GetQty() IEvent_qtyContext { return s.qty }

func (s *EventContext) SetQty(v IEvent_qtyContext) { s.qty = v }

func (s *EventContext) EVENT() antlr.TerminalNode {
	return s.GetToken(SESParserParserEVENT, 0)
}

func (s *EventContext) ID() antlr.TerminalNode {
	return s.GetToken(SESParserParserID, 0)
}

func (s *EventContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SESParserParserWHERE, 0)
}

func (s *EventContext) AllEvent_expression() []IEvent_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEvent_expressionContext); ok {
			len++
		}
	}

	tst := make([]IEvent_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEvent_expressionContext); ok {
			tst[i] = t.(IEvent_expressionContext)
			i++
		}
	}

	return tst
}

func (s *EventContext) Event_expression(i int) IEvent_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEvent_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEvent_expressionContext)
}

func (s *EventContext) Event_qty() IEvent_qtyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEvent_qtyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEvent_qtyContext)
}

func (s *EventContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserAND)
}

func (s *EventContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserAND, i)
}

func (s *EventContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterEvent(s)
	}
}

func (s *EventContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitEvent(s)
	}
}

func (p *SESParserParser) Event() (localctx IEventContext) {
	this := p
	_ = this

	localctx = NewEventContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SESParserParserRULE_event)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(SESParserParserEVENT)
	}
	{
		p.SetState(67)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventContext).name = _m
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserL_CURLY_BRACKET)|(1<<SESParserParserPLUS)|(1<<SESParserParserASTERISK))) != 0 {
		{
			p.SetState(68)

			var _x = p.Event_qty()

			localctx.(*EventContext).qty = _x
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWHERE {
		{
			p.SetState(71)
			p.Match(SESParserParserWHERE)
		}
		{
			p.SetState(72)
			p.Event_expression()
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(73)
					p.Match(SESParserParserAND)
				}
				{
					p.SetState(74)
					p.Event_expression()
				}

			}
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IEvent_qtyContext is an interface to support dynamic dispatch.
type IEvent_qtyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEvent_qtyContext differentiates from other interfaces.
	IsEvent_qtyContext()
}

type Event_qtyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvent_qtyContext() *Event_qtyContext {
	var p = new(Event_qtyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_event_qty
	return p
}

func (*Event_qtyContext) IsEvent_qtyContext() {}

func NewEvent_qtyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Event_qtyContext {
	var p = new(Event_qtyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_event_qty

	return p
}

func (s *Event_qtyContext) GetParser() antlr.Parser { return s.parser }

func (s *Event_qtyContext) CopyFrom(ctx *Event_qtyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Event_qtyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Event_qtyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Qty_preciseContext struct {
	*Event_qtyContext
	from antlr.Token
	to   antlr.Token
}

func NewQty_preciseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Qty_preciseContext {
	var p = new(Qty_preciseContext)

	p.Event_qtyContext = NewEmptyEvent_qtyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Event_qtyContext))

	return p
}

func (s *Qty_preciseContext) GetFrom() antlr.Token { return s.from }

func (s *Qty_preciseContext) GetTo() antlr.Token { return s.to }

func (s *Qty_preciseContext) SetFrom(v antlr.Token) { s.from = v }

func (s *Qty_preciseContext) SetTo(v antlr.Token) { s.to = v }

func (s *Qty_preciseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qty_preciseContext) L_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_CURLY_BRACKET, 0)
}

func (s *Qty_preciseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SESParserParserCOMMA, 0)
}

func (s *Qty_preciseContext) R_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_CURLY_BRACKET, 0)
}

func (s *Qty_preciseContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserNUMBER)
}

func (s *Qty_preciseContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserNUMBER, i)
}

func (s *Qty_preciseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterQty_precise(s)
	}
}

func (s *Qty_preciseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitQty_precise(s)
	}
}

type Qty_plusContext struct {
	*Event_qtyContext
}

func NewQty_plusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Qty_plusContext {
	var p = new(Qty_plusContext)

	p.Event_qtyContext = NewEmptyEvent_qtyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Event_qtyContext))

	return p
}

func (s *Qty_plusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qty_plusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SESParserParserPLUS, 0)
}

func (s *Qty_plusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterQty_plus(s)
	}
}

func (s *Qty_plusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitQty_plus(s)
	}
}

type Qty_precise_altContext struct {
	*Event_qtyContext
	exact antlr.Token
}

func NewQty_precise_altContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Qty_precise_altContext {
	var p = new(Qty_precise_altContext)

	p.Event_qtyContext = NewEmptyEvent_qtyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Event_qtyContext))

	return p
}

func (s *Qty_precise_altContext) GetExact() antlr.Token { return s.exact }

func (s *Qty_precise_altContext) SetExact(v antlr.Token) { s.exact = v }

func (s *Qty_precise_altContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qty_precise_altContext) L_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_CURLY_BRACKET, 0)
}

func (s *Qty_precise_altContext) R_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_CURLY_BRACKET, 0)
}

func (s *Qty_precise_altContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SESParserParserNUMBER, 0)
}

func (s *Qty_precise_altContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterQty_precise_alt(s)
	}
}

func (s *Qty_precise_altContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitQty_precise_alt(s)
	}
}

type Qty_asteriscContext struct {
	*Event_qtyContext
}

func NewQty_asteriscContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Qty_asteriscContext {
	var p = new(Qty_asteriscContext)

	p.Event_qtyContext = NewEmptyEvent_qtyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Event_qtyContext))

	return p
}

func (s *Qty_asteriscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qty_asteriscContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SESParserParserASTERISK, 0)
}

func (s *Qty_asteriscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterQty_asterisc(s)
	}
}

func (s *Qty_asteriscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitQty_asterisc(s)
	}
}

func (p *SESParserParser) Event_qty() (localctx IEvent_qtyContext) {
	this := p
	_ = this

	localctx = NewEvent_qtyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SESParserParserRULE_event_qty)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQty_plusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(SESParserParserPLUS)
		}

	case 2:
		localctx = NewQty_asteriscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(SESParserParserASTERISK)
		}

	case 3:
		localctx = NewQty_preciseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(85)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).from = _m
			}

		}
		{
			p.SetState(88)
			p.Match(SESParserParserCOMMA)
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(89)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).to = _m
			}

		}
		{
			p.SetState(92)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	case 4:
		localctx = NewQty_precise_altContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		{
			p.SetState(94)

			var _m = p.Match(SESParserParserNUMBER)

			localctx.(*Qty_precise_altContext).exact = _m
		}
		{
			p.SetState(95)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	}

	return localctx
}

// ISet_windowContext is an interface to support dynamic dispatch.
type ISet_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSkip returns the skip rule contexts.
	GetSkip() IDateIntervalContext

	// GetWithin returns the within rule contexts.
	GetWithin() IDateIntervalContext

	// SetSkip sets the skip rule contexts.
	SetSkip(IDateIntervalContext)

	// SetWithin sets the within rule contexts.
	SetWithin(IDateIntervalContext)

	// IsSet_windowContext differentiates from other interfaces.
	IsSet_windowContext()
}

type Set_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	skip   IDateIntervalContext
	within IDateIntervalContext
}

func NewEmptySet_windowContext() *Set_windowContext {
	var p = new(Set_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_set_window
	return p
}

func (*Set_windowContext) IsSet_windowContext() {}

func NewSet_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_windowContext {
	var p = new(Set_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_set_window

	return p
}

func (s *Set_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_windowContext) GetSkip() IDateIntervalContext { return s.skip }

func (s *Set_windowContext) GetWithin() IDateIntervalContext { return s.within }

func (s *Set_windowContext) SetSkip(v IDateIntervalContext) { s.skip = v }

func (s *Set_windowContext) SetWithin(v IDateIntervalContext) { s.within = v }

func (s *Set_windowContext) THEN() antlr.TerminalNode {
	return s.GetToken(SESParserParserTHEN, 0)
}

func (s *Set_windowContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserAND)
}

func (s *Set_windowContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserAND, i)
}

func (s *Set_windowContext) SKIP_() antlr.TerminalNode {
	return s.GetToken(SESParserParserSKIP_, 0)
}

func (s *Set_windowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
}

func (s *Set_windowContext) AllDateInterval() []IDateIntervalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateIntervalContext); ok {
			len++
		}
	}

	tst := make([]IDateIntervalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateIntervalContext); ok {
			tst[i] = t.(IDateIntervalContext)
			i++
		}
	}

	return tst
}

func (s *Set_windowContext) DateInterval(i int) IDateIntervalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalContext)
}

func (s *Set_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterSet_window(s)
	}
}

func (s *Set_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitSet_window(s)
	}
}

func (p *SESParserParser) Set_window() (localctx ISet_windowContext) {
	this := p
	_ = this

	localctx = NewSet_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SESParserParserRULE_set_window)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserAND {
		{
			p.SetState(98)
			p.Match(SESParserParserAND)
		}

	}
	{
		p.SetState(101)
		p.Match(SESParserParserTHEN)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserSKIP_ {
		{
			p.SetState(102)
			p.Match(SESParserParserSKIP_)
		}
		{
			p.SetState(103)

			var _x = p.DateInterval()

			localctx.(*Set_windowContext).skip = _x
		}

	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWITHIN || _la == SESParserParserAND {
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(106)
				p.Match(SESParserParserAND)
			}

		}
		{
			p.SetState(109)
			p.Match(SESParserParserWITHIN)
		}
		{
			p.SetState(110)

			var _x = p.DateInterval()

			localctx.(*Set_windowContext).within = _x
		}

	}

	return localctx
}

// IDateIntervalContext is an interface to support dynamic dispatch.
type IDateIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// GetUnit returns the unit token.
	GetUnit() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// SetUnit sets the unit token.
	SetUnit(antlr.Token)

	// GetExtra returns the extra rule contexts.
	GetExtra() IDateIntervalContext

	// SetExtra sets the extra rule contexts.
	SetExtra(IDateIntervalContext)

	// IsDateIntervalContext differentiates from other interfaces.
	IsDateIntervalContext()
}

type DateIntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	num    antlr.Token
	unit   antlr.Token
	extra  IDateIntervalContext
}

func NewEmptyDateIntervalContext() *DateIntervalContext {
	var p = new(DateIntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_dateInterval
	return p
}

func (*DateIntervalContext) IsDateIntervalContext() {}

func NewDateIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateIntervalContext {
	var p = new(DateIntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_dateInterval

	return p
}

func (s *DateIntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *DateIntervalContext) GetNum() antlr.Token { return s.num }

func (s *DateIntervalContext) GetUnit() antlr.Token { return s.unit }

func (s *DateIntervalContext) SetNum(v antlr.Token) { s.num = v }

func (s *DateIntervalContext) SetUnit(v antlr.Token) { s.unit = v }

func (s *DateIntervalContext) GetExtra() IDateIntervalContext { return s.extra }

func (s *DateIntervalContext) SetExtra(v IDateIntervalContext) { s.extra = v }

func (s *DateIntervalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SESParserParserNUMBER, 0)
}

func (s *DateIntervalContext) DATE_UNIT() antlr.TerminalNode {
	return s.GetToken(SESParserParserDATE_UNIT, 0)
}

func (s *DateIntervalContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserAND)
}

func (s *DateIntervalContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserAND, i)
}

func (s *DateIntervalContext) AllDateInterval() []IDateIntervalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDateIntervalContext); ok {
			len++
		}
	}

	tst := make([]IDateIntervalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDateIntervalContext); ok {
			tst[i] = t.(IDateIntervalContext)
			i++
		}
	}

	return tst
}

func (s *DateIntervalContext) DateInterval(i int) IDateIntervalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalContext)
}

func (s *DateIntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateIntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateIntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterDateInterval(s)
	}
}

func (s *DateIntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitDateInterval(s)
	}
}

func (p *SESParserParser) DateInterval() (localctx IDateIntervalContext) {
	this := p
	_ = this

	localctx = NewDateIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SESParserParserRULE_dateInterval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)

		var _m = p.Match(SESParserParserNUMBER)

		localctx.(*DateIntervalContext).num = _m
	}
	{
		p.SetState(114)

		var _m = p.Match(SESParserParserDATE_UNIT)

		localctx.(*DateIntervalContext).unit = _m
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(115)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(116)

				var _x = p.DateInterval()

				localctx.(*DateIntervalContext).extra = _x
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) AbsoluteDate() IAbsoluteDateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsoluteDateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsoluteDateContext)
}

func (s *DateContext) RelativeDate() IRelativeDateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelativeDateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelativeDateContext)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *SESParserParser) Date() (localctx IDateContext) {
	this := p
	_ = this

	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SESParserParserRULE_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.AbsoluteDate()
		}

	case SESParserParserLAST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.RelativeDate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAbsoluteDateContext is an interface to support dynamic dispatch.
type IAbsoluteDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsoluteDateContext differentiates from other interfaces.
	IsAbsoluteDateContext()
}

type AbsoluteDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsoluteDateContext() *AbsoluteDateContext {
	var p = new(AbsoluteDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_absoluteDate
	return p
}

func (*AbsoluteDateContext) IsAbsoluteDateContext() {}

func NewAbsoluteDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsoluteDateContext {
	var p = new(AbsoluteDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_absoluteDate

	return p
}

func (s *AbsoluteDateContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsoluteDateContext) STRING() antlr.TerminalNode {
	return s.GetToken(SESParserParserSTRING, 0)
}

func (s *AbsoluteDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsoluteDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsoluteDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterAbsoluteDate(s)
	}
}

func (s *AbsoluteDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitAbsoluteDate(s)
	}
}

func (p *SESParserParser) AbsoluteDate() (localctx IAbsoluteDateContext) {
	this := p
	_ = this

	localctx = NewAbsoluteDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SESParserParserRULE_absoluteDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(SESParserParserSTRING)
	}

	return localctx
}

// IRelativeDateContext is an interface to support dynamic dispatch.
type IRelativeDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLast returns the last rule contexts.
	GetLast() IDateIntervalContext

	// SetLast sets the last rule contexts.
	SetLast(IDateIntervalContext)

	// IsRelativeDateContext differentiates from other interfaces.
	IsRelativeDateContext()
}

type RelativeDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	last   IDateIntervalContext
}

func NewEmptyRelativeDateContext() *RelativeDateContext {
	var p = new(RelativeDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_relativeDate
	return p
}

func (*RelativeDateContext) IsRelativeDateContext() {}

func NewRelativeDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelativeDateContext {
	var p = new(RelativeDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_relativeDate

	return p
}

func (s *RelativeDateContext) GetParser() antlr.Parser { return s.parser }

func (s *RelativeDateContext) GetLast() IDateIntervalContext { return s.last }

func (s *RelativeDateContext) SetLast(v IDateIntervalContext) { s.last = v }

func (s *RelativeDateContext) LAST() antlr.TerminalNode {
	return s.GetToken(SESParserParserLAST, 0)
}

func (s *RelativeDateContext) DateInterval() IDateIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateIntervalContext)
}

func (s *RelativeDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelativeDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterRelativeDate(s)
	}
}

func (s *RelativeDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitRelativeDate(s)
	}
}

func (p *SESParserParser) RelativeDate() (localctx IRelativeDateContext) {
	this := p
	_ = this

	localctx = NewRelativeDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SESParserParserRULE_relativeDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(SESParserParserLAST)
	}
	{
		p.SetState(129)

		var _x = p.DateInterval()

		localctx.(*RelativeDateContext).last = _x
	}

	return localctx
}

// IEvent_expressionContext is an interface to support dynamic dispatch.
type IEvent_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpr_operandContext

	// GetRight returns the right rule contexts.
	GetRight() IExpr_operandContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpr_operandContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpr_operandContext)

	// IsEvent_expressionContext differentiates from other interfaces.
	IsEvent_expressionContext()
}

type Event_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr_operandContext
	op     antlr.Token
	right  IExpr_operandContext
}

func NewEmptyEvent_expressionContext() *Event_expressionContext {
	var p = new(Event_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_event_expression
	return p
}

func (*Event_expressionContext) IsEvent_expressionContext() {}

func NewEvent_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Event_expressionContext {
	var p = new(Event_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_event_expression

	return p
}

func (s *Event_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Event_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Event_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Event_expressionContext) GetLeft() IExpr_operandContext { return s.left }

func (s *Event_expressionContext) GetRight() IExpr_operandContext { return s.right }

func (s *Event_expressionContext) SetLeft(v IExpr_operandContext) { s.left = v }

func (s *Event_expressionContext) SetRight(v IExpr_operandContext) { s.right = v }

func (s *Event_expressionContext) AllExpr_operand() []IExpr_operandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpr_operandContext); ok {
			len++
		}
	}

	tst := make([]IExpr_operandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpr_operandContext); ok {
			tst[i] = t.(IExpr_operandContext)
			i++
		}
	}

	return tst
}

func (s *Event_expressionContext) Expr_operand(i int) IExpr_operandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_operandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_operandContext)
}

func (s *Event_expressionContext) OP_LOGICAL() antlr.TerminalNode {
	return s.GetToken(SESParserParserOP_LOGICAL, 0)
}

func (s *Event_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Event_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Event_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterEvent_expression(s)
	}
}

func (s *Event_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitEvent_expression(s)
	}
}

func (p *SESParserParser) Event_expression() (localctx IEvent_expressionContext) {
	this := p
	_ = this

	localctx = NewEvent_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SESParserParserRULE_event_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)

		var _x = p.Expr_operand()

		localctx.(*Event_expressionContext).left = _x
	}
	{
		p.SetState(132)

		var _m = p.Match(SESParserParserOP_LOGICAL)

		localctx.(*Event_expressionContext).op = _m
	}
	{
		p.SetState(133)

		var _x = p.Expr_operand()

		localctx.(*Event_expressionContext).right = _x
	}

	return localctx
}

// IExpr_operandContext is an interface to support dynamic dispatch.
type IExpr_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr_operandContext differentiates from other interfaces.
	IsExpr_operandContext()
}

type Expr_operandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_operandContext() *Expr_operandContext {
	var p = new(Expr_operandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_expr_operand
	return p
}

func (*Expr_operandContext) IsExpr_operandContext() {}

func NewExpr_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_operandContext {
	var p = new(Expr_operandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_expr_operand

	return p
}

func (s *Expr_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_operandContext) CopyFrom(ctx *Expr_operandContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Expr_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttrModifiedContext struct {
	*Expr_operandContext
	modifier antlr.Token
}

func NewAttrModifiedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttrModifiedContext {
	var p = new(AttrModifiedContext)

	p.Expr_operandContext = NewEmptyExpr_operandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr_operandContext))

	return p
}

func (s *AttrModifiedContext) GetModifier() antlr.Token { return s.modifier }

func (s *AttrModifiedContext) SetModifier(v antlr.Token) { s.modifier = v }

func (s *AttrModifiedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrModifiedContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_BRACKET, 0)
}

func (s *AttrModifiedContext) EventAttr() IEventAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventAttrContext)
}

func (s *AttrModifiedContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_BRACKET, 0)
}

func (s *AttrModifiedContext) ID() antlr.TerminalNode {
	return s.GetToken(SESParserParserID, 0)
}

func (s *AttrModifiedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterAttrModified(s)
	}
}

func (s *AttrModifiedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitAttrModified(s)
	}
}

type LitContext struct {
	*Expr_operandContext
}

func NewLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitContext {
	var p = new(LitContext)

	p.Expr_operandContext = NewEmptyExpr_operandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr_operandContext))

	return p
}

func (s *LitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterLit(s)
	}
}

func (s *LitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitLit(s)
	}
}

type AttrContext struct {
	*Expr_operandContext
}

func NewAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttrContext {
	var p = new(AttrContext)

	p.Expr_operandContext = NewEmptyExpr_operandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Expr_operandContext))

	return p
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) EventAttr() IEventAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventAttrContext)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *SESParserParser) Expr_operand() (localctx IExpr_operandContext) {
	this := p
	_ = this

	localctx = NewExpr_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SESParserParserRULE_expr_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttrModifiedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			var _m = p.Match(SESParserParserID)

			localctx.(*AttrModifiedContext).modifier = _m
		}
		{
			p.SetState(136)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(137)
			p.EventAttr()
		}
		{
			p.SetState(138)
			p.Match(SESParserParserR_BRACKET)
		}

	case 2:
		localctx = NewAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.EventAttr()
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Literal()
		}

	}

	return localctx
}

// IEventAttrContext is an interface to support dynamic dispatch.
type IEventAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEventName returns the eventName token.
	GetEventName() antlr.Token

	// GetAttrName returns the attrName token.
	GetAttrName() antlr.Token

	// SetEventName sets the eventName token.
	SetEventName(antlr.Token)

	// SetAttrName sets the attrName token.
	SetAttrName(antlr.Token)

	// IsEventAttrContext differentiates from other interfaces.
	IsEventAttrContext()
}

type EventAttrContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	eventName antlr.Token
	attrName  antlr.Token
}

func NewEmptyEventAttrContext() *EventAttrContext {
	var p = new(EventAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_eventAttr
	return p
}

func (*EventAttrContext) IsEventAttrContext() {}

func NewEventAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventAttrContext {
	var p = new(EventAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_eventAttr

	return p
}

func (s *EventAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *EventAttrContext) GetEventName() antlr.Token { return s.eventName }

func (s *EventAttrContext) GetAttrName() antlr.Token { return s.attrName }

func (s *EventAttrContext) SetEventName(v antlr.Token) { s.eventName = v }

func (s *EventAttrContext) SetAttrName(v antlr.Token) { s.attrName = v }

func (s *EventAttrContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserID)
}

func (s *EventAttrContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserID, i)
}

func (s *EventAttrContext) DOT() antlr.TerminalNode {
	return s.GetToken(SESParserParserDOT, 0)
}

func (s *EventAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterEventAttr(s)
	}
}

func (s *EventAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitEventAttr(s)
	}
}

func (p *SESParserParser) EventAttr() (localctx IEventAttrContext) {
	this := p
	_ = this

	localctx = NewEventAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SESParserParserRULE_eventAttr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(144)

			var _m = p.Match(SESParserParserID)

			localctx.(*EventAttrContext).eventName = _m
		}
		{
			p.SetState(145)
			p.Match(SESParserParserDOT)
		}

	}
	{
		p.SetState(148)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventAttrContext).attrName = _m
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Literal_numberContext struct {
	*LiteralContext
}

func NewLiteral_numberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Literal_numberContext {
	var p = new(Literal_numberContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *Literal_numberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_numberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SESParserParserNUMBER, 0)
}

func (s *Literal_numberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterLiteral_number(s)
	}
}

func (s *Literal_numberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitLiteral_number(s)
	}
}

type Literal_stringContext struct {
	*LiteralContext
}

func NewLiteral_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Literal_stringContext {
	var p = new(Literal_stringContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *Literal_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_stringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SESParserParserSTRING, 0)
}

func (s *Literal_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterLiteral_string(s)
	}
}

func (s *Literal_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitLiteral_string(s)
	}
}

func (p *SESParserParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SESParserParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER:
		localctx = NewLiteral_numberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(SESParserParserNUMBER)
		}

	case SESParserParserSTRING:
		localctx = NewLiteral_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(SESParserParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SESParserParserGROUP, 0)
}

func (s *GroupContext) EventAttr() IEventAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEventAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEventAttrContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *SESParserParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SESParserParserRULE_group)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(SESParserParserGROUP)
	}
	{
		p.SetState(155)
		p.EventAttr()
	}

	return localctx
}
