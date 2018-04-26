package interpreter

type ProgramState struct {
	State int
}

func (fs *ProgramState) GetState() int {
	return fs.State
}

func (fs *ProgramState) SetState(s int) {
	fs.State = s
}
