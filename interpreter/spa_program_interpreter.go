/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
)

var log = logging.MustGetLogger("SPAProgramInterpreter")

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
func(v *SPAProgramInterpreter) Interpret(state *ProgramState)  {
	for i := 0; i < v.ast.GetChildCount()-1; i++ {
		stmtContext := v.ast.GetChild(i).(*parser.StmtContext)
		stmtInter := &SPAStmtInterpreter{ast: stmtContext}
		stmtInter.Interpret(state)
	}
}