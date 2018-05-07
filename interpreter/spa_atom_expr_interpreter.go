package interpreter

import (
	"github.com/mkxzy/sparta/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mkxzy/sparta/symbol"
	"github.com/mkxzy/sparta/types"
)

type SPAAtomExprInterpreter struct {
	ast *parser.Atom_exprContext
}

// 实现解释接口
func(v *SPAAtomExprInterpreter) Interpret()  {
	log.Debug("Visit Atom_Expr")

	firstChild := v.ast.GetChild(0)
	switch firstChild.(type) {
	case *parser.Bracket_exprContext:
		bracketInter := &SPABracketExprInterpreter{
			ast:firstChild.(*parser.Bracket_exprContext),
		}
		bracketInter.Interpret()
	case *parser.List_literalContext:
		listInter := &SPAListLiteralInterpreter{ast: firstChild.(*parser.List_literalContext)}
		listInter.Interpret()
	case *parser.Map_literalContext:
		mapInter := &SPAMapLiteralInterpreter{ast: firstChild.(*parser.Map_literalContext)}
		mapInter.Interpret()
	case *parser.Funcall_exprContext:
		funCallInter := &SPAFuncallExprInterpreter{
			ast:firstChild.(*parser.Funcall_exprContext),
		}
		funCallInter.Interpret()
	case *parser.Table_indexContext:
		arrayIndexInter := &SPATableIndexInterpreter{
			ast: firstChild.(*parser.Table_indexContext),
		}
		arrayIndexInter.Interpret()
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
			PushValue(value)
		case parser.SpartaLexerNUMBER_LITERAL:
			PushValue(types.NewNumber(terminalNode.GetText()))
		case parser.SpartaLexerINTEGER_LITERAL:
			PushValue(types.NewInteger(terminalNode.GetText()))
		case parser.SpartaLexerSTRING:
			PushValue(types.NewString(terminalNode.GetText()))
		default:
			panic("类型无效")
		}
	}

	//if v.ast.GetChildCount() == 3 {
	//	firstChild := v.ast.GetChild(0).(antlr.TerminalNode)
	//	switch firstChild.GetText() {
	//	case "(":
	//		// 括号优先表达式
	//		testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
	//		testInter.Interpret()
	//	case "[":
	//		// 列表字面量
	//		testListInter := &SPATestListInterpreter{v.ast.GetChild(1).(*parser.Test_listContext)}
	//		testListInter.Interpret()
	//	}
	//} else {
	//	firstChild := v.ast.GetChild(0)
	//	switch firstChild.(type) {
	//	case *parser.Funcall_exprContext:
	//		funCallInter := &SPAFuncallExprInterpreter{
	//			ast:firstChild.(*parser.Funcall_exprContext),
	//		}
	//		funCallInter.Interpret()
	//	case *parser.Table_indexContext:
	//		arrayIndexInter := &SPATableIndexInterpreter{
	//			ast: firstChild.(*parser.Table_indexContext),
	//		}
	//		arrayIndexInter.Interpret()
	//	default:
	//		terminalNode, ok := firstChild.(*antlr.TerminalNodeImpl)
	//		if !ok {
	//			panic("语法不正确")
	//		}
	//		switch terminalNode.GetSymbol().GetTokenType() {
	//		case parser.SpartaLexerIDENTIFIER:
	//			value := state.currentScope.Resolve(terminalNode.GetText()).(*symbol.SPAVariable).Value
	//			PushValue(value)
	//		case parser.SpartaLexerNUMBER_LITERAL:
	//			PushValue(types.NewNumber(terminalNode.GetText()))
	//		case parser.SpartaLexerINTEGER_LITERAL:
	//			PushValue(types.NewInteger(terminalNode.GetText()))
	//		case parser.SpartaLexerSTRING:
	//			PushValue(types.NewString(terminalNode.GetText()))
	//		default:
	//			panic("类型无效")
	//		}
	//	}
	//}
}
