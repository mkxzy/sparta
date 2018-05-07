package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
)

type SPAForStmtInterpreter struct {
	ast *parser.For_stmtContext
	ff FlowState
}

// 实现解释接口
func(v *SPAForStmtInterpreter) Interpret()  {
	forState := &ForState{State:NORMAL}
	forState.ItemName = v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	sym := symbol.NewVariable(forState.ItemName, types.Null())
	state.Define(sym)

	//v.EvalTest(ctx.GetChild(3).(*parser.TestContext))
	//v.EvalTest(ctx.GetChild(5).(*parser.TestContext))
	fromInter := &SPATestInterpreter{v.ast.GetChild(3).(*parser.TestContext)}
	fromInter.Interpret()

	toInter := &SPATestInterpreter{v.ast.GetChild(5).(*parser.TestContext)}
	toInter.Interpret()
	toNumber, ok := PopValue().(types.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	fromNumber, ok := PopValue().(types.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	if fromNumber > toNumber{
		panic("起始数字不能大于结束数字")
	}

	for i := fromNumber; i <= toNumber; i++ {
		sym.Value = types.SPAInteger(i)
		//v.ExecBlock(ctx.GetChild(6).(*parser.BlockContext), forState)
		blockInter := &SPABlockInterpreter{
			ast: v.ast.GetChild(6).(*parser.BlockContext),
			ff: forState,
		}
		blockInter.Interpret()
		if forState.State == BREAK {
			forState.SetState(NORMAL) //恢复状态
		}
		if forState.State == CONTINUE {
			forState.SetState(NORMAL) //恢复状态
		}
		if forState.State == RETURN {
			forState.SetState(NORMAL) //恢复状态
			v.ff.SetState(RETURN) 	  //传递给调用者
		}
	}
}
