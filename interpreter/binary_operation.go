package interpreter

import "github.com/mkxzy/sparta/vm"

/**
+号运算
 */
func addOp(left, right vm.SPAValue) (result vm.SPAValue, ok bool) {
	switch left.(type) {
	case vm.SPANumber:
		realLeft := left.(vm.SPANumber)
		switch right.(type) {
		case vm.SPANumber:
			realRight := right.(vm.SPANumber)
			result = realLeft + realRight
			ok = true
			return
		case vm.SPAInteger:
			realRight := right.(vm.SPAInteger)
			result = realLeft + vm.SPANumber(realRight)
			ok = true
			return
		}
	}
	result = vm.Null()
	ok = false
	return
}
