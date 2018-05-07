package function

import (
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/symbol"
)

/**
函数调用状态
 */
type FunState struct {
	FunVar    string                   //函数变量名
	ArgCount  int                      //实际传递参数数量
	State     int                      //函数调用状态
	Function  *SPAFunction             //函数定义
	Args      []types.SPAValue         //实际传递参数
	Locals    map[string]symbol.Symbol //局部变量
	PrevState *FunState                //上个函数调用
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
	return fs.Function.FunList
}

func (ms *FunState) Define(s symbol.Symbol) {
	ms.Locals[s.GetName()] = s
}

//解析变量
func (fs *FunState) Resolve(name string) (s symbol.Symbol) {
	for x := fs; x != nil; x = x.Function.FunList {
		s = x.Locals[name]
		if s != nil {
			break
		}
	}
	return
}
// Scope接口实现 end
