/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

var log = logging.MustGetLogger("SPAProgramInterpreter")

var state *ProgramState

func init()  {
	globals := symbol.NewMemorySpace("global")                        //全局内存空间
	printFunc := function.NewInternalFunction("print", []string{"s"}) //print内置函数
	globals.Define(function.NewFunVariable(printFunc))
	state = & ProgramState{
		//currentScope: globals,
		//savedScope:   nil,
		globals:	globals,
		funList:	nil,
	}
}

type SPAProgramInterpreter struct {
	ast parser.IProgramContext
	currentScope symbol.Scope
	savedScope symbol.Scope
}

func NewDirectInterpreter(ast parser.IProgramContext) *SPAProgramInterpreter {
	return &SPAProgramInterpreter{
		ast: ast,
	}
}

// 实现解释接口
func(v *SPAProgramInterpreter) Interpret()  {
	for i := 0; i < v.ast.GetChildCount()-1; i++ {
		stmtContext := v.ast.GetChild(i).(*parser.StmtContext)
		stmtInter := &SPAStmtInterpreter{ast: stmtContext}
		stmtInter.Interpret()
	}
}