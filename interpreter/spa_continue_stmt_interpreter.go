package interpreter

import (
	"github.com/mkxzy/sparta/loop"
)

type SPAContinueStmtInterpreter struct {
	//ast parser.IContinue_stmtContext
	//ff FlowState
}

// 实现解释接口
func(v *SPAContinueStmtInterpreter) Interpret(state *ProgramState)  {
	//v.ff.SetState(CONTINUE)
	loop.Continue()
}
