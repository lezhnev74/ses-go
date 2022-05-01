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
		"R_BRACKET", "L_SQUARE_BRACKET", "R_SQUARE_BRACKET", "COMMA", "SEMI",
		"PLUS", "ASTERISK", "NUMBER", "STRING", "ID",
	}
	staticData.ruleNames = []string{
		"parse", "ses", "event", "event_qty", "ses_window", "dateInterval",
		"event_expression", "expr_operand", "eventAttr", "literal", "group",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 121, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0, 30, 9, 0, 1,
		0, 3, 0, 33, 8, 0, 1, 0, 1, 0, 1, 1, 4, 1, 38, 8, 1, 11, 1, 12, 1, 39,
		1, 2, 1, 2, 1, 2, 3, 2, 45, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 51, 8,
		2, 10, 2, 12, 2, 54, 9, 2, 3, 2, 56, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		62, 8, 3, 1, 3, 1, 3, 3, 3, 66, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 72,
		8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 77, 8, 4, 1, 4, 3, 4, 80, 8, 4, 1, 4, 1,
		4, 3, 4, 84, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 90, 8, 5, 10, 5, 12, 5,
		93, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 106, 8, 7, 1, 8, 1, 8, 3, 8, 110, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		3, 9, 116, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 0, 0, 128, 0, 22, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4,
		41, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 85, 1, 0, 0,
		0, 12, 94, 1, 0, 0, 0, 14, 105, 1, 0, 0, 0, 16, 109, 1, 0, 0, 0, 18, 115,
		1, 0, 0, 0, 20, 117, 1, 0, 0, 0, 22, 28, 3, 2, 1, 0, 23, 24, 3, 8, 4, 0,
		24, 25, 3, 2, 1, 0, 25, 27, 1, 0, 0, 0, 26, 23, 1, 0, 0, 0, 27, 30, 1,
		0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30,
		28, 1, 0, 0, 0, 31, 33, 3, 20, 10, 0, 32, 31, 1, 0, 0, 0, 32, 33, 1, 0,
		0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 5, 0, 0, 1, 35, 1, 1, 0, 0, 0, 36, 38,
		3, 4, 2, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		39, 40, 1, 0, 0, 0, 40, 3, 1, 0, 0, 0, 41, 42, 5, 4, 0, 0, 42, 44, 5, 25,
		0, 0, 43, 45, 3, 6, 3, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 55,
		1, 0, 0, 0, 46, 47, 5, 8, 0, 0, 47, 52, 3, 12, 6, 0, 48, 49, 5, 9, 0, 0,
		49, 51, 3, 12, 6, 0, 50, 48, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1,
		0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55,
		46, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 5, 1, 0, 0, 0, 57, 72, 5, 21, 0,
		0, 58, 72, 5, 22, 0, 0, 59, 61, 5, 17, 0, 0, 60, 62, 5, 23, 0, 0, 61, 60,
		1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 5, 19, 0, 0,
		64, 66, 5, 23, 0, 0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1,
		0, 0, 0, 67, 72, 5, 18, 0, 0, 68, 69, 5, 17, 0, 0, 69, 70, 5, 23, 0, 0,
		70, 72, 5, 18, 0, 0, 71, 57, 1, 0, 0, 0, 71, 58, 1, 0, 0, 0, 71, 59, 1,
		0, 0, 0, 71, 68, 1, 0, 0, 0, 72, 7, 1, 0, 0, 0, 73, 76, 5, 7, 0, 0, 74,
		75, 5, 5, 0, 0, 75, 77, 3, 10, 5, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0,
		0, 0, 77, 83, 1, 0, 0, 0, 78, 80, 5, 9, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 5, 6, 0, 0, 82, 84, 3, 10, 5, 0,
		83, 79, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 9, 1, 0, 0, 0, 85, 86, 5, 23,
		0, 0, 86, 91, 5, 12, 0, 0, 87, 88, 5, 9, 0, 0, 88, 90, 3, 10, 5, 0, 89,
		87, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0,
		0, 92, 11, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 3, 14, 7, 0, 95, 96,
		5, 13, 0, 0, 96, 97, 3, 14, 7, 0, 97, 13, 1, 0, 0, 0, 98, 99, 5, 25, 0,
		0, 99, 100, 5, 15, 0, 0, 100, 101, 3, 16, 8, 0, 101, 102, 5, 16, 0, 0,
		102, 106, 1, 0, 0, 0, 103, 106, 3, 16, 8, 0, 104, 106, 3, 18, 9, 0, 105,
		98, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 15, 1,
		0, 0, 0, 107, 108, 5, 25, 0, 0, 108, 110, 5, 14, 0, 0, 109, 107, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 25, 0, 0,
		112, 17, 1, 0, 0, 0, 113, 116, 5, 23, 0, 0, 114, 116, 5, 24, 0, 0, 115,
		113, 1, 0, 0, 0, 115, 114, 1, 0, 0, 0, 116, 19, 1, 0, 0, 0, 117, 118, 5,
		11, 0, 0, 118, 119, 3, 16, 8, 0, 119, 21, 1, 0, 0, 0, 16, 28, 32, 39, 44,
		52, 55, 61, 65, 71, 76, 79, 83, 91, 105, 109, 115,
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
	SESParserParserEOF              = antlr.TokenEOF
	SESParserParserWS               = 1
	SESParserParserLINE_COMMENT     = 2
	SESParserParserSPACE            = 3
	SESParserParserEVENT            = 4
	SESParserParserSKIP_            = 5
	SESParserParserWITHIN           = 6
	SESParserParserTHEN             = 7
	SESParserParserWHERE            = 8
	SESParserParserAND              = 9
	SESParserParserOR               = 10
	SESParserParserGROUP            = 11
	SESParserParserDATE_UNIT        = 12
	SESParserParserOP_LOGICAL       = 13
	SESParserParserDOT              = 14
	SESParserParserL_BRACKET        = 15
	SESParserParserR_BRACKET        = 16
	SESParserParserL_SQUARE_BRACKET = 17
	SESParserParserR_SQUARE_BRACKET = 18
	SESParserParserCOMMA            = 19
	SESParserParserSEMI             = 20
	SESParserParserPLUS             = 21
	SESParserParserASTERISK         = 22
	SESParserParserNUMBER           = 23
	SESParserParserSTRING           = 24
	SESParserParserID               = 25
)

