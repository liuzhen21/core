// Generated from /home/jesse/myspace/src/tkeel-io/core/pkg/tql/TQL.g4 by ANTLR 4.7.

package parser // TQL

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 94, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 6, 5, 39, 10, 5, 13, 5, 14, 5, 40, 3, 5, 3, 5, 6, 5, 45, 10, 5,
	13, 5, 14, 5, 46, 5, 5, 49, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 7, 5, 60, 10, 5, 12, 5, 14, 5, 63, 11, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 68, 10, 6, 5, 6, 70, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7,
	9, 89, 10, 9, 12, 9, 14, 9, 92, 11, 9, 3, 9, 2, 4, 8, 16, 10, 2, 4, 6,
	8, 10, 12, 14, 16, 2, 6, 3, 2, 36, 38, 3, 2, 39, 40, 4, 2, 11, 11, 13,
	17, 3, 2, 36, 37, 2, 97, 2, 18, 3, 2, 2, 2, 4, 25, 3, 2, 2, 2, 6, 33, 3,
	2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2, 14,
	73, 3, 2, 2, 2, 16, 76, 3, 2, 2, 2, 18, 19, 7, 4, 2, 2, 19, 20, 7, 5, 2,
	2, 20, 21, 5, 6, 4, 2, 21, 22, 7, 21, 2, 2, 22, 23, 5, 4, 3, 2, 23, 24,
	7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 30, 5, 8, 5, 2, 26, 27, 7, 3, 2, 2,
	27, 29, 5, 8, 5, 2, 28, 26, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3,
	2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33,
	34, 7, 44, 2, 2, 34, 7, 3, 2, 2, 2, 35, 36, 8, 5, 1, 2, 36, 49, 5, 10,
	6, 2, 37, 39, 5, 10, 6, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 7, 6, 2,
	2, 43, 45, 5, 12, 7, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44,
	3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 35, 3, 2, 2, 2,
	48, 38, 3, 2, 2, 2, 49, 61, 3, 2, 2, 2, 50, 51, 12, 5, 2, 2, 51, 52, 9,
	2, 2, 2, 52, 60, 5, 8, 5, 6, 53, 54, 12, 4, 2, 2, 54, 55, 9, 3, 2, 2, 55,
	60, 5, 8, 5, 5, 56, 57, 12, 3, 2, 2, 57, 58, 9, 4, 2, 2, 58, 60, 5, 8,
	5, 4, 59, 50, 3, 2, 2, 2, 59, 53, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 60, 63,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 9, 3, 2, 2, 2,
	63, 61, 3, 2, 2, 2, 64, 70, 7, 36, 2, 2, 65, 67, 7, 44, 2, 2, 66, 68, 7,
	45, 2, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69,
	64, 3, 2, 2, 2, 69, 65, 3, 2, 2, 2, 70, 11, 3, 2, 2, 2, 71, 72, 7, 44,
	2, 2, 72, 13, 3, 2, 2, 2, 73, 74, 5, 16, 9, 2, 74, 75, 7, 2, 2, 3, 75,
	15, 3, 2, 2, 2, 76, 77, 8, 9, 1, 2, 77, 78, 7, 47, 2, 2, 78, 90, 3, 2,
	2, 2, 79, 80, 12, 6, 2, 2, 80, 81, 9, 5, 2, 2, 81, 89, 5, 16, 9, 7, 82,
	83, 12, 5, 2, 2, 83, 84, 9, 3, 2, 2, 84, 89, 5, 16, 9, 6, 85, 86, 12, 4,
	2, 2, 86, 87, 9, 4, 2, 2, 87, 89, 5, 16, 9, 5, 88, 79, 3, 2, 2, 2, 88,
	82, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2,
	2, 90, 91, 3, 2, 2, 2, 91, 17, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 12, 30,
	40, 46, 48, 59, 61, 67, 69, 88, 90,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	"'/'", "'%'", "'+'", "'-'", "'.'",
}
var symbolicNames = []string{
	"", "", "INSERT", "INTO", "AS", "AND", "CASE", "ELSE", "END", "EQ", "FROM",
	"GT", "GTE", "LT", "LTE", "NE", "NOT", "NULL", "OR", "SELECT", "THEN",
	"WHERE", "WHEN", "GROUP", "BY", "TUMBLINGWINDOW", "HOPPINGWINDOW", "SLIDINGWINDOW",
	"SESSIONWINDOW", "DD", "HH", "MI", "SS", "MS", "MUL", "DIV", "MOD", "ADD",
	"SUB", "DOT", "TRUE", "FALSE", "ENTITYNAME", "PROPERTYNAME", "TARGETENTITY",
	"NUMBER", "INTEGER", "FLOAT", "STRING", "WHITESPACE",
}

