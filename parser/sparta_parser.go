// Code generated from Sparta.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparta

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 5, 2, 14,
	10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 19, 10, 3, 13, 3, 14, 3, 20, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10,
	2, 2, 2, 28, 2, 13, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 22, 3, 2, 2, 2, 8,
	24, 3, 2, 2, 2, 10, 29, 3, 2, 2, 2, 12, 14, 5, 4, 3, 2, 13, 12, 3, 2, 2,
	2, 13, 14, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3,
	2, 2, 2, 17, 19, 5, 6, 4, 2, 18, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20,
	18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 5, 3, 2, 2, 2, 22, 23, 5, 8, 5,
	2, 23, 7, 3, 2, 2, 2, 24, 25, 7, 3, 2, 2, 25, 26, 7, 6, 2, 2, 26, 27, 7,
	4, 2, 2, 27, 28, 5, 10, 6, 2, 28, 9, 3, 2, 2, 2, 29, 30, 7, 5, 2, 2, 30,
	11, 3, 2, 2, 2, 4, 13, 20,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'var'", "'='",
}
var symbolicNames = []string{
	"", "", "", "NUMBER_LITERAL", "IDENTIFIER", "COMMENT", "LINE_COMMENT",
	"WS", "SHEBANG",
}

var ruleNames = []string{
	"program", "statList", "stat", "varStat", "exp",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SpartaParser struct {
	*antlr.BaseParser
}

func NewSpartaParser(input antlr.TokenStream) *SpartaParser {
	this := new(SpartaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sparta.g4"

	return this
}

// SpartaParser tokens.
const (
	SpartaParserEOF            = antlr.TokenEOF
	SpartaParserT__0           = 1
	SpartaParserT__1           = 2
	SpartaParserNUMBER_LITERAL = 3
	SpartaParserIDENTIFIER     = 4
	SpartaParserCOMMENT        = 5
	SpartaParserLINE_COMMENT   = 6
	SpartaParserWS             = 7
	SpartaParserSHEBANG        = 8
)

// SpartaParser rules.
const (
	SpartaParserRULE_program  = 0
	SpartaParserRULE_statList = 1
	SpartaParserRULE_stat     = 2
	SpartaParserRULE_varStat  = 3
	SpartaParserRULE_exp      = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SpartaParserEOF, 0)
}

func (s *ProgramContext) StatList() IStatListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SpartaParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SpartaParserT__0 {
		{
			p.SetState(10)
			p.StatList()
		}

	}
	{
		p.SetState(13)
		p.Match(SpartaParserEOF)
	}

	return localctx
}

// IStatListContext is an interface to support dynamic dispatch.
type IStatListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatListContext differentiates from other interfaces.
	IsStatListContext()
}

type StatListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatListContext() *StatListContext {
	var p = new(StatListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_statList
	return p
}

func (*StatListContext) IsStatListContext() {}

func NewStatListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatListContext {
	var p = new(StatListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_statList

	return p
}

func (s *StatListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatListContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *StatListContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *StatListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterStatList(s)
	}
}

func (s *StatListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitStatList(s)
	}
}

func (s *StatListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitStatList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) StatList() (localctx IStatListContext) {
	localctx = NewStatListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SpartaParserRULE_statList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SpartaParserT__0 {
		{
			p.SetState(15)
			p.Stat()
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) VarStat() IVarStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitStat(s)
	}
}

func (s *StatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SpartaParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.VarStat()
	}

	return localctx
}

// IVarStatContext is an interface to support dynamic dispatch.
type IVarStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarStatContext differentiates from other interfaces.
	IsVarStatContext()
}

type VarStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarStatContext() *VarStatContext {
	var p = new(VarStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_varStat
	return p
}

func (*VarStatContext) IsVarStatContext() {}

func NewVarStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarStatContext {
	var p = new(VarStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_varStat

	return p
}

func (s *VarStatContext) GetParser() antlr.Parser { return s.parser }

func (s *VarStatContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SpartaParserIDENTIFIER, 0)
}

func (s *VarStatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterVarStat(s)
	}
}

func (s *VarStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitVarStat(s)
	}
}

func (s *VarStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitVarStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) VarStat() (localctx IVarStatContext) {
	localctx = NewVarStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SpartaParserRULE_varStat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(SpartaParserT__0)
	}
	{
		p.SetState(23)
		p.Match(SpartaParserIDENTIFIER)
	}
	{
		p.SetState(24)
		p.Match(SpartaParserT__1)
	}
	{
		p.SetState(25)
		p.Exp()
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SpartaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SpartaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(SpartaParserNUMBER_LITERAL, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SpartaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (s *ExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SpartaVisitor:
		return t.VisitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SpartaParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SpartaParserRULE_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(SpartaParserNUMBER_LITERAL)
	}

	return localctx
}
