package function

/**
模拟函数返回指令
 */
type FunReturn struct {
	error
}

func (fr *FunReturn) Error() string {
	return "return"
}

func Return()  {
	err := &FunReturn{}
	panic(err)
}