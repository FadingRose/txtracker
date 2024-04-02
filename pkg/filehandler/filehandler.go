package filehandler

import (
	"os"
	"path/filepath"
	"txtracker/pkg/common/models"
	"txtracker/pkg/logger"
)

// filehandler.go:
// 1. read smart contract file from dataset/contracts

type FileHandler interface {
	GetContractList() ([]string, error)
	GetContractData(contractName string) (string, error)
}

var (
	contracts []models.Contract
)

// - For testing: main.go -> ../../dataset/contracts
// - For build: txtracker -> ../../dataset/contracts
func init() {
	var DataPath string = "../../dataset/contracts"

	// iterate over all files in the contracts directory
	err := filepath.Walk(DataPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			contracts = append(contracts,
				models.Contract{
					ContractName: info.Name(),
				},
			)
		}
		return nil
	})
	if err != nil {
		logger.Fatal.Println("Error reading contract files:", err)
	}

}
