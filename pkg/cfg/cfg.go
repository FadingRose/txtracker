package cfg

import (
	AST "txtracker/pkg/ast"
	"txtracker/pkg/logger"
	ST "txtracker/pkg/symbol_table"
)

func NewCFG(root *AST.Common, symbolTable *ST.GlobalSymbolTable) *CFG {
	cfg := &CFG{
		symbolTable: symbolTable,
		Visitor:     NewVisitor(),
	}
	cfg.EntryPoints = cfg._constructEntryFuncs(root)

	logger.Info.Println("CFG constructed")
	return cfg
}

func (cfg *CFG) _constructEntryFuncs(root *AST.Common) []*Function {
	contractDefs := cfg._findContractDefinition(root)
	var entryFuncs []*Function
	for _, contractDef := range contractDefs {
		entryFuncs = append(entryFuncs, cfg._findEntryFunc(contractDef)...)
	}

	return entryFuncs
}

func (cfg *CFG) _findContractDefinition(root *AST.Common) []*AST.Common {
	var contractDefs []*AST.Common

	for _, node := range root.Children {
		if node.NodeType == "ContractDefinition" {
			contractDefs = append(contractDefs, node)
		}
	}

	return contractDefs
}

func (cfg *CFG) _findEntryFunc(contractDef *AST.Common) []*Function {
	var entryFuncs []*Function
	contractName := contractDef.ASTNode.(*AST.ContractDefinition).Name
	cfg.Visitor.EnterNamespace(contractName)
	for _, node := range contractDef.Children {

		if node.NodeType == "FunctionDefinition" {
			funcDef := node.ASTNode.(*AST.FunctionDefinition)

			if funcDef.IsPublic() || funcDef.IsExternal() {
				if !funcDef.IsImplemented() {
					continue
				}
				cfg.Visitor.EnterNamespace(funcDef.Name)
				entryFuncs = append(entryFuncs, &Function{
					Name:  funcDef.Name,
					Block: cfg._constructFuncLevelBlock(funcDef),
				})
				cfg.Visitor.ExitNamespace()
			}

		}

	}
	cfg.Visitor.ExitNamespace()

	return entryFuncs
}

func (cfg *CFG) _constructFuncLevelBlock(funcDef *AST.FunctionDefinition) *Block {
	block := &Block{
		ID:        funcDef.ID,
		Namespace: *cfg.Visitor.CurrentNamespace,
	}

	for _, stmt := range funcDef.Body.Statements {
		block.Statements = append(block.Statements, cfg._constructStatement(stmt))
	}

	return block
}

func (cfg *CFG) _constructStatement(stmt *AST.Common) *Statement {

	_type := cfg._getStatementType(stmt)
	modify, depends, declare := cfg._getModifyAndDependsSymbols(stmt, _type)
	return &Statement{
		ASTNode: *stmt,
		Type:    _type,
		Modify:  modify,
		Depends: depends,
		Declare: declare,
	}
}

// returns the symbols that are modified and depends on the given variable declaration statement
func (cfg *CFG) _getModifyAndDependsSymbols(stmt *AST.Common, _type StatementType) ([]ST.Symbol, []ST.Symbol, []ST.Symbol) {
	var modify, depends, declare []ST.Symbol
	switch _type {
	case VariableDeclaration:
		cfg._getModifyAndDependsSymbolsForVariableDeclaration(stmt, &modify, &depends, &declare)
	case Assert:
		cfg._getModifyAndDependsSymbolsForAssert(stmt, &modify, &depends, &declare)
	case Require:
		cfg._getModifyAndDependsSymbolsForRequire(stmt, &modify, &depends, &declare)
	case FunctionCall:
		cfg._getModifyAndDependsSymbolsForFunctionCall(stmt, &modify, &depends, &declare)
	}
	return modify, depends, declare
}

