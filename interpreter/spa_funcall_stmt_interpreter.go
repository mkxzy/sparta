package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/operation"
)

type SPAFuncallStmtInterpreter struct {
	ast parser.IFuncall_stmtContext
}

// 实现解释接口
func(v *SPAFuncallStmtInterpreter) Interpret(state *ProgramState)  {
	//v.EvalFunCallExpr(ctx.GetChild(0).(*parser.Funcall_exprContext))
	funCallExprInter := &SPAFuncallExprInterpreter{
		ast: v.ast.GetChild(0).(*parser.Funcall_exprContext),
	}
	funCallExprInter.Interpret(state)
	operation.PopValue() //丢弃返回值
}
