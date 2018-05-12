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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 276,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2,
	77, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 90, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 102, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 115, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 7, 10, 122, 10, 10, 12, 10, 14, 10, 125, 11, 10, 3, 11, 3, 11, 5, 11,
	129, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 138,
	10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 7, 15, 145, 10, 15, 12, 15,
	14, 15, 148, 11, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 7, 17, 160, 10, 17, 12, 17, 14, 17, 163, 11, 17, 3, 17,
	3, 17, 5, 17, 167, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 7, 19, 179, 10, 19, 12, 19, 14, 19, 182, 11, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 199, 10, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 7, 26, 206, 10, 26, 12, 26, 14, 26, 209, 11, 26, 3, 27, 3,
	27, 3, 27, 7, 27, 214, 10, 27, 12, 27, 14, 27, 217, 11, 27, 3, 28, 5, 28,
	220, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 233, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 5, 31, 241, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 247, 10,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 7, 33, 254, 10, 33, 12, 33, 14,
	33, 257, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 271, 10, 36, 12, 36, 14, 36, 274, 11,
	36, 3, 36, 2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 2, 5, 3, 2, 21, 27, 3, 2, 28, 29, 3, 2, 30, 32, 2, 274, 2,
	75, 3, 2, 2, 2, 4, 89, 3, 2, 2, 2, 6, 91, 3, 2, 2, 2, 8, 101, 3, 2, 2,
	2, 10, 103, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 112,
	3, 2, 2, 2, 18, 118, 3, 2, 2, 2, 20, 126, 3, 2, 2, 2, 22, 130, 3, 2, 2,
	2, 24, 132, 3, 2, 2, 2, 26, 135, 3, 2, 2, 2, 28, 141, 3, 2, 2, 2, 30, 149,
	3, 2, 2, 2, 32, 151, 3, 2, 2, 2, 34, 168, 3, 2, 2, 2, 36, 176, 3, 2, 2,
	2, 38, 185, 3, 2, 2, 2, 40, 187, 3, 2, 2, 2, 42, 189, 3, 2, 2, 2, 44, 192,
	3, 2, 2, 2, 46, 194, 3, 2, 2, 2, 48, 200, 3, 2, 2, 2, 50, 202, 3, 2, 2,
	2, 52, 210, 3, 2, 2, 2, 54, 219, 3, 2, 2, 2, 56, 232, 3, 2, 2, 2, 58, 234,
	3, 2, 2, 2, 60, 238, 3, 2, 2, 2, 62, 244, 3, 2, 2, 2, 64, 250, 3, 2, 2,
	2, 66, 258, 3, 2, 2, 2, 68, 262, 3, 2, 2, 2, 70, 267, 3, 2, 2, 2, 72, 74,
	5, 4, 3, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7,
	2, 2, 3, 79, 3, 3, 2, 2, 2, 80, 90, 5, 6, 4, 2, 81, 90, 5, 10, 6, 2, 82,
	90, 5, 20, 11, 2, 83, 90, 5, 22, 12, 2, 84, 90, 5, 32, 17, 2, 85, 90, 5,
	34, 18, 2, 86, 90, 5, 38, 20, 2, 87, 90, 5, 40, 21, 2, 88, 90, 5, 42, 22,
	2, 89, 80, 3, 2, 2, 2, 89, 81, 3, 2, 2, 2, 89, 82, 3, 2, 2, 2, 89, 83,
	3, 2, 2, 2, 89, 84, 3, 2, 2, 2, 89, 85, 3, 2, 2, 2, 89, 86, 3, 2, 2, 2,
	89, 87, 3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 5, 3, 2, 2, 2, 91, 92, 5, 8,
	5, 2, 92, 93, 7, 3, 2, 2, 93, 94, 5, 44, 23, 2, 94, 7, 3, 2, 2, 2, 95,
	102, 7, 37, 2, 2, 96, 97, 7, 37, 2, 2, 97, 98, 7, 4, 2, 2, 98, 99, 5, 44,
	23, 2, 99, 100, 7, 5, 2, 2, 100, 102, 3, 2, 2, 2, 101, 95, 3, 2, 2, 2,
	101, 96, 3, 2, 2, 2, 102, 9, 3, 2, 2, 2, 103, 104, 7, 6, 2, 2, 104, 105,
	5, 12, 7, 2, 105, 106, 5, 14, 8, 2, 106, 11, 3, 2, 2, 2, 107, 108, 7, 37,
	2, 2, 108, 13, 3, 2, 2, 2, 109, 110, 5, 16, 9, 2, 110, 111, 5, 36, 19,
	2, 111, 15, 3, 2, 2, 2, 112, 114, 7, 7, 2, 2, 113, 115, 5, 18, 10, 2, 114,
	113, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117,
	7, 8, 2, 2, 117, 17, 3, 2, 2, 2, 118, 123, 7, 37, 2, 2, 119, 120, 7, 9,
	2, 2, 120, 122, 7, 37, 2, 2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2,
	123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 19, 3, 2, 2, 2, 125, 123,
	3, 2, 2, 2, 126, 128, 7, 10, 2, 2, 127, 129, 5, 44, 23, 2, 128, 127, 3,
	2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 21, 3, 2, 2, 2, 130, 131, 5, 24, 13,
	2, 131, 23, 3, 2, 2, 2, 132, 133, 5, 12, 7, 2, 133, 134, 5, 26, 14, 2,
	134, 25, 3, 2, 2, 2, 135, 137, 7, 7, 2, 2, 136, 138, 5, 28, 15, 2, 137,
	136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140,
	7, 8, 2, 2, 140, 27, 3, 2, 2, 2, 141, 146, 5, 30, 16, 2, 142, 143, 7, 9,
	2, 2, 143, 145, 5, 30, 16, 2, 144, 142, 3, 2, 2, 2, 145, 148, 3, 2, 2,
	2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 29, 3, 2, 2, 2, 148,
	146, 3, 2, 2, 2, 149, 150, 5, 44, 23, 2, 150, 31, 3, 2, 2, 2, 151, 152,
	7, 11, 2, 2, 152, 153, 5, 44, 23, 2, 153, 161, 5, 36, 19, 2, 154, 155,
	7, 12, 2, 2, 155, 156, 7, 11, 2, 2, 156, 157, 5, 44, 23, 2, 157, 158, 5,
	36, 19, 2, 158, 160, 3, 2, 2, 2, 159, 154, 3, 2, 2, 2, 160, 163, 3, 2,
	2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 166, 3, 2, 2, 2,
	163, 161, 3, 2, 2, 2, 164, 165, 7, 12, 2, 2, 165, 167, 5, 36, 19, 2, 166,
	164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 33, 3, 2, 2, 2, 168, 169, 7,
	13, 2, 2, 169, 170, 7, 37, 2, 2, 170, 171, 7, 14, 2, 2, 171, 172, 5, 44,
	23, 2, 172, 173, 7, 15, 2, 2, 173, 174, 5, 44, 23, 2, 174, 175, 5, 36,
	19, 2, 175, 35, 3, 2, 2, 2, 176, 180, 7, 16, 2, 2, 177, 179, 5, 4, 3, 2,
	178, 177, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180,
	181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 184,
	7, 17, 2, 2, 184, 37, 3, 2, 2, 2, 185, 186, 7, 18, 2, 2, 186, 39, 3, 2,
	2, 2, 187, 188, 7, 19, 2, 2, 188, 41, 3, 2, 2, 2, 189, 190, 7, 20, 2, 2,
	190, 191, 7, 37, 2, 2, 191, 43, 3, 2, 2, 2, 192, 193, 5, 46, 24, 2, 193,
	45, 3, 2, 2, 2, 194, 198, 5, 50, 26, 2, 195, 196, 5, 48, 25, 2, 196, 197,
	5, 50, 26, 2, 197, 199, 3, 2, 2, 2, 198, 195, 3, 2, 2, 2, 198, 199, 3,
	2, 2, 2, 199, 47, 3, 2, 2, 2, 200, 201, 9, 2, 2, 2, 201, 49, 3, 2, 2, 2,
	202, 207, 5, 52, 27, 2, 203, 204, 9, 3, 2, 2, 204, 206, 5, 52, 27, 2, 205,
	203, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208,
	3, 2, 2, 2, 208, 51, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 215, 5, 54,
	28, 2, 211, 212, 9, 4, 2, 2, 212, 214, 5, 54, 28, 2, 213, 211, 3, 2, 2,
	2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216,
	53, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 220, 7, 29, 2, 2, 219, 218,
	3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 5, 56,
	29, 2, 222, 55, 3, 2, 2, 2, 223, 233, 5, 58, 30, 2, 224, 233, 5, 60, 31,
	2, 225, 233, 5, 62, 32, 2, 226, 233, 5, 24, 13, 2, 227, 233, 5, 68, 35,
	2, 228, 233, 7, 37, 2, 2, 229, 233, 7, 35, 2, 2, 230, 233, 7, 36, 2, 2,
	231, 233, 7, 34, 2, 2, 232, 223, 3, 2, 2, 2, 232, 224, 3, 2, 2, 2, 232,
	225, 3, 2, 2, 2, 232, 226, 3, 2, 2, 2, 232, 227, 3, 2, 2, 2, 232, 228,
	3, 2, 2, 2, 232, 229, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 231, 3, 2,
	2, 2, 233, 57, 3, 2, 2, 2, 234, 235, 7, 7, 2, 2, 235, 236, 5, 44, 23, 2,
	236, 237, 7, 8, 2, 2, 237, 59, 3, 2, 2, 2, 238, 240, 7, 4, 2, 2, 239, 241,
	5, 70, 36, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 3,
	2, 2, 2, 242, 243, 7, 5, 2, 2, 243, 61, 3, 2, 2, 2, 244, 246, 7, 16, 2,
	2, 245, 247, 5, 64, 33, 2, 246, 245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2,
	247, 248, 3, 2, 2, 2, 248, 249, 7, 17, 2, 2, 249, 63, 3, 2, 2, 2, 250,
	255, 5, 66, 34, 2, 251, 252, 7, 9, 2, 2, 252, 254, 5, 66, 34, 2, 253, 251,
	3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2,
	2, 2, 256, 65, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 259, 5, 44, 23, 2,
	259, 260, 7, 33, 2, 2, 260, 261, 5, 44, 23, 2, 261, 67, 3, 2, 2, 2, 262,
	263, 7, 37, 2, 2, 263, 264, 7, 4, 2, 2, 264, 265, 5, 44, 23, 2, 265, 266,
	7, 5, 2, 2, 266, 69, 3, 2, 2, 2, 267, 272, 5, 44, 23, 2, 268, 269, 7, 9,
	2, 2, 269, 271, 5, 44, 23, 2, 270, 268, 3, 2, 2, 2, 271, 274, 3, 2, 2,
	2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 71, 3, 2, 2, 2, 274,
	272, 3, 2, 2, 2, 22, 75, 89, 101, 114, 123, 128, 137, 146, 161, 166, 180,
	198, 207, 215, 219, 232, 240, 246, 255, 272,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'['", "']'", "'fun'", "'('", "')'", "','", "'return'", "'if'",
	"'else'", "'for'", "'in'", "'to'", "'{'", "'}'", "'break'", "'continue'",
	"'import'", "'<'", "'>'", "'=='", "'>='", "'<='", "'<>'", "'!='", "'+'",
	"'-'", "'*'", "'/'", "'%'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INTEGER_LITERAL",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt", "assign_stmt", "left_side", "fundef_stmt", "fun_name",
	"fun_body", "fun_par", "namelist", "return_stmt", "funcall_stmt", "funcall_expr",
	"arg_expr", "arg_list", "arg", "if_stmt", "for_stmt", "block", "break_stmt",
	"continue_stmt", "import_stmt", "test", "compare_expr", "comp_op", "arith_expr",
	"term", "factor", "atom_expr", "bracket_expr", "list_literal", "map_literal",
	"entry_list", "entry", "table_index", "test_list",
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
	SpartaParserT__29           = 30
	SpartaParserT__30           = 31
	SpartaParserSTRING          = 32
	SpartaParserINTEGER_LITERAL = 33
	SpartaParserNUMBER_LITERAL  = 34
	SpartaParserIDENTIFIER      = 35
	SpartaParserCOMMENT         = 36
	SpartaParserLINE_COMMENT    = 37
	SpartaParserWS              = 38
	SpartaParserSHEBANG         = 39
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
	SpartaParserRULE_import_stmt   = 20
	SpartaParserRULE_test          = 21
	SpartaParserRULE_compare_expr  = 22
	SpartaParserRULE_comp_op       = 23
	SpartaParserRULE_arith_expr    = 24
	SpartaParserRULE_term          = 25
	SpartaParserRULE_factor        = 26
	SpartaParserRULE_atom_expr     = 27
	SpartaParserRULE_bracket_expr  = 28
	SpartaParserRULE_list_literal  = 29
	SpartaParserRULE_map_literal   = 30
	SpartaParserRULE_entry_list    = 31
	SpartaParserRULE_entry         = 32
	SpartaParserRULE_table_index   = 33
	SpartaParserRULE_test_list     = 34
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(SpartaParserT__3-4))|(1<<(SpartaParserT__7-4))|(1<<(SpartaParserT__8-4))|(1<<(SpartaParserT__10-4))|(1<<(SpartaParserT__15-4))|(1<<(SpartaParserT__16-4))|(1<<(SpartaParserT__17-4))|(1<<(SpartaParserIDENTIFIER-4)))) != 0 {
		{
			p.SetState(70)
			p.Stmt()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
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

func (s *StmtContext) Import_stmt() IImport_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImport_stmtContext)
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

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Fundef_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Return_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.Funcall_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(82)
			p.If_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(83)
			p.For_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(84)
			p.Break_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(85)
			p.Continue_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(86)
			p.Import_stmt()
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
		p.SetState(89)
		p.Left_side()
	}
	{
		p.SetState(90)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(91)
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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(SpartaParserIDENTIFIER)
		}
		{
			p.SetState(95)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(96)
			p.Test()
		}
		{
			p.SetState(97)
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
		p.SetState(101)
		p.Match(SpartaParserT__3)
	}
	{
		p.SetState(102)
		p.Fun_name()
	}
	{
		p.SetState(103)
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
		p.SetState(105)
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
		p.SetState(107)
		p.Fun_par()
	}
	{
		p.SetState(108)
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
		p.SetState(110)
		p.Match(SpartaParserT__4)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserIDENTIFIER {
		{
			p.SetState(111)
			p.Namelist()
		}

	}
	{
		p.SetState(114)
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
		p.SetState(116)
		p.Match(SpartaParserIDENTIFIER)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(117)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(118)
			p.Match(SpartaParserIDENTIFIER)
		}

		p.SetState(123)
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
		p.SetState(124)
		p.Match(SpartaParserT__7)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(125)
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
		p.SetState(128)
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
		p.SetState(130)
		p.Fun_name()
	}
	{
		p.SetState(131)
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
		p.SetState(133)
		p.Match(SpartaParserT__4)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__26))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SpartaParserSTRING-32))|(1<<(SpartaParserINTEGER_LITERAL-32))|(1<<(SpartaParserNUMBER_LITERAL-32))|(1<<(SpartaParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(134)
			p.Arg_list()
		}

	}
	{
		p.SetState(137)
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
		p.SetState(139)
		p.Arg()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(140)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(141)
			p.Arg()
		}

		p.SetState(146)
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
		p.SetState(147)
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
		p.SetState(149)
		p.Match(SpartaParserT__8)
	}
	{
		p.SetState(150)
		p.Test()
	}
	{
		p.SetState(151)
		p.Block()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(152)
				p.Match(SpartaParserT__9)
			}
			{
				p.SetState(153)
				p.Match(SpartaParserT__8)
			}
			{
				p.SetState(154)
				p.Test()
			}
			{
				p.SetState(155)
				p.Block()
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__9 {
		{
			p.SetState(162)
			p.Match(SpartaParserT__9)
		}
		{
			p.SetState(163)
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
		p.SetState(166)
		p.Match(SpartaParserT__10)
	}
	{
		p.SetState(167)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(168)
		p.Match(SpartaParserT__11)
	}
	{
		p.SetState(169)
		p.Test()
	}
	{
		p.SetState(170)
		p.Match(SpartaParserT__12)
	}
	{
		p.SetState(171)
		p.Test()
	}
	{
		p.SetState(172)
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
		p.SetState(174)
		p.Match(SpartaParserT__13)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(SpartaParserT__3-4))|(1<<(SpartaParserT__7-4))|(1<<(SpartaParserT__8-4))|(1<<(SpartaParserT__10-4))|(1<<(SpartaParserT__15-4))|(1<<(SpartaParserT__16-4))|(1<<(SpartaParserT__17-4))|(1<<(SpartaParserIDENTIFIER-4)))) != 0 {
		{
			p.SetState(175)
			p.Stmt()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
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
		p.SetState(183)
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
		p.SetState(185)
		p.Match(SpartaParserT__16)
	}

	return localctx
}

