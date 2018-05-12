/**
直接执行解释器（边解析边执行）
 */
package interpreter

import (
	"github.com/op/go-logging"
	"github.com/mkxzy/sparta/parser"
	"github.com/mkxzy/sparta/symbol"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var log = logging.MustGetLogger("SPAProgramInterpreter")

type SPAProgramInterpreter struct {
	ast parser.IProgramContext
	currentScope symbol.Scope
	savedScope symbol.Scope
}

func NewDirectInterpreterFromFile(filePath string) *SPAProgramInterpreter {
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewSpartaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSpartaParser(stream)
	p.BuildParseTrees = true
	ast := p.Program()                             	// 生成解析树
	return &SPAProgramInterpreter{
		ast: ast,
	}
}

// 实现解释接口
func(v *SPAProgramInterpreter) Interpret(state *ProgramState)  {
	for i := 0; i < v.ast.GetChildCount()-1; i++ {
		stmtContext := v.ast.GetChild(i).(*parser.StmtContext)
		stmtInter := &SPAStmtInterpreter{ast: stmtContext}
		stmtInter.Interpret(state)
	}
}