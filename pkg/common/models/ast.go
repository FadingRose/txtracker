package models

// ast.go: corresponds to the AST JSON file generated by the Solidity compiler.
type ASTNode struct {
	NodeType string `json:"nodeType"`
	Src      string `json:"src"` // location in the source code
	ID       int    `json:"id"`
	Parent   *ASTNode
	Children []*ASTNode
}

// Top-level node for a .sol file
type SourceUnit struct {
	ASTNode
	Nodes []ASTNode `json:"nodes"`
}

// Compiler version node
type PragmaDirective struct {
	ASTNode
	Literals []string `json:"literals"`
}

type ContractDefinition struct {
	ASTNode
	Name                     string                 `json:"name"`
	BaseContracts            []InheritanceSpecifier `json:"baseContracts"`
	ContractDependenceies    []string               `json:"contractDependencies"`
	LinearizaedBaseContracts []string               `json:"linearizedBaseContracts"`
	SubNodes                 []ASTNode              `json:"nodes"`
	ContractKind             string                 `json:"contractKind"`
}

// Block node, which contains a list of statements between '{' and '}'
type Block struct {
	ASTNode
	Statements []ExpressionStatement `json:"statements"`
}

type Identifier struct {
	ASTNode
	ArgumentTypes          []string         `json:"argumentTypes"`
	OverloadedDeclarations []string         `json:"overloadedDeclarations"`
	ReferencedDeclaration  int              `json:"referencedDeclaration"`
	TypeDescriptions       TypeDescriptions `json:"typeDescriptions"`
}

type TypeDescriptions struct {
	TypeString     string `json:"typeString"`
	TypeIdentifier string `json:"typeIdentifier"`
}

type BinaryOperation struct {
	ArgumentTypes    []string         `json:"argumentTypes"`
	CommonType       TypeDescriptions `json:"commonType"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	LeftExpression   ASTNode          `json:"leftExpression"`
	Operator         string           `json:"operator"`
	RightExpression  ASTNode          `json:"rightExpression"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type Literal struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	HexValue         string           `json:"hexValue"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	Kind             string           `json:"kind"`
	LValueRequested  bool             `json:"lValueRequested"`
	Subdenomination  string           `json:"subdenomination"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
	Value            string           `json:"value"`
}

type IfStatement struct {
	ASTNode
	Condition BinaryOperation `json:"condition"`
	FalseBody Block           `json:"falseBody"`
	TrueBody  Block           `json:"trueBody"`
}

type Return struct {
	ASTNode
	Expression               ASTNode `json:"expression"`
	FunctionReturnParameters int     `json:"functionReturnParameters"`
}

