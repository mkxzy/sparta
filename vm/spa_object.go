package vm

/**
标记接口
 */
type SPAValue interface {
	Booler
}

/**
值是否为真
 */
type Booler interface {
	IsTrue() bool
}