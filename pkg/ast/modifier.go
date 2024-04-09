package ast

type ModifierName *Common

type ModifierDefinition struct {
	Common
	BaseModifiers []int                   `json:"baseModifiers"` // int[] | null
	Body          Block                   `json:"body"`          // Block
	Documentation StructuredDocumentation `json:"documentation"`
	Name          string                  `json:"name"`
	NameLocation  string                  `json:"nameLocation"`
	Overrides     OverrideSpecifier       `json:"overrides"` // OverideSpecifier | null
	Parameters    ParameterList           `json:"parameters"`
	Virtual       bool                    `json:"virtual"`
	Visibility    Visibility              `json:"visibility"`
}

func (m *ModifierDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"BaseModifiers": m.BaseModifiers,
		"Body":          m.Body,
		"Documentation": m.Documentation,
		"Name":          m.Name,
		"NameLocation":  m.NameLocation,
		"Overrides":     m.Overrides,
		"Parameters":    m.Parameters,
		"Virtual":       m.Virtual,
		"Visibility":    m.Visibility,
	}
}

func (m *ModifierDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["baseModifiers"].([]int); ok {
		m.BaseModifiers = data
	}

	if data, ok := (*data)["body"].(map[string]interface{}); ok {
		m.Body = *NodeFactory(data).ToBlock()
		m.Body.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		m.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		m.Documentation.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["name"].(string); ok {
		m.Name = data
	}

	if data, ok := (*data)["nameLocation"].(string); ok {
		m.NameLocation = data
	}

	if data, ok := (*data)["overrides"].(map[string]interface{}); ok {
		m.Overrides = *NodeFactory(data).ToOverrideSpecifier()
		m.Overrides.Constructor(&data)
	}

	if data, ok := (*data)["parameters"].(map[string]interface{}); ok {
		m.Parameters = *NodeFactory(data).ToParameterList()
		m.Parameters.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["virtual"].(bool); ok {
		m.Virtual = data
	}

	if data, ok := (*data)["visibility"].(string); ok {
		m.Visibility = Visibility(data)
	}
}
