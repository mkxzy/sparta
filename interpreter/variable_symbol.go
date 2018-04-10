/**
变量符号
 */
package interpreter

type VariableSymbol struct {
	Name string
}

// Symbol接口实现
func (ms *VariableSymbol) GetName() string {
	return ms.Name
}
