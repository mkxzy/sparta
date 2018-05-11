package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/interpreter"
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
	"os"
	"flag"
	"fmt"
)

//var log = logging.MustGetLogger("sparta")

var (
	ShowVersion bool
	ShowHelp bool
	Version = "0.6_pre"
)

func init() {
	//var format = logging.MustStringFormatter(
	//	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	//)
	//backend := logging.NewLogBackend(os.Stdout, "", 0)
	//formatter := logging.NewBackendFormatter(backend, format)
	//logging.SetBackend(backend, formatter)
	//backend1Leveled := logging.AddModuleLevel(backend)
	logging.SetLevel(logging.ERROR, "")
	flag.BoolVar(&ShowVersion, "v", false, "show version and exit")
	flag.BoolVar(&ShowHelp, "h", false, "this help")
	//flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stdout, `sparta %s
Usage: 	sparta [-hv]
		sparta [filename]

Options:
`, Version)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if len(os.Args) == 1{
		usage()
		os.Exit(0)
	}
	if ShowHelp {
		usage()
		os.Exit(0)
	}
	if ShowVersion {
		fmt.Println(Version)
		os.Exit(0)
	}
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	ast := p.Program()                             	// 生成解析树
	state := interpreter.NewProgramState()			// 程序解释器状态数据结构
	inter := interpreter.NewDirectInterpreter(ast) 	// 创建解释器
	inter.Interpret(state)

	//tree.Accept(inter) 	  // 解释执行
	//listener := interpreter.NewInterpreter()
	//antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	//fmt.Println(tree)
}