// IImport_stmtContext is an interface to support dynamic dispatch.
type IImport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_stmtContext differentiates from other interfaces.
	IsImport_stmtContext()
}

type Import_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_stmtContext() *Import_stmtContext {
	var p = new(Import_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_import_stmt
	return p
}

func (*Import_stmtContext) IsImport_stmtContext() {}

func NewImport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_stmtContext {
	var p = new(Import_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_import_stmt

	return p
}

func (s *Import_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_stmtContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Import_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Import_stmt() (localctx IImport_stmtContext) {
	localctx = NewImport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_import_stmt)

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
		p.SetState(187)
		p.Match(SpartaParserT__17)
	}
	{
		p.SetState(188)
		p.Match(SpartaParserIDENTIFIER)
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
	p.EnterRule(localctx, 42, SpartaParserRULE_test)

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
	p.EnterRule(localctx, 44, SpartaParserRULE_compare_expr)
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
		p.SetState(192)
		p.Arith_expr()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__18)|(1<<SpartaParserT__19)|(1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24))) != 0 {
		{
			p.SetState(193)
			p.Comp_op()
		}
		{
			p.SetState(194)
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
	p.EnterRule(localctx, 46, SpartaParserRULE_comp_op)
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
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__18)|(1<<SpartaParserT__19)|(1<<SpartaParserT__20)|(1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24))) != 0) {
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
	p.EnterRule(localctx, 48, SpartaParserRULE_arith_expr)
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
		p.SetState(200)
		p.Term()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__25 || _la == SpartaParserT__26 {
		{
			p.SetState(201)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__25 || _la == SpartaParserT__26) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(202)
			p.Term()
		}

		p.SetState(207)
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
	p.EnterRule(localctx, 50, SpartaParserRULE_term)
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
		p.SetState(208)
		p.Factor()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__27)|(1<<SpartaParserT__28)|(1<<SpartaParserT__29))) != 0 {
		{
			p.SetState(209)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__27)|(1<<SpartaParserT__28)|(1<<SpartaParserT__29))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(210)
			p.Factor()
		}

		p.SetState(215)
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
	p.EnterRule(localctx, 52, SpartaParserRULE_factor)
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
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__26 {
		{
			p.SetState(216)
			p.Match(SpartaParserT__26)
		}

	}
	{
		p.SetState(219)
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

func (s *Atom_exprContext) Bracket_expr() IBracket_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracket_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracket_exprContext)
}

