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
		"absoluteDate", "relativeDate", "dateInterval", "where_expression",
		"binary_expr", "expr_operand", "eventAttr", "literal", "group", "groupAttr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 3, 0, 36, 8, 0, 1, 0, 4, 0, 39, 8, 0, 11, 0, 12, 0,
		40, 1, 0, 3, 0, 44, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 3, 2, 51, 8, 2,
		1, 2, 4, 2, 54, 8, 2, 11, 2, 12, 2, 55, 1, 3, 1, 3, 1, 3, 3, 3, 61, 8,
		3, 1, 3, 1, 3, 3, 3, 65, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 71, 8, 4,
		1, 4, 1, 4, 3, 4, 75, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 81, 8, 4, 1,
		5, 3, 5, 84, 8, 5, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 5, 3, 5, 91, 8, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 96, 8, 5, 1, 5, 1, 5, 3, 5, 100, 8, 5, 1, 5, 3, 5,
		103, 8, 5, 1, 5, 3, 5, 106, 8, 5, 1, 5, 1, 5, 3, 5, 110, 8, 5, 1, 6, 1,
		6, 3, 6, 114, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 5, 9, 125, 8, 9, 10, 9, 12, 9, 128, 9, 9, 1, 10, 1, 10, 1, 10, 5, 10,
		133, 8, 10, 10, 10, 12, 10, 136, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 142, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 150, 8,
		10, 10, 10, 12, 10, 153, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 166, 8, 12, 1, 13, 1, 13, 3,
		13, 170, 8, 13, 1, 13, 1, 13, 1, 13, 5, 13, 175, 8, 13, 10, 13, 12, 13,
		178, 9, 13, 1, 14, 1, 14, 3, 14, 182, 8, 14, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 5, 16, 190, 8, 16, 10, 16, 12, 16, 193, 9, 16, 1, 16, 0,
		1, 20, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		0, 0, 210, 0, 35, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6,
		57, 1, 0, 0, 0, 8, 80, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12, 113, 1, 0,
		0, 0, 14, 115, 1, 0, 0, 0, 16, 117, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20,
		141, 1, 0, 0, 0, 22, 154, 1, 0, 0, 0, 24, 165, 1, 0, 0, 0, 26, 169, 1,
		0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 183, 1, 0, 0, 0, 32, 186, 1, 0, 0, 0,
		34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1,
		0, 0, 0, 37, 39, 3, 4, 2, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 43, 1, 0, 0, 0, 42, 44, 3, 30,
		15, 0, 43, 42, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45,
		46, 5, 0, 0, 1, 46, 1, 1, 0, 0, 0, 47, 48, 5, 4, 0, 0, 48, 3, 1, 0, 0,
		0, 49, 51, 3, 10, 5, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53,
		1, 0, 0, 0, 52, 54, 3, 6, 3, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 5, 1, 0, 0, 0, 57, 58, 5, 5,
		0, 0, 58, 60, 5, 30, 0, 0, 59, 61, 3, 8, 4, 0, 60, 59, 1, 0, 0, 0, 60,
		61, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 63, 5, 13, 0, 0, 63, 65, 3, 20,
		10, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 7, 1, 0, 0, 0, 66, 81,
		5, 26, 0, 0, 67, 81, 5, 27, 0, 0, 68, 70, 5, 22, 0, 0, 69, 71, 5, 28, 0,
		0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74,
		5, 24, 0, 0, 73, 75, 5, 28, 0, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0,
		0, 75, 76, 1, 0, 0, 0, 76, 81, 5, 23, 0, 0, 77, 78, 5, 22, 0, 0, 78, 79,
		5, 28, 0, 0, 79, 81, 5, 23, 0, 0, 80, 66, 1, 0, 0, 0, 80, 67, 1, 0, 0,
		0, 80, 68, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 81, 9, 1, 0, 0, 0, 82, 84, 5,
		14, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85,
		110, 5, 12, 0, 0, 86, 88, 5, 14, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0,
		0, 0, 88, 90, 1, 0, 0, 0, 89, 91, 5, 12, 0, 0, 90, 89, 1, 0, 0, 0, 90,
		91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 5, 10, 0, 0, 93, 99, 3, 18,
		9, 0, 94, 96, 5, 14, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		97, 1, 0, 0, 0, 97, 98, 5, 11, 0, 0, 98, 100, 3, 18, 9, 0, 99, 95, 1, 0,
		0, 0, 99, 100, 1, 0, 0, 0, 100, 110, 1, 0, 0, 0, 101, 103, 5, 14, 0, 0,
		102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104,
		106, 5, 12, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 108, 5, 11, 0, 0, 108, 110, 3, 18, 9, 0, 109, 83, 1, 0,
		0, 0, 109, 87, 1, 0, 0, 0, 109, 102, 1, 0, 0, 0, 110, 11, 1, 0, 0, 0, 111,
		114, 3, 14, 7, 0, 112, 114, 3, 16, 8, 0, 113, 111, 1, 0, 0, 0, 113, 112,
		1, 0, 0, 0, 114, 13, 1, 0, 0, 0, 115, 116, 5, 29, 0, 0, 116, 15, 1, 0,
		0, 0, 117, 118, 5, 9, 0, 0, 118, 119, 3, 18, 9, 0, 119, 17, 1, 0, 0, 0,
		120, 121, 5, 28, 0, 0, 121, 126, 5, 17, 0, 0, 122, 123, 5, 14, 0, 0, 123,
		125, 3, 18, 9, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124,
		1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 19, 1, 0, 0, 0, 128, 126, 1, 0,
		0, 0, 129, 130, 6, 10, -1, 0, 130, 134, 3, 22, 11, 0, 131, 133, 3, 20,
		10, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 142, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137,
		138, 5, 20, 0, 0, 138, 139, 3, 20, 10, 0, 139, 140, 5, 21, 0, 0, 140, 142,
		1, 0, 0, 0, 141, 129, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 151, 1, 0,
		0, 0, 143, 144, 10, 2, 0, 0, 144, 145, 5, 14, 0, 0, 145, 150, 3, 20, 10,
		3, 146, 147, 10, 1, 0, 0, 147, 148, 5, 15, 0, 0, 148, 150, 3, 20, 10, 2,
		149, 143, 1, 0, 0, 0, 149, 146, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151,
		149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 21, 1, 0, 0, 0, 153, 151, 1,
		0, 0, 0, 154, 155, 3, 24, 12, 0, 155, 156, 5, 18, 0, 0, 156, 157, 3, 24,
		12, 0, 157, 23, 1, 0, 0, 0, 158, 159, 5, 30, 0, 0, 159, 160, 5, 20, 0,
		0, 160, 161, 3, 26, 13, 0, 161, 162, 5, 21, 0, 0, 162, 166, 1, 0, 0, 0,
		163, 166, 3, 26, 13, 0, 164, 166, 3, 28, 14, 0, 165, 158, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166, 25, 1, 0, 0, 0, 167, 168, 5,
		30, 0, 0, 168, 170, 5, 19, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 171, 1, 0, 0, 0, 171, 176, 5, 30, 0, 0, 172, 173, 5, 19, 0,
		0, 173, 175, 5, 30, 0, 0, 174, 172, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176,
		174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 27, 1, 0, 0, 0, 178, 176, 1,
		0, 0, 0, 179, 182, 5, 28, 0, 0, 180, 182, 5, 29, 0, 0, 181, 179, 1, 0,
		0, 0, 181, 180, 1, 0, 0, 0, 182, 29, 1, 0, 0, 0, 183, 184, 5, 16, 0, 0,
		184, 185, 3, 32, 16, 0, 185, 31, 1, 0, 0, 0, 186, 191, 3, 26, 13, 0, 187,
		188, 5, 24, 0, 0, 188, 190, 3, 32, 16, 0, 189, 187, 1, 0, 0, 0, 190, 193,
		1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 33, 1, 0,
		0, 0, 193, 191, 1, 0, 0, 0, 29, 35, 40, 43, 50, 55, 60, 64, 70, 74, 80,
		83, 87, 90, 95, 99, 102, 105, 109, 113, 126, 134, 141, 149, 151, 165, 169,
		176, 181, 191,
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
	SESParserParserRULE_where_expression = 10
	SESParserParserRULE_binary_expr      = 11
	SESParserParserRULE_expr_operand     = 12
	SESParserParserRULE_eventAttr        = 13
	SESParserParserRULE_literal          = 14
	SESParserParserRULE_group            = 15
	SESParserParserRULE_groupAttr        = 16
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWINDOW_TEXT {
		{
			p.SetState(34)
			p.Ses_window()
		}

	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserEVENT)|(1<<SESParserParserSKIP_)|(1<<SESParserParserWITHIN)|(1<<SESParserParserTHEN)|(1<<SESParserParserAND))) != 0) {
		{
			p.SetState(37)
			p.Ses()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserGROUP {
		{
			p.SetState(42)
			p.Group()
		}

	}
	{
		p.SetState(45)
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
		p.SetState(47)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SESParserParserSKIP_)|(1<<SESParserParserWITHIN)|(1<<SESParserParserTHEN)|(1<<SESParserParserAND))) != 0 {
		{
			p.SetState(49)
			p.Set_window()
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

func (s *EventContext) Where_expression() IWhere_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_expressionContext)
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SESParserParserWHERE {
		{
			p.SetState(62)
			p.Match(SESParserParserWHERE)
		}
		{
			p.SetState(63)
			p.where_expression(0)
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQty_plusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(SESParserParserPLUS)
		}

	case 2:
		localctx = NewQty_asteriscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Match(SESParserParserASTERISK)
		}

	case 3:
		localctx = NewQty_preciseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(69)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).from = _m
			}

		}
		{
			p.SetState(72)
			p.Match(SESParserParserCOMMA)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserNUMBER {
			{
				p.SetState(73)

				var _m = p.Match(SESParserParserNUMBER)

				localctx.(*Qty_preciseContext).to = _m
			}

		}
		{
			p.SetState(76)
			p.Match(SESParserParserR_CURLY_BRACKET)
		}

	case 4:
		localctx = NewQty_precise_altContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.Match(SESParserParserL_CURLY_BRACKET)
		}
		{
			p.SetState(78)

			var _m = p.Match(SESParserParserNUMBER)

			localctx.(*Qty_precise_altContext).exact = _m
		}
		{
			p.SetState(79)
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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(82)
				p.Match(SESParserParserAND)
			}

		}
		{
			p.SetState(85)
			p.Match(SESParserParserTHEN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(86)
				p.Match(SESParserParserAND)
			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserTHEN {
			{
				p.SetState(89)
				p.Match(SESParserParserTHEN)
			}

		}
		{
			p.SetState(92)
			p.Match(SESParserParserSKIP_)
		}
		{
			p.SetState(93)

			var _x = p.DateInterval()

			localctx.(*Set_windowContext).skip = _x
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

				localctx.(*Set_windowContext).within = _x
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserAND {
			{
				p.SetState(101)
				p.Match(SESParserParserAND)
			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SESParserParserTHEN {
			{
				p.SetState(104)
				p.Match(SESParserParserTHEN)
			}

		}
		{
			p.SetState(107)
			p.Match(SESParserParserWITHIN)
		}
		{
			p.SetState(108)

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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.AbsoluteDate()
		}

	case SESParserParserLAST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
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
		p.SetState(115)
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
		p.SetState(117)
		p.Match(SESParserParserLAST)
	}
	{
		p.SetState(118)

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
		p.SetState(120)

		var _m = p.Match(SESParserParserNUMBER)

		localctx.(*DateIntervalContext).num = _m
	}
	{
		p.SetState(121)

		var _m = p.Match(SESParserParserDATE_UNIT)

		localctx.(*DateIntervalContext).unit = _m
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.Match(SESParserParserAND)
			}
			{
				p.SetState(123)

				var _x = p.DateInterval()

				localctx.(*DateIntervalContext).extra = _x
			}

		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IWhere_expressionContext is an interface to support dynamic dispatch.
type IWhere_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhere_expressionContext differentiates from other interfaces.
	IsWhere_expressionContext()
}

type Where_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_expressionContext() *Where_expressionContext {
	var p = new(Where_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_where_expression
	return p
}

func (*Where_expressionContext) IsWhere_expressionContext() {}

func NewWhere_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_expressionContext {
	var p = new(Where_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_where_expression

	return p
}

func (s *Where_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_expressionContext) CopyFrom(ctx *Where_expressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Where_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_orContext struct {
	*Where_expressionContext
	left  IWhere_expressionContext
	right IWhere_expressionContext
}

func NewExpr_orContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_orContext {
	var p = new(Expr_orContext)

	p.Where_expressionContext = NewEmptyWhere_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Where_expressionContext))

	return p
}

func (s *Expr_orContext) GetLeft() IWhere_expressionContext { return s.left }

func (s *Expr_orContext) GetRight() IWhere_expressionContext { return s.right }

func (s *Expr_orContext) SetLeft(v IWhere_expressionContext) { s.left = v }

func (s *Expr_orContext) SetRight(v IWhere_expressionContext) { s.right = v }

func (s *Expr_orContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_orContext) OR() antlr.TerminalNode {
	return s.GetToken(SESParserParserOR, 0)
}

func (s *Expr_orContext) AllWhere_expression() []IWhere_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhere_expressionContext); ok {
			len++
		}
	}

	tst := make([]IWhere_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhere_expressionContext); ok {
			tst[i] = t.(IWhere_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Expr_orContext) Where_expression(i int) IWhere_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_expressionContext); ok {
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

	return t.(IWhere_expressionContext)
}

func (s *Expr_orContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterExpr_or(s)
	}
}

func (s *Expr_orContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitExpr_or(s)
	}
}

