package ast

// `ContractKind`: "contract" | "interface" | "library"
type ContractKind string

const (
	ContractKind_Contract  ContractKind = "contract"
	ContractKind_Interface ContractKind = "interface"
	ContractKind_Library   ContractKind = "library"
)

type Mutability string

const (
	Mutability_Pure       Mutability = "pure"
	Mutability_View       Mutability = "view"
	Mutability_Nonpayable Mutability = "nonpayable"
	Mutability_Payable    Mutability = "payable"
)

type StorageLocation string

const (
	StorageLocation_Default  StorageLocation = "default"
	StorageLocation_Memory   StorageLocation = "memory"
	StorageLocation_Storage  StorageLocation = "storage"
	StorageLocation_Calldata StorageLocation = "calldata"
)

type Visibility string

const (
	Visibility_External Visibility = "external"
	Visibility_Internal Visibility = "internal"
	Visibility_Public   Visibility = "public"
	Visibility_Private  Visibility = "private"
)
