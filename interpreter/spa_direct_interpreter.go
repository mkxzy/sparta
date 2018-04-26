/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/vm"
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
)

var log = logging.MustGetLogger("SPADirectInterpreter")

type SPADirectInterpreter struct {
	//*parser.BaseSpartaVisitor
	GlobalState *vm.MemorySpace
}

func NewDirectInterpreter(globalState *vm.MemorySpace) *SPADirectInterpreter {
	return &SPADirectInterpreter{
		GlobalState: globalState,
	}
}

// 实现解释接口
func(v *SPADirectInterpreter) Interpret(ctx parser.IProgramContext)  {
	for i := 0; i < ctx.GetChildCount()-1; i++ {
		stmtContext := ctx.GetChild(i).(*parser.StmtContext)
		v.ExecStmt(stmtContext, &ProgramState{State:NORMAL})
	}
}

/**
执行语句块
返回值： 是否被return中断
 */
func (v *SPADirectInterpreter) ExecBlock(ctx *parser.BlockContext, ff FlowState) {
	for i := 1; i < ctx.GetChildCount()-1; i++ {
		stmtContext := ctx.GetChild(i).(*parser.StmtContext)
		v.ExecStmt(stmtContext, ff)
		// return，continue,break都要中断执行块
		if ff.GetState() > NORMAL {
			break
		}
	}
}

// 执行语句
// 返回值： 是否有返回值
func (v *SPADirectInterpreter) ExecStmt(ctx *parser.StmtContext, ff FlowState) {
	log.Debug("Visit Stmt")

	rule := ctx.GetChild(0).(antlr.RuleContext)
	switch rule.GetRuleIndex() {
	case parser.SpartaParserRULE_assign_stmt:
		v.ExecAssignStmt(rule.(*parser.Assign_stmtContext))
	case parser.SpartaParserRULE_fundef_stmt:
		v.ExecFunDefStmt(rule.(*parser.Fundef_stmtContext))
	case parser.SpartaParserRULE_funcall_stmt:
		v.ExecFunCallStmt(rule.(*parser.Funcall_stmtContext))
	case parser.SpartaParserRULE_return_stmt:
		v.ExecReturnStmt(rule.(*parser.Return_stmtContext), ff)
	case parser.SpartaParserRULE_if_stmt:
		v.ExecIfStmt(rule.(*parser.If_stmtContext), ff)
	case parser.SpartaParserRULE_for_stmt:
		v.ExecForStmt(rule.(*parser.For_stmtContext), ff)
	case parser.SpartaParserRULE_continue_stmt:
		v.ExecContinueStmt(rule.(*parser.Continue_stmtContext), ff)
	case parser.SpartaParserRULE_break_stmt:
		v.ExecBreakStmt(rule.(*parser.Break_stmtContext), ff)
	default:
		panic("不支持的语句")
	}
}

/**
赋值语句
 */
func (v *SPADirectInterpreter) ExecAssignStmt(ctx *parser.Assign_stmtContext)  {

	v.EvalTest(ctx.GetChild(2).(*parser.TestContext))
	name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	value := vm.PopValue()
	sym := vm.NewVariable(name, value)

	if vm.HasCallInfo(){
		vm.GetTopCallInfo().Define(sym) //局部变量定义
		//log.Infof("局部变量定义： %v", vm.GetTopCallInfo())
	}else{
		v.GlobalState.Define(sym) 		//全局变量定义
		//log.Infof("全局变量定义: %v", v.GlobalState)
	}
}

/**
定义函数表达式
 */
func (v *SPADirectInterpreter) ExecFunDefStmt(ctx *parser.Fundef_stmtContext)  {
	f := vm.SPAFunction{}
	f.Name = ctx.GetChild(1).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	f.Body = ctx.GetChild(2).(*parser.Fun_bodyContext)
	v.EvalParList(f.Body.GetChild(0).(*parser.Fun_parContext), &f)
	f.Outer = v.GlobalState
	sym := vm.NewFunVariable(f)
	v.GlobalState.Define(sym) //函数定义
	log.Infof("函数定义: %v", v.GlobalState)
}

