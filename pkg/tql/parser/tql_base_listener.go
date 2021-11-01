// Generated from /home/jesse/myspace/src/tkeel-io/core/pkg/tql/TQL.g4 by ANTLR 4.7.

package parser // TQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTQLListener is a complete listener for a parse tree produced by TQLParser.
type BaseTQLListener struct{}

var _ TQLListener = &BaseTQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseTQLListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseTQLListener) ExitRoot(ctx *RootContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseTQLListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseTQLListener) ExitFields(ctx *FieldsContext) {}

// EnterTargetEntity is called when production targetEntity is entered.
func (s *BaseTQLListener) EnterTargetEntity(ctx *TargetEntityContext) {}

// ExitTargetEntity is called when production targetEntity is exited.
func (s *BaseTQLListener) ExitTargetEntity(ctx *TargetEntityContext) {}

// EnterExpression is called when production Expression is entered.
func (s *BaseTQLListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production Expression is exited.
func (s *BaseTQLListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDummyAddSub is called when production DummyAddSub is entered.
func (s *BaseTQLListener) EnterDummyAddSub(ctx *DummyAddSubContext) {}

// ExitDummyAddSub is called when production DummyAddSub is exited.
func (s *BaseTQLListener) ExitDummyAddSub(ctx *DummyAddSubContext) {}

// EnterDummyMulDiv is called when production DummyMulDiv is entered.
func (s *BaseTQLListener) EnterDummyMulDiv(ctx *DummyMulDivContext) {}

// ExitDummyMulDiv is called when production DummyMulDiv is exited.
func (s *BaseTQLListener) ExitDummyMulDiv(ctx *DummyMulDivContext) {}

// EnterDummyCompareValue is called when production DummyCompareValue is entered.
func (s *BaseTQLListener) EnterDummyCompareValue(ctx *DummyCompareValueContext) {}

// ExitDummyCompareValue is called when production DummyCompareValue is exited.
func (s *BaseTQLListener) ExitDummyCompareValue(ctx *DummyCompareValueContext) {}

// EnterSourceEntity is called when production sourceEntity is entered.
func (s *BaseTQLListener) EnterSourceEntity(ctx *SourceEntityContext) {}

// ExitSourceEntity is called when production sourceEntity is exited.
func (s *BaseTQLListener) ExitSourceEntity(ctx *SourceEntityContext) {}

// EnterTargetProperty is called when production targetProperty is entered.
func (s *BaseTQLListener) EnterTargetProperty(ctx *TargetPropertyContext) {}

// ExitTargetProperty is called when production targetProperty is exited.
func (s *BaseTQLListener) ExitTargetProperty(ctx *TargetPropertyContext) {}

// EnterComputing is called when production computing is entered.
func (s *BaseTQLListener) EnterComputing(ctx *ComputingContext) {}

// ExitComputing is called when production computing is exited.
func (s *BaseTQLListener) ExitComputing(ctx *ComputingContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseTQLListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseTQLListener) ExitNumber(ctx *NumberContext) {}

// EnterCompareValue is called when production CompareValue is entered.
func (s *BaseTQLListener) EnterCompareValue(ctx *CompareValueContext) {}

// ExitCompareValue is called when production CompareValue is exited.
func (s *BaseTQLListener) ExitCompareValue(ctx *CompareValueContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTQLListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTQLListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTQLListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTQLListener) ExitAddSub(ctx *AddSubContext) {}
