package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/parser"
	"os"
	//"fmt"
	//"github.com/mkxzy/sparta/interpreter"
	"github.com/mkxzy/sparta/interpreter"
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/scope"
)

var log = logging.MustGetLogger("SPAInterpreter")

func init() {
	//var format = logging.MustStringFormatter(
	//	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	//)
	//backend := logging.NewLogBackend(os.Stdout, "", 0)
	//formatter := logging.NewBackendFormatter(backend, format)
	//logging.SetBackend(backend, formatter)
	//backend1Leveled := logging.AddModuleLevel(backend)
	logging.SetLevel(logging.DEBUG, "")
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	//fmt.Println(input)
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	tree := p.Program()
	//log.Debug(tree.ToStringTree(nil, p))
	state := scope.NewGlobalScope()
	inter := interpreter.NewExpVisitor(state)
	//fmt.Println(visitor)
	//visitor.Visit(tree)
	tree.Accept(inter)

	//listener := interpreter.NewInterpreter()
	//antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	//fmt.Println(tree)
}