/**
获取形参列表
 */
func (v *SPADirectInterpreter) EvalParList(ctx *parser.Fun_parContext, f *vm.SPAFunction) {

	f.Args = []string{}
	if ctx.GetChildCount() == 2{
		return
	}
	context := ctx.GetChild(1).(*parser.NamelistContext)
	//var parNames = make([]string, 0, 10)
	for i := 0; i < context.GetChildCount(); i += 2{
		name := context.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
		f.Args = append(f.Args, name)
	}
}

/**
函数返回表达式
 */
func (v *SPADirectInterpreter) ExecReturnStmt(ctx *parser.Return_stmtContext, ff FlowState) {

	ff.SetState(RETURN)
	if ctx.GetChildCount() == 2 {
		v.EvalTest(ctx.GetChild(1).(*parser.TestContext))
	} else{
		vm.PushValue(vm.Null()) //如果return没有参数，那么插入一个空的返回值
	}
}

/**
函数调用语句
 */
func (v *SPADirectInterpreter) ExecFunCallStmt(ctx *parser.Funcall_stmtContext)  {
	v.EvalFunCallExpr(ctx.GetChild(0).(*parser.Funcall_exprContext))
	vm.PopValue() //丢弃返回值
}

/**
if语句
 */
func (v *SPADirectInterpreter) ExecIfStmt(ctx *parser.If_stmtContext, ff FlowState) {
	v.EvalTest(ctx.GetChild(1).(*parser.TestContext))
	testResult := vm.PopValue()
	if testResult.IsTrue() {
		v.ExecBlock(ctx.GetChild(2).(*parser.BlockContext), ff)
	}
	pos := 3
	for pos + 4 < ctx.GetChildCount() {
		v.EvalTest(ctx.GetChild(pos + 2).(*parser.TestContext))
		testResult = vm.PopValue()
		if testResult.IsTrue() {
			v.ExecBlock(ctx.GetChild(pos + 3).(*parser.BlockContext), ff)
		}
		pos += 4
	}
	if pos < ctx.GetChildCount() {
		v.ExecBlock(ctx.GetChild(pos + 1).(*parser.BlockContext), ff)
	}
}

/**
右值表达式
 */
func (v *SPADirectInterpreter) EvalTest(ctx *parser.TestContext) {
	log.Debug("Visit Postfix Expr")

	v.EvalCompareExpr(ctx.GetChild(0).(*parser.Compare_exprContext))
}

func (v *SPADirectInterpreter) EvalCompareExpr(ctx *parser.Compare_exprContext) {
	v.EvalArithExpr(ctx.GetChild(0).(*parser.Arith_exprContext))
	if ctx.GetChildCount() == 3 {
		v.EvalArithExpr(ctx.GetChild(2).(*parser.Arith_exprContext))
		op := ctx.GetChild(1).(*parser.Comp_opContext).GetChild(0).(antlr.TerminalNode).GetText()
		switch op {
		case "==":
			second := vm.PopValue().(vm.SPAInteger)
			first := vm.PopValue().(vm.SPAInteger)
			result := first == second
			vm.PushValue(vm.SPABool(result))
		default:
			panic("不支持的操作")
		}
	}
}

/**
计算算数表达式
 */
func (v *SPADirectInterpreter) EvalArithExpr(ctx *parser.Arith_exprContext) {
	log.Debug("计算加减法")

	v.EvalTerm(ctx.GetChild(0).(*parser.TermContext))
	if ctx.GetChildCount() > 1 {
		var op = ""
		for i := 1; i < ctx.GetChildCount(); i++ {
			if i % 2 == 1 {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			} else {
				v.EvalTerm(ctx.GetChild(i).(*parser.TermContext))
				arithmetic(op)
			}
		}
	}
}

/**
计算算数表达式
 */
func (v *SPADirectInterpreter) EvalTerm(ctx *parser.TermContext) {
	log.Debug("计算乘除法")
	
	v.EvalFactor(ctx.GetChild(0).(*parser.FactorContext))
	if ctx.GetChildCount() > 1 {
		var op = ""
		for i := 1; i < ctx.GetChildCount(); i++ {
			if i % 2 == 1 {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			} else {
				v.EvalFactor(ctx.GetChild(i).(*parser.FactorContext))
				arithmetic(op)
			}
		}
	}
}

