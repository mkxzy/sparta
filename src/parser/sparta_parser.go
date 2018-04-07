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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 403,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 77, 10,
	3, 12, 3, 14, 3, 80, 11, 3, 3, 3, 5, 3, 83, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 119, 10, 4, 12, 4, 14, 4,
	122, 11, 4, 3, 4, 3, 4, 5, 4, 126, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 138, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 164, 10, 4, 5, 4,
	166, 10, 4, 3, 5, 3, 5, 5, 5, 170, 10, 5, 3, 5, 5, 5, 173, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 182, 10, 7, 12, 7, 14, 7, 185,
	11, 7, 3, 7, 3, 7, 5, 7, 189, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 194, 10, 8,
	12, 8, 14, 8, 197, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 202, 10, 9, 12, 9, 14,
	9, 205, 11, 9, 3, 10, 3, 10, 3, 10, 7, 10, 210, 10, 10, 12, 10, 14, 10,
	213, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 228, 10, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 262, 10,
	11, 12, 11, 14, 11, 265, 11, 11, 3, 12, 3, 12, 7, 12, 269, 10, 12, 12,
	12, 14, 12, 272, 11, 12, 3, 13, 3, 13, 6, 13, 276, 10, 13, 13, 13, 14,
	13, 277, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 285, 10, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 293, 10, 15, 3, 15, 7, 15, 296,
	10, 15, 12, 15, 14, 15, 299, 11, 15, 3, 16, 7, 16, 302, 10, 16, 12, 16,
	14, 16, 305, 11, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 313,
	10, 16, 3, 17, 3, 17, 5, 17, 317, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5,
	18, 323, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 328, 10, 18, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 5, 20, 335, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 5, 21, 344, 10, 21, 3, 21, 5, 21, 347, 10, 21, 3, 22,
	3, 22, 5, 22, 351, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 7,
	23, 359, 10, 23, 12, 23, 14, 23, 362, 11, 23, 3, 23, 5, 23, 365, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 377, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 2, 3, 20, 37, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 10, 4, 2, 3, 3,
	17, 17, 3, 2, 37, 42, 3, 2, 44, 45, 3, 2, 46, 49, 3, 2, 50, 54, 5, 2, 45,
	45, 52, 52, 55, 56, 3, 2, 62, 65, 3, 2, 59, 61, 2, 430, 2, 72, 3, 2, 2,
	2, 4, 78, 3, 2, 2, 2, 6, 165, 3, 2, 2, 2, 8, 167, 3, 2, 2, 2, 10, 174,
	3, 2, 2, 2, 12, 178, 3, 2, 2, 2, 14, 190, 3, 2, 2, 2, 16, 198, 3, 2, 2,
	2, 18, 206, 3, 2, 2, 2, 20, 227, 3, 2, 2, 2, 22, 266, 3, 2, 2, 2, 24, 273,
	3, 2, 2, 2, 26, 284, 3, 2, 2, 2, 28, 292, 3, 2, 2, 2, 30, 303, 3, 2, 2,
	2, 32, 316, 3, 2, 2, 2, 34, 327, 3, 2, 2, 2, 36, 329, 3, 2, 2, 2, 38, 332,
	3, 2, 2, 2, 40, 346, 3, 2, 2, 2, 42, 348, 3, 2, 2, 2, 44, 354, 3, 2, 2,
	2, 46, 376, 3, 2, 2, 2, 48, 378, 3, 2, 2, 2, 50, 380, 3, 2, 2, 2, 52, 382,
	3, 2, 2, 2, 54, 384, 3, 2, 2, 2, 56, 386, 3, 2, 2, 2, 58, 388, 3, 2, 2,
	2, 60, 390, 3, 2, 2, 2, 62, 392, 3, 2, 2, 2, 64, 394, 3, 2, 2, 2, 66, 396,
	3, 2, 2, 2, 68, 398, 3, 2, 2, 2, 70, 400, 3, 2, 2, 2, 72, 73, 5, 4, 3,
	2, 73, 74, 7, 2, 2, 3, 74, 3, 3, 2, 2, 2, 75, 77, 5, 6, 4, 2, 76, 75, 3,
	2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79,
	82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 83, 5, 8, 5, 2, 82, 81, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 166, 7, 3, 2, 2, 85, 86,
	5, 14, 8, 2, 86, 87, 7, 4, 2, 2, 87, 88, 5, 18, 10, 2, 88, 166, 3, 2, 2,
	2, 89, 166, 5, 24, 13, 2, 90, 166, 5, 10, 6, 2, 91, 166, 7, 5, 2, 2, 92,
	93, 7, 6, 2, 2, 93, 166, 7, 58, 2, 2, 94, 95, 7, 7, 2, 2, 95, 96, 5, 4,
	3, 2, 96, 97, 7, 8, 2, 2, 97, 166, 3, 2, 2, 2, 98, 99, 7, 9, 2, 2, 99,
	100, 5, 20, 11, 2, 100, 101, 7, 7, 2, 2, 101, 102, 5, 4, 3, 2, 102, 103,
	7, 8, 2, 2, 103, 166, 3, 2, 2, 2, 104, 105, 7, 10, 2, 2, 105, 106, 5, 4,
	3, 2, 106, 107, 7, 11, 2, 2, 107, 108, 5, 20, 11, 2, 108, 166, 3, 2, 2,
	2, 109, 110, 7, 12, 2, 2, 110, 111, 5, 20, 11, 2, 111, 112, 7, 13, 2, 2,
	112, 120, 5, 4, 3, 2, 113, 114, 7, 14, 2, 2, 114, 115, 5, 20, 11, 2, 115,
	116, 7, 13, 2, 2, 116, 117, 5, 4, 3, 2, 117, 119, 3, 2, 2, 2, 118, 113,
	3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 125, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 15, 2, 2,
	124, 126, 5, 4, 3, 2, 125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	127, 3, 2, 2, 2, 127, 128, 7, 8, 2, 2, 128, 166, 3, 2, 2, 2, 129, 130,
	7, 16, 2, 2, 130, 131, 7, 58, 2, 2, 131, 132, 7, 4, 2, 2, 132, 133, 5,
	20, 11, 2, 133, 134, 7, 17, 2, 2, 134, 137, 5, 20, 11, 2, 135, 136, 7,
	17, 2, 2, 136, 138, 5, 20, 11, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2,
	2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 7, 7, 2, 2, 140, 141, 5, 4, 3, 2,
	141, 142, 7, 8, 2, 2, 142, 166, 3, 2, 2, 2, 143, 144, 7, 16, 2, 2, 144,
	145, 5, 16, 9, 2, 145, 146, 7, 18, 2, 2, 146, 147, 5, 18, 10, 2, 147, 148,
	7, 7, 2, 2, 148, 149, 5, 4, 3, 2, 149, 150, 7, 8, 2, 2, 150, 166, 3, 2,
	2, 2, 151, 152, 7, 19, 2, 2, 152, 153, 5, 12, 7, 2, 153, 154, 5, 38, 20,
	2, 154, 166, 3, 2, 2, 2, 155, 156, 7, 20, 2, 2, 156, 157, 7, 19, 2, 2,
	157, 158, 7, 58, 2, 2, 158, 166, 5, 38, 20, 2, 159, 160, 7, 20, 2, 2, 160,
	163, 5, 16, 9, 2, 161, 162, 7, 4, 2, 2, 162, 164, 5, 18, 10, 2, 163, 161,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 84, 3, 2,
	2, 2, 165, 85, 3, 2, 2, 2, 165, 89, 3, 2, 2, 2, 165, 90, 3, 2, 2, 2, 165,
	91, 3, 2, 2, 2, 165, 92, 3, 2, 2, 2, 165, 94, 3, 2, 2, 2, 165, 98, 3, 2,
	2, 2, 165, 104, 3, 2, 2, 2, 165, 109, 3, 2, 2, 2, 165, 129, 3, 2, 2, 2,
	165, 143, 3, 2, 2, 2, 165, 151, 3, 2, 2, 2, 165, 155, 3, 2, 2, 2, 165,
	159, 3, 2, 2, 2, 166, 7, 3, 2, 2, 2, 167, 169, 7, 21, 2, 2, 168, 170, 5,
	18, 10, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2,
	2, 2, 171, 173, 7, 3, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2,
	173, 9, 3, 2, 2, 2, 174, 175, 7, 22, 2, 2, 175, 176, 7, 58, 2, 2, 176,
	177, 7, 22, 2, 2, 177, 11, 3, 2, 2, 2, 178, 183, 7, 58, 2, 2, 179, 180,
	7, 23, 2, 2, 180, 182, 7, 58, 2, 2, 181, 179, 3, 2, 2, 2, 182, 185, 3,
	2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 188, 3, 2, 2,
	2, 185, 183, 3, 2, 2, 2, 186, 187, 7, 24, 2, 2, 187, 189, 7, 58, 2, 2,
	188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 13, 3, 2, 2, 2, 190, 195,
	5, 28, 15, 2, 191, 192, 7, 17, 2, 2, 192, 194, 5, 28, 15, 2, 193, 191,
	3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2,
	2, 2, 196, 15, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198, 203, 7, 58, 2, 2,
	199, 200, 7, 17, 2, 2, 200, 202, 7, 58, 2, 2, 201, 199, 3, 2, 2, 2, 202,
	205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 17, 3,
	2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 211, 5, 20, 11, 2, 207, 208, 7, 17,
	2, 2, 208, 210, 5, 20, 11, 2, 209, 207, 3, 2, 2, 2, 210, 213, 3, 2, 2,
	2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 19, 3, 2, 2, 2, 213,
	211, 3, 2, 2, 2, 214, 215, 8, 11, 1, 2, 215, 228, 7, 25, 2, 2, 216, 228,
	7, 26, 2, 2, 217, 228, 7, 27, 2, 2, 218, 228, 5, 68, 35, 2, 219, 228, 5,
	70, 36, 2, 220, 228, 7, 28, 2, 2, 221, 228, 5, 36, 19, 2, 222, 228, 5,
	22, 12, 2, 223, 228, 5, 42, 22, 2, 224, 225, 5, 64, 33, 2, 225, 226, 5,
	20, 11, 10, 226, 228, 3, 2, 2, 2, 227, 214, 3, 2, 2, 2, 227, 216, 3, 2,
	2, 2, 227, 217, 3, 2, 2, 2, 227, 218, 3, 2, 2, 2, 227, 219, 3, 2, 2, 2,
	227, 220, 3, 2, 2, 2, 227, 221, 3, 2, 2, 2, 227, 222, 3, 2, 2, 2, 227,
	223, 3, 2, 2, 2, 227, 224, 3, 2, 2, 2, 228, 263, 3, 2, 2, 2, 229, 230,
	12, 11, 2, 2, 230, 231, 5, 66, 34, 2, 231, 232, 5, 20, 11, 11, 232, 262,
	3, 2, 2, 2, 233, 234, 12, 9, 2, 2, 234, 235, 5, 60, 31, 2, 235, 236, 5,
	20, 11, 10, 236, 262, 3, 2, 2, 2, 237, 238, 12, 8, 2, 2, 238, 239, 5, 58,
	30, 2, 239, 240, 5, 20, 11, 9, 240, 262, 3, 2, 2, 2, 241, 242, 12, 7, 2,
	2, 242, 243, 5, 56, 29, 2, 243, 244, 5, 20, 11, 7, 244, 262, 3, 2, 2, 2,
	245, 246, 12, 6, 2, 2, 246, 247, 5, 54, 28, 2, 247, 248, 5, 20, 11, 7,
	248, 262, 3, 2, 2, 2, 249, 250, 12, 5, 2, 2, 250, 251, 5, 52, 27, 2, 251,
	252, 5, 20, 11, 6, 252, 262, 3, 2, 2, 2, 253, 254, 12, 4, 2, 2, 254, 255,
	5, 50, 26, 2, 255, 256, 5, 20, 11, 5, 256, 262, 3, 2, 2, 2, 257, 258, 12,
	3, 2, 2, 258, 259, 5, 62, 32, 2, 259, 260, 5, 20, 11, 4, 260, 262, 3, 2,
	2, 2, 261, 229, 3, 2, 2, 2, 261, 233, 3, 2, 2, 2, 261, 237, 3, 2, 2, 2,
	261, 241, 3, 2, 2, 2, 261, 245, 3, 2, 2, 2, 261, 249, 3, 2, 2, 2, 261,
	253, 3, 2, 2, 2, 261, 257, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 261,
	3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 21, 3, 2, 2, 2, 265, 263, 3, 2,
	2, 2, 266, 270, 5, 26, 14, 2, 267, 269, 5, 32, 17, 2, 268, 267, 3, 2, 2,
	2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271,
	23, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 273, 275, 5, 26, 14, 2, 274, 276,
	5, 32, 17, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275, 3,
	2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 25, 3, 2, 2, 2, 279, 285, 5, 28, 15,
	2, 280, 281, 7, 29, 2, 2, 281, 282, 5, 20, 11, 2, 282, 283, 7, 30, 2, 2,
	283, 285, 3, 2, 2, 2, 284, 279, 3, 2, 2, 2, 284, 280, 3, 2, 2, 2, 285,
	27, 3, 2, 2, 2, 286, 293, 7, 58, 2, 2, 287, 288, 7, 29, 2, 2, 288, 289,
	5, 20, 11, 2, 289, 290, 7, 30, 2, 2, 290, 291, 5, 30, 16, 2, 291, 293,
	3, 2, 2, 2, 292, 286, 3, 2, 2, 2, 292, 287, 3, 2, 2, 2, 293, 297, 3, 2,
	2, 2, 294, 296, 5, 30, 16, 2, 295, 294, 3, 2, 2, 2, 296, 299, 3, 2, 2,
	2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 29, 3, 2, 2, 2, 299,
	297, 3, 2, 2, 2, 300, 302, 5, 32, 17, 2, 301, 300, 3, 2, 2, 2, 302, 305,
	3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 312, 3, 2,
	2, 2, 305, 303, 3, 2, 2, 2, 306, 307, 7, 31, 2, 2, 307, 308, 5, 20, 11,
	2, 308, 309, 7, 32, 2, 2, 309, 313, 3, 2, 2, 2, 310, 311, 7, 23, 2, 2,
	311, 313, 7, 58, 2, 2, 312, 306, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 313,
	31, 3, 2, 2, 2, 314, 315, 7, 24, 2, 2, 315, 317, 7, 58, 2, 2, 316, 314,
	3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 5, 34,
	18, 2, 319, 33, 3, 2, 2, 2, 320, 322, 7, 29, 2, 2, 321, 323, 5, 18, 10,
	2, 322, 321, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324,
	328, 7, 30, 2, 2, 325, 328, 5, 42, 22, 2, 326, 328, 5, 70, 36, 2, 327,
	320, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 326, 3, 2, 2, 2, 328, 35, 3,
	2, 2, 2, 329, 330, 7, 19, 2, 2, 330, 331, 5, 38, 20, 2, 331, 37, 3, 2,
	2, 2, 332, 334, 7, 29, 2, 2, 333, 335, 5, 40, 21, 2, 334, 333, 3, 2, 2,
	2, 334, 335, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 337, 7, 30, 2, 2, 337,
	338, 5, 4, 3, 2, 338, 339, 7, 8, 2, 2, 339, 39, 3, 2, 2, 2, 340, 343, 5,
	16, 9, 2, 341, 342, 7, 17, 2, 2, 342, 344, 7, 28, 2, 2, 343, 341, 3, 2,
	2, 2, 343, 344, 3, 2, 2, 2, 344, 347, 3, 2, 2, 2, 345, 347, 7, 28, 2, 2,
	346, 340, 3, 2, 2, 2, 346, 345, 3, 2, 2, 2, 347, 41, 3, 2, 2, 2, 348, 350,
	7, 33, 2, 2, 349, 351, 5, 44, 23, 2, 350, 349, 3, 2, 2, 2, 350, 351, 3,
	2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 353, 7, 34, 2, 2, 353, 43, 3, 2, 2,
	2, 354, 360, 5, 46, 24, 2, 355, 356, 5, 48, 25, 2, 356, 357, 5, 46, 24,
	2, 357, 359, 3, 2, 2, 2, 358, 355, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360,
	358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 363, 365, 5, 48, 25, 2, 364, 363, 3, 2, 2, 2, 364, 365, 3,
	2, 2, 2, 365, 45, 3, 2, 2, 2, 366, 367, 7, 31, 2, 2, 367, 368, 5, 20, 11,
	2, 368, 369, 7, 32, 2, 2, 369, 370, 7, 4, 2, 2, 370, 371, 5, 20, 11, 2,
	371, 377, 3, 2, 2, 2, 372, 373, 7, 58, 2, 2, 373, 374, 7, 4, 2, 2, 374,
	377, 5, 20, 11, 2, 375, 377, 5, 20, 11, 2, 376, 366, 3, 2, 2, 2, 376, 372,
	3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 47, 3, 2, 2, 2, 378, 379, 9, 2,
	2, 2, 379, 49, 3, 2, 2, 2, 380, 381, 7, 35, 2, 2, 381, 51, 3, 2, 2, 2,
	382, 383, 7, 36, 2, 2, 383, 53, 3, 2, 2, 2, 384, 385, 9, 3, 2, 2, 385,
	55, 3, 2, 2, 2, 386, 387, 7, 43, 2, 2, 387, 57, 3, 2, 2, 2, 388, 389, 9,
	4, 2, 2, 389, 59, 3, 2, 2, 2, 390, 391, 9, 5, 2, 2, 391, 61, 3, 2, 2, 2,
	392, 393, 9, 6, 2, 2, 393, 63, 3, 2, 2, 2, 394, 395, 9, 7, 2, 2, 395, 65,
	3, 2, 2, 2, 396, 397, 7, 57, 2, 2, 397, 67, 3, 2, 2, 2, 398, 399, 9, 8,
	2, 2, 399, 69, 3, 2, 2, 2, 400, 401, 9, 9, 2, 2, 401, 71, 3, 2, 2, 2, 36,
	78, 82, 120, 125, 137, 163, 165, 169, 172, 183, 188, 195, 203, 211, 227,
	261, 263, 270, 277, 284, 292, 297, 303, 312, 316, 322, 327, 334, 343, 346,
	350, 360, 364, 376,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'='", "'break'", "'goto'", "'do'", "'end'", "'while'", "'repeat'",
	"'until'", "'if'", "'then'", "'elseif'", "'else'", "'for'", "','", "'in'",
	"'function'", "'local'", "'return'", "'::'", "'.'", "':'", "'nil'", "'false'",
	"'true'", "'...'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'or'", "'and'",
	"'<'", "'>'", "'<='", "'>='", "'~='", "'=='", "'..'", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'//'", "'&'", "'|'", "'~'", "'<<'", "'>>'", "'not'", "'#'",
	"'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING", "INT", "HEX",
	"FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"chunk", "block", "stat", "retstat", "label", "funcname", "varlist", "namelist",
	"explist", "exp", "prefixexp", "functioncall", "varOrExp", "var1", "varSuffix",
	"nameAndArgs", "args", "functiondef", "funcbody", "parlist", "tableconstructor",
	"fieldlist", "field", "fieldsep", "operatorOr", "operatorAnd", "operatorComparison",
	"operatorStrcat", "operatorAddSub", "operatorMulDivMod", "operatorBitwise",
	"operatorUnary", "operatorPower", "number", "string1",
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
	SpartaParserEOF          = antlr.TokenEOF
	SpartaParserT__0         = 1
	SpartaParserT__1         = 2
	SpartaParserT__2         = 3
	SpartaParserT__3         = 4
	SpartaParserT__4         = 5
	SpartaParserT__5         = 6
	SpartaParserT__6         = 7
	SpartaParserT__7         = 8
	SpartaParserT__8         = 9
	SpartaParserT__9         = 10
	SpartaParserT__10        = 11
	SpartaParserT__11        = 12
	SpartaParserT__12        = 13
	SpartaParserT__13        = 14
	SpartaParserT__14        = 15
	SpartaParserT__15        = 16
	SpartaParserT__16        = 17
	SpartaParserT__17        = 18
	SpartaParserT__18        = 19
	SpartaParserT__19        = 20
	SpartaParserT__20        = 21
	SpartaParserT__21        = 22
	SpartaParserT__22        = 23
	SpartaParserT__23        = 24
	SpartaParserT__24        = 25
	SpartaParserT__25        = 26
	SpartaParserT__26        = 27
	SpartaParserT__27        = 28
	SpartaParserT__28        = 29
	SpartaParserT__29        = 30
	SpartaParserT__30        = 31
	SpartaParserT__31        = 32
	SpartaParserT__32        = 33
	SpartaParserT__33        = 34
	SpartaParserT__34        = 35
	SpartaParserT__35        = 36
	SpartaParserT__36        = 37
	SpartaParserT__37        = 38
	SpartaParserT__38        = 39
	SpartaParserT__39        = 40
	SpartaParserT__40        = 41
	SpartaParserT__41        = 42
	SpartaParserT__42        = 43
	SpartaParserT__43        = 44
	SpartaParserT__44        = 45
	SpartaParserT__45        = 46
	SpartaParserT__46        = 47
	SpartaParserT__47        = 48
	SpartaParserT__48        = 49
	SpartaParserT__49        = 50
	SpartaParserT__50        = 51
	SpartaParserT__51        = 52
	SpartaParserT__52        = 53
	SpartaParserT__53        = 54
	SpartaParserT__54        = 55
	SpartaParserNAME         = 56
	SpartaParserNORMALSTRING = 57
	SpartaParserCHARSTRING   = 58
	SpartaParserLONGSTRING   = 59
	SpartaParserINT          = 60
	SpartaParserHEX          = 61
	SpartaParserFLOAT        = 62
	SpartaParserHEX_FLOAT    = 63
	SpartaParserCOMMENT      = 64
	SpartaParserLINE_COMMENT = 65
	SpartaParserWS           = 66
	SpartaParserSHEBANG      = 67
)

