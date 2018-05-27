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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 305,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 3, 2, 7, 2, 80, 10, 2, 12, 2, 14, 2, 83, 11, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 96, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 108, 10, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 121,
	10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 128, 10, 10, 12, 10, 14,
	10, 131, 11, 10, 3, 11, 3, 11, 5, 11, 135, 10, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 144, 10, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 7, 15, 151, 10, 15, 12, 15, 14, 15, 154, 11, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 166, 10,
	17, 12, 17, 14, 17, 169, 11, 17, 3, 17, 3, 17, 5, 17, 173, 10, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 7, 19, 185,
	10, 19, 12, 19, 14, 19, 188, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 7, 24, 204,
	10, 24, 12, 24, 14, 24, 207, 11, 24, 3, 25, 3, 25, 3, 25, 7, 25, 212, 10,
	25, 12, 25, 14, 25, 215, 11, 25, 3, 26, 3, 26, 3, 26, 5, 26, 220, 10, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 226, 10, 27, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 7, 29, 233, 10, 29, 12, 29, 14, 29, 236, 11, 29, 3, 30, 3, 30,
	3, 30, 7, 30, 241, 10, 30, 12, 30, 14, 30, 244, 11, 30, 3, 31, 5, 31, 247,
	10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 262, 10, 32, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 34, 3, 34, 5, 34, 270, 10, 34, 3, 34, 3, 34, 3, 35, 3, 35, 5, 35,
	276, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 283, 10, 36, 12,
	36, 14, 36, 286, 11, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 7, 39, 300, 10, 39, 12, 39, 14, 39,
	303, 11, 39, 3, 39, 2, 2, 40, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 2, 5, 3, 2, 24, 29, 3, 2, 30, 31, 3, 2,
	32, 34, 2, 305, 2, 81, 3, 2, 2, 2, 4, 95, 3, 2, 2, 2, 6, 97, 3, 2, 2, 2,
	8, 107, 3, 2, 2, 2, 10, 109, 3, 2, 2, 2, 12, 113, 3, 2, 2, 2, 14, 115,
	3, 2, 2, 2, 16, 118, 3, 2, 2, 2, 18, 124, 3, 2, 2, 2, 20, 132, 3, 2, 2,
	2, 22, 136, 3, 2, 2, 2, 24, 138, 3, 2, 2, 2, 26, 141, 3, 2, 2, 2, 28, 147,
	3, 2, 2, 2, 30, 155, 3, 2, 2, 2, 32, 157, 3, 2, 2, 2, 34, 174, 3, 2, 2,
	2, 36, 182, 3, 2, 2, 2, 38, 191, 3, 2, 2, 2, 40, 193, 3, 2, 2, 2, 42, 195,
	3, 2, 2, 2, 44, 198, 3, 2, 2, 2, 46, 200, 3, 2, 2, 2, 48, 208, 3, 2, 2,
	2, 50, 219, 3, 2, 2, 2, 52, 221, 3, 2, 2, 2, 54, 227, 3, 2, 2, 2, 56, 229,
	3, 2, 2, 2, 58, 237, 3, 2, 2, 2, 60, 246, 3, 2, 2, 2, 62, 261, 3, 2, 2,
	2, 64, 263, 3, 2, 2, 2, 66, 267, 3, 2, 2, 2, 68, 273, 3, 2, 2, 2, 70, 279,
	3, 2, 2, 2, 72, 287, 3, 2, 2, 2, 74, 291, 3, 2, 2, 2, 76, 296, 3, 2, 2,
	2, 78, 80, 5, 4, 3, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	84, 85, 7, 2, 2, 3, 85, 3, 3, 2, 2, 2, 86, 96, 5, 6, 4, 2, 87, 96, 5, 10,
	6, 2, 88, 96, 5, 20, 11, 2, 89, 96, 5, 22, 12, 2, 90, 96, 5, 32, 17, 2,
	91, 96, 5, 34, 18, 2, 92, 96, 5, 38, 20, 2, 93, 96, 5, 40, 21, 2, 94, 96,
	5, 42, 22, 2, 95, 86, 3, 2, 2, 2, 95, 87, 3, 2, 2, 2, 95, 88, 3, 2, 2,
	2, 95, 89, 3, 2, 2, 2, 95, 90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92,
	3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 5, 3, 2, 2, 2,
	97, 98, 5, 8, 5, 2, 98, 99, 7, 3, 2, 2, 99, 100, 5, 44, 23, 2, 100, 7,
	3, 2, 2, 2, 101, 108, 7, 41, 2, 2, 102, 103, 7, 41, 2, 2, 103, 104, 7,
	4, 2, 2, 104, 105, 5, 44, 23, 2, 105, 106, 7, 5, 2, 2, 106, 108, 3, 2,
	2, 2, 107, 101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 108, 9, 3, 2, 2, 2, 109,
	110, 7, 6, 2, 2, 110, 111, 5, 12, 7, 2, 111, 112, 5, 14, 8, 2, 112, 11,
	3, 2, 2, 2, 113, 114, 7, 41, 2, 2, 114, 13, 3, 2, 2, 2, 115, 116, 5, 16,
	9, 2, 116, 117, 5, 36, 19, 2, 117, 15, 3, 2, 2, 2, 118, 120, 7, 7, 2, 2,
	119, 121, 5, 18, 10, 2, 120, 119, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	122, 3, 2, 2, 2, 122, 123, 7, 8, 2, 2, 123, 17, 3, 2, 2, 2, 124, 129, 7,
	41, 2, 2, 125, 126, 7, 9, 2, 2, 126, 128, 7, 41, 2, 2, 127, 125, 3, 2,
	2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2,
	130, 19, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 134, 7, 10, 2, 2, 133,
	135, 5, 44, 23, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 21,
	3, 2, 2, 2, 136, 137, 5, 24, 13, 2, 137, 23, 3, 2, 2, 2, 138, 139, 5, 12,
	7, 2, 139, 140, 5, 26, 14, 2, 140, 25, 3, 2, 2, 2, 141, 143, 7, 7, 2, 2,
	142, 144, 5, 28, 15, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	145, 3, 2, 2, 2, 145, 146, 7, 8, 2, 2, 146, 27, 3, 2, 2, 2, 147, 152, 5,
	30, 16, 2, 148, 149, 7, 9, 2, 2, 149, 151, 5, 30, 16, 2, 150, 148, 3, 2,
	2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2,
	153, 29, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 5, 44, 23, 2, 156,
	31, 3, 2, 2, 2, 157, 158, 7, 11, 2, 2, 158, 159, 5, 44, 23, 2, 159, 167,
	5, 36, 19, 2, 160, 161, 7, 12, 2, 2, 161, 162, 7, 11, 2, 2, 162, 163, 5,
	44, 23, 2, 163, 164, 5, 36, 19, 2, 164, 166, 3, 2, 2, 2, 165, 160, 3, 2,
	2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2,
	168, 172, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 171, 7, 12, 2, 2, 171,
	173, 5, 36, 19, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 33,
	3, 2, 2, 2, 174, 175, 7, 13, 2, 2, 175, 176, 7, 41, 2, 2, 176, 177, 7,
	14, 2, 2, 177, 178, 5, 44, 23, 2, 178, 179, 7, 15, 2, 2, 179, 180, 5, 44,
	23, 2, 180, 181, 5, 36, 19, 2, 181, 35, 3, 2, 2, 2, 182, 186, 7, 16, 2,
	2, 183, 185, 5, 4, 3, 2, 184, 183, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186,
	184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 189, 3, 2, 2, 2, 188, 186,
	3, 2, 2, 2, 189, 190, 7, 17, 2, 2, 190, 37, 3, 2, 2, 2, 191, 192, 7, 18,
	2, 2, 192, 39, 3, 2, 2, 2, 193, 194, 7, 19, 2, 2, 194, 41, 3, 2, 2, 2,
	195, 196, 7, 20, 2, 2, 196, 197, 7, 41, 2, 2, 197, 43, 3, 2, 2, 2, 198,
	199, 5, 46, 24, 2, 199, 45, 3, 2, 2, 2, 200, 205, 5, 48, 25, 2, 201, 202,
	7, 21, 2, 2, 202, 204, 5, 48, 25, 2, 203, 201, 3, 2, 2, 2, 204, 207, 3,
	2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 47, 3, 2, 2,
	2, 207, 205, 3, 2, 2, 2, 208, 213, 5, 50, 26, 2, 209, 210, 7, 22, 2, 2,
	210, 212, 5, 50, 26, 2, 211, 209, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213,
	211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 49, 3, 2, 2, 2, 215, 213, 3,
	2, 2, 2, 216, 217, 7, 23, 2, 2, 217, 220, 5, 50, 26, 2, 218, 220, 5, 52,
	27, 2, 219, 216, 3, 2, 2, 2, 219, 218, 3, 2, 2, 2, 220, 51, 3, 2, 2, 2,
	221, 225, 5, 56, 29, 2, 222, 223, 5, 54, 28, 2, 223, 224, 5, 56, 29, 2,
	224, 226, 3, 2, 2, 2, 225, 222, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226,
	53, 3, 2, 2, 2, 227, 228, 9, 2, 2, 2, 228, 55, 3, 2, 2, 2, 229, 234, 5,
	58, 30, 2, 230, 231, 9, 3, 2, 2, 231, 233, 5, 58, 30, 2, 232, 230, 3, 2,
	2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2,
	235, 57, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 242, 5, 60, 31, 2, 238,
	239, 9, 4, 2, 2, 239, 241, 5, 60, 31, 2, 240, 238, 3, 2, 2, 2, 241, 244,
	3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 59, 3, 2,
	2, 2, 244, 242, 3, 2, 2, 2, 245, 247, 7, 31, 2, 2, 246, 245, 3, 2, 2, 2,
	246, 247, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 5, 62, 32, 2, 249,
	61, 3, 2, 2, 2, 250, 262, 5, 64, 33, 2, 251, 262, 5, 66, 34, 2, 252, 262,
	5, 68, 35, 2, 253, 262, 5, 24, 13, 2, 254, 262, 5, 74, 38, 2, 255, 262,
	7, 41, 2, 2, 256, 262, 7, 39, 2, 2, 257, 262, 7, 40, 2, 2, 258, 262, 7,
	36, 2, 2, 259, 262, 7, 37, 2, 2, 260, 262, 7, 38, 2, 2, 261, 250, 3, 2,
	2, 2, 261, 251, 3, 2, 2, 2, 261, 252, 3, 2, 2, 2, 261, 253, 3, 2, 2, 2,
	261, 254, 3, 2, 2, 2, 261, 255, 3, 2, 2, 2, 261, 256, 3, 2, 2, 2, 261,
	257, 3, 2, 2, 2, 261, 258, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 260,
	3, 2, 2, 2, 262, 63, 3, 2, 2, 2, 263, 264, 7, 7, 2, 2, 264, 265, 5, 44,
	23, 2, 265, 266, 7, 8, 2, 2, 266, 65, 3, 2, 2, 2, 267, 269, 7, 4, 2, 2,
	268, 270, 5, 76, 39, 2, 269, 268, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270,
	271, 3, 2, 2, 2, 271, 272, 7, 5, 2, 2, 272, 67, 3, 2, 2, 2, 273, 275, 7,
	16, 2, 2, 274, 276, 5, 70, 36, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2,
	2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 7, 17, 2, 2, 278, 69, 3, 2, 2, 2,
	279, 284, 5, 72, 37, 2, 280, 281, 7, 9, 2, 2, 281, 283, 5, 72, 37, 2, 282,
	280, 3, 2, 2, 2, 283, 286, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285,
	3, 2, 2, 2, 285, 71, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 288, 5, 44,
	23, 2, 288, 289, 7, 35, 2, 2, 289, 290, 5, 44, 23, 2, 290, 73, 3, 2, 2,
	2, 291, 292, 7, 41, 2, 2, 292, 293, 7, 4, 2, 2, 293, 294, 5, 44, 23, 2,
	294, 295, 7, 5, 2, 2, 295, 75, 3, 2, 2, 2, 296, 301, 5, 44, 23, 2, 297,
	298, 7, 9, 2, 2, 298, 300, 5, 44, 23, 2, 299, 297, 3, 2, 2, 2, 300, 303,
	3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 77, 3, 2,
	2, 2, 303, 301, 3, 2, 2, 2, 25, 81, 95, 107, 120, 129, 134, 143, 152, 167,
	172, 186, 205, 213, 219, 225, 234, 242, 246, 261, 269, 275, 284, 301,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'['", "']'", "'fun'", "'('", "')'", "','", "'return'", "'if'",
	"'else'", "'for'", "'in'", "'to'", "'{'", "'}'", "'break'", "'continue'",
	"'import'", "'||'", "'&&'", "'!'", "'<'", "'>'", "'=='", "'>='", "'<='",
	"'!='", "'+'", "'-'", "'*'", "'/'", "'%'", "':'", "", "", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
	"BOOL", "NULL", "INTEGER_LITERAL", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT",
	"LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "stmt", "assign_stmt", "left_side", "fundef_stmt", "fun_name",
	"fun_body", "fun_par", "namelist", "return_stmt", "funcall_stmt", "funcall_expr",
	"arg_expr", "arg_list", "arg", "if_stmt", "for_stmt", "block", "break_stmt",
	"continue_stmt", "import_stmt", "test", "or_test", "and_test", "not_test",
	"compare_expr", "comp_op", "arith_expr", "term", "factor", "atom_expr",
	"bracket_expr", "list_literal", "map_literal", "entry_list", "entry", "table_index",
	"test_list",
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
	SpartaParserT__31           = 32
	SpartaParserT__32           = 33
	SpartaParserSTRING          = 34
	SpartaParserBOOL            = 35
	SpartaParserNULL            = 36
	SpartaParserINTEGER_LITERAL = 37
	SpartaParserNUMBER_LITERAL  = 38
	SpartaParserIDENTIFIER      = 39
	SpartaParserCOMMENT         = 40
	SpartaParserLINE_COMMENT    = 41
	SpartaParserWS              = 42
	SpartaParserSHEBANG         = 43
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
	SpartaParserRULE_or_test       = 22
	SpartaParserRULE_and_test      = 23
	SpartaParserRULE_not_test      = 24
	SpartaParserRULE_compare_expr  = 25
	SpartaParserRULE_comp_op       = 26
	SpartaParserRULE_arith_expr    = 27
	SpartaParserRULE_term          = 28
	SpartaParserRULE_factor        = 29
	SpartaParserRULE_atom_expr     = 30
	SpartaParserRULE_bracket_expr  = 31
	SpartaParserRULE_list_literal  = 32
	SpartaParserRULE_map_literal   = 33
	SpartaParserRULE_entry_list    = 34
	SpartaParserRULE_entry         = 35
	SpartaParserRULE_table_index   = 36
	SpartaParserRULE_test_list     = 37
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__3)|(1<<SpartaParserT__7)|(1<<SpartaParserT__8)|(1<<SpartaParserT__10)|(1<<SpartaParserT__15)|(1<<SpartaParserT__16)|(1<<SpartaParserT__17))) != 0) || _la == SpartaParserIDENTIFIER {
		{
			p.SetState(76)
			p.Stmt()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Assign_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Fundef_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Return_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Funcall_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.If_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)
			p.For_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(90)
			p.Break_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(91)
			p.Continue_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(92)
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
		p.SetState(95)
		p.Left_side()
	}
	{
		p.SetState(96)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(97)
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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.Match(SpartaParserIDENTIFIER)
		}
		{
			p.SetState(101)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(102)
			p.Test()
		}
		{
			p.SetState(103)
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
		p.SetState(107)
		p.Match(SpartaParserT__3)
	}
	{
		p.SetState(108)
		p.Fun_name()
	}
	{
		p.SetState(109)
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
		p.SetState(111)
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
		p.SetState(113)
		p.Fun_par()
	}
	{
		p.SetState(114)
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
		p.SetState(116)
		p.Match(SpartaParserT__4)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserIDENTIFIER {
		{
			p.SetState(117)
			p.Namelist()
		}

	}
	{
		p.SetState(120)
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
		p.SetState(122)
		p.Match(SpartaParserIDENTIFIER)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(123)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(124)
			p.Match(SpartaParserIDENTIFIER)
		}

		p.SetState(129)
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
		p.SetState(130)
		p.Match(SpartaParserT__7)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(131)
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
		p.SetState(134)
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
		p.SetState(136)
		p.Fun_name()
	}
	{
		p.SetState(137)
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
		p.SetState(139)
		p.Match(SpartaParserT__4)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__20)|(1<<SpartaParserT__28))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SpartaParserSTRING-34))|(1<<(SpartaParserBOOL-34))|(1<<(SpartaParserNULL-34))|(1<<(SpartaParserINTEGER_LITERAL-34))|(1<<(SpartaParserNUMBER_LITERAL-34))|(1<<(SpartaParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(140)
			p.Arg_list()
		}

	}
	{
		p.SetState(143)
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
		p.SetState(145)
		p.Arg()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(146)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(147)
			p.Arg()
		}

		p.SetState(152)
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
		p.SetState(153)
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
		p.SetState(155)
		p.Match(SpartaParserT__8)
	}
	{
		p.SetState(156)
		p.Test()
	}
	{
		p.SetState(157)
		p.Block()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(158)
				p.Match(SpartaParserT__9)
			}
			{
				p.SetState(159)
				p.Match(SpartaParserT__8)
			}
			{
				p.SetState(160)
				p.Test()
			}
			{
				p.SetState(161)
				p.Block()
			}

		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__9 {
		{
			p.SetState(168)
			p.Match(SpartaParserT__9)
		}
		{
			p.SetState(169)
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
		p.SetState(172)
		p.Match(SpartaParserT__10)
	}
	{
		p.SetState(173)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(174)
		p.Match(SpartaParserT__11)
	}
	{
		p.SetState(175)
		p.Test()
	}
	{
		p.SetState(176)
		p.Match(SpartaParserT__12)
	}
	{
		p.SetState(177)
		p.Test()
	}
	{
		p.SetState(178)
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
		p.SetState(180)
		p.Match(SpartaParserT__13)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__3)|(1<<SpartaParserT__7)|(1<<SpartaParserT__8)|(1<<SpartaParserT__10)|(1<<SpartaParserT__15)|(1<<SpartaParserT__16)|(1<<SpartaParserT__17))) != 0) || _la == SpartaParserIDENTIFIER {
		{
			p.SetState(181)
			p.Stmt()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
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
		p.SetState(189)
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
		p.SetState(191)
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
		p.SetState(193)
		p.Match(SpartaParserT__17)
	}
	{
		p.SetState(194)
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
		p.SetState(196)
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

func (p *SpartaParser) Or_test() (localctx IOr_testContext) {
	localctx = NewOr_testContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SpartaParserRULE_or_test)
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
		p.And_test()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__18 {
		{
			p.SetState(199)
			p.Match(SpartaParserT__18)
		}
		{
			p.SetState(200)
			p.And_test()
		}

		p.SetState(205)
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
	p.EnterRule(localctx, 46, SpartaParserRULE_and_test)
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
		p.SetState(206)
		p.Not_test()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__19 {
		{
			p.SetState(207)
			p.Match(SpartaParserT__19)
		}
		{
			p.SetState(208)
			p.Not_test()
		}

		p.SetState(213)
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
	p.EnterRule(localctx, 48, SpartaParserRULE_not_test)

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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(SpartaParserT__20)
		}
		{
			p.SetState(215)
			p.Not_test()
		}

	case SpartaParserT__1, SpartaParserT__4, SpartaParserT__13, SpartaParserT__28, SpartaParserSTRING, SpartaParserBOOL, SpartaParserNULL, SpartaParserINTEGER_LITERAL, SpartaParserNUMBER_LITERAL, SpartaParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
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
	p.EnterRule(localctx, 50, SpartaParserRULE_compare_expr)
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
		p.SetState(219)
		p.Arith_expr()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25)|(1<<SpartaParserT__26))) != 0 {
		{
			p.SetState(220)
			p.Comp_op()
		}
		{
			p.SetState(221)
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
	p.EnterRule(localctx, 52, SpartaParserRULE_comp_op)
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
		p.SetState(225)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__21)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25)|(1<<SpartaParserT__26))) != 0) {
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
	p.EnterRule(localctx, 54, SpartaParserRULE_arith_expr)
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
		p.SetState(227)
		p.Term()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__27 || _la == SpartaParserT__28 {
		{
			p.SetState(228)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SpartaParserT__27 || _la == SpartaParserT__28) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(229)
			p.Term()
		}

		p.SetState(234)
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
	p.EnterRule(localctx, 56, SpartaParserRULE_term)
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
		p.SetState(235)
		p.Factor()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(SpartaParserT__29-30))|(1<<(SpartaParserT__30-30))|(1<<(SpartaParserT__31-30)))) != 0 {
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(SpartaParserT__29-30))|(1<<(SpartaParserT__30-30))|(1<<(SpartaParserT__31-30)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)
			p.Factor()
		}

		p.SetState(242)
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
	p.EnterRule(localctx, 58, SpartaParserRULE_factor)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__28 {
		{
			p.SetState(243)
			p.Match(SpartaParserT__28)
		}

	}
	{
		p.SetState(246)
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

func (s *Atom_exprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SpartaParserBOOL, 0)
}

