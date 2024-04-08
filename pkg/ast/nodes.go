package ast

// ----------------------------------------------------------------------------
// Inline nodes of the top-level nodes
// ----------------------------------------------------------------------------

type InheritanceSpecifierBaseNameInterface interface {
	DescribeInheritanceSpecifierBaseName() string
	Constructor(data *map[string]interface{})
}

type InheritanceSpecifier struct {
	Common
	Arguments []Expression                          `json:"arguments"` // Expression[] | null
	BaseName  InheritanceSpecifierBaseNameInterface // UserDefinedTypeName | IdentifierPath
}

func (i *InheritanceSpecifier) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Arguments": i.Arguments,
		"BaseName":  i.BaseName,
	}
}

func (i *InheritanceSpecifier) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["arguments"].([]interface{}); ok {
		i.Arguments = make([]Expression, len(data))
		for cnt, v := range data {
			v := v.(map[string]interface{})
			i.Arguments[cnt].Constructor(&v)
		}
	}

	if data, ok := (*data)["baseName"].(map[string]interface{}); ok {
		i.BaseName.Constructor(&data)
	}
}

type StructuredDocumentation struct {
	Common
	Text string `json:"text"`
}

func (s *StructuredDocumentation) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Text": s.Text,
	}
}

func (s *StructuredDocumentation) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["text"].(string); ok {
		s.Text = data
	}
}

type UserDefinedTypeName struct {
	Common
	ContractScope        string           `json:"contractScope"`        // string | null
	Name                 string           `json:"name"`                 // string | null
	PathNode             IdentifierPath   `json:"pathNode"`             // IdentifierPath | null
	ReferenceDeclaration int              `json:"referenceDeclaration"` // int | null
	TypeDescriptions     TypeDescriptions `json:"typeDescriptions"`     // TypeDescriptions | null
}

func (u *UserDefinedTypeName) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"ContractScope":        u.ContractScope,
		"Name":                 u.Name,
		"PathNode":             u.PathNode,
		"ReferenceDeclaration": u.ReferenceDeclaration,
		"TypeDescriptions":     u.TypeDescriptions,
	}
}

func (u *UserDefinedTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["contractScope"].(string); ok {
		u.ContractScope = data
	}

	if data, ok := (*data)["name"].(string); ok {
		u.Name = data
	}

	if data, ok := (*data)["pathNode"].(map[string]interface{}); ok {
		u.PathNode.Constructor(&data)
	}

	if data, ok := (*data)["referenceDeclaration"].(int); ok {
		u.ReferenceDeclaration = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		u.TypeDescriptions.Constructor(&data)
	}
}

func (u *UserDefinedTypeName) DescribeOverrideSpecifier() string {
	return "UserDefinedTypeName"
}

type IdentifierPath struct {
	Common
	Name                  string        `json:"name"`
	NameLocations         NameLocations `json:"nameLocation"` // string[] | null
	ReferencedDeclaration int           `json:"referencedDeclaration"`
}

func (i *IdentifierPath) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Name":                  i.Name,
		"NameLocations":         i.NameLocations,
		"ReferencedDeclaration": i.ReferencedDeclaration,
	}
}

func (i *IdentifierPath) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["name"].(string); ok {
		i.Name = data
	}

	if data, ok := (*data)["nameLocation"].([]string); ok {
		i.NameLocations = make(NameLocations, len(data))
		i.NameLocations.Constructor(&data)
	}

	if data, ok := (*data)["referencedDeclaration"].(int); ok {
		i.ReferencedDeclaration = data
	}
}

func (i *IdentifierPath) DescribeOverrideSpecifier() string {
	return "IdentifierPath"
}

type TypeDescriptions struct {
	Common
	TypeIdentifier string `json:"typeIdentifier"` // string | null
	TypeString     string `json:"typeString"`     // string | null
}

func (t *TypeDescriptions) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"TypeIdentifier": t.TypeIdentifier,
		"TypeString":     t.TypeString,
	}
}

func (t *TypeDescriptions) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["typeIdentifier"].(string); ok {
		t.TypeIdentifier = data
	}

	if data, ok := (*data)["typeString"].(string); ok {
		t.TypeString = data
	}
}

type EnumValue struct {
	Common
	Name         string `json:"name"`
	NameLocation string `json:"nameLocation"`
}

func (e *EnumValue) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Name":         e.Name,
		"NameLocation": e.NameLocation,
	}
}

func (e *EnumValue) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["name"].(string); ok {
		e.Name = data
	}
	if data, ok := (*data)["nameLocation"].(string); ok {
		e.NameLocation = data
	}
}

type ParameterList struct {
	Common
	Parameters []VariableDeclaration `json:"parameters"`
}

func (p *ParameterList) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Parameters": p.Parameters,
	}
}

func (p *ParameterList) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["parameters"].([]interface{}); ok {
		p.Parameters = make([]VariableDeclaration, len(data))
		for cnt, v := range data {
			v := v.(map[string]interface{})
			p.Parameters[cnt].Constructor(&v)
		}
	}
}

func (p *ParameterList) New() *ParameterList {
	return &ParameterList{
		Parameters: make([]VariableDeclaration, 0),
	}
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

func (v *VariableDeclaration) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["baseFunctions"].([]float64); ok {
		v.BaseFunctions = make(BaseFunctions, len(data))
		v.BaseFunctions.Constructor(&data)
	}

	if data, ok := (*data)["constant"].(bool); ok {
		v.Constant = data
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		v.Documentation.Constructor(&data)
	}

	if data, ok := (*data)["functionSelector"].(string); ok {
		v.FunctionSelector = data
	}

	if data, ok := (*data)["indexed"].(bool); ok {
		v.Indexed = data
	}

	if data, ok := (*data)["mutability"].(string); ok {
		v.Mutability = Mutability(data)
	}

	if data, ok := (*data)["name"].(string); ok {
		v.Name = data
	}

	if data, ok := (*data)["nameLocation"].(string); ok {
		v.NameLocation = data
	}

	if data, ok := (*data)["overrides"].(map[string]interface{}); ok {
		v.Overrides.Constructor(&data)
	}

	if data, ok := (*data)["scope"].(int); ok {
		v.Scope = data
	}

	if data, ok := (*data)["stateVariable"].(bool); ok {
		v.StateVariable = data
	}

	if data, ok := (*data)["storageLocation"].(string); ok {
		v.StorageLocation = StorageLocation(data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		v.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*data)["typeName"].(map[string]interface{}); ok {
		v.TypeName.Constructor(&data)
	}

	if data, ok := (*data)["value"].(map[string]interface{}); ok {
		v.Value.Constructor(&data)
	}

	if data, ok := (*data)["visibility"].(string); ok {
		v.Visibility = Visibility(data)
	}
}

type OverideSpecifierInterface interface {
	DescribeOverrideSpecifier() string
	Constructor(data *map[string]interface{})
}

type OverrideSpecifier struct {
	Common
	Overrides []OverideSpecifierInterface `json:"overrides"`
}

func (o *OverrideSpecifier) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Overrides": o.Overrides,
	}
}

func (o *OverrideSpecifier) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["overrides"].([]interface{}); ok {
		o.Overrides = make([]OverideSpecifierInterface, len(data))
		for cnt, v := range data {
			v := v.(map[string]interface{})
			o.Overrides[cnt].Constructor(&v)
		}
	}
}
