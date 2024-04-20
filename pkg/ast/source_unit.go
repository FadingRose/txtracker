package ast

// SourceUnit's nodes are follows:
//
// (ContractDefinition | EnumDefinition | ErrorDefinition | FunctionDefinition | ImportDirective | PragmaDirective | StructDefinition | UserDefinedValueTypeDefinition | UsingForDirective | VariableDeclaration)[]
type SourceUnit struct {
	Common
	AbsolutePath         string          `json:"absolutePath"`
	ExperimentalSolidity bool            `json:"experimentalSolidity"` // boolean || null -> false
	ExportedSymbols      ExportedSymbols `json:"exportedSymbols"`
	License              string          `json:"license"` // string | null
}

func (s *SourceUnit) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"AbsolutePath":         s.AbsolutePath,
		"ExperimentalSolidity": s.ExperimentalSolidity,
		"ExportedSymbols":      s.ExportedSymbols,
		"License":              s.License,
	}
}

func (s *SourceUnit) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["absolutePath"].(string); ok {
		s.AbsolutePath = data
	}

	if data, ok := (*data)["experimentalSolidity"].(bool); ok {
		s.ExperimentalSolidity = data
	} else {
		s.ExperimentalSolidity = false
	}

	if data, ok := (*data)["exportedSymbols"].(map[string][]float64); ok {
		s.ExportedSymbols = make(ExportedSymbols)
		s.ExportedSymbols.Constructor(&data)
	}

	if data, ok := (*data)["license"].(string); ok {

		s.License = data
	} else {
		s.License = ""
	}
}

type ContractDefinition struct {
	Common
	Abstract                bool                    `json:"abstract"`
	BaseContracts           []InheritanceSpecifier  `json:"baseContracts"`
	CanonicaName            string                  `json:"canonicalName"` // string | null
	ContractDependencies    ContractDependencies    `json:"contractDependencies"`
	ContractKind            ContractKind            `json:"contractKind"`
	Documentation           StructuredDocumentation `json:"documentation"` // StructuredDocumentation | null
	FullyImplemented        bool                    `json:"fullyImplemented"`
	InternalFunctionIDs     InternalFunctionIDs     `json:"internalFunctionIDs"`
	LinearizedBaseContracts LinearizedBaseContracts `json:"linearizedBaseContracts"`
	Name                    string                  `json:"name"`
	Scope                   int                     `json:"scope"`
	UsedErrors              UsedErrors              `json:"usedErrors"`
	UsedEvents              UsedEvents              `json:"usedEvents"`
}

func (c *ContractDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Abstract":                c.Abstract,
		"BaseContracts":           c.BaseContracts,
		"CanonicaName":            c.CanonicaName,
		"ContractDependencies":    c.ContractDependencies,
		"ContractKind":            c.ContractKind,
		"Documentation":           c.Documentation,
		"FullyImplemented":        c.FullyImplemented,
		"InternalFunctionIDs":     c.InternalFunctionIDs,
		"LinearizedBaseContracts": c.LinearizedBaseContracts,
		"Name":                    c.Name,
		"Scope":                   c.Scope,
		"UsedErrors":              c.UsedErrors,
		"UsedEvents":              c.UsedEvents,
	}
}

func (c *ContractDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["abstract"].(bool); ok {
		c.Abstract = data
	}

	if data, ok := (*data)["baseContracts"].([]interface{}); ok {
		c.BaseContracts = make([]InheritanceSpecifier, len(data))
		for i, v := range data {
			v := v.(map[string]interface{})
			c.BaseContracts[i].Constructor(&v)
		}
	}

	if data, ok := (*data)["canonicalName"].(string); ok {
		c.CanonicaName = data
	}

	if data, ok := (*data)["contractDependencies"].([]float64); ok {
		c.ContractDependencies = make(ContractDependencies, len(data))
		c.ContractDependencies.Constructor(&data)
	}

	if data, ok := (*data)["contractKind"].(string); ok {
		c.ContractKind = ContractKind(data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		c.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		c.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["fullyImplemented"].(bool); ok {
		c.FullyImplemented = data
	}

	if data, ok := (*data)["internalFunctionIDs"].(map[string]float64); ok {
		c.InternalFunctionIDs = make(InternalFunctionIDs)
		c.InternalFunctionIDs.Constructor(&data)
	}

	if data, ok := (*data)["linearizedBaseContracts"].([]float64); ok {
		c.LinearizedBaseContracts = make(LinearizedBaseContracts, len(data))
		c.LinearizedBaseContracts.Constructor(&data)
	}

	if data, ok := (*data)["name"].(string); ok {
		c.Name = data
	}

	if data, ok := (*data)["scope"].(int); ok {
		c.Scope = data
	}

	if data, ok := (*data)["usedErrors"].([]float64); ok {
		c.UsedErrors = make(UsedErrors, len(data))
		c.UsedErrors.Constructor(&data)
	}

	if data, ok := (*data)["usedEvents"].([]float64); ok {
		c.UsedEvents = make(UsedEvents, len(data))
		c.UsedEvents.Constructor(&data)
	}
}

type EnumDefinition struct {
	CanonicaName  string                  `json:"canonicalName"`
	Documentation StructuredDocumentation `json:"documentation"`
	Members       []EnumValue             `json:"members"`
	Name          string                  `json:"name"`
	NameLocation  string                  `json:"nameLocation"`
}

func (e *EnumDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"CanonicaName":  e.CanonicaName,
		"Documentation": e.Documentation,
		"Members":       e.Members,
		"Name":          e.Name,
		"NameLocation":  e.NameLocation,
	}
}

func (e *EnumDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["canonicalName"].(string); ok {
		e.CanonicaName = data
	}
	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		e.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		e.Documentation.Constructor(&data)
	}
	if data, ok := (*data)["members"].([]interface{}); ok {
		//e.Members = make([]EnumValue, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			member := NodeFactory(v).ToEnumValue()
			member.Constructor(&v)
			e.Members = append(e.Members, *member)
		}
	}
	if data, ok := (*data)["name"].(string); ok {
		e.Name = data
	}
	if data, ok := (*data)["nameLocation"].(string); ok {
		e.NameLocation = data
	}
}