func (s *Atom_exprContext) List_literal() IList_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_literalContext)
}

func (s *Atom_exprContext) Map_literal() IMap_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_literalContext)
}

func (s *Atom_exprContext) Funcall_expr() IFuncall_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncall_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncall_exprContext)
}

func (s *Atom_exprContext) Table_index() ITable_indexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_indexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_indexContext)
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
	p.EnterRule(localctx, 54, SpartaParserRULE_atom_expr)

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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Bracket_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.List_literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.Map_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
			p.Funcall_expr()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(225)
			p.Table_index()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(226)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(227)
			p.Match(SpartaParserINTEGER_LITERAL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(228)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(229)
			p.Match(SpartaParserSTRING)
		}

	}

	return localctx
}

// IBracket_exprContext is an interface to support dynamic dispatch.
type IBracket_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracket_exprContext differentiates from other interfaces.
	IsBracket_exprContext()
}

type Bracket_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracket_exprContext() *Bracket_exprContext {
	var p = new(Bracket_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_bracket_expr
	return p
}

func (*Bracket_exprContext) IsBracket_exprContext() {}

func NewBracket_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bracket_exprContext {
	var p = new(Bracket_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_bracket_expr

	return p
}

func (s *Bracket_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Bracket_exprContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Bracket_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bracket_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Bracket_expr() (localctx IBracket_exprContext) {
	localctx = NewBracket_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SpartaParserRULE_bracket_expr)

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
		p.SetState(232)
		p.Match(SpartaParserT__4)
	}
	{
		p.SetState(233)
		p.Test()
	}
	{
		p.SetState(234)
		p.Match(SpartaParserT__5)
	}

	return localctx
}

// IList_literalContext is an interface to support dynamic dispatch.
type IList_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_literalContext differentiates from other interfaces.
	IsList_literalContext()
}

