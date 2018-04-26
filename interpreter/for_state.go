package interpreter

/**
循环状态
 */
type ForState struct {
	State int
	ItemName string
}

func (fs *ForState) GetState() int {
	return fs.State
}

func (fs *ForState) SetState(s int) {
	fs.State = s
}
