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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 154,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 53, 10, 3, 13, 3, 14, 3,
	54, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 7, 8, 70, 10, 8, 12, 8, 14, 8, 73, 11, 8, 3, 9, 3, 9, 3, 9, 7,
	9, 78, 10, 9, 12, 9, 14, 9, 81, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 86,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 92, 10, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 101, 10, 14, 12, 14, 14, 14, 104,
	11, 14, 3, 15, 3, 15, 3, 15, 7, 15, 109, 10, 15, 12, 15, 14, 15, 112, 11,
	15, 3, 16, 3, 16, 3, 16, 5, 16, 117, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17,
	122, 10, 17, 3, 18, 3, 18, 5, 18, 126, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 134, 10, 19, 3, 20, 3, 20, 5, 20, 138, 10, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 7, 21, 145, 10, 21, 12, 21, 14, 21,
	148, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2,
	5, 3, 2, 7, 13, 3, 2, 14, 15, 3, 2, 16, 18, 2, 146, 2, 47, 3, 2, 2, 2,
	4, 52, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 60, 3, 2,
	2, 2, 12, 64, 3, 2, 2, 2, 14, 66, 3, 2, 2, 2, 16, 74, 3, 2, 2, 2, 18, 85,
	3, 2, 2, 2, 20, 87, 3, 2, 2, 2, 22, 93, 3, 2, 2, 2, 24, 95, 3, 2, 2, 2,
	26, 97, 3, 2, 2, 2, 28, 105, 3, 2, 2, 2, 30, 116, 3, 2, 2, 2, 32, 118,
	3, 2, 2, 2, 34, 123, 3, 2, 2, 2, 36, 133, 3, 2, 2, 2, 38, 135, 3, 2, 2,
	2, 40, 141, 3, 2, 2, 2, 42, 149, 3, 2, 2, 2, 44, 151, 3, 2, 2, 2, 46, 48,
	5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2,
	49, 50, 7, 2, 2, 3, 50, 3, 3, 2, 2, 2, 51, 53, 5, 6, 4, 2, 52, 51, 3, 2,
	2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 5,
	3, 2, 2, 2, 56, 57, 5, 8, 5, 2, 57, 7, 3, 2, 2, 2, 58, 59, 5, 10, 6, 2,
	59, 9, 3, 2, 2, 2, 60, 61, 7, 26, 2, 2, 61, 62, 7, 3, 2, 2, 62, 63, 5,
	12, 7, 2, 63, 11, 3, 2, 2, 2, 64, 65, 5, 14, 8, 2, 65, 13, 3, 2, 2, 2,
	66, 71, 5, 16, 9, 2, 67, 68, 7, 4, 2, 2, 68, 70, 5, 16, 9, 2, 69, 67, 3,
	2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72,
	15, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 79, 5, 18, 10, 2, 75, 76, 7, 5,
	2, 2, 76, 78, 5, 18, 10, 2, 77, 75, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79,
	77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 17, 3, 2, 2, 2, 81, 79, 3, 2, 2,
	2, 82, 83, 7, 6, 2, 2, 83, 86, 5, 18, 10, 2, 84, 86, 5, 20, 11, 2, 85,
	82, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 19, 3, 2, 2, 2, 87, 91, 5, 24,
	13, 2, 88, 89, 5, 22, 12, 2, 89, 90, 5, 24, 13, 2, 90, 92, 3, 2, 2, 2,
	91, 88, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 21, 3, 2, 2, 2, 93, 94, 9,
	2, 2, 2, 94, 23, 3, 2, 2, 2, 95, 96, 5, 26, 14, 2, 96, 25, 3, 2, 2, 2,
	97, 102, 5, 28, 15, 2, 98, 99, 9, 3, 2, 2, 99, 101, 5, 28, 15, 2, 100,
	98, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3,
	2, 2, 2, 103, 27, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 110, 5, 30, 16,
	2, 106, 107, 9, 4, 2, 2, 107, 109, 5, 30, 16, 2, 108, 106, 3, 2, 2, 2,
	109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	29, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 114, 9, 3, 2, 2, 114, 117, 5,
	30, 16, 2, 115, 117, 5, 32, 17, 2, 116, 113, 3, 2, 2, 2, 116, 115, 3, 2,
	2, 2, 117, 31, 3, 2, 2, 2, 118, 121, 5, 34, 18, 2, 119, 120, 7, 19, 2,
	2, 120, 122, 5, 30, 16, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 33, 3, 2, 2, 2, 123, 125, 5, 36, 19, 2, 124, 126, 5, 38, 20, 2, 125,
	124, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 35, 3, 2, 2, 2, 127, 128, 7,
	20, 2, 2, 128, 129, 5, 44, 23, 2, 129, 130, 7, 21, 2, 2, 130, 134, 3, 2,
	2, 2, 131, 134, 7, 23, 2, 2, 132, 134, 7, 26, 2, 2, 133, 127, 3, 2, 2,
	2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 37, 3, 2, 2, 2, 135,
	137, 7, 20, 2, 2, 136, 138, 5, 40, 21, 2, 137, 136, 3, 2, 2, 2, 137, 138,
	3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 7, 21, 2, 2, 140, 39, 3, 2,
	2, 2, 141, 146, 5, 42, 22, 2, 142, 143, 7, 22, 2, 2, 143, 145, 5, 42, 22,
	2, 144, 142, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 41, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 5,
	12, 7, 2, 150, 43, 3, 2, 2, 2, 151, 152, 5, 12, 7, 2, 152, 45, 3, 2, 2,
	2, 16, 47, 54, 71, 79, 85, 91, 102, 110, 116, 121, 125, 133, 137, 146,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'or'", "'and'", "'not'", "'<'", "'>'", "'=='", "'>='", "'<='",
	"'<>'", "'!='", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'('", "')'",
	"','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "NUMBER_LITERAL", "FLOAT_LITERAL", "INT_LITERAL", "IDENTIFIER",
	"COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt_list", "stmt", "expr_stmt", "assign_stmt", "test", "or_test",
	"and_test", "not_test", "comparison", "comp_op", "expr", "arith_expr",
	"term", "factor", "power", "atom_expr", "atom", "trailer", "arg_list",
	"argument", "testlist_comp",
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
	SpartaParserNUMBER_LITERAL = 21
	SpartaParserFLOAT_LITERAL  = 22
	SpartaParserINT_LITERAL    = 23
	SpartaParserIDENTIFIER     = 24
	SpartaParserCOMMENT        = 25
	SpartaParserLINE_COMMENT   = 26
	SpartaParserWS             = 27
	SpartaParserSHEBANG        = 28
)

