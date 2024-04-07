package main

import (
	"fmt"
	"os"
	"txtracker/pkg/compiler"
	"txtracker/pkg/filehandler"
	"txtracker/pkg/logger"
	"txtracker/pkg/parser"
	"txtracker/pkg/printer"
)

func main() {

	var SPECIFIC_CONTRACT string
	if len(os.Args) >= 2 {
		SPECIFIC_CONTRACT = os.Args[1]
	} else {
		SPECIFIC_CONTRACT = ""

	}

	filehandler, err := filehandler.NewFileHandler("../../dataset/contracts", SPECIFIC_CONTRACT)
	if err != nil {
		panic(err)
	}

	// Compile then Parse AST
	compiler := compiler.NewSolidityCompiler()
	parser := parser.NewASTParser()
	solFilePaths := filehandler.GetContractSolPathList()
	for _, path := range solFilePaths {
		fmt.Println("Processing:", path)
		err := compiler.SolidityToAST_JSON(path)
		if err != nil {
			logger.Fatal.Println("Error compiling:", err)
			panic(err)
		}
		astFilePath := path + ".ast.json"
		root := parser.ParseAST_JSON(astFilePath)
		ast_printer := printer.NewASTPrinter(root)
		ast_printer.PrintAST()
	}

}
