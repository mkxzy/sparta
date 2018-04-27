package symbol

import (
	"github.com/mkxzy/sparta/types"
)

/**
变量符号
 */
type SPAVariable struct {
	Name  string
	Value types.SPAValue
}

func NewVariable(name string, value types.SPAValue) *SPAVariable {
	return &SPAVariable{
		Name: name,
		Value: value,
	}
}

func NewNullVariable(name string) *SPAVariable {
	return &SPAVariable{
		Name:  name,
		Value: types.Null(),
	}
}

// Symbol接口实现
func (ms  *SPAVariable) GetName() string {
	return ms.Name
}
