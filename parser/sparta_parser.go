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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 234,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 75, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 87, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9,
	100, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 107, 10, 10, 12, 10,
	14, 10, 110, 11, 10, 3, 11, 3, 11, 5, 11, 114, 10, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 123, 10, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 15, 7, 15, 130, 10, 15, 12, 15, 14, 15, 133, 11, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 145,
	10, 17, 12, 17, 14, 17, 148, 11, 17, 3, 17, 3, 17, 5, 17, 152, 10, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 7,
	19, 164, 10, 19, 12, 19, 14, 19, 167, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 181, 10,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 7, 25, 188, 10, 25, 12, 25, 14,
	25, 191, 11, 25, 3, 26, 3, 26, 3, 26, 7, 26, 196, 10, 26, 12, 26, 14, 26,
	199, 11, 26, 3, 27, 5, 27, 202, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 224, 10, 28, 3, 29, 3, 29, 3,
	29, 7, 29, 229, 10, 29, 12, 29, 14, 29, 232, 11, 29, 3, 29, 2, 2, 30, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 2, 5, 3, 2, 20, 26, 3, 2, 27, 28, 3, 2,
	29, 31, 2, 234, 2, 61, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 76, 3, 2, 2, 2,
	8, 86, 3, 2, 2, 2, 10, 88, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14, 94, 3, 2,
	2, 2, 16, 97, 3, 2, 2, 2, 18, 103, 3, 2, 2, 2, 20, 111, 3, 2, 2, 2, 22,
	115, 3, 2, 2, 2, 24, 117, 3, 2, 2, 2, 26, 120, 3, 2, 2, 2, 28, 126, 3,
	2, 2, 2, 30, 134, 3, 2, 2, 2, 32, 136, 3, 2, 2, 2, 34, 153, 3, 2, 2, 2,
	36, 161, 3, 2, 2, 2, 38, 170, 3, 2, 2, 2, 40, 172, 3, 2, 2, 2, 42, 174,
	3, 2, 2, 2, 44, 176, 3, 2, 2, 2, 46, 182, 3, 2, 2, 2, 48, 184, 3, 2, 2,
	2, 50, 192, 3, 2, 2, 2, 52, 201, 3, 2, 2, 2, 54, 223, 3, 2, 2, 2, 56, 225,
	3, 2, 2, 2, 58, 60, 5, 4, 3, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2,
	61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 64, 65, 7, 2, 2, 3, 65, 3, 3, 2, 2, 2, 66, 75, 5, 6, 4, 2, 67,
	75, 5, 10, 6, 2, 68, 75, 5, 20, 11, 2, 69, 75, 5, 22, 12, 2, 70, 75, 5,
	32, 17, 2, 71, 75, 5, 34, 18, 2, 72, 75, 5, 38, 20, 2, 73, 75, 5, 40, 21,
	2, 74, 66, 3, 2, 2, 2, 74, 67, 3, 2, 2, 2, 74, 68, 3, 2, 2, 2, 74, 69,
	3, 2, 2, 2, 74, 70, 3, 2, 2, 2, 74, 71, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	74, 73, 3, 2, 2, 2, 75, 5, 3, 2, 2, 2, 76, 77, 5, 8, 5, 2, 77, 78, 7, 3,
	2, 2, 78, 79, 5, 42, 22, 2, 79, 7, 3, 2, 2, 2, 80, 87, 7, 35, 2, 2, 81,
	82, 7, 35, 2, 2, 82, 83, 7, 4, 2, 2, 83, 84, 5, 42, 22, 2, 84, 85, 7, 5,
	2, 2, 85, 87, 3, 2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 81, 3, 2, 2, 2, 87, 9,
	3, 2, 2, 2, 88, 89, 7, 6, 2, 2, 89, 90, 5, 12, 7, 2, 90, 91, 5, 14, 8,
	2, 91, 11, 3, 2, 2, 2, 92, 93, 7, 35, 2, 2, 93, 13, 3, 2, 2, 2, 94, 95,
	5, 16, 9, 2, 95, 96, 5, 36, 19, 2, 96, 15, 3, 2, 2, 2, 97, 99, 7, 7, 2,
	2, 98, 100, 5, 18, 10, 2, 99, 98, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100,
	101, 3, 2, 2, 2, 101, 102, 7, 8, 2, 2, 102, 17, 3, 2, 2, 2, 103, 108, 7,
	35, 2, 2, 104, 105, 7, 9, 2, 2, 105, 107, 7, 35, 2, 2, 106, 104, 3, 2,
	2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2,
	109, 19, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 113, 7, 10, 2, 2, 112,
	114, 5, 42, 22, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 21,
	3, 2, 2, 2, 115, 116, 5, 24, 13, 2, 116, 23, 3, 2, 2, 2, 117, 118, 5, 12,
	7, 2, 118, 119, 5, 26, 14, 2, 119, 25, 3, 2, 2, 2, 120, 122, 7, 7, 2, 2,
	121, 123, 5, 28, 15, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 125, 7, 8, 2, 2, 125, 27, 3, 2, 2, 2, 126, 131, 5,
	30, 16, 2, 127, 128, 7, 9, 2, 2, 128, 130, 5, 30, 16, 2, 129, 127, 3, 2,
	2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2,
	132, 29, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 135, 5, 42, 22, 2, 135,
	31, 3, 2, 2, 2, 136, 137, 7, 11, 2, 2, 137, 138, 5, 42, 22, 2, 138, 146,
	5, 36, 19, 2, 139, 140, 7, 12, 2, 2, 140, 141, 7, 11, 2, 2, 141, 142, 5,
	42, 22, 2, 142, 143, 5, 36, 19, 2, 143, 145, 3, 2, 2, 2, 144, 139, 3, 2,
	2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2,
	147, 151, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 7, 12, 2, 2, 150,
	152, 5, 36, 19, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 33,
	3, 2, 2, 2, 153, 154, 7, 13, 2, 2, 154, 155, 7, 35, 2, 2, 155, 156, 7,
	14, 2, 2, 156, 157, 5, 42, 22, 2, 157, 158, 7, 15, 2, 2, 158, 159, 5, 42,
	22, 2, 159, 160, 5, 36, 19, 2, 160, 35, 3, 2, 2, 2, 161, 165, 7, 16, 2,
	2, 162, 164, 5, 4, 3, 2, 163, 162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 165,
	3, 2, 2, 2, 168, 169, 7, 17, 2, 2, 169, 37, 3, 2, 2, 2, 170, 171, 7, 18,
	2, 2, 171, 39, 3, 2, 2, 2, 172, 173, 7, 19, 2, 2, 173, 41, 3, 2, 2, 2,
	174, 175, 5, 44, 23, 2, 175, 43, 3, 2, 2, 2, 176, 180, 5, 48, 25, 2, 177,
	178, 5, 46, 24, 2, 178, 179, 5, 48, 25, 2, 179, 181, 3, 2, 2, 2, 180, 177,
	3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 45, 3, 2, 2, 2, 182, 183, 9, 2,
	2, 2, 183, 47, 3, 2, 2, 2, 184, 189, 5, 50, 26, 2, 185, 186, 9, 3, 2, 2,
	186, 188, 5, 50, 26, 2, 187, 185, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189,
	187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 49, 3, 2, 2, 2, 191, 189, 3,
	2, 2, 2, 192, 197, 5, 52, 27, 2, 193, 194, 9, 4, 2, 2, 194, 196, 5, 52,
	27, 2, 195, 193, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 51, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 202,
	7, 28, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 204, 5, 54, 28, 2, 204, 53, 3, 2, 2, 2, 205, 206, 7, 7, 2, 2,
	206, 207, 5, 42, 22, 2, 207, 208, 7, 8, 2, 2, 208, 224, 3, 2, 2, 2, 209,
	210, 7, 4, 2, 2, 210, 211, 5, 56, 29, 2, 211, 212, 7, 5, 2, 2, 212, 224,
	3, 2, 2, 2, 213, 224, 5, 24, 13, 2, 214, 224, 7, 35, 2, 2, 215, 216, 7,
	35, 2, 2, 216, 217, 7, 4, 2, 2, 217, 218, 5, 42, 22, 2, 218, 219, 7, 5,
	2, 2, 219, 224, 3, 2, 2, 2, 220, 224, 7, 33, 2, 2, 221, 224, 7, 34, 2,
	2, 222, 224, 7, 32, 2, 2, 223, 205, 3, 2, 2, 2, 223, 209, 3, 2, 2, 2, 223,
	213, 3, 2, 2, 2, 223, 214, 3, 2, 2, 2, 223, 215, 3, 2, 2, 2, 223, 220,
	3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 55, 3, 2,
	2, 2, 225, 230, 5, 42, 22, 2, 226, 227, 7, 9, 2, 2, 227, 229, 5, 42, 22,
	2, 228, 226, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 57, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 19, 61, 74,
	86, 99, 108, 113, 122, 131, 146, 151, 165, 180, 189, 197, 201, 223, 230,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'['", "']'", "'fun'", "'('", "')'", "','", "'return'", "'if'",
	"'else'", "'for'", "'in'", "'to'", "'{'", "'}'", "'break'", "'continue'",
	"'<'", "'>'", "'=='", "'>='", "'<='", "'<>'", "'!='", "'+'", "'-'", "'*'",
	"'/'", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INTEGER_LITERAL",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt", "assign_stmt", "left_side", "fundef_stmt", "fun_name",
	"fun_body", "fun_par", "namelist", "return_stmt", "funcall_stmt", "funcall_expr",
	"arg_expr", "arg_list", "arg", "if_stmt", "for_stmt", "block", "break_stmt",
	"continue_stmt", "test", "compare_expr", "comp_op", "arith_expr", "term",
	"factor", "atom_expr", "test_list",
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
	SpartaParserEOF             = antlr.TokenEOF
	SpartaParserT__0            = 1
	SpartaParserT__1            = 2
	SpartaParserT__2            = 3
	SpartaParserT__3            = 4
	SpartaParserT__4            = 5
	SpartaParserT__5            = 6
	SpartaParserT__6            = 7
	SpartaParserT__7            = 8
	SpartaParserT__8            = 9
	SpartaParserT__9            = 10
	SpartaParserT__10           = 11
	SpartaParserT__11           = 12
	SpartaParserT__12           = 13
	SpartaParserT__13           = 14
	SpartaParserT__14           = 15
	SpartaParserT__15           = 16
	SpartaParserT__16           = 17
	SpartaParserT__17           = 18
	SpartaParserT__18           = 19
	SpartaParserT__19           = 20
	SpartaParserT__20           = 21
	SpartaParserT__21           = 22
	SpartaParserT__22           = 23
	SpartaParserT__23           = 24
	SpartaParserT__24           = 25
	SpartaParserT__25           = 26
	SpartaParserT__26           = 27
	SpartaParserT__27           = 28
	SpartaParserT__28           = 29
	SpartaParserSTRING          = 30
	SpartaParserINTEGER_LITERAL = 31
	SpartaParserNUMBER_LITERAL  = 32
	SpartaParserIDENTIFIER      = 33
	SpartaParserCOMMENT         = 34
	SpartaParserLINE_COMMENT    = 35
	SpartaParserWS              = 36
	SpartaParserSHEBANG         = 37
)

