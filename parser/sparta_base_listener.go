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

// EnterExp is called when production exp is entered.
func (s *BaseSpartaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseSpartaListener) ExitExp(ctx *ExpContext) {}
