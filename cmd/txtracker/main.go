package main

import (
	"fmt"
	"os"
	CFG "txtracker/internal/cfg"
	"txtracker/internal/compiler"
	"txtracker/internal/filehandler"
	"txtracker/internal/logger"
	"txtracker/internal/parser"
	"txtracker/internal/printer"
	symboltable "txtracker/internal/symbol_table"
)

func main() {

	var SPECIFIC_CONTRACT string
	var PRINTER PrinterType
	SPECIFIC_CONTRACT, PRINTER = cmd(os.Args)
	// if len(os.Args) >= 2 {
	// 	SPECIFIC_CONTRACT = os.Args[1]
	// } else {
	// 	SPECIFIC_CONTRACT = ""
	// }

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
		symbol_table := symboltable.NewGlobalSymbolTable(root)
		cfg := CFG.NewCFG(root, symbol_table)

		cfg_printer := printer.NewCFGPrinter(cfg)

		switch PRINTER {
		case CFG_PRINTER:
			cfg_printer.Print()
		}

	}

}

type PrinterType string

const (
	CFG_PRINTER       PrinterType = "cfg"
	CALLGRAPH_PRINTER PrinterType = "callgraph"
)

type SPECIFIC_CONTRACT = string

func cmd(args []string) (SPECIFIC_CONTRACT, PrinterType) {
	var SPECIFIC_CONTRACT string
	var PRINTER PrinterType
	SPECIFIC_CONTRACT, PRINTER = "", CFG_PRINTER
	if len(args) >= 2 {
		SPECIFIC_CONTRACT = args[1]
	} else {
		SPECIFIC_CONTRACT = ""
	}
	return SPECIFIC_CONTRACT, PRINTER
}
