package filehandler

import (
	"os"
	"path/filepath"
	"strings"
	"txtracker/pkg/common/models"
	"txtracker/pkg/logger"
)

// filehandler.go:
// 1. read smart contract file from dataset/contracts

type FileHandler interface {
	GetContractList() ([]string, error)
	GetContractData(contractName string) (string, error)
}

type SolidityFileHandler struct {
	DataPath  string
	contracts []models.Contract
}

func NewSolidityFileHandler(dataPath string) *SolidityFileHandler {
	handler := &SolidityFileHandler{
		DataPath: dataPath,
	}
	handler.loadContracts()
	return handler
}

// loadContracts replaces the init function and is called within the constructor
func (s *SolidityFileHandler) loadContracts() {
	err := filepath.Walk(s.DataPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileName, fileExtension := filepath.Base(path), filepath.Ext(path)

			var solCode models.SoliditySourceCode
			var evmCode models.EVMByteCode

			switch fileExtension {
			case ".sol":
				solCode = readThenDumpFileContent(path)
			case ".evm":
				evmCode = readThenDumpFileContent(path)
			default:
				logger.Fatal.Println("Unknown file extension:", fileExtension)
			}

			s.contracts = append(s.contracts,
				models.Contract{
					ContractName: strings.Split(fileName, ".")[0],
					SolidityCode: models.SolidityCode{SourceCode: solCode},
					EVMCode:      models.EVMCode{ByteCode: evmCode},
				},
			)
		}
		return nil
	})
	if err != nil {
		logger.Fatal.Println("Error reading contract files:", err)
	}
}

// readThenDumpFileContent modified to return the file content directly
func readThenDumpFileContent(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Fatal.Println("Error reading and dumping file:", err)
		return ""
	}
	return string(content)
}

// GetContractList and GetContractData methods remain largely unchanged
func (s SolidityFileHandler) GetContractList() ([]string, error) {
	var contractList []string
	for _, contract := range s.contracts {
		contractList = append(contractList, contract.ContractName)
	}
	return contractList, nil
}

func (s SolidityFileHandler) GetContractData(contractName string) (models.SoliditySourceCode, models.EVMByteCode, error) {
	for _, contract := range s.contracts {
		if contract.ContractName == contractName {
			return contract.SolidityCode.SourceCode, contract.EVMCode.ByteCode, nil // Assuming you want to return the Solidity source code
		}
	}
	return "", "", nil // Consider returning an error if the contract is not found
}