// SpartaParser rules.
const (
	SpartaParserRULE_chunk              = 0
	SpartaParserRULE_block              = 1
	SpartaParserRULE_stat               = 2
	SpartaParserRULE_retstat            = 3
	SpartaParserRULE_label              = 4
	SpartaParserRULE_funcname           = 5
	SpartaParserRULE_varlist            = 6
	SpartaParserRULE_namelist           = 7
	SpartaParserRULE_explist            = 8
	SpartaParserRULE_exp                = 9
	SpartaParserRULE_prefixexp          = 10
	SpartaParserRULE_functioncall       = 11
	SpartaParserRULE_varOrExp           = 12
	SpartaParserRULE_var1               = 13
	SpartaParserRULE_varSuffix          = 14
	SpartaParserRULE_nameAndArgs        = 15
	SpartaParserRULE_args               = 16
	SpartaParserRULE_functiondef        = 17
	SpartaParserRULE_funcbody           = 18
	SpartaParserRULE_parlist            = 19
	SpartaParserRULE_tableconstructor   = 20
	SpartaParserRULE_fieldlist          = 21
	SpartaParserRULE_field              = 22
	SpartaParserRULE_fieldsep           = 23
	SpartaParserRULE_operatorOr         = 24
	SpartaParserRULE_operatorAnd        = 25
	SpartaParserRULE_operatorComparison = 26
	SpartaParserRULE_operatorStrcat     = 27
	SpartaParserRULE_operatorAddSub     = 28
	SpartaParserRULE_operatorMulDivMod  = 29
	SpartaParserRULE_operatorBitwise    = 30
	SpartaParserRULE_operatorUnary      = 31
	SpartaParserRULE_operatorPower      = 32
	SpartaParserRULE_number             = 33
	SpartaParserRULE_string1            = 34
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(SpartaParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitChunk(s)
	}
}

