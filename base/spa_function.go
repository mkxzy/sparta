package base

/**
函数定义
 */
type SPAFunction struct {
	argc int //参数数量（定义的时候需要）
	args []string //参数名
	locals map[string]SPAValue //局部变量
}

func NewFunction(argc int, args []string) *SPAFunction  {
	return &SPAFunction{
		argc: argc,
		args: args,
		locals: make(map[string]SPAValue),
	}
}

func(s SPAFunction) IsTrue() bool  {
	return true
}