package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
)

type SPATestListInterpreter struct {
	ast *parser.Test_listContext
}

func (v *SPATestListInterpreter) Interpret() {
	elementCount := (v.ast.GetChildCount() + 1) / 2
	list := types.NewList(elementCount)
	for i := 0; i < v.ast.GetChildCount(); i += 2{
		//v.EvalTest(ctx.GetChild(i).(*parser.TestContext))
		testInter := &SPATestInterpreter{v.ast.GetChild(i).(*parser.TestContext)}
		testInter.Interpret()
		value := PopValue()
		list.Append(value)
	}
	PushValue(list)
	log.Info(list)
}