func (p *SpartaParser) Chunk() (localctx IChunkContext) {
	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SpartaParserRULE_chunk)

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
		p.Block()
	}
	{
		p.SetState(71)
		p.Match(SpartaParserEOF)
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

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) Retstat() IRetstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SpartaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SpartaParserRULE_block)
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
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__0)|(1<<SpartaParserT__2)|(1<<SpartaParserT__3)|(1<<SpartaParserT__4)|(1<<SpartaParserT__6)|(1<<SpartaParserT__7)|(1<<SpartaParserT__9)|(1<<SpartaParserT__13)|(1<<SpartaParserT__16)|(1<<SpartaParserT__17)|(1<<SpartaParserT__19)|(1<<SpartaParserT__26))) != 0) || _la == SpartaParserNAME {
		{
			p.SetState(73)
			p.Stat()
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__18 {
		{
			p.SetState(79)
			p.Retstat()
		}

	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Varlist() IVarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *StatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *StatContext) Functioncall() IFunctioncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctioncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *StatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *StatContext) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *StatContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *StatContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *StatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *StatContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *StatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *SpartaParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpartaParserRULE_stat)
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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(SpartaParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Varlist()
		}
		{
			p.SetState(84)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(85)
			p.Explist()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Functioncall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)
			p.Match(SpartaParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)
			p.Match(SpartaParserT__3)
		}
		{
			p.SetState(91)
			p.Match(SpartaParserNAME)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(93)
			p.Block()
		}
		{
			p.SetState(94)
			p.Match(SpartaParserT__5)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(96)
			p.Match(SpartaParserT__6)
		}
		{
			p.SetState(97)
			p.exp(0)
		}
		{
			p.SetState(98)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(99)
			p.Block()
		}
		{
			p.SetState(100)
			p.Match(SpartaParserT__5)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(102)
			p.Match(SpartaParserT__7)
		}
		{
			p.SetState(103)
			p.Block()
		}
		{
			p.SetState(104)
			p.Match(SpartaParserT__8)
		}
		{
			p.SetState(105)
			p.exp(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(107)
			p.Match(SpartaParserT__9)
		}
		{
			p.SetState(108)
			p.exp(0)
		}
		{
			p.SetState(109)
			p.Match(SpartaParserT__10)
		}
		{
			p.SetState(110)
			p.Block()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SpartaParserT__11 {
			{
				p.SetState(111)
				p.Match(SpartaParserT__11)
			}
			{
				p.SetState(112)
				p.exp(0)
			}
			{
				p.SetState(113)
				p.Match(SpartaParserT__10)
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

		if _la == SpartaParserT__12 {
			{
				p.SetState(121)
				p.Match(SpartaParserT__12)
			}
			{
				p.SetState(122)
				p.Block()
			}

		}
		{
			p.SetState(125)
			p.Match(SpartaParserT__5)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(127)
			p.Match(SpartaParserT__13)
		}
		{
			p.SetState(128)
			p.Match(SpartaParserNAME)
		}
		{
			p.SetState(129)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(130)
			p.exp(0)
		}
		{
			p.SetState(131)
			p.Match(SpartaParserT__14)
		}
		{
			p.SetState(132)
			p.exp(0)
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SpartaParserT__14 {
			{
				p.SetState(133)
				p.Match(SpartaParserT__14)
			}
			{
				p.SetState(134)
				p.exp(0)
			}

		}
		{
			p.SetState(137)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(138)
			p.Block()
		}
		{
			p.SetState(139)
			p.Match(SpartaParserT__5)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(141)
			p.Match(SpartaParserT__13)
		}
		{
			p.SetState(142)
			p.Namelist()
		}
		{
			p.SetState(143)
			p.Match(SpartaParserT__15)
		}
		{
			p.SetState(144)
			p.Explist()
		}
		{
			p.SetState(145)
			p.Match(SpartaParserT__4)
		}
		{
			p.SetState(146)
			p.Block()
		}
		{
			p.SetState(147)
			p.Match(SpartaParserT__5)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(149)
			p.Match(SpartaParserT__16)
		}
		{
			p.SetState(150)
			p.Funcname()
		}
		{
			p.SetState(151)
			p.Funcbody()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(153)
			p.Match(SpartaParserT__17)
		}
		{
			p.SetState(154)
			p.Match(SpartaParserT__16)
		}
		{
			p.SetState(155)
			p.Match(SpartaParserNAME)
		}
		{
			p.SetState(156)
			p.Funcbody()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(157)
			p.Match(SpartaParserT__17)
		}
		{
			p.SetState(158)
			p.Namelist()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SpartaParserT__1 {
			{
				p.SetState(159)
				p.Match(SpartaParserT__1)
			}
			{
				p.SetState(160)
				p.Explist()
			}

		}

	}

	return localctx
}

// IRetstatContext is an interface to support dynamic dispatch.
type IRetstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetstatContext differentiates from other interfaces.
	IsRetstatContext()
}

type RetstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatContext() *RetstatContext {
	var p = new(RetstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_retstat
	return p
}

func (*RetstatContext) IsRetstatContext() {}

func NewRetstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatContext {
	var p = new(RetstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_retstat

	return p
}

func (s *RetstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *RetstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterRetstat(s)
	}
}

func (s *RetstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitRetstat(s)
	}
}

func (p *SpartaParser) Retstat() (localctx IRetstatContext) {
	localctx = NewRetstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpartaParserRULE_retstat)
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
		p.SetState(165)
		p.Match(SpartaParserT__18)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__16)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25)|(1<<SpartaParserT__26)|(1<<SpartaParserT__30))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(SpartaParserT__42-43))|(1<<(SpartaParserT__49-43))|(1<<(SpartaParserT__52-43))|(1<<(SpartaParserT__53-43))|(1<<(SpartaParserNAME-43))|(1<<(SpartaParserNORMALSTRING-43))|(1<<(SpartaParserCHARSTRING-43))|(1<<(SpartaParserLONGSTRING-43))|(1<<(SpartaParserINT-43))|(1<<(SpartaParserHEX-43))|(1<<(SpartaParserFLOAT-43))|(1<<(SpartaParserHEX_FLOAT-43)))) != 0) {
		{
			p.SetState(166)
			p.Explist()
		}

	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__0 {
		{
			p.SetState(169)
			p.Match(SpartaParserT__0)
		}

	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *SpartaParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpartaParserRULE_label)

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
		p.Match(SpartaParserT__19)
	}
	{
		p.SetState(173)
		p.Match(SpartaParserNAME)
	}
	{
		p.SetState(174)
		p.Match(SpartaParserT__19)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(SpartaParserNAME)
}

