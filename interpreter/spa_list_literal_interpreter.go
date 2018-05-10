package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/operation"
)

type SPAListLiteralInterpreter struct {
	ast *parser.List_literalContext
}

// 实现解释接口
func(v *SPAListLiteralInterpreter) Interpret(state *ProgramState)  {
	if v.ast.GetChildCount() == 3 {
		testListContext := v.ast.GetChild(1).(*parser.Test_listContext)
		elements := (testListContext.GetChildCount() + 1) / 2
		li := types.NewList(elements)
		for i := 0; i < testListContext.GetChildCount(); i+=2 {
			testInter := &SPATestInterpreter{ast: testListContext.GetChild(i).(*parser.TestContext)}
			testInter.Interpret(state)
			item := operation.PopValue()
			li.Append(item)
		}
		operation.PushValue(li)
	}else{
		operation.PushValue(types.NewList(0))
	}
}
