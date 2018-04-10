package interpreter

import (
	"fmt"
	"github.com/mkxzy/sparta/parser"
	//"reflect"
)

type Interpreter struct {
	parser.BaseSpartaListener

	Symbols map[string]Symbol
	GlobalSpace []interface{}
	GlobalIndex int
	CurrentScope Scope
}

func(s *Interpreter) defineSymbol(symbol Symbol)  {
	s.Symbols[symbol.GetName()] = symbol
}

func(s *Interpreter) pushGlobalVar(v interface{})  {
	s.GlobalSpace[s.GlobalIndex] = v
	s.GlobalIndex++
}

func(s *Interpreter) setScope(scope Scope)  {
	s.CurrentScope = scope
}

func(s *Interpreter) String()  {
	fmt.Println(s.Symbols)
	fmt.Println(s.GlobalSpace)
	fmt.Println(s.GlobalIndex)
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
		GlobalSpace: make([]interface{}, 500),
		GlobalIndex: 0,
		CurrentScope: nil,
	}
	return interpreter
}

func (s *Interpreter) EnterChunk(ctx *parser.ChunkContext) {
	s.CurrentScope = NewGlobalScope()
}

func (s *Interpreter) ExitChunk(ctx *parser.ChunkContext) {
	fmt.Println(s.CurrentScope.GetScopeName())
	fmt.Println(s.GlobalSpace)
	fmt.Println(s.GlobalIndex)
}

// EnterStat is called when production stat is entered.
func (s *Interpreter) EnterStat(ctx *parser.StatContext) {
	//for i:=0;i<ctx.GetChildCount();i++{
	//	fmt.Println(ctx.GetChild(i))
	//}
	s.defineVariable("a", "abcdefg")
	s.defineVariable("b", "abcdefg")
}

// ExitStat is called when production stat is exited.
func (s *Interpreter) ExitStat(ctx *parser.StatContext) {

}

// EnterStr is called when production str is entered.
func (s *Interpreter) EnterStr(ctx *parser.StrContext) {}

// ExitStr is called when production str is exited.
func (s *Interpreter) ExitStr(ctx *parser.StrContext) {}