type List_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_literalContext() *List_literalContext {
	var p = new(List_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_list_literal
	return p
}

func (*List_literalContext) IsList_literalContext() {}

func NewList_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_literalContext {
	var p = new(List_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_list_literal

	return p
}

func (s *List_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *List_literalContext) Test_list() ITest_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITest_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITest_listContext)
}

func (s *List_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) List_literal() (localctx IList_literalContext) {
	localctx = NewList_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SpartaParserRULE_list_literal)
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
		p.SetState(236)
		p.Match(SpartaParserT__1)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__26))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SpartaParserSTRING-32))|(1<<(SpartaParserINTEGER_LITERAL-32))|(1<<(SpartaParserNUMBER_LITERAL-32))|(1<<(SpartaParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(237)
			p.Test_list()
		}

	}
	{
		p.SetState(240)
		p.Match(SpartaParserT__2)
	}

	return localctx
}

// IMap_literalContext is an interface to support dynamic dispatch.
type IMap_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_literalContext differentiates from other interfaces.
	IsMap_literalContext()
}

type Map_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_literalContext() *Map_literalContext {
	var p = new(Map_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_map_literal
	return p
}

func (*Map_literalContext) IsMap_literalContext() {}

func NewMap_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_literalContext {
	var p = new(Map_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_map_literal

	return p
}

func (s *Map_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_literalContext) Entry_list() IEntry_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntry_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntry_listContext)
}

