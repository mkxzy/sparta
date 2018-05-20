package types

/**
表示语言的bool类型
 */
type SPABool bool

func NewBool(literal string) SPABool {
	switch literal {
	case "true":
		return SPABool(true)
	case "false":
		return SPABool(false)
	default:
		panic("表达式无效")
	}
}

func True() SPABool {
	return true
}

func False() SPABool {
	return false
}

// 返回值对应的相反bool值
func InvertBool(v SPAValue) SPABool {
	return SPABool(!v.IsTrue())
}

func(s SPABool) IsTrue() bool  {
	return bool(s)
}

func(s SPABool) IsFalse() bool {
	return !s.IsTrue()
}

func(spaBool SPABool) ToInt() SPAInteger {
	if spaBool == true{
		return SPAInteger(1)
	} else{
		return SPAInteger(0)
	}
}

func(n SPABool) String() string  {
	if n.IsTrue(){
		return "true"
	} else {
		return "false"
	}
}
