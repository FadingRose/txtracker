package parser

import (
	"testing"
	"txtracker/pkg/parser"
)

func setupTestEnvironment() string {
	testPath := "./test_ast_dataset/0x0a3f9678d6b631386c2dd3de8809b48b0d1bbd56.sol.ast.json"
	return testPath
}

func TestASTParserImpl_ParseAST_JSON(t *testing.T) {
	DataPath := setupTestEnvironment()
	root := parser.NewASTParser().ParseAST_JSON(DataPath)

	if root == nil {
		t.Errorf("Expected root to be non-nil")
		return
	}

	if root.NodeType != "SourceUnit" {
		t.Errorf("Expected root to have name SourceUnit, got %s", root.NodeType)
	}

	if root.Children[0].NodeType != "PragmaDirective" {
		t.Errorf("Expected first child to be PragmaDirective, got %s", root.Children[0].NodeType)
	}

	if root.Children[1].ID != 97 {
		t.Errorf("Expected second child to have ID 97, got %d", root.Children[1].ID)
	}

	if root.Children[1].Children[0].ID != 34 {
		t.Errorf("Expected second child's first child to have ID 34, got %d", root.Children[1].Children[0].ID)
	}
}
