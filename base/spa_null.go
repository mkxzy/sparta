package base

type nullValue struct {
}

var n = &nullValue{}

func Null() SPAValue {
	return n
}

func IsNull(value SPAValue) bool {
	return value == n
}
