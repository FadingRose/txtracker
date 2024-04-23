package cfg

import (
	"strings"
	AST "txtracker/pkg/ast"
	"txtracker/pkg/logger"
	ST "txtracker/pkg/symbol_table"
)

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
	case Assignment:
		cfg._getModifyAndDependsSymbolsForAssignment(stmt, &modify, &depends, &declare)
	case Return:
		cfg._getModifyAndDependsSymbolsForReturn(stmt, &modify, &depends, &declare)
	case Emit:
		cfg._getSymbolsForEmit(stmt, &modify, &depends, &declare)
	default:
		logger.Warning.Println("Unhandled statement type:", _type)
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
		if initialValue := stmt.ASTNode.(*AST.VariableDeclarationStatement).GetInitialValue(); initialValue != nil {
			//fmt.Println("InitialValue:", initialValue)
			extractSymbolsFromExpression(initialValue, depends)
		}

	}

}

// recrusively extract symbols from the given expression
func extractSymbolsFromExpression(expr *AST.Common, symbols *[]ST.Symbol) {
	if expr == nil {
		return
	}

	switch expr.NodeType {
	case "Identifier":
		*symbols = append(*symbols, ST.Symbol{
			Namespace:  nil,
			Identifier: expr.ASTNode.(*AST.Identifier).Name,
			Type: func() ST.SymbolType {
				ts := expr.ASTNode.(*AST.Identifier).TypeDescriptions.TypeString
				// ts start with "function" means it is a function
				if strings.Split(ts, " ")[0] == "function" {
					return ST.Function
				}
				return ST.Unknown
			}(), // check the symbol table
		})
		return
	case "IndexAccess":
		extractSymbolsFromExpression(expr.ASTNode.(*AST.IndexAccess).BaseExpression, symbols)
		extractSymbolsFromExpression(expr.ASTNode.(*AST.IndexAccess).IndexExpression, symbols)
	case "MemberAccess":
		extractSymbolsFromExpression(expr.ASTNode.(*AST.MemberAccess).Expression, symbols)
	case "BinaryOperation":
		extractSymbolsFromExpression(expr.ASTNode.(*AST.BinaryOperation).LeftExpression, symbols)
		extractSymbolsFromExpression(expr.ASTNode.(*AST.BinaryOperation).RightExpression, symbols)
	case "FunctionCall":
		for _, arg := range expr.ASTNode.(*AST.FunctionCall).Arguments {
			extractSymbolsFromExpression(arg, symbols)
		}
		extractSymbolsFromExpression(expr.ASTNode.(*AST.FunctionCall).Expression, symbols)
	case "UnaryOperation":
		extractSymbolsFromExpression(expr.ASTNode.(*AST.UnaryOperation).SubExpression, symbols)
	case "Assignment":
		extractSymbolsFromExpression(expr.ASTNode.(*AST.Assignment).LeftHandSide, symbols)
		extractSymbolsFromExpression(expr.ASTNode.(*AST.Assignment).RightHandSide, symbols)
	case "Literal":
		return
	case "ElementaryTypeNameExpression":
		return
	case "NewExpression":
		return
	default:
		logger.Warning.Println("Unhandle expression type:", expr.NodeType)
	}
}

func (cfg *CFG) _getSymbolsForEmit(stmt *AST.Common, modify, depends, declare *[]ST.Symbol) {
	// stmt.NodeType() == "EmitStatement"
	emit := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall)
	*depends = append(*depends, ST.Symbol{
		Namespace:  *cfg.Visitor.CurrentNamespace,
		Identifier: emit.Expression.ASTNode.(*AST.Identifier).Name + "()",
		Type:       ST.Event,
	})
}

func (cfg *CFG) _getModifyAndDependsSymbolsForReturn(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	// stmt.NodeType() == "Return"

}

func (cfg *CFG) _getModifyAndDependsSymbolsForAssignment(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	// stmt.NodeType() == "ExpressionStatement"
	assignment := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.Assignment)
	extractSymbolsFromExpression(assignment.LeftHandSide, modify)
	extractSymbolsFromExpression(assignment.RightHandSide, depends)
}

func (cfg *CFG) _getModifyAndDependsSymbolsForAssert(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
}

func (cfg *CFG) _getModifyAndDependsSymbolsForRequire(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	arguents := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Arguments
	for _, arg := range arguents {
		extractSymbolsFromExpression(arg, depends)
	}
	// extractSymbolsFromExpression(stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Expression, depends)
}

func (cfg *CFG) _getModifyAndDependsSymbolsForFunctionCall(stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	arguments := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Arguments
	for _, arg := range arguments {
		extractSymbolsFromExpression(arg, depends)
	}
}
