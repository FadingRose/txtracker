package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"txtracker/pkg/common/models"
	"txtracker/pkg/logger"
)

type ASTParser interface {
	ParseAST_JSON(astFilePath string) error
}

type ASTParserImpl struct {
}

func NewASTParser() ASTParser {
	return &ASTParserImpl{}
}

func (a *ASTParserImpl) ParseAST_JSON(astFilePath string) error {
	logger.Info.Println("ParseAST_JSON called with path:", astFilePath)
	JsonData := collectJSON(astFilePath)

	var root models.Common
	parseAST(JsonData, &root)

	return nil
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
		dest.SetParent(root.Instace())
		root.AddChild(dest.Instace())

		// Special fields set at here
		models.Reflect_constructor(&dest, &data)

		// // Reflectively set the fields of the ASTNode
		// v := reflect.ValueOf(dest)
		// if v.Kind() != reflect.Ptr || v.IsNil() {
		// 	logger.Fatal.Fatalf("Not a pointer: %v", v.Kind())
		// 	panic("Reflect failed: not a pointer")
		// }

		// v = v.Elem() // & -> *

		// t := v.Type()
		// for i := 0; i < v.NumField(); i++ {

		// 	jsonTag := t.Field(i).Tag.Get("json")

		// 	if jsonTag == "nodes" || jsonTag == "" {
		// 		continue
		// 	}

		// 	if value, ok := data[jsonTag]; ok {
		// 		fieldValue := v.Field(i)
		// 		if fieldValue.IsValid() && fieldValue.CanSet() {
		// 			switch fieldValue.Kind() {
		// 			case reflect.String:
		// 				fieldValue.SetString(value.(string))
		// 			case reflect.Int:
		// 				fieldValue.SetInt(int64(value.(float64)))
		// 			case reflect.Bool:
		// 				fieldValue.SetBool(value.(bool))
		// 			case reflect.Map:
		// 				_map_handler(&value, fieldValue)
		// 			case reflect.Slice:
		// 				_slice_handler(&value, fieldValue)
		// 			default:
		// 				logger.Fatal.Fatalf("Unknown type: %v", fieldValue.Kind())
		// 				panic("Unknown type")
		// 			}
		// 		}
		// 	}
		// }

		// Recursively parse children

		childs, ok := data["nodes"].([]interface{})
		if !ok {
			return
		}
		for _, child := range childs {
			parseAST(child, dest.Instace())
		}

	case []interface{}:
		for _, value := range data {
			parseAST(value, root)
		}
	default:
		logger.Fatal.Fatalf("Unknown type: %v", data)
	}
}

func _map_handler(_map *interface{}, fieldValue reflect.Value) {
	// Handle the case where the field is a map
	res := make(map[string]int)
	for key, values := range (*_map).(map[string]interface{}) {
		v := values.([]interface{})[0].(float64)
		res[key] = int(v)
	}
	fieldValue.Set(reflect.ValueOf(res))
}

func _slice_handler(_slice *interface{}, fieldValue reflect.Value) {
	// Handle the case where the field is a slice

	// WARNING: WE ASSUME THAT THE SLICE IS A SLICE OF STRINGS
	_type := fieldValue.Type().Elem().Kind().String()
	fmt.Println("Type:", _type)
	if len((*_slice).([]interface{})) == 0 {
		fieldValue.Set(reflect.ValueOf([]string{}))
		return
	}

	res := make([]string, 0)
	for _, value := range (*_slice).([]interface{}) {
		res = append(res, value.(string))
	}

	fieldValue.Set(reflect.ValueOf(res))
}
