package interpreter

import "github.com/mkxzy/sparta/vm"

type FunState struct {
	FunName  string
	ArgCount int
	State    int
	Function vm.SPAFunction
}

func (fs *FunState) GetState() int {
	return fs.State
}

func (fs *FunState) SetState(s int) {
	fs.State = s
}
