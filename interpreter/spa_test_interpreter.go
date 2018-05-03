package interpreter

import "github.com/mkxzy/sparta/parser"

type SPATestInterpreter struct {
	ast *parser.TestContext
}

// 实现解释接口
func(v *SPATestInterpreter) Interpret()  {

	//v.EvalCompareExpr(ctx.GetChild(0).(*parser.Compare_exprContext))

	compareInter := &SPACompareExprInterpreter{
		ast: v.ast.GetChild(0).(*parser.Compare_exprContext),
	}
	compareInter.Interpret()
}
