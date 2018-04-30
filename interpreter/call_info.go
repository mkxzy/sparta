package interpreter

import (
	"bytes"
	"fmt"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

/**
函数调用信息
 */
type CallInfo struct {
	function.SPAFunction 				//函数定义
	locals  map[string]symbol.Symbol 	//局部变量
	upvals  map[string]symbol.Symbol	//自由变量（闭包）
}

func NewCallInfo(fs *FunState)  *CallInfo {
	ci := &CallInfo{
		SPAFunction: fs.Function,
		locals: make(map[string]symbol.Symbol),
		upvals: make(map[string]symbol.Symbol),
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

// begein Scope接口实现
func (ms *CallInfo) GetScopeName() string  {
	return ms.Name
}

func (ms *CallInfo) GetEnclosingScope() symbol.Scope {
	return ms.Outer
}

func (ms *CallInfo) Define(symbol symbol.Symbol) {
	ms.locals[symbol.GetName()] = symbol
}

// end Scope接口实现

func (ci *CallInfo) Resolve(name string) symbol.Symbol {

	sym := ci.locals[name]
	if sym != nil {
		return sym
	}
	sym = ci.upvals[name]
	if sym != nil {
		return sym
	}
	if ci.Outer != nil{
		sym = ci.Outer.Resolve(name)
		if sym != nil{
			return sym
		}
	}
	return nil
}
// Scope接口实现 end

func(ci *CallInfo) String() string  {
	var buffer bytes.Buffer
	for k, v := range ci.locals {
		s := fmt.Sprintf("%s: %v, ", k, v)
		buffer.WriteString(s)
	}
	return buffer.String()
}