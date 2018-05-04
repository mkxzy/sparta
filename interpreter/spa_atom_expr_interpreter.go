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

	if v.ast.GetChildCount() == 3 {
		firstChild := v.ast.GetChild(0).(antlr.TerminalNode)
		switch firstChild.GetText() {
		case "(":
			// 括号优先表达式
			testInter := &SPATestInterpreter{v.ast.GetChild(1).(*parser.TestContext)}
			testInter.Interpret()
		case "[":
			// 列表字面量
			testListInter := &SPATestListInterpreter{v.ast.GetChild(1).(*parser.Test_listContext)}
			testListInter.Interpret()
		}
	} else if v.ast.GetChildCount() == 4 {
		//列表访问
		testInter := &SPATestInterpreter{v.ast.GetChild(2).(*parser.TestContext)}
		testInter.Interpret()
		name := v.ast.GetToken(parser.SpartaLexerIDENTIFIER, 0).GetText()
		sym := state.currentScope.Resolve(name).(*symbol.SPAVariable)
		value := sym.Value.(*types.SPAList)
		index := PopValue().(types.SPAInteger)
		PushValue(value.Get(index))
	} else {
		funCallContext, ok := v.ast.GetChild(0).(*parser.Funcall_exprContext)
		// 函数调用
		if ok {
			funCallInter := &SPAFuncallExprInterpreter{
				ast:funCallContext,
			}
			funCallInter.Interpret()
		} else {
			//变量或常量
			terminalNode := v.ast.GetChild(0).(antlr.TerminalNode)
			tt := terminalNode.GetSymbol().GetTokenType()
			switch tt {
			case parser.SpartaLexerIDENTIFIER:
				varName := terminalNode.GetText()
				log.Infof("变量名称：%s", varName)
				log.Debug(state.currentScope)
				value := state.currentScope.Resolve(varName).(*symbol.SPAVariable).Value
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
	}
}