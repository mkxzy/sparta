package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/operation"
)

type SPATableIndexInterpreter struct {
	ast *parser.Table_indexContext
}

func(v *SPATableIndexInterpreter) Interpret(state *ProgramState)  {

	//列表访问
	testInter := &SPATestInterpreter{v.ast.GetChild(2).(*parser.TestContext)}
	testInter.Interpret(state)
	name := v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	value := state.Resolve(name).(*symbol.SPAVariable).Value
	switch value.(type) {
	case *types.SPAList:
		index := operation.PopValue().(types.SPAInteger)
		array := value.(*types.SPAList)
		operation.PushValue(array.Get(index))
	case *types.SPAMap:
		key := operation.PopValue()
		spaMap := value.(*types.SPAMap)
		operation.PushValue(spaMap.Get(key))
	}
}
