package interpreter

import (
	"github.com/mkxzy/sparta/parser"
)

type SPAEntryInterpreter struct {
	ast *parser.EntryContext
}

func(v *SPAEntryInterpreter) Interpret(state *ProgramState)  {

	keyInter := &SPATestInterpreter{ast: v.ast.GetChild(0).(*parser.TestContext)}
	keyInter.Interpret(state)

	valueInter := &SPATestInterpreter{ast: v.ast.GetChild(2).(*parser.TestContext)}
	valueInter.Interpret(state)
}
