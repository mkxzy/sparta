package main

import (
	//"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/parser"
	"os"
	"fmt"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	//fmt.Println(input)
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	tree := p.Chunk()
	fmt.Println(tree.ToStringTree(nil,p))
	visitor := &DemoVisitor{}
	//visitor.Visit(tree)
	tree.Accept(visitor)
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	//fmt.Println(tree)
}
