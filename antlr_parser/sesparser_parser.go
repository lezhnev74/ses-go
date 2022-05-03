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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'.'", "'('",
		"')'", "'{'", "'}'", "','", "';'", "'+'", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "LINE_COMMENT", "SPACE", "EVENT", "SKIP_", "WITHIN", "THEN",
		"WHERE", "AND", "OR", "GROUP", "DATE_UNIT", "OP_LOGICAL", "DOT", "L_BRACKET",
		"R_BRACKET", "L_CURLY_BRACKET", "R_CURLY_BRACKET", "COMMA", "SEMI",
		"PLUS", "ASTERISK", "NUMBER", "STRING", "ID",
	}
	staticData.ruleNames = []string{
		"parse", "window", "ses", "windowed_ses", "event", "event_qty", "ses_window",
		"dateInterval", "event_expression", "expr_operand", "eventAttr", "literal",
		"group",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 137, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 3, 0, 28, 8, 0, 1, 0, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 2, 4, 2, 46, 8, 2, 11, 2, 12, 2, 47, 1, 3, 3, 3, 51,
		8, 3, 1, 3, 4, 3, 54, 8, 3, 11, 3, 12, 3, 55, 1, 4, 1, 4, 1, 4, 3, 4, 61,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 67, 8, 4, 10, 4, 12, 4, 70, 9, 4, 3,
		4, 72, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 78, 8, 5, 1, 5, 1, 5, 3, 5,
		82, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 6, 1, 6, 1, 6, 3,
		6, 93, 8, 6, 1, 6, 3, 6, 96, 8, 6, 1, 6, 1, 6, 3, 6, 100, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 106, 8, 7, 10, 7, 12, 7, 109, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 122, 8, 9, 1,
		10, 1, 10, 3, 10, 126, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 132, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 0, 0, 145, 0, 27, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 45,
		1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 87, 1, 0, 0, 0, 12,
		89, 1, 0, 0, 0, 14, 101, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18, 121, 1, 0,
		0, 0, 20, 125, 1, 0, 0, 0, 22, 131, 1, 0, 0, 0, 24, 133, 1, 0, 0, 0, 26,
		28, 3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0,
		0, 29, 33, 3, 4, 2, 0, 30, 32, 3, 6, 3, 0, 31, 30, 1, 0, 0, 0, 32, 35,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 36, 38, 3, 24, 12, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1,
		0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 5, 0, 0, 1, 40, 1, 1, 0, 0, 0, 41,
		42, 5, 6, 0, 0, 42, 43, 3, 14, 7, 0, 43, 3, 1, 0, 0, 0, 44, 46, 3, 8, 4,
		0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 51, 3, 12, 6, 0, 50, 49, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 53, 1, 0, 0, 0, 52, 54, 3, 8, 4, 0, 53, 52, 1,
		0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		7, 1, 0, 0, 0, 57, 58, 5, 4, 0, 0, 58, 60, 5, 25, 0, 0, 59, 61, 3, 10,
		5, 0, 60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 71, 1, 0, 0, 0, 62, 63,
		5, 8, 0, 0, 63, 68, 3, 16, 8, 0, 64, 65, 5, 9, 0, 0, 65, 67, 3, 16, 8,
		0, 66, 64, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 62, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 9, 1, 0, 0, 0, 73, 88, 5, 21, 0, 0, 74, 88, 5,
		22, 0, 0, 75, 77, 5, 17, 0, 0, 76, 78, 5, 23, 0, 0, 77, 76, 1, 0, 0, 0,
		77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 5, 19, 0, 0, 80, 82, 5,
		23, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83,
		88, 5, 18, 0, 0, 84, 85, 5, 17, 0, 0, 85, 86, 5, 23, 0, 0, 86, 88, 5, 18,
		0, 0, 87, 73, 1, 0, 0, 0, 87, 74, 1, 0, 0, 0, 87, 75, 1, 0, 0, 0, 87, 84,
		1, 0, 0, 0, 88, 11, 1, 0, 0, 0, 89, 92, 5, 7, 0, 0, 90, 91, 5, 5, 0, 0,
		91, 93, 3, 14, 7, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 99, 1,
		0, 0, 0, 94, 96, 5, 9, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		97, 1, 0, 0, 0, 97, 98, 5, 6, 0, 0, 98, 100, 3, 14, 7, 0, 99, 95, 1, 0,
		0, 0, 99, 100, 1, 0, 0, 0, 100, 13, 1, 0, 0, 0, 101, 102, 5, 23, 0, 0,
		102, 107, 5, 12, 0, 0, 103, 104, 5, 9, 0, 0, 104, 106, 3, 14, 7, 0, 105,
		103, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108,
		1, 0, 0, 0, 108, 15, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 3, 18,
		9, 0, 111, 112, 5, 13, 0, 0, 112, 113, 3, 18, 9, 0, 113, 17, 1, 0, 0, 0,
		114, 115, 5, 25, 0, 0, 115, 116, 5, 15, 0, 0, 116, 117, 3, 20, 10, 0, 117,
		118, 5, 16, 0, 0, 118, 122, 1, 0, 0, 0, 119, 122, 3, 20, 10, 0, 120, 122,
		3, 22, 11, 0, 121, 114, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1,
		0, 0, 0, 122, 19, 1, 0, 0, 0, 123, 124, 5, 25, 0, 0, 124, 126, 5, 14, 0,
		0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127,
		128, 5, 25, 0, 0, 128, 21, 1, 0, 0, 0, 129, 132, 5, 23, 0, 0, 130, 132,
		5, 24, 0, 0, 131, 129, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 23, 1, 0,
		0, 0, 133, 134, 5, 11, 0, 0, 134, 135, 3, 20, 10, 0, 135, 25, 1, 0, 0,
		0, 19, 27, 33, 37, 47, 50, 55, 60, 68, 71, 77, 81, 87, 92, 95, 99, 107,
		121, 125, 131,
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
	SESParserParserSKIP_           = 5
	SESParserParserWITHIN          = 6
	SESParserParserTHEN            = 7
	SESParserParserWHERE           = 8
	SESParserParserAND             = 9
	SESParserParserOR              = 10
	SESParserParserGROUP           = 11
	SESParserParserDATE_UNIT       = 12
	SESParserParserOP_LOGICAL      = 13
	SESParserParserDOT             = 14
	SESParserParserL_BRACKET       = 15
	SESParserParserR_BRACKET       = 16
	SESParserParserL_CURLY_BRACKET = 17
	SESParserParserR_CURLY_BRACKET = 18
	SESParserParserCOMMA           = 19
	SESParserParserSEMI            = 20
	SESParserParserPLUS            = 21
	SESParserParserASTERISK        = 22
	SESParserParserNUMBER          = 23
	SESParserParserSTRING          = 24
	SESParserParserID              = 25
)

