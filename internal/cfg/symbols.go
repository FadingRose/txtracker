package cfg

import (
	"strings"
	AST "txtracker/internal/ast"
	"txtracker/internal/logger"
	ST "txtracker/internal/symbol_table"
)

// returns the symbols that are modified and depends on the given variable declaration statement
func (cfg *CFG) _getModifyAndDependsSymbols(stmt *AST.Common, _type StatementType) ([]ST.Symbol, []ST.Symbol, []ST.Symbol) {
	var modify, depends, declare []ST.Symbol

	handlers := map[StatementType]SymbolHandler{
		VariableDeclaration: &VariableDeclarationHandler{},
		Emit:                &EmitHandler{},
		Return:              &ReturnHandler{},
		Assignment:          &AssignmentHandler{},
		Assert:              &AssertHandler{},
		Require:             &RequireHandler{},
		FunctionCall:        &FunctionCallHandler{},
	}
	if handler, ok := handlers[_type]; ok {
		handler.GetSymbols(*cfg.Visitor.CurrentNamespace, stmt, &modify, &depends, &declare)
	} else {
		logger.Warning.Println("Unhandled statement type:", _type)
	}
	return modify, depends, declare
}

type SymbolHandler interface {
	GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends, declare *[]ST.Symbol)
}

type VariableDeclarationHandler struct {
}

func (h *VariableDeclarationHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	// stmt.NodeType() == "VariableDeclareStatement"
	declarations, statevars := stmt.ASTNode.(*AST.VariableDeclarationStatement).GetDeclarations()
	for i, declaration := range declarations {
		decl := &ST.Symbol{
			Namespace:  namespace,
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

type EmitHandler struct {
}

func (h *EmitHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends, declare *[]ST.Symbol) {
	// stmt.NodeType() == "EmitStatement"
	emit := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall)
	*depends = append(*depends, ST.Symbol{
		Namespace:  namespace,
		Identifier: emit.Expression.ASTNode.(*AST.Identifier).Name + "()",
		Type:       ST.Event,
	})
}

type ReturnHandler struct {
}

func (h *ReturnHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	expr := stmt.ASTNode.(*AST.Return).Expression
	if expr.NodeType == "FunctionCall" {
		extractFuncSymbols(namespace, expr, depends)
	} else {
		extractSymbolsFromExpression(expr, depends)
	}
}

type AssignmentHandler struct {
}

func (h *AssignmentHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	// stmt.NodeType() == "ExpressionStatement"
	assignment := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.Assignment)
	extractSymbolsFromExpression(assignment.LeftHandSide, modify)
	extractSymbolsFromExpression(assignment.RightHandSide, depends)
}

type AssertHandler struct {
}

func (h *AssertHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
}

type RequireHandler struct {
}

func (h *RequireHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	arguents := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Arguments
	for _, arg := range arguents {
		extractSymbolsFromExpression(arg, depends)
	}
	// extractSymbolsFromExpression(stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Expression, depends)
}

type FunctionCallHandler struct {
}

func (h *FunctionCallHandler) GetSymbols(namespace ST.Namespace, stmt *AST.Common, modify, depends *[]ST.Symbol, declare *[]ST.Symbol) {
	arguments := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Arguments
	for _, arg := range arguments {
		extractSymbolsFromExpression(arg, depends)
	}

	funcRef := stmt.ASTNode.(*AST.ExpressionStatement).Expression.ASTNode.(*AST.FunctionCall).Expression

	// Member access upgradeAgent.upgradeFrom(...)
	if funcRef.NodeType == "MemberAccess" || funcRef.NodeType == "Identifier" {
		extractFuncSymbols(namespace, funcRef, declare)
	}

}

// helper function:
// recrusively extract symbols from the given function reference
func extractFuncSymbols(namespace ST.Namespace, expr *AST.Common, symbols *[]ST.Symbol) {
	if expr == nil {
		return
	}

	switch expr.NodeType {
	case "Identifier":
		*symbols = append(*symbols, ST.Symbol{
			Namespace:  namespace,
			Identifier: expr.ASTNode.(*AST.Identifier).Name,
			Type:       ST.Function,
		})
		return
	case "MemberAccess":
		*symbols = append(*symbols, ST.Symbol{
			Namespace:  namespace,
			Identifier: expr.ASTNode.(*AST.MemberAccess).MemberName,
			Type:       ST.Function,
		})
		extractFuncSymbols(namespace, expr.ASTNode.(*AST.MemberAccess).Expression, symbols)
	case "FunctionCall":
		extractFuncSymbols(namespace, expr.ASTNode.(*AST.FunctionCall).Expression, symbols)
	default:
		logger.Warning.Println("Unhandle FunctionCall internal expression type:", expr.NodeType)
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
