package interpreter

import "github.com/mkxzy/sparta/vm"

/**
取负数
 */
func minus(v vm.SPAValue) (result vm.SPAValue, ok bool) {
	switch v.(type) {
	case vm.SPANumber:
		real := v.(vm.SPANumber)
		result = -real
		ok = true
		return
	case vm.SPAInteger:
		real := v.(vm.SPAInteger)
		result = -real
		ok = true
		return
	}
	result = vm.Null()
	ok = false
	return
}