// SpartaParser rules.
const (
	SpartaParserRULE_program       = 0
	SpartaParserRULE_stmt          = 1
	SpartaParserRULE_assign_stmt   = 2
	SpartaParserRULE_left_side     = 3
	SpartaParserRULE_fundef_stmt   = 4
	SpartaParserRULE_fun_name      = 5
	SpartaParserRULE_fun_body      = 6
	SpartaParserRULE_fun_par       = 7
	SpartaParserRULE_namelist      = 8
	SpartaParserRULE_return_stmt   = 9
	SpartaParserRULE_funcall_stmt  = 10
	SpartaParserRULE_funcall_expr  = 11
	SpartaParserRULE_arg_expr      = 12
	SpartaParserRULE_arg_list      = 13
	SpartaParserRULE_arg           = 14
	SpartaParserRULE_if_stmt       = 15
	SpartaParserRULE_for_stmt      = 16
	SpartaParserRULE_block         = 17
	SpartaParserRULE_break_stmt    = 18
	SpartaParserRULE_continue_stmt = 19
	SpartaParserRULE_test          = 20
	SpartaParserRULE_compare_expr  = 21
	SpartaParserRULE_comp_op       = 22
	SpartaParserRULE_arith_expr    = 23
	SpartaParserRULE_term          = 24
	SpartaParserRULE_factor        = 25
	SpartaParserRULE_atom_expr     = 26
	SpartaParserRULE_test_list     = 27
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(SpartaParserT__3-4))|(1<<(SpartaParserT__7-4))|(1<<(SpartaParserT__8-4))|(1<<(SpartaParserT__10-4))|(1<<(SpartaParserT__15-4))|(1<<(SpartaParserT__16-4))|(1<<(SpartaParserIDENTIFIER-4)))) != 0 {
		{
			p.SetState(56)
			p.Stmt()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
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

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Break_stmt() IBreak_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *StmtContext) Continue_stmt() IContinue_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Fundef_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Return_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(67)
			p.Funcall_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)
			p.If_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(69)
			p.For_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(70)
			p.Break_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(71)
			p.Continue_stmt()
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