func (s *FuncnameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, i)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (p *SpartaParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SpartaParserRULE_funcname)
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
		p.SetState(176)
		p.Match(SpartaParserNAME)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__20 {
		{
			p.SetState(177)
			p.Match(SpartaParserT__20)
		}
		{
			p.SetState(178)
			p.Match(SpartaParserNAME)
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__21 {
		{
			p.SetState(184)
			p.Match(SpartaParserT__21)
		}
		{
			p.SetState(185)
			p.Match(SpartaParserNAME)
		}

	}

	return localctx
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_varlist
	return p
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVar1() []IVar1Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar1Context)(nil)).Elem())
	var tst = make([]IVar1Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar1Context)
		}
	}

	return tst
}

func (s *VarlistContext) Var1(i int) IVar1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar1Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar1Context)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVarlist(s)
	}
}

func (s *VarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVarlist(s)
	}
}

func (p *SpartaParser) Varlist() (localctx IVarlistContext) {
	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SpartaParserRULE_varlist)
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
		p.SetState(188)
		p.Var1()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__14 {
		{
			p.SetState(189)
			p.Match(SpartaParserT__14)
		}
		{
			p.SetState(190)
			p.Var1()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *NamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(SpartaParserNAME)
}

func (s *NamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterNamelist(s)
	}
}

