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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 188,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 7, 2, 50, 10, 2, 12, 2, 14, 2, 53, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 5, 3, 59, 10, 3, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 75, 10, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 83, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 7, 10, 92, 10, 10, 12, 10, 14, 10, 95, 11, 10, 3,
	11, 3, 11, 7, 11, 99, 10, 11, 12, 11, 14, 11, 102, 11, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 7, 12, 109, 10, 12, 12, 12, 14, 12, 112, 11, 12, 3,
	13, 3, 13, 3, 13, 7, 13, 117, 10, 13, 12, 13, 14, 13, 120, 11, 13, 3, 14,
	3, 14, 3, 14, 5, 14, 125, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 131,
	10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 138, 10, 17, 12, 17,
	14, 17, 141, 11, 17, 3, 18, 3, 18, 3, 18, 7, 18, 146, 10, 18, 12, 18, 14,
	18, 149, 11, 18, 3, 19, 3, 19, 3, 19, 5, 19, 154, 10, 19, 3, 20, 3, 20,
	3, 20, 5, 20, 159, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 5, 21, 170, 10, 21, 3, 22, 3, 22, 5, 22, 174, 10, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23, 181, 10, 23, 12, 23, 14, 23,
	184, 11, 23, 3, 24, 3, 24, 3, 24, 2, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 2, 5, 3, 2,
	14, 20, 3, 2, 21, 22, 3, 2, 23, 25, 2, 186, 2, 51, 3, 2, 2, 2, 4, 58, 3,
	2, 2, 2, 6, 60, 3, 2, 2, 2, 8, 74, 3, 2, 2, 2, 10, 76, 3, 2, 2, 2, 12,
	78, 3, 2, 2, 2, 14, 80, 3, 2, 2, 2, 16, 86, 3, 2, 2, 2, 18, 88, 3, 2, 2,
	2, 20, 96, 3, 2, 2, 2, 22, 105, 3, 2, 2, 2, 24, 113, 3, 2, 2, 2, 26, 124,
	3, 2, 2, 2, 28, 126, 3, 2, 2, 2, 30, 132, 3, 2, 2, 2, 32, 134, 3, 2, 2,
	2, 34, 142, 3, 2, 2, 2, 36, 153, 3, 2, 2, 2, 38, 155, 3, 2, 2, 2, 40, 169,
	3, 2, 2, 2, 42, 171, 3, 2, 2, 2, 44, 177, 3, 2, 2, 2, 46, 185, 3, 2, 2,
	2, 48, 50, 5, 4, 3, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2,
	54, 55, 7, 2, 2, 3, 55, 3, 3, 2, 2, 2, 56, 59, 5, 8, 5, 2, 57, 59, 5, 6,
	4, 2, 58, 56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 5, 3, 2, 2, 2, 60, 62,
	7, 3, 2, 2, 61, 63, 5, 12, 7, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 7, 3, 2, 2, 2, 64, 65, 5, 10, 6, 2, 65, 66, 7, 4, 2, 2, 66, 67, 5,
	12, 7, 2, 67, 75, 3, 2, 2, 2, 68, 75, 5, 12, 7, 2, 69, 70, 7, 5, 2, 2,
	70, 71, 7, 29, 2, 2, 71, 72, 5, 14, 8, 2, 72, 73, 5, 20, 11, 2, 73, 75,
	3, 2, 2, 2, 74, 64, 3, 2, 2, 2, 74, 68, 3, 2, 2, 2, 74, 69, 3, 2, 2, 2,
	75, 9, 3, 2, 2, 2, 76, 77, 7, 29, 2, 2, 77, 11, 3, 2, 2, 2, 78, 79, 5,
	22, 12, 2, 79, 13, 3, 2, 2, 2, 80, 82, 7, 6, 2, 2, 81, 83, 5, 16, 9, 2,
	82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7,
	7, 2, 2, 85, 15, 3, 2, 2, 2, 86, 87, 5, 18, 10, 2, 87, 17, 3, 2, 2, 2,
	88, 93, 7, 29, 2, 2, 89, 90, 7, 8, 2, 2, 90, 92, 7, 29, 2, 2, 91, 89, 3,
	2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94,
	19, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 100, 7, 9, 2, 2, 97, 99, 5, 4,
	3, 2, 98, 97, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100,
	101, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 104,
	7, 10, 2, 2, 104, 21, 3, 2, 2, 2, 105, 110, 5, 24, 13, 2, 106, 107, 7,
	11, 2, 2, 107, 109, 5, 24, 13, 2, 108, 106, 3, 2, 2, 2, 109, 112, 3, 2,
	2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 23, 3, 2, 2, 2,
	112, 110, 3, 2, 2, 2, 113, 118, 5, 26, 14, 2, 114, 115, 7, 12, 2, 2, 115,
	117, 5, 26, 14, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 25, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 121, 122, 7, 13, 2, 2, 122, 125, 5, 26, 14, 2, 123, 125, 5, 28, 15,
	2, 124, 121, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 27, 3, 2, 2, 2, 126,
	130, 5, 32, 17, 2, 127, 128, 5, 30, 16, 2, 128, 129, 5, 32, 17, 2, 129,
	131, 3, 2, 2, 2, 130, 127, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 29, 3,
	2, 2, 2, 132, 133, 9, 2, 2, 2, 133, 31, 3, 2, 2, 2, 134, 139, 5, 34, 18,
	2, 135, 136, 9, 3, 2, 2, 136, 138, 5, 34, 18, 2, 137, 135, 3, 2, 2, 2,
	138, 141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140,
	33, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 147, 5, 36, 19, 2, 143, 144,
	9, 4, 2, 2, 144, 146, 5, 36, 19, 2, 145, 143, 3, 2, 2, 2, 146, 149, 3,
	2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 35, 3, 2, 2,
	2, 149, 147, 3, 2, 2, 2, 150, 151, 7, 22, 2, 2, 151, 154, 5, 36, 19, 2,
	152, 154, 5, 38, 20, 2, 153, 150, 3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154,
	37, 3, 2, 2, 2, 155, 158, 5, 40, 21, 2, 156, 157, 7, 26, 2, 2, 157, 159,
	5, 36, 19, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 39, 3, 2,
	2, 2, 160, 161, 7, 6, 2, 2, 161, 162, 5, 12, 7, 2, 162, 163, 7, 7, 2, 2,
	163, 170, 3, 2, 2, 2, 164, 165, 7, 29, 2, 2, 165, 170, 5, 42, 22, 2, 166,
	170, 7, 29, 2, 2, 167, 170, 7, 28, 2, 2, 168, 170, 7, 27, 2, 2, 169, 160,
	3, 2, 2, 2, 169, 164, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2, 169, 167, 3, 2,
	2, 2, 169, 168, 3, 2, 2, 2, 170, 41, 3, 2, 2, 2, 171, 173, 7, 6, 2, 2,
	172, 174, 5, 44, 23, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174,
	175, 3, 2, 2, 2, 175, 176, 7, 7, 2, 2, 176, 43, 3, 2, 2, 2, 177, 182, 5,
	46, 24, 2, 178, 179, 7, 8, 2, 2, 179, 181, 5, 46, 24, 2, 180, 178, 3, 2,
	2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2,
	183, 45, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186, 5, 12, 7, 2, 186,
	47, 3, 2, 2, 2, 20, 51, 58, 62, 74, 82, 93, 100, 110, 118, 124, 130, 139,
	147, 153, 158, 169, 173, 182,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'return'", "'='", "'fun'", "'('", "')'", "','", "'{'", "'}'", "'or'",
	"'and'", "'not'", "'<'", "'>'", "'=='", "'>='", "'<='", "'<>'", "'!='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'**'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "STRING", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT",
	"LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt", "return_stmt", "expr_stmt", "primary_expr", "postfix_expr",
	"par_seq", "parlist", "namelist", "block", "or_test", "and_test", "not_test",
	"compare_expr", "comp_op", "arith_expr", "term", "factor", "power", "atom_expr",
	"arg_seq", "arg_list", "argument",
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
	SpartaParserSTRING         = 25
	SpartaParserNUMBER_LITERAL = 26
	SpartaParserIDENTIFIER     = 27
	SpartaParserCOMMENT        = 28
	SpartaParserLINE_COMMENT   = 29
	SpartaParserWS             = 30
	SpartaParserSHEBANG        = 31
)

