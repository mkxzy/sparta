// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SpartaParser.
type SpartaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SpartaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SpartaParser#stmt_list.
	VisitStmt_list(ctx *Stmt_listContext) interface{}

	// Visit a parse tree produced by SpartaParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by SpartaParser#expr_stmt.
	VisitExpr_stmt(ctx *Expr_stmtContext) interface{}

	// Visit a parse tree produced by SpartaParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by SpartaParser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by SpartaParser#or_test.
	VisitOr_test(ctx *Or_testContext) interface{}

	// Visit a parse tree produced by SpartaParser#and_test.
	VisitAnd_test(ctx *And_testContext) interface{}

	// Visit a parse tree produced by SpartaParser#not_test.
	VisitNot_test(ctx *Not_testContext) interface{}

	// Visit a parse tree produced by SpartaParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by SpartaParser#comp_op.
	VisitComp_op(ctx *Comp_opContext) interface{}

	// Visit a parse tree produced by SpartaParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SpartaParser#arith_expr.
	VisitArith_expr(ctx *Arith_exprContext) interface{}

	// Visit a parse tree produced by SpartaParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SpartaParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by SpartaParser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by SpartaParser#atom_expr.
	VisitAtom_expr(ctx *Atom_exprContext) interface{}

	// Visit a parse tree produced by SpartaParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by SpartaParser#trailer.
	VisitTrailer(ctx *TrailerContext) interface{}

	// Visit a parse tree produced by SpartaParser#arg_list.
	VisitArg_list(ctx *Arg_listContext) interface{}

	// Visit a parse tree produced by SpartaParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by SpartaParser#testlist_comp.
	VisitTestlist_comp(ctx *Testlist_compContext) interface{}
}
