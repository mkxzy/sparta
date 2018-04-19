package scope

import (
	"github.com/mkxzy/sparta/base"
)

type FunFrame struct {
	Fun *SPAFunction
	Locals []*VariableSymbol //局部变量
	ReturnValue base.SPAValue
}

func NewFrame(f *SPAFunction)  *FunFrame{
	return &FunFrame{
		Fun: f,
		Locals: make([]*VariableSymbol, 0 , 10),
		ReturnValue: base.Null(),
	}
}

func(ff *FunFrame) PushArgs(args []base.SPAValue)  {
	for i := 0; i < len(ff.Fun.Args); i++{
		ff.Locals = append(ff.Locals, NewVariable(ff.Fun.Args[i], args[i]))
	}
}