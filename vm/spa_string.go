package vm

type SPAString struct {
	str string
}

func(s SPAString) Original() interface{}{
	return s.str
}

func(s SPAString) String() string {
	return s.str
}

func NewString(literal string) *SPAString {
	spa := SPAString{
		str: escapeString(literal),
	}
	return &spa
}

/**
转义字符串（状态机）
 */
func escapeString(text string) string {
	var result = make([]rune, 0 , len(text))
	escaped := false //是否处于转义状态
	for _, c := range text {
		if escaped {
			result = append(result, escapeMapping(c))
			escaped = false //关闭转义状态
		}else{
			switch c {
			case '"':
				continue
			case '\\':
				escaped = true
			default:
				result = append(result, c)
			}
		}
	}
	s := string(result)
	return s
}

/**
转义字符映射
 */
func escapeMapping(c rune) rune {
	switch c {
	case '"', '\\':
		return c
	case 't':
		return '\t'
	case 'r':
		return '\r'
	case 'n':
		return '\n'
	default:
		return c
	}
}

func(s SPAString) IsTrue() bool  {
	return len(s.str) > 0
}