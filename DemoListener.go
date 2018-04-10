package main

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"github.com/mkxzy/sparta/interpreter"
)

type DemoListener struct {
	parser.BaseSpartaListener
	symbols *interpreter.SymbolTable
}

func (v *DemoListener) EnterEveryRule(ctx antlr.ParserRuleContext)  {
	//fmt.Println(ctx.GetText())
}

func (v *DemoListener) EnterStr(ctx *parser.StrContext) {
	fmt.Println(ctx.GetText())
}

func (s *DemoListener) EnterFuncname(ctx *parser.FuncnameContext){
	symbol := &interpreter.FunctionSymbol{
		Name: ctx.GetText(),
		//Type: nil,
	}
	s.symbols.Define(symbol)
	fmt.Println(ctx.GetText())
	//s.symbols.Define(ctx.GetText())
}

func (s *DemoListener) EnterChunk(ctx *parser.ChunkContext) {
	s.symbols = interpreter.NewSymbolTable()
}

func (s *DemoListener) ExitChunk(ctx *parser.ChunkContext) {
	fmt.Println(s.symbols)
}
