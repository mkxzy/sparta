package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/function"
)

type SPAFundefStmtInterpreter struct {
	ast parser.IFundef_stmtContext
}

// 实现解释接口
func(v *SPAFundefStmtInterpreter) Interpret(state *ProgramState)  {
	f := function.SPAFunction{}
	f.Name = v.ast.GetChild(1).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	f.Body = v.ast.GetChild(2).(*parser.Fun_bodyContext)
	f.Args = []string{}
	funParContext := f.Body.GetChild(0)
	if funParContext.GetChildCount() > 2 {
		namelistContext := funParContext.GetChild(1).(*parser.NamelistContext)
		for i := 0; i < namelistContext.GetChildCount(); i += 2{
			name := namelistContext.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			f.Args = append(f.Args, name)
		}
	}
	//保存函数调用链
	f.FS = state.CurrentState()       //保存函数闭包空间
	sym := function.NewFunVariable(f) //创建函数变量
	state.Define(sym)                 //定义函数
	log.Infof("函数定义: %v", state)
}
