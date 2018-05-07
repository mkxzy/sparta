package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
)

type SPATableIndexInterpreter struct {
	ast *parser.Table_indexContext
}

func(v *SPATableIndexInterpreter) Interpret()  {

	//列表访问
	testInter := &SPATestInterpreter{v.ast.GetChild(2).(*parser.TestContext)}
	testInter.Interpret()
	name := v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	value := state.Resolve(name).(*symbol.SPAVariable).Value
	switch value.(type) {
	case *types.SPAList:
		index := PopValue().(types.SPAInteger)
		array := value.(*types.SPAList)
		PushValue(array.Get(index))
	case *types.SPAMap:
		key := PopValue()
		spaMap := value.(*types.SPAMap)
		PushValue(spaMap.Get(key))
	}
}
