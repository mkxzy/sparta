package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/types"
	"fmt"
)

type SPAFuncallExprInterpreter struct {
	ast *parser.Funcall_exprContext
	fs 	*FunState
}

// 实现解释接口
func(v *SPAFuncallExprInterpreter) Interpret()  {
	//fs := &FunState{}
	v.fs = &FunState{}
	v.EvalFunName(v.ast.GetChild(0).(*parser.Fun_nameContext)) 		//获取函数名
	v.EvalFunArgs(v.ast.GetChild(1).(*parser.Arg_exprContext)) 		//参数入栈
	va, ok := state.currentScope.Resolve(v.fs.FunName).(*symbol.SPAVariable)  //获取函数定义
	//log.Debug(v.fs.FunName)
	//log.Debug(ok)
	//log.Debug(currentScope)
	if !ok {
		panic("未找到函数定义")
	}
	f, ok := va.Value.(function.SPAFunction)
	if !ok {
		panic("未找到函数定义")
	}
	v.fs.Function = f
	v.CallFunc()
}

// 获取函数名
func(v *SPAFuncallExprInterpreter) EvalFunName(ctx *parser.Fun_nameContext) {
	v.fs.FunName = ctx.GetChild(0).(antlr.TerminalNode).GetText()

}

// 加载参数到操作数栈（顺序）
func(v *SPAFuncallExprInterpreter) EvalFunArgs(ctx *parser.Arg_exprContext) {
	v.fs.ArgCount = 0
	if ctx.GetChildCount() == 2{
		return
	}
	argListContext := ctx.GetChild(1).(*parser.Arg_listContext)
	for i := 0; i < argListContext.GetChildCount(); i += 2 {
		v.EvalArgument(argListContext.GetChild(i).(*parser.ArgContext))
		v.fs.ArgCount++
	}
}

// 把参数放入操作树栈
func (v *SPAFuncallExprInterpreter) EvalArgument(ctx *parser.ArgContext) {
	//v.EvalTest(ctx.GetChild(0).(*parser.TestContext))
	testInter := &SPATestInterpreter{ctx.GetChild(0).(*parser.TestContext)}
	testInter.Interpret()
}

// 调用函数
func (v *SPAFuncallExprInterpreter) CallFunc() {

	// 内置函数
	if v.fs.Function.Internal {
		switch v.fs.FunName {
		case "print":
			var arg types.SPAValue
			for i := v.fs.ArgCount; i > 0; i--{
				arg = PopValue()
			}
			fmt.Println(arg)
			PushNullValue() //函数体内没有返回语句，自动插入空返回值
		}
	} else {
		v.loadArgs()
		ci := NewCallInfo(v.fs) //创建函数调用信息
		state.ChangeScope(ci)

		defer state.RestoreScope()

		//去掉大括号
		blockCtx := v.fs.Function.Body.GetChild(1)
		//v.ExecBlock(bockCtx.(*parser.BlockContext), v.fs)
		blockInter := &SPABlockInterpreter{
			ast: blockCtx.(*parser.BlockContext),
			ff: v.fs,
		}
		blockInter.Interpret()
		if v.fs.State != RETURN {
			PushNullValue()
		}
	}
}

func (v *SPAFuncallExprInterpreter) loadArgs() {
	v.fs.Args = make([]types.SPAValue, v.fs.ArgCount, v.fs.ArgCount)
	for i := v.fs.ArgCount - 1; i >= 0; i-- {
		value := PopValue()
		v.fs.Args[i] = value
	}
}