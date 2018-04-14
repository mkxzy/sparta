// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 193,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 53, 10, 3, 13, 3, 14, 3,
	54, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 78,
	10, 8, 3, 9, 3, 9, 3, 9, 7, 9, 83, 10, 9, 12, 9, 14, 9, 86, 11, 9, 3, 10,
	3, 10, 3, 10, 7, 10, 91, 10, 10, 12, 10, 14, 10, 94, 11, 10, 3, 11, 3,
	11, 3, 11, 5, 11, 99, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 105, 10,
	12, 12, 12, 14, 12, 108, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 123, 10, 13,
	3, 14, 3, 14, 3, 14, 7, 14, 128, 10, 14, 12, 14, 14, 14, 131, 11, 14, 3,
	15, 3, 15, 3, 15, 7, 15, 136, 10, 15, 12, 15, 14, 15, 139, 11, 15, 3, 16,
	3, 16, 3, 16, 7, 16, 144, 10, 16, 12, 16, 14, 16, 147, 11, 16, 3, 17, 3,
	17, 3, 17, 7, 17, 152, 10, 17, 12, 17, 14, 17, 155, 11, 17, 3, 18, 3, 18,
	3, 18, 7, 18, 160, 10, 18, 12, 18, 14, 18, 163, 11, 18, 3, 19, 3, 19, 3,
	19, 7, 19, 168, 10, 19, 12, 19, 14, 19, 171, 11, 19, 3, 20, 3, 20, 3, 20,
	5, 20, 176, 10, 20, 3, 21, 3, 21, 3, 21, 5, 21, 181, 10, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 191, 10, 23, 3, 23,
	2, 2, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 2, 5, 3, 2, 22, 23, 3, 2, 24, 25, 3, 2, 26, 28, 2,
	198, 2, 47, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 60, 3,
	2, 2, 2, 10, 62, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16,
	79, 3, 2, 2, 2, 18, 87, 3, 2, 2, 2, 20, 98, 3, 2, 2, 2, 22, 100, 3, 2,
	2, 2, 24, 122, 3, 2, 2, 2, 26, 124, 3, 2, 2, 2, 28, 132, 3, 2, 2, 2, 30,
	140, 3, 2, 2, 2, 32, 148, 3, 2, 2, 2, 34, 156, 3, 2, 2, 2, 36, 164, 3,
	2, 2, 2, 38, 175, 3, 2, 2, 2, 40, 177, 3, 2, 2, 2, 42, 182, 3, 2, 2, 2,
	44, 190, 3, 2, 2, 2, 46, 48, 5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3,
	2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 7, 2, 2, 3, 50, 3, 3, 2, 2, 2, 51,
	53, 5, 6, 4, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2,
	2, 54, 55, 3, 2, 2, 2, 55, 5, 3, 2, 2, 2, 56, 57, 5, 8, 5, 2, 57, 7, 3,
	2, 2, 2, 58, 61, 5, 10, 6, 2, 59, 61, 5, 12, 7, 2, 60, 58, 3, 2, 2, 2,
	60, 59, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 63, 7, 3, 2, 2, 63, 64, 7, 33,
	2, 2, 64, 65, 7, 4, 2, 2, 65, 66, 5, 14, 8, 2, 66, 11, 3, 2, 2, 2, 67,
	68, 7, 33, 2, 2, 68, 69, 7, 4, 2, 2, 69, 70, 5, 14, 8, 2, 70, 13, 3, 2,
	2, 2, 71, 77, 5, 16, 9, 2, 72, 73, 7, 5, 2, 2, 73, 74, 5, 16, 9, 2, 74,
	75, 7, 6, 2, 2, 75, 76, 5, 14, 8, 2, 76, 78, 3, 2, 2, 2, 77, 72, 3, 2,
	2, 2, 77, 78, 3, 2, 2, 2, 78, 15, 3, 2, 2, 2, 79, 84, 5, 18, 10, 2, 80,
	81, 7, 7, 2, 2, 81, 83, 5, 18, 10, 2, 82, 80, 3, 2, 2, 2, 83, 86, 3, 2,
	2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 17, 3, 2, 2, 2, 86, 84,
	3, 2, 2, 2, 87, 92, 5, 20, 11, 2, 88, 89, 7, 8, 2, 2, 89, 91, 5, 20, 11,
	2, 90, 88, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 19, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96, 7, 9, 2, 2,
	96, 99, 5, 20, 11, 2, 97, 99, 5, 22, 12, 2, 98, 95, 3, 2, 2, 2, 98, 97,
	3, 2, 2, 2, 99, 21, 3, 2, 2, 2, 100, 106, 5, 26, 14, 2, 101, 102, 5, 24,
	13, 2, 102, 103, 5, 26, 14, 2, 103, 105, 3, 2, 2, 2, 104, 101, 3, 2, 2,
	2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	23, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 123, 7, 10, 2, 2, 110, 123,
	7, 11, 2, 2, 111, 123, 7, 12, 2, 2, 112, 123, 7, 13, 2, 2, 113, 123, 7,
	14, 2, 2, 114, 123, 7, 15, 2, 2, 115, 123, 7, 16, 2, 2, 116, 123, 7, 17,
	2, 2, 117, 118, 7, 9, 2, 2, 118, 123, 7, 17, 2, 2, 119, 123, 7, 18, 2,
	2, 120, 121, 7, 18, 2, 2, 121, 123, 7, 9, 2, 2, 122, 109, 3, 2, 2, 2, 122,
	110, 3, 2, 2, 2, 122, 111, 3, 2, 2, 2, 122, 112, 3, 2, 2, 2, 122, 113,
	3, 2, 2, 2, 122, 114, 3, 2, 2, 2, 122, 115, 3, 2, 2, 2, 122, 116, 3, 2,
	2, 2, 122, 117, 3, 2, 2, 2, 122, 119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2,
	123, 25, 3, 2, 2, 2, 124, 129, 5, 28, 15, 2, 125, 126, 7, 19, 2, 2, 126,
	128, 5, 28, 15, 2, 127, 125, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 27, 3, 2, 2, 2, 131, 129, 3, 2,
	2, 2, 132, 137, 5, 30, 16, 2, 133, 134, 7, 20, 2, 2, 134, 136, 5, 30, 16,
	2, 135, 133, 3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 29, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 145, 5,
	32, 17, 2, 141, 142, 7, 21, 2, 2, 142, 144, 5, 32, 17, 2, 143, 141, 3,
	2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2,
	2, 146, 31, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 153, 5, 34, 18, 2, 149,
	150, 9, 2, 2, 2, 150, 152, 5, 34, 18, 2, 151, 149, 3, 2, 2, 2, 152, 155,
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 33, 3, 2,
	2, 2, 155, 153, 3, 2, 2, 2, 156, 161, 5, 36, 19, 2, 157, 158, 9, 3, 2,
	2, 158, 160, 5, 36, 19, 2, 159, 157, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2,
	161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 35, 3, 2, 2, 2, 163, 161,
	3, 2, 2, 2, 164, 169, 5, 38, 20, 2, 165, 166, 9, 4, 2, 2, 166, 168, 5,
	38, 20, 2, 167, 165, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2,
	2, 2, 169, 170, 3, 2, 2, 2, 170, 37, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2,
	172, 173, 9, 3, 2, 2, 173, 176, 5, 38, 20, 2, 174, 176, 5, 40, 21, 2, 175,
	172, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 39, 3, 2, 2, 2, 177, 180, 5,
	42, 22, 2, 178, 179, 7, 29, 2, 2, 179, 181, 5, 38, 20, 2, 180, 178, 3,
	2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 41, 3, 2, 2, 2, 182, 183, 5, 44, 23,
	2, 183, 43, 3, 2, 2, 2, 184, 185, 7, 30, 2, 2, 185, 186, 5, 34, 18, 2,
	186, 187, 7, 31, 2, 2, 187, 191, 3, 2, 2, 2, 188, 191, 7, 32, 2, 2, 189,
	191, 7, 33, 2, 2, 190, 184, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 189,
	3, 2, 2, 2, 191, 45, 3, 2, 2, 2, 20, 47, 54, 60, 77, 84, 92, 98, 106, 122,
	129, 137, 145, 153, 161, 169, 175, 180, 190,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'var'", "'='", "'if'", "'else'", "'or'", "'and'", "'not'", "'<'",
	"'>'", "'=='", "'>='", "'<='", "'<>'", "'!='", "'in'", "'is'", "'|'", "'^'",
	"'&'", "'<<'", "'>>'", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'('",
	"')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "NUMBER_LITERAL", "IDENTIFIER",
	"COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt_list", "stmt", "expr_stmt", "var_stmt", "assign_stmt",
	"test", "or_test", "and_test", "not_test", "comparison", "comp_op", "expr",
	"xor_expr", "and_expr", "shift_expr", "arith_expr", "term", "factor", "power",
	"atom_expr", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SpartaParser struct {
	*antlr.BaseParser
}

