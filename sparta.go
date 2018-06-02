package main

import (
	"github.com/mkxzy/sparta/interpreter"
	"github.com/op/go-logging"
	"os"
	"flag"
	"fmt"
	"path/filepath"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
)

var (
	ShowVersion bool
	ShowHelp bool
	Version = "v0.6"
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
		fmt.Printf("Sparta %s", Version)
		os.Exit(0)
	}

	fileName, _ := filepath.Abs(os.Args[1])
	fileDir := filepath.Dir(fileName)
	//currentFile := filepath.Dir(fileName)
	fmt.Println(fileName)
	inter := interpreter.NewDirectInterpreterFromFile(fileName)	//创建解释器

	state := interpreter.NewProgramState()						//程序解释器状态数据结构
	state.Define(symbol.NewVariable("CURRENT_FILE", types.SPAString(fileName)))
	state.Define(symbol.NewVariable("CURRENT_DIR", types.SPAString(fileDir)))
	inter.Interpret(state)
}
