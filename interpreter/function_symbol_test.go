package interpreter

import (
	"testing"
	"reflect"
)

func TestType(t *testing.T) {
	var symbol Symbol = &VariableSymbol{}
	switch symbol.(type) {
	case *FunctionSymbol:
		t.Log("函数")
	case *VariableSymbol:
		t.Log("变量")
	}
	t.Log(reflect.TypeOf(symbol))
}