func (s *Map_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Map_literal() (localctx IMap_literalContext) {
	localctx = NewMap_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SpartaParserRULE_map_literal)
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
		p.SetState(242)
		p.Match(SpartaParserT__13)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__26))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(SpartaParserSTRING-32))|(1<<(SpartaParserINTEGER_LITERAL-32))|(1<<(SpartaParserNUMBER_LITERAL-32))|(1<<(SpartaParserIDENTIFIER-32)))) != 0) {
		{
			p.SetState(243)
			p.Entry_list()
		}

	}
	{
		p.SetState(246)
		p.Match(SpartaParserT__14)
	}

	return localctx
}

// IEntry_listContext is an interface to support dynamic dispatch.
type IEntry_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntry_listContext differentiates from other interfaces.
	IsEntry_listContext()
}

type Entry_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntry_listContext() *Entry_listContext {
	var p = new(Entry_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_entry_list
	return p
}

func (*Entry_listContext) IsEntry_listContext() {}

func NewEntry_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entry_listContext {
	var p = new(Entry_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_entry_list

	return p
}

func (s *Entry_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Entry_listContext) AllEntry() []IEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntryContext)(nil)).Elem())
	var tst = make([]IEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntryContext)
		}
	}

	return tst
}

func (s *Entry_listContext) Entry(i int) IEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
}