// SpartaParser rules.
const (
	SpartaParserRULE_program      = 0
	SpartaParserRULE_stmt         = 1
	SpartaParserRULE_return_stmt  = 2
	SpartaParserRULE_expr_stmt    = 3
	SpartaParserRULE_primary_expr = 4
	SpartaParserRULE_postfix_expr = 5
	SpartaParserRULE_par_seq      = 6
	SpartaParserRULE_parlist      = 7
	SpartaParserRULE_namelist     = 8
	SpartaParserRULE_block        = 9
	SpartaParserRULE_or_test      = 10
	SpartaParserRULE_and_test     = 11
	SpartaParserRULE_not_test     = 12
	SpartaParserRULE_compare_expr = 13
	SpartaParserRULE_comp_op      = 14
	SpartaParserRULE_arith_expr   = 15
	SpartaParserRULE_term         = 16
	SpartaParserRULE_factor       = 17
	SpartaParserRULE_power        = 18
	SpartaParserRULE_atom_expr    = 19
	SpartaParserRULE_arg_seq      = 20
	SpartaParserRULE_arg_list     = 21
	SpartaParserRULE_argument     = 22
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

func (s *ProgramContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__0)|(1<<SpartaParserT__2)|(1<<SpartaParserT__3)|(1<<SpartaParserT__10)|(1<<SpartaParserT__19)|(1<<SpartaParserSTRING)|(1<<SpartaParserNUMBER_LITERAL)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(46)
			p.Stmt()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(SpartaParserEOF)
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

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SpartaParserRULE_stmt)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__2, SpartaParserT__3, SpartaParserT__10, SpartaParserT__19, SpartaParserSTRING, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Expr_stmt()
		}

	case SpartaParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Return_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_return_stmt
	return p
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpartaParserRULE_return_stmt)

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
		p.Match(SpartaParserT__0)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(59)
			p.Postfix_expr()
		}

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

