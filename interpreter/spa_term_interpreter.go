package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SPATermInterpreter struct {
	ast *parser.TermContext
}

func(v *SPATermInterpreter) Interpret()  {
	log.Debug("计算乘除法")

	//v.EvalFactor(v.ast.GetChild(0).(*parser.FactorContext))
	factorInter := &SPAFactorInterpreter{v.ast.GetChild(0).(*parser.FactorContext)}
	factorInter.Interpret()
	if v.ast.GetChildCount() > 1 {
		var op = ""
		for i := 1; i < v.ast.GetChildCount(); i++ {
			if i % 2 == 1 {
				op = v.ast.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			} else {
				//v.EvalFactor(v.ast.GetChild(i).(*parser.FactorContext))
				factorInter := &SPAFactorInterpreter{v.ast.GetChild(i).(*parser.FactorContext)}
				factorInter.Interpret()
				arithmetic(op)
			}
		}
	}
}
