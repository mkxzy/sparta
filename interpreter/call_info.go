package interpreter

import (
	"bytes"
	"fmt"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

/**
函数调用状态
 */
type CallInfo struct {
	function.SPAFunction //函数定义
	*symbol.MemorySpace //函数内存空间
}

func NewCallInfo(fs *FunState)  *CallInfo {
	ci := &CallInfo{
		SPAFunction: fs.Function,
		MemorySpace: symbol.NewMemorySpace(""),
	}
	for i := 0; i < len(fs.Function.Args); i++{
		if i < len(fs.Args){
			ci.Define(symbol.NewVariable(fs.Function.Args[i], fs.Args[i]))
		} else {
			ci.Define(symbol.NewNullVariable(fs.Function.Args[i]))
		}
	}
	return ci
}

func (ci *CallInfo) Resolve(name string) symbol.Symbol {

	sym := ci.MemorySpace.Resolve(name)
	if sym != nil {
		return sym
	}
	if ci.Outer != nil {
		sym := ci.Outer.Resolve(name)
		if sym != nil{
			return sym
		}
	}
	return nil
}
// Scope接口实现 end

func(ci *CallInfo) String() string  {
	var buffer bytes.Buffer
	for k, v := range ci.Symbols {
		s := fmt.Sprintf("%s: %v, ", k, v)
		buffer.WriteString(s)
	}
	return buffer.String()
}