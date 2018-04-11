package interpreter

import (
	"fmt"
	"github.com/mkxzy/sparta/parser"
	//"reflect"
	//"reflect"
	//"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Interpreter struct {
	parser.BaseSpartaListener

	Symbols map[string]Symbol
	GlobalSpace []interface{}
	CurrentScope Scope
}

func(s *Interpreter) defineSymbol(symbol Symbol)  {
	s.Symbols[symbol.GetName()] = symbol
}

func(s *Interpreter) pushGlobalVar(v interface{})  {
	s.GlobalSpace = append(s.GlobalSpace, v)
}

func(s *Interpreter) setScope(scope Scope)  {
	s.CurrentScope = scope
}

func(s *Interpreter) String()  {
	fmt.Println(s.Symbols)
	fmt.Println(s.GlobalSpace)
}

func(s *Interpreter) defineVariable(name string, v interface{})  {
	sym := VariableSymbol{
		Name: name,
	}
	s.defineSymbol(sym)
	s.pushGlobalVar(v)
}

//执行指令
func exec(address int)  {
}

func NewInterpreter() *Interpreter{
	interpreter := &Interpreter{
		Symbols: make(map[string]Symbol),
		GlobalSpace: make([]interface{}, 0, 500),
		CurrentScope: nil,
	}
	return interpreter
}

func (s *Interpreter) EnterProgram(ctx *parser.ProgramContext) {
	s.CurrentScope = NewGlobalScope()
}

func (s *Interpreter) ExitProgram(ctx *parser.ProgramContext) {
	fmt.Println(s.CurrentScope.GetScopeName())
	fmt.Println(s.GlobalSpace)
}

// EnterStat is called when production stat is entered.
func (s *Interpreter) EnterStat(ctx *parser.StatContext) {
	//for i:=0;i<ctx.GetChildCount();i++{
	//	fmt.Println(ctx.GetChild(i))
	//}
	fmt.Println(ctx.ToStringTree(nil, ctx.GetParser()))
	name := ctx.GetChild(0).GetChild(1).GetChild(0).(antlr.TerminalNode)
	value := ctx.GetChild(0).GetChild(3).GetChild(0).GetChild(0).GetChild(0).(antlr.TerminalNode)
	fmt.Println(value.(antlr.ParseTree).ToStringTree(nil, ctx.GetParser()))
	//fmt.Println(value(nil, ctx.GetParser()))
	fmt.Println(reflect.TypeOf(value))
	s.defineVariable(name.GetText(), value.GetText())
	//s.defineVariable("b", "abcdefg")
}

// ExitStat is called when production stat is exited.
func (s *Interpreter) ExitStat(ctx *parser.StatContext) {

}

// EnterStr is called when production str is entered.
func (s *Interpreter) EnterStr(ctx *parser.StrContext) {}

// ExitStr is called when production str is exited.
func (s *Interpreter) ExitStr(ctx *parser.StrContext) {}