package interpreter

import "github.com/mkxzy/sparta/parser"

type SPABlockInterpreter struct {
	ast parser.IBlockContext
	ff FlowState
}

// 实现解释接口
func(v *SPABlockInterpreter) Interpret(state *ProgramState)  {
	for i := 1; i < v.ast.GetChildCount()-1; i++ {
		stmtContext := v.ast.GetChild(i).(*parser.StmtContext)
		stmtInter := &SPAStmtInterpreter{stmtContext, v.ff}
		stmtInter.Interpret(state)

		// return，continue,break都要中断执行块
		if v.ff.GetState() > NORMAL {
			break
		}
	}
}
