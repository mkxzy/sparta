package interpreter

import "github.com/mkxzy/sparta/parser"

type SPATestInterpreter struct {
	ast *parser.TestContext
}

// 实现解释接口
func(v *SPATestInterpreter) Interpret(state *ProgramState)  {

	//compareInter := &SPACompareExprInterpreter{
	//	ast: v.ast.GetChild(0).(*parser.Compare_exprContext),
	//}
	//compareInter.Interpret(state)

	orInter := &SPAOrTestInterpreter{
		ast: v.ast.GetChild(0).(*parser.Or_testContext),
	}
	orInter.Interpret(state)
}
