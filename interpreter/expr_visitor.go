package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"strconv"
	"fmt"
	"github.com/mkxzy/sparta/base"
)

var log = logging.MustGetLogger("ExprVisitor")

type ExpVisitor struct {
	*parser.BaseSpartaVisitor
	vars map[string]interface{}
}

func NewExpVisitor() *ExpVisitor {
	return &ExpVisitor{
		vars: make(map[string]interface{}),
	}
}

func (v *ExpVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	//nu := base.Null()
	//log.Debug(nu.IsTrue())
	log.Debug("Visit Program")
	for i := 0; i < ctx.GetChildCount()-1; i++ {
		v.VisitStmt(ctx.GetChild(i).(*parser.StmtContext))
	}
	log.Debug(v.vars)
	return nil
}

func (v *ExpVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	log.Debug("Visit Stmt")
	v.VisitSimple_stmt(ctx.GetChild(0).(*parser.Simple_stmtContext))
	return nil
}

func (v *ExpVisitor) VisitSimple_stmt(ctx *parser.Simple_stmtContext) interface{} {
	log.Debug("Visit Simple_Stmt")

	//return v.VisitChildren(ctx)
	v.VisitExpr_stmt(ctx.GetChild(0).(*parser.Expr_stmtContext))
	return nil
}

func (v *ExpVisitor) VisitExpr_stmt(ctx *parser.Expr_stmtContext) interface{} {
	log.Debug("Visit Expr_Stmt")

	if ctx.GetChildCount() == 1 {
		v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
	} else if ctx.GetChildCount() == 4 {
		log.Debug("函数定义")
	} else {
		// 赋值
		name := v.VisitPrimary_expr(ctx.GetChild(0).(*parser.Primary_exprContext)).(string)
		value := v.VisitPostfix_expr(ctx.GetChild(2).(*parser.Postfix_exprContext))
		v.putVar(name, value)
		log.Debugf("%s = %t", name, value)
	}
	return nil
}

func (v *ExpVisitor) VisitPrimary_expr(ctx *parser.Primary_exprContext) interface{} {
	return ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
}

func (v *ExpVisitor) VisitPostfix_expr(ctx *parser.Postfix_exprContext) interface{} {
	//return v.VisitChildren(ctx)
	if ctx.GetChildCount() == 1{
		return v.VisitOr_test(ctx.GetChild(0).(*parser.Or_testContext))
	} else {
		log.Debug("函数定义")
		return base.SPAFunction{}
	}
}

func (v *ExpVisitor) VisitOr_test(ctx *parser.Or_testContext) interface{} {
	log.Debug("Visit Or_Test")

	var result interface{}
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitAnd_test(ctx.GetChild(0).(*parser.And_testContext))
		if isTrue(result){
			return result
		}
	}
	return result
}

func (v *ExpVisitor) VisitAnd_test(ctx *parser.And_testContext) interface{} {
	log.Debug("Visit And_Test")

	var result interface{}
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		result = v.VisitNot_test(ctx.GetChild(i).(*parser.Not_testContext))
		if !isTrue(result) {
			return result
		}
	}
	return result
}

func (v *ExpVisitor) VisitNot_test(ctx *parser.Not_testContext) interface{} {
	log.Debug("Visit Not_Test")

	if ctx.GetChildCount() == 2 {
		result := v.VisitNot_test(ctx.GetChild(1).(*parser.Not_testContext))
		return getInvert(result)
	} else {
		return v.VisitComparison(ctx.GetChild(0).(*parser.Compare_exprContext))
	}
}

func (v *ExpVisitor) VisitComparison(ctx *parser.Compare_exprContext) interface{} {
	log.Debug("Visit Comparison")

	// 目前先不支持比较，只返回表达式
	return v.VisitArith_expr(ctx.GetChild(0).(*parser.Arith_exprContext))
}

func (v *ExpVisitor) VisitArith_expr(ctx *parser.Arith_exprContext) interface{} {
	log.Debug("Visit Arith_Expr")

	if ctx.GetChildCount() == 1 {
		return v.VisitTerm(ctx.GetChild(0).(*parser.TermContext))
	}else{
		var termResult = 0.0 //默认为0，累加的时候不需要判断了
		var op = "+" // 统一处理
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitTerm(ctx.GetChild(i).(*parser.TermContext))
				if r == nil {
					return nil
				}
				termResult = arithmetic(termResult, r.(float64), op)
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *ExpVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	log.Debug("Visit Term")

	//r := v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	if ctx.GetChildCount() == 1 {
		return v.VisitFactor(ctx.GetChild(0).(*parser.FactorContext))
	} else {
		var termResult = 1.0 //默认为1方便处理
		var op = "*"
		for i := 0; i < ctx.GetChildCount(); i++ {
			if i % 2 == 0 {
				r := v.VisitFactor(ctx.GetChild(i).(*parser.FactorContext))
				if r == nil{
					return nil
				}
				switch r.(type) {
				case float64:
					termResult = arithmetic(termResult, r.(float64), op)
				default:
					panic("类型无效")
				}
			} else {
				op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
			}
		}
		return termResult
	}
}

