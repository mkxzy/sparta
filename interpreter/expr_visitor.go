package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"math"
	"strconv"
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

	log.Debug("Visit Program")
	if(ctx.GetChildCount() == 2){
		v.VisitStatList(ctx.GetChild(0).(*parser.Stmt_listContext))
		log.Debug(v.vars)
	}
	return nil
}

func (v *ExpVisitor) VisitStatList(ctx *parser.Stmt_listContext) interface{} {
	log.Debug("Visit Stmt_List")
	for i := 0; i < ctx.GetChildCount(); i++ {
		v.VisitStat(ctx.GetChild(i).(*parser.StmtContext))
	}
	return nil
}

func (v *ExpVisitor) VisitStat(ctx *parser.StmtContext) interface{} {
	log.Debug("Visit Stmt")
	v.VisitExpr_stmt(ctx.GetChild(0).(*parser.Expr_stmtContext))
	return nil
}

func (v *ExpVisitor) VisitExpr_stmt(ctx *parser.Expr_stmtContext) interface{} {
	log.Debug("Visit Expr_Stmt")
	v.VisitAssign_stmt(ctx.GetChild(0).(*parser.Assign_stmtContext))
	return nil
}

func (v *ExpVisitor) VisitAssign_stmt(ctx *parser.Assign_stmtContext) interface{} {
	//return v.VisitChildren(ctx)
	log.Debug("Visit Assign_Stmt")
	token := ctx.GetToken(parser.SpartaParserIDENTIFIER, 0)
	if token == nil {
		panic("标识符不正确")
	}
	name := token.GetText()
	value := v.VisitTest(ctx.GetChild(2).(*parser.TestContext))
	v.putVar(name, value)
	log.Debugf("%s = %v", name, value)
	return nil
}

func (v *ExpVisitor) VisitTest(ctx *parser.TestContext) interface{} {
	log.Debug("Visit Test")

	return v.VisitOr_test(ctx.GetChild(0).(*parser.Or_testContext))
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
		return v.VisitComparison(ctx.GetChild(0).(*parser.ComparisonContext))
	}
}

func (v *ExpVisitor) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	log.Debug("Visit Comparison")

	// 目前先不支持比较，只返回表达式
	return v.VisitExpr(ctx.GetChild(0).(*parser.ExprContext))
}

func (v *ExpVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	log.Debug("Visit Expr")

	return v.VisitArith_expr(ctx.GetChild(0).(*parser.Arith_exprContext))
}

func (v *ExpVisitor) VisitArith_expr(ctx *parser.Arith_exprContext) interface{} {
	log.Debug("Visit Arith_Expr")

	//termResult := v.VisitTerm(ctx.GetChild(0).(*parser.TermContext)).(float64)
	//log.Debug(termResult)

	var termResult = 0.0 //默认为0，累加的时候不需要判断了
	var op = "+" // 统一处理
	for i := 0; i < ctx.GetChildCount(); i++ {
		if i % 2 == 0 {
			r := v.VisitTerm(ctx.GetChild(i).(*parser.TermContext)).(float64)
			termResult = arithmetic(termResult, r, op)
		} else {
			op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
		}
	}
	return termResult
}

func (v *ExpVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	log.Debug("Visit Term")

	var termResult = 1.0 //默认为1方便处理
	var op = "*"
	for i := 0; i < ctx.GetChildCount(); i++ {
		if i % 2 == 0 {
			r := v.VisitFactor(ctx.GetChild(i).(*parser.FactorContext)).(float64)
			termResult = arithmetic(termResult, r, op)
		} else {
			op = ctx.GetChild(i).(*antlr.TerminalNodeImpl).GetText()
		}
	}
	return termResult
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

	atom := v.VisitAtom_expr(ctx.GetChild(0).(*parser.Atom_exprContext)).(float64)
	if ctx.GetChildCount() == 1 {
		return atom
	} else {
		factor := v.VisitFactor(ctx.GetChild(2).(*parser.FactorContext)).(float64)
		return math.Pow(atom, factor)
	}
}

func (v *ExpVisitor) VisitAtom_expr(ctx *parser.Atom_exprContext) interface{} {
	log.Debug("Visit Atom_Expr")

	return v.VisitAtom(ctx.GetChild(0).(*parser.AtomContext))
}

func (v *ExpVisitor) VisitAtom(ctx *parser.AtomContext) interface{} {
	log.Debug("Visit Atom")

	if ctx.GetChildCount() == 3 {
		return v.VisitTestlist_comp(ctx.GetChild(1).(*parser.Testlist_compContext))
	} else {
		terminalNode := ctx.GetChild(0).(*antlr.TerminalNodeImpl)
		tt := terminalNode.GetSymbol().GetTokenType()
		log.Debug(terminalNode.GetText())
		switch tt {
		case parser.SpartaLexerIDENTIFIER:
			return v.getVar(terminalNode.GetText()).(float64)
		case parser.SpartaLexerNUMBER_LITERAL:
			return parseNumber(terminalNode.GetText())
		default:
			panic("类型无效")
		}
	}
}

func (v *ExpVisitor) VisitTestlist_comp(ctx *parser.Testlist_compContext) interface{} {
	log.Debug("Visit Test_Comp")

	return v.VisitTest(ctx.GetChild(0).(*parser.TestContext))
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