// SESParserParser rules.
const (
	SESParserParserRULE_parse            = 0
	SESParserParserRULE_ses              = 1
	SESParserParserRULE_event            = 2
	SESParserParserRULE_event_qty        = 3
	SESParserParserRULE_ses_window       = 4
	SESParserParserRULE_dateInterval     = 5
	SESParserParserRULE_event_expression = 6
	SESParserParserRULE_expr_operand     = 7
	SESParserParserRULE_eventAttr        = 8
	SESParserParserRULE_literal          = 9
	SESParserParserRULE_group            = 10
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

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(SESParserParserEOF, 0)
}

func (s *ParseContext) AllSes_window() []ISes_windowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISes_windowContext); ok {
			len++
		}
	}

	tst := make([]ISes_windowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISes_windowContext); ok {
			tst[i] = t.(ISes_windowContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Ses_window(i int) ISes_windowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISes_windowContext); ok {
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

	return t.(ISes_windowContext)
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
	{
		p.SetState(22)
		p.Ses()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SESParserParserTHEN {
		{
			p.SetState(23)
			p.Ses_window()
		}
		{
			p.SetState(24)
			p.Ses()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserGROUP {
		{
			p.SetState(31)
			p.Group()
		}

	}
	{
		p.SetState(34)
		p.Match(SESParserParserEOF)
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
	p.EnterRule(localctx, 2, SESParserParserRULE_ses)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SESParserParserEVENT {
		{
			p.SetState(36)
			p.Event()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 4, SESParserParserRULE_event)
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
		p.SetState(41)
		p.Match(SESParserParserEVENT)
	}
	{
		p.SetState(42)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventContext).name = _m
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserL_SQUARE_BRACKET)|(1<<SESParserParserPLUS)|(1<<SESParserParserASTERISK))) != 0 {
		{
			p.SetState(43)

			var _x = p.Event_qty()

			localctx.(*EventContext).qty = _x
		}

	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWHERE {
		{
			p.SetState(46)
			p.Match(SESParserParserWHERE)
		}
		{
			p.SetState(47)
			p.Event_expression()
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SESParserParserAND {
			{
				p.SetState(48)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(49)
				p.Event_expression()
			}

			p.SetState(54)
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

func (s *Qty_preciseContext) L_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_SQUARE_BRACKET, 0)
}

func (s *Qty_preciseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SESParserParserCOMMA, 0)
}

func (s *Qty_preciseContext) R_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_SQUARE_BRACKET, 0)
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

func (s *Qty_precise_altContext) L_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_SQUARE_BRACKET, 0)
}

func (s *Qty_precise_altContext) R_SQUARE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_SQUARE_BRACKET, 0)
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
	p.EnterRule(localctx, 6, SESParserParserRULE_event_qty)
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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQty_plusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(SESParserParserPLUS)
		}

	case 2:
		localctx = NewQty_asteriscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(SESParserParserASTERISK)
		}

	case 3:
		localctx = NewQty_preciseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Match(SESParserParserL_SQUARE_BRACKET)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(60)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).from = _m
			}

		}
		{
			p.SetState(63)
			p.Match(SESParserParserCOMMA)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(64)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).to = _m
			}

		}
		{
			p.SetState(67)
			p.Match(SESParserParserR_SQUARE_BRACKET)
		}

	case 4:
		localctx = NewQty_precise_altContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(SESParserParserL_SQUARE_BRACKET)
		}
		{
			p.SetState(69)

			var _m = p.Match(SESParserParserNUMBER)

			localctx.(*Qty_precise_altContext).exact = _m
		}
		{
			p.SetState(70)
			p.Match(SESParserParserR_SQUARE_BRACKET)
		}

	}

	return localctx
}