type Expr_binContext struct {
	*Where_expressionContext
	extra IWhere_expressionContext
}

func NewExpr_binContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_binContext {
	var p = new(Expr_binContext)

	p.Where_expressionContext = NewEmptyWhere_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Where_expressionContext))

	return p
}

func (s *Expr_binContext) GetExtra() IWhere_expressionContext { return s.extra }

func (s *Expr_binContext) SetExtra(v IWhere_expressionContext) { s.extra = v }

func (s *Expr_binContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_binContext) Binary_expr() IBinary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinary_exprContext)
}

func (s *Expr_binContext) AllWhere_expression() []IWhere_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhere_expressionContext); ok {
			len++
		}
	}

	tst := make([]IWhere_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhere_expressionContext); ok {
			tst[i] = t.(IWhere_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Expr_binContext) Where_expression(i int) IWhere_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_expressionContext); ok {
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

	return t.(IWhere_expressionContext)
}

func (s *Expr_binContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterExpr_bin(s)
	}
}

func (s *Expr_binContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitExpr_bin(s)
	}
}

type Expr_bracketedContext struct {
	*Where_expressionContext
}

func NewExpr_bracketedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_bracketedContext {
	var p = new(Expr_bracketedContext)

	p.Where_expressionContext = NewEmptyWhere_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Where_expressionContext))

	return p
}

