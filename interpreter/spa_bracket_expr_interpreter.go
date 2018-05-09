package interpreter

import (
	"github.com/mkxzy/sparta/parser"
)

type SPABracketExprInterpreter struct {
	ast *parser.Bracket_exprContext
}

func(v *SPABracketExprInterpreter) Interpret(state *ProgramState)  {

	testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
	testInter.Interpret(state)
}
