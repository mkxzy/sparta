package interpreter

import (
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/function"
)

type FunState struct {
	FunName  string
	ArgCount int		//实际传递参数数量
	State    int
	Function function.SPAFunction
	Args []types.SPAValue //实际传递参数
}

func (fs *FunState) GetState() int {
	return fs.State
}

func (fs *FunState) SetState(s int) {
	fs.State = s
}
