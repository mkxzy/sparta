package base

type SPAFunction struct {

}

func(f *SPAFunction) call(args...SPAValue) SPAValue  {
	return Null()
}
