// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpartaListener is a complete listener for a parse tree produced by SpartaParser.
type SpartaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatList is called when entering the statList production.
	EnterStatList(c *StatListContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterVarStat is called when entering the varStat production.
	EnterVarStat(c *VarStatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterBasicLit is called when entering the basicLit production.
	EnterBasicLit(c *BasicLitContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatList is called when exiting the statList production.
	ExitStatList(c *StatListContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitVarStat is called when exiting the varStat production.
	ExitVarStat(c *VarStatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitBasicLit is called when exiting the basicLit production.
	ExitBasicLit(c *BasicLitContext)
}
