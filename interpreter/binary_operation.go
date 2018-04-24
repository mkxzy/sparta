package interpreter

import (
	"github.com/mkxzy/sparta/vm"
	"fmt"
	"bytes"
)

//操作符和操作关联
var operations = make(map[string]func(left, right vm.SPAValue) (result vm.SPAValue, ok bool))

//初始化
func init()  {
	operations["+"] = add
	operations["-"] = sub
	operations["*"] = mul
	operations["/"] = div
	operations["%"] = mod
}

/**
+号运算
 */
func add(left, right vm.SPAValue) (result vm.SPAValue, ok bool) {
	switch left.(type) {
	case vm.SPANumber:
		realLeft := left.(vm.SPANumber)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = realLeft + realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft + vm.SPANumber(realRight)
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft + vm.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case vm.SPAInteger:
		realLeft := left.(vm.SPAInteger)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft) + realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft + realRight
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft + realRight.ToInt()
			ok = true
			return
		}
	case vm.SPAString:
		realLeft := left.(vm.SPAString)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPAString(fmt.Sprintf("%s%f", realLeft, realRight))
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = vm.SPAString(fmt.Sprintf("%s%d", realLeft, realRight))
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}

/**
减号操作
 */
func sub(left, right vm.SPAValue) (result vm.SPAValue, ok bool) {
	switch left.(type) {
	case vm.SPANumber:
		realLeft := left.(vm.SPANumber)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = realLeft - realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft - vm.SPANumber(realRight)
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft - vm.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case vm.SPAInteger:
		realLeft := left.(vm.SPAInteger)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft) - realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft - realRight
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft - realRight.ToInt()
			ok = true
			return
		}
	case vm.SPABool:
		realLeft := left.(vm.SPABool)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft.ToInt()) - realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft.ToInt() - realRight
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft.ToInt() - realRight.ToInt()
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}

/**
乘法操作
 */
func mul(left, right vm.SPAValue) (result vm.SPAValue, ok bool)  {
	switch left.(type) {
	case vm.SPANumber:
		realLeft := left.(vm.SPANumber)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = realLeft * realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft * vm.SPANumber(realRight)
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft * vm.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case vm.SPAInteger:
		realLeft := left.(vm.SPAInteger)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft) * realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft * realRight
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft * realRight.ToInt()
			ok = true
			return
		}
	case vm.SPABool:
		realLeft := left.(vm.SPABool)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft.ToInt()) * realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft.ToInt() * realRight
			ok = true
			return
		case vm.SPABool:
			realRight := right.(vm.SPABool)
			result = realLeft.ToInt() * realRight.ToInt()
			ok = true
			return
		}
	case vm.SPAString:
		realLeft := left.(vm.SPAString)
		switch right.(type) {
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			var buffer bytes.Buffer
			for i := 0; i < int(realRight); i++ {
				buffer.WriteString(string(realLeft))
			}
			result = vm.SPAString(buffer.String())
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}

/**
除法操作
 */
func div(left, right vm.SPAValue) (result vm.SPAValue, ok bool)  {
	switch left.(type) {
	case vm.SPANumber:
		realLeft := left.(vm.SPANumber)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = realLeft + realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft + vm.SPANumber(realRight)
			ok = true
			return
		}
	case vm.SPAInteger:
		realLeft := left.(vm.SPAInteger)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = vm.SPANumber(realLeft) + realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft + realRight
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}

/**
取模操作
 */
func mod(left, right vm.SPAValue) (result vm.SPAValue, ok bool)  {
	switch left.(type) {
	case vm.SPAInteger:
		realLeft := left.(vm.SPAInteger)
		switch right.(type) {
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft % realRight
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}
