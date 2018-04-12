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

	// Visit a parse tree produced by SpartaParser#exp.
	VisitExp(ctx *ExpContext) interface{}
}
