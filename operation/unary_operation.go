package operation

import "github.com/mkxzy/sparta/types"

/**
取负数
 */
func Minus(v types.SPAValue) (result types.SPAValue, ok bool) {
	switch v.(type) {
	case types.SPANumber:
		real := v.(types.SPANumber)
		result = -real
		ok = true
		return
	case types.SPAInteger:
		real := v.(types.SPAInteger)
		result = -real
		ok = true
		return
	}
	result = types.Null()
	ok = false
	return
}
