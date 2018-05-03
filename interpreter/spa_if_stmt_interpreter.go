package interpreter

import "github.com/mkxzy/sparta/parser"

type SPAIfStmtInterpreter struct {
	ast *parser.If_stmtContext
	ff FlowState
}

// 实现解释接口
func(v *SPAIfStmtInterpreter) Interpret()  {
	testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
	testInter.Interpret()
	testResult := PopValue()
	if testResult.IsTrue() {
		blockInter := &SPABlockInterpreter{v.ast.GetChild(2).(*parser.BlockContext), v.ff}
		blockInter.Interpret()
		return
	}
	pos := 3
	for pos + 4 < v.ast.GetChildCount() {
		testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
		testInter.Interpret()
		testResult = PopValue()
		if testResult.IsTrue() {
			blockInter := &SPABlockInterpreter{v.ast.GetChild(pos + 3).(*parser.BlockContext), v.ff}
			blockInter.Interpret()
			return
		}
		pos += 4
	}
	if pos < v.ast.GetChildCount() {
		blockInter := &SPABlockInterpreter{v.ast.GetChild(pos + 1).(*parser.BlockContext), v.ff}
		blockInter.Interpret()
	}
}