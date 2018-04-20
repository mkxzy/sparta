package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"github.com/mkxzy/sparta/base"
)

var log = logging.MustGetLogger("ExprVisitor")

var calls [10000]*base.CallInfo            //栈空间最多10000
var fp = -1                                 //函数栈指针
//var returnValue base.SPAValue = base.Null() //函数返回值
var state = base.NewMemorySpace("global")  //全局内存空间

// 函数调用信息入栈
func pushCallInfo(ci *base.CallInfo)  {
	log.Debug("函数调用前", state)
	//v.state = f
	fp++
	calls[fp] = ci
}

// 函数调用信息出栈
func popCallInfo()  {
	log.Debug("函数调用后", state)
	//state = f.Outer
	if hasCallInfo() {
		calls[fp] = nil
		fp--
	} else {
		panic("函数栈为空")
	}
}

// 获取栈顶函数调用信息
func getTopCallInfo() *base.CallInfo  {
	if hasCallInfo() {
		return calls[fp]
	}
	panic("函数栈为空")
}

// 判断是否有函数调用
func hasCallInfo() bool  {
	return fp > -1
}

type SPAInterpreter struct {
	*parser.BaseSpartaVisitor
}

func NewInterpreter() *SPAInterpreter {
	return &SPAInterpreter{
	}
}

func (v *SPAInterpreter) VisitProgram(ctx *parser.ProgramContext) interface{} {

	//nu := base.Null()
	//log.Debug(nu.IsTrue())
	log.Debug("Visit Program")
	for i := 0; i < ctx.GetChildCount()-1; i++ {
		v.VisitStmt(ctx.GetChild(i).(*parser.StmtContext))
	}
	log.Debug(state)
	return base.Null()
}

func (v *SPAInterpreter) VisitStmt(ctx *parser.StmtContext) interface{} {
	log.Debug("Visit Stmt")
	v.VisitSimple_stmt(ctx.GetChild(0).(*parser.Simple_stmtContext))
	return base.Null()
}

func (v *SPAInterpreter) VisitSimple_stmt(ctx *parser.Simple_stmtContext) interface{} {
	log.Debug("Visit Simple_Stmt")

	//return v.VisitChildren(ctx)
	
	switch ctx.GetChild(0).(type) {
	case *parser.Expr_stmtContext:
		v.VisitExpr_stmt(ctx.GetChild(0).(*parser.Expr_stmtContext))
	case *parser.Return_stmtContext:
		v.VisitReturn_stmt(ctx.GetChild(0).(*parser.Return_stmtContext))
	}
	return base.Null()
}

func (v *SPAInterpreter) VisitReturn_stmt(ctx *parser.Return_stmtContext) interface{} {
	if ctx.GetChildCount() == 2 {
		return v.VisitPostfix_expr(ctx.GetChild(1).(*parser.Postfix_exprContext)).(base.SPAValue)
		//log.Debug("函数返回值", returnValue)
	}
	return base.Null()
}

func (v *SPAInterpreter) VisitExpr_stmt(ctx *parser.Expr_stmtContext) interface{} {
	log.Debug("Visit Expr_Stmt")

	if ctx.GetChildCount() == 1 {
		v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
	} else if ctx.GetChildCount() == 4 {
		name := ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
		pars := ctx.GetChild(2).(*parser.Par_seqContext)
		body := ctx.GetChild(3).(*parser.BlockContext)
		parNames := getParList(pars)
		f := base.NewFunction(name, parNames, body)
		f.Outer = state
		state.Define(f) //函数定义
		log.Infof("函数定义: %v", state)
	} else {
		// 赋值
		name := v.VisitPrimary_expr(ctx.GetChild(0).(*parser.Primary_exprContext)).(string)
		value := v.VisitPostfix_expr(ctx.GetChild(2).(*parser.Postfix_exprContext)).(base.SPAValue)
		sym := base.NewVariable(name, value)
		state.Define(sym)
		log.Infof("变量定义: %v", state)
		//log.Infof("%s = %t", name, value)
	}
	return base.Null()
}

