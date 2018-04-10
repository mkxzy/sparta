package interpreter

import "testing"

func TestLocalScope_Resolve(t *testing.T) {
	globalScope := NewGlobalScope()
	globalScope.Define(NewFuncionSymbol("main"))

	localScope := NewLocalScope(globalScope)
	localScope.Define(NewFuncionSymbol("print"))

	const f = "main"
	if localScope.Resolve(f) == nil{
		t.Errorf("找不到变量%s", f)
	}else{
		t.Log(localScope.Resolve(f))
	}

	const f2  = "xxx"
	if localScope.Resolve(f2) == nil{
		t.Log(localScope.Resolve(f2))
	}
}
