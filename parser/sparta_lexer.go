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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 181,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 89, 10, 15,
	12, 15, 14, 15, 92, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 98, 10,
	16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 6, 18, 106, 10, 18, 13, 18,
	14, 18, 107, 3, 18, 5, 18, 111, 10, 18, 3, 19, 3, 19, 7, 19, 115, 10, 19,
	12, 19, 14, 19, 118, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 124, 10,
	20, 12, 20, 14, 20, 127, 11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 7, 21, 138, 10, 21, 12, 21, 14, 21, 141, 11, 21,
	3, 21, 3, 21, 3, 22, 6, 22, 146, 10, 22, 13, 22, 14, 22, 147, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 7, 23, 155, 10, 23, 12, 23, 14, 23, 158, 11, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 5, 24, 164, 10, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 170, 10, 25, 3, 26, 3, 26, 3, 26, 7, 26, 175, 10, 26, 12, 26,
	14, 26, 178, 11, 26, 5, 26, 180, 10, 26, 3, 125, 2, 27, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 2, 33, 2, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22,
	47, 2, 49, 2, 51, 2, 3, 2, 12, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 10,
	2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118,
	118, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15, 34, 34,
	6, 2, 38, 38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2,
	55298, 56321, 3, 2, 56322, 57345, 3, 2, 51, 59, 2, 189, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39,
	3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3,
	53, 3, 2, 2, 2, 5, 55, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 61, 3, 2, 2, 2,
	11, 63, 3, 2, 2, 2, 13, 65, 3, 2, 2, 2, 15, 67, 3, 2, 2, 2, 17, 69, 3,
	2, 2, 2, 19, 76, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23, 80, 3, 2, 2, 2, 25,
	82, 3, 2, 2, 2, 27, 84, 3, 2, 2, 2, 29, 86, 3, 2, 2, 2, 31, 97, 3, 2, 2,
	2, 33, 99, 3, 2, 2, 2, 35, 110, 3, 2, 2, 2, 37, 112, 3, 2, 2, 2, 39, 119,
	3, 2, 2, 2, 41, 133, 3, 2, 2, 2, 43, 145, 3, 2, 2, 2, 45, 151, 3, 2, 2,
	2, 47, 163, 3, 2, 2, 2, 49, 169, 3, 2, 2, 2, 51, 179, 3, 2, 2, 2, 53, 54,
	7, 63, 2, 2, 54, 4, 3, 2, 2, 2, 55, 56, 7, 104, 2, 2, 56, 57, 7, 119, 2,
	2, 57, 58, 7, 112, 2, 2, 58, 6, 3, 2, 2, 2, 59, 60, 7, 125, 2, 2, 60, 8,
	3, 2, 2, 2, 61, 62, 7, 127, 2, 2, 62, 10, 3, 2, 2, 2, 63, 64, 7, 42, 2,
	2, 64, 12, 3, 2, 2, 2, 65, 66, 7, 43, 2, 2, 66, 14, 3, 2, 2, 2, 67, 68,
	7, 46, 2, 2, 68, 16, 3, 2, 2, 2, 69, 70, 7, 116, 2, 2, 70, 71, 7, 103,
	2, 2, 71, 72, 7, 118, 2, 2, 72, 73, 7, 119, 2, 2, 73, 74, 7, 116, 2, 2,
	74, 75, 7, 112, 2, 2, 75, 18, 3, 2, 2, 2, 76, 77, 7, 45, 2, 2, 77, 20,
	3, 2, 2, 2, 78, 79, 7, 47, 2, 2, 79, 22, 3, 2, 2, 2, 80, 81, 7, 44, 2,
	2, 81, 24, 3, 2, 2, 2, 82, 83, 7, 49, 2, 2, 83, 26, 3, 2, 2, 2, 84, 85,
	7, 39, 2, 2, 85, 28, 3, 2, 2, 2, 86, 90, 7, 36, 2, 2, 87, 89, 5, 31, 16,
	2, 88, 87, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 36, 2, 2,
	94, 30, 3, 2, 2, 2, 95, 98, 10, 2, 2, 2, 96, 98, 5, 33, 17, 2, 97, 95,
	3, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 32, 3, 2, 2, 2, 99, 100, 7, 94, 2,
	2, 100, 101, 9, 3, 2, 2, 101, 34, 3, 2, 2, 2, 102, 103, 5, 51, 26, 2, 103,
	105, 7, 48, 2, 2, 104, 106, 9, 4, 2, 2, 105, 104, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2,
	2, 2, 109, 111, 5, 51, 26, 2, 110, 102, 3, 2, 2, 2, 110, 109, 3, 2, 2,
	2, 111, 36, 3, 2, 2, 2, 112, 116, 5, 49, 25, 2, 113, 115, 5, 47, 24, 2,
	114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116,
	117, 3, 2, 2, 2, 117, 38, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 7,
	49, 2, 2, 120, 121, 7, 44, 2, 2, 121, 125, 3, 2, 2, 2, 122, 124, 11, 2,
	2, 2, 123, 122, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 126, 128, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 128,
	129, 7, 44, 2, 2, 129, 130, 7, 49, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132,
	8, 20, 2, 2, 132, 40, 3, 2, 2, 2, 133, 134, 7, 49, 2, 2, 134, 135, 7, 49,
	2, 2, 135, 139, 3, 2, 2, 2, 136, 138, 10, 5, 2, 2, 137, 136, 3, 2, 2, 2,
	138, 141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140,
	142, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 143, 8, 21, 2, 2, 143, 42,
	3, 2, 2, 2, 144, 146, 9, 6, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2,
	2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2,
	149, 150, 8, 22, 2, 2, 150, 44, 3, 2, 2, 2, 151, 152, 7, 37, 2, 2, 152,
	156, 7, 35, 2, 2, 153, 155, 10, 5, 2, 2, 154, 153, 3, 2, 2, 2, 155, 158,
	3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2,
	2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 8, 23, 3, 2, 160, 46, 3, 2, 2, 2,
	161, 164, 5, 49, 25, 2, 162, 164, 9, 4, 2, 2, 163, 161, 3, 2, 2, 2, 163,
	162, 3, 2, 2, 2, 164, 48, 3, 2, 2, 2, 165, 170, 9, 7, 2, 2, 166, 170, 10,
	8, 2, 2, 167, 168, 9, 9, 2, 2, 168, 170, 9, 10, 2, 2, 169, 165, 3, 2, 2,
	2, 169, 166, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 50, 3, 2, 2, 2, 171,
	180, 7, 50, 2, 2, 172, 176, 9, 11, 2, 2, 173, 175, 9, 4, 2, 2, 174, 173,
	3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177, 3, 2,
	2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 179, 171, 3, 2, 2, 2,
	179, 172, 3, 2, 2, 2, 180, 52, 3, 2, 2, 2, 16, 2, 90, 97, 107, 110, 116,
	125, 139, 147, 156, 163, 169, 176, 179, 4, 8, 2, 2, 2, 3, 2,
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
	"", "'='", "'fun'", "'{'", "'}'", "'('", "')'", "','", "'return'", "'+'",
	"'-'", "'*'", "'/'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER_LITERAL",
	"IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "STRING", "StringCharacter", "EscapeSequence",
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
	SpartaLexerSTRING         = 14
	SpartaLexerNUMBER_LITERAL = 15
	SpartaLexerIDENTIFIER     = 16
	SpartaLexerCOMMENT        = 17
	SpartaLexerLINE_COMMENT   = 18
	SpartaLexerWS             = 19
	SpartaLexerSHEBANG        = 20
)
