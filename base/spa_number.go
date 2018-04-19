package base

import (
	"strconv"
	"fmt"
)

type SPANumber float64

func NewNumber(literal string) SPANumber {
	var spa = SPANumber(parseNumber(literal))
	return spa
}

func(n SPANumber) String() string {
	return fmt.Sprintf("%f", n)
}

/**
解析数字
 */
func parseNumber(text string) float64 {
	r, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic("数字转化错误")
	}
	return r
}

func(n SPANumber) IsTrue() bool {
	return n > 0.0
}