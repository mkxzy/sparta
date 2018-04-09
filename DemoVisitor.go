package main

import (
	"github.com/mkxzy/sparta/parser"
)

type DemoVisitor struct {
	parser.BaseSpartaVisitor
}

func (v *DemoVisitor) VisitChunk(ctx *parser.ChunkContext) interface{} {
	return v.VisitChildren(ctx)
}