func (s *NamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitNamelist(s)
	}
}

func (p *SpartaParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SpartaParserRULE_namelist)

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
		p.SetState(196)
		p.Match(SpartaParserNAME)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(197)
				p.Match(SpartaParserT__14)
			}
			{
				p.SetState(198)
				p.Match(SpartaParserNAME)
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IExplistContext is an interface to support dynamic dispatch.
type IExplistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExplistContext differentiates from other interfaces.
	IsExplistContext()
}

type ExplistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExplistContext() *ExplistContext {
	var p = new(ExplistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_explist
	return p
}

func (*ExplistContext) IsExplistContext() {}

func NewExplistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExplistContext {
	var p = new(ExplistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_explist

	return p
}

func (s *ExplistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExplistContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExplistContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExplistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExplistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterExplist(s)
	}
}

func (s *ExplistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitExplist(s)
	}
}

func (p *SpartaParser) Explist() (localctx IExplistContext) {
	localctx = NewExplistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SpartaParserRULE_explist)
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
		p.SetState(204)
		p.exp(0)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SpartaParserT__14 {
		{
			p.SetState(205)
			p.Match(SpartaParserT__14)
		}
		{
			p.SetState(206)
			p.exp(0)
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ExpContext) String1() IString1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString1Context)
}

func (s *ExpContext) Functiondef() IFunctiondefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiondefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *ExpContext) Prefixexp() IPrefixexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixexpContext)
}

func (s *ExpContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ExpContext) OperatorUnary() IOperatorUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorUnaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) OperatorPower() IOperatorPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorPowerContext)
}

func (s *ExpContext) OperatorMulDivMod() IOperatorMulDivModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorMulDivModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorMulDivModContext)
}

func (s *ExpContext) OperatorAddSub() IOperatorAddSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAddSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAddSubContext)
}

func (s *ExpContext) OperatorStrcat() IOperatorStrcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorStrcatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorStrcatContext)
}

func (s *ExpContext) OperatorComparison() IOperatorComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorComparisonContext)
}

func (s *ExpContext) OperatorAnd() IOperatorAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAndContext)
}

func (s *ExpContext) OperatorOr() IOperatorOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorOrContext)
}

func (s *ExpContext) OperatorBitwise() IOperatorBitwiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorBitwiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorBitwiseContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *SpartaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *SpartaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SpartaParserRULE_exp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__22:
		{
			p.SetState(213)
			p.Match(SpartaParserT__22)
		}

	case SpartaParserT__23:
		{
			p.SetState(214)
			p.Match(SpartaParserT__23)
		}

	case SpartaParserT__24:
		{
			p.SetState(215)
			p.Match(SpartaParserT__24)
		}

	case SpartaParserINT, SpartaParserHEX, SpartaParserFLOAT, SpartaParserHEX_FLOAT:
		{
			p.SetState(216)
			p.Number()
		}

	case SpartaParserNORMALSTRING, SpartaParserCHARSTRING, SpartaParserLONGSTRING:
		{
			p.SetState(217)
			p.String1()
		}

	case SpartaParserT__25:
		{
			p.SetState(218)
			p.Match(SpartaParserT__25)
		}

	case SpartaParserT__16:
		{
			p.SetState(219)
			p.Functiondef()
		}

	case SpartaParserT__26, SpartaParserNAME:
		{
			p.SetState(220)
			p.Prefixexp()
		}

	case SpartaParserT__30:
		{
			p.SetState(221)
			p.Tableconstructor()
		}

	case SpartaParserT__42, SpartaParserT__49, SpartaParserT__52, SpartaParserT__53:
		{
			p.SetState(222)
			p.OperatorUnary()
		}
		{
			p.SetState(223)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(228)
					p.OperatorPower()
				}
				{
					p.SetState(229)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(231)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(232)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(233)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(235)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(236)
					p.OperatorAddSub()
				}
				{
					p.SetState(237)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(240)
					p.OperatorStrcat()
				}
				{
					p.SetState(241)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(244)
					p.OperatorComparison()
				}
				{
					p.SetState(245)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(248)
					p.OperatorAnd()
				}
				{
					p.SetState(249)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(252)
					p.OperatorOr()
				}
				{
					p.SetState(253)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SpartaParserRULE_exp)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(256)
					p.OperatorBitwise()
				}
				{
					p.SetState(257)
					p.exp(2)
				}

			}

		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IPrefixexpContext is an interface to support dynamic dispatch.
type IPrefixexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixexpContext differentiates from other interfaces.
	IsPrefixexpContext()
}

type PrefixexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixexpContext() *PrefixexpContext {
	var p = new(PrefixexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_prefixexp
	return p
}

func (*PrefixexpContext) IsPrefixexpContext() {}

func NewPrefixexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixexpContext {
	var p = new(PrefixexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_prefixexp

	return p
}

func (s *PrefixexpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixexpContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *PrefixexpContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *PrefixexpContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *PrefixexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterPrefixexp(s)
	}
}

func (s *PrefixexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitPrefixexp(s)
	}
}

