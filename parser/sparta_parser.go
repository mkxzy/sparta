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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 170,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 7, 2,
	46, 10, 2, 12, 2, 14, 2, 49, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 58, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 75, 10, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11, 9, 3, 10, 3, 10,
	5, 10, 89, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5,
	13, 98, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 105, 10, 14,
	12, 14, 14, 14, 108, 11, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 119, 10, 16, 12, 16, 14, 16, 122, 11, 16, 3,
	16, 3, 16, 5, 16, 126, 10, 16, 3, 17, 3, 17, 7, 17, 130, 10, 17, 12, 17,
	14, 17, 133, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7,
	19, 142, 10, 19, 12, 19, 14, 19, 145, 11, 19, 3, 20, 3, 20, 3, 20, 7, 20,
	150, 10, 20, 12, 20, 14, 20, 153, 11, 20, 3, 21, 5, 21, 156, 10, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	168, 10, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 4, 3, 2, 14, 15, 3, 2, 16, 18, 2,
	168, 2, 47, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 59, 3, 2, 2, 2, 8, 63, 3,
	2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 69, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16,
	78, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 90, 3, 2, 2, 2, 22, 92, 3, 2, 2,
	2, 24, 95, 3, 2, 2, 2, 26, 101, 3, 2, 2, 2, 28, 109, 3, 2, 2, 2, 30, 111,
	3, 2, 2, 2, 32, 127, 3, 2, 2, 2, 34, 136, 3, 2, 2, 2, 36, 138, 3, 2, 2,
	2, 38, 146, 3, 2, 2, 2, 40, 155, 3, 2, 2, 2, 42, 167, 3, 2, 2, 2, 44, 46,
	5, 4, 3, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2,
	47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 51, 7,
	2, 2, 3, 51, 3, 3, 2, 2, 2, 52, 58, 5, 6, 4, 2, 53, 58, 5, 8, 5, 2, 54,
	58, 5, 18, 10, 2, 55, 58, 5, 20, 11, 2, 56, 58, 5, 30, 16, 2, 57, 52, 3,
	2, 2, 2, 57, 53, 3, 2, 2, 2, 57, 54, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57,
	56, 3, 2, 2, 2, 58, 5, 3, 2, 2, 2, 59, 60, 7, 21, 2, 2, 60, 61, 7, 3, 2,
	2, 61, 62, 5, 34, 18, 2, 62, 7, 3, 2, 2, 2, 63, 64, 7, 4, 2, 2, 64, 65,
	5, 10, 6, 2, 65, 66, 5, 12, 7, 2, 66, 9, 3, 2, 2, 2, 67, 68, 7, 21, 2,
	2, 68, 11, 3, 2, 2, 2, 69, 70, 5, 14, 8, 2, 70, 71, 5, 32, 17, 2, 71, 13,
	3, 2, 2, 2, 72, 74, 7, 5, 2, 2, 73, 75, 5, 16, 9, 2, 74, 73, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 7, 6, 2, 2, 77, 15, 3,
	2, 2, 2, 78, 83, 7, 21, 2, 2, 79, 80, 7, 7, 2, 2, 80, 82, 7, 21, 2, 2,
	81, 79, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3,
	2, 2, 2, 84, 17, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88, 7, 8, 2, 2, 87,
	89, 5, 34, 18, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 19, 3, 2,
	2, 2, 90, 91, 5, 22, 12, 2, 91, 21, 3, 2, 2, 2, 92, 93, 5, 10, 6, 2, 93,
	94, 5, 24, 13, 2, 94, 23, 3, 2, 2, 2, 95, 97, 7, 5, 2, 2, 96, 98, 5, 26,
	14, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99,
	100, 7, 6, 2, 2, 100, 25, 3, 2, 2, 2, 101, 106, 5, 28, 15, 2, 102, 103,
	7, 7, 2, 2, 103, 105, 5, 28, 15, 2, 104, 102, 3, 2, 2, 2, 105, 108, 3,
	2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 27, 3, 2, 2,
	2, 108, 106, 3, 2, 2, 2, 109, 110, 5, 34, 18, 2, 110, 29, 3, 2, 2, 2, 111,
	112, 7, 9, 2, 2, 112, 113, 5, 34, 18, 2, 113, 120, 5, 32, 17, 2, 114, 115,
	7, 10, 2, 2, 115, 116, 5, 34, 18, 2, 116, 117, 5, 32, 17, 2, 117, 119,
	3, 2, 2, 2, 118, 114, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 120, 121, 3, 2, 2, 2, 121, 125, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2,
	123, 124, 7, 11, 2, 2, 124, 126, 5, 32, 17, 2, 125, 123, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 31, 3, 2, 2, 2, 127, 131, 7, 12, 2, 2, 128, 130,
	5, 4, 3, 2, 129, 128, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2,
	134, 135, 7, 13, 2, 2, 135, 33, 3, 2, 2, 2, 136, 137, 5, 36, 19, 2, 137,
	35, 3, 2, 2, 2, 138, 143, 5, 38, 20, 2, 139, 140, 9, 2, 2, 2, 140, 142,
	5, 38, 20, 2, 141, 139, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3,
	2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 37, 3, 2, 2, 2, 145, 143, 3, 2, 2,
	2, 146, 151, 5, 40, 21, 2, 147, 148, 9, 3, 2, 2, 148, 150, 5, 40, 21, 2,
	149, 147, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151,
	152, 3, 2, 2, 2, 152, 39, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 156, 7,
	15, 2, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 3, 2, 2,
	2, 157, 158, 5, 42, 22, 2, 158, 41, 3, 2, 2, 2, 159, 160, 7, 5, 2, 2, 160,
	161, 5, 34, 18, 2, 161, 162, 7, 6, 2, 2, 162, 168, 3, 2, 2, 2, 163, 168,
	5, 22, 12, 2, 164, 168, 7, 21, 2, 2, 165, 168, 7, 20, 2, 2, 166, 168, 7,
	19, 2, 2, 167, 159, 3, 2, 2, 2, 167, 163, 3, 2, 2, 2, 167, 164, 3, 2, 2,
	2, 167, 165, 3, 2, 2, 2, 167, 166, 3, 2, 2, 2, 168, 43, 3, 2, 2, 2, 16,
	47, 57, 74, 83, 88, 97, 106, 120, 125, 131, 143, 151, 155, 167,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'fun'", "'('", "')'", "','", "'return'", "'if'", "'else if'",
	"'else'", "'{'", "'}'", "'+'", "'-'", "'*'", "'/'", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt", "assign_stmt", "fundef_stmt", "fun_name", "fun_body",
	"fun_par", "namelist", "return_stmt", "funcall_stmt", "funcall_expr", "arg_expr",
	"arg_list", "arg", "if_stmt", "block", "postfix_expr", "arith_expr", "term",
	"factor", "atom_expr",
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
	SpartaParserSTRING         = 17
	SpartaParserNUMBER_LITERAL = 18
	SpartaParserIDENTIFIER     = 19
	SpartaParserCOMMENT        = 20
	SpartaParserLINE_COMMENT   = 21
	SpartaParserWS             = 22
	SpartaParserSHEBANG        = 23
)

