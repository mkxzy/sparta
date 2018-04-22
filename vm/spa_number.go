package vm

import (
	"strconv"
	"fmt"
)

type SPANumber float64

func NewNumber(literal string) SPANumber {
	r, err := strconv.ParseFloat(literal, 64)
	if err != nil {
		panic("数字转化错误")
	}
	return SPANumber(r)
}

func(n SPANumber) String() string {
	return fmt.Sprintf("%.2f", n)
}

func(n SPANumber) IsTrue() bool {
	return n > 0.0
}