func (cfg *CFG) _getModifyAndDependsSymbolsForVariableDeclaration(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	// stmt.NodeType() == "VariableDeclareStatement"
	declarations, statevars := stmt.ASTNode.(*AST.VariableDeclarationStatement).GetDeclarations()
	for i, declaration := range declarations {
		decl := &ST.Symbol{
			Namespace:  *cfg.Visitor.CurrentNamespace,
			Identifier: declaration,
			Type: func() ST.SymbolType {
				if statevars[i] {
					return ST.StateVariable
				} else {
					return ST.LocalVariable
				}
			}(),
		}
		*declare = append(*declare, *decl)
	}
}

func (cfg *CFG) _getModifyAndDependsSymbolsForAssert(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
}

func (cfg *CFG) _getModifyAndDependsSymbolsForRequire(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
}

func (cfg *CFG) _getModifyAndDependsSymbolsForFunctionCall(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {

}

func (cfg *CFG) _getStatementType(stmt *AST.Common) StatementType {
	switch stmt.NodeType {
	case "IfStatement":
		return If
	case "ForStatement":
		return For
	case "WhileStatement":
		return While
	case "DoWhileStatement":
		return DoWhile
	case "Return":
		return Return
	case "EmitStatement":
		return Emit
	case "VariableDeclarationStatement":
		return VariableDeclaration
	case "ExpressionStatement":
		// whether the expression is a require?
		if cfg._isRequire(stmt.ASTNode.(*AST.ExpressionStatement)) {
			return Require
		} else if cfg._isAssert(stmt.ASTNode.(*AST.ExpressionStatement)) {
			return Assert
		} else if cfg._isModify(stmt.ASTNode.(*AST.ExpressionStatement)) {
			return Modify
		} else if cfg._isFunctionCall(stmt.ASTNode.(*AST.ExpressionStatement)) {
			return FunctionCall
		} else {
			return Expression
		}
	case "BreakStatement":
		return Break
	case "ContinueStatement":
		return Continue
	case "PlaceholderStatement":
		return Placeholder
	case "RevertStatement":
		return Revert
	case "TryStatement":
		return Try
	case "UncheckedBlock":
		return UncheckedBlock
	case "InlineAssemblyStatement":
		return InlineAssembly
	default:
		logger.Warning.Println("Unknown statement type:", stmt.NodeType)
		return -1
	}
}

func (cfg *CFG) _isRequire(stmt *AST.ExpressionStatement) bool {

	if funCall, ok := stmt.Expression.ASTNode.(*AST.FunctionCall); !ok {
		return false
	} else if funCall.Expression.NodeType != "Identifier" {
		return false
	} else if funCall.Expression.ASTNode.(*AST.Identifier).Name != "require" {
		return false
	}

	return true
}

func (cfg *CFG) _isAssert(stmt *AST.ExpressionStatement) bool {

	if funCall, ok := stmt.Expression.ASTNode.(*AST.FunctionCall); !ok {
		return false
	} else if funCall.Expression.NodeType != "Identifier" {
		return false
	} else if funCall.Expression.ASTNode.(*AST.Identifier).Name != "assert" {
		return false
	}

	return true
}

func (cfg *CFG) _isModify(stmt *AST.ExpressionStatement) bool {
	// Assigment | FunctionCall | MemberAccess | BinaryOperation

	// Assignment
	if assign, ok := stmt.Expression.ASTNode.(*AST.Assignment); ok {

		// mapping: index access
		if lhs, ok := assign.LeftHandSide.ASTNode.(*AST.IndexAccess); ok {
			if idt, ok := lhs.BaseExpression.ASTNode.(*AST.Identifier); ok {
				return cfg._isStateVariable(idt.Name)
			}
		}
	}

	return false
}

func (cfg *CFG) _isFunctionCall(stmt *AST.ExpressionStatement) bool {
	if _, ok := stmt.Expression.ASTNode.(*AST.FunctionCall); !ok {
		return false
	}

	return true
}

func (cfg *CFG) _isStateVariable(v string) bool {
	return cfg.symbolTable.IsExistWithIdentifierOnly(v)
}
