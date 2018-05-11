package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/types"
	"fmt"
	"github.com/mkxzy/sparta/operation"
)

type SPAFuncallExprInterpreter struct {
	ast *parser.Funcall_exprContext
}

// 实现解释接口
func(v *SPAFuncallExprInterpreter) Interpret(state *ProgramState)  {

	fs := &function.FunState{
		Locals: make(map[string]symbol.Symbol),
	}

	evalFunName(v.ast.GetChild(0).(*parser.Fun_nameContext), fs) 			//获取函数名
	log.Notice(fs.FunVar)

	evalFunArgs(v.ast.GetChild(1).(*parser.Arg_exprContext), state, fs) 	//参数入栈

	va, ok := state.Resolve(fs.FunVar).(*symbol.SPAVariable) 				//获取函数定义
	if !ok {
		panic("未找到函数定义")
	}
	f, ok := va.Value.(function.SPAFunction)
	if !ok {
		panic("未找到函数定义")
	}
	fs.Function = &f

	savedState := state.GetCurrent()
	state.SetCurrent(fs)

	defer callReturn(savedState, state)

	call(state)
	noReturn(savedState, state)
}

//函数有返回值的情况处理
func callReturn(savedState * function.FunState, state *ProgramState) {
	if err := recover(); err != nil{
		switch err.(type) {
		case *function.FunReturn:
			state.SetCurrent(savedState) //交还控制权给调用者
		default:
			panic(err)
		}
	}
}

//函数没有返回值的情况处理
func noReturn(savedState * function.FunState, state *ProgramState)  {
	operation.PushNullValue()    //没有返回值的情况下插入空值
	state.SetCurrent(savedState) //交还控制权给调用者
}

// 获取函数名
func evalFunName(ctx *parser.Fun_nameContext, fs *function.FunState) {
	funVar := ctx.GetChild(0).(antlr.TerminalNode).GetText()
	//log.Debugf("name: %s", funVar)
	fs.FunVar = funVar
}

// 加载参数到操作数栈（顺序）
func evalFunArgs(ctx *parser.Arg_exprContext, state *ProgramState, fs *function.FunState) {
	fs.ArgCount = 0
	if ctx.GetChildCount() == 2{
		return
	}
	argListContext := ctx.GetChild(1).(*parser.Arg_listContext)
	for i := 0; i < argListContext.GetChildCount(); i += 2 {
		evalArgument(argListContext.GetChild(i).(*parser.ArgContext), state)
		fs.ArgCount++
	}
}

// 把参数放入操作树栈
func evalArgument(ctx *parser.ArgContext, state *ProgramState) {

	testInter := &SPATestInterpreter{ctx.GetChild(0).(*parser.TestContext)}
	testInter.Interpret(state)
}

// 调用函数
func call(state *ProgramState) {

	fs := state.GetCurrent()
	// 内置函数
	if fs.Function.Internal {
		switch fs.FunVar {
		case "print":
			var arg types.SPAValue
			for i := fs.ArgCount; i > 0; i--{
				arg = operation.PopValue()
			}
			fmt.Println(arg)
		}
	} else {

		// 传递参数
		fs.Args = make([]types.SPAValue, fs.ArgCount, fs.ArgCount)
		for i := fs.ArgCount - 1; i >= 0; i-- {
			value := operation.PopValue()
			fs.Args[i] = value
		}
		for i := 0; i < len(fs.Function.Args); i++{
			if i < len(fs.Args){
				fs.Define(symbol.NewVariable(fs.Function.Args[i], fs.Args[i]))
			} else {
				fs.Define(symbol.NewNullVariable(fs.Function.Args[i]))
			}
		}

		//去掉大括号
		blockCtx := fs.Function.Body.GetChild(1)
		blockInter := &SPABlockInterpreter{
			ast: blockCtx.(*parser.BlockContext),
		}
		blockInter.Interpret(state)
	}
}