// SpartaParser rules.
const (
	SpartaParserRULE_program       = 0
	SpartaParserRULE_stmt_list     = 1
	SpartaParserRULE_stmt          = 2
	SpartaParserRULE_expr_stmt     = 3
	SpartaParserRULE_assign_stmt   = 4
	SpartaParserRULE_test          = 5
	SpartaParserRULE_or_test       = 6
	SpartaParserRULE_and_test      = 7
	SpartaParserRULE_not_test      = 8
	SpartaParserRULE_comparison    = 9
	SpartaParserRULE_comp_op       = 10
	SpartaParserRULE_expr          = 11
	SpartaParserRULE_arith_expr    = 12
	SpartaParserRULE_term          = 13
	SpartaParserRULE_factor        = 14
	SpartaParserRULE_power         = 15
	SpartaParserRULE_atom_expr     = 16
	SpartaParserRULE_atom          = 17
	SpartaParserRULE_trailer       = 18
	SpartaParserRULE_arg_list      = 19
	SpartaParserRULE_argument      = 20
	SpartaParserRULE_testlist_comp = 21
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

	if _la == SpartaParserIDENTIFIER {
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

	for ok := true; ok; ok = _la == SpartaParserIDENTIFIER {
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Assign_stmt()
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
	p.EnterRule(localctx, 8, SpartaParserRULE_assign_stmt)

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
		p.SetState(58)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(59)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(60)
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

func (s *TestContext) Or_test() IOr_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_testContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_testContext)
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
	p.EnterRule(localctx, 10, SpartaParserRULE_test)

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
		p.SetState(62)
		p.Or_test()
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
	p.EnterRule(localctx, 12, SpartaParserRULE_or_test)
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
		p.SetState(64)
		p.And_test()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__1 {
		{
			p.SetState(65)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(66)
			p.And_test()
		}

		p.SetState(71)
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
	p.EnterRule(localctx, 14, SpartaParserRULE_and_test)
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
		p.SetState(72)
		p.Not_test()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__2 {
		{
			p.SetState(73)
			p.Match(SpartaParserT__2)
		}
		{
			p.SetState(74)
			p.Not_test()
		}

		p.SetState(79)
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
	p.EnterRule(localctx, 16, SpartaParserRULE_not_test)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(SpartaParserT__3)
		}
		{
			p.SetState(81)
			p.Not_test()
		}

	case SpartaParserT__11, SpartaParserT__12, SpartaParserT__17, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
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

func (s *ComparisonContext) Comp_op() IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 18, SpartaParserRULE_comparison)
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
		p.Expr()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__4)|(1<<SpartaParserT__5)|(1<<SpartaParserT__6)|(1<<SpartaParserT__7)|(1<<SpartaParserT__8)|(1<<SpartaParserT__9)|(1<<SpartaParserT__10))) != 0 {
		{
			p.SetState(86)
			p.Comp_op()
		}
		{
			p.SetState(87)
			p.Expr()
		}

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
	p.EnterRule(localctx, 20, SpartaParserRULE_comp_op)
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
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__4)|(1<<SpartaParserT__5)|(1<<SpartaParserT__6)|(1<<SpartaParserT__7)|(1<<SpartaParserT__8)|(1<<SpartaParserT__9)|(1<<SpartaParserT__10))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

