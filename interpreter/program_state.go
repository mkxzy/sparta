package interpreter

import (
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

/**
程序状态
 */
type ProgramState struct {
	globals     *symbol.MemorySpace //全局变量空间
	currentFunc *function.FunState  //函数调用链
}

func NewProgramState() *ProgramState {
	globals := symbol.NewMemorySpace("global")                        //全局内存空间
	printFunc := function.NewInternalFunction("print", []string{"s"}) //print内置函数
	globals.Define(function.NewFunVariable(printFunc))
	state := & ProgramState{

		globals:     globals,
		currentFunc: nil,
	}
	return state
}

func (ps *ProgramState) Define(s symbol.Symbol)  {
	if ps.currentFunc != nil {
		ps.currentFunc.Define(s)
	}else{
		ps.globals.Define(s)
	}
}

func (ps *ProgramState) Resolve(name string) (s symbol.Symbol) {
	s = ps.currentFunc.Resolve(name)
	if s == nil {
		s = ps.globals.Resolve(name)
	}
	return
}
