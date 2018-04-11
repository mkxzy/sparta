package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/interpreter"
	"github.com/mkxzy/sparta/parser"
)

type DemoListener struct {
	parser.BaseSpartaListener
	symbols *interpreter.GlobalScope
}

func (v *DemoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}

func (v *DemoListener) EnterStr(ctx *parser.StrContext) {
	fmt.Println(ctx.GetText())
}

func (s *DemoListener) EnterFuncname(ctx *parser.FuncnameContext) {
	symbol := &interpreter.FunctionSymbol{
		Name: ctx.GetText(),
		//Type: nil,
	}
	s.symbols.Define(symbol)
	fmt.Println(ctx.GetText())
	//s.symbols.Define(ctx.GetText())
}

func (s *DemoListener) EnterChunk(ctx *parser.ProgramContext) {
	s.symbols = interpreter.NewGlobalScope()
}

func (s *DemoListener) ExitChunk(ctx *parser.ProgramContext) {
	fmt.Println(s.symbols)
}