func (s *Assign_stmtContext) Left_side() ILeft_sideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeft_sideContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeft_sideContext)
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
		p.SetState(74)
		p.Left_side()
	}
	{
		p.SetState(75)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(76)
		p.Test()
	}

	return localctx
}

// ILeft_sideContext is an interface to support dynamic dispatch.
type ILeft_sideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeft_sideContext differentiates from other interfaces.
	IsLeft_sideContext()
}

type Left_sideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeft_sideContext() *Left_sideContext {
	var p = new(Left_sideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_left_side
	return p
}

func (*Left_sideContext) IsLeft_sideContext() {}

func NewLeft_sideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Left_sideContext {
	var p = new(Left_sideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_left_side

	return p
}

func (s *Left_sideContext) GetParser() antlr.Parser { return s.parser }

func (s *Left_sideContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Left_sideContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Left_sideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Left_sideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Left_side() (localctx ILeft_sideContext) {
	localctx = NewLeft_sideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpartaParserRULE_left_side)

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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(SpartaParserIDENTIFIER)
		}
		{
			p.SetState(80)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(81)
			p.Test()
		}
		{
			p.SetState(82)
			p.Match(SpartaParserT__2)
		}

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
	p.EnterRule(localctx, 8, SpartaParserRULE_fundef_stmt)

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
		p.Match(SpartaParserT__3)
	}
	{
		p.SetState(87)
		p.Fun_name()
	}
	{
		p.SetState(88)
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
	p.EnterRule(localctx, 10, SpartaParserRULE_fun_name)

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
	p.EnterRule(localctx, 12, SpartaParserRULE_fun_body)

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
		p.SetState(92)
		p.Fun_par()
	}
	{
		p.SetState(93)
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
	p.EnterRule(localctx, 14, SpartaParserRULE_fun_par)
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
		p.Match(SpartaParserT__4)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserIDENTIFIER {
		{
			p.SetState(96)
			p.Namelist()
		}

	}
	{
		p.SetState(99)
		p.Match(SpartaParserT__5)
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
		p.SetState(101)
		p.Match(SpartaParserIDENTIFIER)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(102)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(103)
			p.Match(SpartaParserIDENTIFIER)
		}

		p.SetState(108)
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

func (s *Return_stmtContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SpartaParserRULE_return_stmt)

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
		p.Match(SpartaParserT__7)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Test()
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
	p.EnterRule(localctx, 20, SpartaParserRULE_funcall_stmt)

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
		p.SetState(113)
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
	p.EnterRule(localctx, 22, SpartaParserRULE_funcall_expr)

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
		p.Fun_name()
	}
	{
		p.SetState(116)
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
	p.EnterRule(localctx, 24, SpartaParserRULE_arg_expr)
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
		p.SetState(118)
		p.Match(SpartaParserT__4)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(SpartaParserT__1-2))|(1<<(SpartaParserT__4-2))|(1<<(SpartaParserT__25-2))|(1<<(SpartaParserSTRING-2))|(1<<(SpartaParserINTEGER_LITERAL-2))|(1<<(SpartaParserNUMBER_LITERAL-2))|(1<<(SpartaParserIDENTIFIER-2)))) != 0 {
		{
			p.SetState(119)
			p.Arg_list()
		}

	}
	{
		p.SetState(122)
		p.Match(SpartaParserT__5)
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
	p.EnterRule(localctx, 26, SpartaParserRULE_arg_list)
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
		p.Arg()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(125)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(126)
			p.Arg()
		}

		p.SetState(131)
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

