package vm

import (
	"bytes"
	"fmt"
)

/**
函数调用状态
 */
type CallInfo struct {
	SPAFunction //函数定义
	*MemorySpace //函数内存空间
}

func NewCallInfo(f SPAFunction, args []SPAValue)  *CallInfo {
	ci := &CallInfo{
		SPAFunction: f,
		MemorySpace: NewMemorySpace(""),
	}
	for i := 0; i < len(f.Args); i++{
		if i < len(args){
			ci.Define(NewVariable(f.Args[i], args[i]))
		} else {
			ci.Define(NewNullVariable(f.Args[i]))
		}
	}
	return ci
}

func (ci *CallInfo) Resolve(name string) Symbol {

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