/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/vm"
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
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
		v.VisitStmt(ctx.GetChild(i).(*parser.StmtContext))
	}
	//log.Debug(v.GlobalState)
}

// 执行语句
func (v *SPADirectInterpreter) VisitStmt(ctx *parser.StmtContext) {
	log.Debug("Visit Stmt")

	switch ctx.GetChild(0).(type) {
	case *parser.Expr_stmtContext:
		v.VisitExpr_stmt(ctx.GetChild(0).(*parser.Expr_stmtContext))
	case *parser.Return_stmtContext:
		v.VisitReturn_stmt(ctx.GetChild(0).(*parser.Return_stmtContext))
	}
}

func (v *SPADirectInterpreter) VisitReturn_stmt(ctx *parser.Return_stmtContext) interface{} {
	if ctx.GetChildCount() == 2 {
		return v.VisitPostfix_expr(ctx.GetChild(1).(*parser.Postfix_exprContext)).(vm.SPAValue)
	}
	return vm.Null()
}

func (v *SPADirectInterpreter) VisitExpr_stmt(ctx *parser.Expr_stmtContext) interface{} {
	log.Debug("Visit Expr_Stmt")

	if ctx.GetChildCount() == 1 {
		v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
	} else if ctx.GetChildCount() == 4 {
		name := ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
		pars := ctx.GetChild(2).(*parser.Par_seqContext)
		body := ctx.GetChild(3).(*parser.BlockContext)
		parNames := getParList(pars)
		f := vm.NewFunction(name, parNames, body)
		f.Outer = v.GlobalState
		v.GlobalState.Define(f) //函数定义
		log.Infof("函数定义: %v", v.GlobalState)
	} else {
		// 赋值
		name := v.VisitPrimary_expr(ctx.GetChild(0).(*parser.Primary_exprContext)).(string)
		value := v.VisitPostfix_expr(ctx.GetChild(2).(*parser.Postfix_exprContext)).(vm.SPAValue)
		sym := vm.NewVariable(name, value)
		v.GlobalState.Define(sym)
		log.Infof("变量定义: %v", v.GlobalState)
		//log.Infof("%s = %t", name, value)
	}
	return vm.Null()
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

func (v *SPADirectInterpreter) VisitPrimary_expr(ctx *parser.Primary_exprContext) interface{} {
	return ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
}

func (v *SPADirectInterpreter) VisitPostfix_expr(ctx *parser.Postfix_exprContext) interface{} {
	//return v.VisitChildren(ctx)
	log.Debug("Visit Postfix Expr")

	return v.VisitOr_test(ctx.GetChild(0).(*parser.Or_testContext))
}

func (v *SPADirectInterpreter) VisitBlock(ctx *parser.BlockContext){
	//var result vm.SPAValue
	for i := 1; i < ctx.GetChildCount()-1; i++ {
		v.VisitStmt(ctx.GetChild(i).(*parser.StmtContext))
	}
	//return result
}

func (v *SPADirectInterpreter) VisitOr_test(ctx *parser.Or_testContext) interface{} {
	log.Debug("Visit Or_Test")

	var result vm.SPAValue
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitAnd_test(ctx.GetChild(0).(*parser.And_testContext)).(vm.SPAValue)
		if result.IsTrue() {
			return result
		}
	}
	return result
}

func (v *SPADirectInterpreter) VisitAnd_test(ctx *parser.And_testContext) interface{} {
	log.Debug("Visit And_Test")

	var result vm.SPAValue
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitNot_test(ctx.GetChild(i).(*parser.Not_testContext)).(vm.SPAValue)
		if !result.IsTrue() {
			return result
		}
	}
	return result
}

func (v *SPADirectInterpreter) VisitNot_test(ctx *parser.Not_testContext) interface{} {
	log.Debug("Visit Not_Test")

	if ctx.GetChildCount() == 2 {
		result := v.VisitNot_test(ctx.GetChild(1).(*parser.Not_testContext)).(vm.SPAValue)
		return vm.InvertBool(result)
	} else {
		return v.VisitComparison(ctx.GetChild(0).(*parser.Compare_exprContext))
	}
}

func (v *SPADirectInterpreter) VisitComparison(ctx *parser.Compare_exprContext) interface{} {
	log.Debug("Visit Comparison")

	// 目前先不支持比较，只返回表达式
	return v.VisitArith_expr(ctx.GetChild(0).(*parser.Arith_exprContext))
}

