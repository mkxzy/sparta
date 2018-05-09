package interpreter

import "github.com/mkxzy/sparta/parser"

type SPAIfStmtInterpreter struct {
	ast *parser.If_stmtContext
	ff *ForState
}

// 实现解释接口
func(v *SPAIfStmtInterpreter) Interpret(state *ProgramState)  {
	testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
	testInter.Interpret(state)
	testResult := PopValue()
	if testResult.IsTrue() {
		blockInter := &SPABlockInterpreter{v.ast.GetChild(2).(*parser.BlockContext), v.ff}
		blockInter.Interpret(state)
		return
	}
	pos := 3
	for pos + 4 < v.ast.GetChildCount() {
		testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
		testInter.Interpret(state)
		testResult = PopValue()
		if testResult.IsTrue() {
			blockInter := &SPABlockInterpreter{v.ast.GetChild(pos + 3).(*parser.BlockContext), v.ff}
			blockInter.Interpret(state)
			return
		}
		pos += 4
	}
	if pos < v.ast.GetChildCount() {
		blockInter := &SPABlockInterpreter{v.ast.GetChild(pos + 1).(*parser.BlockContext), v.ff}
		blockInter.Interpret(state)
	}
}