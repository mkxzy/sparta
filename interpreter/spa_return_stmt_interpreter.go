package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
)

type SPAReturnStmtInterpreter struct {
	ast *parser.Return_stmtContext
	ff FlowState
}

// 实现解释接口
func(v *SPAReturnStmtInterpreter) Interpret()  {
	v.ff.SetState(RETURN)
	if v.ast.GetChildCount() == 2 {
		//计算返回值
		testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
		testInter.Interpret()
	} else{
		//如果return没有参数，那么插入一个空的返回值
		PushValue(types.Null())
	}
}
