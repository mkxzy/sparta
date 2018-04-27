package interpreter

import "github.com/mkxzy/sparta/types"

const MaxOperandStack = 10

var operands [MaxOperandStack]types.SPAValue
var sp = -1

func PushValue(v types.SPAValue)  {
	if sp >= len(operands) - 1{
		panic("操作树栈溢出")
	}
	sp++
	operands[sp] = v

	log.Infof("操作数压栈： %v", operands)
}

func PushNullValue()  {
	PushValue(types.Null())
}

func PopValue() types.SPAValue  {
	if sp > -1 {
		v := operands[sp]
		operands[sp] = nil
		sp--

		log.Infof("操作数出栈： %v", operands)
		return v
	}
	panic("操作树栈为空")
}

func PeekValue() types.SPAValue {
	if sp > -1 {
		return operands[sp]
	}
	return nil
}
