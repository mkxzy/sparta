package base

type SPANumber float64

func(n SPANumber) IsTrue() bool {
	return n > 0.0
}