// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSpartaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSpartaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitStmt_list(ctx *Stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitExpr_stmt(ctx *Expr_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitVar_stmt(ctx *Var_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitOr_test(ctx *Or_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitAnd_test(ctx *And_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitNot_test(ctx *Not_testContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitComp_op(ctx *Comp_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitXor_expr(ctx *Xor_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitAnd_expr(ctx *And_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitShift_expr(ctx *Shift_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitArith_expr(ctx *Arith_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitAtom_expr(ctx *Atom_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}
