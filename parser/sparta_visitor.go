// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SpartaParser.
type SpartaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SpartaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SpartaParser#statList.
	VisitStatList(ctx *StatListContext) interface{}

	// Visit a parse tree produced by SpartaParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by SpartaParser#varStat.
	VisitVarStat(ctx *VarStatContext) interface{}

	// Visit a parse tree produced by SpartaParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SpartaParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by SpartaParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by SpartaParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by SpartaParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SpartaParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by SpartaParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}
}
