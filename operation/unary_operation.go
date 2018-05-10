package operation

import "github.com/mkxzy/sparta/types"

/**
取负数
 */
func Minus() {
	v := PopValue()
	switch v.(type) {
	case types.SPANumber:
		real := v.(types.SPANumber)
		v = -real
	case types.SPAInteger:
		real := v.(types.SPAInteger)
		v = -real
	}
	PushValue(v)
}
