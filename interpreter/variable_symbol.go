/**
变量符号
 */
package interpreter

import (
	"github.com/mkxzy/sparta/base"
)

type VariableSymbol struct {
	Name string
	Value base.SPAValue
}

func NewVariable(name string, value base.SPAValue) VariableSymbol {
	return VariableSymbol{
		Name: name,
		Value: value,
	}
}

// Symbol接口实现
func (ms  VariableSymbol) GetName() string {
	return ms.Name
}

func (ms  VariableSymbol) GetValue() base.SPAValue  {
	return ms.GetValue()
}
