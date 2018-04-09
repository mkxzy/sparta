package interpreter

import (
	"fmt"
	"bytes"
)

/**
符号
 */
type Symbol struct {
	Name string
	Type SymbolType
}

/**
符号类型
 */
type SymbolType interface {
	getName() string
}

/**
符号表
 */
type SymbolTable struct {
	Symbols map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	table := &SymbolTable{
		Symbols: make(map[string]Symbol),
	}
	return table
}

func (table *SymbolTable) GetScopeName() string  {
	return "global"
}

func (table *SymbolTable) GetEnclosingScope() *Scope {
	return nil
}

func (table *SymbolTable) Define(symbol Symbol) {
	table.Symbols[symbol.Name] = symbol
}

func (table *SymbolTable) Resolve(name string) Symbol {
	return table.Symbols[name]
}

func (table *SymbolTable) String() string {
	var buffer bytes.Buffer
	for k, v := range table.Symbols {
		entryString := fmt.Sprintf("%s: %v\n", k, v)
		buffer.WriteString(entryString)
	}
	return buffer.String()
}