package ast

type Statement *Common

type Block struct {
	Common
	Statements []*Common `json:"statements"`
}

func (b *Block) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Statements": b.Statements,
	}
}

func (b *Block) Constructor(data *map[string]interface{}) {

	if data, ok := (*data)["statements"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			sm := NodeFactory(v)
			sm.ASTNode.Constructor(&v)
			b.Statements = append(b.Statements, sm)
		}
	}

}
func (b *Block) DescribeStatement() string {
	return "Block"
}

type IfStatement struct {
	Common
	Condition     Expression              `json:"condition"` // Expression | null
	Documentation StructuredDocumentation `json:"documentation"`
	FalseBody     Statement               `json:"falseBody"` // Statement | null
	TrueBody      Statement               `json:"trueBody"`  // Statement | null
}

func (i *IfStatement) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Condition":     i.Condition,
		"Documentation": i.Documentation,
		"FalseBody":     i.FalseBody,
		"TrueBody":      i.TrueBody,
	}
}

func (i *IfStatement) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["condition"].(map[string]interface{}); ok {
		i.Condition = NodeFactory(data)
		i.Condition.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		i.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["falseBody"].(map[string]interface{}); ok {
		i.FalseBody = NodeFactory(data)
		i.FalseBody.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["trueBody"].(map[string]interface{}); ok {
		i.TrueBody = NodeFactory(data)
		i.TrueBody.ASTNode.Constructor(&data)
	}
}
func (i *IfStatement) DescribeStatement() string {
	return "IfStatement"
}

type Return struct {
	Common
	Documentation            StructuredDocumentation `json:"documentation"` // StructuredDocumentation | null
	Expression               Expression              `json:"expression"`    // Expression | null
	FunctionReturnParameters int                     `json:"functionReturnParameters"`
}

func (r *Return) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Documentation":            r.Documentation,
		"Expression":               r.Expression,
		"FunctionReturnParameters": r.FunctionReturnParameters,
	}
}

func (r *Return) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		r.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["expression"].(map[string]interface{}); ok {
		r.Expression = NodeFactory(data)
		r.Expression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["functionReturnParameters"].(int); ok {
		r.FunctionReturnParameters = data
	}
}

func (r *Return) DescribeStatement() string {
	return "Return"
}

type VariableDeclarationStatement struct {
	Common
	Assignments   []int                   `json:"assignments"`  // int[] | null
	Declarations  []*VariableDeclaration  `json:"declarations"` // VariableDeclaration
	Documentation StructuredDocumentation `json:"documentation"`
	InitialValue  Expression              `json:"initialValue"` // Expression | null
}

func (v *VariableDeclarationStatement) GetDeclarations() ([]string, []bool) {
	var declarations []string
	var statevariable []bool
	for _, vd := range v.Declarations {
		declaration, state := vd.GetDeclaration()
		declarations = append(declarations, declaration)
		statevariable = append(statevariable, state)
	}
	return declarations, statevariable
}

func (v *VariableDeclarationStatement) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Assignments":   v.Assignments,
		"Declarations":  v.Declarations,
		"Documentation": v.Documentation,
		"InitialValue":  v.InitialValue,
	}
}

func (v *VariableDeclarationStatement) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["assignments"].([]interface{}); ok {
		for _, dt := range data {
			dt := int(dt.(float64))
			v.Assignments = append(v.Assignments, dt)
		}
	}

	if data, ok := (*data)["declarations"].([]interface{}); ok {
		v.Declarations = make([]*VariableDeclaration, 0)
		for _, dt := range data {
			dt := dt.(map[string]interface{})
			vd := &VariableDeclaration{}
			vd.Constructor(&dt)
			v.Declarations = append(v.Declarations, vd)
		}
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		v.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["initialValue"].(map[string]interface{}); ok {
		v.InitialValue = NodeFactory(data)
		v.InitialValue.ASTNode.Constructor(&data)
	}
}

func (v *VariableDeclarationStatement) DescribeStatement() string {
	return "VariableDeclarationStatement"
}

