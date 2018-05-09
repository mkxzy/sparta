package interpreter

import "github.com/mkxzy/sparta/parser"

type SPAFactorInterpreter struct {
	ast *parser.FactorContext
}

func(v *SPAFactorInterpreter) Interpret(state *ProgramState)  {
	log.Debug("Visit Factor")

	if v.ast.GetChildCount() == 1{
		atomInter := &SPAAtomExprInterpreter{v.ast.GetChild(0).(*parser.Atom_exprContext)}
		atomInter.Interpret(state)
		//v.EvalAtomExpr(ctx.GetChild(0).(*parser.Atom_exprContext))
	} else{
		atomInter := &SPAAtomExprInterpreter{v.ast.GetChild(1).(*parser.Atom_exprContext)}
		atomInter.Interpret(state)
		//v.EvalAtomExpr(ctx.GetChild(1).(*parser.Atom_exprContext))
		//取反
		first := PopValue()
		result, _ := minus(first)
		PushValue(result)
	}
}