func NewSpartaParser(input antlr.TokenStream) *SpartaParser {
	this := new(SpartaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sparta.g4"

	return this
}

// SpartaParser tokens.
const (
	SpartaParserEOF            = antlr.TokenEOF
	SpartaParserT__0           = 1
	SpartaParserT__1           = 2
	SpartaParserT__2           = 3
	SpartaParserT__3           = 4
	SpartaParserT__4           = 5
	SpartaParserT__5           = 6
	SpartaParserT__6           = 7
	SpartaParserT__7           = 8
	SpartaParserT__8           = 9
	SpartaParserT__9           = 10
	SpartaParserT__10          = 11
	SpartaParserT__11          = 12
	SpartaParserT__12          = 13
	SpartaParserT__13          = 14
	SpartaParserT__14          = 15
	SpartaParserT__15          = 16
	SpartaParserT__16          = 17
	SpartaParserT__17          = 18
	SpartaParserT__18          = 19
	SpartaParserT__19          = 20
	SpartaParserT__20          = 21
	SpartaParserT__21          = 22
	SpartaParserT__22          = 23
	SpartaParserT__23          = 24
	SpartaParserT__24          = 25
	SpartaParserT__25          = 26
	SpartaParserT__26          = 27
	SpartaParserT__27          = 28
	SpartaParserT__28          = 29
	SpartaParserNUMBER_LITERAL = 30
	SpartaParserIDENTIFIER     = 31
	SpartaParserCOMMENT        = 32
	SpartaParserLINE_COMMENT   = 33
	SpartaParserWS             = 34
	SpartaParserSHEBANG        = 35
)

// SpartaParser rules.
const (
	SpartaParserRULE_program     = 0
	SpartaParserRULE_stmt_list   = 1
	SpartaParserRULE_stmt        = 2
	SpartaParserRULE_expr_stmt   = 3
	SpartaParserRULE_var_stmt    = 4
	SpartaParserRULE_assign_stmt = 5
	SpartaParserRULE_test        = 6
	SpartaParserRULE_or_test     = 7
	SpartaParserRULE_and_test    = 8
	SpartaParserRULE_not_test    = 9
	SpartaParserRULE_comparison  = 10
	SpartaParserRULE_comp_op     = 11
	SpartaParserRULE_expr        = 12
	SpartaParserRULE_xor_expr    = 13
	SpartaParserRULE_and_expr    = 14
	SpartaParserRULE_shift_expr  = 15
	SpartaParserRULE_arith_expr  = 16
	SpartaParserRULE_term        = 17
	SpartaParserRULE_factor      = 18
	SpartaParserRULE_power       = 19
	SpartaParserRULE_atom_expr   = 20
	SpartaParserRULE_atom        = 21
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SpartaParserEOF, 0)
}

