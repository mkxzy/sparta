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

// EnterBlock is called when production block is entered.
func (s *BaseSpartaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSpartaListener) ExitBlock(ctx *BlockContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseSpartaListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseSpartaListener) ExitStat(ctx *StatContext) {}

// EnterVarstat is called when production varstat is entered.
func (s *BaseSpartaListener) EnterVarstat(ctx *VarstatContext) {}

// ExitVarstat is called when production varstat is exited.
func (s *BaseSpartaListener) ExitVarstat(ctx *VarstatContext) {}

// EnterAssignstat is called when production assignstat is entered.
func (s *BaseSpartaListener) EnterAssignstat(ctx *AssignstatContext) {}

// ExitAssignstat is called when production assignstat is exited.
func (s *BaseSpartaListener) ExitAssignstat(ctx *AssignstatContext) {}

// EnterBreakstat is called when production breakstat is entered.
func (s *BaseSpartaListener) EnterBreakstat(ctx *BreakstatContext) {}

// ExitBreakstat is called when production breakstat is exited.
func (s *BaseSpartaListener) ExitBreakstat(ctx *BreakstatContext) {}

// EnterIfstat is called when production ifstat is entered.
func (s *BaseSpartaListener) EnterIfstat(ctx *IfstatContext) {}

// ExitIfstat is called when production ifstat is exited.
func (s *BaseSpartaListener) ExitIfstat(ctx *IfstatContext) {}

// EnterWhilestat is called when production whilestat is entered.
func (s *BaseSpartaListener) EnterWhilestat(ctx *WhilestatContext) {}

// ExitWhilestat is called when production whilestat is exited.
func (s *BaseSpartaListener) ExitWhilestat(ctx *WhilestatContext) {}

// EnterForstat is called when production forstat is entered.
func (s *BaseSpartaListener) EnterForstat(ctx *ForstatContext) {}

// ExitForstat is called when production forstat is exited.
func (s *BaseSpartaListener) ExitForstat(ctx *ForstatContext) {}

// EnterRetstat is called when production retstat is entered.
func (s *BaseSpartaListener) EnterRetstat(ctx *RetstatContext) {}

// ExitRetstat is called when production retstat is exited.
func (s *BaseSpartaListener) ExitRetstat(ctx *RetstatContext) {}

// EnterFundef is called when production fundef is entered.
func (s *BaseSpartaListener) EnterFundef(ctx *FundefContext) {}

