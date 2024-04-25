package cfg

import (
	AST "txtracker/internal/ast"
	"txtracker/internal/logger"
	ST "txtracker/internal/symbol_table"
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
	Assignment // Assignment
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
		"Assign",
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

func StatementToString(s *Statement) string {
	switch s.Type {
	case VariableDeclaration:
		return variableDeclarationToString(s)
	case Emit:
		return emitToString(s)
	case Return:
		return returnToString(s)
	case Assignment:
		return assignmentToString(s)
	case FunctionCall:
		return functionCallToString(s)
	case Require:
		return requireToString(s)
	}
	logger.Warning.Println("Unhandled statement to string:", s.Type)
	return ""
}

func variableDeclarationToString(s *Statement) string {
	return printDeclare(s.Declare) + " <- " + printDepends(s.Depends)
}

func emitToString(s *Statement) string {
	return printDepends(s.Depends)
}

func returnToString(s *Statement) string {
	return printDepends(s.Depends)
}

func assignmentToString(s *Statement) string {
	return printModify(s.Modify) + " <- " + printDepends(s.Depends)
}

func functionCallToString(s *Statement) string {
	funcString := func(funcs *[]ST.Symbol) string {
		// reverse the order of the funcs
		for i, j := 0, len(*funcs)-1; i < j; i, j = i+1, j-1 {
			(*funcs)[i], (*funcs)[j] = (*funcs)[j], (*funcs)[i]
		}
		// funcs[0].funcs[1].funcs[2]()
		var res string
		for i, f := range *funcs {
			res += f.Identifier
			if i != len(*funcs)-1 {
				res += "."
			} else {
				res += "()"
			}
		}
		return res
	}(&s.Declare)
	return funcString + printDepends(s.Depends)
}

func requireToString(s *Statement) string {
	return printDepends(s.Depends)
}

func printDeclare(declare []ST.Symbol) string {
	var res string
	for _, d := range declare {
		res += "[" + d.Identifier + func() string {
			if d.Type == ST.Function {
				return "()" + " "
			}
			return ""
		}() + "]" + " "
	}
	return res
}

func printModify(modify []ST.Symbol) string {
	var res string
	for i, d := range modify {
		res += "[" + d.Identifier + "]"
		if i == 0 {
			res += "* "
		}

	}
	return res
}

func printDepends(depends []ST.Symbol) string {
	var res string
	for _, d := range depends {
		res += "[" + d.Identifier + func() string {
			if d.Type == ST.Function {
				return "()"
			}
			return ""
		}() + "]" + " "
	}
	return res
}
