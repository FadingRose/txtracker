package cfg

import (
	AST "txtracker/pkg/ast"
	ST "txtracker/pkg/symbol_table"
)

type StatementType int

const (
	Break StatementType = iota
	Continue
	DoWhile
	For
	Emit
	Expression
	If
	InlineAssembly
	Placeholder
	Return
	Revert
	Try
	UncheckedBlock
	VariableDeclaration
	While
	// Expression Statements includes require and assert
	Require
	Assert
	FunctionCall
	// If statement's condition includes StateVariable
	Authorize
	// Modify: Write to state variable
	Assignment
)

func (s StatementType) String() string {
	return [...]string{
		"Break",
		"Continue",
		"DoWhile",
		"For",
		"Emit",
		"Expression",
		"If",
		"InlineAssembly",
		"Placeholder",
		"Return",
		"Revert",
		"Try",
		"UncheckedBlock",
		"VariableDeclaration",
		"While",
		"Require",
		"Assert",
		"FunctionCall",
		"Authorize",
		"Modify",
	}[s]
}

type CFG struct {
	EntryPoints []*Function `json:"entryPoints"`
	Blocks      []*Block    `json:"blocks"`
	Edges       []*Edge     `json:"edges"`
	symbolTable *ST.GlobalSymbolTable
	Visitor     *Visitor
}

type Visitor struct {
	CurrentNamespace *ST.Namespace
}

func NewVisitor() *Visitor {
	return &Visitor{
		CurrentNamespace: &ST.Namespace{},
	}
}

func (v *Visitor) EnterNamespace(namespace string) {
	v.CurrentNamespace.Push(namespace)
}

func (v *Visitor) ExitNamespace() {
	v.CurrentNamespace.Pop()
}

type Function struct {
	Name       string `json:"name"`
	SrcID      int    `json:"src"`
	Block      *Block
	Parameters []*ST.Symbol
	// TODO add modifiers
}

type Block struct {
	ID                int          `json:"id"`
	Namespace         ST.Namespace `json:"namespace"`
	Statements        []*Statement `json:"statements"`
	SuccessorsEdges   []*Edge      `json:"successors"`
	PredecessorsEdges []*Edge      `json:"predecessors"`
}

func (b *Block) Collect(s *Statement) {
	b.Statements = append(b.Statements, s)
}

func (b *Block) AddSuccessorEdge(e *Edge) {
	b.SuccessorsEdges = append(b.SuccessorsEdges, e)
}

type Edge struct {
	Source      *Block   `json:"source"`
	Destination *Block   `json:"destination"`
	Type        EdgeType `json:"type"`
}

type EdgeType int

const (
	ConditionalTrue EdgeType = iota
	ConditionalFalse
	Unconditional
)

type Statement struct {
	ASTNode AST.Common `json:"astNode"`
	Type    StatementType
	Modify  []ST.Symbol
	Depends []ST.Symbol
	Declare []ST.Symbol
}
