package types

import "github.com/op/go-logging"

var log = logging.MustGetLogger("operands")

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