func (p *SpartaParser) Prefixexp() (localctx IPrefixexpContext) {
	localctx = NewPrefixexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SpartaParserRULE_prefixexp)

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
		p.SetState(264)
		p.VarOrExp()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(265)
				p.NameAndArgs()
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_functioncall
	return p
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *FunctioncallContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *FunctioncallContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (p *SpartaParser) Functioncall() (localctx IFunctioncallContext) {
	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SpartaParserRULE_functioncall)

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
		p.SetState(271)
		p.VarOrExp()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(272)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IVarOrExpContext is an interface to support dynamic dispatch.
type IVarOrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarOrExpContext differentiates from other interfaces.
	IsVarOrExpContext()
}

type VarOrExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrExpContext() *VarOrExpContext {
	var p = new(VarOrExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_varOrExp
	return p
}

func (*VarOrExpContext) IsVarOrExpContext() {}

func NewVarOrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrExpContext {
	var p = new(VarOrExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_varOrExp

	return p
}

func (s *VarOrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrExpContext) Var1() IVar1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar1Context)
}

func (s *VarOrExpContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarOrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVarOrExp(s)
	}
}

func (s *VarOrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVarOrExp(s)
	}
}

func (p *SpartaParser) VarOrExp() (localctx IVarOrExpContext) {
	localctx = NewVarOrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SpartaParserRULE_varOrExp)

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

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.Var1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(SpartaParserT__26)
		}
		{
			p.SetState(279)
			p.exp(0)
		}
		{
			p.SetState(280)
			p.Match(SpartaParserT__27)
		}

	}

	return localctx
}

// IVar1Context is an interface to support dynamic dispatch.
type IVar1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar1Context differentiates from other interfaces.
	IsVar1Context()
}

type Var1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar1Context() *Var1Context {
	var p = new(Var1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_var1
	return p
}

func (*Var1Context) IsVar1Context() {}

func NewVar1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var1Context {
	var p = new(Var1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_var1

	return p
}

func (s *Var1Context) GetParser() antlr.Parser { return s.parser }

func (s *Var1Context) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *Var1Context) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *Var1Context) AllVarSuffix() []IVarSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem())
	var tst = make([]IVarSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarSuffixContext)
		}
	}

	return tst
}

func (s *Var1Context) VarSuffix(i int) IVarSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarSuffixContext)
}

func (s *Var1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVar1(s)
	}
}

func (s *Var1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVar1(s)
	}
}

func (p *SpartaParser) Var1() (localctx IVar1Context) {
	localctx = NewVar1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SpartaParserRULE_var1)

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
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserNAME:
		{
			p.SetState(284)
			p.Match(SpartaParserNAME)
		}

	case SpartaParserT__26:
		{
			p.SetState(285)
			p.Match(SpartaParserT__26)
		}
		{
			p.SetState(286)
			p.exp(0)
		}
		{
			p.SetState(287)
			p.Match(SpartaParserT__27)
		}
		{
			p.SetState(288)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(292)
				p.VarSuffix()
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IVarSuffixContext is an interface to support dynamic dispatch.
type IVarSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarSuffixContext differentiates from other interfaces.
	IsVarSuffixContext()
}

type VarSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSuffixContext() *VarSuffixContext {
	var p = new(VarSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_varSuffix
	return p
}

func (*VarSuffixContext) IsVarSuffixContext() {}

func NewVarSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSuffixContext {
	var p = new(VarSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_varSuffix

	return p
}

func (s *VarSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSuffixContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarSuffixContext) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *VarSuffixContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *VarSuffixContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *VarSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVarSuffix(s)
	}
}

func (s *VarSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVarSuffix(s)
	}
}

func (p *SpartaParser) VarSuffix() (localctx IVarSuffixContext) {
	localctx = NewVarSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SpartaParserRULE_varSuffix)
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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__21)|(1<<SpartaParserT__26)|(1<<SpartaParserT__30))) != 0) || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SpartaParserNORMALSTRING-57))|(1<<(SpartaParserCHARSTRING-57))|(1<<(SpartaParserLONGSTRING-57)))) != 0) {
		{
			p.SetState(298)
			p.NameAndArgs()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__28:
		{
			p.SetState(304)
			p.Match(SpartaParserT__28)
		}
		{
			p.SetState(305)
			p.exp(0)
		}
		{
			p.SetState(306)
			p.Match(SpartaParserT__29)
		}

	case SpartaParserT__20:
		{
			p.SetState(308)
			p.Match(SpartaParserT__20)
		}
		{
			p.SetState(309)
			p.Match(SpartaParserNAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameAndArgsContext is an interface to support dynamic dispatch.
type INameAndArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameAndArgsContext differentiates from other interfaces.
	IsNameAndArgsContext()
}

type NameAndArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAndArgsContext() *NameAndArgsContext {
	var p = new(NameAndArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_nameAndArgs
	return p
}

func (*NameAndArgsContext) IsNameAndArgsContext() {}

func NewNameAndArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAndArgsContext {
	var p = new(NameAndArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_nameAndArgs

	return p
}

func (s *NameAndArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAndArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NameAndArgsContext) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *NameAndArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAndArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAndArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitNameAndArgs(s)
	}
}

func (p *SpartaParser) NameAndArgs() (localctx INameAndArgsContext) {
	localctx = NewNameAndArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SpartaParserRULE_nameAndArgs)
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
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__21 {
		{
			p.SetState(312)
			p.Match(SpartaParserT__21)
		}
		{
			p.SetState(313)
			p.Match(SpartaParserNAME)
		}

	}
	{
		p.SetState(316)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ArgsContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ArgsContext) String1() IString1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString1Context)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *SpartaParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SpartaParserRULE_args)
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

	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(SpartaParserT__26)
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__16)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25)|(1<<SpartaParserT__26)|(1<<SpartaParserT__30))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(SpartaParserT__42-43))|(1<<(SpartaParserT__49-43))|(1<<(SpartaParserT__52-43))|(1<<(SpartaParserT__53-43))|(1<<(SpartaParserNAME-43))|(1<<(SpartaParserNORMALSTRING-43))|(1<<(SpartaParserCHARSTRING-43))|(1<<(SpartaParserLONGSTRING-43))|(1<<(SpartaParserINT-43))|(1<<(SpartaParserHEX-43))|(1<<(SpartaParserFLOAT-43))|(1<<(SpartaParserHEX_FLOAT-43)))) != 0) {
			{
				p.SetState(319)
				p.Explist()
			}

		}
		{
			p.SetState(322)
			p.Match(SpartaParserT__27)
		}

	case SpartaParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Tableconstructor()
		}

	case SpartaParserNORMALSTRING, SpartaParserCHARSTRING, SpartaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.String1()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_functiondef
	return p
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFunctiondef(s)
	}
}

func (p *SpartaParser) Functiondef() (localctx IFunctiondefContext) {
	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SpartaParserRULE_functiondef)

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
		p.SetState(327)
		p.Match(SpartaParserT__16)
	}
	{
		p.SetState(328)
		p.Funcbody()
	}

	return localctx
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_funcbody
	return p
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncbodyContext) Parlist() IParlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (p *SpartaParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SpartaParserRULE_funcbody)
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
		p.SetState(330)
		p.Match(SpartaParserT__26)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__25 || _la == SpartaParserNAME {
		{
			p.SetState(331)
			p.Parlist()
		}

	}
	{
		p.SetState(334)
		p.Match(SpartaParserT__27)
	}
	{
		p.SetState(335)
		p.Block()
	}
	{
		p.SetState(336)
		p.Match(SpartaParserT__5)
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

func (s *ParlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterParlist(s)
	}
}

