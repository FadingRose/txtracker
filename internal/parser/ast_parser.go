package parser

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"txtracker/internal/ast"
	"txtracker/internal/logger"
)

type ASTParser interface {
	ParseAST_JSON(astFilePath string) *ast.Common
}

type ASTParserImpl struct {
}

func NewASTParser() ASTParser {
	return &ASTParserImpl{}
}

func (a *ASTParserImpl) ParseAST_JSON(astFilePath string) *ast.Common {
	logger.Info.Println("ParseAST_JSON called with path:", astFilePath)
	JsonData := collectJSON(astFilePath)

	var root ast.Common
	parseAST(JsonData, &root)

	root = *root.Children[0]
	root.SetParent(&root)
	return &root
}

func collectJSON(filePath string) interface{} {
	// read file, collect lines, and return
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatal.Println("Error reading file:", err)
		return nil
	}
	defer file.Close()

	// Delete Until meet the first '{'
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "{") {
			break
		}
	}

	// Collect the rest of the lines
	var builder strings.Builder
	builder.WriteString("{\n") // Add the first '{'
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Fatal.Println("Error scanning file:", err)
	}

	cleanJSON := builder.String()
	var result interface{}

	// Unmarshal the JSON data
	err = json.Unmarshal([]byte(cleanJSON), &result)
	if err != nil {
		logger.Fatal.Println("Error unmarshalling JSON data:", err)
	}
	return result
}

func parseAST(jsonNode interface{}, root *ast.Common) {

	switch data := jsonNode.(type) {
	case map[string]interface{}:
		// Try create ASTNode from JSON
		// Common fields set at here
		dest := ast.NodeFactory(data)

		// Set relationship between parent and child
		dest.SetParent(root.Instance())
		root.AddChild(dest.Instance())

		// Special fields set at here
		dest.ASTNode.Constructor(&data)

		// Recursively parse children from "nodes" or "body"
		childs, ok := data["nodes"].([]interface{})

		if !ok {
			return
		}
		for _, child := range childs {
			logger.Info.Println("Visiting child:" + child.(map[string]interface{})["nodeType"].(string))
			parseAST(child, dest.Instance())
		}
		return

	case []interface{}:
		for _, value := range data {
			logger.Info.Println("Visiting array element")
			parseAST(value, root)
		}
		return
	default:
		logger.Fatal.Fatalf("Unknown type: %v", data)
		panic("Unknown type")
	}
}
