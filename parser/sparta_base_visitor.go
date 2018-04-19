// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSpartaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSpartaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitSimple_stmt(ctx *Simple_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitExpr_stmt(ctx *Expr_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitPrimary_expr(ctx *Primary_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitPostfix_expr(ctx *Postfix_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitPar_seq(ctx *Par_seqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitParlist(ctx *ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitNamelist(ctx *NamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitBlock(ctx *BlockContext) interface{} {
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

func (v *BaseSpartaVisitor) VisitCompare_expr(ctx *Compare_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitComp_op(ctx *Comp_opContext) interface{} {
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

func (v *BaseSpartaVisitor) VisitArg_seq(ctx *Arg_seqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitArg_list(ctx *Arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSpartaVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}
