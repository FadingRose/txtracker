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

type StateMutability string

const (
	StateMutability_Pure       StateMutability = "pure"
	StateMutability_View       StateMutability = "view"
	StateMutability_Nonpayable StateMutability = "nonpayable"
	StateMutability_Payable    StateMutability = "payable"
)

type FunctionKind string

const (
	FunctionKind_Constructor  FunctionKind = "constructor"
	FunctionKind_Fallback     FunctionKind = "fallback"
	FunctionKind_Function     FunctionKind = "function"
	FunctionKind_Receive      FunctionKind = "receive"
	FunctionKind_FreeFunction FunctionKind = "freeFunction"
)

type ModifierKind string

const (
	ModifierKind_ModifierInvocation       = "modifierInvocation"
	ModifierKind_BaseConstructorSpecifier = "baseConstructorSpecifier"
)

type Operator string

const (
	Operator_Addition           Operator = "+"
	Operator_Subtraction        Operator = "-"
	Operator_Multiplication     Operator = "*"
	Operator_Division           Operator = "/"
	Operator_Modulo             Operator = "%"
	Operator_Exponentiation     Operator = "**"
	Operator_And                Operator = "&&"
	Operator_Or                 Operator = "||"
	Operator_StrictNotEqual     Operator = "!="
	Operator_StrictEqual        Operator = "=="
	Operator_LessThan           Operator = "<"
	Operator_LessThanOrEqual    Operator = "<="
	Operator_GreaterThan        Operator = ">"
	Operator_GreaterThanOrEqual Operator = ">="
	Operator_BitwiseXor         Operator = "^"
	Operator_BitwiseAnd         Operator = "&"
	Operator_BitwiseOr          Operator = "|"
	Operator_ShiftLeft          Operator = "<<"
	Operator_ShiftRight         Operator = ">>"
)

type AssignmentOperator string

const (
	AssignmentOperator_Assignment     AssignmentOperator = "="
	AssignmentOperator_Addition       AssignmentOperator = "+="
	AssignmentOperator_Subtraction    AssignmentOperator = "-="
	AssignmentOperator_Multiplication AssignmentOperator = "*="
	AssignmentOperator_Division       AssignmentOperator = "/="
	AssignmentOperator_Modulo         AssignmentOperator = "%="
	AssignmentOperator_BitwiseOr      AssignmentOperator = "|="
	AssignmentOperator_BitwiseAnd     AssignmentOperator = "&="
	AssignmentOperator_BitwiseXor     AssignmentOperator = "^="
	AssignmentOperator_ShiftLeft      AssignmentOperator = "<<="
	AssignmentOperator_ShiftRight     AssignmentOperator = ">>="
)

type LiteralKind string

const (
	LiteralKind_Boolean       LiteralKind = "bool"
	LiteralKind_String        LiteralKind = "string"
	LiteralKind_Integer       LiteralKind = "number"
	LiteralKind_HexString     LiteralKind = "hexString"
	LiteralKind_UnicodeString LiteralKind = "unicodeString"
)

type Subdenomination string

const (
	Subdenomination_Wei     Subdenomination = "wei"
	Subdenomination_Gwei    Subdenomination = "gwei"
	Subdenomination_Ether   Subdenomination = "ether"
	Subdenomination_Finny   Subdenomination = "finny"
	Subdenomination_Szabo   Subdenomination = "szabo"
	Subdenomination_Weeks   Subdenomination = "weeks"
	Subdenomination_Days    Subdenomination = "days"
	Subdenomination_Hours   Subdenomination = "hours"
	Subdenomination_Minutes Subdenomination = "minutes"
	Subdenomination_Seconds Subdenomination = "seconds"
)

type FunctionCallKind string

const (
	FunctionCallKind_FunctionCall          FunctionCallKind = "functionCall"
	FunctionCallKind_TypeConversion        FunctionCallKind = "typeConversion"
	FunctionCallKind_StructConstructorCall FunctionCallKind = "structConstructorCall"
)
