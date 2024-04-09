package ast

import "txtracker/pkg/logger"

type ASTNode interface {
	Attributes() *map[string]interface{}
	Constructor(data *map[string]interface{})
}

// Common is the base struct for all AST nodes.
type Common struct {
	NodeType string `json:"nodeType"`
	Src      string `json:"src"` // location in the source code
	ID       int    `json:"id"`
	Parent   *Common
	Children []*Common //refactor here
	ASTNode  ASTNode
	// ASTNode provides methods:
	// Attributes() *map[string]interface{}
	// Constructor(data *map[string]interface{})
}

func (c *Common) Instance() *Common {
	return c
}

func (c *Common) AddChild(child *Common) {
	c.Children = append(c.Children, child)
}

func (c *Common) SetParent(parent *Common) {
	c.Parent = parent
}

// Type Conversion
func (c *Common) ToBlock() *Block {
	if c.NodeType == "Block" {
		return &Block{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to Block")
		panic("Cannot convert to Block")
	}
}

func (c *Common) ToParameterList() *ParameterList {
	if c.NodeType == "ParameterList" {
		return &ParameterList{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to ParameterList")
		panic("Cannot convert to ParameterList")
	}
}

func (c *Common) ToStructuredDocumentation() *StructuredDocumentation {
	if c.NodeType == "Documentation" {
		return &StructuredDocumentation{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to StructuredDocumentation")
		panic("Cannot convert to StructuredDocumentation")
	}
}

func (c *Common) ToEnumValue() *EnumValue {
	if c.NodeType == "EnumValue" {
		return &EnumValue{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to EnumValue")
		panic("Cannot convert to EnumValue")
	}
}

func (c *Common) ToModifierInvocation() *ModifierInvocation {
	if c.NodeType == "ModifierInvocation" {
		return &ModifierInvocation{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to ModifierInvocation")
		panic("Cannot convert to ModifierInvocation")
	}
}

func (c *Common) ToOverrideSpecifier() *OverrideSpecifier {
	if c.NodeType == "OverrideSpecifier" {
		return &OverrideSpecifier{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to ModifierDefinition")
		panic("Cannot convert to ModifierDefinition")
	}
}

func (c *Common) ToTypeDescriptions() *TypeDescriptions {
	if c.NodeType == "TypeDescriptions" {
		return &TypeDescriptions{
			Common: *c,
		}
	} else {
		logger.Fatal.Println("Cannot convert to TypeDescriptions")
		panic("Cannot convert to TypeDescriptions")
	}
}