var ruleNames = []string{
	"root", "fields", "targetEntity", "expr", "sourceEntity", "targetProperty",
	"computing", "numExp",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TQLParser struct {
	*antlr.BaseParser
}

func NewTQLParser(input antlr.TokenStream) *TQLParser {
	this := new(TQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TQL.g4"

	return this
}

// TQLParser tokens.
const (
	TQLParserEOF            = antlr.TokenEOF
	TQLParserT__0           = 1
	TQLParserINSERT         = 2
	TQLParserINTO           = 3
	TQLParserAS             = 4
	TQLParserAND            = 5
	TQLParserCASE           = 6
	TQLParserELSE           = 7
	TQLParserEND            = 8
	TQLParserEQ             = 9
	TQLParserFROM           = 10
	TQLParserGT             = 11
	TQLParserGTE            = 12
	TQLParserLT             = 13
	TQLParserLTE            = 14
	TQLParserNE             = 15
	TQLParserNOT            = 16
	TQLParserNULL           = 17
	TQLParserOR             = 18
	TQLParserSELECT         = 19
	TQLParserTHEN           = 20
	TQLParserWHERE          = 21
	TQLParserWHEN           = 22
	TQLParserGROUP          = 23
	TQLParserBY             = 24
	TQLParserTUMBLINGWINDOW = 25
	TQLParserHOPPINGWINDOW  = 26
	TQLParserSLIDINGWINDOW  = 27
	TQLParserSESSIONWINDOW  = 28
	TQLParserDD             = 29
	TQLParserHH             = 30
	TQLParserMI             = 31
	TQLParserSS             = 32
	TQLParserMS             = 33
	TQLParserMUL            = 34
	TQLParserDIV            = 35
	TQLParserMOD            = 36
	TQLParserADD            = 37
	TQLParserSUB            = 38
	TQLParserDOT            = 39
	TQLParserTRUE           = 40
	TQLParserFALSE          = 41
	TQLParserENTITYNAME     = 42
	TQLParserPROPERTYNAME   = 43
	TQLParserTARGETENTITY   = 44
	TQLParserNUMBER         = 45
	TQLParserINTEGER        = 46
	TQLParserFLOAT          = 47
	TQLParserSTRING         = 48
	TQLParserWHITESPACE     = 49
)

// TQLParser rules.
const (
	TQLParserRULE_root           = 0
	TQLParserRULE_fields         = 1
	TQLParserRULE_targetEntity   = 2
	TQLParserRULE_expr           = 3
	TQLParserRULE_sourceEntity   = 4
	TQLParserRULE_targetProperty = 5
	TQLParserRULE_computing      = 6
	TQLParserRULE_numExp         = 7
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) INSERT() antlr.TerminalNode {
	return s.GetToken(TQLParserINSERT, 0)
}

func (s *RootContext) INTO() antlr.TerminalNode {
	return s.GetToken(TQLParserINTO, 0)
}

func (s *RootContext) TargetEntity() ITargetEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetEntityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetEntityContext)
}

func (s *RootContext) SELECT() antlr.TerminalNode {
	return s.GetToken(TQLParserSELECT, 0)
}

