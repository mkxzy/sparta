package interpreter

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/base"
)

/**
解释器接口
 */
type SPAExpression interface {
	interpret(ctx antlr.Tree) base.SPAValue
}

type SPANumberExpression struct {
}

func(expr *SPANumberExpression) interpret(ctx antlr.Tree) base.SPAValue {
	real := ctx.(*antlr.TerminalNodeImpl)
	return base.NewNumber(real.GetText())
}

type SPANamedFunDefExpression struct {
}

func(expr *SPANamedFunDefExpression) interpret(ctx *antlr.Tree) base.SPAValue {
	return base.Null()
}