func (v *ExpVisitor) VisitFactor(ctx *parser.FactorContext) interface{} {
	log.Debug("Visit Factor")
	//return v.VisitChildren(ctx)
	if ctx.GetChildCount() == 1 {
		return v.VisitPower(ctx.GetChild(0).(*parser.PowerContext))
	} else {
		op := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
		r := v.VisitFactor(ctx.GetChild(1).(*parser.FactorContext)).(float64)
		return arithmetic(0.0, r, op)
	}
}

func (v *ExpVisitor) VisitPower(ctx *parser.PowerContext) interface{} {
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

func (v *ExpVisitor) VisitAtom_expr(ctx *parser.Atom_exprContext) interface{} {
	log.Debug("Visit Atom_Expr")

	if ctx.GetChildCount() == 3 {
		return v.VisitPostfix_expr(ctx.GetChild(1).(*parser.Postfix_exprContext))
	}
	if ctx.GetChildCount() == 4 {
		name := ctx.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
		args := v.VisitArg_list(ctx.GetChild(2).(*parser.Arg_listContext))
		return callInternalFunc(name, args.([]interface{}))
	}

	//log.Debug(v.vars)
	//终结符
	terminalNode := ctx.GetChild(0).(*antlr.TerminalNodeImpl)
	tt := terminalNode.GetSymbol().GetTokenType()
	//log.Debug(terminalNode.GetText())
	switch tt {
	case parser.SpartaLexerIDENTIFIER:
		return v.getVar(terminalNode.GetText()).(float64)
	case parser.SpartaLexerNUMBER_LITERAL:
		return parseNumber(terminalNode.GetText())
	case parser.SpartaLexerSTRING:
		return escapeString(terminalNode.GetText())
	default:
		panic("类型无效")
	}
}

func (v *ExpVisitor) VisitArg_list(ctx *parser.Arg_listContext) interface{} {
	//return v.VisitChildren(ctx)
	argList := make([]interface{}, 0, 10)
	for i := 0; i < ctx.GetChildCount(); i += 2 {
		v := v.VisitArgument(ctx.GetChild(i).(*parser.ArgumentContext))
		argList = append(argList, v)
	}
	return argList
}

func (v *ExpVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	return v.VisitPostfix_expr(ctx.GetChild(0).(*parser.Postfix_exprContext))
}

func callInternalFunc(name string, args []interface{}) interface{} {
	switch name {
	case "print":
		fmt.Println(args...)
		return nil
	default:
		panic("function not found")
	}
}

func(v *ExpVisitor) putVar(name string, value interface{})  {
	v.vars[name] = value
}

func(v *ExpVisitor) getVar(name string) interface{}  {
	return v.vars[name]
}

/**
判断数据是否为真
 */
func isTrue(v interface{}) bool  {
	switch v.(type) {
	case bool:
		return v.(bool)
	case float64:
		f := v.(float64)
		return math.Floor(f) > 0
	default:
		return false
	}
}

/**
数据取反
 */
func getInvert(v interface{}) bool {
	r := isTrue(v)
	return !r
}

/**
算数运算
 */
func arithmetic(first, second float64, op string) float64  {
	switch op {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		panic("不支持的操作")
	}
}

/**
解析数字
 */
func parseNumber(text string) float64 {
	r, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic("数字转化错误")
	}
	return r
}

/**
转义字符串（状态机）
 */
func escapeString(text string) string {
	var result = make([]rune, 0 , len(text))
	escaped := false //是否处于转义状态
	for _, c := range text {
		if escaped {
			result = append(result, escapeMapping(c))
			escaped = false //关闭转义状态
		}else{
			switch c {
			case '"':
				continue
			case '\\':
				escaped = true
			default:
				result = append(result, c)
			}
		}
	}
	s := string(result)
	return s
}

/**
转义字符映射
 */
func escapeMapping(c rune) rune {
	switch c {
	case '"', '\\':
		return c
	case 't':
		return '\t'
	case 'r':
		return '\r'
	case 'n':
		return '\n'
	default:
		return c
	}
}