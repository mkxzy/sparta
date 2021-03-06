package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/operation"
)

type SPACompareExprInterpreter struct {
	ast parser.ICompare_exprContext
}

// 实现解释接口
func(v *SPACompareExprInterpreter) Interpret(state *ProgramState)  {
	//v.EvalArithExpr(v.ast.GetChild(0).(*parser.Arith_exprContext))
	arithInter := &SPAArithExprInterpreter{v.ast.GetChild(0).(*parser.Arith_exprContext)}
	arithInter.Interpret(state)
	if v.ast.GetChildCount() == 3 {
		//v.EvalArithExpr(v.ast.GetChild(2).(*parser.Arith_exprContext))
		arithInter := &SPAArithExprInterpreter{v.ast.GetChild(2).(*parser.Arith_exprContext)}
		arithInter.Interpret(state)
		op := v.ast.GetChild(1).(*parser.Comp_opContext).GetChild(0).(antlr.TerminalNode).GetText()
		operation.Calculate(op)
		//switch op {
		//case "==":
		//	second := operation.PopValue().(types.SPAInteger)
		//	first := operation.PopValue().(types.SPAInteger)
		//	result := first == second
		//	operation.PushValue(types.SPABool(result))
		//default:
		//	panic("不支持的操作")
		//}
	}
}