func (s *ArgContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SpartaParserRULE_arg)

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
		p.SetState(132)
		p.Test()
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

func (s *If_stmtContext) AllTest() []ITestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITestContext)(nil)).Elem())
	var tst = make([]ITestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITestContext)
		}
	}

	return tst
}

func (s *If_stmtContext) Test(i int) ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
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
	p.EnterRule(localctx, 30, SpartaParserRULE_if_stmt)
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
		p.SetState(134)
		p.Match(SpartaParserT__8)
	}
	{
		p.SetState(135)
		p.Test()
	}
	{
		p.SetState(136)
		p.Block()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(137)
				p.Match(SpartaParserT__9)
			}
			{
				p.SetState(138)
				p.Match(SpartaParserT__8)
			}
			{
				p.SetState(139)
				p.Test()
			}
			{
				p.SetState(140)
				p.Block()
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__9 {
		{
			p.SetState(147)
			p.Match(SpartaParserT__9)
		}
		{
			p.SetState(148)
			p.Block()
		}

	}

	return localctx
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_for_stmt
	return p
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *For_stmtContext) AllTest() []ITestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITestContext)(nil)).Elem())
	var tst = make([]ITestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITestContext)
		}
	}

	return tst
}

func (s *For_stmtContext) Test(i int) ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *For_stmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SpartaParserRULE_for_stmt)

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
		p.SetState(151)
		p.Match(SpartaParserT__10)
	}
	{
		p.SetState(152)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(153)
		p.Match(SpartaParserT__11)
	}
	{
		p.SetState(154)
		p.Test()
	}
	{
		p.SetState(155)
		p.Match(SpartaParserT__12)
	}
	{
		p.SetState(156)
		p.Test()
	}
	{
		p.SetState(157)
		p.Block()
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
	p.EnterRule(localctx, 34, SpartaParserRULE_block)
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
		p.SetState(159)
		p.Match(SpartaParserT__13)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(SpartaParserT__3-4))|(1<<(SpartaParserT__7-4))|(1<<(SpartaParserT__8-4))|(1<<(SpartaParserT__10-4))|(1<<(SpartaParserT__15-4))|(1<<(SpartaParserT__16-4))|(1<<(SpartaParserIDENTIFIER-4)))) != 0 {
		{
			p.SetState(160)
			p.Stmt()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(SpartaParserT__14)
	}

	return localctx
}

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_break_stmt
	return p
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Break_stmt() (localctx IBreak_stmtContext) {
	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_break_stmt)

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
		p.SetState(168)
		p.Match(SpartaParserT__15)
	}

	return localctx
}

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_continue_stmt
	return p
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Continue_stmt() (localctx IContinue_stmtContext) {
	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_continue_stmt)

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
		p.SetState(170)
		p.Match(SpartaParserT__16)
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

