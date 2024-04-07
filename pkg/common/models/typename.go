package models

// `TypeName`: ArrayTypeName | ElementaryTypeName | FunctionTypeName | Mapping | UserDefinedTypeName
type TypeName interface {
	DescribeTypeName() string
	Constructor(*map[string]interface{})
}

type StateMutability string

const (
	Pure       StateMutability = "pure"
	View       StateMutability = "view"
	Nonpayable StateMutability = "nonpayable"
	Payable    StateMutability = "payable"
)

// `Visibility`: "external" | "internal" | "public" | "private"
type Visibility string

const (
	External Visibility = "external"
	Internal Visibility = "internal"
	Public   Visibility = "public"
	Private  Visibility = "private"
)

// ----------------------------------------------------------------------------
// ArrayTypeName
type ArrayTypeName struct {
	Common
	BaseType         TypeName         `json:"baseType"`
	Length           Expression       `json:"length"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

func (a *ArrayTypeName) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"BaseType":         a.BaseType,
		"Length":           a.Length,
		"TypeDescriptions": a.TypeDescriptions,
	}
}

func (a *ArrayTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["baseType"].(map[string]interface{}); ok {
		var res TypeName
		res.Constructor(&data)
		a.BaseType = res
	}
	if data, ok := (*data)["length"].(map[string]interface{}); ok {
		var res Expression
		res.Constructor(&data)
		a.Length = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		a.TypeDescriptions = res
	}
}

func (a *ArrayTypeName) DescribeTypeName() string {
	return "ArrayTypeName"
}

// ----------------------------------------------------------------------------
// ElementaryTypeName
type ElementaryTypeName struct {
	Common
	Name             string           `json:"name"`
	StateMutability  StateMutability  `json:"stateMutability"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

func (e *ElementaryTypeName) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Name":             e.Name,
		"StateMutability":  e.StateMutability,
		"TypeDescriptions": e.TypeDescriptions,
	}
}

func (e *ElementaryTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["name"].(string); ok {
		var res string
		res = data
		e.Name = res
	}
	if data, ok := (*data)["stateMutability"].(string); ok {
		var res StateMutability
		res = StateMutability(data)
		e.StateMutability = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		e.TypeDescriptions = res
	}
}

func (e *ElementaryTypeName) DescribeTypeName() string {
	return "ElementaryTypeName"
}

// ----------------------------------------------------------------------------
// FunctionTypeName
type FunctionTypeName struct {
	Common
	ParameterTypes   ParameterList    `json:"parameterTypes"`
	ReturnTypes      ParameterList    `json:"returnTypes"`
	StateMutability  StateMutability  `json:"stateMutability"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
	Visibility       Visibility       `json:"visibility"`
}

func (f *FunctionTypeName) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ParameterTypes":   f.ParameterTypes,
		"ReturnTypes":      f.ReturnTypes,
		"StateMutability":  f.StateMutability,
		"TypeDescriptions": f.TypeDescriptions,
		"Visibility":       f.Visibility,
	}
}

func (f *FunctionTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["parameterTypes"]; ok {
		var res ParameterList
		data := data.(map[string]interface{})
		res.Constructor(&data)
		f.ParameterTypes = res
	}
	if data, ok := (*data)["returnTypes"]; ok {
		var res ParameterList
		data := data.(map[string]interface{})
		res.Constructor(&data)
		f.ReturnTypes = res
	}
	if data, ok := (*data)["stateMutability"].(string); ok {
		var res StateMutability
		res = StateMutability(data)
		f.StateMutability = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		f.TypeDescriptions = res
	}
	if data, ok := (*data)["visibility"].(string); ok {
		var res Visibility
		res = Visibility(data)
		f.Visibility = res
	}
}

func (f *FunctionTypeName) DescribeTypeName() string {
	return "FunctionTypeName"
}

// ----------------------------------------------------------------------------
// Mapping
type Mapping struct {
	Common
	KeyName           string           `json:"keyName"`
	KeyNameLocation   string           `json:"keyNameLocation"`
	KeyType           TypeName         `json:"keyType"`
	TypeDescriptions  TypeDescriptions `json:"typeDescriptions"`
	ValueName         string           `json:"valueName"`
	ValueNameLocation string           `json:"valueNameLocation"`
	ValueType         TypeName         `json:"valueType"`
}

func (m *Mapping) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"KeyName":           m.KeyName,
		"KeyNameLocation":   m.KeyNameLocation,
		"KeyType":           m.KeyType,
		"TypeDescriptions":  m.TypeDescriptions,
		"ValueName":         m.ValueName,
		"ValueNameLocation": m.ValueNameLocation,
		"ValueType":         m.ValueType,
	}
}

func (m *Mapping) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["keyName"].(string); ok {
		var res string
		res = data
		m.KeyName = res
	}
	if data, ok := (*data)["keyNameLocation"].(string); ok {
		var res string
		res = data
		m.KeyNameLocation = res
	}
	if data, ok := (*data)["keyType"].(map[string]interface{}); ok {
		var res TypeName
		res.Constructor(&data)
		m.KeyType = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		m.TypeDescriptions = res
	}
	if data, ok := (*data)["valueName"].(string); ok {
		var res string
		res = data
		m.ValueName = res
	}
	if data, ok := (*data)["valueNameLocation"].(string); ok {
		var res string
		res = data
		m.ValueNameLocation = res
	}
	if data, ok := (*data)["valueType"].(map[string]interface{}); ok {
		var res TypeName
		res.Constructor(&data)
		m.ValueType = res
	}
}

func (m *Mapping) DescribeTypeName() string {
	return "Mapping"
}

// ----------------------------------------------------------------------------
// UserDefinedTypeName
type UserDefinedTypeName struct {
	Common
	ContractScope        string           `json:"contractScope"`
	Name                 string           `json:"name"`
	PathNode             IdentifierPath   `json:"pathNode"`
	ReferenceDeclaration int              `json:"referenceDeclaration"`
	TypeDescriptions     TypeDescriptions `json:"typeDescriptions"`
}

func (u *UserDefinedTypeName) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ContractScope":        u.ContractScope,
		"Name":                 u.Name,
		"PathNode":             u.PathNode,
		"ReferenceDeclaration": u.ReferenceDeclaration,
		"TypeDescriptions":     u.TypeDescriptions,
	}
}

func (u *UserDefinedTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["contractScope"].(string); ok {
		var res string
		res = data
		u.ContractScope = res
	}
	if data, ok := (*data)["name"].(string); ok {
		var res string
		res = data
		u.Name = res
	}
	if data, ok := (*data)["pathNode"].(map[string]interface{}); ok {
		var res IdentifierPath
		res.Constructor(&data)
		u.PathNode = res
	}
	if data, ok := (*data)["referenceDeclaration"].(int); ok {
		var res int
		res = data
		u.ReferenceDeclaration = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		u.TypeDescriptions = res
	}
}

func (u *UserDefinedTypeName) DescribeTypeName() string {
	return "UserDefinedTypeName"
}