func (s *ParlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitParlist(s)
	}
}

func (p *SpartaParser) Parlist() (localctx IParlistContext) {
	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SpartaParserRULE_parlist)
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

	p.SetState(344)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SpartaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Namelist()
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SpartaParserT__14 {
			{
				p.SetState(339)
				p.Match(SpartaParserT__14)
			}
			{
				p.SetState(340)
				p.Match(SpartaParserT__25)
			}

		}

	case SpartaParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(SpartaParserT__25)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableconstructorContext is an interface to support dynamic dispatch.
type ITableconstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableconstructorContext differentiates from other interfaces.
	IsTableconstructorContext()
}

type TableconstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableconstructorContext() *TableconstructorContext {
	var p = new(TableconstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_tableconstructor
	return p
}

func (*TableconstructorContext) IsTableconstructorContext() {}

func NewTableconstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableconstructorContext {
	var p = new(TableconstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_tableconstructor

	return p
}

func (s *TableconstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableconstructorContext) Fieldlist() IFieldlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldlistContext)
}

func (s *TableconstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableconstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableconstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterTableconstructor(s)
	}
}

func (s *TableconstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitTableconstructor(s)
	}
}

func (p *SpartaParser) Tableconstructor() (localctx ITableconstructorContext) {
	localctx = NewTableconstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SpartaParserRULE_tableconstructor)
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
		p.SetState(346)
		p.Match(SpartaParserT__30)
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SpartaParserT__16)|(1<<SpartaParserT__22)|(1<<SpartaParserT__23)|(1<<SpartaParserT__24)|(1<<SpartaParserT__25)|(1<<SpartaParserT__26)|(1<<SpartaParserT__28)|(1<<SpartaParserT__30))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(SpartaParserT__42-43))|(1<<(SpartaParserT__49-43))|(1<<(SpartaParserT__52-43))|(1<<(SpartaParserT__53-43))|(1<<(SpartaParserNAME-43))|(1<<(SpartaParserNORMALSTRING-43))|(1<<(SpartaParserCHARSTRING-43))|(1<<(SpartaParserLONGSTRING-43))|(1<<(SpartaParserINT-43))|(1<<(SpartaParserHEX-43))|(1<<(SpartaParserFLOAT-43))|(1<<(SpartaParserHEX_FLOAT-43)))) != 0) {
		{
			p.SetState(347)
			p.Fieldlist()
		}

	}
	{
		p.SetState(350)
		p.Match(SpartaParserT__31)
	}

	return localctx
}

// IFieldlistContext is an interface to support dynamic dispatch.
type IFieldlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldlistContext differentiates from other interfaces.
	IsFieldlistContext()
}

type FieldlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldlistContext() *FieldlistContext {
	var p = new(FieldlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fieldlist
	return p
}

func (*FieldlistContext) IsFieldlistContext() {}

func NewFieldlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldlistContext {
	var p = new(FieldlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fieldlist

	return p
}

func (s *FieldlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldlistContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldlistContext) AllFieldsep() []IFieldsepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldsepContext)(nil)).Elem())
	var tst = make([]IFieldsepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldsepContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Fieldsep(i int) IFieldsepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldsepContext)
}

func (s *FieldlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFieldlist(s)
	}
}

func (s *FieldlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFieldlist(s)
	}
}

func (p *SpartaParser) Fieldlist() (localctx IFieldlistContext) {
	localctx = NewFieldlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SpartaParserRULE_fieldlist)
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
		p.SetState(352)
		p.Field()
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(353)
				p.Fieldsep()
			}
			{
				p.SetState(354)
				p.Field()
			}

		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__0 || _la == SpartaParserT__14 {
		{
			p.SetState(361)
			p.Fieldsep()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *FieldContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(SpartaParserNAME, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *SpartaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SpartaParserRULE_field)

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

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Match(SpartaParserT__28)
		}
		{
			p.SetState(365)
			p.exp(0)
		}
		{
			p.SetState(366)
			p.Match(SpartaParserT__29)
		}
		{
			p.SetState(367)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(368)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Match(SpartaParserNAME)
		}
		{
			p.SetState(371)
			p.Match(SpartaParserT__1)
		}
		{
			p.SetState(372)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.exp(0)
		}

	}

	return localctx
}

// IFieldsepContext is an interface to support dynamic dispatch.
type IFieldsepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsepContext differentiates from other interfaces.
	IsFieldsepContext()
}

type FieldsepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsepContext() *FieldsepContext {
	var p = new(FieldsepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_fieldsep
	return p
}

func (*FieldsepContext) IsFieldsepContext() {}

func NewFieldsepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsepContext {
	var p = new(FieldsepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_fieldsep

	return p
}

func (s *FieldsepContext) GetParser() antlr.Parser { return s.parser }
func (s *FieldsepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterFieldsep(s)
	}
}

func (s *FieldsepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitFieldsep(s)
	}
}

func (p *SpartaParser) Fieldsep() (localctx IFieldsepContext) {
	localctx = NewFieldsepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SpartaParserRULE_fieldsep)
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
		p.SetState(376)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SpartaParserT__0 || _la == SpartaParserT__14) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorOrContext is an interface to support dynamic dispatch.
type IOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorOrContext differentiates from other interfaces.
	IsOperatorOrContext()
}

type OperatorOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorOrContext() *OperatorOrContext {
	var p = new(OperatorOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorOr
	return p
}

func (*OperatorOrContext) IsOperatorOrContext() {}

func NewOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorOrContext {
	var p = new(OperatorOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorOr

	return p
}

func (s *OperatorOrContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorOr(s)
	}
}

func (s *OperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorOr(s)
	}
}

func (p *SpartaParser) OperatorOr() (localctx IOperatorOrContext) {
	localctx = NewOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SpartaParserRULE_operatorOr)

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
		p.SetState(378)
		p.Match(SpartaParserT__32)
	}

	return localctx
}

// IOperatorAndContext is an interface to support dynamic dispatch.
type IOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAndContext differentiates from other interfaces.
	IsOperatorAndContext()
}

type OperatorAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAndContext() *OperatorAndContext {
	var p = new(OperatorAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorAnd
	return p
}

func (*OperatorAndContext) IsOperatorAndContext() {}

func NewOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAndContext {
	var p = new(OperatorAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorAnd

	return p
}

func (s *OperatorAndContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorAnd(s)
	}
}

func (s *OperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorAnd(s)
	}
}

func (p *SpartaParser) OperatorAnd() (localctx IOperatorAndContext) {
	localctx = NewOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SpartaParserRULE_operatorAnd)

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
		p.SetState(380)
		p.Match(SpartaParserT__33)
	}

	return localctx
}

// IOperatorComparisonContext is an interface to support dynamic dispatch.
type IOperatorComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorComparisonContext differentiates from other interfaces.
	IsOperatorComparisonContext()
}

type OperatorComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorComparisonContext() *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorComparison
	return p
}

func (*OperatorComparisonContext) IsOperatorComparisonContext() {}

func NewOperatorComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorComparison

	return p
}