func (v *SPADirectInterpreter) VisitArith_expr(ctx *parser.Arith_exprContext) interface{} {
	log.Debug("Visit Arith_Expr")

	if ctx.GetChildCount() == 1 {
		return v.VisitTerm(ctx.GetChild(0).(*parser.TermContext))
	}else{
		var termResult vm.SPAValue = vm.SPANumber(1.0) //默认为0，累加的时候不需要判断了
		var op = "+"                                   // 统一处理
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitTerm(ctx.GetChild(i).(*parser.TermContext))
				termResult = arithmetic(termResult, r.(vm.SPAValue), op)
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *SPADirectInterpreter) VisitTerm(ctx *parser.TermContext) interface{} {
	log.Debug("Visit Term")

	//r := v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	if ctx.GetChildCount() == 1 {
		return v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	} else {
		var termResult vm.SPAValue = vm.SPANumber(1.0) //默认为1方便处理
		var op = "*"
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitFactor(ctx.GetChild(i).(*parser.FactorContext))
				termResult = arithmetic(termResult, r.(vm.SPAValue), op)
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *SPADirectInterpreter) VisitFactor(ctx *parser.FactorContext) interface{} {
	log.Debug("Visit Factor")
	//return v.VisitChildren(ctx)
	if ctx.GetChildCount() == 1 {
		return v.VisitPower(ctx.GetChild(0).(*parser.PowerContext))
	} else {
		//op := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		r := v.VisitFactor(ctx.GetChild(1).(*parser.FactorContext)).(vm.SPANumber)
		return vm.Negative(r)
	}
}

func (v *SPADirectInterpreter) VisitPower(ctx *parser.PowerContext) interface{} {
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

func (v *SPADirectInterpreter) VisitAtom_expr(ctx *parser.Atom_exprContext) interface{} {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		return v.VisitPostfix_expr(ctx.GetChild(1).(*parser.Postfix_exprContext))
	}
	if ctx.GetChildCount() == 2 {
		// 函数调用
		name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()                  //获取函数名
		args := v.VisitArg_seq(ctx.GetChild(1).(*parser.Arg_seqContext)).([]vm.SPAValue) //获取参数列表
		f, ok := v.GlobalState.Resolve(name).(*vm.SPAFunction)                           //获取函数定义
		if !ok {
			panic("函数未定义")
		}
		ci := vm.NewCallInfo(f)                //创建函数调用信息
		ci.PushArgs(args)                      //传入参数
		vm.PushCallInfo(ci)                    //保存到函数调用栈
		v.VisitBlock(vm.GetTopCallInfo().Body) //调用函数
		vm.PopCallInfo()                       //结束调用
		return vm.Null()
	}

	//终结符
	terminalNode := ctx.GetChild(0).(*antlr.TerminalNodeImpl)
	tt := terminalNode.GetSymbol().GetTokenType()
	switch tt {
	case parser.SpartaLexerIDENTIFIER:
		if vm.HasCallInfo(){
			return vm.GetTopCallInfo().Resolve(terminalNode.GetText()).(vm.VariableSymbol).Value
		}
		return v.GlobalState.Resolve(terminalNode.GetText()).(vm.VariableSymbol).Value
	case parser.SpartaLexerNUMBER_LITERAL:
		v := vm.NewNumber(terminalNode.GetText())
		return v
	case parser.SpartaLexerSTRING:
		return vm.NewString(terminalNode.GetText())
	default:
		panic("类型无效")
	}
}

func (v *SPADirectInterpreter) VisitArg_seq(ctx *parser.Arg_seqContext) interface{} {
	if ctx.GetChildCount() == 2{
		return []vm.SPAValue{}
	}
	return v.VisitArg_list(ctx.GetChild(1).(*parser.Arg_listContext))
}

func (v *SPADirectInterpreter) VisitArg_list(ctx *parser.Arg_listContext) interface{} {
	//return v.VisitChildren(ctx)
	argList := make([]vm.SPAValue, 0, 10)
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		v := v.VisitArgument(ctx.GetChild(i).(*parser.ArgumentContext)).(vm.SPAValue)
		argList = append(argList, v)
	}
	return argList
}

func (v *SPADirectInterpreter) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	return v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
}

/**
算数运算
 */
func arithmetic(first, second vm.SPAValue, op string) vm.SPAValue  {
	//switch op {
	//case "+":
	//	return first.(*vm.SPANumber) + second.(*vm.SPANumber)
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