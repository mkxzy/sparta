package interpreter

import (
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

/**
程序状态
 */
type ProgramState struct {
	Globals     *symbol.MemorySpace //全局变量空间
	CurrentFunc *function.FunState  //函数调用链
}

func NewProgramState() *ProgramState {
	globals := symbol.NewMemorySpace("global")                        //全局内存空间
	printFunc := function.NewInternalFunction("print", []string{"s"}) //print内置函数
	globals.Define(function.NewFunVariable(printFunc))
	state := &ProgramState{
		Globals:     globals,
		CurrentFunc: nil,
	}
	return state
}

//定义变量
func (ps *ProgramState) Define(s symbol.Symbol)  {
	if ps.CurrentFunc != nil {
		ps.CurrentFunc.Define(s)
	}else{
		ps.Globals.Define(s)
	}
}

//解析变量
func (ps *ProgramState) Resolve(name string) (s symbol.Symbol) {
	s = ps.CurrentFunc.Resolve(name)
	if s == nil {
		s = ps.Globals.Resolve(name)
	}
	return
}

func (ps *ProgramState) GetCurrent() *function.FunState {
	return ps.CurrentFunc
}

func (ps *ProgramState) SetCurrent(fs *function.FunState) {
	ps.CurrentFunc = fs
}