func (s *Expr_bracketedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_bracketedContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserL_BRACKET, 0)
}

func (s *Expr_bracketedContext) Where_expression() IWhere_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhere_expressionContext)
}

func (s *Expr_bracketedContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(SESParserParserR_BRACKET, 0)
}

func (s *Expr_bracketedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterExpr_bracketed(s)
	}
}

func (s *Expr_bracketedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitExpr_bracketed(s)
	}
}

type Expr_andContext struct {
	*Where_expressionContext
	left  IWhere_expressionContext
	right IWhere_expressionContext
}

func NewExpr_andContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_andContext {
	var p = new(Expr_andContext)

	p.Where_expressionContext = NewEmptyWhere_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Where_expressionContext))

	return p
}

func (s *Expr_andContext) GetLeft() IWhere_expressionContext { return s.left }

func (s *Expr_andContext) GetRight() IWhere_expressionContext { return s.right }

func (s *Expr_andContext) SetLeft(v IWhere_expressionContext) { s.left = v }

func (s *Expr_andContext) SetRight(v IWhere_expressionContext) { s.right = v }

func (s *Expr_andContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_andContext) AND() antlr.TerminalNode {
	return s.GetToken(SESParserParserAND, 0)
}

func (s *Expr_andContext) AllWhere_expression() []IWhere_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhere_expressionContext); ok {
			len++
		}
	}

	tst := make([]IWhere_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhere_expressionContext); ok {
			tst[i] = t.(IWhere_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Expr_andContext) Where_expression(i int) IWhere_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhere_expressionContext); ok {
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

	return t.(IWhere_expressionContext)
}

