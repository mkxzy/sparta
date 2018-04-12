package interpreter

import (
	"fmt"
	"github.com/mkxzy/sparta/parser"
	"reflect"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"bytes"
)

type DemoInterpreter struct {
	parser.BaseSpartaListener

	globalTop int
	Symbols map[string]Symbol
	GlobalSpace []interface{}
	CurrentScope Scope
}

func NewInterpreter() *DemoInterpreter {
	interpreter := &DemoInterpreter{
		Symbols: make(map[string]Symbol),
		GlobalSpace: make([]interface{}, 0, 500),
		CurrentScope: nil,
		globalTop: -1,
	}
	return interpreter
}

func(s *DemoInterpreter) defineSymbol(symbol Symbol)  {
	s.Symbols[symbol.GetName()] = symbol
}

func(s *DemoInterpreter) pushGlobalVar(v interface{}) (addr int)  {
	s.GlobalSpace = append(s.GlobalSpace, v)
	s.globalTop++
	addr = s.globalTop
	return
}

func(s *DemoInterpreter) setScope(scope Scope)  {
	s.CurrentScope = scope
}

func(s *DemoInterpreter) defineVariable(name string, v interface{})  {
	sym := &VariableSymbol{
		Name: name,
	}
	s.defineSymbol(sym)
	s.pushGlobalVar(v)
}

func(s DemoInterpreter) String() string {

	var buf bytes.Buffer
	for k, v := range s.Symbols{
		buf.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	buf.WriteString(fmt.Sprint(s.GlobalSpace))
	return buf.String()
}

func (s *DemoInterpreter) EnterProgram(ctx *parser.ProgramContext) {
	s.CurrentScope = NewGlobalScope()
}

func (s *DemoInterpreter) ExitProgram(ctx *parser.ProgramContext) {
	//fmt.Println(s.CurrentScope.GetScopeName())
	//fmt.Println(s.Symbols)
	fmt.Println(s)
}

// EnterStat is called when production stat is entered.
func (s *DemoInterpreter) EnterStat(ctx *parser.StatContext) {
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
func (s *DemoInterpreter) ExitStat(ctx *parser.StatContext) {

}

// EnterStr is called when production str is entered.
func (s *DemoInterpreter) EnterStr(ctx *parser.StrContext) {}

// ExitStr is called when production str is exited.
func (s *DemoInterpreter) ExitStr(ctx *parser.StrContext) {}