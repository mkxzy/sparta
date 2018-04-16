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

// EnterStmt_list is called when production stmt_list is entered.
func (s *BaseSpartaListener) EnterStmt_list(ctx *Stmt_listContext) {}

// ExitStmt_list is called when production stmt_list is exited.
func (s *BaseSpartaListener) ExitStmt_list(ctx *Stmt_listContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSpartaListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSpartaListener) ExitStmt(ctx *StmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BaseSpartaListener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BaseSpartaListener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseSpartaListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseSpartaListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterTest is called when production test is entered.
func (s *BaseSpartaListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BaseSpartaListener) ExitTest(ctx *TestContext) {}

// EnterOr_test is called when production or_test is entered.
func (s *BaseSpartaListener) EnterOr_test(ctx *Or_testContext) {}

// ExitOr_test is called when production or_test is exited.
func (s *BaseSpartaListener) ExitOr_test(ctx *Or_testContext) {}

// EnterAnd_test is called when production and_test is entered.
func (s *BaseSpartaListener) EnterAnd_test(ctx *And_testContext) {}

// ExitAnd_test is called when production and_test is exited.
func (s *BaseSpartaListener) ExitAnd_test(ctx *And_testContext) {}

// EnterNot_test is called when production not_test is entered.
func (s *BaseSpartaListener) EnterNot_test(ctx *Not_testContext) {}

// ExitNot_test is called when production not_test is exited.
func (s *BaseSpartaListener) ExitNot_test(ctx *Not_testContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSpartaListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSpartaListener) ExitComparison(ctx *ComparisonContext) {}

// EnterComp_op is called when production comp_op is entered.
func (s *BaseSpartaListener) EnterComp_op(ctx *Comp_opContext) {}

// ExitComp_op is called when production comp_op is exited.
func (s *BaseSpartaListener) ExitComp_op(ctx *Comp_opContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSpartaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSpartaListener) ExitExpr(ctx *ExprContext) {}

// EnterArith_expr is called when production arith_expr is entered.
func (s *BaseSpartaListener) EnterArith_expr(ctx *Arith_exprContext) {}

// ExitArith_expr is called when production arith_expr is exited.
func (s *BaseSpartaListener) ExitArith_expr(ctx *Arith_exprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSpartaListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSpartaListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseSpartaListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseSpartaListener) ExitFactor(ctx *FactorContext) {}

// EnterPower is called when production power is entered.
func (s *BaseSpartaListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseSpartaListener) ExitPower(ctx *PowerContext) {}

// EnterAtom_expr is called when production atom_expr is entered.
func (s *BaseSpartaListener) EnterAtom_expr(ctx *Atom_exprContext) {}

// ExitAtom_expr is called when production atom_expr is exited.
func (s *BaseSpartaListener) ExitAtom_expr(ctx *Atom_exprContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSpartaListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSpartaListener) ExitAtom(ctx *AtomContext) {}

// EnterTrailer is called when production trailer is entered.
func (s *BaseSpartaListener) EnterTrailer(ctx *TrailerContext) {}

// ExitTrailer is called when production trailer is exited.
func (s *BaseSpartaListener) ExitTrailer(ctx *TrailerContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BaseSpartaListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BaseSpartaListener) ExitArg_list(ctx *Arg_listContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseSpartaListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseSpartaListener) ExitArgument(ctx *ArgumentContext) {}

// EnterTestlist_comp is called when production testlist_comp is entered.
func (s *BaseSpartaListener) EnterTestlist_comp(ctx *Testlist_compContext) {}

// ExitTestlist_comp is called when production testlist_comp is exited.
func (s *BaseSpartaListener) ExitTestlist_comp(ctx *Testlist_compContext) {}
