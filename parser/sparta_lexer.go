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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 61, 519,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 7, 49, 288, 10, 49, 12,
	49, 14, 49, 291, 11, 49, 3, 50, 3, 50, 3, 50, 7, 50, 296, 10, 50, 12, 50,
	14, 50, 299, 11, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 7, 51, 306, 10,
	51, 12, 51, 14, 51, 309, 11, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 7, 53, 323, 10, 53, 12, 53,
	14, 53, 326, 11, 53, 3, 53, 5, 53, 329, 10, 53, 3, 54, 6, 54, 332, 10,
	54, 13, 54, 14, 54, 333, 3, 55, 3, 55, 3, 55, 6, 55, 339, 10, 55, 13, 55,
	14, 55, 340, 3, 56, 6, 56, 344, 10, 56, 13, 56, 14, 56, 345, 3, 56, 3,
	56, 7, 56, 350, 10, 56, 12, 56, 14, 56, 353, 11, 56, 3, 56, 5, 56, 356,
	10, 56, 3, 56, 3, 56, 6, 56, 360, 10, 56, 13, 56, 14, 56, 361, 3, 56, 5,
	56, 365, 10, 56, 3, 56, 6, 56, 368, 10, 56, 13, 56, 14, 56, 369, 3, 56,
	3, 56, 5, 56, 374, 10, 56, 3, 57, 3, 57, 3, 57, 6, 57, 379, 10, 57, 13,
	57, 14, 57, 380, 3, 57, 3, 57, 7, 57, 385, 10, 57, 12, 57, 14, 57, 388,
	11, 57, 3, 57, 5, 57, 391, 10, 57, 3, 57, 3, 57, 3, 57, 3, 57, 6, 57, 397,
	10, 57, 13, 57, 14, 57, 398, 3, 57, 5, 57, 402, 10, 57, 3, 57, 3, 57, 3,
	57, 6, 57, 407, 10, 57, 13, 57, 14, 57, 408, 3, 57, 3, 57, 5, 57, 413,
	10, 57, 3, 58, 3, 58, 5, 58, 417, 10, 58, 3, 58, 6, 58, 420, 10, 58, 13,
	58, 14, 58, 421, 3, 59, 3, 59, 5, 59, 426, 10, 59, 3, 59, 6, 59, 429, 10,
	59, 13, 59, 14, 59, 430, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 437, 10, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 443, 10, 60, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61, 456, 10, 61,
	3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 6,
	63, 468, 10, 63, 13, 63, 14, 63, 469, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65,
	3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 7, 66, 482, 10, 66, 12, 66, 14, 66,
	485, 11, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3,
	67, 7, 67, 496, 10, 67, 12, 67, 14, 67, 499, 11, 67, 3, 67, 3, 67, 3, 68,
	6, 68, 504, 10, 68, 13, 68, 14, 68, 505, 3, 68, 3, 68, 3, 69, 3, 69, 3,
	69, 7, 69, 513, 10, 69, 12, 69, 14, 69, 516, 11, 69, 3, 69, 3, 69, 4, 324,
	483, 2, 70, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47,
	93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 2, 107, 54, 109,
	55, 111, 56, 113, 57, 115, 2, 117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127,
	2, 129, 2, 131, 58, 133, 59, 135, 60, 137, 61, 3, 2, 16, 5, 2, 67, 92,
	97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 36, 36, 94,
	94, 4, 2, 41, 41, 94, 94, 4, 2, 90, 90, 122, 122, 4, 2, 71, 71, 103, 103,
	4, 2, 45, 45, 47, 47, 4, 2, 82, 82, 114, 114, 12, 2, 36, 36, 41, 41, 94,
	94, 99, 100, 104, 104, 112, 112, 116, 116, 118, 118, 120, 120, 124, 124,
	3, 2, 50, 52, 3, 2, 50, 59, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 12, 12,
	15, 15, 5, 2, 11, 12, 14, 15, 34, 34, 2, 550, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2,
	2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137,
	3, 2, 2, 2, 3, 139, 3, 2, 2, 2, 5, 141, 3, 2, 2, 2, 7, 145, 3, 2, 2, 2,
	9, 147, 3, 2, 2, 2, 11, 153, 3, 2, 2, 2, 13, 156, 3, 2, 2, 2, 15, 158,
	3, 2, 2, 2, 17, 160, 3, 2, 2, 2, 19, 167, 3, 2, 2, 2, 21, 172, 3, 2, 2,
	2, 23, 178, 3, 2, 2, 2, 25, 182, 3, 2, 2, 2, 27, 184, 3, 2, 2, 2, 29, 191,
	3, 2, 2, 2, 31, 195, 3, 2, 2, 2, 33, 197, 3, 2, 2, 2, 35, 199, 3, 2, 2,
	2, 37, 203, 3, 2, 2, 2, 39, 209, 3, 2, 2, 2, 41, 214, 3, 2, 2, 2, 43, 216,
	3, 2, 2, 2, 45, 218, 3, 2, 2, 2, 47, 220, 3, 2, 2, 2, 49, 222, 3, 2, 2,
	2, 51, 226, 3, 2, 2, 2, 53, 229, 3, 2, 2, 2, 55, 233, 3, 2, 2, 2, 57, 235,
	3, 2, 2, 2, 59, 237, 3, 2, 2, 2, 61, 240, 3, 2, 2, 2, 63, 243, 3, 2, 2,
	2, 65, 246, 3, 2, 2, 2, 67, 249, 3, 2, 2, 2, 69, 252, 3, 2, 2, 2, 71, 254,
	3, 2, 2, 2, 73, 256, 3, 2, 2, 2, 75, 258, 3, 2, 2, 2, 77, 260, 3, 2, 2,
	2, 79, 262, 3, 2, 2, 2, 81, 265, 3, 2, 2, 2, 83, 267, 3, 2, 2, 2, 85, 269,
	3, 2, 2, 2, 87, 271, 3, 2, 2, 2, 89, 274, 3, 2, 2, 2, 91, 277, 3, 2, 2,
	2, 93, 281, 3, 2, 2, 2, 95, 283, 3, 2, 2, 2, 97, 285, 3, 2, 2, 2, 99, 292,
	3, 2, 2, 2, 101, 302, 3, 2, 2, 2, 103, 312, 3, 2, 2, 2, 105, 328, 3, 2,
	2, 2, 107, 331, 3, 2, 2, 2, 109, 335, 3, 2, 2, 2, 111, 373, 3, 2, 2, 2,
	113, 412, 3, 2, 2, 2, 115, 414, 3, 2, 2, 2, 117, 423, 3, 2, 2, 2, 119,
	442, 3, 2, 2, 2, 121, 455, 3, 2, 2, 2, 123, 457, 3, 2, 2, 2, 125, 462,
	3, 2, 2, 2, 127, 473, 3, 2, 2, 2, 129, 475, 3, 2, 2, 2, 131, 477, 3, 2,
	2, 2, 133, 491, 3, 2, 2, 2, 135, 503, 3, 2, 2, 2, 137, 509, 3, 2, 2, 2,
	139, 140, 7, 61, 2, 2, 140, 4, 3, 2, 2, 2, 141, 142, 7, 120, 2, 2, 142,
	143, 7, 99, 2, 2, 143, 144, 7, 116, 2, 2, 144, 6, 3, 2, 2, 2, 145, 146,
	7, 63, 2, 2, 146, 8, 3, 2, 2, 2, 147, 148, 7, 100, 2, 2, 148, 149, 7, 116,
	2, 2, 149, 150, 7, 103, 2, 2, 150, 151, 7, 99, 2, 2, 151, 152, 7, 109,
	2, 2, 152, 10, 3, 2, 2, 2, 153, 154, 7, 107, 2, 2, 154, 155, 7, 104, 2,
	2, 155, 12, 3, 2, 2, 2, 156, 157, 7, 125, 2, 2, 157, 14, 3, 2, 2, 2, 158,
	159, 7, 127, 2, 2, 159, 16, 3, 2, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162,
	7, 110, 2, 2, 162, 163, 7, 117, 2, 2, 163, 164, 7, 103, 2, 2, 164, 165,
	7, 107, 2, 2, 165, 166, 7, 104, 2, 2, 166, 18, 3, 2, 2, 2, 167, 168, 7,
	103, 2, 2, 168, 169, 7, 110, 2, 2, 169, 170, 7, 117, 2, 2, 170, 171, 7,
	103, 2, 2, 171, 20, 3, 2, 2, 2, 172, 173, 7, 121, 2, 2, 173, 174, 7, 106,
	2, 2, 174, 175, 7, 107, 2, 2, 175, 176, 7, 110, 2, 2, 176, 177, 7, 103,
	2, 2, 177, 22, 3, 2, 2, 2, 178, 179, 7, 104, 2, 2, 179, 180, 7, 113, 2,
	2, 180, 181, 7, 116, 2, 2, 181, 24, 3, 2, 2, 2, 182, 183, 7, 46, 2, 2,
	183, 26, 3, 2, 2, 2, 184, 185, 7, 116, 2, 2, 185, 186, 7, 103, 2, 2, 186,
	187, 7, 118, 2, 2, 187, 188, 7, 119, 2, 2, 188, 189, 7, 116, 2, 2, 189,
	190, 7, 112, 2, 2, 190, 28, 3, 2, 2, 2, 191, 192, 7, 104, 2, 2, 192, 193,
	7, 119, 2, 2, 193, 194, 7, 112, 2, 2, 194, 30, 3, 2, 2, 2, 195, 196, 7,
	48, 2, 2, 196, 32, 3, 2, 2, 2, 197, 198, 7, 60, 2, 2, 198, 34, 3, 2, 2,
	2, 199, 200, 7, 112, 2, 2, 200, 201, 7, 107, 2, 2, 201, 202, 7, 110, 2,
	2, 202, 36, 3, 2, 2, 2, 203, 204, 7, 104, 2, 2, 204, 205, 7, 99, 2, 2,
	205, 206, 7, 110, 2, 2, 206, 207, 7, 117, 2, 2, 207, 208, 7, 103, 2, 2,
	208, 38, 3, 2, 2, 2, 209, 210, 7, 118, 2, 2, 210, 211, 7, 116, 2, 2, 211,
	212, 7, 119, 2, 2, 212, 213, 7, 103, 2, 2, 213, 40, 3, 2, 2, 2, 214, 215,
	7, 42, 2, 2, 215, 42, 3, 2, 2, 2, 216, 217, 7, 43, 2, 2, 217, 44, 3, 2,
	2, 2, 218, 219, 7, 93, 2, 2, 219, 46, 3, 2, 2, 2, 220, 221, 7, 95, 2, 2,
	221, 48, 3, 2, 2, 2, 222, 223, 7, 48, 2, 2, 223, 224, 7, 48, 2, 2, 224,
	225, 7, 48, 2, 2, 225, 50, 3, 2, 2, 2, 226, 227, 7, 113, 2, 2, 227, 228,
	7, 116, 2, 2, 228, 52, 3, 2, 2, 2, 229, 230, 7, 99, 2, 2, 230, 231, 7,
	112, 2, 2, 231, 232, 7, 102, 2, 2, 232, 54, 3, 2, 2, 2, 233, 234, 7, 62,
	2, 2, 234, 56, 3, 2, 2, 2, 235, 236, 7, 64, 2, 2, 236, 58, 3, 2, 2, 2,
	237, 238, 7, 62, 2, 2, 238, 239, 7, 63, 2, 2, 239, 60, 3, 2, 2, 2, 240,
	241, 7, 64, 2, 2, 241, 242, 7, 63, 2, 2, 242, 62, 3, 2, 2, 2, 243, 244,
	7, 128, 2, 2, 244, 245, 7, 63, 2, 2, 245, 64, 3, 2, 2, 2, 246, 247, 7,
	63, 2, 2, 247, 248, 7, 63, 2, 2, 248, 66, 3, 2, 2, 2, 249, 250, 7, 48,
	2, 2, 250, 251, 7, 48, 2, 2, 251, 68, 3, 2, 2, 2, 252, 253, 7, 45, 2, 2,
	253, 70, 3, 2, 2, 2, 254, 255, 7, 47, 2, 2, 255, 72, 3, 2, 2, 2, 256, 257,
	7, 44, 2, 2, 257, 74, 3, 2, 2, 2, 258, 259, 7, 49, 2, 2, 259, 76, 3, 2,
	2, 2, 260, 261, 7, 39, 2, 2, 261, 78, 3, 2, 2, 2, 262, 263, 7, 49, 2, 2,
	263, 264, 7, 49, 2, 2, 264, 80, 3, 2, 2, 2, 265, 266, 7, 40, 2, 2, 266,
	82, 3, 2, 2, 2, 267, 268, 7, 126, 2, 2, 268, 84, 3, 2, 2, 2, 269, 270,
	7, 128, 2, 2, 270, 86, 3, 2, 2, 2, 271, 272, 7, 62, 2, 2, 272, 273, 7,
	62, 2, 2, 273, 88, 3, 2, 2, 2, 274, 275, 7, 64, 2, 2, 275, 276, 7, 64,
	2, 2, 276, 90, 3, 2, 2, 2, 277, 278, 7, 112, 2, 2, 278, 279, 7, 113, 2,
	2, 279, 280, 7, 118, 2, 2, 280, 92, 3, 2, 2, 2, 281, 282, 7, 37, 2, 2,
	282, 94, 3, 2, 2, 2, 283, 284, 7, 96, 2, 2, 284, 96, 3, 2, 2, 2, 285, 289,
	9, 2, 2, 2, 286, 288, 9, 3, 2, 2, 287, 286, 3, 2, 2, 2, 288, 291, 3, 2,
	2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 98, 3, 2, 2, 2,
	291, 289, 3, 2, 2, 2, 292, 297, 7, 36, 2, 2, 293, 296, 5, 119, 60, 2, 294,
	296, 10, 4, 2, 2, 295, 293, 3, 2, 2, 2, 295, 294, 3, 2, 2, 2, 296, 299,
	3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 300, 3, 2,
	2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 36, 2, 2, 301, 100, 3, 2, 2, 2,
	302, 307, 7, 41, 2, 2, 303, 306, 5, 119, 60, 2, 304, 306, 10, 5, 2, 2,
	305, 303, 3, 2, 2, 2, 305, 304, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307,
	305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 310, 3, 2, 2, 2, 309, 307,
	3, 2, 2, 2, 310, 311, 7, 41, 2, 2, 311, 102, 3, 2, 2, 2, 312, 313, 7, 93,
	2, 2, 313, 314, 5, 105, 53, 2, 314, 315, 7, 95, 2, 2, 315, 104, 3, 2, 2,
	2, 316, 317, 7, 63, 2, 2, 317, 318, 5, 105, 53, 2, 318, 319, 7, 63, 2,
	2, 319, 329, 3, 2, 2, 2, 320, 324, 7, 93, 2, 2, 321, 323, 11, 2, 2, 2,
	322, 321, 3, 2, 2, 2, 323, 326, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 324,
	322, 3, 2, 2, 2, 325, 327, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 327, 329,
	7, 95, 2, 2, 328, 316, 3, 2, 2, 2, 328, 320, 3, 2, 2, 2, 329, 106, 3, 2,
	2, 2, 330, 332, 5, 127, 64, 2, 331, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2,
	2, 333, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 108, 3, 2, 2, 2, 335,
	336, 7, 50, 2, 2, 336, 338, 9, 6, 2, 2, 337, 339, 5, 129, 65, 2, 338, 337,
	3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 340, 341, 3, 2,
	2, 2, 341, 110, 3, 2, 2, 2, 342, 344, 5, 127, 64, 2, 343, 342, 3, 2, 2,
	2, 344, 345, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 351, 7, 48, 2, 2, 348, 350, 5, 127, 64, 2, 349, 348,
	3, 2, 2, 2, 350, 353, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2,
	2, 2, 352, 355, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 354, 356, 5, 115, 58,
	2, 355, 354, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 374, 3, 2, 2, 2, 357,
	359, 7, 48, 2, 2, 358, 360, 5, 127, 64, 2, 359, 358, 3, 2, 2, 2, 360, 361,
	3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 364, 3, 2,
	2, 2, 363, 365, 5, 115, 58, 2, 364, 363, 3, 2, 2, 2, 364, 365, 3, 2, 2,
	2, 365, 374, 3, 2, 2, 2, 366, 368, 5, 127, 64, 2, 367, 366, 3, 2, 2, 2,
	368, 369, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370,
	371, 3, 2, 2, 2, 371, 372, 5, 115, 58, 2, 372, 374, 3, 2, 2, 2, 373, 343,
	3, 2, 2, 2, 373, 357, 3, 2, 2, 2, 373, 367, 3, 2, 2, 2, 374, 112, 3, 2,
	2, 2, 375, 376, 7, 50, 2, 2, 376, 378, 9, 6, 2, 2, 377, 379, 5, 129, 65,
	2, 378, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380,
	381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 386, 7, 48, 2, 2, 383, 385,
	5, 129, 65, 2, 384, 383, 3, 2, 2, 2, 385, 388, 3, 2, 2, 2, 386, 384, 3,
	2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 390, 3, 2, 2, 2, 388, 386, 3, 2, 2,
	2, 389, 391, 5, 117, 59, 2, 390, 389, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2,
	391, 413, 3, 2, 2, 2, 392, 393, 7, 50, 2, 2, 393, 394, 9, 6, 2, 2, 394,
	396, 7, 48, 2, 2, 395, 397, 5, 129, 65, 2, 396, 395, 3, 2, 2, 2, 397, 398,
	3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 401, 3, 2,
	2, 2, 400, 402, 5, 117, 59, 2, 401, 400, 3, 2, 2, 2, 401, 402, 3, 2, 2,
	2, 402, 413, 3, 2, 2, 2, 403, 404, 7, 50, 2, 2, 404, 406, 9, 6, 2, 2, 405,
	407, 5, 129, 65, 2, 406, 405, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 406,
	3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 411, 5, 117,
	59, 2, 411, 413, 3, 2, 2, 2, 412, 375, 3, 2, 2, 2, 412, 392, 3, 2, 2, 2,
	412, 403, 3, 2, 2, 2, 413, 114, 3, 2, 2, 2, 414, 416, 9, 7, 2, 2, 415,
	417, 9, 8, 2, 2, 416, 415, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 419,
	3, 2, 2, 2, 418, 420, 5, 127, 64, 2, 419, 418, 3, 2, 2, 2, 420, 421, 3,
	2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 116, 3, 2, 2,
	2, 423, 425, 9, 9, 2, 2, 424, 426, 9, 8, 2, 2, 425, 424, 3, 2, 2, 2, 425,
	426, 3, 2, 2, 2, 426, 428, 3, 2, 2, 2, 427, 429, 5, 127, 64, 2, 428, 427,
	3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2,
	2, 2, 431, 118, 3, 2, 2, 2, 432, 433, 7, 94, 2, 2, 433, 443, 9, 10, 2,
	2, 434, 436, 7, 94, 2, 2, 435, 437, 7, 15, 2, 2, 436, 435, 3, 2, 2, 2,
	436, 437, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 443, 7, 12, 2, 2, 439,
	443, 5, 121, 61, 2, 440, 443, 5, 123, 62, 2, 441, 443, 5, 125, 63, 2, 442,
	432, 3, 2, 2, 2, 442, 434, 3, 2, 2, 2, 442, 439, 3, 2, 2, 2, 442, 440,
	3, 2, 2, 2, 442, 441, 3, 2, 2, 2, 443, 120, 3, 2, 2, 2, 444, 445, 7, 94,
	2, 2, 445, 456, 5, 127, 64, 2, 446, 447, 7, 94, 2, 2, 447, 448, 5, 127,
	64, 2, 448, 449, 5, 127, 64, 2, 449, 456, 3, 2, 2, 2, 450, 451, 7, 94,
	2, 2, 451, 452, 9, 11, 2, 2, 452, 453, 5, 127, 64, 2, 453, 454, 5, 127,
	64, 2, 454, 456, 3, 2, 2, 2, 455, 444, 3, 2, 2, 2, 455, 446, 3, 2, 2, 2,
	455, 450, 3, 2, 2, 2, 456, 122, 3, 2, 2, 2, 457, 458, 7, 94, 2, 2, 458,
	459, 7, 122, 2, 2, 459, 460, 5, 129, 65, 2, 460, 461, 5, 129, 65, 2, 461,
	124, 3, 2, 2, 2, 462, 463, 7, 94, 2, 2, 463, 464, 7, 119, 2, 2, 464, 465,
	7, 125, 2, 2, 465, 467, 3, 2, 2, 2, 466, 468, 5, 129, 65, 2, 467, 466,
	3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 469, 470, 3, 2,
	2, 2, 470, 471, 3, 2, 2, 2, 471, 472, 7, 127, 2, 2, 472, 126, 3, 2, 2,
	2, 473, 474, 9, 12, 2, 2, 474, 128, 3, 2, 2, 2, 475, 476, 9, 13, 2, 2,
	476, 130, 3, 2, 2, 2, 477, 478, 7, 49, 2, 2, 478, 479, 7, 44, 2, 2, 479,
	483, 3, 2, 2, 2, 480, 482, 11, 2, 2, 2, 481, 480, 3, 2, 2, 2, 482, 485,
	3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 484, 486, 3, 2,
	2, 2, 485, 483, 3, 2, 2, 2, 486, 487, 7, 44, 2, 2, 487, 488, 7, 49, 2,
	2, 488, 489, 3, 2, 2, 2, 489, 490, 8, 66, 2, 2, 490, 132, 3, 2, 2, 2, 491,
	492, 7, 49, 2, 2, 492, 493, 7, 49, 2, 2, 493, 497, 3, 2, 2, 2, 494, 496,
	10, 14, 2, 2, 495, 494, 3, 2, 2, 2, 496, 499, 3, 2, 2, 2, 497, 495, 3,
	2, 2, 2, 497, 498, 3, 2, 2, 2, 498, 500, 3, 2, 2, 2, 499, 497, 3, 2, 2,
	2, 500, 501, 8, 67, 2, 2, 501, 134, 3, 2, 2, 2, 502, 504, 9, 15, 2, 2,
	503, 502, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 505,
	506, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 508, 8, 68, 2, 2, 508, 136,
	3, 2, 2, 2, 509, 510, 7, 37, 2, 2, 510, 514, 7, 35, 2, 2, 511, 513, 10,
	14, 2, 2, 512, 511, 3, 2, 2, 2, 513, 516, 3, 2, 2, 2, 514, 512, 3, 2, 2,
	2, 514, 515, 3, 2, 2, 2, 515, 517, 3, 2, 2, 2, 516, 514, 3, 2, 2, 2, 517,
	518, 8, 69, 3, 2, 518, 138, 3, 2, 2, 2, 38, 2, 289, 295, 297, 305, 307,
	324, 328, 333, 340, 345, 351, 355, 361, 364, 369, 373, 380, 386, 390, 398,
	401, 408, 412, 416, 421, 425, 430, 436, 442, 455, 469, 483, 497, 505, 514,
	4, 8, 2, 2, 2, 3, 2,
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
	"", "';'", "'var'", "'='", "'break'", "'if'", "'{'", "'}'", "'elseif'",
	"'else'", "'while'", "'for'", "','", "'return'", "'fun'", "'.'", "':'",
	"'nil'", "'false'", "'true'", "'('", "')'", "'['", "']'", "'...'", "'or'",
	"'and'", "'<'", "'>'", "'<='", "'>='", "'~='", "'=='", "'..'", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'//'", "'&'", "'|'", "'~'", "'<<'", "'>>'", "'not'",
	"'#'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "NAME", "NORMALSTRING",
	"CHARSTRING", "LONGSTRING", "INT", "HEX", "FLOAT", "HEX_FLOAT", "COMMENT",
	"LINE_COMMENT", "WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "NAME", "NORMALSTRING",
	"CHARSTRING", "LONGSTRING", "NESTED_STR", "INT", "HEX", "FLOAT", "HEX_FLOAT",
	"ExponentPart", "HexExponentPart", "EscapeSequence", "DecimalEscape", "HexEscape",
	"UtfEscape", "Digit", "HexDigit", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
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
	SpartaLexerT__0         = 1
	SpartaLexerT__1         = 2
	SpartaLexerT__2         = 3
	SpartaLexerT__3         = 4
	SpartaLexerT__4         = 5
	SpartaLexerT__5         = 6
	SpartaLexerT__6         = 7
	SpartaLexerT__7         = 8
	SpartaLexerT__8         = 9
	SpartaLexerT__9         = 10
	SpartaLexerT__10        = 11
	SpartaLexerT__11        = 12
	SpartaLexerT__12        = 13
	SpartaLexerT__13        = 14
	SpartaLexerT__14        = 15
	SpartaLexerT__15        = 16
	SpartaLexerT__16        = 17
	SpartaLexerT__17        = 18
	SpartaLexerT__18        = 19
	SpartaLexerT__19        = 20
	SpartaLexerT__20        = 21
	SpartaLexerT__21        = 22
	SpartaLexerT__22        = 23
	SpartaLexerT__23        = 24
	SpartaLexerT__24        = 25
	SpartaLexerT__25        = 26
	SpartaLexerT__26        = 27
	SpartaLexerT__27        = 28
	SpartaLexerT__28        = 29
	SpartaLexerT__29        = 30
	SpartaLexerT__30        = 31
	SpartaLexerT__31        = 32
	SpartaLexerT__32        = 33
	SpartaLexerT__33        = 34
	SpartaLexerT__34        = 35
	SpartaLexerT__35        = 36
	SpartaLexerT__36        = 37
	SpartaLexerT__37        = 38
	SpartaLexerT__38        = 39
	SpartaLexerT__39        = 40
	SpartaLexerT__40        = 41
	SpartaLexerT__41        = 42
	SpartaLexerT__42        = 43
	SpartaLexerT__43        = 44
	SpartaLexerT__44        = 45
	SpartaLexerT__45        = 46
	SpartaLexerT__46        = 47
	SpartaLexerNAME         = 48
	SpartaLexerNORMALSTRING = 49
	SpartaLexerCHARSTRING   = 50
	SpartaLexerLONGSTRING   = 51
	SpartaLexerINT          = 52
	SpartaLexerHEX          = 53
	SpartaLexerFLOAT        = 54
	SpartaLexerHEX_FLOAT    = 55
	SpartaLexerCOMMENT      = 56
	SpartaLexerLINE_COMMENT = 57
	SpartaLexerWS           = 58
	SpartaLexerSHEBANG      = 59
)