// SpartaParser rules.
const (
	SpartaParserRULE_program      = 0
	SpartaParserRULE_stmt         = 1
	SpartaParserRULE_assign_stmt  = 2
	SpartaParserRULE_fundef_stmt  = 3
	SpartaParserRULE_fun_name     = 4
	SpartaParserRULE_fun_body     = 5
	SpartaParserRULE_fun_par      = 6
	SpartaParserRULE_namelist     = 7
	SpartaParserRULE_return_stmt  = 8
	SpartaParserRULE_funcall_stmt = 9
	SpartaParserRULE_funcall_expr = 10
	SpartaParserRULE_arg_expr     = 11
	SpartaParserRULE_arg_list     = 12
	SpartaParserRULE_arg          = 13
	SpartaParserRULE_if_stmt      = 14
	SpartaParserRULE_block        = 15
	SpartaParserRULE_postfix_expr = 16
	SpartaParserRULE_arith_expr   = 17
	SpartaParserRULE_term         = 18
	SpartaParserRULE_factor       = 19
	SpartaParserRULE_atom_expr    = 20
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__5)|(1<<SpartaParserT__6)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(42)
			p.Stmt()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
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

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Fundef_stmt() IFundef_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFundef_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFundef_stmtContext)
}

func (s *StmtContext) Return_stmt() IReturn_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StmtContext) Funcall_stmt() IFuncall_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncall_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncall_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Fundef_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Return_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Funcall_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.If_stmt()
		}

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

func (s *Assign_stmtContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpartaParserRULE_assign_stmt)

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
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(58)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(59)
		p.Postfix_expr()
	}

	return localctx
}

// IFundef_stmtContext is an interface to support dynamic dispatch.
type IFundef_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFundef_stmtContext differentiates from other interfaces.
	IsFundef_stmtContext()
}

type Fundef_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFundef_stmtContext() *Fundef_stmtContext {
	var p = new(Fundef_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fundef_stmt
	return p
}

func (*Fundef_stmtContext) IsFundef_stmtContext() {}

func NewFundef_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fundef_stmtContext {
	var p = new(Fundef_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fundef_stmt

	return p
}

func (s *Fundef_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Fundef_stmtContext) Fun_name() IFun_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFun_nameContext)
}

