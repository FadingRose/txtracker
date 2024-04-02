package main

import (
	"fmt"
	"txtracker/pkg/filehandler"
)

func main() {
	filehandler, err := filehandler.NewFileHandler("../../dataset/contracts")
	if err != nil {
		panic(err)
	}

	contracts, err := filehandler.GetContractList()
	fmt.Println(contracts)
	if err != nil {
		panic(err)
	}

}
