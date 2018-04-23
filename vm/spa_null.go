package vm

/**
表示空类型，受保护类型，外部只能通过IsNull判断是否是空类型
 */
type spaNull struct {
}

var n = &spaNull{}

func Null() *spaNull {
	return n
}

/**
判断任意值是否为空
 */
func IsNull(value SPAValue) bool {
	return value == n
}

/**
空值永远是false
 */
func(s *spaNull) IsTrue() bool  {
	return false
}

func(s *spaNull) String() string {
	return "null"
}