func (s *Atom_exprContext) NULL() antlr.TerminalNode {
	return s.GetToken(SpartaParserNULL, 0)
}

func (s *Atom_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *SpartaParser) Atom_expr() (localctx IAtom_exprContext) {
	localctx = NewAtom_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SpartaParserRULE_atom_expr)

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

	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Bracket_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.List_literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)
			p.Map_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(251)
			p.Funcall_expr()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(252)
			p.Table_index()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(253)
			p.Match(SpartaParserIDENTIFIER)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(254)
			p.Match(SpartaParserINTEGER_LITERAL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(255)
			p.Match(SpartaParserNUMBER_LITERAL)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(256)
			p.Match(SpartaParserSTRING)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(257)
			p.Match(SpartaParserBOOL)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(258)
			p.Match(SpartaParserNULL)
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
	p.EnterRule(localctx, 62, SpartaParserRULE_bracket_expr)

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
		p.SetState(261)
		p.Match(SpartaParserT__4)
	}
	{
		p.SetState(262)
		p.Test()
	}
	{
		p.SetState(263)
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
	p.EnterRule(localctx, 64, SpartaParserRULE_list_literal)
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
		p.Match(SpartaParserT__1)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__20)|(1<<SpartaParserT__28))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SpartaParserSTRING-34))|(1<<(SpartaParserBOOL-34))|(1<<(SpartaParserNULL-34))|(1<<(SpartaParserINTEGER_LITERAL-34))|(1<<(SpartaParserNUMBER_LITERAL-34))|(1<<(SpartaParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(266)
			p.Test_list()
		}

	}
	{
		p.SetState(269)
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
	p.EnterRule(localctx, 66, SpartaParserRULE_map_literal)
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
		p.SetState(271)
		p.Match(SpartaParserT__13)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__1)|(1<<SpartaParserT__4)|(1<<SpartaParserT__13)|(1<<SpartaParserT__20)|(1<<SpartaParserT__28))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SpartaParserSTRING-34))|(1<<(SpartaParserBOOL-34))|(1<<(SpartaParserNULL-34))|(1<<(SpartaParserINTEGER_LITERAL-34))|(1<<(SpartaParserNUMBER_LITERAL-34))|(1<<(SpartaParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(272)
			p.Entry_list()
		}

	}
	{
		p.SetState(275)
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
	p.EnterRule(localctx, 68, SpartaParserRULE_entry_list)
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
		p.SetState(277)
		p.Entry()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(278)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(279)
			p.Entry()
		}

		p.SetState(284)
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
	p.EnterRule(localctx, 70, SpartaParserRULE_entry)

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
		p.SetState(285)
		p.Test()
	}
	{
		p.SetState(286)
		p.Match(SpartaParserT__32)
	}
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 72, SpartaParserRULE_table_index)

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
		p.SetState(289)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(290)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(291)
		p.Test()
	}
	{
		p.SetState(292)
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
	p.EnterRule(localctx, 74, SpartaParserRULE_test_list)
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
		p.SetState(294)
		p.Test()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__6 {
		{
			p.SetState(295)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(296)
			p.Test()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