func (s *OperatorComparisonContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorComparison(s)
	}
}

func (s *OperatorComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorComparison(s)
	}
}

func (p *SpartaParser) OperatorComparison() (localctx IOperatorComparisonContext) {
	localctx = NewOperatorComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SpartaParserRULE_operatorComparison)
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
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(SpartaParserT__34-35))|(1<<(SpartaParserT__35-35))|(1<<(SpartaParserT__36-35))|(1<<(SpartaParserT__37-35))|(1<<(SpartaParserT__38-35))|(1<<(SpartaParserT__39-35)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorStrcatContext is an interface to support dynamic dispatch.
type IOperatorStrcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorStrcatContext differentiates from other interfaces.
	IsOperatorStrcatContext()
}

type OperatorStrcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorStrcatContext() *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorStrcat
	return p
}

func (*OperatorStrcatContext) IsOperatorStrcatContext() {}

func NewOperatorStrcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorStrcat

	return p
}

func (s *OperatorStrcatContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStrcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorStrcat(s)
	}
}

func (p *SpartaParser) OperatorStrcat() (localctx IOperatorStrcatContext) {
	localctx = NewOperatorStrcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SpartaParserRULE_operatorStrcat)

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
		p.SetState(384)
		p.Match(SpartaParserT__40)
	}

	return localctx
}

// IOperatorAddSubContext is an interface to support dynamic dispatch.
type IOperatorAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAddSubContext differentiates from other interfaces.
	IsOperatorAddSubContext()
}

type OperatorAddSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAddSubContext() *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorAddSub
	return p
}

func (*OperatorAddSubContext) IsOperatorAddSubContext() {}

func NewOperatorAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorAddSub

	return p
}

func (s *OperatorAddSubContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorAddSub(s)
	}
}

func (p *SpartaParser) OperatorAddSub() (localctx IOperatorAddSubContext) {
	localctx = NewOperatorAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SpartaParserRULE_operatorAddSub)
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
		p.SetState(386)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SpartaParserT__41 || _la == SpartaParserT__42) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorMulDivModContext is an interface to support dynamic dispatch.
type IOperatorMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorMulDivModContext differentiates from other interfaces.
	IsOperatorMulDivModContext()
}

type OperatorMulDivModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorMulDivModContext() *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorMulDivMod
	return p
}

func (*OperatorMulDivModContext) IsOperatorMulDivModContext() {}

func NewOperatorMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorMulDivMod

	return p
}

func (s *OperatorMulDivModContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorMulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorMulDivMod(s)
	}
}

func (p *SpartaParser) OperatorMulDivMod() (localctx IOperatorMulDivModContext) {
	localctx = NewOperatorMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SpartaParserRULE_operatorMulDivMod)
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
		p.SetState(388)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(SpartaParserT__43-44))|(1<<(SpartaParserT__44-44))|(1<<(SpartaParserT__45-44))|(1<<(SpartaParserT__46-44)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorBitwiseContext is an interface to support dynamic dispatch.
type IOperatorBitwiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorBitwiseContext differentiates from other interfaces.
	IsOperatorBitwiseContext()
}

type OperatorBitwiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorBitwiseContext() *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorBitwise
	return p
}

func (*OperatorBitwiseContext) IsOperatorBitwiseContext() {}

func NewOperatorBitwiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorBitwise

	return p
}

func (s *OperatorBitwiseContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorBitwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorBitwiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorBitwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorBitwise(s)
	}
}

func (s *OperatorBitwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorBitwise(s)
	}
}

func (p *SpartaParser) OperatorBitwise() (localctx IOperatorBitwiseContext) {
	localctx = NewOperatorBitwiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SpartaParserRULE_operatorBitwise)
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
		p.SetState(390)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(SpartaParserT__47-48))|(1<<(SpartaParserT__48-48))|(1<<(SpartaParserT__49-48))|(1<<(SpartaParserT__50-48))|(1<<(SpartaParserT__51-48)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorUnaryContext is an interface to support dynamic dispatch.
type IOperatorUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorUnaryContext differentiates from other interfaces.
	IsOperatorUnaryContext()
}

type OperatorUnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorUnaryContext() *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorUnary
	return p
}

func (*OperatorUnaryContext) IsOperatorUnaryContext() {}

func NewOperatorUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorUnary

	return p
}

func (s *OperatorUnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorUnary(s)
	}
}

func (p *SpartaParser) OperatorUnary() (localctx IOperatorUnaryContext) {
	localctx = NewOperatorUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SpartaParserRULE_operatorUnary)
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
		p.SetState(392)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(SpartaParserT__42-43))|(1<<(SpartaParserT__49-43))|(1<<(SpartaParserT__52-43))|(1<<(SpartaParserT__53-43)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorPowerContext is an interface to support dynamic dispatch.
type IOperatorPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorPowerContext differentiates from other interfaces.
	IsOperatorPowerContext()
}

type OperatorPowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPowerContext() *OperatorPowerContext {
	var p = new(OperatorPowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_operatorPower
	return p
}

func (*OperatorPowerContext) IsOperatorPowerContext() {}

func NewOperatorPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPowerContext {
	var p = new(OperatorPowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_operatorPower

	return p
}

func (s *OperatorPowerContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterOperatorPower(s)
	}
}

func (s *OperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitOperatorPower(s)
	}
}

func (p *SpartaParser) OperatorPower() (localctx IOperatorPowerContext) {
	localctx = NewOperatorPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SpartaParserRULE_operatorPower)

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
		p.SetState(394)
		p.Match(SpartaParserT__54)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(SpartaParserINT, 0)
}

func (s *NumberContext) HEX() antlr.TerminalNode {
	return s.GetToken(SpartaParserHEX, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SpartaParserFLOAT, 0)
}

func (s *NumberContext) HEX_FLOAT() antlr.TerminalNode {
	return s.GetToken(SpartaParserHEX_FLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SpartaParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SpartaParserRULE_number)
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
		p.SetState(396)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(SpartaParserINT-60))|(1<<(SpartaParserHEX-60))|(1<<(SpartaParserFLOAT-60))|(1<<(SpartaParserHEX_FLOAT-60)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IString1Context is an interface to support dynamic dispatch.
type IString1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString1Context differentiates from other interfaces.
	IsString1Context()
}

type String1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString1Context() *String1Context {
	var p = new(String1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_string1
	return p
}

func (*String1Context) IsString1Context() {}

func NewString1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String1Context {
	var p = new(String1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_string1

	return p
}

func (s *String1Context) GetParser() antlr.Parser { return s.parser }

func (s *String1Context) NORMALSTRING() antlr.TerminalNode {
	return s.GetToken(SpartaParserNORMALSTRING, 0)
}

func (s *String1Context) CHARSTRING() antlr.TerminalNode {
	return s.GetToken(SpartaParserCHARSTRING, 0)
}

func (s *String1Context) LONGSTRING() antlr.TerminalNode {
	return s.GetToken(SpartaParserLONGSTRING, 0)
}

func (s *String1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterString1(s)
	}
}

func (s *String1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitString1(s)
	}
}

func (p *SpartaParser) String1() (localctx IString1Context) {
	localctx = NewString1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SpartaParserRULE_string1)
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
		p.SetState(398)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SpartaParserNORMALSTRING-57))|(1<<(SpartaParserCHARSTRING-57))|(1<<(SpartaParserLONGSTRING-57)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SpartaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SpartaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
