/**
函数符号
 */

package interpreter

type FunctionSymbol struct {
	Name   string
	Args   []*VariableSymbol
	Locals []*VariableSymbol
	Addr   int
	Outer  *FunctionSymbol
}

// Symbol接口实现
func (ms *FunctionSymbol) GetName() string {
	return ms.Name
}

// Scope接口实现
func (table *FunctionSymbol) GetScopeName() string  {

	return table.Name
}

func (table *FunctionSymbol) GetEnclosingScope() Scope {

	return table.Outer
}

func (table *FunctionSymbol) Define(symbol Symbol) {
	switch symbol.(type) {
	case *FunctionSymbol:
		var realSymbol = symbol.(*FunctionSymbol)
		realSymbol.Outer = table
	case *VariableSymbol:
		var realSymbol = symbol.(*VariableSymbol)
		table.Locals = append(table.Locals, realSymbol)
	default:
		panic("语法错误")
	}
}

func (table *FunctionSymbol) Resolve(name string) Symbol {
	for _, x := range table.Args {
		if x.GetName() == name{
			s := &FunctionSymbol{
				Name: name,
			}
			return s
		}
	}
	for _, x := range table.Locals {
		if x.GetName() == name{
			s := &FunctionSymbol{
				Name: name,
			}
			return s
		}
	}
	return nil
}
// Scope接口实现

func NewFuncionSymbol(name string) *FunctionSymbol{
	f := &FunctionSymbol{
		Name: name,
	}
	return f
}