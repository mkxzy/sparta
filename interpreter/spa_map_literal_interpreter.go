package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/operation"
)

type SPAMapLiteralInterpreter struct {
	ast *parser.Map_literalContext
}

func(v *SPAMapLiteralInterpreter) Interpret(state *ProgramState) {
	result := types.NewMap()
	if v.ast.GetChildCount() == 3 {
		entryListContex := v.ast.GetChild(1).(*parser.Entry_listContext)
		for i := 0; i < entryListContex.GetChildCount(); i+=2 {
			entryInter := &SPAEntryInterpreter{entryListContex.GetChild(i).(*parser.EntryContext)}
			entryInter.Interpret(state)
			value := operation.PopValue()
			key := operation.PopValue()
			result.Put(key, value)
		}
	}
	operation.PushValue(result)
}
