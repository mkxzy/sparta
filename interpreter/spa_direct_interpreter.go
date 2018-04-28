/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/function"
)

var log = logging.MustGetLogger("SPADirectInterpreter")

type SPADirectInterpreter struct {
	//*parser.BaseSpartaVisitor
	GlobalState *symbol.MemorySpace
}

func NewDirectInterpreter(globalState *symbol.MemorySpace) *SPADirectInterpreter {
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
	//name := ctx.GetChild(0).(*parser.Left_sideContext).GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()

	as := &AssignState{}
	v.EvalLeftSide(ctx.GetChild(0).(*parser.Left_sideContext), as)

	if as.State == NewVariable {
		value := PopValue()
		sym := symbol.NewVariable(as.Name, value)
		if HasCallInfo(){
			GetTopCallInfo().Define(sym) //局部变量定义
			//log.Infof("局部变量定义： %v", vm.GetTopCallInfo())
		}else{
			v.GlobalState.Define(sym) 		//全局变量定义
			//log.Infof("全局变量定义: %v", v.GlobalState)
		}
	} else {
		var sym symbol.Symbol
		if HasCallInfo(){
			sym = GetTopCallInfo().Resolve(as.Name) 	//局部变量定义
		}else{
			sym = v.GlobalState.Resolve(as.Name) 		//全局变量定义
		}
		if sym == nil {
			panic("")
		}
		list := sym.(*symbol.SPAVariable).Value.(*types.SPAList)
		index, ok := PopValue().(types.SPAInteger)
		if !ok {
			panic("index必须是整数")
		}
		v := PopValue()
		list.Set(index, v)
	}
}

func(v *SPADirectInterpreter) EvalLeftSide(ctx *parser.Left_sideContext, as *AssignState)  {
	name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	if ctx.GetChildCount() == 1 {
		as.Name = name
		as.State = NewVariable
	} else {
		as.Name = name
		as.State = ListAccess
		v.EvalTest(ctx.GetChild(2).(*parser.TestContext))
	}
}

/**
定义函数表达式
 */
func (v *SPADirectInterpreter) ExecFunDefStmt(ctx *parser.Fundef_stmtContext)  {
	f := function.SPAFunction{}
	f.Name = ctx.GetChild(1).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	f.Body = ctx.GetChild(2).(*parser.Fun_bodyContext)
	v.EvalParList(f.Body.GetChild(0).(*parser.Fun_parContext), &f)
	f.Outer = v.GlobalState
	sym := function.NewFunVariable(f)
	v.GlobalState.Define(sym) //函数定义
	log.Infof("函数定义: %v", v.GlobalState)
}

/**
获取形参列表
 */
func (v *SPADirectInterpreter) EvalParList(ctx *parser.Fun_parContext, f *function.SPAFunction) {

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
		PushValue(types.Null()) //如果return没有参数，那么插入一个空的返回值
	}
}

/**
函数调用语句
 */
func (v *SPADirectInterpreter) ExecFunCallStmt(ctx *parser.Funcall_stmtContext)  {
	v.EvalFunCallExpr(ctx.GetChild(0).(*parser.Funcall_exprContext))
	PopValue() //丢弃返回值
}

/**
if语句执行
 */
