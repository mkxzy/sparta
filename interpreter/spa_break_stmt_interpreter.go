package interpreter

import (
	"github.com/mkxzy/sparta/loop"
)

type SPABreakStmtInterpreter struct {
	//ast parser.IBreak_stmtContext
	//ff FlowState
}

// 实现解释接口
func(v *SPABreakStmtInterpreter) Interpret(state *ProgramState)  {
	//v.ff.SetState(BREAK)
	loop.Break()
}
