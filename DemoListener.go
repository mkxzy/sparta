package main

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
)

type DemoListener struct {
	parser.BaseSpartaListener
}

func (v *DemoListener) EnterEveryRule(ctx antlr.ParserRuleContext)  {
	fmt.Println(ctx.GetText())
}
