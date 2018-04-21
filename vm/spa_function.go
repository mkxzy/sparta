package vm

import (
	"github.com/mkxzy/sparta/parser"
	"strings"
)

/**
函数符号定义
 */
type SPAFunction struct {

	Name string 	//函数名
	Args []string 	//参数名
	Outer Scope 	//外部作用域
	Body *parser.BlockContext //函数体（解析树）
}

func NewFunction(name string, args []string, body *parser.BlockContext) *SPAFunction {
	return &SPAFunction{
		Name: name,
		Args: args,
		Body: body,
	}
}

// Symbol接口实现
func (f *SPAFunction) GetName() string {
	return f.Name
}

//Stringer接口实现
func(f *SPAFunction) String() string {
	return f.Name + "(" + strings.Join(f.Args, ", ") + ")"
}