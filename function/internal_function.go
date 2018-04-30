package function

var internals =  make(map[string]InternalFunction)

func init()  {
	internals["print"] = InternalFunction{
		Name: "print",
		Args: []string{"s"},
	}
}

/**
内置函数
 */
type InternalFunction struct {

	Name string 	//函数名
	Args []string 	//参数名
}

// Symbol接口实现
func (ms InternalFunction) GetName() string {
	return ms.Name
}