func (s *Expr_stmtContext) Primary_expr() IPrimary_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Expr_stmtContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Expr_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Expr_stmtContext) Par_seq() IPar_seqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPar_seqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPar_seqContext)
}

func (s *Expr_stmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Expr_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Primary_expr()
		}
		{
			p.SetState(63)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(64)
			p.Postfix_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Postfix_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(SpartaParserT__2)
		}
		{
			p.SetState(68)
			p.Match(SpartaParserIDENTIFIER)
		}
		{
			p.SetState(69)
			p.Par_seq()
		}
		{
			p.SetState(70)
			p.Block()
		}

	}

	return localctx
}

// IPrimary_exprContext is an interface to support dynamic dispatch.
type IPrimary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_exprContext differentiates from other interfaces.
	IsPrimary_exprContext()
}

type Primary_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_exprContext() *Primary_exprContext {
	var p = new(Primary_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_primary_expr
	return p
}

func (*Primary_exprContext) IsPrimary_exprContext() {}

func NewPrimary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_exprContext {
	var p = new(Primary_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_primary_expr

	return p
}

func (s *Primary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_exprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Primary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Primary_expr() (localctx IPrimary_exprContext) {
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpartaParserRULE_primary_expr)

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
		p.SetState(74)
		p.Match(SpartaParserIDENTIFIER)
	}

	return localctx
}

// IPostfix_exprContext is an interface to support dynamic dispatch.
type IPostfix_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfix_exprContext differentiates from other interfaces.
	IsPostfix_exprContext()
}

type Postfix_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfix_exprContext() *Postfix_exprContext {
	var p = new(Postfix_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_postfix_expr
	return p
}

func (*Postfix_exprContext) IsPostfix_exprContext() {}

func NewPostfix_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Postfix_exprContext {
	var p = new(Postfix_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_postfix_expr

	return p
}

func (s *Postfix_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Postfix_exprContext) Or_test() IOr_testContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_testContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_testContext)
}

func (s *Postfix_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Postfix_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Postfix_expr() (localctx IPostfix_exprContext) {
	localctx = NewPostfix_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SpartaParserRULE_postfix_expr)

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
		p.SetState(76)
		p.Or_test()
	}

	return localctx
}

// IPar_seqContext is an interface to support dynamic dispatch.
type IPar_seqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPar_seqContext differentiates from other interfaces.
	IsPar_seqContext()
}

