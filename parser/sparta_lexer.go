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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 274,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 7, 32, 183, 10, 32, 12,
	32, 14, 32, 186, 11, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34,
	6, 34, 195, 10, 34, 13, 34, 14, 34, 196, 3, 35, 3, 35, 7, 35, 201, 10,
	35, 12, 35, 14, 35, 204, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 210,
	10, 36, 12, 36, 14, 36, 213, 11, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 224, 10, 37, 12, 37, 14, 37, 227, 11,
	37, 3, 37, 3, 37, 3, 38, 6, 38, 232, 10, 38, 13, 38, 14, 38, 233, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 7, 39, 241, 10, 39, 12, 39, 14, 39, 244, 11,
	39, 3, 39, 3, 39, 3, 40, 3, 40, 5, 40, 250, 10, 40, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 5, 42, 257, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 263,
	10, 43, 3, 44, 3, 44, 3, 44, 7, 44, 268, 10, 44, 12, 44, 14, 44, 271, 11,
	44, 5, 44, 273, 10, 44, 3, 211, 2, 45, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 2, 81, 2, 83, 2, 85, 2, 87,
	2, 3, 2, 12, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15,
	34, 34, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 10, 2, 36, 36, 41, 41, 94,
	94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 6, 2, 38, 38, 67,
	92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321, 3,
	2, 56322, 57345, 3, 2, 51, 59, 2, 281, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 89, 3, 2, 2, 2, 5, 91, 3, 2, 2, 2, 7,
	93, 3, 2, 2, 2, 9, 95, 3, 2, 2, 2, 11, 99, 3, 2, 2, 2, 13, 101, 3, 2, 2,
	2, 15, 103, 3, 2, 2, 2, 17, 105, 3, 2, 2, 2, 19, 112, 3, 2, 2, 2, 21, 115,
	3, 2, 2, 2, 23, 120, 3, 2, 2, 2, 25, 124, 3, 2, 2, 2, 27, 127, 3, 2, 2,
	2, 29, 130, 3, 2, 2, 2, 31, 132, 3, 2, 2, 2, 33, 134, 3, 2, 2, 2, 35, 140,
	3, 2, 2, 2, 37, 149, 3, 2, 2, 2, 39, 151, 3, 2, 2, 2, 41, 153, 3, 2, 2,
	2, 43, 156, 3, 2, 2, 2, 45, 159, 3, 2, 2, 2, 47, 162, 3, 2, 2, 2, 49, 165,
	3, 2, 2, 2, 51, 168, 3, 2, 2, 2, 53, 170, 3, 2, 2, 2, 55, 172, 3, 2, 2,
	2, 57, 174, 3, 2, 2, 2, 59, 176, 3, 2, 2, 2, 61, 178, 3, 2, 2, 2, 63, 180,
	3, 2, 2, 2, 65, 189, 3, 2, 2, 2, 67, 191, 3, 2, 2, 2, 69, 198, 3, 2, 2,
	2, 71, 205, 3, 2, 2, 2, 73, 219, 3, 2, 2, 2, 75, 231, 3, 2, 2, 2, 77, 237,
	3, 2, 2, 2, 79, 249, 3, 2, 2, 2, 81, 251, 3, 2, 2, 2, 83, 256, 3, 2, 2,
	2, 85, 262, 3, 2, 2, 2, 87, 272, 3, 2, 2, 2, 89, 90, 7, 63, 2, 2, 90, 4,
	3, 2, 2, 2, 91, 92, 7, 93, 2, 2, 92, 6, 3, 2, 2, 2, 93, 94, 7, 95, 2, 2,
	94, 8, 3, 2, 2, 2, 95, 96, 7, 104, 2, 2, 96, 97, 7, 119, 2, 2, 97, 98,
	7, 112, 2, 2, 98, 10, 3, 2, 2, 2, 99, 100, 7, 42, 2, 2, 100, 12, 3, 2,
	2, 2, 101, 102, 7, 43, 2, 2, 102, 14, 3, 2, 2, 2, 103, 104, 7, 46, 2, 2,
	104, 16, 3, 2, 2, 2, 105, 106, 7, 116, 2, 2, 106, 107, 7, 103, 2, 2, 107,
	108, 7, 118, 2, 2, 108, 109, 7, 119, 2, 2, 109, 110, 7, 116, 2, 2, 110,
	111, 7, 112, 2, 2, 111, 18, 3, 2, 2, 2, 112, 113, 7, 107, 2, 2, 113, 114,
	7, 104, 2, 2, 114, 20, 3, 2, 2, 2, 115, 116, 7, 103, 2, 2, 116, 117, 7,
	110, 2, 2, 117, 118, 7, 117, 2, 2, 118, 119, 7, 103, 2, 2, 119, 22, 3,
	2, 2, 2, 120, 121, 7, 104, 2, 2, 121, 122, 7, 113, 2, 2, 122, 123, 7, 116,
	2, 2, 123, 24, 3, 2, 2, 2, 124, 125, 7, 107, 2, 2, 125, 126, 7, 112, 2,
	2, 126, 26, 3, 2, 2, 2, 127, 128, 7, 118, 2, 2, 128, 129, 7, 113, 2, 2,
	129, 28, 3, 2, 2, 2, 130, 131, 7, 125, 2, 2, 131, 30, 3, 2, 2, 2, 132,
	133, 7, 127, 2, 2, 133, 32, 3, 2, 2, 2, 134, 135, 7, 100, 2, 2, 135, 136,
	7, 116, 2, 2, 136, 137, 7, 103, 2, 2, 137, 138, 7, 99, 2, 2, 138, 139,
	7, 109, 2, 2, 139, 34, 3, 2, 2, 2, 140, 141, 7, 101, 2, 2, 141, 142, 7,
	113, 2, 2, 142, 143, 7, 112, 2, 2, 143, 144, 7, 118, 2, 2, 144, 145, 7,
	107, 2, 2, 145, 146, 7, 112, 2, 2, 146, 147, 7, 119, 2, 2, 147, 148, 7,
	103, 2, 2, 148, 36, 3, 2, 2, 2, 149, 150, 7, 62, 2, 2, 150, 38, 3, 2, 2,
	2, 151, 152, 7, 64, 2, 2, 152, 40, 3, 2, 2, 2, 153, 154, 7, 63, 2, 2, 154,
	155, 7, 63, 2, 2, 155, 42, 3, 2, 2, 2, 156, 157, 7, 64, 2, 2, 157, 158,
	7, 63, 2, 2, 158, 44, 3, 2, 2, 2, 159, 160, 7, 62, 2, 2, 160, 161, 7, 63,
	2, 2, 161, 46, 3, 2, 2, 2, 162, 163, 7, 62, 2, 2, 163, 164, 7, 64, 2, 2,
	164, 48, 3, 2, 2, 2, 165, 166, 7, 35, 2, 2, 166, 167, 7, 63, 2, 2, 167,
	50, 3, 2, 2, 2, 168, 169, 7, 45, 2, 2, 169, 52, 3, 2, 2, 2, 170, 171, 7,
	47, 2, 2, 171, 54, 3, 2, 2, 2, 172, 173, 7, 44, 2, 2, 173, 56, 3, 2, 2,
	2, 174, 175, 7, 49, 2, 2, 175, 58, 3, 2, 2, 2, 176, 177, 7, 39, 2, 2, 177,
	60, 3, 2, 2, 2, 178, 179, 7, 60, 2, 2, 179, 62, 3, 2, 2, 2, 180, 184, 7,
	36, 2, 2, 181, 183, 5, 79, 40, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2,
	2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2,
	186, 184, 3, 2, 2, 2, 187, 188, 7, 36, 2, 2, 188, 64, 3, 2, 2, 2, 189,
	190, 5, 87, 44, 2, 190, 66, 3, 2, 2, 2, 191, 192, 5, 87, 44, 2, 192, 194,
	7, 48, 2, 2, 193, 195, 9, 2, 2, 2, 194, 193, 3, 2, 2, 2, 195, 196, 3, 2,
	2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 68, 3, 2, 2, 2,
	198, 202, 5, 85, 43, 2, 199, 201, 5, 83, 42, 2, 200, 199, 3, 2, 2, 2, 201,
	204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 70, 3,
	2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206, 7, 49, 2, 2, 206, 207, 7, 44,
	2, 2, 207, 211, 3, 2, 2, 2, 208, 210, 11, 2, 2, 2, 209, 208, 3, 2, 2, 2,
	210, 213, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212,
	214, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 215, 7, 44, 2, 2, 215, 216,
	7, 49, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 8, 36, 2, 2, 218, 72, 3, 2,
	2, 2, 219, 220, 7, 49, 2, 2, 220, 221, 7, 49, 2, 2, 221, 225, 3, 2, 2,
	2, 222, 224, 10, 3, 2, 2, 223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 228, 229, 8, 37, 2, 2, 229, 74, 3, 2, 2, 2, 230, 232, 9, 4,
	2, 2, 231, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2,
	233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 8, 38, 2, 2, 236,
	76, 3, 2, 2, 2, 237, 238, 7, 37, 2, 2, 238, 242, 7, 35, 2, 2, 239, 241,
	10, 3, 2, 2, 240, 239, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240, 3, 2,
	2, 2, 242, 243, 3, 2, 2, 2, 243, 245, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2,
	245, 246, 8, 39, 3, 2, 246, 78, 3, 2, 2, 2, 247, 250, 10, 5, 2, 2, 248,
	250, 5, 81, 41, 2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 80,
	3, 2, 2, 2, 251, 252, 7, 94, 2, 2, 252, 253, 9, 6, 2, 2, 253, 82, 3, 2,
	2, 2, 254, 257, 5, 85, 43, 2, 255, 257, 9, 2, 2, 2, 256, 254, 3, 2, 2,
	2, 256, 255, 3, 2, 2, 2, 257, 84, 3, 2, 2, 2, 258, 263, 9, 7, 2, 2, 259,
	263, 10, 8, 2, 2, 260, 261, 9, 9, 2, 2, 261, 263, 9, 10, 2, 2, 262, 258,
	3, 2, 2, 2, 262, 259, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 86, 3, 2,
	2, 2, 264, 273, 7, 50, 2, 2, 265, 269, 9, 11, 2, 2, 266, 268, 9, 2, 2,
	2, 267, 266, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 264,
	3, 2, 2, 2, 272, 265, 3, 2, 2, 2, 273, 88, 3, 2, 2, 2, 15, 2, 184, 196,
	202, 211, 225, 233, 242, 249, 256, 262, 269, 272, 4, 8, 2, 2, 2, 3, 2,
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
	"", "'='", "'['", "']'", "'fun'", "'('", "')'", "','", "'return'", "'if'",
	"'else'", "'for'", "'in'", "'to'", "'{'", "'}'", "'break'", "'continue'",
	"'<'", "'>'", "'=='", "'>='", "'<='", "'<>'", "'!='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INTEGER_LITERAL",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "STRING", "INTEGER_LITERAL",
	"NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
	"StringCharacter", "EscapeSequence", "LetterOrDigit", "Letter", "DecimalIntegerLiteral",
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
	SpartaLexerT__0            = 1
	SpartaLexerT__1            = 2
	SpartaLexerT__2            = 3
	SpartaLexerT__3            = 4
	SpartaLexerT__4            = 5
	SpartaLexerT__5            = 6
	SpartaLexerT__6            = 7
	SpartaLexerT__7            = 8
	SpartaLexerT__8            = 9
	SpartaLexerT__9            = 10
	SpartaLexerT__10           = 11
	SpartaLexerT__11           = 12
	SpartaLexerT__12           = 13
	SpartaLexerT__13           = 14
	SpartaLexerT__14           = 15
	SpartaLexerT__15           = 16
	SpartaLexerT__16           = 17
	SpartaLexerT__17           = 18
	SpartaLexerT__18           = 19
	SpartaLexerT__19           = 20
	SpartaLexerT__20           = 21
	SpartaLexerT__21           = 22
	SpartaLexerT__22           = 23
	SpartaLexerT__23           = 24
	SpartaLexerT__24           = 25
	SpartaLexerT__25           = 26
	SpartaLexerT__26           = 27
	SpartaLexerT__27           = 28
	SpartaLexerT__28           = 29
	SpartaLexerT__29           = 30
	SpartaLexerSTRING          = 31
	SpartaLexerINTEGER_LITERAL = 32
	SpartaLexerNUMBER_LITERAL  = 33
	SpartaLexerIDENTIFIER      = 34
	SpartaLexerCOMMENT         = 35
	SpartaLexerLINE_COMMENT    = 36
	SpartaLexerWS              = 37
	SpartaLexerSHEBANG         = 38
)