type VariableDeclaration struct {
	Common
	BaseFunctions    BaseFunctions           `json:"baseFunctions"`    // int[] | null
	Constant         bool                    `json:"constant"`         // boolean
	Documentation    StructuredDocumentation `json:"documentation"`    // StructuredDocumentation | null
	FunctionSelector string                  `json:"functionSelector"` // string | null
	Indexed          bool                    `json:"indexed"`          // boolean
	Mutability       Mutability              `json:"mutability"`       // string
	Name             string                  `json:"name"`             // string
	NameLocation     string                  `json:"nameLocation"`     // string | null
	Overrides        OverrideSpecifier       `json:"overrides"`        // OverrideSpecifier | null
	Scope            int                     `json:"scope"`            // int
	StateVariable    bool                    `json:"stateVariable"`    // boolean
	StorageLocation  StorageLocation         `json:"storageLocation"`  // string
	TypeDescriptions TypeDescriptions        `json:"typeDescriptions"` // TypeDescriptions
	TypeName         TypeName                `json:"typeName"`         // TypeName
	Value            Expression              `json:"value"`            // Expression | null
	Visibility       Visibility              `json:"visibility"`       // string
}

func (v *VariableDeclaration) GetDeclaration() (string, bool) {
	return v.Name, v.StateVariable
}

func (v *VariableDeclaration) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"BaseFunctions":    v.BaseFunctions,
		"Constant":         v.Constant,
		"Documentation":    v.Documentation,
		"FunctionSelector": v.FunctionSelector,
		"Indexed":          v.Indexed,
		"Mutability":       v.Mutability,
		"Name":             v.Name,
		"NameLocation":     v.NameLocation,
		"Overrides":        v.Overrides,
		"Scope":            v.Scope,
		"StateVariable":    v.StateVariable,
		"StorageLocation":  v.StorageLocation,
		"TypeDescriptions": v.TypeDescriptions,
		"TypeName":         v.TypeName,
		"Value":            v.Value,
		"Visibility":       v.Visibility,
	}
}

func (v *VariableDeclaration) Constructor(_data *map[string]interface{}) {
	if data, ok := (*_data)["baseFunctions"].([]float64); ok {
		v.BaseFunctions = make(BaseFunctions, len(data))
		v.BaseFunctions.Constructor(&data)
	}

	if data, ok := (*_data)["constant"].(bool); ok {
		v.Constant = data
	}

	if data, ok := (*_data)["documentation"].(map[string]interface{}); ok {
		v.Documentation.Constructor(&data)
	}

	if data, ok := (*_data)["functionSelector"].(string); ok {
		v.FunctionSelector = data
	}

	if data, ok := (*_data)["indexed"].(bool); ok {
		v.Indexed = data
	}

	if data, ok := (*_data)["mutability"].(string); ok {
		v.Mutability = Mutability(data)
	}

	if data, ok := (*_data)["name"].(string); ok {
		v.Name = data
	}

	if data, ok := (*_data)["nameLocation"].(string); ok {
		v.NameLocation = data
	}

	if data, ok := (*_data)["overrides"].(map[string]interface{}); ok {
		v.Overrides.Constructor(&data)
	}

	if data, ok := (*_data)["scope"].(float64); ok {
		v.Scope = (int(data))
	}

	if data, ok := (*_data)["stateVariable"].(bool); ok {
		v.StateVariable = data
	}

	if data, ok := (*_data)["storageLocation"].(string); ok {
		v.StorageLocation = StorageLocation(data)
	}

	if data, ok := (*_data)["typeDescriptions"].(map[string]interface{}); ok {

		v.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*_data)["typeName"].(map[string]interface{}); ok {
		v.TypeName = NodeFactory(data)
		v.TypeName.ASTNode.Constructor(&data)
	}

	if data, ok := (*_data)["value"].(map[string]interface{}); ok {
		v.Value = NodeFactory(data)
		v.Value.ASTNode.Constructor(&data)
	}

	if data, ok := (*_data)["visibility"].(string); ok {
		v.Visibility = Visibility(data)
	}
}

type ExpressionStatement struct {
	Common
	Expression    Expression              `json:"expression"` // Expression
	Documentation StructuredDocumentation `json:"documentation"`
}