func (s *ExprContext) Arith_expr() IArith_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArith_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArith_exprContext)
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
	p.EnterRule(localctx, 22, SpartaParserRULE_expr)

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
		p.SetState(93)
		p.Arith_expr()
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
	p.EnterRule(localctx, 24, SpartaParserRULE_arith_expr)
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
		p.SetState(95)
		p.Term()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__11 || _la == SpartaParserT__12 {
		{
			p.SetState(96)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__11 || _la == SpartaParserT__12) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(97)
			p.Term()
		}

		p.SetState(102)
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
	p.EnterRule(localctx, 26, SpartaParserRULE_term)
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
		p.SetState(103)
		p.Factor()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15))) != 0 {
		{
			p.SetState(104)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(105)
			p.Factor()
		}

		p.SetState(110)
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
	p.EnterRule(localctx, 28, SpartaParserRULE_factor)
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

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__11, SpartaParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__11 || _la == SpartaParserT__12) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(112)
			p.Factor()
		}

	case SpartaParserT__17, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
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
	p.EnterRule(localctx, 30, SpartaParserRULE_power)
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
		p.SetState(116)
		p.Atom_expr()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__16 {
		{
			p.SetState(117)
			p.Match(SpartaParserT__16)
		}
		{
			p.SetState(118)
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

func (s *Atom_exprContext) Trailer() ITrailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrailerContext)
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
	p.EnterRule(localctx, 32, SpartaParserRULE_atom_expr)
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
		p.SetState(121)
		p.Atom()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__17 {
		{
			p.SetState(122)
			p.Trailer()
		}

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

func (s *AtomContext) Testlist_comp() ITestlist_compContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestlist_compContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestlist_compContext)
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
	p.EnterRule(localctx, 34, SpartaParserRULE_atom)

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
	case SpartaParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(SpartaParserT__17)
		}
		{
			p.SetState(126)
			p.Testlist_comp()
		}
		{
			p.SetState(127)
			p.Match(SpartaParserT__18)
		}

	case SpartaParserNUMBER_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(SpartaParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITrailerContext is an interface to support dynamic dispatch.
type ITrailerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrailerContext differentiates from other interfaces.
	IsTrailerContext()
}

type TrailerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrailerContext() *TrailerContext {
	var p = new(TrailerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_trailer
	return p
}

func (*TrailerContext) IsTrailerContext() {}

func NewTrailerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrailerContext {
	var p = new(TrailerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_trailer

	return p
}

func (s *TrailerContext) GetParser() antlr.Parser { return s.parser }

func (s *TrailerContext) Arg_list() IArg_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *TrailerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrailerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrailerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterTrailer(s)
	}
}

func (s *TrailerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitTrailer(s)
	}
}

func (s *TrailerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitTrailer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Trailer() (localctx ITrailerContext) {
	localctx = NewTrailerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_trailer)
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
		p.SetState(133)
		p.Match(SpartaParserT__17)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__3)|(1<<SpartaParserT__11)|(1<<SpartaParserT__12)|(1<<SpartaParserT__17)|(1<<SpartaParserNUMBER_LITERAL)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(134)
			p.Arg_list()
		}

	}
	{
		p.SetState(137)
		p.Match(SpartaParserT__18)
	}

	return localctx
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_arg_list
	return p
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *Arg_listContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (s *Arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitArg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_arg_list)
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
		p.SetState(139)
		p.Argument()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__19 {
		{
			p.SetState(140)
			p.Match(SpartaParserT__19)
		}
		{
			p.SetState(141)
			p.Argument()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_argument)

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
		p.SetState(147)
		p.Test()
	}

	return localctx
}

// ITestlist_compContext is an interface to support dynamic dispatch.
type ITestlist_compContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestlist_compContext differentiates from other interfaces.
	IsTestlist_compContext()
}

type Testlist_compContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestlist_compContext() *Testlist_compContext {
	var p = new(Testlist_compContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_testlist_comp
	return p
}

func (*Testlist_compContext) IsTestlist_compContext() {}

func NewTestlist_compContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Testlist_compContext {
	var p = new(Testlist_compContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_testlist_comp

	return p
}

func (s *Testlist_compContext) GetParser() antlr.Parser { return s.parser }

func (s *Testlist_compContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Testlist_compContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Testlist_compContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Testlist_compContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterTestlist_comp(s)
	}
}

func (s *Testlist_compContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitTestlist_comp(s)
	}
}

func (s *Testlist_compContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitTestlist_comp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Testlist_comp() (localctx ITestlist_compContext) {
	localctx = NewTestlist_compContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SpartaParserRULE_testlist_comp)

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
		p.SetState(149)
		p.Test()
	}

	return localctx
}