func (s *ProgramContext) Stmt_list() IStmt_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmt_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmt_listContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SpartaParserRULE_program)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__0 || _la == SpartaParserIDENTIFIER {
		{
			p.SetState(44)
			p.Stmt_list()
		}

	}
	{
		p.SetState(47)
		p.Match(SpartaParserEOF)
	}

	return localctx
}

// IStmt_listContext is an interface to support dynamic dispatch.
type IStmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmt_listContext differentiates from other interfaces.
	IsStmt_listContext()
}

type Stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_listContext() *Stmt_listContext {
	var p = new(Stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_stmt_list
	return p
}

func (*Stmt_listContext) IsStmt_listContext() {}

func NewStmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_listContext {
	var p = new(Stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_stmt_list

	return p
}

func (s *Stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_listContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *Stmt_listContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stmt_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterStmt_list(s)
	}
}

func (s *Stmt_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitStmt_list(s)
	}
}

func (s *Stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitStmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Stmt_list() (localctx IStmt_listContext) {
	localctx = NewStmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SpartaParserRULE_stmt_list)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SpartaParserT__0 || _la == SpartaParserIDENTIFIER {
		{
			p.SetState(49)
			p.Stmt()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Expr_stmt() IExpr_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpartaParserRULE_stmt)

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
		p.SetState(54)
		p.Expr_stmt()
	}

	return localctx
}

