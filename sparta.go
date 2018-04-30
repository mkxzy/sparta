package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/interpreter"
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
	"os"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

//var log = logging.MustGetLogger("sparta")

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
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	ctx := p.Program()        // 生成解析树
	inter := getInterpreter() // 创建解释器
	inter.Interpret(ctx)

	//tree.Accept(inter) 	  // 解释执行
	//listener := interpreter.NewInterpreter()
	//antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	//fmt.Println(tree)
}

func getInterpreter() interpreter.SPAInterpreter {
	globalState := symbol.NewMemorySpace("global")                  //全局内存空间
	printFunc := function.NewInternalFunction("print", []string{"s"}) //print内置函数
	globalState.Define(function.NewFunVariable(printFunc))
	return interpreter.NewDirectInterpreter(globalState)
}