func (s *Fundef_stmtContext) Fun_body() IFun_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFun_bodyContext)
}

func (s *Fundef_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fundef_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Fundef_stmt() (localctx IFundef_stmtContext) {
	localctx = NewFundef_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpartaParserRULE_fundef_stmt)

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
		p.SetState(61)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(62)
		p.Fun_name()
	}
	{
		p.SetState(63)
		p.Fun_body()
	}

	return localctx
}

// IFun_nameContext is an interface to support dynamic dispatch.
type IFun_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFun_nameContext differentiates from other interfaces.
	IsFun_nameContext()
}

type Fun_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_nameContext() *Fun_nameContext {
	var p = new(Fun_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fun_name
	return p
}

func (*Fun_nameContext) IsFun_nameContext() {}

func NewFun_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_nameContext {
	var p = new(Fun_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fun_name

	return p
}

func (s *Fun_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Fun_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Fun_name() (localctx IFun_nameContext) {
	localctx = NewFun_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpartaParserRULE_fun_name)

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

	return localctx
}

// IFun_bodyContext is an interface to support dynamic dispatch.
type IFun_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFun_bodyContext differentiates from other interfaces.
	IsFun_bodyContext()
}

type Fun_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_bodyContext() *Fun_bodyContext {
	var p = new(Fun_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fun_body
	return p
}

func (*Fun_bodyContext) IsFun_bodyContext() {}

func NewFun_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_bodyContext {
	var p = new(Fun_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fun_body

	return p
}

func (s *Fun_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_bodyContext) Fun_par() IFun_parContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_parContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFun_parContext)
}

func (s *Fun_bodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Fun_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Fun_body() (localctx IFun_bodyContext) {
	localctx = NewFun_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SpartaParserRULE_fun_body)

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
		p.SetState(67)
		p.Fun_par()
	}
	{
		p.SetState(68)
		p.Block()
	}

	return localctx
}

// IFun_parContext is an interface to support dynamic dispatch.
type IFun_parContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFun_parContext differentiates from other interfaces.
	IsFun_parContext()
}

type Fun_parContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_parContext() *Fun_parContext {
	var p = new(Fun_parContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fun_par
	return p
}

func (*Fun_parContext) IsFun_parContext() {}

func NewFun_parContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_parContext {
	var p = new(Fun_parContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fun_par

	return p
}

func (s *Fun_parContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_parContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *Fun_parContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_parContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Fun_par() (localctx IFun_parContext) {
	localctx = NewFun_parContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SpartaParserRULE_fun_par)
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
		p.SetState(70)
		p.Match(SpartaParserT__2)
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserIDENTIFIER {
		{
			p.SetState(71)
			p.Namelist()
		}

	}
	{
		p.SetState(74)
		p.Match(SpartaParserT__3)
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
	p.EnterRule(localctx, 14, SpartaParserRULE_namelist)
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
		p.SetState(76)
		p.Match(SpartaParserIDENTIFIER)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__4 {
		{
			p.SetState(77)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(78)
			p.Match(SpartaParserIDENTIFIER)
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 16, SpartaParserRULE_return_stmt)

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
		p.Match(SpartaParserT__5)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(85)
			p.Postfix_expr()
		}

	}

	return localctx
}

// IFuncall_stmtContext is an interface to support dynamic dispatch.
type IFuncall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncall_stmtContext differentiates from other interfaces.
	IsFuncall_stmtContext()
}

type Funcall_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncall_stmtContext() *Funcall_stmtContext {
	var p = new(Funcall_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_funcall_stmt
	return p
}

func (*Funcall_stmtContext) IsFuncall_stmtContext() {}

func NewFuncall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcall_stmtContext {
	var p = new(Funcall_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_funcall_stmt

	return p
}

func (s *Funcall_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcall_stmtContext) Funcall_expr() IFuncall_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncall_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncall_exprContext)
}

func (s *Funcall_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcall_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Funcall_stmt() (localctx IFuncall_stmtContext) {
	localctx = NewFuncall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SpartaParserRULE_funcall_stmt)

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
		p.SetState(88)
		p.Funcall_expr()
	}

	return localctx
}

// IFuncall_exprContext is an interface to support dynamic dispatch.
type IFuncall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncall_exprContext differentiates from other interfaces.
	IsFuncall_exprContext()
}

type Funcall_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncall_exprContext() *Funcall_exprContext {
	var p = new(Funcall_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_funcall_expr
	return p
}

func (*Funcall_exprContext) IsFuncall_exprContext() {}

func NewFuncall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcall_exprContext {
	var p = new(Funcall_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_funcall_expr

	return p
}

func (s *Funcall_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcall_exprContext) Fun_name() IFun_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFun_nameContext)
}

