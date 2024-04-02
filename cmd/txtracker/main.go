package main

import (
	"txtracker/pkg/compiler"
	"txtracker/pkg/filehandler"
)

func main() {
	filehandler, err := filehandler.NewFileHandler("../../dataset/contracts")
	if err != nil {
		panic(err)
	}

	// To AST
	compiler := compiler.NewSolidityCompiler()
	solFilePaths := filehandler.GetContractSolPathList()
	for _, path := range solFilePaths {
		err := compiler.SolidityToAST_JSON(path)
		if err != nil {
			panic(err)
		}

	}

}
