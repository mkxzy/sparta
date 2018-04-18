package base

type SPAString string

func(s SPAString) IsTrue() bool  {
	return len(s) > 0
}