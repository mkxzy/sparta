// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 214,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 121, 10,
	22, 12, 22, 14, 22, 124, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 130,
	10, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 7, 25, 138, 10, 25, 12,
	25, 14, 25, 141, 11, 25, 3, 25, 5, 25, 144, 10, 25, 3, 26, 3, 26, 7, 26,
	148, 10, 26, 12, 26, 14, 26, 151, 11, 26, 3, 27, 3, 27, 3, 27, 3, 27, 7,
	27, 157, 10, 27, 12, 27, 14, 27, 160, 11, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 171, 10, 28, 12, 28, 14, 28,
	174, 11, 28, 3, 28, 3, 28, 3, 29, 6, 29, 179, 10, 29, 13, 29, 14, 29, 180,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 7, 30, 188, 10, 30, 12, 30, 14, 30,
	191, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 197, 10, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 5, 32, 203, 10, 32, 3, 33, 3, 33, 3, 33, 7, 33, 208,
	10, 33, 12, 33, 14, 33, 211, 11, 33, 5, 33, 213, 10, 33, 3, 158, 2, 34,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 2, 47, 2, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29,
	61, 2, 63, 2, 65, 2, 3, 2, 12, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 10,
	2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118,
	118, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15, 34, 34,
	6, 2, 38, 38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2,
	55298, 56321, 3, 2, 56322, 57345, 3, 2, 51, 59, 2, 222, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2,
	2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 3, 67, 3, 2, 2,
	2, 5, 69, 3, 2, 2, 2, 7, 72, 3, 2, 2, 2, 9, 76, 3, 2, 2, 2, 11, 80, 3,
	2, 2, 2, 13, 82, 3, 2, 2, 2, 15, 84, 3, 2, 2, 2, 17, 87, 3, 2, 2, 2, 19,
	90, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 96, 3, 2, 2, 2, 25, 99, 3, 2, 2,
	2, 27, 101, 3, 2, 2, 2, 29, 103, 3, 2, 2, 2, 31, 105, 3, 2, 2, 2, 33, 107,
	3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37, 112, 3, 2, 2, 2, 39, 114, 3, 2, 2,
	2, 41, 116, 3, 2, 2, 2, 43, 118, 3, 2, 2, 2, 45, 129, 3, 2, 2, 2, 47, 131,
	3, 2, 2, 2, 49, 143, 3, 2, 2, 2, 51, 145, 3, 2, 2, 2, 53, 152, 3, 2, 2,
	2, 55, 166, 3, 2, 2, 2, 57, 178, 3, 2, 2, 2, 59, 184, 3, 2, 2, 2, 61, 196,
	3, 2, 2, 2, 63, 202, 3, 2, 2, 2, 65, 212, 3, 2, 2, 2, 67, 68, 7, 63, 2,
	2, 68, 4, 3, 2, 2, 2, 69, 70, 7, 113, 2, 2, 70, 71, 7, 116, 2, 2, 71, 6,
	3, 2, 2, 2, 72, 73, 7, 99, 2, 2, 73, 74, 7, 112, 2, 2, 74, 75, 7, 102,
	2, 2, 75, 8, 3, 2, 2, 2, 76, 77, 7, 112, 2, 2, 77, 78, 7, 113, 2, 2, 78,
	79, 7, 118, 2, 2, 79, 10, 3, 2, 2, 2, 80, 81, 7, 62, 2, 2, 81, 12, 3, 2,
	2, 2, 82, 83, 7, 64, 2, 2, 83, 14, 3, 2, 2, 2, 84, 85, 7, 63, 2, 2, 85,
	86, 7, 63, 2, 2, 86, 16, 3, 2, 2, 2, 87, 88, 7, 64, 2, 2, 88, 89, 7, 63,
	2, 2, 89, 18, 3, 2, 2, 2, 90, 91, 7, 62, 2, 2, 91, 92, 7, 63, 2, 2, 92,
	20, 3, 2, 2, 2, 93, 94, 7, 62, 2, 2, 94, 95, 7, 64, 2, 2, 95, 22, 3, 2,
	2, 2, 96, 97, 7, 35, 2, 2, 97, 98, 7, 63, 2, 2, 98, 24, 3, 2, 2, 2, 99,
	100, 7, 45, 2, 2, 100, 26, 3, 2, 2, 2, 101, 102, 7, 47, 2, 2, 102, 28,
	3, 2, 2, 2, 103, 104, 7, 44, 2, 2, 104, 30, 3, 2, 2, 2, 105, 106, 7, 49,
	2, 2, 106, 32, 3, 2, 2, 2, 107, 108, 7, 39, 2, 2, 108, 34, 3, 2, 2, 2,
	109, 110, 7, 44, 2, 2, 110, 111, 7, 44, 2, 2, 111, 36, 3, 2, 2, 2, 112,
	113, 7, 42, 2, 2, 113, 38, 3, 2, 2, 2, 114, 115, 7, 43, 2, 2, 115, 40,
	3, 2, 2, 2, 116, 117, 7, 46, 2, 2, 117, 42, 3, 2, 2, 2, 118, 122, 7, 36,
	2, 2, 119, 121, 5, 45, 23, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2,
	2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124,
	122, 3, 2, 2, 2, 125, 126, 7, 36, 2, 2, 126, 44, 3, 2, 2, 2, 127, 130,
	10, 2, 2, 2, 128, 130, 5, 47, 24, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3,
	2, 2, 2, 130, 46, 3, 2, 2, 2, 131, 132, 7, 94, 2, 2, 132, 133, 9, 3, 2,
	2, 133, 48, 3, 2, 2, 2, 134, 135, 5, 65, 33, 2, 135, 139, 7, 48, 2, 2,
	136, 138, 9, 4, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139,
	137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 144, 3, 2, 2, 2, 141, 139,
	3, 2, 2, 2, 142, 144, 5, 65, 33, 2, 143, 134, 3, 2, 2, 2, 143, 142, 3,
	2, 2, 2, 144, 50, 3, 2, 2, 2, 145, 149, 5, 63, 32, 2, 146, 148, 5, 61,
	31, 2, 147, 146, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 52, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 153,
	7, 49, 2, 2, 153, 154, 7, 44, 2, 2, 154, 158, 3, 2, 2, 2, 155, 157, 11,
	2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 159, 3, 2, 2,
	2, 158, 156, 3, 2, 2, 2, 159, 161, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161,
	162, 7, 44, 2, 2, 162, 163, 7, 49, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165,
	8, 27, 2, 2, 165, 54, 3, 2, 2, 2, 166, 167, 7, 49, 2, 2, 167, 168, 7, 49,
	2, 2, 168, 172, 3, 2, 2, 2, 169, 171, 10, 5, 2, 2, 170, 169, 3, 2, 2, 2,
	171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173,
	175, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175, 176, 8, 28, 2, 2, 176, 56,
	3, 2, 2, 2, 177, 179, 9, 6, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3, 2,
	2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2,
	182, 183, 8, 29, 2, 2, 183, 58, 3, 2, 2, 2, 184, 185, 7, 37, 2, 2, 185,
	189, 7, 35, 2, 2, 186, 188, 10, 5, 2, 2, 187, 186, 3, 2, 2, 2, 188, 191,
	3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 192, 3, 2,
	2, 2, 191, 189, 3, 2, 2, 2, 192, 193, 8, 30, 3, 2, 193, 60, 3, 2, 2, 2,
	194, 197, 5, 63, 32, 2, 195, 197, 9, 4, 2, 2, 196, 194, 3, 2, 2, 2, 196,
	195, 3, 2, 2, 2, 197, 62, 3, 2, 2, 2, 198, 203, 9, 7, 2, 2, 199, 203, 10,
	8, 2, 2, 200, 201, 9, 9, 2, 2, 201, 203, 9, 10, 2, 2, 202, 198, 3, 2, 2,
	2, 202, 199, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 64, 3, 2, 2, 2, 204,
	213, 7, 50, 2, 2, 205, 209, 9, 11, 2, 2, 206, 208, 9, 4, 2, 2, 207, 206,
	3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2,
	2, 2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 204, 3, 2, 2, 2,
	212, 205, 3, 2, 2, 2, 213, 66, 3, 2, 2, 2, 16, 2, 122, 129, 139, 143, 149,
	158, 172, 180, 189, 196, 202, 209, 212, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'or'", "'and'", "'not'", "'<'", "'>'", "'=='", "'>='", "'<='",
	"'<>'", "'!='", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'('", "')'",
	"','",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "STRING", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT",
	"WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "STRING", "StringCharacter", "EscapeSequence",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
	"LetterOrDigit", "Letter", "DecimalIntegerLiteral",
}

type SpartaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSpartaLexer(input antlr.CharStream) *SpartaLexer {

	l := new(SpartaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sparta.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SpartaLexer tokens.
const (
	SpartaLexerT__0           = 1
	SpartaLexerT__1           = 2
	SpartaLexerT__2           = 3
	SpartaLexerT__3           = 4
	SpartaLexerT__4           = 5
	SpartaLexerT__5           = 6
	SpartaLexerT__6           = 7
	SpartaLexerT__7           = 8
	SpartaLexerT__8           = 9
	SpartaLexerT__9           = 10
	SpartaLexerT__10          = 11
	SpartaLexerT__11          = 12
	SpartaLexerT__12          = 13
	SpartaLexerT__13          = 14
	SpartaLexerT__14          = 15
	SpartaLexerT__15          = 16
	SpartaLexerT__16          = 17
	SpartaLexerT__17          = 18
	SpartaLexerT__18          = 19
	SpartaLexerT__19          = 20
	SpartaLexerSTRING         = 21
	SpartaLexerNUMBER_LITERAL = 22
	SpartaLexerIDENTIFIER     = 23
	SpartaLexerCOMMENT        = 24
	SpartaLexerLINE_COMMENT   = 25
	SpartaLexerWS             = 26
	SpartaLexerSHEBANG        = 27
)
