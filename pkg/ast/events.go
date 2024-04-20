package ast

// ----------------------------------------------------------------------------
// Event nodes
// ----------------------------------------------------------------------------

type EventDefinition struct {
	Common
	Anonymous     bool                    `json:"anonymous"`
	Documentation StructuredDocumentation `json:"documentation"`
	EventSelector string                  `json:"eventSelector"`
	Name          string                  `json:"name"`
	NameLocation  string                  `json:"nameLocation"`
	Parameters    ParameterList           `json:"parameters"`
}

func (e *EventDefinition) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"Anonymous":     e.Anonymous,
		"Documentation": e.Documentation,
		"EventSelector": e.EventSelector,
		"Name":          e.Name,
		"NameLocation":  e.NameLocation,
		"Parameters":    e.Parameters,
	}
}

func (e *EventDefinition) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["anonymous"].(bool); ok {
		e.Anonymous = data
	}

	if data, ok := (*data)["documentation"].(map[string]interface{}); ok {
		e.Documentation = *NodeFactory(data).ToStructuredDocumentation()
		e.Documentation.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["eventSelector"].(string); ok {
		e.EventSelector = data
	}

	if data, ok := (*data)["name"].(string); ok {
		e.Name = data
	}

	if data, ok := (*data)["nameLocation"].(string); ok {
		e.NameLocation = data
	}

	if data, ok := (*data)["parameters"].(map[string]interface{}); ok {
		e.Parameters = *NodeFactory(data).ToParameterList()
		e.Parameters.ASTNode.Constructor(&data)
	}
}
