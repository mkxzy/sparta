package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// 变量类型
const (
	Scalar      = iota
	Array
)

type SPALeftSideInterpreter struct {
	ast     *parser.Left_sideContext
	varName string //变量名称
	varType int    //赋值类型
}

// 实现解释接口
func(v *SPALeftSideInterpreter) Interpret(state *ProgramState)  {

	v.varName = v.ast.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	v.varType = Scalar
	if v.ast.GetChildCount() >= 4 {
		v.varType = Array
		testInter := &SPATestInterpreter{}
		testInter.Interpret(state)
	}
}