type Par_seqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPar_seqContext() *Par_seqContext {
	var p = new(Par_seqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_par_seq
	return p
}

func (*Par_seqContext) IsPar_seqContext() {}

func NewPar_seqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Par_seqContext {
	var p = new(Par_seqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_par_seq

	return p
}

func (s *Par_seqContext) GetParser() antlr.Parser { return s.parser }

func (s *Par_seqContext) Parlist() IParlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *Par_seqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Par_seqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Par_seq() (localctx IPar_seqContext) {
	localctx = NewPar_seqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SpartaParserRULE_par_seq)
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
		p.SetState(78)
		p.Match(SpartaParserT__3)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserIDENTIFIER {
		{
			p.SetState(79)
			p.Parlist()
		}

	}
	{
		p.SetState(82)
		p.Match(SpartaParserT__4)
	}

	return localctx
}

// IParlistContext is an interface to support dynamic dispatch.
type IParlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParlistContext differentiates from other interfaces.
	IsParlistContext()
}

type ParlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParlistContext() *ParlistContext {
	var p = new(ParlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_parlist
	return p
}

func (*ParlistContext) IsParlistContext() {}

func NewParlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParlistContext {
	var p = new(ParlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_parlist

	return p
}

func (s *ParlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParlistContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ParlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Parlist() (localctx IParlistContext) {
	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SpartaParserRULE_parlist)

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
		p.SetState(84)
		p.Namelist()
	}

	return localctx
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_namelist
	return p
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SpartaParserIDENTIFIER)
}

func (s *NamelistContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SpartaParserRULE_namelist)
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
		p.SetState(86)
		p.Match(SpartaParserIDENTIFIER)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__5 {
		{
			p.SetState(87)
			p.Match(SpartaParserT__5)
		}
		{
			p.SetState(88)
			p.Match(SpartaParserIDENTIFIER)
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SpartaParserRULE_block)
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
		p.SetState(94)
		p.Match(SpartaParserT__6)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__0)|(1<<SpartaParserT__2)|(1<<SpartaParserT__3)|(1<<SpartaParserT__10)|(1<<SpartaParserT__19)|(1<<SpartaParserSTRING)|(1<<SpartaParserNUMBER_LITERAL)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(95)
			p.Stmt()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(SpartaParserT__7)
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

func (p *SpartaParser) Or_test() (localctx IOr_testContext) {
	localctx = NewOr_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SpartaParserRULE_or_test)
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
		p.And_test()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__8 {
		{
			p.SetState(104)
			p.Match(SpartaParserT__8)
		}
		{
			p.SetState(105)
			p.And_test()
		}

		p.SetState(110)
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

func (p *SpartaParser) And_test() (localctx IAnd_testContext) {
	localctx = NewAnd_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SpartaParserRULE_and_test)
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
		p.SetState(111)
		p.Not_test()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__9 {
		{
			p.SetState(112)
			p.Match(SpartaParserT__9)
		}
		{
			p.SetState(113)
			p.Not_test()
		}

		p.SetState(118)
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

func (s *Not_testContext) Compare_expr() ICompare_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompare_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompare_exprContext)
}

func (s *Not_testContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_testContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Not_test() (localctx INot_testContext) {
	localctx = NewNot_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SpartaParserRULE_not_test)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(SpartaParserT__10)
		}
		{
			p.SetState(120)
			p.Not_test()
		}

	case SpartaParserT__3, SpartaParserT__19, SpartaParserSTRING, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Compare_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompare_exprContext is an interface to support dynamic dispatch.
type ICompare_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompare_exprContext differentiates from other interfaces.
	IsCompare_exprContext()
}

type Compare_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompare_exprContext() *Compare_exprContext {
	var p = new(Compare_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_compare_expr
	return p
}

func (*Compare_exprContext) IsCompare_exprContext() {}

func NewCompare_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compare_exprContext {
	var p = new(Compare_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_compare_expr

	return p
}

func (s *Compare_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Compare_exprContext) AllArith_expr() []IArith_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArith_exprContext)(nil)).Elem())
	var tst = make([]IArith_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArith_exprContext)
		}
	}

	return tst
}

