package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
)

type SPAMapLiteralInterpreter struct {
	ast *parser.Map_literalContext
}

func(v *SPAMapLiteralInterpreter) Interpret() {
	result := types.NewMap()
	if v.ast.GetChildCount() == 3 {
		entryListContex := v.ast.GetChild(1).(*parser.Entry_listContext)
		for i := 0; i < entryListContex.GetChildCount(); i+=2 {
			entryInter := &SPAEntryInterpreter{entryListContex.GetChild(i).(*parser.EntryContext)}
			entryInter.Interpret()
			value := PopValue()
			key := PopValue()
			result.Put(key, value)
		}
	}
	PushValue(result)
}
