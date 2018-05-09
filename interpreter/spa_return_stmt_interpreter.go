package interpreter

import (
	"github.com/mkxzy/sparta/parser"
)

type SPAReturnStmtInterpreter struct {
	ast *parser.Return_stmtContext
	ff FlowState
}

// 实现解释接口
func(v *SPAReturnStmtInterpreter) Interpret(state *ProgramState)  {
	//v.ff.LoadState(RETURN)
	if state.currentFunc == nil{
		return
	}
	if v.ast.GetChildCount() == 2 {
		//计算返回值
		testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
		testInter.Interpret(state)
	} else{
		//如果return没有参数，那么插入一个空的返回值
		PushNullValue()
	}
	state.currentFunc.Return()
}