func (s *Compare_exprContext) Arith_expr(i int) IArith_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArith_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArith_exprContext)
}

func (s *Compare_exprContext) Comp_op() IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *Compare_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compare_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Compare_expr() (localctx ICompare_exprContext) {
	localctx = NewCompare_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SpartaParserRULE_compare_expr)
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
		p.SetState(124)
		p.Arith_expr()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__11)|(1<<SpartaParserT__12)|(1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15)|(1<<SpartaParserT__16)|(1<<SpartaParserT__17))) != 0 {
		{
			p.SetState(125)
			p.Comp_op()
		}
		{
			p.SetState(126)
			p.Arith_expr()
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

func (p *SpartaParser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SpartaParserRULE_comp_op)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__11)|(1<<SpartaParserT__12)|(1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15)|(1<<SpartaParserT__16)|(1<<SpartaParserT__17))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (p *SpartaParser) Arith_expr() (localctx IArith_exprContext) {
	localctx = NewArith_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SpartaParserRULE_arith_expr)
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
		p.SetState(132)
		p.Term()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(133)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SpartaParserT__18 || _la == SpartaParserT__19) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(134)
				p.Term()
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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

func (p *SpartaParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SpartaParserRULE_term)
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
		p.SetState(140)
		p.Factor()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22))) != 0 {
		{
			p.SetState(141)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(142)
			p.Factor()
		}

		p.SetState(147)
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

func (p *SpartaParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SpartaParserRULE_factor)

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(SpartaParserT__19)
		}
		{
			p.SetState(149)
			p.Factor()
		}

	case SpartaParserT__3, SpartaParserSTRING, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
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

func (p *SpartaParser) Power() (localctx IPowerContext) {
	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_power)
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
		p.SetState(153)
		p.Atom_expr()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__23 {
		{
			p.SetState(154)
			p.Match(SpartaParserT__23)
		}
		{
			p.SetState(155)
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

func (s *Atom_exprContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Atom_exprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Atom_exprContext) Arg_seq() IArg_seqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_seqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_seqContext)
}

func (s *Atom_exprContext) NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SpartaParserNUMBER_LITERAL, 0)
}

func (s *Atom_exprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SpartaParserSTRING, 0)
}

func (s *Atom_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Atom_expr() (localctx IAtom_exprContext) {
	localctx = NewAtom_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_atom_expr)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(SpartaParserT__3)
		}
		{
			p.SetState(159)
			p.Postfix_expr()
		}
		{
			p.SetState(160)
			p.Match(SpartaParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(SpartaParserIDENTIFIER)
		}
		{
			p.SetState(163)
			p.Arg_seq()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(165)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(166)
			p.Match(SpartaParserSTRING)
		}

	}

	return localctx
}

// IArg_seqContext is an interface to support dynamic dispatch.
type IArg_seqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_seqContext differentiates from other interfaces.
	IsArg_seqContext()
}

type Arg_seqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_seqContext() *Arg_seqContext {
	var p = new(Arg_seqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_arg_seq
	return p
}

func (*Arg_seqContext) IsArg_seqContext() {}

func NewArg_seqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_seqContext {
	var p = new(Arg_seqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_arg_seq

	return p
}

func (s *Arg_seqContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_seqContext) Arg_list() IArg_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Arg_seqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_seqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Arg_seq() (localctx IArg_seqContext) {
	localctx = NewArg_seqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_arg_seq)
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
		p.SetState(169)
		p.Match(SpartaParserT__3)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__3)|(1<<SpartaParserT__10)|(1<<SpartaParserT__19)|(1<<SpartaParserSTRING)|(1<<SpartaParserNUMBER_LITERAL)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(170)
			p.Arg_list()
		}

	}
	{
		p.SetState(173)
		p.Match(SpartaParserT__4)
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

func (p *SpartaParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SpartaParserRULE_arg_list)
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
		p.Argument()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__5 {
		{
			p.SetState(176)
			p.Match(SpartaParserT__5)
		}
		{
			p.SetState(177)
			p.Argument()
		}

		p.SetState(182)
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

func (s *ArgumentContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SpartaParserRULE_argument)

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
		p.Postfix_expr()
	}

	return localctx
}
