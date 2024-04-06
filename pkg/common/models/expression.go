package models

import (
	"fmt"
)

// Expression: Assignment | BinaryOperation | Conditional | ElementaryTypeNameExpression | FunctionCall | FunctionCallOptions | Identifier | IndexAccess | IndexRangeAccess | Literal | MemberAccess | NewExpression | TupleExpression | UnaryOperation
type Expression interface {
	DescribeExpression() string
	Constructor(*map[string]interface{})
}

// operator: "=" | "+=" | "-=" | "*=" | "/=" | "%=" | "|=" | "&=" | "^=" | ">>=" | "<<="
type Operator string

const (
	EqualOperator           Operator = "="
	AddEqualOperator        Operator = "+="
	SubEqualOperator        Operator = "-="
	MulEqualOperator        Operator = "*="
	DivEqualOperator        Operator = "/="
	ModEqualOperator        Operator = "%="
	OrEqualOperator         Operator = "|="
	AndEqualOperator        Operator = "&="
	XorEqualOperator        Operator = "^="
	RightShiftEqualOperator Operator = ">>="
	LeftShiftEqualOperator  Operator = "<<="
)

// ----------------------------------------------------------------------------
// Assignment node
type Assignment struct {
	Common
	ArgumentTypes    []string         `json:"argumentTypes"`
	IsConstant       bool             `json:"isConstant"`
	IsLValue         bool             `json:"isLValue"`
	IsPure           bool             `json:"isPure"`
	LValueRequested  bool             `json:"lValueRequested"`
	LeftHandSide     Expression       `json:"leftHandSide"`
	Operator         Operator         `json:"operator"`
	RightHandSide    Expression       `json:"rightHandSide"`
	TypeDescriptions TypeDescriptions `json:"typeDescriptions"`
}

func (a *Assignment) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    a.ArgumentTypes,
		"IsConstant":       a.IsConstant,
		"IsLValue":         a.IsLValue,
		"IsPure":           a.IsPure,
		"LValueRequested":  a.LValueRequested,
		"LeftHandSide":     a.LeftHandSide,
		"Operator":         a.Operator,
		"RightHandSide":    a.RightHandSide,
		"TypeDescriptions": a.TypeDescriptions,
	}
}

func (a *Assignment) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []string
		for _, v := range data.([]interface{}) {
			res = append(res, v.(string))
		}
		a.ArgumentTypes = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		a.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		a.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		a.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		a.LValueRequested = res
	}
	if data, ok := (*data)["LeftHandSide"]; ok {
		var res Expression
		res.Constructor(&data)
		a.LeftHandSide = res
	}
	if data, ok := (*data)["Operator"]; ok {
		var res Operator
		res = Operator(data.(string))
		a.Operator = res
	}
	if data, ok := (*data)["RightHandSide"]; ok {
		a.RightHandSide = data.(Expression)
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		a.TypeDescriptions = data.(TypeDescriptions)
	}

}

func (a *Assignment) DescribeExpression() string {
	return fmt.Sprintf("This is an assignment expression.")
}

// ----------------------------------------------------------------------------
// BinaryOperation node
type BinaryOperation struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	CommonType       TypeDescriptions   `json:"commonType"`
	Function         int                `json:"function"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	LeftExpression   Expression         `json:"leftExpression"`
	Operator         Operator           `json:"operator"`
	RightExpression  Expression         `json:"rightExpression"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (b *BinaryOperation) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    b.ArgumentTypes,
		"CommonType":       b.CommonType,
		"Function":         b.Function,
		"IsConstant":       b.IsConstant,
		"IsLValue":         b.IsLValue,
		"IsPure":           b.IsPure,
		"LValueRequested":  b.LValueRequested,
		"LeftExpression":   b.LeftExpression,
		"Operator":         b.Operator,
		"RightExpression":  b.RightExpression,
		"TypeDescriptions": b.TypeDescriptions,
	}
}

func (b *BinaryOperation) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		b.ArgumentTypes = res
	}
	if data, ok := (*data)["CommonType"]; ok {
		b.CommonType = data.(TypeDescriptions)
	}
	if data, ok := (*data)["Function"]; ok {
		var res int
		res = data.(int)
		b.Function = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		b.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		b.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		b.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		b.LValueRequested = res
	}
	if data, ok := (*data)["LeftExpression"]; ok {
		var res Expression
		res.Constructor(&data)
		b.LeftExpression = res
	}
	if data, ok := (*data)["Operator"]; ok {
		var res Operator
		res = Operator(data.(string))
		b.Operator = res
	}
	if data, ok := (*data)["RightExpression"]; ok {
		b.RightExpression = data.(Expression)
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		b.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (b *BinaryOperation) DescribeExpression() string {
	return fmt.Sprintf("This is a binary operation expression.")
}

// ----------------------------------------------------------------------------
// Conditional node
type Conditional struct {
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	Condition        Expression         `json:"condition"`
	FalseExpression  Expression         `json:"falseExpression"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TrueExpression   Expression         `json:"trueExpression"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (c *Conditional) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    c.ArgumentTypes,
		"Condition":        c.Condition,
		"FalseExpression":  c.FalseExpression,
		"IsConstant":       c.IsConstant,
		"IsLValue":         c.IsLValue,
		"IsPure":           c.IsPure,
		"LValueRequested":  c.LValueRequested,
		"TrueExpression":   c.TrueExpression,
		"TypeDescriptions": c.TypeDescriptions,
	}
}

func (c *Conditional) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		c.ArgumentTypes = res
	}
	if data, ok := (*data)["Condition"]; ok {
		var res Expression
		res.Constructor(&data)
		c.Condition = res
	}
	if data, ok := (*data)["FalseExpression"]; ok {
		c.FalseExpression = data.(Expression)
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		c.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		c.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		c.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		c.LValueRequested = res
	}
	if data, ok := (*data)["TrueExpression"]; ok {
		c.TrueExpression = data.(Expression)
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		c.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (c *Conditional) DescribeExpression() string {
	return fmt.Sprintf("This is a conditional expression.")
}