func (s *Expr_andContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterExpr_and(s)
	}
}

func (s *Expr_andContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitExpr_and(s)
	}
}

func (p *SESParserParser) Where_expression() (localctx IWhere_expressionContext) {
	return p.where_expression(0)
}

func (p *SESParserParser) where_expression(_p int) (localctx IWhere_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewWhere_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IWhere_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SESParserParserRULE_where_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER, SESParserParserSTRING, SESParserParserID:
		localctx = NewExpr_binContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(130)
			p.Binary_expr()
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(131)

					var _x = p.where_expression(0)

					localctx.(*Expr_binContext).extra = _x
				}

			}
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}

	case SESParserParserL_BRACKET:
		localctx = NewExpr_bracketedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(138)
			p.where_expression(0)
		}
		{
			p.SetState(139)
			p.Match(SESParserParserR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_andContext(p, NewWhere_expressionContext(p, _parentctx, _parentState))
				localctx.(*Expr_andContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SESParserParserRULE_where_expression)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(144)
					p.Match(SESParserParserAND)
				}
				{
					p.SetState(145)

					var _x = p.where_expression(3)

					localctx.(*Expr_andContext).right = _x
				}

			case 2:
				localctx = NewExpr_orContext(p, NewWhere_expressionContext(p, _parentctx, _parentState))
				localctx.(*Expr_orContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SESParserParserRULE_where_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(147)
					p.Match(SESParserParserOR)
				}
				{
					p.SetState(148)

					var _x = p.where_expression(2)

					localctx.(*Expr_orContext).right = _x
				}

			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IBinary_exprContext is an interface to support dynamic dispatch.
type IBinary_exprContext interface {
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

	// IsBinary_exprContext differentiates from other interfaces.
	IsBinary_exprContext()
}

type Binary_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExpr_operandContext
	op     antlr.Token
	right  IExpr_operandContext
}

