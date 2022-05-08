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
		"", "", "'.'", "'('", "')'", "'{'", "'}'", "','", "';'", "'+'", "'*'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "LINE_COMMENT", "SPACE", "WINDOW_TEXT", "EVENT", "WINDOW",
		"FROM", "TO", "LAST", "SKIP_", "WITHIN", "THEN", "WHERE", "AND", "OR",
		"GROUP", "DATE_UNIT", "OP_LOGICAL", "DOT", "L_BRACKET", "R_BRACKET",
		"L_CURLY_BRACKET", "R_CURLY_BRACKET", "COMMA", "SEMI", "PLUS", "ASTERISK",
		"NUMBER", "STRING", "ID",
	}
	staticData.ruleNames = []string{
		"parse", "ses_window", "ses", "event", "event_qty", "set_window", "date",
		"absoluteDate", "relativeDate", "dateInterval", "event_expression",
		"expr_operand", "eventAttr", "literal", "group", "groupAttr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 175, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 3, 0, 34, 8, 0, 1, 0, 4, 0, 37, 8, 0, 11, 0, 12, 0, 38, 1, 0, 3,
		0, 42, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 3, 2, 49, 8, 2, 1, 2, 4, 2,
		52, 8, 2, 11, 2, 12, 2, 53, 1, 3, 1, 3, 1, 3, 3, 3, 59, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 65, 8, 3, 10, 3, 12, 3, 68, 9, 3, 3, 3, 70, 8, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 76, 8, 4, 1, 4, 1, 4, 3, 4, 80, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 86, 8, 4, 1, 5, 3, 5, 89, 8, 5, 1, 5, 1, 5,
		3, 5, 93, 8, 5, 1, 5, 3, 5, 96, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 101, 8, 5,
		1, 5, 1, 5, 3, 5, 105, 8, 5, 1, 5, 3, 5, 108, 8, 5, 1, 5, 3, 5, 111, 8,
		5, 1, 5, 1, 5, 3, 5, 115, 8, 5, 1, 6, 1, 6, 3, 6, 119, 8, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 130, 8, 9, 10, 9, 12, 9,
		133, 9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 146, 8, 11, 1, 12, 1, 12, 3, 12, 150, 8, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 155, 8, 12, 10, 12, 12, 12, 158, 9, 12, 1, 13, 1,
		13, 3, 13, 162, 8, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15,
		170, 8, 15, 10, 15, 12, 15, 173, 9, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 0, 188, 0, 33, 1, 0, 0,
		0, 2, 45, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 55, 1, 0, 0, 0, 8, 85, 1, 0,
		0, 0, 10, 114, 1, 0, 0, 0, 12, 118, 1, 0, 0, 0, 14, 120, 1, 0, 0, 0, 16,
		122, 1, 0, 0, 0, 18, 125, 1, 0, 0, 0, 20, 134, 1, 0, 0, 0, 22, 145, 1,
		0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 161, 1, 0, 0, 0, 28, 163, 1, 0, 0, 0,
		30, 166, 1, 0, 0, 0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1,
		0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 37, 3, 4, 2, 0, 36, 35, 1, 0, 0, 0, 37,
		38, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 41, 1, 0, 0,
		0, 40, 42, 3, 28, 14, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43,
		1, 0, 0, 0, 43, 44, 5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 46, 5, 4, 0, 0,
		46, 3, 1, 0, 0, 0, 47, 49, 3, 10, 5, 0, 48, 47, 1, 0, 0, 0, 48, 49, 1,
		0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 52, 3, 6, 3, 0, 51, 50, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 5, 1, 0, 0,
		0, 55, 56, 5, 5, 0, 0, 56, 58, 5, 30, 0, 0, 57, 59, 3, 8, 4, 0, 58, 57,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 69, 1, 0, 0, 0, 60, 61, 5, 13, 0, 0,
		61, 66, 3, 20, 10, 0, 62, 63, 5, 14, 0, 0, 63, 65, 3, 20, 10, 0, 64, 62,
		1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0,
		67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 60, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 7, 1, 0, 0, 0, 71, 86, 5, 26, 0, 0, 72, 86, 5, 27, 0, 0, 73,
		75, 5, 22, 0, 0, 74, 76, 5, 28, 0, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0,
		0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 5, 24, 0, 0, 78, 80, 5, 28, 0, 0, 79,
		78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 86, 5, 23,
		0, 0, 82, 83, 5, 22, 0, 0, 83, 84, 5, 28, 0, 0, 84, 86, 5, 23, 0, 0, 85,
		71, 1, 0, 0, 0, 85, 72, 1, 0, 0, 0, 85, 73, 1, 0, 0, 0, 85, 82, 1, 0, 0,
		0, 86, 9, 1, 0, 0, 0, 87, 89, 5, 14, 0, 0, 88, 87, 1, 0, 0, 0, 88, 89,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 115, 5, 12, 0, 0, 91, 93, 5, 14, 0,
		0, 92, 91, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 96,
		5, 12, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0,
		97, 98, 5, 10, 0, 0, 98, 104, 3, 18, 9, 0, 99, 101, 5, 14, 0, 0, 100, 99,
		1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 5, 11,
		0, 0, 103, 105, 3, 18, 9, 0, 104, 100, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0,
		105, 115, 1, 0, 0, 0, 106, 108, 5, 14, 0, 0, 107, 106, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 111, 5, 12, 0, 0, 110, 109,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 5, 11,
		0, 0, 113, 115, 3, 18, 9, 0, 114, 88, 1, 0, 0, 0, 114, 92, 1, 0, 0, 0,
		114, 107, 1, 0, 0, 0, 115, 11, 1, 0, 0, 0, 116, 119, 3, 14, 7, 0, 117,
		119, 3, 16, 8, 0, 118, 116, 1, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 13,
		1, 0, 0, 0, 120, 121, 5, 29, 0, 0, 121, 15, 1, 0, 0, 0, 122, 123, 5, 9,
		0, 0, 123, 124, 3, 18, 9, 0, 124, 17, 1, 0, 0, 0, 125, 126, 5, 28, 0, 0,
		126, 131, 5, 17, 0, 0, 127, 128, 5, 14, 0, 0, 128, 130, 3, 18, 9, 0, 129,
		127, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 19, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 3, 22,
		11, 0, 135, 136, 5, 18, 0, 0, 136, 137, 3, 22, 11, 0, 137, 21, 1, 0, 0,
		0, 138, 139, 5, 30, 0, 0, 139, 140, 5, 20, 0, 0, 140, 141, 3, 24, 12, 0,
		141, 142, 5, 21, 0, 0, 142, 146, 1, 0, 0, 0, 143, 146, 3, 24, 12, 0, 144,
		146, 3, 26, 13, 0, 145, 138, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 144,
		1, 0, 0, 0, 146, 23, 1, 0, 0, 0, 147, 148, 5, 30, 0, 0, 148, 150, 5, 19,
		0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 156, 5, 30, 0, 0, 152, 153, 5, 19, 0, 0, 153, 155, 5, 30, 0, 0, 154,
		152, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 25, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 159, 162, 5, 28,
		0, 0, 160, 162, 5, 29, 0, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0,
		162, 27, 1, 0, 0, 0, 163, 164, 5, 16, 0, 0, 164, 165, 3, 30, 15, 0, 165,
		29, 1, 0, 0, 0, 166, 171, 3, 24, 12, 0, 167, 168, 5, 24, 0, 0, 168, 170,
		3, 30, 15, 0, 169, 167, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1,
		0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 31, 1, 0, 0, 0, 173, 171, 1, 0, 0,
		0, 26, 33, 38, 41, 48, 53, 58, 66, 69, 75, 79, 85, 88, 92, 95, 100, 104,
		107, 110, 114, 118, 131, 145, 149, 156, 161, 171,
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
	SESParserParserWINDOW_TEXT     = 4
	SESParserParserEVENT           = 5
	SESParserParserWINDOW          = 6
	SESParserParserFROM            = 7
	SESParserParserTO              = 8
	SESParserParserLAST            = 9
	SESParserParserSKIP_           = 10
	SESParserParserWITHIN          = 11
	SESParserParserTHEN            = 12
	SESParserParserWHERE           = 13
	SESParserParserAND             = 14
	SESParserParserOR              = 15
	SESParserParserGROUP           = 16
	SESParserParserDATE_UNIT       = 17
	SESParserParserOP_LOGICAL      = 18
	SESParserParserDOT             = 19
	SESParserParserL_BRACKET       = 20
	SESParserParserR_BRACKET       = 21
	SESParserParserL_CURLY_BRACKET = 22
	SESParserParserR_CURLY_BRACKET = 23
	SESParserParserCOMMA           = 24
	SESParserParserSEMI            = 25
	SESParserParserPLUS            = 26
	SESParserParserASTERISK        = 27
	SESParserParserNUMBER          = 28
	SESParserParserSTRING          = 29
	SESParserParserID              = 30
)

