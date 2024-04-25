package ast

// TypeName: ArrayTypeName | ElementaryTypeName | FunctionTypeName | Mapping | UserDefinedTypeName
type TypeName = *Common

type ElementaryTypeName struct {
	Common
	Name             string           `json:"name"`
	StateMutability  StateMutability  `json:"stateMutability"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

func (e *ElementaryTypeName) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Name":             e.Name,
		"StateMutability":  e.StateMutability,
		"TypeDescriptions": e.TypeDescriptions,
	}
}

func (e *ElementaryTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["name"].(string); ok {
		e.Name = data
	}

	if data, ok := (*data)["stateMutability"].(string); ok {
		e.StateMutability = StateMutability(data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		e.TypeDescriptions.Constructor(&data)
	}
}

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

func (m *Mapping) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
		m.KeyName = data
	}

	if data, ok := (*data)["keyNameLocation"].(string); ok {
		m.KeyNameLocation = data
	}

	if data, ok := (*data)["keyType"].(map[string]interface{}); ok {
		m.KeyType = NodeFactory(data)
		m.KeyType.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		m.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*data)["valueName"].(string); ok {
		m.ValueName = data
	}

	if data, ok := (*data)["valueNameLocation"].(string); ok {
		m.ValueNameLocation = data
	}

	if data, ok := (*data)["valueType"].(map[string]interface{}); ok {
		m.ValueType = NodeFactory(data)
		m.ValueType.ASTNode.Constructor(&data)
	}
}

type ArrayTypeName struct {
	BaseType         TypeName         `json:"baseType"`
	Length           Expression       `json:"length"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

func (a *ArrayTypeName) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"BaseType":         a.BaseType,
		"Length":           a.Length,
		"TypeDescriptions": a.TypeDescriptions,
	}
}

func (a *ArrayTypeName) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["baseType"].(map[string]interface{}); ok {
		a.BaseType = NodeFactory(data)
		a.BaseType.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["length"].(map[string]interface{}); ok {
		a.Length = NodeFactory(data)
		a.Length.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		a.TypeDescriptions.Constructor(&data)
	}
}
