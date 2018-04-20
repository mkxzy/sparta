package base

/**
变量符号
 */
type VariableSymbol struct {
	Name string
	Value SPAValue
}

func NewVariable(name string, value SPAValue) VariableSymbol {
	return VariableSymbol{
		Name: name,
		Value: value,
	}
}

// Symbol接口实现
func (ms  VariableSymbol) GetName() string {
	return ms.Name
}

//符号关联的值
func (ms  VariableSymbol) GetValue() SPAValue  {
	return ms.GetValue()
}
