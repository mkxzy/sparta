package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/function"
)

type SPAFundefStmtInterpreter struct {
	ast parser.IFundef_stmtContext
	f *function.SPAFunction
}

// 实现解释接口
func(v *SPAFundefStmtInterpreter) Interpret()  {
	v.f = &function.SPAFunction{}
	v.f.Name = v.ast.GetChild(1).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	v.f.Body = v.ast.GetChild(2).(*parser.Fun_bodyContext)
	v.f.Args = []string{}
	funParContext := v.f.Body.GetChild(0)
	if funParContext.GetChildCount() > 2 {
		namelistContext := funParContext.GetChild(1).(*parser.NamelistContext)
		for i := 0; i < namelistContext.GetChildCount(); i += 2{
			name := namelistContext.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			v.f.Args = append(v.f.Args, name)
		}
	}
	//保存函数调用链
	v.f.FunList = state.funList
	sym := function.NewFunVariable(*v.f)
	state.Define(sym) 		//函数定义
	log.Infof("函数定义: %v", state)
}
