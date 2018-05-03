package interpreter

import "github.com/mkxzy/sparta/parser"

type SPABreakStmtInterpreter struct {
	ast parser.IBreak_stmtContext
	ff FlowState
}

// 实现解释接口
func(v *SPABreakStmtInterpreter) Interpret()  {
	v.ff.SetState(BREAK)
}