/**
计算因子
 */
func (v *SPADirectInterpreter) EvalFactor(ctx *parser.FactorContext) {
	log.Debug("Visit Factor")

	if ctx.GetChildCount() == 1{
		v.EvalAtomExpr(ctx.GetChild(0).(*parser.Atom_exprContext))
	} else{
		v.EvalAtomExpr(ctx.GetChild(1).(*parser.Atom_exprContext))
		//取反
		first := vm.PopValue()
		result, _ := minus(first)
		vm.PushValue(result)
	}
}

/**
原子表达式
 */
func (v *SPADirectInterpreter) EvalAtomExpr(ctx *parser.Atom_exprContext) {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		// 括号优先表达式
		v.EvalTest(ctx.GetChild(1).(*parser.TestContext))
	} else {
		funCallExpr, ok := ctx.GetChild(0).(*parser.Funcall_exprContext)
		// 函数调用
		if ok {
			v.EvalFunCallExpr(funCallExpr)
		} else {
			//终结符
			terminalNode := ctx.GetChild(0).(antlr.TerminalNode)
			tt := terminalNode.GetSymbol().GetTokenType()
			switch tt {
			case parser.SpartaLexerIDENTIFIER:
				var value vm.SPAValue
				if vm.HasCallInfo(){
					value = vm.GetTopCallInfo().Resolve(terminalNode.GetText()).(*vm.VariableSymbol).Value
				} else {
					value = v.GlobalState.Resolve(terminalNode.GetText()).(*vm.VariableSymbol).Value
				}
				vm.PushValue(value)
			case parser.SpartaLexerNUMBER_LITERAL:
				vm.PushValue(vm.NewNumber(terminalNode.GetText()))
			case parser.SpartaLexerINTEGER_LITERAL:
				vm.PushValue(vm.NewInteger(terminalNode.GetText()))
			case parser.SpartaLexerSTRING:
				vm.PushValue(vm.NewString(terminalNode.GetText()))
			default:
				panic("类型无效")
			}
		}
	}
}

/**
执行函数调用表达式
 */
func(v *SPADirectInterpreter) EvalFunCallExpr(funCallExpr *parser.Funcall_exprContext)  {
	fs := &FunState{}
	v.EvalFunName(funCallExpr.GetChild(0).(*parser.Fun_nameContext), fs) //获取函数名
	v.EvalFunArgs(funCallExpr.GetChild(1).(*parser.Arg_exprContext), fs) //参数入栈
	va, ok := v.GlobalState.Resolve(fs.FunName).(*vm.VariableSymbol)     //获取函数定义
	if !ok {
		panic("未找到函数定义")
	}
	f, ok := va.Value.(vm.SPAFunction)
	if !ok {
		panic("未找到函数定义")
	}
	fs.Function = f
	v.CallFunc(fs)
}

// 获取函数名
func(v *SPADirectInterpreter) EvalFunName(ctx *parser.Fun_nameContext, state *FunState) {
	state.FunName = ctx.GetChild(0).(antlr.TerminalNode).GetText()
}

// 保存参数到操作数栈（顺序）
func(v *SPADirectInterpreter) EvalFunArgs(ctx *parser.Arg_exprContext, state *FunState) {
	state.ArgCount = 0
	if ctx.GetChildCount() == 2{
		return
	}
	argListContext := ctx.GetChild(1).(*parser.Arg_listContext)
	for i := 0; i < argListContext.GetChildCount(); i += 2 {
		v.EvalArgument(argListContext.GetChild(i).(*parser.ArgContext))
		state.ArgCount++
	}
}

// 把参数放入操作树栈
func (v *SPADirectInterpreter) EvalArgument(ctx *parser.ArgContext) {
	v.EvalTest(ctx.GetChild(0).(*parser.TestContext))
}

