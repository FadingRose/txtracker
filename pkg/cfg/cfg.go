package cfg

import (
	AST "txtracker/pkg/ast"
	ST "txtracker/pkg/symbol_table"
)

func NewCFG(root *AST.Common, symbolTable *ST.GlobalSymbolTable) *CFG {
	cfg := &CFG{}
	cfg.EntryPoints = _constructEntryFuncs(root, symbolTable)

	return cfg
}

func _constructEntryFuncs(root *AST.Common, symbolTable *ST.GlobalSymbolTable) []*Function {

	return nil
}
