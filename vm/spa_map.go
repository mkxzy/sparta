package vm

type SPAMap struct {
	mp map[SPAValue]SPAValue
}

func(s SPAMap) IsTrue() bool  {
	return len(s.mp) > 0
}
