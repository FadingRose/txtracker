package filehandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"txtracker/internal/common/models"
	"txtracker/internal/logger"
)

// filehandler.go:
// 1. read smart contract file from dataset/contracts

type FileHandler interface {
	GetContractList() ([]string, error)
	GetContractSolPathList() []string
	GetContractData(contractName string) (models.SoliditySourceCode, models.EVMByteCode, error)
	WriteASTFile(contractName string, data string) error
}

type SolidityFileHandler struct {
	DataPath   string
	contracts  []models.Contract
	isSpecific bool
}

// loadContracts replaces the init function and is called within the constructor

func NewFileHandler(dataPath string, specific_file string) (FileHandler, error) {
	dataPath = filepath.Join(dataPath, specific_file)

	handler := &SolidityFileHandler{
		DataPath: dataPath,
	}

	if specific_file != "" {
		handler.isSpecific = true
	} else {
		handler.isSpecific = false
	}

	fmt.Println("NewFileHandler called with path:", dataPath)
	handler.loadContracts()

	return handler, nil
}

func (s *SolidityFileHandler) loadContracts() {
	var solCode models.SoliditySourceCode
	var evmCode models.EVMByteCode

	if !s.isSpecific {
		// Walk through the directory and read all files
		err := filepath.Walk(s.DataPath, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fileName, fileExtension := filepath.Base(path), filepath.Ext(path)

				if fileExtension == ".sol" {
					s.contracts = append(s.contracts,
						models.Contract{
							ContractName: strings.Split(fileName, ".")[0],
						},
					)
				}
			}
			return nil
		})
		if err != nil {
			logger.Fatal.Println("Error reading contract files:", err)
			panic(err)
		}

	} else {
		// Only load the specific file
		fileName, fileExtension := filepath.Base(s.DataPath), filepath.Ext(s.DataPath)
		switch fileExtension {
		case ".sol":
			solCode = readThenDumpFileContent(s.DataPath)
		case ".evm":
			evmCode = readThenDumpFileContent(s.DataPath)
		default:
			logger.Warning.Println("Unknown file extension:", fileExtension)
		}

		s.contracts = append(s.contracts,
			models.Contract{
				ContractName: strings.Split(fileName, ".")[0],
				SolidityCode: models.SolidityCode{SourceCode: solCode},
				EVMCode:      models.EVMCode{ByteCode: evmCode},
			},
		)
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

func (s SolidityFileHandler) GetContractSolPathList() []string {
	var contractPathList []string
	if s.isSpecific {
		contractPathList = append(contractPathList, s.DataPath)
	} else {
		for _, contract := range s.contracts {
			contractPathList = append(contractPathList,
				filepath.Join(s.DataPath, contract.ContractName+".sol"))
		}
	}

	return contractPathList
}

func (s SolidityFileHandler) WriteASTFile(contractName string, data string) error {

	return nil
}
