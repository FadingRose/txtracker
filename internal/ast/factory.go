package ast

import (
	"txtracker/internal/logger"
)

func NodeFactory(data map[string]interface{}) *Common {
	common, nodeType := commonFactory(data)
	common.ASTNode = astNodes[nodeType]()
	return common
}

var astNodes = map[string]func() ASTNode{
	"SourceUnit":          func() ASTNode { return &SourceUnit{} },
	"ContractDefinition":  func() ASTNode { return &ContractDefinition{} },
	"PragmaDirective":     func() ASTNode { return &PragmaDirective{} },
	"FunctionDefinition":  func() ASTNode { return &FunctionDefinition{} },
	"VariableDeclaration": func() ASTNode { return &VariableDeclaration{} },
	"UsingForDirective":   func() ASTNode { return &UsingForDirective{} },

	// Conditional
	"Conditional": func() ASTNode { return &Conditional{} },

	// Enum
	"EnumDefinition": func() ASTNode { return &EnumDefinition{} },
	"EnumValue":      func() ASTNode { return &EnumValue{} },

	// Events
	"EventDefinition": func() ASTNode { return &EventDefinition{} },

	// Functions
	"FunctionCall": func() ASTNode { return &FunctionCall{} },

	// Parameters
	"ParameterList": func() ASTNode { return &ParameterList{} },

	// Modifiers
	"ModifierDefinition": func() ASTNode { return &ModifierDefinition{} },
	"ModifierInvocation": func() ASTNode { return &ModifierInvocation{} },

	// Statements
	"Block":                        func() ASTNode { return &Block{} },
	"IfStatement":                  func() ASTNode { return &IfStatement{} },
	"Return":                       func() ASTNode { return &Return{} },
	"VariableDeclarationStatement": func() ASTNode { return &VariableDeclarationStatement{} },
	"ExpressionStatement":          func() ASTNode { return &ExpressionStatement{} },
	"PlaceholderStatement":         func() ASTNode { return &PlaceholderStatement{} },
	"ForStatement":                 func() ASTNode { return &ForStatement{} },
	"Break":                        func() ASTNode { return &Break{} },

	// Expressions
	"BinaryOperation":              func() ASTNode { return &BinaryOperation{} },
	"Identifier":                   func() ASTNode { return &Identifier{} },
	"Literal":                      func() ASTNode { return &Literal{} },
	"Assignment":                   func() ASTNode { return &Assignment{} },
	"MemberAccess":                 func() ASTNode { return &MemberAccess{} },
	"IndexAccess":                  func() ASTNode { return &IndexAccess{} },
	"UnaryOperation":               func() ASTNode { return &UnaryOperation{} },
	"ElementaryTypeNameExpression": func() ASTNode { return &ElementaryTypeNameExpression{} },
	"TupleExpression":              func() ASTNode { return &TupleExpression{} },
	"NewExpression":                func() ASTNode { return &NewExpression{} },

	// TypeNames
	"ElementaryTypeName":  func() ASTNode { return &ElementaryTypeName{} },
	"UserDefinedTypeName": func() ASTNode { return &UserDefinedTypeName{} },
	"Mapping":             func() ASTNode { return &Mapping{} },
	"StructDefinition":    func() ASTNode { return &StructDefinition{} },
	"ArrayTypeName":       func() ASTNode { return &ArrayTypeName{} },
}

func commonFactory(data map[string]interface{}) (*Common, string) {
	nodeType, ok := data["nodeType"].(string)
	if !ok || nodeType == "" {
		logger.Fatal.Println("nodeType not found")
		panic("nodeType not found")
	}
	src := data["src"].(string)
	if !ok || src == "" {
		logger.Fatal.Println("src not found")
		panic("src not found")
	}
	id := int(data["id"].(float64))
	if !ok || id == 0 {
		logger.Fatal.Println("id not found or id is 0")
		panic("id not found or id is 0")
	}
	common := &Common{
		NodeType: nodeType,
		Src:      src,
		ID:       id,
	}
	return common, nodeType
}
