// 符号表， 管理程序符号实体

package interpreter

import (
	"bytes"
	"fmt"
)

/**
符号表
 */
type SymbolTable struct {
	Symbols map[string]Symbol
}

// begein Scope接口实现
func (table *SymbolTable) GetScopeName() string  {
	return "global"
}

func (table *SymbolTable) GetEnclosingScope() *Scope {
	return nil
}

func (table *SymbolTable) Define(symbol Symbol) {
	table.Symbols[symbol.GetName()] = symbol
}

func (table *SymbolTable) Resolve(name string) Symbol {
	return table.Symbols[name]
}
// end Scope接口实现

// Stringer接口实现
func (table *SymbolTable) String() string {
	var buffer bytes.Buffer
	for k, v := range table.Symbols {
		entryString := fmt.Sprintf("%s: %v\n", k, v)
		buffer.WriteString(entryString)
	}
	return buffer.String()
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{
		Symbols: make(map[string]Symbol),
	}
	return table
}
