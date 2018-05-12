package operation

import (
	"github.com/mkxzy/sparta/types"
	"fmt"
	"bytes"
)

//操作符和操作关联
var binOperations = make(map[string]func(left, right types.SPAValue) (result types.SPAValue, ok bool))

//初始化
func init()  {
	binOperations["+"] = add
	binOperations["-"] = sub
	binOperations["*"] = mul
	binOperations["/"] = div
	binOperations["%"] = mod
	binOperations["=="] = equals
}

/**
算数运算
 */
func Calculate(op string){
	second := PopValue()
	first := PopValue()
	operation := binOperations[op]
	if operation == nil{
		panic("不支持的操作")
	}
	result, _ := operation(first, second)
	PushValue(result)
}

/**
+号运算
 */
func add(left, right types.SPAValue) (result types.SPAValue, ok bool) {
	switch left.(type) {
	case types.SPANumber:
		realLeft := left.(types.SPANumber)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = realLeft + realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft + types.SPANumber(realRight)
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft + types.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft) + realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft + realRight
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft + realRight.ToInt()
			ok = true
			return
		}
	case types.SPAString:
		realLeft := left.(types.SPAString)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPAString(fmt.Sprintf("%s%f", realLeft, realRight))
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = types.SPAString(fmt.Sprintf("%s%d", realLeft, realRight))
			ok = true
			return
		}
	}
	result = types.Null()
	ok = false
	return
}

/**
减号操作
 */
func sub(left, right types.SPAValue) (result types.SPAValue, ok bool) {
	switch left.(type) {
	case types.SPANumber:
		realLeft := left.(types.SPANumber)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = realLeft - realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft - types.SPANumber(realRight)
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft - types.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft) - realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft - realRight
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft - realRight.ToInt()
			ok = true
			return
		}
	case types.SPABool:
		realLeft := left.(types.SPABool)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft.ToInt()) - realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft.ToInt() - realRight
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft.ToInt() - realRight.ToInt()
			ok = true
			return
		}
	}
	result = types.Null()
	ok = false
	return
}

/**
乘法操作
 */
func mul(left, right types.SPAValue) (result types.SPAValue, ok bool)  {
	switch left.(type) {
	case types.SPANumber:
		realLeft := left.(types.SPANumber)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = realLeft * realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft * types.SPANumber(realRight)
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft * types.SPANumber(realRight.ToInt())
			ok = true
			return
		}
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft) * realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft * realRight
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft * realRight.ToInt()
			ok = true
			return
		}
	case types.SPABool:
		realLeft := left.(types.SPABool)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft.ToInt()) * realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft.ToInt() * realRight
			ok = true
			return
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = realLeft.ToInt() * realRight.ToInt()
			ok = true
			return
		}
	case types.SPAString:
		realLeft := left.(types.SPAString)
		switch right.(type) {
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			var buffer bytes.Buffer
			for i := 0; i < int(realRight); i++ {
				buffer.WriteString(string(realLeft))
			}
			result = types.SPAString(buffer.String())
			ok = true
			return
		}
	}
	result = types.Null()
	ok = false
	return
}

/**
除法操作
 */
func div(left, right types.SPAValue) (result types.SPAValue, ok bool)  {
	switch left.(type) {
	case types.SPANumber:
		realLeft := left.(types.SPANumber)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = realLeft + realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft + types.SPANumber(realRight)
			ok = true
			return
		}
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPANumber:
			realRight := right.(types.SPANumber)
			result = types.SPANumber(realLeft) + realRight
			ok = true
			return
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft + realRight
			ok = true
			return
		}
	}
	result = types.Null()
	ok = false
	return
}

/**
取模操作
 */
func mod(left, right types.SPAValue) (result types.SPAValue, ok bool)  {
	switch left.(type) {
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = realLeft % realRight
			ok = true
			return
		}
	}
	result = types.Null()
	ok = false
	return
}

/**
比较相等
 */
func equals(left, right types.SPAValue) (result types.SPAValue, ok bool) {
	result = types.Null()
	ok = false
	switch left.(type) {
	case types.SPABool:
		realLeft := left.(types.SPABool)
		switch right.(type) {
		case types.SPABool:
			realRight := right.(types.SPABool)
			result = types.SPABool(realLeft == realRight)
			ok = true
		}
	case types.SPAInteger:
		realLeft := left.(types.SPAInteger)
		switch right.(type) {
		case types.SPAInteger:
			realRight := right.(types.SPAInteger)
			result = types.SPABool(realLeft == realRight)
			ok = true
		case types.SPANumber:
			x := realLeft.ToNumber()
			realRight := right.(types.SPANumber)
			result = types.SPABool(x == realRight)
			ok = true
		}
	}
	return
}
