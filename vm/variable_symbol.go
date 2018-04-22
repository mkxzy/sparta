package vm

/**
变量符号
 */
type VariableSymbol struct {
	Name string
	Value SPAValue
}

func NewVariable(name string, value SPAValue) *VariableSymbol {
	return &VariableSymbol{
		Name: name,
		Value: value,
	}
}

func NewFunVariable(f SPAFunction) *VariableSymbol {
	return &VariableSymbol{
		Name: f.Name,
		Value: f,
	}
}

// Symbol接口实现
func (ms  *VariableSymbol) GetName() string {
	return ms.Name
}