type VariableDeclaration struct {
	ASTNode
	Constant         bool               `json:"constant"`
	Name             string             `json:"name"`
	Scope            int                `json:"scope"`
	StateVariable    bool               `json:"stateVariable"`
	StorageLocation  string             `json:"storageLocation"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
	TypeName         ElementaryTypeName `json:"typeName"`
	Value            ASTNode            `json:"value"`
	Visibility       string             `json:"visibility"`
}

type ElementaryTypeName struct {
	ASTNode
	Name             string           `json:"name"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type VariableDeclarationStatement struct {
	ASTNode
	Assignments   []int                 `json:"assignments"`
	Declarrations []VariableDeclaration `json:"declarations"`
	InitialValue  ASTNode               `json:"initialValue"`
}

type FunctionCall struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	Arguments        []ASTNode        `json:"arguments"`
	Expression       ASTNode          `json:"expression"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	Kind             string           `json:"kind"`
	LValueRequested  bool             `json:"lValueRequested"`
	Names            []string         `json:"names"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type ExpressionStatement struct {
	ASTNode
	Expression ASTNode `json:"expression"`
}

type FunctionDefinition struct {
	ASTNode
	Implemented      bool          `json:"implemented"`
	IsConstructor    bool          `json:"isConstructor"`
	IsDeclaredConst  bool          `json:"isDeclaredConst"`
	Modifiers        []ASTNode     `json:"modifiers"`
	Name             string        `json:"name"`
	Parameters       ParameterList `json:"parameters"`
	Payable          bool          `json:"payable"`
	ReturnParameters ParameterList `json:"returnParameters"`
	Scope            int           `json:"scope"`
	StateMutability  string        `json:"stateMutability"`
	SuperFunction    string        `json:"superFunction"`
	Visibility       string        `json:"visibility"`
}

type ParameterList struct {
	ASTNode
	Parameters []ASTNode `json:"parameters"`
}

type EventDefinition struct {
	Anonymous  bool          `json:"anonymous"`
	Name       string        `json:"name"`
	Parameters ParameterList `json:"parameters"`
}

type Assignment struct {
	ArgumentTypes    []string         `json:"argumentTypes"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	LeftHandSide     ASTNode          `json:"leftHandSide"`
	Operator         string           `json:"operator"`
	RightHandSide    ASTNode          `json:"rightHandSide"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type MemberAccess struct {
	ArgumentTypes         []string         `json:"argumentTypes"`
	Expression            ASTNode          `json:"expression"`
	IsConstant            bool             `json:"isConstant"`
	IsLValue              bool             `json:"isLValue"`
	IsPure                bool             `json:"isPure"`
	LValueRequested       bool             `json:"lValueRequested"`
	MemberName            string           `json:"memberName"`
	ReferencedDeclaration int              `json:"referencedDeclaration"`
	TypeDescriptions      TypeDescriptions `json:"typeDescriptions"`
}

type PlacehoderStatement struct {
	ASTNode
}

type ModifierDefinition struct {
	ASTNode
	Name       string        `json:"name"`
	Parameters ParameterList `json:"parameters"`
	Visibility string        `json:"visibility"`
}

type ElementaryTypeNameExpression struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
	TypeName         string           `json:"typeName"`
}

type ModifierInvocation struct {
	ASTNode
	Arguments    []ASTNode `json:"arguments"`
	ModifierName ASTNode   `json:"modifierName"`
}

type UserDefinedTypeName struct {
	ASTNode
	ContractScope        string           `json:"contractScope"`
	ReferenceDeclaration int              `json:"referencedDeclaration"`
	TypeDescriptions     TypeDescriptions `json:"typeDescriptions"`
}

type InheritanceSpecifier struct {
	ASTNode
	Arguments []ASTNode `json:"arguments"`
	BaseName  ASTNode   `json:"baseName"`
}

type UnaryOperation struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	Operator         string           `json:"operator"`
	Prefix           bool             `json:"prefix"`
	SubExpression    ASTNode          `json:"subExpression"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type Mapping struct {
	ASTNode
	KeyType          ASTNode          `json:"keyType"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
	ValueType        ASTNode          `json:"valueType"`
}

type StructDefinition struct {
	ASTNode
	CanonicalName string    `json:"canonicalName"`
	Members       []ASTNode `json:"members"`
	Name          string    `json:"name"`
	Scope         int       `json:"scope"`
	Visibility    string    `json:"visibility"`
}

type UsingForDirective struct {
	ASTNode
	LibraryName ASTNode `json:"libraryName"`
	TypeName    ASTNode `json:"typeName"`
}

type IndexAccess struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	BaseExpression   []ASTNode        `json:"baseExpression"`
	IndexExpression  []ASTNode        `json:"indexExpression"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type EnumValue struct {
	ASTNode
	Name string `json:"name"`
}

type EnumDefinition struct {
	ASTNode
	CanonicalName string      `json:"canonicalName"`
	Members       []EnumValue `json:"members"`
	Name          string      `json:"name"`
}

type ArraryTypeName struct {
	ASTNode
	BaseType         ASTNode          `json:"baseType"`
	Length           string           `json:"length"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type ForStatement struct {
	ASTNode
	Body                     Block           `json:"body"`
	Condition                BinaryOperation `json:"condition"`
	InitializationExpression ASTNode         `json:"initializationExpression"`
	LoopExpression           ASTNode         `json:"loopExpression"`
}

type Break struct {
	ASTNode
}

type TupleExpression struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	Components       []ASTNode        `json:"components"`
	IsConstant       bool             `json:"isConstant"`
	IsInlineArray    bool             `json:"isInlineArray"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

type NewExpression struct {
	ASTNode
	ArgumentTypes    []string         `json:"argumentTypes"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
	TypeName         ASTNode          `json:"typeName"`
}
