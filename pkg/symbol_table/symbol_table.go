package symboltable

import (
	"strings"
	"txtracker/pkg/ast"
)

type GlobalSymbolTable struct {
	Table map[string]Symbol
}

type Symbol struct {
	Namespace Namespace
	Type      SymbolType
}

type Namespace []string

func (n Namespace) String() string {
	return strings.Join(n, "::")
}

type SymbolType int

const (
	StateVariable SymbolType = iota
	Function
)

func NewGlobalSymbolTable(root *ast.Common) *GlobalSymbolTable {
	gst := &GlobalSymbolTable{Table: make(map[string]Symbol)}
	gst.constructor(root)
	return gst
}

func (gst *GlobalSymbolTable) InsertSymbol(symbol Symbol) {
	gst.Table[symbol.Namespace.String()] = symbol
}

func (gst *GlobalSymbolTable) LookupSymbol(symbolName string) Symbol {
	return gst.Table[symbolName]
}

func (gst *GlobalSymbolTable) constructor(root *ast.Common) *GlobalSymbolTable {

}
