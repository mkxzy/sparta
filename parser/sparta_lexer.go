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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 111,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 35, 10, 4,
	12, 4, 14, 4, 38, 11, 4, 3, 4, 5, 4, 41, 10, 4, 3, 5, 3, 5, 7, 5, 45, 10,
	5, 12, 5, 14, 5, 48, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 54, 10, 6, 12,
	6, 14, 6, 57, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 7, 7, 68, 10, 7, 12, 7, 14, 7, 71, 11, 7, 3, 7, 3, 7, 3, 8, 6, 8, 76,
	10, 8, 13, 8, 14, 8, 77, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 85, 10, 9,
	12, 9, 14, 9, 88, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 94, 10, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 100, 10, 11, 3, 12, 3, 12, 3, 12, 7, 12,
	105, 10, 12, 12, 12, 14, 12, 108, 11, 12, 5, 12, 110, 10, 12, 3, 55, 2,
	13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 2, 21, 2,
	23, 2, 3, 2, 10, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14,
	15, 34, 34, 6, 2, 38, 38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298,
	56321, 3, 2, 55298, 56321, 3, 2, 56322, 57345, 3, 2, 51, 59, 2, 119, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	3, 25, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 40, 3, 2, 2, 2, 9, 42, 3, 2, 2,
	2, 11, 49, 3, 2, 2, 2, 13, 63, 3, 2, 2, 2, 15, 75, 3, 2, 2, 2, 17, 81,
	3, 2, 2, 2, 19, 93, 3, 2, 2, 2, 21, 99, 3, 2, 2, 2, 23, 109, 3, 2, 2, 2,
	25, 26, 7, 120, 2, 2, 26, 27, 7, 99, 2, 2, 27, 28, 7, 116, 2, 2, 28, 4,
	3, 2, 2, 2, 29, 30, 7, 63, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 5, 23, 12,
	2, 32, 36, 7, 48, 2, 2, 33, 35, 9, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 38,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 41, 3, 2, 2, 2,
	38, 36, 3, 2, 2, 2, 39, 41, 5, 23, 12, 2, 40, 31, 3, 2, 2, 2, 40, 39, 3,
	2, 2, 2, 41, 8, 3, 2, 2, 2, 42, 46, 5, 21, 11, 2, 43, 45, 5, 19, 10, 2,
	44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 10, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 49, 2, 2, 50,
	51, 7, 44, 2, 2, 51, 55, 3, 2, 2, 2, 52, 54, 11, 2, 2, 2, 53, 52, 3, 2,
	2, 2, 54, 57, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56, 58,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 7, 44, 2, 2, 59, 60, 7, 49, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 62, 8, 6, 2, 2, 62, 12, 3, 2, 2, 2, 63, 64,
	7, 49, 2, 2, 64, 65, 7, 49, 2, 2, 65, 69, 3, 2, 2, 2, 66, 68, 10, 3, 2,
	2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 8, 7, 2, 2,
	73, 14, 3, 2, 2, 2, 74, 76, 9, 4, 2, 2, 75, 74, 3, 2, 2, 2, 76, 77, 3,
	2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79,
	80, 8, 8, 2, 2, 80, 16, 3, 2, 2, 2, 81, 82, 7, 37, 2, 2, 82, 86, 7, 35,
	2, 2, 83, 85, 10, 3, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 89, 90, 8, 9, 3, 2, 90, 18, 3, 2, 2, 2, 91, 94, 5, 21, 11, 2, 92, 94,
	9, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 20, 3, 2, 2, 2,
	95, 100, 9, 5, 2, 2, 96, 100, 10, 6, 2, 2, 97, 98, 9, 7, 2, 2, 98, 100,
	9, 8, 2, 2, 99, 95, 3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2,
	100, 22, 3, 2, 2, 2, 101, 110, 7, 50, 2, 2, 102, 106, 9, 9, 2, 2, 103,
	105, 9, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2,
	2, 2, 109, 101, 3, 2, 2, 2, 109, 102, 3, 2, 2, 2, 110, 24, 3, 2, 2, 2,
	14, 2, 36, 40, 46, 55, 69, 77, 86, 93, 99, 106, 109, 4, 8, 2, 2, 2, 3,
	2,
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
	"", "'var'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT",
	"WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT",
	"WS", "SHEBANG", "LetterOrDigit", "Letter", "DecimalIntegerLiteral",
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
	SpartaLexerNUMBER_LITERAL = 3
	SpartaLexerIDENTIFIER     = 4
	SpartaLexerCOMMENT        = 5
	SpartaLexerLINE_COMMENT   = 6
	SpartaLexerWS             = 7
	SpartaLexerSHEBANG        = 8
)
