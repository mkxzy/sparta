/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/vm"
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
		v.ExecStmt(stmtContext)
	}
}

// 执行语句
// 返回值： 是否是返回语句
func (v *SPADirectInterpreter) ExecStmt(ctx *parser.StmtContext) bool {
	log.Debug("Visit Stmt")

	rule := ctx.GetChild(0).(antlr.RuleContext)
	switch rule.GetRuleIndex() {
	case parser.SpartaParserRULE_assign_stmt:
		v.ExecAssignStmt(rule.(*parser.Assign_stmtContext))
		return false
	case parser.SpartaParserRULE_fundef_stmt:
		v.ExecFunDefStmt(rule.(*parser.Fundef_stmtContext))
		return false
	case parser.SpartaParserRULE_return_stmt:
		v.ExecReturnStmt(rule.(*parser.Return_stmtContext))
		return true
	case parser.SpartaParserRULE_funcall_stmt:
		v.ExecFunCallStmt(rule.(*parser.Funcall_stmtContext))
		return false
	default:
		return false
	}
}

/**
赋值语句
 */
func (v *SPADirectInterpreter) ExecAssignStmt(ctx *parser.Assign_stmtContext)  {

	v.EvalPostExpr(ctx.GetChild(2).(*parser.Postfix_exprContext))
	name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	value := vm.PopValue()
	sym := vm.NewVariable(name, value)

	if vm.HasCallInfo(){
		vm.GetTopCallInfo().Define(sym)
		//log.Infof("局部变量定义： %v", vm.GetTopCallInfo())
	}else{
		v.GlobalState.Define(sym)
		//log.Infof("全局变量定义: %v", v.GlobalState)
	}
}

/**
定义函数表达式
 */
func (v *SPADirectInterpreter) ExecFunDefStmt(ctx *parser.Fundef_stmtContext)  {
	name := ctx.GetChild(1).GetChild(0).(*antlr.TerminalNodeImpl).GetText()
	//pars := ctx.GetChild(2).(*parser.Fun_parContext)
	body := ctx.GetChild(2).(*parser.Fun_bodyContext)
	parNames := getParList(body.GetChild(0).(*parser.Fun_parContext))
	log.Infof("参数： %v", parNames)
	f := vm.NewFunction(name, parNames, body)
	f.Outer = v.GlobalState
	sym := vm.NewVariable(name, f)
	v.GlobalState.Define(sym) //函数定义
	log.Infof("函数定义: %v", v.GlobalState)
}

/**
函数返回表达式
 */
func (v *SPADirectInterpreter) ExecReturnStmt(ctx *parser.Return_stmtContext) {
	if !vm.HasCallInfo(){
		panic("未执行函数不能用返回语句")
	}
	if ctx.GetChildCount() == 2 {
		v.EvalPostExpr(ctx.GetChild(1).(*parser.Postfix_exprContext))
	} else{
		vm.PushValue(vm.Null()) //如果return没有参数，那么插入一个空的返回值
	}
}

/**
函数返回表达式
 */
func (v *SPADirectInterpreter) ExecFunCallStmt(ctx *parser.Funcall_stmtContext)  {
	v.EvalFunCallExpr(ctx.GetChild(0).(*parser.Funcall_exprContext))
}

/**
获取形参列表
 */
func getParList(ctx *parser.Fun_parContext) []string {

	context := ctx.GetChild(1).(*parser.NamelistContext)
	var parNames = make([]string, 0, 10)
	for i := 0; i < context.GetChildCount(); i += 2{
		name := context.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
		parNames = append(parNames, name)
	}
	return parNames
}

/**
右值表达式
 */
func (v *SPADirectInterpreter) EvalPostExpr(ctx *parser.Postfix_exprContext) {
	//return v.VisitChildren(ctx)
	log.Debug("Visit Postfix Expr")

	v.EvalArithExpr(ctx.GetChild(0).(*parser.Arith_exprContext))
}

// 计算算术表达式
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

// 计算算术表达式
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
		minusValue() //取反
	}
}

