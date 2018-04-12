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

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatList is called when exiting the statList production.
	ExitStatList(c *StatListContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitVarStat is called when exiting the varStat production.
	ExitVarStat(c *VarStatContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)
}
