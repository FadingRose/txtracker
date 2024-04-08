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
		e.Documentation.Constructor(&data)
	}
	if data, ok := (*data)["members"].([]interface{}); ok {
		e.Members = make([]EnumValue, len(data))
		for i, v := range data {
			v := v.(map[string]interface{})
			e.Members[i].Constructor(&v)
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
		e.Parameters.Constructor(&data)
	}
}