// ExitFundef is called when production fundef is exited.
func (s *BaseSpartaListener) ExitFundef(ctx *FundefContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseSpartaListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseSpartaListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterVarlist is called when production varlist is entered.
func (s *BaseSpartaListener) EnterVarlist(ctx *VarlistContext) {}

// ExitVarlist is called when production varlist is exited.
func (s *BaseSpartaListener) ExitVarlist(ctx *VarlistContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseSpartaListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseSpartaListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExplist is called when production explist is entered.
func (s *BaseSpartaListener) EnterExplist(ctx *ExplistContext) {}

// ExitExplist is called when production explist is exited.
func (s *BaseSpartaListener) ExitExplist(ctx *ExplistContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseSpartaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseSpartaListener) ExitExp(ctx *ExpContext) {}

// EnterPrefixexp is called when production prefixexp is entered.
func (s *BaseSpartaListener) EnterPrefixexp(ctx *PrefixexpContext) {}

// ExitPrefixexp is called when production prefixexp is exited.
func (s *BaseSpartaListener) ExitPrefixexp(ctx *PrefixexpContext) {}

// EnterFuncall is called when production funcall is entered.
func (s *BaseSpartaListener) EnterFuncall(ctx *FuncallContext) {}

// ExitFuncall is called when production funcall is exited.
func (s *BaseSpartaListener) ExitFuncall(ctx *FuncallContext) {}

// EnterVarOrExp is called when production varOrExp is entered.
func (s *BaseSpartaListener) EnterVarOrExp(ctx *VarOrExpContext) {}

// ExitVarOrExp is called when production varOrExp is exited.
func (s *BaseSpartaListener) ExitVarOrExp(ctx *VarOrExpContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSpartaListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSpartaListener) ExitVariable(ctx *VariableContext) {}

// EnterVarSuffix is called when production varSuffix is entered.
func (s *BaseSpartaListener) EnterVarSuffix(ctx *VarSuffixContext) {}

// ExitVarSuffix is called when production varSuffix is exited.
func (s *BaseSpartaListener) ExitVarSuffix(ctx *VarSuffixContext) {}

// EnterNameAndArgs is called when production nameAndArgs is entered.
func (s *BaseSpartaListener) EnterNameAndArgs(ctx *NameAndArgsContext) {}

// ExitNameAndArgs is called when production nameAndArgs is exited.
func (s *BaseSpartaListener) ExitNameAndArgs(ctx *NameAndArgsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseSpartaListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseSpartaListener) ExitArgs(ctx *ArgsContext) {}

// EnterFunexp is called when production funexp is entered.
func (s *BaseSpartaListener) EnterFunexp(ctx *FunexpContext) {}

// ExitFunexp is called when production funexp is exited.
func (s *BaseSpartaListener) ExitFunexp(ctx *FunexpContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseSpartaListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseSpartaListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterParlist is called when production parlist is entered.
func (s *BaseSpartaListener) EnterParlist(ctx *ParlistContext) {}

// ExitParlist is called when production parlist is exited.
func (s *BaseSpartaListener) ExitParlist(ctx *ParlistContext) {}

// EnterTableconstructor is called when production tableconstructor is entered.
func (s *BaseSpartaListener) EnterTableconstructor(ctx *TableconstructorContext) {}

// ExitTableconstructor is called when production tableconstructor is exited.
func (s *BaseSpartaListener) ExitTableconstructor(ctx *TableconstructorContext) {}

// EnterFieldlist is called when production fieldlist is entered.
func (s *BaseSpartaListener) EnterFieldlist(ctx *FieldlistContext) {}

// ExitFieldlist is called when production fieldlist is exited.
func (s *BaseSpartaListener) ExitFieldlist(ctx *FieldlistContext) {}

// EnterField is called when production field is entered.
func (s *BaseSpartaListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSpartaListener) ExitField(ctx *FieldContext) {}

// EnterFieldsep is called when production fieldsep is entered.
func (s *BaseSpartaListener) EnterFieldsep(ctx *FieldsepContext) {}

// ExitFieldsep is called when production fieldsep is exited.
func (s *BaseSpartaListener) ExitFieldsep(ctx *FieldsepContext) {}

// EnterOperatorOr is called when production operatorOr is entered.
func (s *BaseSpartaListener) EnterOperatorOr(ctx *OperatorOrContext) {}

// ExitOperatorOr is called when production operatorOr is exited.
func (s *BaseSpartaListener) ExitOperatorOr(ctx *OperatorOrContext) {}

// EnterOperatorAnd is called when production operatorAnd is entered.
func (s *BaseSpartaListener) EnterOperatorAnd(ctx *OperatorAndContext) {}

// ExitOperatorAnd is called when production operatorAnd is exited.
func (s *BaseSpartaListener) ExitOperatorAnd(ctx *OperatorAndContext) {}

// EnterOperatorComparison is called when production operatorComparison is entered.
func (s *BaseSpartaListener) EnterOperatorComparison(ctx *OperatorComparisonContext) {}

// ExitOperatorComparison is called when production operatorComparison is exited.
func (s *BaseSpartaListener) ExitOperatorComparison(ctx *OperatorComparisonContext) {}

// EnterOperatorStrcat is called when production operatorStrcat is entered.
func (s *BaseSpartaListener) EnterOperatorStrcat(ctx *OperatorStrcatContext) {}

// ExitOperatorStrcat is called when production operatorStrcat is exited.
func (s *BaseSpartaListener) ExitOperatorStrcat(ctx *OperatorStrcatContext) {}

// EnterOperatorAddSub is called when production operatorAddSub is entered.
func (s *BaseSpartaListener) EnterOperatorAddSub(ctx *OperatorAddSubContext) {}

// ExitOperatorAddSub is called when production operatorAddSub is exited.
func (s *BaseSpartaListener) ExitOperatorAddSub(ctx *OperatorAddSubContext) {}

// EnterOperatorMulDivMod is called when production operatorMulDivMod is entered.
func (s *BaseSpartaListener) EnterOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// ExitOperatorMulDivMod is called when production operatorMulDivMod is exited.
func (s *BaseSpartaListener) ExitOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// EnterOperatorBitwise is called when production operatorBitwise is entered.
func (s *BaseSpartaListener) EnterOperatorBitwise(ctx *OperatorBitwiseContext) {}

// ExitOperatorBitwise is called when production operatorBitwise is exited.
func (s *BaseSpartaListener) ExitOperatorBitwise(ctx *OperatorBitwiseContext) {}

// EnterOperatorUnary is called when production operatorUnary is entered.
func (s *BaseSpartaListener) EnterOperatorUnary(ctx *OperatorUnaryContext) {}

// ExitOperatorUnary is called when production operatorUnary is exited.
func (s *BaseSpartaListener) ExitOperatorUnary(ctx *OperatorUnaryContext) {}

// EnterOperatorPower is called when production operatorPower is entered.
func (s *BaseSpartaListener) EnterOperatorPower(ctx *OperatorPowerContext) {}

// ExitOperatorPower is called when production operatorPower is exited.
func (s *BaseSpartaListener) ExitOperatorPower(ctx *OperatorPowerContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSpartaListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSpartaListener) ExitNumber(ctx *NumberContext) {}

// EnterStr is called when production str is entered.
func (s *BaseSpartaListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseSpartaListener) ExitStr(ctx *StrContext) {}
