// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSpartaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSpartaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitStatList(ctx *StatListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitVarStat(ctx *VarStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitBasicLit(ctx *BasicLitContext) interface{} {
	return v.VisitChildren(ctx)
}
