package ast

import "fmt"

type ASTNode interface {
	Attributes() *map[string]interface{}
	Constructor(data *map[string]interface{})
}

func NodeFactory(data map[string]interface{}) (*Common, error) {
	nodeType, ok := data["nodeType"].(string)
	if !ok || nodeType == "" {
		return nil, fmt.Errorf("nodeType not found")
	}
	src := data["src"].(string)
	if !ok || src == "" {
		return nil, fmt.Errorf("src not found")
	}
	id := int(data["id"].(float64))
	if !ok || id == 0 {
		return nil, fmt.Errorf("id not found or id is 0")
	}
	common := &Common{
		NodeType: nodeType,
		Src:      src,
		ID:       id,
	}
	common.ASTNode = astNodes[nodeType]()
	return common, nil
}

var astNodes = map[string]func() ASTNode{
	"SourceUnit":         func() ASTNode { return &SourceUnit{} },
	"ContractDefinition": func() ASTNode { return &ContractDefinition{} },
}