func (s *TestContext) Compare_expr() ICompare_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompare_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompare_exprContext)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_test)

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
		p.SetState(172)
		p.Compare_expr()
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
	p.EnterRule(localctx, 42, SpartaParserRULE_compare_expr)
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
		p.SetState(174)
		p.Arith_expr()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__17)|(1<<SpartaParserT__18)|(1<<SpartaParserT__19)|(1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23))) != 0 {
		{
			p.SetState(175)
			p.Comp_op()
		}
		{
			p.SetState(176)
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
	p.EnterRule(localctx, 44, SpartaParserRULE_comp_op)
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
		p.SetState(180)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__17)|(1<<SpartaParserT__18)|(1<<SpartaParserT__19)|(1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23))) != 0) {
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
	p.EnterRule(localctx, 46, SpartaParserRULE_arith_expr)
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
		p.SetState(182)
		p.Term()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__24 || _la == SpartaParserT__25 {
		{
			p.SetState(183)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__24 || _la == SpartaParserT__25) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(184)
			p.Term()
		}

		p.SetState(189)
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
	p.EnterRule(localctx, 48, SpartaParserRULE_term)
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
		p.SetState(190)
		p.Factor()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__26)|(1<<SpartaParserT__27)|(1<<SpartaParserT__28))) != 0 {
		{
			p.SetState(191)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__26)|(1<<SpartaParserT__27)|(1<<SpartaParserT__28))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(192)
			p.Factor()
		}

		p.SetState(197)
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
	p.EnterRule(localctx, 50, SpartaParserRULE_factor)
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
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__25 {
		{
			p.SetState(198)
			p.Match(SpartaParserT__25)
		}

	}
	{
		p.SetState(201)
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

func (s *Atom_exprContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Atom_exprContext) Test_list() ITest_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITest_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITest_listContext)
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

