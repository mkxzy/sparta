package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/symbol"
	"path/filepath"
)

type SPAImportStmtInterpreter struct {
	ast *parser.Import_stmtContext
}

// 实现解释接口
func(v *SPAImportStmtInterpreter) Interpret(state *ProgramState)  {
	moduleName := v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
	moduleFile := moduleName + ".spa"
	//log.Info(filePath)
	currentDir := state.Resolve("CURRENT_DIR").(*symbol.SPAVariable).Value.(types.SPAString)
	absFile := filepath.Join(string(currentDir), moduleFile)
	inter := NewDirectInterpreterFromFile(absFile)
	inter.Interpret(state)
}
