package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/loop"
)

type SPAForStmtInterpreter struct {
	ast *parser.For_stmtContext
}

// 实现解释接口
func(v *SPAForStmtInterpreter) Interpret(state *ProgramState)  {
	//forState := &ForState{State:NORMAL}
	itemName := v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()

	fromInter := &SPATestInterpreter{v.ast.GetChild(3).(*parser.TestContext)}
	fromInter.Interpret(state)

	toInter := &SPATestInterpreter{v.ast.GetChild(5).(*parser.TestContext)}
	toInter.Interpret(state)
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

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *loop.ForBreak:
				return
			default:
				panic(err)
			}
		}
	}()

	sym := symbol.NewVariable(itemName, types.Null())
	state.Define(sym)
	for i := fromNumber; i <= toNumber; i++ {
		v.execLoopBody(sym, i, state)
	}
}

func(v *SPAForStmtInterpreter) execLoopBody(sym *symbol.SPAVariable, i types.SPAInteger, state *ProgramState) {

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *loop.ForContinue:
				return
			default:
				panic(err)
			}
		}
	}()

	sym.Value = i
	blockInter := &SPABlockInterpreter {
		ast: v.ast.GetChild(6).(*parser.BlockContext),
		//ff: forState,
	}
	blockInter.Interpret(state)
}
