package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"txtracker/pkg/common/models"
	"txtracker/pkg/logger"
)

type ASTParser interface {
	ParseAST_JSON(astFilePath string) *models.Common
}

type ASTParserImpl struct {
}

func NewASTParser() ASTParser {
	return &ASTParserImpl{}
}

func (a *ASTParserImpl) ParseAST_JSON(astFilePath string) *models.Common {
	logger.Info.Println("ParseAST_JSON called with path:", astFilePath)
	JsonData := collectJSON(astFilePath)

	var root models.Common
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

func parseAST(jsonNode interface{}, root *models.Common) {

	switch data := jsonNode.(type) {
	case map[string]interface{}:
		// Try create ASTNode from JSON
		// Common fields set at here
		dest, err := models.StringToASTNode(data)
		if err != nil {
			fmt.Println("Error converting JSON to ASTNode:", err)
		}

		// Set relationship between parent and child
		dest.SetParent(root.Instance())
		root.AddChild(dest.Instance())

		// Special fields set at here
		dest.Constructor(&data)

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
