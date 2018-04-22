package vm

import (
	"github.com/mkxzy/sparta/parser"
	"strings"
)

/**
函数定义
 */
type SPAFunction struct {

	Internal bool	//是否内置函数
	Name string 	//函数名
	Args []string 	//参数名
	Outer Scope 	//外部作用域
	Body *parser.Fun_bodyContext //函数体（解析树）
}

func NewFunction(name string, args []string, body *parser.Fun_bodyContext) SPAFunction {
	return SPAFunction{
		Internal: false,
		Name: name,
		Args: args,
		Body: body,
	}
}

func NewInternalFunction(name string, args []string) SPAFunction  {
	return SPAFunction{
		Internal: true,
		Name: name,
		Args: args,
	}
}

// Booler接口
func(s SPAFunction) IsTrue() bool  {
	return true
}

//Stringer接口实现
func(f SPAFunction) String() string {
	return f.Name + "(" + strings.Join(f.Args, ", ") + ")"
}