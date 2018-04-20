package base

/**
函数调用状态
 */
type CallInfo struct {
	*SPAFunction //函数定义
	*MemorySpace //函数内存空间
}

func NewCallInfo(f *SPAFunction)  *CallInfo {
	return &CallInfo{
		SPAFunction: f,
		MemorySpace: NewMemorySpace(""),
	}
}

func(ci *CallInfo) PushArgs(args []SPAValue)  {
	for i := 0; i < len(ci.Args); i++{
		//ci.Locals = append(ci.Locals, NewVariable(ci.Args[i], args[i]))
		ci.Define(NewVariable(ci.Args[i], args[i]))
		if i >= len(args) {
			//ci.Locals = append(ci.Locals, NewVariable(ci.Args[i], base.Null())) //不足的参数用NULL来补
			ci.Define(NewVariable(ci.Args[i], Null()))
		}
	}
}

func (ci *CallInfo) Resolve(name string) Symbol {

	sym := ci.MemorySpace.Resolve(name)
	if sym != nil {
		return sym
	}
	if ci.Outer != nil {
		sym := ci.Outer.Resolve(name)
		if sym != nil{
			return sym
		}
	}
	return nil
}
// Scope接口实现 end