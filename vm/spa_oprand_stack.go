package vm

import "github.com/op/go-logging"

var log = logging.MustGetLogger("operands")

var operands [100]SPAValue
var sp = -1

func PushValue(v SPAValue)  {
	if sp >= len(operands) - 1{
		panic("操作树栈溢出")
	}
	sp++
	operands[sp] = v

	log.Infof("操作数压栈： %v", operands)
}

func PopValue() SPAValue  {
	if sp > -1 {
		v := operands[sp]
		operands[sp] = nil
		sp--

		log.Infof("操作数出栈： %v", operands)
		return v
	}
	panic("操作树栈为空")
}

func PeekValue() SPAValue {
	if sp > -1 {
		return operands[sp]
	}
	return nil
}
