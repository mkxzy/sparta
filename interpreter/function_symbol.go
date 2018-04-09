package interpreter

type FunctionSymbol struct {
	Name string
	Args []string
	Locals []string
	Addr int
	Parent *FunctionSymbol
}

func (ms FunctionSymbol) GetName() string {
	return ms.Name
}

func (table *FunctionSymbol) GetScopeName() string  {
	return table.Name
}

func (table *FunctionSymbol) GetEnclosingScope() Scope {
	//switch table.Parent.(type) {
	//case Scope:
	//	return table.Parent
	//default:
	//	return nil
	//}
	return table.Parent
}

func (table *FunctionSymbol) Define(symbol Symbol) {
	//table.Symbols[symbol.GetName()] = symbol
}

func (table *FunctionSymbol) Resolve(name string) Symbol {
	//return table.Symbols[name]
	//var s Symbol = nil
	for _, x := range table.Args {
		if x == name{
			s := FunctionSymbol{
				Name: name,
			}
			return s
		}
	}
	for _, x := range table.Locals {
		if x == name{
			s := FunctionSymbol{
				Name: name,
			}
			return s
		}
	}
	return nil
}
