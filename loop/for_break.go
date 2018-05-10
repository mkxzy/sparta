package loop

type ForBreak struct {
	error
}

func (fb *ForBreak) Error() string {
	return "break"
}

func Break()  {
	err := &ForBreak{}
	panic(err)
}
