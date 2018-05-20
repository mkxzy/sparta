package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/operation"
	"github.com/mkxzy/sparta/types"
)

type SPAAndTestInterpreter struct {
	ast *parser.And_testContext
}

// 实现解释接口
func(v *SPAAndTestInterpreter) Interpret(state *ProgramState) {

	var testValue types.SPAValue
	for i := 0; i < v.ast.GetChildCount(); i += 2{
		notInter := &SPANotTestInterpreter{
			ast: v.ast.GetChild(i).(*parser.Not_testContext),
		}
		notInter.Interpret(state)
		testValue = operation.PopValue().(types.SPAValue)
		if !testValue.IsTrue() {
			goto SET_VAL
		}
	}

SET_VAL:
	operation.PushValue(testValue)
}
