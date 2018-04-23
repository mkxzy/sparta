package vm

import "strconv"

type SPAInteger int64

func NewInteger(literal string) SPAInteger {
	r, err := strconv.Atoi(literal)
	if err != nil {
		panic("数字转化错误")
	}
	return SPAInteger(r)
}

func(n SPAInteger) IsTrue() bool {
	return n > 0
}
