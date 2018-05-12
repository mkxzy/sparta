package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type SPAStmtInterpreter struct {
	ast parser.IStmtContext
	//ff *ForState
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
		return &SPAAssignStmtInterpreter{rule.(*parser.Assign_stmtContext)}
	case parser.SpartaParserRULE_fundef_stmt:
		return &SPAFundefStmtInterpreter{ast: rule.(*parser.Fundef_stmtContext)}
	case parser.SpartaParserRULE_funcall_stmt:
		return &SPAFuncallStmtInterpreter{rule.(*parser.Funcall_stmtContext)}
	case parser.SpartaParserRULE_return_stmt:
		return &SPAReturnStmtInterpreter{rule.(*parser.Return_stmtContext)}
	case parser.SpartaParserRULE_if_stmt:
		return &SPAIfStmtInterpreter{rule.(*parser.If_stmtContext)}
	case parser.SpartaParserRULE_for_stmt:
		return &SPAForStmtInterpreter{rule.(*parser.For_stmtContext)}
	case parser.SpartaParserRULE_continue_stmt:
		return &SPAContinueStmtInterpreter{}
	case parser.SpartaParserRULE_break_stmt:
		return &SPABreakStmtInterpreter{}
	case parser.SpartaParserRULE_import_stmt:
		return &SPAImportStmtInterpreter{ast: rule.(*parser.Import_stmtContext)}
	default:
		panic("不支持的语句")
	}
}