package interpreter

import "github.com/mkxzy/sparta/symbol"

type ProgramState struct {
	State int
	currentScope symbol.Scope
	savedScope symbol.Scope
}

func (fs *ProgramState) GetState() int {
	return fs.State
}

func (fs *ProgramState) SetState(s int) {
	fs.State = s
}

func (fs *ProgramState) ChangeScope(s symbol.Scope)  {
	fs.savedScope = fs.currentScope
	fs.currentScope = s
}

func (fs *ProgramState) RestoreScope()  {
	fs.currentScope = fs.savedScope
}
