package symboltable

import "txtracker/pkg/ast"

type SymbolTable interface {
	InsertSymbol(symbol Symbol)
	LookupSymbol(symbolName string) Symbol
	Constructor(root *ast.Common) *SymbolTable
}
