package ast

// ----------------------------------------------------------------------------
// Inline nodes of the top-level nodes
// ----------------------------------------------------------------------------

type InheritanceSpecifier struct {
	Common
	Arguments []Expression `json:"arguments"` // Expression[] | null
	BaseName  *Common      // UserDefinedTypeName | IdentifierPath
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
			i.Arguments[cnt].ASTNode.Constructor(&v)
		}
	}

	if data, ok := (*data)["baseName"].(map[string]interface{}); ok {
		i.BaseName.ASTNode.Constructor(&data)
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

type Overrides *Common //UserDefinedTypeName IdentifierPath

type OverrideSpecifier struct {
	Common
	Overrides []Overrides `json:"overrides"`
}

func (o *OverrideSpecifier) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Overrides": o.Overrides,
	}
}

func (o *OverrideSpecifier) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["overrides"].([]interface{}); ok {
		//o.Overrides = make([]OverrideSpecifier, len(data))
		for _, v := range data {
			v := v.(map[string]interface{})
			override := NodeFactory(v)
			override.ASTNode.Constructor(&v)
			o.Overrides = append(o.Overrides, override)
		}
	}
}
