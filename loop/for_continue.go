package loop

type ForContinue struct {
	error
}

func (fc *ForContinue) Error() string {
	return "continue"
}

func Continue()  {
	err := &ForContinue{}
	panic(err)
}
