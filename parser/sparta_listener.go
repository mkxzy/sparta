// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpartaListener is a complete listener for a parse tree produced by SpartaParser.
type SpartaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterPrimary_expr is called when entering the primary_expr production.
	EnterPrimary_expr(c *Primary_exprContext)

	// EnterPostfix_expr is called when entering the postfix_expr production.
	EnterPostfix_expr(c *Postfix_exprContext)

	// EnterOr_test is called when entering the or_test production.
	EnterOr_test(c *Or_testContext)

	// EnterAnd_test is called when entering the and_test production.
	EnterAnd_test(c *And_testContext)

	// EnterNot_test is called when entering the not_test production.
	EnterNot_test(c *Not_testContext)

	// EnterCompare_expr is called when entering the compare_expr production.
	EnterCompare_expr(c *Compare_exprContext)

	// EnterComp_op is called when entering the comp_op production.
	EnterComp_op(c *Comp_opContext)

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

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitPrimary_expr is called when exiting the primary_expr production.
	ExitPrimary_expr(c *Primary_exprContext)

	// ExitPostfix_expr is called when exiting the postfix_expr production.
	ExitPostfix_expr(c *Postfix_exprContext)

	// ExitOr_test is called when exiting the or_test production.
	ExitOr_test(c *Or_testContext)

	// ExitAnd_test is called when exiting the and_test production.
	ExitAnd_test(c *And_testContext)

	// ExitNot_test is called when exiting the not_test production.
	ExitNot_test(c *Not_testContext)

	// ExitCompare_expr is called when exiting the compare_expr production.
	ExitCompare_expr(c *Compare_exprContext)

	// ExitComp_op is called when exiting the comp_op production.
	ExitComp_op(c *Comp_opContext)

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

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)
}
