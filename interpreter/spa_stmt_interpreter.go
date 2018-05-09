package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SPAStmtInterpreter struct {
	ast parser.IStmtContext
	ff *ForState
}

// 实现解释接口
func(v *SPAStmtInterpreter) Interpret(state *ProgramState)  {
	inter := v.getRealStmtInterpreter()
	inter.Interpret(state)
}

func(v *SPAStmtInterpreter) getRealStmtInterpreter() SPAInterpreter {
	rule := v.ast.GetChild(0).(antlr.RuleContext)
	switch rule.GetRuleIndex() {
	case parser.SpartaParserRULE_assign_stmt:
		return &SPAAssignStmtInterpreter{rule.(parser.IAssign_stmtContext)}
	case parser.SpartaParserRULE_fundef_stmt:
		return &SPAFundefStmtInterpreter{ast: rule.(parser.IFundef_stmtContext)}
	case parser.SpartaParserRULE_funcall_stmt:
		return &SPAFuncallStmtInterpreter{rule.(parser.IFuncall_stmtContext)}
	case parser.SpartaParserRULE_return_stmt:
		return &SPAReturnStmtInterpreter{rule.(*parser.Return_stmtContext), v.ff}
		//v.ExecReturnStmt(rule.(*parser.Return_stmtContext), ff)
	case parser.SpartaParserRULE_if_stmt:
		return &SPAIfStmtInterpreter{rule.(*parser.If_stmtContext), v.ff}
		//v.ExecIfStmt(rule.(*parser.If_stmtContext), ff)
	case parser.SpartaParserRULE_for_stmt:
		return &SPAForStmtInterpreter{rule.(*parser.For_stmtContext), v.ff}
		//v.ExecForStmt(rule.(*parser.For_stmtContext), ff)
	case parser.SpartaParserRULE_continue_stmt:
		return &SPAContinueStmtInterpreter{rule.(parser.IContinue_stmtContext), v.ff}
		//v.ExecContinueStmt(rule.(*parser.Continue_stmtContext), ff)
	case parser.SpartaParserRULE_break_stmt:
		return &SPABreakStmtInterpreter{rule.(parser.IBreak_stmtContext), v.ff}
		//v.ExecBreakStmt(rule.(*parser.Break_stmtContext), ff)
	default:
		panic("不支持的语句")
	}
}