// IExpr_stmtContext is an interface to support dynamic dispatch.
type IExpr_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr_stmtContext differentiates from other interfaces.
	IsExpr_stmtContext()
}

type Expr_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_stmtContext() *Expr_stmtContext {
	var p = new(Expr_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_expr_stmt
	return p
}

func (*Expr_stmtContext) IsExpr_stmtContext() {}

func NewExpr_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_stmtContext {
	var p = new(Expr_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_expr_stmt

	return p
}

func (s *Expr_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_stmtContext) Var_stmt() IVar_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_stmtContext)
}

func (s *Expr_stmtContext) Assign_stmt() IAssign_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *Expr_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterExpr_stmt(s)
	}
}

func (s *Expr_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitExpr_stmt(s)
	}
}

func (s *Expr_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitExpr_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Expr_stmt() (localctx IExpr_stmtContext) {
	localctx = NewExpr_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpartaParserRULE_expr_stmt)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Var_stmt()
		}

	case SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Assign_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVar_stmtContext is an interface to support dynamic dispatch.
type IVar_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_stmtContext differentiates from other interfaces.
	IsVar_stmtContext()
}

type Var_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_stmtContext() *Var_stmtContext {
	var p = new(Var_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_var_stmt
	return p
}

func (*Var_stmtContext) IsVar_stmtContext() {}

func NewVar_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_stmtContext {
	var p = new(Var_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_var_stmt

	return p
}

func (s *Var_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Var_stmtContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Var_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVar_stmt(s)
	}
}

func (s *Var_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVar_stmt(s)
	}
}

func (s *Var_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitVar_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Var_stmt() (localctx IVar_stmtContext) {
	localctx = NewVar_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpartaParserRULE_var_stmt)

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
		p.SetState(60)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(61)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(62)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(63)
		p.Test()
	}

	return localctx
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_assign_stmt
	return p
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Assign_stmtContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitAssign_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SpartaParserRULE_assign_stmt)

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
		p.SetState(65)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(66)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(67)
		p.Test()
	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) AllOr_test() []IOr_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOr_testContext)(nil)).Elem())
	var tst = make([]IOr_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOr_testContext)
		}
	}

	return tst
}

func (s *TestContext) Or_test(i int) IOr_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOr_testContext)
}

func (s *TestContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitTest(s)
	}
}

func (s *TestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SpartaParserRULE_test)
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
		p.SetState(69)
		p.Or_test()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__2 {
		{
			p.SetState(70)
			p.Match(SpartaParserT__2)
		}
		{
			p.SetState(71)
			p.Or_test()
		}
		{
			p.SetState(72)
			p.Match(SpartaParserT__3)
		}
		{
			p.SetState(73)
			p.Test()
		}

	}

	return localctx
}

// IOr_testContext is an interface to support dynamic dispatch.
type IOr_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOr_testContext differentiates from other interfaces.
	IsOr_testContext()
}

type Or_testContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOr_testContext() *Or_testContext {
	var p = new(Or_testContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_or_test
	return p
}

func (*Or_testContext) IsOr_testContext() {}

func NewOr_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_testContext {
	var p = new(Or_testContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_or_test

	return p
}

func (s *Or_testContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_testContext) AllAnd_test() []IAnd_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnd_testContext)(nil)).Elem())
	var tst = make([]IAnd_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnd_testContext)
		}
	}

	return tst
}

