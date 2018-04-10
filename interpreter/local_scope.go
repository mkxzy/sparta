package interpreter

type LocalScope struct {
	Symbols   map[string]Symbol
	Enclosing Scope
}

// begein Scope接口实现
func (table *LocalScope) GetScopeName() string  {
	return "global"
}

func (table *LocalScope) GetEnclosingScope() Scope {
	return table.Enclosing
}

func (table *LocalScope) Define(symbol Symbol) {
	table.Symbols[symbol.GetName()] = symbol
}

func (table *LocalScope) Resolve(name string) Symbol {

	if table == nil {
		return nil
	}
	sym := table.Symbols[name]
	if sym != nil{
		return sym
	}else{
		var s = table.GetEnclosingScope()
		return s.Resolve(name)
	}
}
// end Scope接口实现

func NewLocalScope(enclosing Scope) *LocalScope {
	scope := &LocalScope{
		Symbols: make(map[string]Symbol),
		Enclosing: enclosing,
	}
	return scope
}
