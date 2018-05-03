package interpreter

import "github.com/mkxzy/sparta/parser"

// 变量类型
const (
	Scalar      = iota
	Array
	//MapAccess
)

type SPALeftSideInterpreter struct {
	ast     parser.ILeft_sideContext
	varName string //变量名称
	varType int    //赋值类型
}

// 实现解释接口
func(v *SPALeftSideInterpreter) Interpret()  {
}
