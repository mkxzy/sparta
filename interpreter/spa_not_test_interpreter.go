package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/operation"
	"github.com/mkxzy/sparta/types"
)

type SPANotTestInterpreter struct {
	ast *parser.Not_testContext
}

// 实现解释接口
func(v *SPANotTestInterpreter) Interpret(state *ProgramState) {
	if v.ast.GetChildCount() == 1 {
		compareInter := &SPACompareExprInterpreter{
			ast: v.ast.GetChild(0).(*parser.Compare_exprContext),
		}
		compareInter.Interpret(state)
	} else {
		testValue := operation.PopValue().(types.SPAValue)
		operation.PushValue(types.InvertBool(testValue))
	}
}
