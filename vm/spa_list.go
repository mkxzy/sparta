package vm

type SPAList struct {
	arr []SPAValue
}

func(s SPAList) IsTrue() bool  {
	return len(s.arr) > 0
}
