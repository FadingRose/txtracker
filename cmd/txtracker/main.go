package main

import (
	"fmt"
	"txtracker/pkg/compiler"
	"txtracker/pkg/filehandler"
	"txtracker/pkg/parser"
	"txtracker/pkg/printer"
)

func main() {
	filehandler, err := filehandler.NewFileHandler("../../dataset/contracts")
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
			panic(err)
		}
		astFilePath := path + ".ast.json"
		root := parser.ParseAST_JSON(astFilePath)
		ast_printer := printer.NewASTPrinter(root)
		ast_printer.PrintAST()
	}

}