func NewEmptyBinary_exprContext() *Binary_exprContext {
	var p = new(Binary_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SESParserParserRULE_binary_expr
	return p
}

func (*Binary_exprContext) IsBinary_exprContext() {}

func NewBinary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_exprContext {
	var p = new(Binary_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SESParserParserRULE_binary_expr

	return p
}

func (s *Binary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_exprContext) GetOp() antlr.Token { return s.op }

func (s *Binary_exprContext) SetOp(v antlr.Token) { s.op = v }

func (s *Binary_exprContext) GetLeft() IExpr_operandContext { return s.left }

func (s *Binary_exprContext) GetRight() IExpr_operandContext { return s.right }

func (s *Binary_exprContext) SetLeft(v IExpr_operandContext) { s.left = v }

func (s *Binary_exprContext) SetRight(v IExpr_operandContext) { s.right = v }

func (s *Binary_exprContext) AllExpr_operand() []IExpr_operandContext {
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

func (s *Binary_exprContext) Expr_operand(i int) IExpr_operandContext {
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

func (s *Binary_exprContext) OP_LOGICAL() antlr.TerminalNode {
	return s.GetToken(SESParserParserOP_LOGICAL, 0)
}

func (s *Binary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.EnterBinary_expr(s)
	}
}

func (s *Binary_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SESParserListener); ok {
		listenerT.ExitBinary_expr(s)
	}
}

func (p *SESParserParser) Binary_expr() (localctx IBinary_exprContext) {
	this := p
	_ = this

	localctx = NewBinary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SESParserParserRULE_binary_expr)

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

		var _x = p.Expr_operand()

		localctx.(*Binary_exprContext).left = _x
	}
	{
		p.SetState(155)

		var _m = p.Match(SESParserParserOP_LOGICAL)

		localctx.(*Binary_exprContext).op = _m
	}
	{
		p.SetState(156)

		var _x = p.Expr_operand()

		localctx.(*Binary_exprContext).right = _x
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
	p.EnterRule(localctx, 24, SESParserParserRULE_expr_operand)

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

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttrModifiedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)

			var _m = p.Match(SESParserParserID)

			localctx.(*AttrModifiedContext).modifier = _m
		}
		{
			p.SetState(159)
			p.Match(SESParserParserL_BRACKET)
		}
		{
			p.SetState(160)
			p.EventAttr()
		}
		{
			p.SetState(161)
			p.Match(SESParserParserR_BRACKET)
		}

	case 2:
		localctx = NewAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.EventAttr()
		}

	case 3:
		localctx = NewLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
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
	p.EnterRule(localctx, 26, SESParserParserRULE_eventAttr)

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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(167)

			var _m = p.Match(SESParserParserID)

			localctx.(*EventAttrContext).eventName = _m
		}
		{
			p.SetState(168)
			p.Match(SESParserParserDOT)
		}

	}
	{
		p.SetState(171)

		var _m = p.Match(SESParserParserID)

		localctx.(*EventAttrContext).attrName = _m
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				p.Match(SESParserParserDOT)
			}
			{
				p.SetState(173)
				p.Match(SESParserParserID)
			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, SESParserParserRULE_literal)

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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SESParserParserNUMBER:
		localctx = NewLiteral_numberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Match(SESParserParserNUMBER)
		}

	case SESParserParserSTRING:
		localctx = NewLiteral_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
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
	p.EnterRule(localctx, 30, SESParserParserRULE_group)

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
		p.SetState(183)
		p.Match(SESParserParserGROUP)
	}
	{
		p.SetState(184)
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
	p.EnterRule(localctx, 32, SESParserParserRULE_groupAttr)

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
		p.SetState(186)

		var _x = p.EventAttr()

		localctx.(*GroupAttrContext).attr1 = _x
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(187)
				p.Match(SESParserParserCOMMA)
			}
			{
				p.SetState(188)

				var _x = p.GroupAttr()

				localctx.(*GroupAttrContext).extraAttr = _x
			}

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

func (p *SESParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *Where_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Where_expressionContext)
		}
		return p.Where_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SESParserParser) Where_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