func (s *Or_testContext) And_test(i int) IAnd_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnd_testContext)
}

func (s *Or_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOr_test(s)
	}
}

func (s *Or_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOr_test(s)
	}
}

func (s *Or_testContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitOr_test(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Or_test() (localctx IOr_testContext) {
	localctx = NewOr_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SpartaParserRULE_or_test)
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
		p.SetState(77)
		p.And_test()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__4 {
		{
			p.SetState(78)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(79)
			p.And_test()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnd_testContext is an interface to support dynamic dispatch.
type IAnd_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_testContext differentiates from other interfaces.
	IsAnd_testContext()
}

type And_testContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_testContext() *And_testContext {
	var p = new(And_testContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_and_test
	return p
}

func (*And_testContext) IsAnd_testContext() {}

func NewAnd_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_testContext {
	var p = new(And_testContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_and_test

	return p
}

func (s *And_testContext) GetParser() antlr.Parser { return s.parser }

func (s *And_testContext) AllNot_test() []INot_testContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INot_testContext)(nil)).Elem())
	var tst = make([]INot_testContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INot_testContext)
		}
	}

	return tst
}

func (s *And_testContext) Not_test(i int) INot_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_testContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INot_testContext)
}

func (s *And_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterAnd_test(s)
	}
}

func (s *And_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitAnd_test(s)
	}
}

func (s *And_testContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitAnd_test(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) And_test() (localctx IAnd_testContext) {
	localctx = NewAnd_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SpartaParserRULE_and_test)
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
		p.SetState(85)
		p.Not_test()
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__5 {
		{
			p.SetState(86)
			p.Match(SpartaParserT__5)
		}
		{
			p.SetState(87)
			p.Not_test()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INot_testContext is an interface to support dynamic dispatch.
type INot_testContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_testContext differentiates from other interfaces.
	IsNot_testContext()
}

type Not_testContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_testContext() *Not_testContext {
	var p = new(Not_testContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_not_test
	return p
}

func (*Not_testContext) IsNot_testContext() {}

func NewNot_testContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_testContext {
	var p = new(Not_testContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_not_test

	return p
}

func (s *Not_testContext) GetParser() antlr.Parser { return s.parser }

func (s *Not_testContext) Not_test() INot_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_testContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_testContext)
}

func (s *Not_testContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *Not_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_testContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterNot_test(s)
	}
}

func (s *Not_testContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitNot_test(s)
	}
}

func (s *Not_testContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitNot_test(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Not_test() (localctx INot_testContext) {
	localctx = NewNot_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SpartaParserRULE_not_test)

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

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(94)
			p.Not_test()
		}

	case SpartaParserT__21, SpartaParserT__22, SpartaParserT__27, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Comparison()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparisonContext) AllComp_op() []IComp_opContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComp_opContext)(nil)).Elem())
	var tst = make([]IComp_opContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComp_opContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Comp_op(i int) IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SpartaParserRULE_comparison)
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
		p.SetState(98)
		p.Expr()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__6)|(1<<SpartaParserT__7)|(1<<SpartaParserT__8)|(1<<SpartaParserT__9)|(1<<SpartaParserT__10)|(1<<SpartaParserT__11)|(1<<SpartaParserT__12)|(1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15))) != 0 {
		{
			p.SetState(99)
			p.Comp_op()
		}
		{
			p.SetState(100)
			p.Expr()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComp_opContext is an interface to support dynamic dispatch.
type IComp_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComp_opContext differentiates from other interfaces.
	IsComp_opContext()
}

type Comp_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_opContext() *Comp_opContext {
	var p = new(Comp_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_comp_op
	return p
}

func (*Comp_opContext) IsComp_opContext() {}

func NewComp_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_opContext {
	var p = new(Comp_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_comp_op

	return p
}

func (s *Comp_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Comp_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comp_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterComp_op(s)
	}
}

func (s *Comp_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitComp_op(s)
	}
}

func (s *Comp_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitComp_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SpartaParserRULE_comp_op)

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

	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Match(SpartaParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(SpartaParserT__8)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Match(SpartaParserT__9)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(SpartaParserT__10)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.Match(SpartaParserT__11)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(112)
			p.Match(SpartaParserT__12)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(113)
			p.Match(SpartaParserT__13)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(114)
			p.Match(SpartaParserT__14)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(115)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(116)
			p.Match(SpartaParserT__14)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(117)
			p.Match(SpartaParserT__15)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(118)
			p.Match(SpartaParserT__15)
		}
		{
			p.SetState(119)
			p.Match(SpartaParserT__6)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllXor_expr() []IXor_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXor_exprContext)(nil)).Elem())
	var tst = make([]IXor_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXor_exprContext)
		}
	}

	return tst
}

