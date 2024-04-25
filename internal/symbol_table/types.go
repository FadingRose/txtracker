package symboltable

type NodeType string

const (
	SourceUnit          NodeType = "SourceUnit"
	ContractDefinition  NodeType = "ContractDefinition"
	VariableDeclaration NodeType = "VariableDeclaration"
	FunctionDefinition  NodeType = "FunctionDefinition"
	ModifierDefinition  NodeType = "ModifierDefinition"
)
