package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
	"github.com/mkxzy/sparta/operation"
)

type SPAAtomExprInterpreter struct {
	ast *parser.Atom_exprContext
}

// 实现解释接口
func(v *SPAAtomExprInterpreter) Interpret(state *ProgramState)  {
	log.Debug("Visit Atom_Expr")

	firstChild := v.ast.GetChild(0)
	switch firstChild.(type) {
	case *parser.Bracket_exprContext:
		bracketInter := &SPABracketExprInterpreter{
			ast:firstChild.(*parser.Bracket_exprContext),
		}
		bracketInter.Interpret(state)
	case *parser.List_literalContext:
		listInter := &SPAListLiteralInterpreter{ast: firstChild.(*parser.List_literalContext)}
		listInter.Interpret(state)
	case *parser.Map_literalContext:
		mapInter := &SPAMapLiteralInterpreter{ast: firstChild.(*parser.Map_literalContext)}
		mapInter.Interpret(state)
	case *parser.Funcall_exprContext:
		funCallInter := &SPAFuncallExprInterpreter{
			ast:firstChild.(*parser.Funcall_exprContext),
		}
		funCallInter.Interpret(state)
	case *parser.Table_indexContext:
		arrayIndexInter := &SPATableIndexInterpreter{
			ast: firstChild.(*parser.Table_indexContext),
		}
		arrayIndexInter.Interpret(state)
	default:
		terminalNode, ok := firstChild.(antlr.TerminalNode)
		if !ok {
			panic("语法不正确")
		}
		switch terminalNode.GetSymbol().GetTokenType() {
		case parser.SpartaLexerIDENTIFIER:
			log.Notice(*state)
			log.Notice(terminalNode.GetText())
			value := state.Resolve(terminalNode.GetText()).(*symbol.SPAVariable).Value
			operation.PushValue(value)
		case parser.SpartaLexerNUMBER_LITERAL:
			operation.PushValue(types.NewNumber(terminalNode.GetText()))
		case parser.SpartaLexerINTEGER_LITERAL:
			operation.PushValue(types.NewInteger(terminalNode.GetText()))
		case parser.SpartaLexerSTRING:
			operation.PushValue(types.NewString(terminalNode.GetText()))
		default:
			panic("类型无效")
		}
	}
}