func (s *ExprContext) Xor_expr(i int) IXor_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXor_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXor_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SpartaParserRULE_expr)
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
		p.SetState(122)
		p.Xor_expr()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__16 {
		{
			p.SetState(123)
			p.Match(SpartaParserT__16)
		}
		{
			p.SetState(124)
			p.Xor_expr()
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IXor_exprContext is an interface to support dynamic dispatch.
type IXor_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXor_exprContext differentiates from other interfaces.
	IsXor_exprContext()
}

type Xor_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXor_exprContext() *Xor_exprContext {
	var p = new(Xor_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_xor_expr
	return p
}

func (*Xor_exprContext) IsXor_exprContext() {}

func NewXor_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Xor_exprContext {
	var p = new(Xor_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_xor_expr

	return p
}

func (s *Xor_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Xor_exprContext) AllAnd_expr() []IAnd_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnd_exprContext)(nil)).Elem())
	var tst = make([]IAnd_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnd_exprContext)
		}
	}

	return tst
}

func (s *Xor_exprContext) And_expr(i int) IAnd_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnd_exprContext)
}

func (s *Xor_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Xor_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Xor_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterXor_expr(s)
	}
}

func (s *Xor_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitXor_expr(s)
	}
}

func (s *Xor_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitXor_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Xor_expr() (localctx IXor_exprContext) {
	localctx = NewXor_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SpartaParserRULE_xor_expr)
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
		p.SetState(130)
		p.And_expr()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__17 {
		{
			p.SetState(131)
			p.Match(SpartaParserT__17)
		}
		{
			p.SetState(132)
			p.And_expr()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnd_exprContext is an interface to support dynamic dispatch.
type IAnd_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_exprContext differentiates from other interfaces.
	IsAnd_exprContext()
}

type And_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_exprContext() *And_exprContext {
	var p = new(And_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_and_expr
	return p
}

func (*And_exprContext) IsAnd_exprContext() {}

func NewAnd_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_exprContext {
	var p = new(And_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_and_expr

	return p
}

func (s *And_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *And_exprContext) AllShift_expr() []IShift_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShift_exprContext)(nil)).Elem())
	var tst = make([]IShift_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShift_exprContext)
		}
	}

	return tst
}

func (s *And_exprContext) Shift_expr(i int) IShift_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShift_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShift_exprContext)
}

func (s *And_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterAnd_expr(s)
	}
}

func (s *And_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitAnd_expr(s)
	}
}

func (s *And_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitAnd_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) And_expr() (localctx IAnd_exprContext) {
	localctx = NewAnd_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SpartaParserRULE_and_expr)
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
		p.SetState(138)
		p.Shift_expr()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__18 {
		{
			p.SetState(139)
			p.Match(SpartaParserT__18)
		}
		{
			p.SetState(140)
			p.Shift_expr()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IShift_exprContext is an interface to support dynamic dispatch.
type IShift_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShift_exprContext differentiates from other interfaces.
	IsShift_exprContext()
}

type Shift_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShift_exprContext() *Shift_exprContext {
	var p = new(Shift_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_shift_expr
	return p
}

func (*Shift_exprContext) IsShift_exprContext() {}

func NewShift_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shift_exprContext {
	var p = new(Shift_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_shift_expr

	return p
}

func (s *Shift_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Shift_exprContext) AllArith_expr() []IArith_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArith_exprContext)(nil)).Elem())
	var tst = make([]IArith_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArith_exprContext)
		}
	}

	return tst
}