func getParList(ctx *parser.Par_seqContext) []string {
	var parNames []string = make([]string, 0, 10)
	if ctx.GetChildCount() == 3 {
		nameList := ctx.GetChild(1).GetChild(0).(*parser.NamelistContext)
		for i := 0; i < nameList.GetChildCount(); i += 2{
			name := nameList.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			parNames = append(parNames, name)
		}
	}
	return parNames
}

func (v *SPAInterpreter) VisitPrimary_expr(ctx *parser.Primary_exprContext) interface{} {
	return ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
}

func (v *SPAInterpreter) VisitPostfix_expr(ctx *parser.Postfix_exprContext) interface{} {
	//return v.VisitChildren(ctx)
	log.Debug("Visit Postfix Expr")

	return v.VisitOr_test(ctx.GetChild(0).(*parser.Or_testContext))

	//if ctx.GetChildCount() == 1{
	//	return v.VisitOr_test(ctx.GetChild(0).(*parser.Or_testContext))
	//} else {
	//	log.Debug("函数定义")
	//	pars := ctx.GetChild(1).(*parser.Par_seqContext)
	//	body := ctx.GetChild(2).(*parser.BlockContext)
	//	return v.defineFun(pars, body)
	//}
}

func (v *SPAInterpreter) VisitBlock(ctx *parser.BlockContext) interface{} {
	var result base.SPAValue
	for i := 1; i < ctx.GetChildCount()-1; i++ {
		result = v.VisitStmt(ctx.GetChild(i).(*parser.StmtContext)).(base.SPAValue)
	}
	return result
}

func (v *SPAInterpreter) VisitOr_test(ctx *parser.Or_testContext) interface{} {
	log.Debug("Visit Or_Test")

	var result base.SPAValue
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitAnd_test(ctx.GetChild(0).(*parser.And_testContext)).(base.SPAValue)
		if result.IsTrue() {
			return result
		}
	}
	return result
}

func (v *SPAInterpreter) VisitAnd_test(ctx *parser.And_testContext) interface{} {
	log.Debug("Visit And_Test")

	var result base.SPAValue
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitNot_test(ctx.GetChild(i).(*parser.Not_testContext)).(base.SPAValue)
		if !result.IsTrue() {
			return result
		}
	}
	return result
}

func (v *SPAInterpreter) VisitNot_test(ctx *parser.Not_testContext) interface{} {
	log.Debug("Visit Not_Test")

	if ctx.GetChildCount() == 2 {
		result := v.VisitNot_test(ctx.GetChild(1).(*parser.Not_testContext)).(base.SPAValue)
		return base.InvertBool(result)
	} else {
		return v.VisitComparison(ctx.GetChild(0).(*parser.Compare_exprContext))
	}
}

func (v *SPAInterpreter) VisitComparison(ctx *parser.Compare_exprContext) interface{} {
	log.Debug("Visit Comparison")

	// 目前先不支持比较，只返回表达式
	return v.VisitArith_expr(ctx.GetChild(0).(*parser.Arith_exprContext))
}