func (v *SPADirectInterpreter) ExecIfStmt(ctx *parser.If_stmtContext, ff FlowState) {
	v.EvalTest(ctx.GetChild(1).(*parser.TestContext))
	testResult := PopValue()
	if testResult.IsTrue() {
		v.ExecBlock(ctx.GetChild(2).(*parser.BlockContext), ff)
	}
	pos := 3
	for pos + 4 < ctx.GetChildCount() {
		v.EvalTest(ctx.GetChild(pos + 2).(*parser.TestContext))
		testResult = PopValue()
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

/**
比较表达式
 */
func (v *SPADirectInterpreter) EvalCompareExpr(ctx *parser.Compare_exprContext) {
	v.EvalArithExpr(ctx.GetChild(0).(*parser.Arith_exprContext))
	if ctx.GetChildCount() == 3 {
		v.EvalArithExpr(ctx.GetChild(2).(*parser.Arith_exprContext))
		op := ctx.GetChild(1).(*parser.Comp_opContext).GetChild(0).(antlr.TerminalNode).GetText()
		switch op {
		case "==":
			second := PopValue().(types.SPAInteger)
			first := PopValue().(types.SPAInteger)
			result := first == second
			PushValue(types.SPABool(result))
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
		first := PopValue()
		result, _ := minus(first)
		PushValue(result)
	}
}

/**
原子表达式
 */
func (v *SPADirectInterpreter) EvalAtomExpr(ctx *parser.Atom_exprContext) {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		firstChild := ctx.GetChild(0).(antlr.TerminalNode)
		switch firstChild.GetText() {
		case "(":
			// 括号优先表达式
			v.EvalTest(ctx.GetChild(1).(*parser.TestContext))
		case "[":
			v.EvalTestList(ctx.GetChild(1).(*parser.Test_listContext))
		}
	} else if ctx.GetChildCount() == 4 {
		v.EvalTest(ctx.GetChild(2).(*parser.TestContext))
		name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
		var sym *symbol.SPAVariable
		if HasCallInfo(){
			sym = GetTopCallInfo().Resolve(name).(*symbol.SPAVariable)
		} else {
			sym = v.GlobalState.Resolve(name).(*symbol.SPAVariable)
		}
		value := sym.Value.(*types.SPAList)
		index := PopValue().(types.SPAInteger)
		PushValue(value.Get(index))
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
				var value types.SPAValue
				if HasCallInfo(){
					value = GetTopCallInfo().Resolve(terminalNode.GetText()).(*symbol.SPAVariable).Value
				} else {
					value = v.GlobalState.Resolve(terminalNode.GetText()).(*symbol.SPAVariable).Value
				}
				PushValue(value)
			case parser.SpartaLexerNUMBER_LITERAL:
				PushValue(types.NewNumber(terminalNode.GetText()))
			case parser.SpartaLexerINTEGER_LITERAL:
				PushValue(types.NewInteger(terminalNode.GetText()))
			case parser.SpartaLexerSTRING:
				PushValue(types.NewString(terminalNode.GetText()))
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
	va, ok := v.GlobalState.Resolve(fs.FunName).(*symbol.SPAVariable)    //获取函数定义
	if !ok {
		panic("未找到函数定义")
	}
	f, ok := va.Value.(function.SPAFunction)
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
			var arg types.SPAValue
			for i := fs.ArgCount; i > 0; i--{
				arg = PopValue()
			}
			fmt.Println(arg)
			PushNullValue() //函数体内没有返回语句，自动插入空返回值
		}
	} else {
		PassArgs(fs)
		ci := NewCallInfo(fs) //创建函数调用信息
		//passArgs(ci, fs.ArgCount)		//传递参数
		PushCallInfo(ci)    	//保存到函数调用栈

		defer PopCallInfo() //函数退出时弹出调用栈

		//var ret bool
		//去掉大括号
		bockCtx := fs.Function.Body.GetChild(1)
		v.ExecBlock(bockCtx.(*parser.BlockContext), fs)
		if fs.State != RETURN {
			PushNullValue()
		}
	}
}


func PassArgs(fs *FunState) {
	fs.Args = make([]types.SPAValue, fs.ArgCount, fs.ArgCount)
	for i := fs.ArgCount - 1; i >= 0; i-- {
		v := PopValue()
		fs.Args[i] = v
	}
}

// 循环体
func (v *SPADirectInterpreter) ExecForStmt(ctx *parser.For_stmtContext, ff FlowState) {

	forState := &ForState{State:NORMAL}
	forState.ItemName = ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	sym := symbol.NewVariable(forState.ItemName, types.Null())

	if HasCallInfo(){
		GetTopCallInfo().Define(sym) //局部变量定义
		//log.Infof("局部变量定义： %v", vm.GetTopCallInfo())
	}else{
		v.GlobalState.Define(sym) 		//全局变量定义
		//log.Infof("全局变量定义: %v", v.GlobalState)
	}
	v.EvalTest(ctx.GetChild(3).(*parser.TestContext))
	v.EvalTest(ctx.GetChild(5).(*parser.TestContext))
	toNumber, ok := PopValue().(types.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	fromNumber, ok := PopValue().(types.SPAInteger)
	if !ok {
		panic("类型不正确")
	}
	if fromNumber > toNumber{
		panic("起始数字不能大于结束数字")
	}

	for i := fromNumber; i <= toNumber; i++ {
		sym.Value = types.SPAInteger(i)
		v.ExecBlock(ctx.GetChild(6).(*parser.BlockContext), forState)
		if forState.State == BREAK {
			forState.SetState(NORMAL) //恢复状态
		}
		if forState.State == CONTINUE {
			forState.SetState(NORMAL) //恢复状态
		}
		if forState.State == RETURN {
			forState.SetState(NORMAL) //恢复状态
			ff.SetState(RETURN) 	  //传递给调用者
		}
	}
}

/**
循环继续
 */
func (v *SPADirectInterpreter) ExecContinueStmt(ctx *parser.Continue_stmtContext, state FlowState) {
	state.SetState(CONTINUE)
}

/**
中断循环
 */
func (v *SPADirectInterpreter) ExecBreakStmt(ctx *parser.Break_stmtContext, state FlowState) {
	state.SetState(BREAK)
}

func (v *SPADirectInterpreter) EvalTestList(ctx *parser.Test_listContext) {
	elementCount := (ctx.GetChildCount() + 1) / 2
	list := types.NewList(elementCount)
	for i := 0; i < ctx.GetChildCount(); i += 2{
		v.EvalTest(ctx.GetChild(i).(*parser.TestContext))
		value := PopValue()
		list.Append(value)
	}
	PushValue(list)
	log.Info(list)
}

/**
算数运算
 */
func arithmetic(op string) {
	second := PopValue()
	first := PopValue()
	operation := operations[op]
	if operation == nil{
		panic("不支持的操作")
	}
	result, _ := operation(first, second)
	PushValue(result)
}