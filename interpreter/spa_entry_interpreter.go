package interpreter

import (
	"github.com/mkxzy/sparta/parser"
)

type SPAEntryInterpreter struct {
	ast *parser.EntryContext
}

func(v *SPAEntryInterpreter) Interpret()  {

	keyInter := &SPATestInterpreter{ast: v.ast.GetChild(0).(*parser.TestContext)}
	keyInter.Interpret()

	valueInter := &SPATestInterpreter{ast: v.ast.GetChild(2).(*parser.TestContext)}
	valueInter.Interpret()
}