// SESParserParser rules.
const (
	SESParserParserRULE_parse            = 0
	SESParserParserRULE_ses_window       = 1
	SESParserParserRULE_ses              = 2
	SESParserParserRULE_event            = 3
	SESParserParserRULE_event_qty        = 4
	SESParserParserRULE_set_window       = 5
	SESParserParserRULE_date             = 6
	SESParserParserRULE_absoluteDate     = 7
	SESParserParserRULE_relativeDate     = 8
	SESParserParserRULE_dateInterval     = 9
	SESParserParserRULE_event_expression = 10
	SESParserParserRULE_expr_operand     = 11
	SESParserParserRULE_eventAttr        = 12
	SESParserParserRULE_literal          = 13
	SESParserParserRULE_group            = 14
	SESParserParserRULE_groupAttr        = 15
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

func (s *ParseContext) Ses_window() ISes_windowContext {
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWINDOW_TEXT {
		{
			p.SetState(32)
			p.Ses_window()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserEVENT)|(1<<SESParserParserSKIP_)|(1<<SESParserParserWITHIN)|(1<<SESParserParserTHEN)|(1<<SESParserParserAND))) != 0) {
		{
			p.SetState(35)
			p.Ses()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserGROUP {
		{
			p.SetState(40)
			p.Group()
		}

	}
	{
		p.SetState(43)
		p.Match(SESParserParserEOF)
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

func (s *Ses_windowContext) WINDOW_TEXT() antlr.TerminalNode {
	return s.GetToken(SESParserParserWINDOW_TEXT, 0)
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
	p.EnterRule(localctx, 2, SESParserParserRULE_ses_window)

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
		p.SetState(45)
		p.Match(SESParserParserWINDOW_TEXT)
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
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserSKIP_)|(1<<SESParserParserWITHIN)|(1<<SESParserParserTHEN)|(1<<SESParserParserAND))) != 0 {
		{
			p.SetState(47)
			p.Set_window()
		}

	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(50)
				p.Event()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
		p.SetState(55)
		p.Match(SESParserParserEVENT)
	}
	{
		p.SetState(56)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventContext).name = _m
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserL_CURLY_BRACKET)|(1<<SESParserParserPLUS)|(1<<SESParserParserASTERISK))) != 0 {
		{
			p.SetState(57)

			var _x = p.Event_qty()

			localctx.(*EventContext).qty = _x
		}

	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWHERE {
		{
			p.SetState(60)
			p.Match(SESParserParserWHERE)
		}
		{
			p.SetState(61)
			p.Event_expression()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(62)
					p.Match(SESParserParserAND)
				}
				{
					p.SetState(63)
					p.Event_expression()
				}

			}
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQty_plusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(SESParserParserPLUS)
		}

	case 2:
		localctx = NewQty_asteriscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(SESParserParserASTERISK)
		}

	case 3:
		localctx = NewQty_preciseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(74)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).from = _m
			}

		}
		{
			p.SetState(77)
			p.Match(SESParserParserCOMMA)
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(78)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).to = _m
			}

		}
		{
			p.SetState(81)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	case 4:
		localctx = NewQty_precise_altContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(82)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		{
			p.SetState(83)

			var _m = p.Match(SESParserParserNUMBER)

			localctx.(*Qty_precise_altContext).exact = _m
		}
		{
			p.SetState(84)
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

func (s *Set_windowContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(SESParserParserWITHIN, 0)
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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(87)
				p.Match(SESParserParserAND)
			}

		}
		{
			p.SetState(90)
			p.Match(SESParserParserTHEN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(91)
				p.Match(SESParserParserAND)
			}

		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserTHEN {
			{
				p.SetState(94)
				p.Match(SESParserParserTHEN)
			}

		}
		{
			p.SetState(97)
			p.Match(SESParserParserSKIP_)
		}
		{
			p.SetState(98)

			var _x = p.DateInterval()

			localctx.(*Set_windowContext).skip = _x
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserWITHIN || _la == SESParserParserAND {
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SESParserParserAND {
				{
					p.SetState(99)
					p.Match(SESParserParserAND)
				}

			}
			{
				p.SetState(102)
				p.Match(SESParserParserWITHIN)
			}
			{
				p.SetState(103)

				var _x = p.DateInterval()

				localctx.(*Set_windowContext).within = _x
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(106)
				p.Match(SESParserParserAND)
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserTHEN {
			{
				p.SetState(109)
				p.Match(SESParserParserTHEN)
			}

		}
		{
			p.SetState(112)
			p.Match(SESParserParserWITHIN)
		}
		{
			p.SetState(113)

			var _x = p.DateInterval()

			localctx.(*Set_windowContext).within = _x
		}

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
	p.EnterRule(localctx, 12, SESParserParserRULE_date)

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.AbsoluteDate()
		}

	case SESParserParserLAST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
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
	p.EnterRule(localctx, 14, SESParserParserRULE_absoluteDate)

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
		p.SetState(120)
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
	p.EnterRule(localctx, 16, SESParserParserRULE_relativeDate)

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
		p.SetState(122)
		p.Match(SESParserParserLAST)
	}
	{
		p.SetState(123)

		var _x = p.DateInterval()

		localctx.(*RelativeDateContext).last = _x
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
	p.EnterRule(localctx, 18, SESParserParserRULE_dateInterval)

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
		p.SetState(125)

		var _m = p.Match(SESParserParserNUMBER)

		localctx.(*DateIntervalContext).num = _m
	}
	{
		p.SetState(126)

		var _m = p.Match(SESParserParserDATE_UNIT)

		localctx.(*DateIntervalContext).unit = _m
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(127)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(128)

				var _x = p.DateInterval()

				localctx.(*DateIntervalContext).extra = _x
			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
		p.SetState(134)

		var _x = p.Expr_operand()

		localctx.(*Event_expressionContext).left = _x
	}
	{
		p.SetState(135)

		var _m = p.Match(SESParserParserOP_LOGICAL)

		localctx.(*Event_expressionContext).op = _m
	}
	{
		p.SetState(136)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttrModifiedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)

			var _m = p.Match(SESParserParserID)

			localctx.(*AttrModifiedContext).modifier = _m
		}
		{
			p.SetState(139)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(140)
			p.EventAttr()
		}
		{
			p.SetState(141)
			p.Match(SESParserParserR_BRACKET)
		}

	case 2:
		localctx = NewAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.EventAttr()
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
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

func (s *EventAttrContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserDOT)
}

func (s *EventAttrContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserDOT, i)
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
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(147)

			var _m = p.Match(SESParserParserID)

			localctx.(*EventAttrContext).eventName = _m
		}
		{
			p.SetState(148)
			p.Match(SESParserParserDOT)
		}

	}
	{
		p.SetState(151)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventAttrContext).attrName = _m
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SESParserParserDOT {
		{
			p.SetState(152)
			p.Match(SESParserParserDOT)
		}
		{
			p.SetState(153)
			p.Match(SESParserParserID)
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER:
		localctx = NewLiteral_numberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(SESParserParserNUMBER)
		}

	case SESParserParserSTRING:
		localctx = NewLiteral_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
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

func (s *GroupContext) GroupAttr() IGroupAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupAttrContext)
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
		p.SetState(163)
		p.Match(SESParserParserGROUP)
	}
	{
		p.SetState(164)
		p.GroupAttr()
	}

	return localctx
}

