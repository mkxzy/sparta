package types

import (
	"github.com/op/go-logging"
	"fmt"
)

var log = logging.MustGetLogger("operands")

/**
标记接口
 */
type SPAValue interface {
	Booler
	fmt.Stringer
}

/**
值是否为真
 */
type Booler interface {
	IsTrue() bool
}