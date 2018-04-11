// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpartaListener is a complete listener for a parse tree produced by SpartaParser.
type SpartaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterVarstat is called when entering the varstat production.
	EnterVarstat(c *VarstatContext)

	// EnterAssignstat is called when entering the assignstat production.
	EnterAssignstat(c *AssignstatContext)

	// EnterBreakstat is called when entering the breakstat production.
	EnterBreakstat(c *BreakstatContext)

	// EnterIfstat is called when entering the ifstat production.
	EnterIfstat(c *IfstatContext)

	// EnterWhilestat is called when entering the whilestat production.
	EnterWhilestat(c *WhilestatContext)

	// EnterForstat is called when entering the forstat production.
	EnterForstat(c *ForstatContext)

	// EnterRetstat is called when entering the retstat production.
	EnterRetstat(c *RetstatContext)

	// EnterFundef is called when entering the fundef production.
	EnterFundef(c *FundefContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterVarlist is called when entering the varlist production.
	EnterVarlist(c *VarlistContext)

	// EnterNamelist is called when entering the namelist production.
	EnterNamelist(c *NamelistContext)

	// EnterExplist is called when entering the explist production.
	EnterExplist(c *ExplistContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterPrefixexp is called when entering the prefixexp production.
	EnterPrefixexp(c *PrefixexpContext)

	// EnterFuncall is called when entering the funcall production.
	EnterFuncall(c *FuncallContext)

	// EnterVarOrExp is called when entering the varOrExp production.
	EnterVarOrExp(c *VarOrExpContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVarSuffix is called when entering the varSuffix production.
	EnterVarSuffix(c *VarSuffixContext)

	// EnterNameAndArgs is called when entering the nameAndArgs production.
	EnterNameAndArgs(c *NameAndArgsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterFunexp is called when entering the funexp production.
	EnterFunexp(c *FunexpContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterParlist is called when entering the parlist production.
	EnterParlist(c *ParlistContext)

	// EnterTableconstructor is called when entering the tableconstructor production.
	EnterTableconstructor(c *TableconstructorContext)

	// EnterFieldlist is called when entering the fieldlist production.
	EnterFieldlist(c *FieldlistContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldsep is called when entering the fieldsep production.
	EnterFieldsep(c *FieldsepContext)

	// EnterOperatorOr is called when entering the operatorOr production.
	EnterOperatorOr(c *OperatorOrContext)

	// EnterOperatorAnd is called when entering the operatorAnd production.
	EnterOperatorAnd(c *OperatorAndContext)

	// EnterOperatorComparison is called when entering the operatorComparison production.
	EnterOperatorComparison(c *OperatorComparisonContext)

	// EnterOperatorStrcat is called when entering the operatorStrcat production.
	EnterOperatorStrcat(c *OperatorStrcatContext)

	// EnterOperatorAddSub is called when entering the operatorAddSub production.
	EnterOperatorAddSub(c *OperatorAddSubContext)

	// EnterOperatorMulDivMod is called when entering the operatorMulDivMod production.
	EnterOperatorMulDivMod(c *OperatorMulDivModContext)

	// EnterOperatorBitwise is called when entering the operatorBitwise production.
	EnterOperatorBitwise(c *OperatorBitwiseContext)

	// EnterOperatorUnary is called when entering the operatorUnary production.
	EnterOperatorUnary(c *OperatorUnaryContext)

	// EnterOperatorPower is called when entering the operatorPower production.
	EnterOperatorPower(c *OperatorPowerContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitVarstat is called when exiting the varstat production.
	ExitVarstat(c *VarstatContext)

	// ExitAssignstat is called when exiting the assignstat production.
	ExitAssignstat(c *AssignstatContext)

	// ExitBreakstat is called when exiting the breakstat production.
	ExitBreakstat(c *BreakstatContext)

	// ExitIfstat is called when exiting the ifstat production.
	ExitIfstat(c *IfstatContext)

	// ExitWhilestat is called when exiting the whilestat production.
	ExitWhilestat(c *WhilestatContext)

	// ExitForstat is called when exiting the forstat production.
	ExitForstat(c *ForstatContext)

	// ExitRetstat is called when exiting the retstat production.
	ExitRetstat(c *RetstatContext)

	// ExitFundef is called when exiting the fundef production.
	ExitFundef(c *FundefContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitVarlist is called when exiting the varlist production.
	ExitVarlist(c *VarlistContext)

	// ExitNamelist is called when exiting the namelist production.
	ExitNamelist(c *NamelistContext)

	// ExitExplist is called when exiting the explist production.
	ExitExplist(c *ExplistContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitPrefixexp is called when exiting the prefixexp production.
	ExitPrefixexp(c *PrefixexpContext)

	// ExitFuncall is called when exiting the funcall production.
	ExitFuncall(c *FuncallContext)

	// ExitVarOrExp is called when exiting the varOrExp production.
	ExitVarOrExp(c *VarOrExpContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVarSuffix is called when exiting the varSuffix production.
	ExitVarSuffix(c *VarSuffixContext)

	// ExitNameAndArgs is called when exiting the nameAndArgs production.
	ExitNameAndArgs(c *NameAndArgsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitFunexp is called when exiting the funexp production.
	ExitFunexp(c *FunexpContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitParlist is called when exiting the parlist production.
	ExitParlist(c *ParlistContext)

	// ExitTableconstructor is called when exiting the tableconstructor production.
	ExitTableconstructor(c *TableconstructorContext)

	// ExitFieldlist is called when exiting the fieldlist production.
	ExitFieldlist(c *FieldlistContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldsep is called when exiting the fieldsep production.
	ExitFieldsep(c *FieldsepContext)

	// ExitOperatorOr is called when exiting the operatorOr production.
	ExitOperatorOr(c *OperatorOrContext)

	// ExitOperatorAnd is called when exiting the operatorAnd production.
	ExitOperatorAnd(c *OperatorAndContext)

	// ExitOperatorComparison is called when exiting the operatorComparison production.
	ExitOperatorComparison(c *OperatorComparisonContext)

	// ExitOperatorStrcat is called when exiting the operatorStrcat production.
	ExitOperatorStrcat(c *OperatorStrcatContext)

	// ExitOperatorAddSub is called when exiting the operatorAddSub production.
	ExitOperatorAddSub(c *OperatorAddSubContext)

	// ExitOperatorMulDivMod is called when exiting the operatorMulDivMod production.
	ExitOperatorMulDivMod(c *OperatorMulDivModContext)

	// ExitOperatorBitwise is called when exiting the operatorBitwise production.
	ExitOperatorBitwise(c *OperatorBitwiseContext)

	// ExitOperatorUnary is called when exiting the operatorUnary production.
	ExitOperatorUnary(c *OperatorUnaryContext)

	// ExitOperatorPower is called when exiting the operatorPower production.
	ExitOperatorPower(c *OperatorPowerContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)
}
