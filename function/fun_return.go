package function

type FunReturn struct {
	error
}

func (this *FunReturn) Error() string {
	return "函数返回"
}
