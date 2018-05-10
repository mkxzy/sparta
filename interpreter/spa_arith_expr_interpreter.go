package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/operation"
)

type SPAArithExprInterpreter struct {
	ast *parser.Arith_exprContext
}

func(v *SPAArithExprInterpreter) Interpret(state *ProgramState)  {
	log.Debug("计算加减法")

	//v.EvalTerm(ctx.GetChild(0).(*parser.TermContext))
	termInter := &SPATermInterpreter{v.ast.GetChild(0).(*parser.TermContext)}
	termInter.Interpret(state)
	if v.ast.GetChildCount() > 1 {
		var op = ""
		for i := 1; i < v.ast.GetChildCount(); i++ {
			if i % 2 == 1 {
				op = v.ast.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			} else {
				termInter := &SPATermInterpreter{v.ast.GetChild(i).(*parser.TermContext)}
				termInter.Interpret(state)
				//v.EvalTerm(v.ast.GetChild(i).(*parser.TermContext))
				operation.Arithmetic(op)
			}
		}
	}
}