func (s *RootContext) Fields() IFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(TQLParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *TQLParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TQLParserRULE_root)

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
		p.SetState(16)
		p.Match(TQLParserINSERT)
	}
	{
		p.SetState(17)
		p.Match(TQLParserINTO)
	}
	{
		p.SetState(18)
		p.TargetEntity()
	}
	{
		p.SetState(19)
		p.Match(TQLParserSELECT)
	}
	{
		p.SetState(20)
		p.Fields()
	}
	{
		p.SetState(21)
		p.Match(TQLParserEOF)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FieldsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *TQLParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TQLParserRULE_fields)
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
	{
		p.SetState(23)
		p.expr(0)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TQLParserT__0 {
		{
			p.SetState(24)
			p.Match(TQLParserT__0)
		}
		{
			p.SetState(25)
			p.expr(0)
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITargetEntityContext is an interface to support dynamic dispatch.
type ITargetEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetEntityContext differentiates from other interfaces.
	IsTargetEntityContext()
}

type TargetEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetEntityContext() *TargetEntityContext {
	var p = new(TargetEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_targetEntity
	return p
}

func (*TargetEntityContext) IsTargetEntityContext() {}

func NewTargetEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetEntityContext {
	var p = new(TargetEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_targetEntity

	return p
}

func (s *TargetEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetEntityContext) ENTITYNAME() antlr.TerminalNode {
	return s.GetToken(TQLParserENTITYNAME, 0)
}

func (s *TargetEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterTargetEntity(s)
	}
}

func (s *TargetEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitTargetEntity(s)
	}
}

func (p *TQLParser) TargetEntity() (localctx ITargetEntityContext) {
	localctx = NewTargetEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TQLParserRULE_targetEntity)

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
		p.SetState(31)
		p.Match(TQLParserENTITYNAME)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionContext struct {
	*ExprContext
}

func NewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionContext {
	var p = new(ExpressionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) AllSourceEntity() []ISourceEntityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceEntityContext)(nil)).Elem())
	var tst = make([]ISourceEntityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceEntityContext)
		}
	}

	return tst
}

func (s *ExpressionContext) SourceEntity(i int) ISourceEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceEntityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceEntityContext)
}

func (s *ExpressionContext) AS() antlr.TerminalNode {
	return s.GetToken(TQLParserAS, 0)
}

func (s *ExpressionContext) AllTargetProperty() []ITargetPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITargetPropertyContext)(nil)).Elem())
	var tst = make([]ITargetPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITargetPropertyContext)
		}
	}

	return tst
}

func (s *ExpressionContext) TargetProperty(i int) ITargetPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITargetPropertyContext)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

type DummyAddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewDummyAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DummyAddSubContext {
	var p = new(DummyAddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DummyAddSubContext) GetOp() antlr.Token { return s.op }

func (s *DummyAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *DummyAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DummyAddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DummyAddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DummyAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterDummyAddSub(s)
	}
}

func (s *DummyAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitDummyAddSub(s)
	}
}

type DummyMulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewDummyMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DummyMulDivContext {
	var p = new(DummyMulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DummyMulDivContext) GetOp() antlr.Token { return s.op }

func (s *DummyMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *DummyMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DummyMulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DummyMulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DummyMulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterDummyMulDiv(s)
	}
}

func (s *DummyMulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitDummyMulDiv(s)
	}
}

type DummyCompareValueContext struct {
	*ExprContext
	op antlr.Token
}

func NewDummyCompareValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DummyCompareValueContext {
	var p = new(DummyCompareValueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DummyCompareValueContext) GetOp() antlr.Token { return s.op }

func (s *DummyCompareValueContext) SetOp(v antlr.Token) { s.op = v }

func (s *DummyCompareValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DummyCompareValueContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DummyCompareValueContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DummyCompareValueContext) EQ() antlr.TerminalNode {
	return s.GetToken(TQLParserEQ, 0)
}

func (s *DummyCompareValueContext) GT() antlr.TerminalNode {
	return s.GetToken(TQLParserGT, 0)
}

func (s *DummyCompareValueContext) LT() antlr.TerminalNode {
	return s.GetToken(TQLParserLT, 0)
}

func (s *DummyCompareValueContext) GTE() antlr.TerminalNode {
	return s.GetToken(TQLParserGTE, 0)
}

func (s *DummyCompareValueContext) LTE() antlr.TerminalNode {
	return s.GetToken(TQLParserLTE, 0)
}

func (s *DummyCompareValueContext) NE() antlr.TerminalNode {
	return s.GetToken(TQLParserNE, 0)
}

func (s *DummyCompareValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterDummyCompareValue(s)
	}
}

func (s *DummyCompareValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitDummyCompareValue(s)
	}
}

