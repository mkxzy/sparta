package base

/**
表示语言的bool类型
 */
type SPABool struct {
	b bool
}

func NewBoolValue(literal string) SPABool {
	switch literal {
	case "true":
		return SPABool{true}
	case "false":
		return SPABool{false}
	default:
		panic("表达式无效")
	}
}

func(s SPABool) IsTrue() bool  {
	return s.b
}
