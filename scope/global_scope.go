// 符号表， 管理程序符号实体

package scope

import (
	"bytes"
	"fmt"
)

/**
符号表
 */
type GlobalScope struct {
	Symbols map[string]Symbol
}

func NewGlobalScope() *GlobalScope {
	table := &GlobalScope{
		Symbols: make(map[string]Symbol),
	}
	return table
}

// begein Scope接口实现
func (table *GlobalScope) GetScopeName() string  {
	return "global"
}

func (table *GlobalScope) GetEnclosingScope() Scope {
	return nil
}

func (table *GlobalScope) Define(symbol Symbol) {
	table.Symbols[symbol.GetName()] = symbol
}

func (table *GlobalScope) Resolve(name string) Symbol {
	return table.Symbols[name]
}
// end Scope接口实现

// Stringer接口实现
func (table *GlobalScope) String() string {
	var buffer bytes.Buffer
	for k, v := range table.Symbols {
		entryString := fmt.Sprintf("%s: %v\n", k, v)
		buffer.WriteString(entryString)
	}
	return buffer.String()
}