func (s *Shift_exprContext) Arith_expr(i int) IArith_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArith_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArith_exprContext)
}

func (s *Shift_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shift_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shift_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterShift_expr(s)
	}
}

func (s *Shift_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitShift_expr(s)
	}
}

func (s *Shift_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitShift_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Shift_expr() (localctx IShift_exprContext) {
	localctx = NewShift_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SpartaParserRULE_shift_expr)
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
		p.SetState(146)
		p.Arith_expr()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__19 || _la == SpartaParserT__20 {
		{
			p.SetState(147)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__19 || _la == SpartaParserT__20) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(148)
			p.Arith_expr()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArith_exprContext is an interface to support dynamic dispatch.
type IArith_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArith_exprContext differentiates from other interfaces.
	IsArith_exprContext()
}

type Arith_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArith_exprContext() *Arith_exprContext {
	var p = new(Arith_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_arith_expr
	return p
}

func (*Arith_exprContext) IsArith_exprContext() {}

func NewArith_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arith_exprContext {
	var p = new(Arith_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_arith_expr

	return p
}

func (s *Arith_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Arith_exprContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Arith_exprContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Arith_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arith_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arith_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterArith_expr(s)
	}
}

func (s *Arith_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitArith_expr(s)
	}
}

func (s *Arith_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitArith_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Arith_expr() (localctx IArith_exprContext) {
	localctx = NewArith_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SpartaParserRULE_arith_expr)
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
		p.SetState(154)
		p.Term()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__21 || _la == SpartaParserT__22 {
		{
			p.SetState(155)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__21 || _la == SpartaParserT__22) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(156)
			p.Term()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SpartaParserRULE_term)
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
		p.SetState(162)
		p.Factor()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25))) != 0 {
		{
			p.SetState(163)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(164)
			p.Factor()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) Power() IPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_factor)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__21, SpartaParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__21 || _la == SpartaParserT__22) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(171)
			p.Factor()
		}

	case SpartaParserT__27, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Power()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPowerContext is an interface to support dynamic dispatch.
type IPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowerContext differentiates from other interfaces.
	IsPowerContext()
}

type PowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerContext() *PowerContext {
	var p = new(PowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_power
	return p
}

func (*PowerContext) IsPowerContext() {}

func NewPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerContext {
	var p = new(PowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_power

	return p
}

func (s *PowerContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerContext) Atom_expr() IAtom_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_exprContext)
}

func (s *PowerContext) Factor() IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Power() (localctx IPowerContext) {
	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_power)
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
		p.SetState(175)
		p.Atom_expr()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__26 {
		{
			p.SetState(176)
			p.Match(SpartaParserT__26)
		}
		{
			p.SetState(177)
			p.Factor()
		}

	}

	return localctx
}

// IAtom_exprContext is an interface to support dynamic dispatch.
type IAtom_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtom_exprContext differentiates from other interfaces.
	IsAtom_exprContext()
}

type Atom_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtom_exprContext() *Atom_exprContext {
	var p = new(Atom_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_atom_expr
	return p
}

func (*Atom_exprContext) IsAtom_exprContext() {}

func NewAtom_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atom_exprContext {
	var p = new(Atom_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_atom_expr

	return p
}

func (s *Atom_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Atom_exprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Atom_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atom_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterAtom_expr(s)
	}
}

func (s *Atom_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitAtom_expr(s)
	}
}

func (s *Atom_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitAtom_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Atom_expr() (localctx IAtom_exprContext) {
	localctx = NewAtom_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_atom_expr)

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
		p.SetState(180)
		p.Atom()
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Arith_expr() IArith_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArith_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArith_exprContext)
}

func (s *AtomContext) NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SpartaParserNUMBER_LITERAL, 0)
}

func (s *AtomContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SpartaParserRULE_atom)

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

	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__27:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Match(SpartaParserT__27)
		}
		{
			p.SetState(183)
			p.Arith_expr()
		}
		{
			p.SetState(184)
			p.Match(SpartaParserT__28)
		}

	case SpartaParserNUMBER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(SpartaParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
