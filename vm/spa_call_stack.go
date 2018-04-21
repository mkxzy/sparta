package vm

var calls [10000]*CallInfo            //栈空间最多10000
var fp = -1                                 //函数栈指针

// 函数调用信息入栈
func PushCallInfo(ci *CallInfo)  {

	fp++
	calls[fp] = ci
}

// 函数调用信息出栈
func PopCallInfo()  {

	if HasCallInfo() {
		calls[fp] = nil
		fp--
	} else {
		panic("函数栈为空")
	}
}

// 获取栈顶函数调用信息
func GetTopCallInfo() *CallInfo  {
	if HasCallInfo() {
		return calls[fp]
	}
	panic("函数栈为空")
}

// 判断是否有函数调用
func HasCallInfo() bool  {
	return fp > -1
}
