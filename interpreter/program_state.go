package interpreter

import (
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

/**
程序状态
 */
type ProgramState struct {
	globals		*symbol.MemorySpace		//全局变量空间
	funList		*function.FunState		//函数调用链
}

func (ps *ProgramState) Define(s symbol.Symbol)  {
	if ps.funList != nil {
		ps.funList.Define(s)
	}else{
		ps.globals.Define(s)
	}
}

func (ps *ProgramState) Resolve(name string) (s symbol.Symbol) {
	s = ps.funList.Resolve(name)
	if s == nil {
		s = ps.globals.Resolve(name)
	}
	return
}