func (p *TQLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TQLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, TQLParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(34)
			p.SourceEntity()
		}

	case 2:
		localctx = NewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == TQLParserMUL || _la == TQLParserENTITYNAME {
			{
				p.SetState(35)
				p.SourceEntity()
			}

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(40)
			p.Match(TQLParserAS)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(41)
					p.TargetProperty()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDummyMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(49)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DummyMulDivContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(TQLParserMUL-34))|(1<<(TQLParserDIV-34))|(1<<(TQLParserMOD-34)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DummyMulDivContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(50)
					p.expr(4)
				}

			case 2:
				localctx = NewDummyAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_expr)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(52)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DummyAddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == TQLParserADD || _la == TQLParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DummyAddSubContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(53)
					p.expr(3)
				}

			case 3:
				localctx = NewDummyCompareValueContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_expr)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				p.SetState(55)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DummyCompareValueContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TQLParserEQ)|(1<<TQLParserGT)|(1<<TQLParserGTE)|(1<<TQLParserLT)|(1<<TQLParserLTE)|(1<<TQLParserNE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DummyCompareValueContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(56)
					p.expr(2)
				}

			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ISourceEntityContext is an interface to support dynamic dispatch.
type ISourceEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceEntityContext differentiates from other interfaces.
	IsSourceEntityContext()
}

type SourceEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceEntityContext() *SourceEntityContext {
	var p = new(SourceEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_sourceEntity
	return p
}

func (*SourceEntityContext) IsSourceEntityContext() {}

func NewSourceEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceEntityContext {
	var p = new(SourceEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_sourceEntity

	return p
}

func (s *SourceEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceEntityContext) ENTITYNAME() antlr.TerminalNode {
	return s.GetToken(TQLParserENTITYNAME, 0)
}

func (s *SourceEntityContext) PROPERTYNAME() antlr.TerminalNode {
	return s.GetToken(TQLParserPROPERTYNAME, 0)
}

func (s *SourceEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterSourceEntity(s)
	}
}

func (s *SourceEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitSourceEntity(s)
	}
}

func (p *TQLParser) SourceEntity() (localctx ISourceEntityContext) {
	localctx = NewSourceEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TQLParserRULE_sourceEntity)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TQLParserMUL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(TQLParserMUL)
		}

	case TQLParserENTITYNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(TQLParserENTITYNAME)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(64)
				p.Match(TQLParserPROPERTYNAME)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITargetPropertyContext is an interface to support dynamic dispatch.
type ITargetPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetPropertyContext differentiates from other interfaces.
	IsTargetPropertyContext()
}

type TargetPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetPropertyContext() *TargetPropertyContext {
	var p = new(TargetPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_targetProperty
	return p
}

func (*TargetPropertyContext) IsTargetPropertyContext() {}

func NewTargetPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetPropertyContext {
	var p = new(TargetPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_targetProperty

	return p
}

func (s *TargetPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetPropertyContext) ENTITYNAME() antlr.TerminalNode {
	return s.GetToken(TQLParserENTITYNAME, 0)
}

func (s *TargetPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterTargetProperty(s)
	}
}

func (s *TargetPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitTargetProperty(s)
	}
}

func (p *TQLParser) TargetProperty() (localctx ITargetPropertyContext) {
	localctx = NewTargetPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TQLParserRULE_targetProperty)

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
		p.SetState(69)
		p.Match(TQLParserENTITYNAME)
	}

	return localctx
}

// IComputingContext is an interface to support dynamic dispatch.
type IComputingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComputingContext differentiates from other interfaces.
	IsComputingContext()
}

type ComputingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComputingContext() *ComputingContext {
	var p = new(ComputingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_computing
	return p
}

func (*ComputingContext) IsComputingContext() {}

func NewComputingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComputingContext {
	var p = new(ComputingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_computing

	return p
}

func (s *ComputingContext) GetParser() antlr.Parser { return s.parser }

func (s *ComputingContext) NumExp() INumExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumExpContext)
}

func (s *ComputingContext) EOF() antlr.TerminalNode {
	return s.GetToken(TQLParserEOF, 0)
}

func (s *ComputingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComputingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterComputing(s)
	}
}

func (s *ComputingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitComputing(s)
	}
}

func (p *TQLParser) Computing() (localctx IComputingContext) {
	localctx = NewComputingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TQLParserRULE_computing)

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
		p.SetState(71)
		p.numExp(0)
	}
	{
		p.SetState(72)
		p.Match(TQLParserEOF)
	}

	return localctx
}

