package function

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
