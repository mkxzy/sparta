// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpartaListener is a complete listener for a parse tree produced by SpartaParser.
type SpartaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmt_list is called when entering the stmt_list production.
	EnterStmt_list(c *Stmt_listContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterOr_test is called when entering the or_test production.
	EnterOr_test(c *Or_testContext)

	// EnterAnd_test is called when entering the and_test production.
	EnterAnd_test(c *And_testContext)

	// EnterNot_test is called when entering the not_test production.
	EnterNot_test(c *Not_testContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterArith_expr is called when entering the arith_expr production.
	EnterArith_expr(c *Arith_exprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterAtom_expr is called when entering the atom_expr production.
	EnterAtom_expr(c *Atom_exprContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterTrailer is called when entering the trailer production.
	EnterTrailer(c *TrailerContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterTestlist_comp is called when entering the testlist_comp production.
	EnterTestlist_comp(c *Testlist_compContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmt_list is called when exiting the stmt_list production.
	ExitStmt_list(c *Stmt_listContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitOr_test is called when exiting the or_test production.
	ExitOr_test(c *Or_testContext)

	// ExitAnd_test is called when exiting the and_test production.
	ExitAnd_test(c *And_testContext)

	// ExitNot_test is called when exiting the not_test production.
	ExitNot_test(c *Not_testContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitArith_expr is called when exiting the arith_expr production.
	ExitArith_expr(c *Arith_exprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitAtom_expr is called when exiting the atom_expr production.
	ExitAtom_expr(c *Atom_exprContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitTrailer is called when exiting the trailer production.
	ExitTrailer(c *TrailerContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitTestlist_comp is called when exiting the testlist_comp production.
	ExitTestlist_comp(c *Testlist_compContext)
}