// ISes_windowContext is an interface to support dynamic dispatch.
type ISes_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSes_windowContext differentiates from other interfaces.
	IsSes_windowContext()
}

type Ses_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Ses_windowContext) THEN() antlr.TerminalNode {
	return s.GetToken(SESParserParserTHEN, 0)
}

func (s *Ses_windowContext) SKIP_() antlr.TerminalNode {
	return s.GetToken(SESParserParserSKIP_, 0)
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

func (s *Ses_windowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
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
	p.EnterRule(localctx, 8, SESParserParserRULE_ses_window)
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
		p.SetState(73)
		p.Match(SESParserParserTHEN)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserSKIP_ {
		{
			p.SetState(74)
			p.Match(SESParserParserSKIP_)
		}
		{
			p.SetState(75)
			p.DateInterval()
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWITHIN || _la == SESParserParserAND {
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(78)
				p.Match(SESParserParserAND)
			}

		}
		{
			p.SetState(81)
			p.Match(SESParserParserWITHIN)
		}
		{
			p.SetState(82)
			p.DateInterval()
		}

	}

	return localctx
}

// IDateIntervalContext is an interface to support dynamic dispatch.
type IDateIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateIntervalContext differentiates from other interfaces.
	IsDateIntervalContext()
}

type DateIntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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
	p.EnterRule(localctx, 10, SESParserParserRULE_dateInterval)

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
		p.SetState(85)
		p.Match(SESParserParserNUMBER)
	}
	{
		p.SetState(86)
		p.Match(SESParserParserDATE_UNIT)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(87)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(88)
				p.DateInterval()
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, SESParserParserRULE_event_expression)

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
		p.SetState(94)

		var _x = p.Expr_operand()

		localctx.(*Event_expressionContext).left = _x
	}
	{
		p.SetState(95)

		var _m = p.Match(SESParserParserOP_LOGICAL)

		localctx.(*Event_expressionContext).op = _m
	}
	{
		p.SetState(96)

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
	p.EnterRule(localctx, 14, SESParserParserRULE_expr_operand)

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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttrModifiedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)

			var _m = p.Match(SESParserParserID)

			localctx.(*AttrModifiedContext).modifier = _m
		}
		{
			p.SetState(99)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(100)
			p.EventAttr()
		}
		{
			p.SetState(101)
			p.Match(SESParserParserR_BRACKET)
		}

	case 2:
		localctx = NewAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.EventAttr()
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
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
	p.EnterRule(localctx, 16, SESParserParserRULE_eventAttr)

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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)

			var _m = p.Match(SESParserParserID)

			localctx.(*EventAttrContext).eventName = _m
		}
		{
			p.SetState(108)
			p.Match(SESParserParserDOT)
		}

	}
	{
		p.SetState(111)

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
	p.EnterRule(localctx, 18, SESParserParserRULE_literal)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER:
		localctx = NewLiteral_numberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(SESParserParserNUMBER)
		}

	case SESParserParserSTRING:
		localctx = NewLiteral_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
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
	p.EnterRule(localctx, 20, SESParserParserRULE_group)

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
		p.SetState(117)
		p.Match(SESParserParserGROUP)
	}
	{
		p.SetState(118)
		p.EventAttr()
	}

	return localctx
}
