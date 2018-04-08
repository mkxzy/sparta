package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/src/parser"
	"os"
)

type TreeShapeListener struct {
	*parser.BaseSpartaListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	//fmt.Println(input)
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	tree := p.Chunk()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	//fmt.Println(tree)
}
