package scope

import (
	"github.com/mkxzy/sparta/parser"
	"strings"
	"github.com/mkxzy/sparta/base"
)

/**
函数值
 */
type SPAFunction struct {
	Name string //函数名
	Args []string //参数名
	Locals []*VariableSymbol //局部变量
	Outer  Scope //外部作用域
	Body *parser.BlockContext //函数体（模拟指令地址）
}

func NewFunction(name string, args []string, body *parser.BlockContext) *SPAFunction  {
	return &SPAFunction{
		Name: name,
		Args: args,
		Body: body,
	}
}

func(f *SPAFunction) String() string {
	return "(" + strings.Join(f.Args, ", ") + ")"
}

// Symbol接口实现
func (ms *SPAFunction) GetName() string {
	return ms.Name
}

// Scope接口实现
func (table *SPAFunction) GetScopeName() string  {

	return table.Name
}

func (table *SPAFunction) GetEnclosingScope() Scope {

	return table.Outer
}

func (table *SPAFunction) Define(symbol Symbol) {
	switch symbol.(type) {
	case *VariableSymbol:
		var realSymbol = symbol.(*VariableSymbol)
		table.Locals = append(table.Locals, realSymbol)
	default:
		panic("语法错误")
	}
}

func (table *SPAFunction) Resolve(name string) Symbol {

	for _, x := range table.Locals {
		if x.GetName() == name{
			return x
		}
	}
	if table.Outer != nil {
		sym := table.Outer.Resolve(name)
		if sym != nil{
			return sym
		}
	}
	return nil
}

func(s SPAFunction) IsTrue() bool  {
	return true
}

func(ff *SPAFunction) PushArgs(args []base.SPAValue)  {
	for i := 0; i < len(ff.Args); i++{
		ff.Locals = append(ff.Locals, NewVariable(ff.Args[i], args[i]))
	}
}