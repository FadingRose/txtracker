package models

type SoliditySourceCode = string
type EVMByteCode = string

type Contract struct {
	ContractName string       `json:"contractName"`
	SolidityCode SolidityCode `json:"solidityCode"`
	EVMCode      EVMCode      `json:"evmCode"`
}

type SolidityCode struct {
	SourceCode SoliditySourceCode
}

type EVMCode struct {
	ByteCode EVMByteCode
}