// SESParserParser rules.
const (
	SESParserParserRULE_parse            = 0
	SESParserParserRULE_window           = 1
	SESParserParserRULE_ses              = 2
	SESParserParserRULE_windowed_ses     = 3
	SESParserParserRULE_event            = 4
	SESParserParserRULE_event_qty        = 5
	SESParserParserRULE_ses_window       = 6
	SESParserParserRULE_dateInterval     = 7
	SESParserParserRULE_event_expression = 8
	SESParserParserRULE_expr_operand     = 9
	SESParserParserRULE_eventAttr        = 10
	SESParserParserRULE_literal          = 11
	SESParserParserRULE_group            = 12
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

func (s *ParseContext) Ses() ISesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISesContext)
}

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

func (s *ParseContext) AllWindowed_ses() []IWindowed_sesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWindowed_sesContext); ok {
			len++
		}
	}

	tst := make([]IWindowed_sesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWindowed_sesContext); ok {
			tst[i] = t.(IWindowed_sesContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Windowed_ses(i int) IWindowed_sesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowed_sesContext); ok {
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

	return t.(IWindowed_sesContext)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWITHIN {
		{
			p.SetState(26)
			p.Window()
		}

	}
	{
		p.SetState(29)
		p.Ses()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SESParserParserEVENT || _la == SESParserParserTHEN {
		{
			p.SetState(30)
			p.Windowed_ses()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserGROUP {
		{
			p.SetState(36)
			p.Group()
		}

	}
	{
		p.SetState(39)
		p.Match(SESParserParserEOF)
	}

	return localctx
}

// IWindowContext is an interface to support dynamic dispatch.
type IWindowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWithin returns the within rule contexts.
	GetWithin() IDateIntervalContext

	// SetWithin sets the within rule contexts.
	SetWithin(IDateIntervalContext)

	// IsWindowContext differentiates from other interfaces.
	IsWindowContext()
}

type WindowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *WindowContext) GetWithin() IDateIntervalContext { return s.within }

func (s *WindowContext) SetWithin(v IDateIntervalContext) { s.within = v }

func (s *WindowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
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
		p.SetState(41)
		p.Match(SESParserParserWITHIN)
	}
	{
		p.SetState(42)

		var _x = p.DateInterval()

		localctx.(*WindowContext).within = _x
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(44)
				p.Event()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IWindowed_sesContext is an interface to support dynamic dispatch.
type IWindowed_sesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowed_sesContext differentiates from other interfaces.
	IsWindowed_sesContext()
}

type Windowed_sesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowed_sesContext() *Windowed_sesContext {
	var p = new(Windowed_sesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_windowed_ses
	return p
}

func (*Windowed_sesContext) IsWindowed_sesContext() {}

func NewWindowed_sesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Windowed_sesContext {
	var p = new(Windowed_sesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_windowed_ses

	return p
}

func (s *Windowed_sesContext) GetParser() antlr.Parser { return s.parser }

func (s *Windowed_sesContext) Ses_window() ISes_windowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISes_windowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISes_windowContext)
}

func (s *Windowed_sesContext) AllEvent() []IEventContext {
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

func (s *Windowed_sesContext) Event(i int) IEventContext {
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

func (s *Windowed_sesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Windowed_sesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Windowed_sesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterWindowed_ses(s)
	}
}

func (s *Windowed_sesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitWindowed_ses(s)
	}
}

func (p *SESParserParser) Windowed_ses() (localctx IWindowed_sesContext) {
	this := p
	_ = this

	localctx = NewWindowed_sesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SESParserParserRULE_windowed_ses)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserTHEN {
		{
			p.SetState(49)
			p.Ses_window()
		}

	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(52)
				p.Event()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 8, SESParserParserRULE_event)
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
		p.SetState(57)
		p.Match(SESParserParserEVENT)
	}
	{
		p.SetState(58)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventContext).name = _m
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserL_CURLY_BRACKET)|(1<<SESParserParserPLUS)|(1<<SESParserParserASTERISK))) != 0 {
		{
			p.SetState(59)

			var _x = p.Event_qty()

			localctx.(*EventContext).qty = _x
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWHERE {
		{
			p.SetState(62)
			p.Match(SESParserParserWHERE)
		}
		{
			p.SetState(63)
			p.Event_expression()
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SESParserParserAND {
			{
				p.SetState(64)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(65)
				p.Event_expression()
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 10, SESParserParserRULE_event_qty)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQty_plusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(SESParserParserPLUS)
		}

	case 2:
		localctx = NewQty_asteriscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(SESParserParserASTERISK)
		}

	case 3:
		localctx = NewQty_preciseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(76)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).from = _m
			}

		}
		{
			p.SetState(79)
			p.Match(SESParserParserCOMMA)
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(80)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).to = _m
			}

		}
		{
			p.SetState(83)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	case 4:
		localctx = NewQty_precise_altContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(84)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		{
			p.SetState(85)

			var _m = p.Match(SESParserParserNUMBER)

			localctx.(*Qty_precise_altContext).exact = _m
		}
		{
			p.SetState(86)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	}

	return localctx
}