type ErrorDefinition struct {
	Documentation StructuredDocumentation `json:"documentation"` // StructuredDocumentation | null
	ErrorSelector string                  `json:"errorSelector"` // string | null
	Name          string                  `json:"name"`
	NameLocation  string                  `json:"nameLocation"`
	Parameters    ParameterList           `json:"parameters"`
}

func (e *ErrorDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Documentation": e.Documentation,
		"ErrorSelector": e.ErrorSelector,
		"Name":          e.Name,
		"NameLocation":  e.NameLocation,
		"Parameters":    e.Parameters,
	}
}

func (e *ErrorDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		e.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		e.Documentation.Constructor(&data)
	}
	if data, ok := (*data)["errorSelector"].(string); ok {
		e.ErrorSelector = data
	} else {
		e.ErrorSelector = ""

	}
	if data, ok := (*data)["name"].(string); ok {
		e.Name = data
	}
	if data, ok := (*data)["nameLocation"].(string); ok {
		e.NameLocation = data
	}
	if data, ok := (*data)["parameters"].(map[string]interface{}); ok {
		e.Parameters = *NodeFactory(data).ToParameterList()
		e.Parameters.Constructor(&data)
	}
}

type PragmaDirective struct {
	Common
	Literals Literals `json:"literals"`
}

func (p *PragmaDirective) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Literals": p.Literals,
	}
}

func (p *PragmaDirective) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["literals"].([]string); ok {
		p.Literals = make(Literals, len(data))
		p.Literals.Constructor(&data)
	}
}

type FunctionDefinition struct {
	Common
	Body             Block                   `json:"body"`             // Block | null
	Documentation    StructuredDocumentation `json:"documentation"`    // StructuredDocumentation | null
	FunctionSelector string                  `json:"functionSelector"` // string | null
	Implemented      bool                    `json:"implemented"`
	Kind             FunctionKind            `json:"kind"`
	Modifiers        []ModifierInvocation    `json:"modifiers"`
	Name             string                  `json:"name"`
	NameLocation     string                  `json:"nameLocation"` // string | null
	Overrides        OverrideSpecifier       `json:"overrides"`    // OverrideSpecifier | null
	Parameters       ParameterList           `json:"parameters"`
	ReturnParameters ParameterList           `json:"returnParameters"`
	Scope            int                     `json:"scope"`
	StateMutability  StateMutability         `json:"stateMutability"`
	Virtual          bool                    `json:"virtual"`
	Visibility       Visibility              `json:"visibility"`
}

func (f *FunctionDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Body":             f.Body,
		"Documentation":    f.Documentation,
		"FunctionSelector": f.FunctionSelector,
		"Implemented":      f.Implemented,
		"Kind":             f.Kind,
		"Modifiers":        f.Modifiers,
		"Name":             f.Name,
		"NameLocation":     f.NameLocation,
		"Overrides":        f.Overrides,
		"Parameters":       f.Parameters,
		"ReturnParameters": f.ReturnParameters,
		"Scope":            f.Scope,
		"StateMutability":  f.StateMutability,
		"Virtual":          f.Virtual,
		"Visibility":       f.Visibility,
	}
}