// INumExpContext is an interface to support dynamic dispatch.
type INumExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumExpContext differentiates from other interfaces.
	IsNumExpContext()
}

type NumExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumExpContext() *NumExpContext {
	var p = new(NumExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TQLParserRULE_numExp
	return p
}

func (*NumExpContext) IsNumExpContext() {}

func NewNumExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumExpContext {
	var p = new(NumExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TQLParserRULE_numExp

	return p
}

func (s *NumExpContext) GetParser() antlr.Parser { return s.parser }

func (s *NumExpContext) CopyFrom(ctx *NumExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	*NumExpContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.NumExpContext = NewEmptyNumExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumExpContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TQLParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitNumber(s)
	}
}

type CompareValueContext struct {
	*NumExpContext
	op antlr.Token
}

func NewCompareValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareValueContext {
	var p = new(CompareValueContext)

	p.NumExpContext = NewEmptyNumExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumExpContext))

	return p
}

func (s *CompareValueContext) GetOp() antlr.Token { return s.op }

func (s *CompareValueContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareValueContext) AllNumExp() []INumExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumExpContext)(nil)).Elem())
	var tst = make([]INumExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumExpContext)
		}
	}

	return tst
}

func (s *CompareValueContext) NumExp(i int) INumExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumExpContext)
}

func (s *CompareValueContext) EQ() antlr.TerminalNode {
	return s.GetToken(TQLParserEQ, 0)
}

func (s *CompareValueContext) GT() antlr.TerminalNode {
	return s.GetToken(TQLParserGT, 0)
}

func (s *CompareValueContext) LT() antlr.TerminalNode {
	return s.GetToken(TQLParserLT, 0)
}

func (s *CompareValueContext) GTE() antlr.TerminalNode {
	return s.GetToken(TQLParserGTE, 0)
}

func (s *CompareValueContext) LTE() antlr.TerminalNode {
	return s.GetToken(TQLParserLTE, 0)
}

func (s *CompareValueContext) NE() antlr.TerminalNode {
	return s.GetToken(TQLParserNE, 0)
}

func (s *CompareValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterCompareValue(s)
	}
}

func (s *CompareValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitCompareValue(s)
	}
}

type MulDivContext struct {
	*NumExpContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.NumExpContext = NewEmptyNumExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumExpContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllNumExp() []INumExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumExpContext)(nil)).Elem())
	var tst = make([]INumExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumExpContext)
		}
	}

	return tst
}

func (s *MulDivContext) NumExp(i int) INumExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumExpContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*NumExpContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.NumExpContext = NewEmptyNumExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumExpContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllNumExp() []INumExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumExpContext)(nil)).Elem())
	var tst = make([]INumExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumExpContext)
		}
	}

	return tst
}

func (s *AddSubContext) NumExp(i int) INumExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumExpContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TQLListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (p *TQLParser) NumExp() (localctx INumExpContext) {
	return p.numExp(0)
}

func (p *TQLParser) numExp(_p int) (localctx INumExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewNumExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, TQLParserRULE_numExp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewNumberContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(75)
		p.Match(TQLParserNUMBER)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewNumExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_numExp)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(78)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MulDivContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == TQLParserMUL || _la == TQLParserDIV) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MulDivContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(79)
					p.numExp(5)
				}

			case 2:
				localctx = NewAddSubContext(p, NewNumExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_numExp)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(81)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == TQLParserADD || _la == TQLParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AddSubContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(82)
					p.numExp(4)
				}

			case 3:
				localctx = NewCompareValueContext(p, NewNumExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TQLParserRULE_numExp)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(84)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*CompareValueContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TQLParserEQ)|(1<<TQLParserGT)|(1<<TQLParserGTE)|(1<<TQLParserLT)|(1<<TQLParserLTE)|(1<<TQLParserNE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*CompareValueContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(85)
					p.numExp(3)
				}

			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

func (p *TQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 7:
		var t *NumExpContext = nil
		if localctx != nil {
			t = localctx.(*NumExpContext)
		}
		return p.NumExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TQLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TQLParser) NumExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
