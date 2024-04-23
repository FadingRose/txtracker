package cfg

import (
	"strings"
	AST "txtracker/pkg/ast"
	"txtracker/pkg/logger"
)

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
			return Assignment
		} else if cfg._isFunctionCall(stmt.ASTNode.(*AST.ExpressionStatement)) {
			if cfg._isEvent(stmt.ASTNode.(*AST.ExpressionStatement)) {
				return Emit
			}
			return FunctionCall
		} else {
			logger.Warning.Println("Unknown handled expression statement type:", stmt.NodeType)
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
	if _, ok := stmt.Expression.ASTNode.(*AST.Assignment); ok {

		return true
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

func (cfg *CFG) _isEvent(stmt *AST.ExpressionStatement) bool {
	if funCall, ok := stmt.Expression.ASTNode.(*AST.FunctionCall); !ok {
		return false
	} else if idt, ok := funCall.Expression.ASTNode.(*AST.Identifier); !ok {
		return false
	} else {
		ti := idt.TypeDescriptions.TypeIdentifier
		if strings.Contains(ti, "event") {
			return true
		}
	}

	return false
}