func (f *FunctionDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["body"].(map[string]interface{}); ok {
		f.Body = *NodeFactory(data).ToBlock()
		f.Body.Constructor(&data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		f.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		f.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["functionSelector"].(string); ok {
		f.FunctionSelector = data
	}

	if data, ok := (*data)["implemented"].(bool); ok {
		f.Implemented = data
	}

	if data, ok := (*data)["kind"].(string); ok {
		f.Kind = FunctionKind(data)
	}

	if data, ok := (*data)["modifiers"].([]interface{}); ok {
		//f.Modifiers = make([]ModifierInvocation, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			mi := NodeFactory(v).ToModifierInvocation()
			mi.Constructor(&v)
			f.Modifiers = append(f.Modifiers, *mi)
		}
	}

	if data, ok := (*data)["name"].(string); ok {
		f.Name = data
	}

	if data, ok := (*data)["nameLocation"].(string); ok {
		f.NameLocation = data
	}

	if data, ok := (*data)["overrides"].(map[string]interface{}); ok {
		f.Overrides = *NodeFactory(data).ToOverrideSpecifier()
		f.Overrides.Constructor(&data)
	}

	if data, ok := (*data)["parameters"].(map[string]interface{}); ok {
		f.Parameters = *NodeFactory(data).ToParameterList()
		f.Parameters.Constructor(&data)
	}

	if data, ok := (*data)["returnParameters"].(map[string]interface{}); ok {
		f.ReturnParameters = *NodeFactory(data).ToParameterList()
		f.ReturnParameters.Constructor(&data)
	}

	if data, ok := (*data)["scope"].(int); ok {
		f.Scope = data
	}

	if data, ok := (*data)["stateMutability"].(string); ok {
		f.StateMutability = StateMutability(data)
	}

	if data, ok := (*data)["virtual"].(bool); ok {
		f.Virtual = data
	}

	if data, ok := (*data)["visibility"].(string); ok {
		f.Visibility = Visibility(data)
	}
}

func (f *FunctionDefinition) IsPublic() bool {
	return f.Visibility == "public"
}

func (f *FunctionDefinition) IsExternal() bool {
	return f.Visibility == "external"
}

type ModifierInvocation struct {
	Common
	Arguments    []Expression `json:"arguments"` // Expression[] || null
	Kind         ModifierKind `json:"kind"`
	ModifierName ModifierName `json:"modifierName"` // Identifier | IdentifierPath
}

func (m *ModifierInvocation) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Arguments":    m.Arguments,
		"Kind":         m.Kind,
		"ModifierName": m.ModifierName,
	}
}

func (m *ModifierInvocation) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["arguments"].([]interface{}); ok {
		//m.Arguments = make([]Expression, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			//m.Arguments[i].ASTNode.Constructor(&v)
			expr := NodeFactory(v)
			expr.ASTNode.Constructor(&v)
		}
	}

	if data, ok := (*data)["kind"].(string); ok {
		m.Kind = ModifierKind(data)
	}

	if data, ok := (*data)["modifierName"].(map[string]interface{}); ok {
		m.ModifierName = NodeFactory(data)
		m.ModifierName.ASTNode.Constructor(&data)
	}
}

type FunctionCall struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	Arguments        []Expression       `json:"arguments"`
	Expression       Expression         `json:"expression"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	Kind             FunctionCallKind   `json:"kind"`
	LValueRequested  bool               `json:"lValueRequested"`
	NameLocations    []string           `json:"nameLocations"`
	Names            []string           `json:"names"`
	TryCall          bool               `json:"tryCall"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (f *FunctionCall) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"ArgumentTypes":    f.ArgumentTypes,
		"Arguments":        f.Arguments,
		"Expression":       f.Expression,
		"IsConstant":       f.IsConstant,
		"IsLValue":         f.IsLValue,
		"IsPure":           f.IsPure,
		"Kind":             f.Kind,
		"LValueRequested":  f.LValueRequested,
		"NameLocations":    f.NameLocations,
		"Names":            f.Names,
		"TryCall":          f.TryCall,
		"TypeDescriptions": f.TypeDescriptions,
	}
}

func (f *FunctionCall) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		//f.ArgumentTypes = make([]TypeDescriptions, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			td := NodeFactory(v).ToTypeDescriptions()
			td.Constructor(&v)
			f.ArgumentTypes = append(f.ArgumentTypes, *td)
		}
	}

	if data, ok := (*data)["arguments"].([]interface{}); ok {
		//f.Arguments = make([]Expression, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			expr := NodeFactory(v)
			expr.ASTNode.Constructor(&v)
			f.Arguments = append(f.Arguments, expr)
		}
	}

	if data, ok := (*data)["expression"].(map[string]interface{}); ok {
		f.Expression = NodeFactory(data)
		f.Expression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		f.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		f.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		f.IsPure = data
	}

	if data, ok := (*data)["kind"].(string); ok {
		f.Kind = FunctionCallKind(data)
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		f.LValueRequested = data
	}

	if data, ok := (*data)["nameLocations"].([]string); ok {
		f.NameLocations = data
	}

	if data, ok := (*data)["names"].([]string); ok {
		f.Names = data
	}

	if data, ok := (*data)["tryCall"].(bool); ok {
		f.TryCall = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		f.TypeDescriptions = TypeDescriptions{}
		f.TypeDescriptions.Constructor(&data)
	}
}
