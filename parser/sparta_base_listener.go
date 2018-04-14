// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSpartaListener is a complete listener for a parse tree produced by SpartaParser.
type BaseSpartaListener struct{}

var _ SpartaListener = &BaseSpartaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSpartaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSpartaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSpartaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSpartaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSpartaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSpartaListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatList is called when production statList is entered.
func (s *BaseSpartaListener) EnterStatList(ctx *StatListContext) {}

// ExitStatList is called when production statList is exited.
func (s *BaseSpartaListener) ExitStatList(ctx *StatListContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseSpartaListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseSpartaListener) ExitStat(ctx *StatContext) {}

// EnterVarStat is called when production varStat is entered.
func (s *BaseSpartaListener) EnterVarStat(ctx *VarStatContext) {}

// ExitVarStat is called when production varStat is exited.
func (s *BaseSpartaListener) ExitVarStat(ctx *VarStatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSpartaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSpartaListener) ExitExpr(ctx *ExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseSpartaListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseSpartaListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseSpartaListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseSpartaListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseSpartaListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseSpartaListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSpartaListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSpartaListener) ExitLiteral(ctx *LiteralContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseSpartaListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseSpartaListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterBasicLit is called when production basicLit is entered.
func (s *BaseSpartaListener) EnterBasicLit(ctx *BasicLitContext) {}

// ExitBasicLit is called when production basicLit is exited.
func (s *BaseSpartaListener) ExitBasicLit(ctx *BasicLitContext) {}
