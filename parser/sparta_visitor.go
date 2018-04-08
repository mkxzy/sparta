// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SpartaParser.
type SpartaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SpartaParser#chunk.
	VisitChunk(ctx *ChunkContext) interface{}

	// Visit a parse tree produced by SpartaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by SpartaParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by SpartaParser#retstat.
	VisitRetstat(ctx *RetstatContext) interface{}

	// Visit a parse tree produced by SpartaParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by SpartaParser#funcname.
	VisitFuncname(ctx *FuncnameContext) interface{}

	// Visit a parse tree produced by SpartaParser#varlist.
	VisitVarlist(ctx *VarlistContext) interface{}

	// Visit a parse tree produced by SpartaParser#namelist.
	VisitNamelist(ctx *NamelistContext) interface{}

	// Visit a parse tree produced by SpartaParser#explist.
	VisitExplist(ctx *ExplistContext) interface{}

	// Visit a parse tree produced by SpartaParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by SpartaParser#prefixexp.
	VisitPrefixexp(ctx *PrefixexpContext) interface{}

	// Visit a parse tree produced by SpartaParser#functioncall.
	VisitFunctioncall(ctx *FunctioncallContext) interface{}

	// Visit a parse tree produced by SpartaParser#varOrExp.
	VisitVarOrExp(ctx *VarOrExpContext) interface{}

	// Visit a parse tree produced by SpartaParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SpartaParser#varSuffix.
	VisitVarSuffix(ctx *VarSuffixContext) interface{}

	// Visit a parse tree produced by SpartaParser#nameAndArgs.
	VisitNameAndArgs(ctx *NameAndArgsContext) interface{}

	// Visit a parse tree produced by SpartaParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by SpartaParser#functiondef.
	VisitFunctiondef(ctx *FunctiondefContext) interface{}

	// Visit a parse tree produced by SpartaParser#funcbody.
	VisitFuncbody(ctx *FuncbodyContext) interface{}

	// Visit a parse tree produced by SpartaParser#parlist.
	VisitParlist(ctx *ParlistContext) interface{}

	// Visit a parse tree produced by SpartaParser#tableconstructor.
	VisitTableconstructor(ctx *TableconstructorContext) interface{}

	// Visit a parse tree produced by SpartaParser#fieldlist.
	VisitFieldlist(ctx *FieldlistContext) interface{}

	// Visit a parse tree produced by SpartaParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by SpartaParser#fieldsep.
	VisitFieldsep(ctx *FieldsepContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorOr.
	VisitOperatorOr(ctx *OperatorOrContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorAnd.
	VisitOperatorAnd(ctx *OperatorAndContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorComparison.
	VisitOperatorComparison(ctx *OperatorComparisonContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorStrcat.
	VisitOperatorStrcat(ctx *OperatorStrcatContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorAddSub.
	VisitOperatorAddSub(ctx *OperatorAddSubContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorMulDivMod.
	VisitOperatorMulDivMod(ctx *OperatorMulDivModContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorBitwise.
	VisitOperatorBitwise(ctx *OperatorBitwiseContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorUnary.
	VisitOperatorUnary(ctx *OperatorUnaryContext) interface{}

	// Visit a parse tree produced by SpartaParser#operatorPower.
	VisitOperatorPower(ctx *OperatorPowerContext) interface{}

	// Visit a parse tree produced by SpartaParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by SpartaParser#str.
	VisitStr(ctx *StrContext) interface{}
}
