package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/operation"
	"github.com/mkxzy/sparta/types"
)

type SPAOrTestInterpreter struct {
	ast *parser.Or_testContext
}

// 实现解释接口
func(v *SPAOrTestInterpreter) Interpret(state *ProgramState) {

	var testValue types.SPAValue
	for i := 0; i < v.ast.GetChildCount(); i += 2{
		andInter := &SPAAndTestInterpreter{
			ast: v.ast.GetChild(i).(*parser.And_testContext),
		}
		andInter.Interpret(state)
		testValue = operation.PopValue().(types.SPAValue)
		log.Info(testValue)
		if testValue.IsTrue() {
			goto SET_VAL
		}
	}

SET_VAL:
	operation.PushValue(testValue)
}