// ISes_windowContext is an interface to support dynamic dispatch.
type ISes_windowContext interface {
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

	// IsSes_windowContext differentiates from other interfaces.
	IsSes_windowContext()
}

type Ses_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	skip   IDateIntervalContext
	within IDateIntervalContext
}

func NewEmptySes_windowContext() *Ses_windowContext {
	var p = new(Ses_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_ses_window
	return p
}

func (*Ses_windowContext) IsSes_windowContext() {}

func NewSes_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ses_windowContext {
	var p = new(Ses_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_ses_window

	return p
}

func (s *Ses_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Ses_windowContext) GetSkip() IDateIntervalContext { return s.skip }

func (s *Ses_windowContext) GetWithin() IDateIntervalContext { return s.within }

func (s *Ses_windowContext) SetSkip(v IDateIntervalContext) { s.skip = v }

func (s *Ses_windowContext) SetWithin(v IDateIntervalContext) { s.within = v }

func (s *Ses_windowContext) THEN() antlr.TerminalNode {
	return s.GetToken(SESParserParserTHEN, 0)
}

func (s *Ses_windowContext) SKIP_() antlr.TerminalNode {
	return s.GetToken(SESParserParserSKIP_, 0)
}

func (s *Ses_windowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
}

func (s *Ses_windowContext) AllDateInterval() []IDateIntervalContext {
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

func (s *Ses_windowContext) DateInterval(i int) IDateIntervalContext {
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

func (s *Ses_windowContext) AND() antlr.TerminalNode {
	return s.GetToken(SESParserParserAND, 0)
}

func (s *Ses_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ses_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ses_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterSes_window(s)
	}
}

func (s *Ses_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitSes_window(s)
	}
}

func (p *SESParserParser) Ses_window() (localctx ISes_windowContext) {
	this := p
	_ = this

	localctx = NewSes_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SESParserParserRULE_ses_window)
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
		p.SetState(89)
		p.Match(SESParserParserTHEN)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserSKIP_ {
		{
			p.SetState(90)
			p.Match(SESParserParserSKIP_)
		}
		{
			p.SetState(91)

			var _x = p.DateInterval()

			localctx.(*Ses_windowContext).skip = _x
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWITHIN || _la == SESParserParserAND {
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(94)
				p.Match(SESParserParserAND)
			}

		}
		{
			p.SetState(97)
			p.Match(SESParserParserWITHIN)
		}
		{
			p.SetState(98)

			var _x = p.DateInterval()

			localctx.(*Ses_windowContext).within = _x
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
	p.EnterRule(localctx, 14, SESParserParserRULE_dateInterval)

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
		p.SetState(101)

		var _m = p.Match(SESParserParserNUMBER)

		localctx.(*DateIntervalContext).num = _m
	}
	{
		p.SetState(102)

		var _m = p.Match(SESParserParserDATE_UNIT)

		localctx.(*DateIntervalContext).unit = _m
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(103)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(104)

				var _x = p.DateInterval()

				localctx.(*DateIntervalContext).extra = _x
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, SESParserParserRULE_event_expression)

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
		p.SetState(110)

		var _x = p.Expr_operand()

		localctx.(*Event_expressionContext).left = _x
	}
	{
		p.SetState(111)

		var _m = p.Match(SESParserParserOP_LOGICAL)

		localctx.(*Event_expressionContext).op = _m
	}
	{
		p.SetState(112)

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
	p.EnterRule(localctx, 18, SESParserParserRULE_expr_operand)

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

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttrModifiedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

			var _m = p.Match(SESParserParserID)

			localctx.(*AttrModifiedContext).modifier = _m
		}
		{
			p.SetState(115)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(116)
			p.EventAttr()
		}
		{
			p.SetState(117)
			p.Match(SESParserParserR_BRACKET)
		}

	case 2:
		localctx = NewAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.EventAttr()
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
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
	p.EnterRule(localctx, 20, SESParserParserRULE_eventAttr)

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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(123)

			var _m = p.Match(SESParserParserID)

			localctx.(*EventAttrContext).eventName = _m
		}
		{
			p.SetState(124)
			p.Match(SESParserParserDOT)
		}

	}
	{
		p.SetState(127)

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
	p.EnterRule(localctx, 22, SESParserParserRULE_literal)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER:
		localctx = NewLiteral_numberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(SESParserParserNUMBER)
		}

	case SESParserParserSTRING:
		localctx = NewLiteral_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 24, SESParserParserRULE_group)

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
		p.SetState(133)
		p.Match(SESParserParserGROUP)
	}
	{
		p.SetState(134)
		p.EventAttr()
	}

	return localctx
}
