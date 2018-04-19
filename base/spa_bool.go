package base

/**
表示语言的bool类型
 */
type SPABool struct {
	b bool
}

func NewBool(literal string) SPABool {
	switch literal {
	case "true":
		return SPABool{true}
	case "false":
		return SPABool{false}
	default:
		panic("表达式无效")
	}
}

// 返回值对应的相反bool值
func InvertBool(v SPAValue) SPABool {
	return SPABool{!v.IsTrue()}
}

func(s SPABool) IsTrue() bool  {
	return s.b
}

func(s SPABool) Add(v SPAValue) SPAValue {
	switch v.(type) {
	case SPANumber:
		number := v.(SPANumber)
		return number + toNumber(s)
	case SPABool:
		b2 := v.(SPABool)
		return toNumber(s) + toNumber(b2)
	default:
		panic("操作不支持")
	}
}

func toNumber(spaBool SPABool) SPANumber {
	if spaBool.IsTrue(){
		return SPANumber(1)
	} else{
		return SPANumber(0)
	}
}
