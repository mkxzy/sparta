package base

/**
函数定义
 */
type SPAFunction struct {
	args []string //参数名
}

func NewFunction(args []string) *SPAFunction  {
	return &SPAFunction{
		args: args,
	}
}

func(f *SPAFunction) getArgs() int {
	return len(f.args)
}

func(s SPAFunction) IsTrue() bool  {
	return true
}