// IGroupAttrContext is an interface to support dynamic dispatch.
type IGroupAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAttr1 returns the attr1 rule contexts.
	GetAttr1() IEventAttrContext

	// GetExtraAttr returns the extraAttr rule contexts.
	GetExtraAttr() IGroupAttrContext

	// SetAttr1 sets the attr1 rule contexts.
	SetAttr1(IEventAttrContext)

	// SetExtraAttr sets the extraAttr rule contexts.
	SetExtraAttr(IGroupAttrContext)

	// IsGroupAttrContext differentiates from other interfaces.
	IsGroupAttrContext()
}

type GroupAttrContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	attr1     IEventAttrContext
	extraAttr IGroupAttrContext
}

func NewEmptyGroupAttrContext() *GroupAttrContext {
	var p = new(GroupAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_groupAttr
	return p
}

func (*GroupAttrContext) IsGroupAttrContext() {}

func NewGroupAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupAttrContext {
	var p = new(GroupAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_groupAttr

	return p
}

func (s *GroupAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupAttrContext) GetAttr1() IEventAttrContext { return s.attr1 }

func (s *GroupAttrContext) GetExtraAttr() IGroupAttrContext { return s.extraAttr }

func (s *GroupAttrContext) SetAttr1(v IEventAttrContext) { s.attr1 = v }

func (s *GroupAttrContext) SetExtraAttr(v IGroupAttrContext) { s.extraAttr = v }

func (s *GroupAttrContext) EventAttr() IEventAttrContext {
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

func (s *GroupAttrContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SESParserParserCOMMA)
}

func (s *GroupAttrContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SESParserParserCOMMA, i)
}

func (s *GroupAttrContext) AllGroupAttr() []IGroupAttrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupAttrContext); ok {
			len++
		}
	}

	tst := make([]IGroupAttrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupAttrContext); ok {
			tst[i] = t.(IGroupAttrContext)
			i++
		}
	}

	return tst
}

func (s *GroupAttrContext) GroupAttr(i int) IGroupAttrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupAttrContext); ok {
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

	return t.(IGroupAttrContext)
}

func (s *GroupAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterGroupAttr(s)
	}
}

func (s *GroupAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitGroupAttr(s)
	}
}

func (p *SESParserParser) GroupAttr() (localctx IGroupAttrContext) {
	this := p
	_ = this

	localctx = NewGroupAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SESParserParserRULE_groupAttr)

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
		p.SetState(166)

		var _x = p.EventAttr()

		localctx.(*GroupAttrContext).attr1 = _x
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(167)
				p.Match(SESParserParserCOMMA)
			}
			{
				p.SetState(168)

				var _x = p.GroupAttr()

				localctx.(*GroupAttrContext).extraAttr = _x
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}
