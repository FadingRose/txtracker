package ast

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
