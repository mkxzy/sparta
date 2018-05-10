package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/operation"
)

type SPAAssignStmtInterpreter struct {
	ast parser.IAssign_stmtContext
}

// 实现解释接口
func(v *SPAAssignStmtInterpreter) Interpret(state *ProgramState)  {

	// 解析表达式
	testInter := &SPATestInterpreter{v.ast.GetChild(2).(*parser.TestContext)}
	testInter.Interpret(state)

	leftSideInter := &SPALeftSideInterpreter{
		ast: v.ast.GetChild(0).(*parser.Left_sideContext),
	}
	leftSideInter.Interpret(state)

	if leftSideInter.varType == Scalar {
		value := operation.PopValue()
		sym := symbol.NewVariable(leftSideInter.varName, value)
		state.Define(sym)
	} else {
		var sym symbol.Symbol
		sym = state.Resolve(leftSideInter.varName)
		if sym == nil {
			panic("")
		}
		list := sym.(*symbol.SPAVariable).Value.(*types.SPAList)
		index, ok := operation.PopValue().(types.SPAInteger)
		if !ok {
			panic("index必须是整数")
		}
		v := operation.PopValue()
		list.Set(index, v)
	}
}