// 调用函数
func (v *SPADirectInterpreter) CallFunc(fs *FunState) {

	// 内置函数
	if fs.Function.Internal {
		switch fs.FunName {
		case "print":
			var arg vm.SPAValue
			for i := fs.ArgCount; i > 0; i--{
				arg = vm.PopValue()
			}
			fmt.Println(arg)
			vm.PushNullValue() //函数体内没有返回语句，自动插入空返回值
		}
	} else {
		argList := FetchArgs(fs.ArgCount)
		ci := vm.NewCallInfo(fs.Function, argList) //创建函数调用信息
		//passArgs(ci, fs.ArgCount)		//传递参数
		vm.PushCallInfo(ci)    	//保存到函数调用栈

		defer vm.PopCallInfo() //函数退出时弹出调用栈

		//var ret bool
		//去掉大括号
		bockCtx := fs.Function.Body.GetChild(1)
		v.ExecBlock(bockCtx.(*parser.BlockContext), fs)
		if fs.State != RETURN {
			vm.PushNullValue()
		}
	}
}

func FetchArgs(argCount int) []vm.SPAValue {
	argList := make([]vm.SPAValue, argCount, argCount)
	for i := argCount - 1; i >= 0; i-- {
		v := vm.PopValue()
		argList[i] = v
	}
	return argList
}

//// 参数压栈的时候是顺序的，因此遍历要倒序
//func passArgs(ci *vm.CallInfo, args int)  {
//
//	// 实参多于或等于形参
//	if args >= len(ci.Args){
//		for i := args - 1; i >= 0; i-- {
//			v := vm.PopValue()
//			//参数对齐，多余参数丢弃
//			if i < len(ci.Args) {
//				ci.Define(vm.NewVariable(ci.Args[i], v))
//			}
//		}
//	} else {
//		// 实参少于形参 args > len(ci.ArgCount)
//		for i := len(ci.Args) - 1; i >= 0; i-- {
//
//			if i < args {
//				v := vm.PopValue()
//				ci.Define(vm.NewVariable(ci.Args[i], v))
//			} else {
//				v := vm.Null()
//				ci.Define(vm.NewVariable(ci.Args[i], v))
//			}
//		}
//	}
//}

// 循环体
func (v *SPADirectInterpreter) ExecForStmt(ctx *parser.For_stmtContext, ff FlowState) {

	forState := &ForState{State:NORMAL}
	forState.ItemName = ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	sym := vm.NewVariable(forState.ItemName, vm.Null())

	if vm.HasCallInfo(){
		vm.GetTopCallInfo().Define(sym) //局部变量定义
		//log.Infof("局部变量定义： %v", vm.GetTopCallInfo())
	}else{
		v.GlobalState.Define(sym) 		//全局变量定义
		//log.Infof("全局变量定义: %v", v.GlobalState)
	}
	v.EvalTest(ctx.GetChild(3).(*parser.TestContext))
	v.EvalTest(ctx.GetChild(5).(*parser.TestContext))
	toNumber, ok := vm.PopValue().(vm.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	fromNumber, ok := vm.PopValue().(vm.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	if fromNumber > toNumber{
		panic("起始数字不能大于结束数字")
	}

	for i := fromNumber; i <= toNumber; i++ {
		sym.Value = vm.SPAInteger(i)
		v.ExecBlock(ctx.GetChild(6).(*parser.BlockContext), forState)
		if forState.State == BREAK {
			break
		}
		if forState.State == CONTINUE {
			continue
		}
		if forState.State == RETURN {
			ff.SetState(RETURN) 		//传递给调用者
			return
		}
	}
}

/**
循环继续
 */
func (v *SPADirectInterpreter) ExecContinueStmt(ctx *parser.Continue_stmtContext, state FlowState) {
	state.SetState(CONTINUE)
}


func (v *SPADirectInterpreter) ExecBreakStmt(ctx *parser.Break_stmtContext, state FlowState) {
	state.SetState(BREAK)
}

/**
算数运算
 */
func arithmetic(op string) {
	second := vm.PopValue()
	first := vm.PopValue()
	operation := operations[op]
	if operation == nil{
		panic("不支持的操作")
	}
	result, _ := operation(first, second)
	vm.PushValue(result)
}