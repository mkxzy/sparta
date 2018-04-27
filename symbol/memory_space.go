// 符号表， 管理程序符号实体

package symbol

import (
	"bytes"
	"fmt"
)

/**
变量存储空间
 */
type MemorySpace struct {
	Name string //符号名
	Symbols map[string]Symbol //符号
}

func NewMemorySpace(name string) *MemorySpace {
	table := &MemorySpace{
		Name: name,
		Symbols: make(map[string]Symbol),
	}
	return table
}

// begein Scope接口实现
func (ms *MemorySpace) GetScopeName() string  {
	return ms.Name
}

func (ms *MemorySpace) GetEnclosingScope() Scope {
	return nil
}

func (ms *MemorySpace) Define(symbol Symbol) {
	ms.Symbols[symbol.GetName()] = symbol
}

func (ms *MemorySpace) Resolve(name string) Symbol {
	return ms.Symbols[name]
}
// end Scope接口实现

// Stringer接口实现
func (ms *MemorySpace) String() string {
	var buffer bytes.Buffer
	for k, v := range ms.Symbols {
		entryString := fmt.Sprintf("%s: %v, ", k, v)
		buffer.WriteString(entryString)
	}
	return buffer.String()
}