func (v *SPAInterpreter) VisitArith_expr(ctx *parser.Arith_exprContext) interface{} {
	log.Debug("Visit Arith_Expr")

	if ctx.GetChildCount() == 1 {
		return v.VisitTerm(ctx.GetChild(0).(*parser.TermContext))
	}else{
		var termResult base.SPAValue = base.SPANumber(1.0) //默认为0，累加的时候不需要判断了
		var op = "+" // 统一处理
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitTerm(ctx.GetChild(i).(*parser.TermContext))
				termResult = arithmetic(termResult, r.(base.SPAValue), op)
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *SPAInterpreter) VisitTerm(ctx *parser.TermContext) interface{} {
	log.Debug("Visit Term")

	//r := v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	if ctx.GetChildCount() == 1 {
		return v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	} else {
		var termResult base.SPAValue = base.SPANumber(1.0) //默认为1方便处理
		var op = "*"
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitFactor(ctx.GetChild(i).(*parser.FactorContext))
				termResult = arithmetic(termResult, r.(base.SPAValue), op)
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *SPAInterpreter) VisitFactor(ctx *parser.FactorContext) interface{} {
	log.Debug("Visit Factor")
	//return v.VisitChildren(ctx)
	if ctx.GetChildCount() == 1 {
		return v.VisitPower(ctx.GetChild(0).(*parser.PowerContext))
	} else {
		//op := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		r := v.VisitFactor(ctx.GetChild(1).(*parser.FactorContext)).(base.SPANumber)
		return base.Negative(r)
	}
}

func (v *SPAInterpreter) VisitPower(ctx *parser.PowerContext) interface{} {
	log.Debug("Visit Power")

	atom := v.VisitAtom_expr(ctx.GetChild(0).(*parser.Atom_exprContext))
	if atom == nil {
		return nil
	}
	if ctx.GetChildCount() == 1 {
		return atom
	} else {
		factor := v.VisitFactor(ctx.GetChild(2).(*parser.FactorContext)).(float64)
		return math.Pow(atom.(float64), factor)
	}
}

func (v *SPAInterpreter) VisitAtom_expr(ctx *parser.Atom_exprContext) interface{} {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		return v.VisitPostfix_expr(ctx.GetChild(1).(*parser.Postfix_exprContext))
	}
	if ctx.GetChildCount() == 2 {
		// 函数调用
		name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()                    	//获取函数名
		args := v.VisitArg_seq(ctx.GetChild(1).(*parser.Arg_seqContext)).([]base.SPAValue) 	//获取参数列表
		f, ok := state.Resolve(name).(*base.SPAFunction)                                  		//获取函数定义
		if !ok {
			panic("函数未定义")
		}
		ci := base.NewCallInfo(f) //创建函数调用信息
		ci.PushArgs(args) //传入参数
		pushCallInfo(ci) //保存到函数调用栈
		v.VisitBlock(getTopCallInfo().Body) //调用函数
		popCallInfo() //结束调用
		return base.Null()
	}

	//终结符
	terminalNode := ctx.GetChild(0).(*antlr.TerminalNodeImpl)
	tt := terminalNode.GetSymbol().GetTokenType()
	switch tt {
	case parser.SpartaLexerIDENTIFIER:
		if hasCallInfo(){
			return getTopCallInfo().Resolve(terminalNode.GetText()).(base.VariableSymbol).Value
		}
		return state.Resolve(terminalNode.GetText()).(base.VariableSymbol).Value
	case parser.SpartaLexerNUMBER_LITERAL:
		v := base.NewNumber(terminalNode.GetText())
		return v
	case parser.SpartaLexerSTRING:
		return base.NewString(terminalNode.GetText())
	default:
		panic("类型无效")
	}
}

func (v *SPAInterpreter) VisitArg_seq(ctx *parser.Arg_seqContext) interface{} {
	if ctx.GetChildCount() == 2{
		return []base.SPAValue{}
	}
	return v.VisitArg_list(ctx.GetChild(1).(*parser.Arg_listContext))
}

func (v *SPAInterpreter) VisitArg_list(ctx *parser.Arg_listContext) interface{} {
	//return v.VisitChildren(ctx)
	argList := make([]base.SPAValue, 0, 10)
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		v := v.VisitArgument(ctx.GetChild(i).(*parser.ArgumentContext)).(base.SPAValue)
		argList = append(argList, v)
	}
	return argList
}

func (v *SPAInterpreter) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	return v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
}

/**
算数运算
 */
func arithmetic(first, second base.SPAValue, op string) base.SPAValue  {
	//switch op {
	//case "+":
	//	return first.(*base.SPANumber) + second.(*base.SPANumber)
	//case "-":
	//	return first - second
	//case "*":
	//	return first * second
	//case "/":
	//	return first / second
	//default:
	//	panic("不支持的操作")
	//}
	return first
}