package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/operation"
)

type SPAFactorInterpreter struct {
	ast *parser.FactorContext
}

func(v *SPAFactorInterpreter) Interpret(state *ProgramState)  {
	log.Debug("Visit Factor")

	if v.ast.GetChildCount() == 1{
		atomInter := &SPAAtomExprInterpreter{v.ast.GetChild(0).(*parser.Atom_exprContext)}
		atomInter.Interpret(state)
	} else{
		atomInter := &SPAAtomExprInterpreter{v.ast.GetChild(1).(*parser.Atom_exprContext)}
		atomInter.Interpret(state)
		operation.Minus()
	}
}
