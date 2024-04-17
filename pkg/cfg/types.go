package cfg

type StatementType int

const (
	Break StatementType = iota
	Continue
	DoWhile
	For
	Emit
	Expression
	If
	InlineAssembly
	Placeholder
	Return
	Revert
	Try
	UncheckedBlock
	VariableDeclaration
	While
)

type CFG struct {
	EntryPoints []*Function `json:"entryPoints"`
	Blocks      []*Block    `json:"blocks"`
	Edges       []*Edge     `json:"edges"`
}

type Function struct {
	Name      string `json:"name"`
	Successor *Block `json:"successor"`
}

type Block struct {
	ID                int         `json:"id"`
	Statements        []Statement `json:"statements"`
	SuccessorsEdges   []*Edge     `json:"successors"`
	PredecessorsEdges []*Edge     `json:"predecessors"`
}

type Edge struct {
	Source      *Block   `json:"source"`
	Destination *Block   `json:"destination"`
	Type        EdgeType `json:"type"`
}

type EdgeType int

const (
	ConditionalTrue EdgeType = iota
	ConditionalFalse
	Unconditional
)

type Statement struct {
	Type       StatementType
	Content    string
	Attributes map[string]interface{}
}
