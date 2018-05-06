package interpreter

import (
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/function"
	"github.com/mkxzy/sparta/symbol"
)

type FunState struct {
	FunName  string						//函数名
	ArgCount int						//实际传递参数数量
	State    int						//函数调用状态
	Function function.SPAFunction		//函数定义
	Args []types.SPAValue 				//实际传递参数
	locals  map[string]symbol.Symbol 	//局部变量
	prevState *FunState					//上个函数调用
}

func (fs *FunState) GetState() int {
	return fs.State
}

func (fs *FunState) SetState(s int) {
	fs.State = s
}

// begein Scope接口实现
func (fs *FunState) GetScopeName() string  {
	return fs.Function.Name
}

func (fs *FunState) GetEnclosingScope() symbol.Scope {
	return fs.Function.Outer
}

func (ms *FunState) Define(s symbol.Symbol) {
	ms.locals[s.GetName()] = s
}

func (fs *FunState) Resolve(name string) symbol.Symbol {

	sym := fs.locals[name]
	if sym != nil {
		return sym
	}
	if fs.GetEnclosingScope() != nil{
		sym = fs.GetEnclosingScope().Resolve(name)
		if sym != nil{
			return sym
		}
	}
	return nil
}
// Scope接口实现 end