func (s *Entry_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entry_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Entry_list() (localctx IEntry_listContext) {
	localctx = NewEntry_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SpartaParserRULE_entry_list)
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
		p.SetState(248)
		p.Entry()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(249)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(250)
			p.Entry()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_entry
	return p
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) AllTest() []ITestContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITestContext)(nil)).Elem())
	var tst = make([]ITestContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITestContext)
		}
	}

	return tst
}

func (s *EntryContext) Test(i int) ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SpartaParserRULE_entry)

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
		p.SetState(256)
		p.Test()
	}
	{
		p.SetState(257)
		p.Match(SpartaParserT__30)
	}
	{
		p.SetState(258)
		p.Test()
	}

	return localctx
}

// ITable_indexContext is an interface to support dynamic dispatch.
type ITable_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_indexContext differentiates from other interfaces.
	IsTable_indexContext()
}

type Table_indexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_indexContext() *Table_indexContext {
	var p = new(Table_indexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_table_index
	return p
}

func (*Table_indexContext) IsTable_indexContext() {}

func NewTable_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_indexContext {
	var p = new(Table_indexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_table_index

	return p
}

func (s *Table_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_indexContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *Table_indexContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *Table_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Table_index() (localctx ITable_indexContext) {
	localctx = NewTable_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SpartaParserRULE_table_index)

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
		p.SetState(260)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(261)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(262)
		p.Test()
	}
	{
		p.SetState(263)
		p.Match(SpartaParserT__2)
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
	p.EnterRule(localctx, 68, SpartaParserRULE_test_list)
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
		p.SetState(265)
		p.Test()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(266)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(267)
			p.Test()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