func (s *Atom_exprContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SpartaParserINTEGER_LITERAL, 0)
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
	p.EnterRule(localctx, 52, SpartaParserRULE_atom_expr)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(204)
			p.Test()
		}
		{
			p.SetState(205)
			p.Match(SpartaParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(208)
			p.Test_list()
		}
		{
			p.SetState(209)
			p.Match(SpartaParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.Funcall_expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)
			p.Match(SpartaParserIDENTIFIER)
		}

		{
			p.SetState(214)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(215)
			p.Test()
		}
		{
			p.SetState(216)
			p.Match(SpartaParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(218)
			p.Match(SpartaParserINTEGER_LITERAL)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(219)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(220)
			p.Match(SpartaParserSTRING)
		}

	}

	return localctx
}

// ITest_listContext is an interface to support dynamic dispatch.
type ITest_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTest_listContext differentiates from other interfaces.
	IsTest_listContext()
}

type Test_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTest_listContext() *Test_listContext {
	var p = new(Test_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_test_list
	return p
}

func (*Test_listContext) IsTest_listContext() {}

func NewTest_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Test_listContext {
	var p = new(Test_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_test_list

	return p
}

func (s *Test_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Test_listContext) AllTest() []ITestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITestContext)(nil)).Elem())
	var tst = make([]ITestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITestContext)
		}
	}

	return tst
}

func (s *Test_listContext) Test(i int) ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Test_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Test_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Test_list() (localctx ITest_listContext) {
	localctx = NewTest_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SpartaParserRULE_test_list)
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
		p.SetState(223)
		p.Test()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(224)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(225)
			p.Test()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