func (e *ExpressionStatement) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Expression":    e.Expression,
		"Documentation": e.Documentation,
	}
}

func (e *ExpressionStatement) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["expression"].(map[string]interface{}); ok {
		e.Expression = NodeFactory(data)
		e.Expression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		e.Documentation.Constructor(&data)
	}
}

type PlaceholderStatement struct {
	Common
	Documentation StructuredDocumentation `json:"documentation"`
}

func (p *PlaceholderStatement) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Documentation": p.Documentation,
	}
}

func (p *PlaceholderStatement) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		p.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		p.Documentation.Constructor(&data)
	}
}

type ForStatement struct {
	Common
	Body                     Statement               `json:"body"`      // Statement
	Condition                Expression              `json:"condition"` // Expression | null
	Documentation            StructuredDocumentation `json:"documentation"`
	InitializationExpression Expression              `json:"initializationExpression"` // Expression | null | VariableDeclarationStatement
	IsSimpleCounterLoop      bool                    `json:"isSimpleCounterLoop"`      // boolean
	LoopExpression           Expression              `json:"loopExpression"`           // Expression | null
}

func (f *ForStatement) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Body":                     f.Body,
		"Condition":                f.Condition,
		"Documentation":            f.Documentation,
		"InitializationExpression": f.InitializationExpression,
		"IsSimpleCounterLoop":      f.IsSimpleCounterLoop,
		"LoopExpression":           f.LoopExpression,
	}
}

func (f *ForStatement) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["body"].(map[string]interface{}); ok {
		f.Body = NodeFactory(data)
		f.Body.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["condition"].(map[string]interface{}); ok {
		f.Condition = NodeFactory(data)
		f.Condition.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		f.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["initializationExpression"].(map[string]interface{}); ok {
		f.InitializationExpression = NodeFactory(data)
		f.InitializationExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["isSimpleCounterLoop"].(bool); ok {
		f.IsSimpleCounterLoop = data
	}

	if data, ok := (*data)["loopExpression"].(map[string]interface{}); ok {
		f.LoopExpression = NodeFactory(data)
		f.LoopExpression.ASTNode.Constructor(&data)
	}
}

type UnaryOperation struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`    // TypeDescriptions
	Function         int                `json:"function"`         // int
	IsConstant       bool               `json:"isConstant"`       // boolean
	IsLValue         bool               `json:"isLValue"`         // boolean
	IsPure           bool               `json:"isPure"`           // boolean
	LValueRequested  bool               `json:"lValueRequested"`  // boolean
	Operator         UnaryOperator      `json:"operator"`         // string
	Prefix           bool               `json:"prefix"`           // boolean
	SubExpression    Expression         `json:"subExpression"`    // Expression
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions
}

func (u *UnaryOperation) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"ArgumentTypes":    u.ArgumentTypes,
		"Function":         u.Function,
		"IsConstant":       u.IsConstant,
		"IsLValue":         u.IsLValue,
		"IsPure":           u.IsPure,
		"LValueRequested":  u.LValueRequested,
		"Operator":         u.Operator,
		"Prefix":           u.Prefix,
		"SubExpression":    u.SubExpression,
		"TypeDescriptions": u.TypeDescriptions,
	}
}

func (u *UnaryOperation) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, dt := range data {
			dt := dt.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&dt)
			u.ArgumentTypes = append(u.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["function"].(float64); ok {
		u.Function = int(data)
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		u.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		u.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		u.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		u.LValueRequested = data
	}

	if data, ok := (*data)["operator"].(string); ok {
		u.Operator = UnaryOperator(data)
	}

	if data, ok := (*data)["prefix"].(bool); ok {
		u.Prefix = data
	}

	if data, ok := (*data)["subExpression"].(map[string]interface{}); ok {
		u.SubExpression = NodeFactory(data)
		u.SubExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		u.TypeDescriptions.Constructor(&data)
	}
}

type Break struct {
	Common
	Documentation StructuredDocumentation `json:"documentation"`
}

func (b *Break) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Documentation": b.Documentation,
	}
}

func (b *Break) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		b.Documentation.Constructor(&data)
	}
}