func (s *Funcall_exprContext) Arg_expr() IArg_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_exprContext)
}

func (s *Funcall_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcall_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Funcall_expr() (localctx IFuncall_exprContext) {
	localctx = NewFuncall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SpartaParserRULE_funcall_expr)

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
		p.SetState(90)
		p.Fun_name()
	}
	{
		p.SetState(91)
		p.Arg_expr()
	}

	return localctx
}

// IArg_exprContext is an interface to support dynamic dispatch.
type IArg_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArg_exprContext differentiates from other interfaces.
	IsArg_exprContext()
}

type Arg_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_exprContext() *Arg_exprContext {
	var p = new(Arg_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_arg_expr
	return p
}

func (*Arg_exprContext) IsArg_exprContext() {}

func NewArg_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_exprContext {
	var p = new(Arg_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_arg_expr

	return p
}

func (s *Arg_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_exprContext) Arg_list() IArg_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArg_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Arg_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Arg_expr() (localctx IArg_exprContext) {
	localctx = NewArg_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SpartaParserRULE_arg_expr)
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
		p.SetState(93)
		p.Match(SpartaParserT__2)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__2)|(1<<SpartaParserT__12)|(1<<SpartaParserSTRING)|(1<<SpartaParserNUMBER_LITERAL)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(94)
			p.Arg_list()
		}

	}
	{
		p.SetState(97)
		p.Match(SpartaParserT__3)
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

func (s *Arg_listContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *Arg_listContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SpartaParserRULE_arg_list)
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
		p.SetState(99)
		p.Arg()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__4 {
		{
			p.SetState(100)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(101)
			p.Arg()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Postfix_expr() IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SpartaParserRULE_arg)

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
		p.SetState(107)
		p.Postfix_expr()
	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) AllPostfix_expr() []IPostfix_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem())
	var tst = make([]IPostfix_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPostfix_exprContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Postfix_expr(i int) IPostfix_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfix_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPostfix_exprContext)
}

func (s *If_stmtContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SpartaParserRULE_if_stmt)
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
		p.SetState(109)
		p.Match(SpartaParserT__6)
	}
	{
		p.SetState(110)
		p.Postfix_expr()
	}
	{
		p.SetState(111)
		p.Block()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__7 {
		{
			p.SetState(112)
			p.Match(SpartaParserT__7)
		}
		{
			p.SetState(113)
			p.Postfix_expr()
		}
		{
			p.SetState(114)
			p.Block()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__8 {
		{
			p.SetState(121)
			p.Match(SpartaParserT__8)
		}
		{
			p.SetState(122)
			p.Block()
		}

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
	p.EnterRule(localctx, 30, SpartaParserRULE_block)
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
		p.SetState(125)
		p.Match(SpartaParserT__9)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__5)|(1<<SpartaParserT__6)|(1<<SpartaParserIDENTIFIER))) != 0 {
		{
			p.SetState(126)
			p.Stmt()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(SpartaParserT__10)
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

func (s *Postfix_exprContext) Arith_expr() IArith_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArith_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArith_exprContext)
}

func (s *Postfix_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Postfix_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Postfix_expr() (localctx IPostfix_exprContext) {
	localctx = NewPostfix_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SpartaParserRULE_postfix_expr)

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

func (p *SpartaParser) Arith_expr() (localctx IArith_exprContext) {
	localctx = NewArith_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SpartaParserRULE_arith_expr)
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
		p.SetState(136)
		p.Term()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__11 || _la == SpartaParserT__12 {
		{
			p.SetState(137)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__11 || _la == SpartaParserT__12) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(138)
			p.Term()
		}

		p.SetState(143)
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

func (p *SpartaParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_term)
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
		p.SetState(144)
		p.Factor()
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15))) != 0 {
		{
			p.SetState(145)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__13)|(1<<SpartaParserT__14)|(1<<SpartaParserT__15))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.Factor()
		}

		p.SetState(151)
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

func (s *FactorContext) Atom_expr() IAtom_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_exprContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_factor)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__12 {
		{
			p.SetState(152)
			p.Match(SpartaParserT__12)
		}

	}
	{
		p.SetState(155)
		p.Atom_expr()
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

func (s *Atom_exprContext) Funcall_expr() IFuncall_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncall_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncall_exprContext)
}

func (s *Atom_exprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
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

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(SpartaParserT__2)
		}
		{
			p.SetState(158)
			p.Postfix_expr()
		}
		{
			p.SetState(159)
			p.Match(SpartaParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Funcall_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.Match(SpartaParserSTRING)
		}

	}

	return localctx
}
