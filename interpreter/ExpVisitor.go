package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("ExpVisitor")

type ExpVisitor struct {
	*parser.BaseSpartaVisitor
}

func NewExpVisitor() *ExpVisitor {
	return &ExpVisitor{}
}

func (v *ExpVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	//defer func () {
	//	if err := recover(); err != nil{
	//		log.Debug(err)
	//	}
	//}()
	//return v.VisitChildren(ctx)
	if(ctx.GetChildCount() == 2){
		v.VisitStatList(ctx.GetChild(0).(*parser.Stmt_listContext))
	}
	return nil
}

func (v *ExpVisitor) VisitStatList(ctx *parser.Stmt_listContext) interface{} {
	log.Debug("Visit StatList")
	for i:=0; i<ctx.GetChildCount(); i++ {
		v.VisitStat(ctx.GetChild(i).(*parser.StmtContext))
	}
	return nil
}

func (v *ExpVisitor) VisitStat(ctx *parser.StmtContext) interface{} {
	log.Debug("Visit Stat")
	v.VisitVarStat(ctx.GetChild(0).(*parser.Var_stmtContext))
	return nil
}

func (v *ExpVisitor) VisitVarStat(ctx *parser.Var_stmtContext) interface{} {

	log.Debug("Visit VarStat")
	token := ctx.GetToken(parser.SpartaParserIDENTIFIER, 0)
	if token == nil {
		panic("标识符不正确")
	}
	name := ctx.GetToken(parser.SpartaParserIDENTIFIER, 0).GetText()
	value := v.VisitExpr(ctx.GetChild(3).(*parser.ExprContext))
	log.Debugf("%s = %v", name, value)
	return nil
}

func (v *ExpVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	//i, err := strconv.ParseFloat(ctx.GetToken(parser.SpartaParserNUMBER_LITERAL, 0).GetText(), 32)
	//if err != nil {
	//	//log.Debugf("数字解析异常")
	//	panic("数字解析异常")
	//}
	//return i
	return nil
}