/**
原子表达式
 */
func (v *SPADirectInterpreter) EvalAtomExpr(ctx *parser.Atom_exprContext) {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		// 括号优先表达式
		v.EvalPostExpr(ctx.GetChild(1).(*parser.Postfix_exprContext))
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
			case parser.SpartaLexerSTRING:
				vm.PushValue(vm.NewString(terminalNode.GetText()))
			default:
				panic("类型无效")
			}
		}
	}
}

func(v *SPADirectInterpreter) EvalFunCallExpr(funCallExpr *parser.Funcall_exprContext)  {
	name := getFunName(funCallExpr.GetChild(0).(*parser.Fun_nameContext)) 		//获取函数名
	argCount := v.pushArgs(funCallExpr.GetChild(1).(*parser.Arg_exprContext))		//参数入栈
	va, ok := v.GlobalState.Resolve(name).(*vm.VariableSymbol)                     	//获取函数定义
	if !ok {
		panic("未找到函数定义")
	}
	f, ok := va.Value.(*vm.SPAFunction)
	if !ok {
		panic("未找到函数定义")
	}
	v.CallFunc(f, argCount) //调用函数
}

// 判断是否函数调用表达式
func isFunCallExpr(expr antlr.Tree) bool {
	_, ok := expr.(*parser.Funcall_exprContext)
	return ok
}

// 获取函数名
func getFunName(ctx *parser.Fun_nameContext) string {
	return ctx.GetChild(0).(antlr.TerminalNode).GetText()
}

// 参数压栈
func(v *SPADirectInterpreter) pushArgs(ctx *parser.Arg_exprContext) int {
	if ctx.GetChildCount() == 2{
		return 0
	}
	args := 0
	argListContext := ctx.GetChild(1).(*parser.Arg_listContext)
	for i := 0; i < argListContext.GetChildCount(); i += 2 {
		v.EvalArgument(argListContext.GetChild(i).(*parser.ArgContext))						//参数压栈
		args++
	}
	return args
}

// 调用函数
func (v *SPADirectInterpreter) CallFunc(f *vm.SPAFunction, args int) {
	log.Info(f)
	ci := vm.NewCallInfo(f) //创建函数调用信息
	//ci.PushArgs(args)      	//传入参数
	passArgs(ci, args)
	log.Info(ci.Symbols)
	vm.PushCallInfo(ci)    	//保存到函数调用栈
	log.Info(vm.GetTopCallInfo())

	defer vm.PopCallInfo() //函数退出时弹出调用栈

	//var ret bool
	//去掉大括号
	for i := 2; i < f.Body.GetChildCount()-1; i++ {
		ret := v.ExecStmt(f.Body.GetChild(i).(*parser.StmtContext))
		if ret {
			return
		}
	}
	vm.PushValue(vm.Null())
}

// 传递参数
func passArgs(ci *vm.CallInfo, args int)  {
	for i := 0; i < args; i++{
		ci.Define(vm.NewVariable(ci.Args[i], vm.PopValue()))
		if i >= args {
			ci.Define(vm.NewVariable(ci.Args[i], vm.Null())) //不足的参数用NULL来补
		}
	}
}

// 参数入栈
func (v *SPADirectInterpreter) EvalArgument(ctx *parser.ArgContext) {
	v.EvalPostExpr(ctx.GetChild(0).(*parser.Postfix_exprContext))
}

/**
算数运算
 */
func arithmetic(op string) {
	second := vm.PopValue().(vm.SPANumber)
	first := vm.PopValue().(vm.SPANumber)
	switch op {
	case "+":
		result := first + second
		vm.PushValue(result)
	case "-":
		result := first - second
		vm.PushValue(result)
	case "*":
		result := first * second
		vm.PushValue(result)
	case "/":
		result := first / second
		vm.PushValue(result)
	default:
		panic("不支持的操作")
	}
}

/**
取反
 */
func minusValue()  {
	first := vm.PopValue().(vm.SPANumber)
	first = -first
	vm.PushValue(first)
}