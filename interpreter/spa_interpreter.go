/**
解释器接口
 */
package interpreter

import "github.com/mkxzy/sparta/parser"

type SPAInterpreter interface {
	Interpret(ctx parser.IProgramContext)
}
