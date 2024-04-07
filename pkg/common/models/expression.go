package models

import (
	"fmt"
)

// Expression: Assignment | BinaryOperation | Conditional | ElementaryTypeNameExpression | FunctionCall | FunctionCallOptions | Identifier | IndexAccess | IndexRangeAccess | Literal | MemberAccess | NewExpression | TupleExpression | UnaryOperation
type Expression interface {
	DescribeExpression() string
	Constructor(*map[string]interface{})
}

// ----------------------------------------------------------------------------
// Expression Factory
func ExpressionFactory(data *map[string]interface{}) Expression {
	common := &Common{
		ID:       (int)((*data)["id"].(float64)),
		NodeType: (*data)["nodeType"].(string),
		Src:      (*data)["src"].(string),
	}
	if data, ok := (*data)["nodeType"]; ok {
		switch data {
		case "Assignment":
			return &Assignment{
				Common: *common,
			}
		case "BinaryOperation":
			return &BinaryOperation{
				Common: *common,
			}
		case "Conditional":
			return &Conditional{
				Common: *common,
			}
		case "ElementaryTypeNameExpression":
			return &ElementaryTypeNameExpression{
				Common: *common,
			}
		case "FunctionCall":
			return &FunctionCall{
				Common: *common,
			}
		case "FunctionCallOptions":
			return &FunctionCallOptions{
				Common: *common,
			}
		case "Identifier":
			return &Identifier{
				Common: *common,
			}
		case "IndexAccess":
			return &IndexAccess{
				Common: *common,
			}
		case "Literal":
			return &Literal{
				Common: *common,
			}
		case "MemberAccess":
			return &MemberAccess{
				Common: *common,
			}
		case "NewExpression":
			return &NewExpression{
				Common: *common,
			}
		case "TupleExpression":
			return &TupleExpression{
				Common: *common,
			}
		case "UnaryOperation":
			return &UnaryOperation{
				Common: *common,
			}
		}

	}
	return nil
}

// CONST ----------------------------------------------------------------------------
type LiteralKind string

const (
	StringLiteral  LiteralKind = "string"
	NumberLiteral  LiteralKind = "number"
	BooleanLiteral LiteralKind = "bool"
	HexLiteral     LiteralKind = "hexString"
	UnicodeLiteral LiteralKind = "unicodeString"
)

type LiteralSubdenomination string

const (
	WeeksSubdenomination   LiteralSubdenomination = "weeks"
	DaysSubdenomination    LiteralSubdenomination = "days"
	HoursSubdenomination   LiteralSubdenomination = "hours"
	MinutesSubdenomination LiteralSubdenomination = "minutes"
	SecondsSubdenomination LiteralSubdenomination = "seconds"
	WeiSubdenomination     LiteralSubdenomination = "wei"
	GweiSubdenomination    LiteralSubdenomination = "gwei"
	EtherSubdenomination   LiteralSubdenomination = "ether"
	FinnySubdenomination   LiteralSubdenomination = "finny"
	SzaboSubdenomination   LiteralSubdenomination = "szabo"
)

type FunctionCallKind string

const (
	functionCall          FunctionCallKind = "functionCall"
	typeConversion        FunctionCallKind = "typeConversion"
	structConstructorCall FunctionCallKind = "structConstructorCall"
)

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

type UnaryOperator string

const (
	NotOperator        UnaryOperator = "!"
	ComplementOperator UnaryOperator = "~"
	IncrementOperator  UnaryOperator = "++"
	DecrementOperator  UnaryOperator = "--"
	NegativeOperator   UnaryOperator = "-"
	DeleteOperator     UnaryOperator = "delete"
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
		data := data.(map[string]interface{})
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
	if data, ok := (*data)["argumentTypes"]; ok {
		var res []TypeDescriptions
		if data != nil {
			for _, v := range data.([]interface{}) {
				res = append(res, v.(TypeDescriptions))
			}
		} else {
			res = make([]TypeDescriptions, 0)
		}
		b.ArgumentTypes = res
	}
	if data, ok := (*data)["commonType"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		b.CommonType = res
	}
	if data, ok := (*data)["function"]; ok {
		var res int
		if data != nil {
			res = (int)(data.(float64))
		} else {
			res = -1
		}
		b.Function = res
	}
	if data, ok := (*data)["isConstant"]; ok {
		var res bool
		res = data.(bool)
		b.IsConstant = res
	}
	if data, ok := (*data)["isLValue"]; ok {
		var res bool
		res = data.(bool)
		b.IsLValue = res
	}
	if data, ok := (*data)["isPure"]; ok {
		var res bool
		res = data.(bool)
		b.IsPure = res
	}
	if data, ok := (*data)["lValueRequested"]; ok {
		var res bool
		res = data.(bool)
		b.LValueRequested = res
	}
	if data, ok := (*data)["leftExpression"].(map[string]interface{}); ok {
		var res Expression
		res = ExpressionFactory(&data)
		res.Constructor(&data)
		b.LeftExpression = res
	}
	if data, ok := (*data)["operator"]; ok {
		var res Operator
		res = Operator(data.(string))
		b.Operator = res
	}
	if data, ok := (*data)["rightExpression"].(map[string]interface{}); ok {
		var res Expression
		res = ExpressionFactory(&data)
		res.Constructor(&data)
		b.RightExpression = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		b.TypeDescriptions = res
	}
}

func (b *BinaryOperation) DescribeExpression() string {
	return fmt.Sprintf("This is a binary operation expression.")
}

// ----------------------------------------------------------------------------
// Conditional node
type Conditional struct {
	Common
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
		data := data.(map[string]interface{})
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

// ----------------------------------------------------------------------------
// ElementaryTypeNameExpression node

type ElementaryTypeNameExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
	TypeName         ElementaryTypeName `json:"typeName"`
}

func (e *ElementaryTypeNameExpression) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    e.ArgumentTypes,
		"IsConstant":       e.IsConstant,
		"IsLValue":         e.IsLValue,
		"IsPure":           e.IsPure,
		"LValueRequested":  e.LValueRequested,
		"TypeDescriptions": e.TypeDescriptions,
		"TypeName":         e.TypeName,
	}
}

func (e *ElementaryTypeNameExpression) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		e.ArgumentTypes = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		e.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		e.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		e.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		e.LValueRequested = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		e.TypeDescriptions = data.(TypeDescriptions)
	}
	if data, ok := (*data)["TypeName"]; ok {
		e.TypeName = data.(ElementaryTypeName)
	}
}

func (e *ElementaryTypeNameExpression) DescribeExpression() string {
	return fmt.Sprintf("This is an elementary type name expression.")
}

// ----------------------------------------------------------------------------
// FunctionCall node
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
	NameLocations    []string           `json:"nameLocation"`
	Names            []string           `json:"names"`
	TryCall          bool               `json:"tryCall"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (f *FunctionCall) Attributes() map[string]interface{} {
	return map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"]; ok {
		var res []TypeDescriptions
		if data != nil {
			for _, v := range data.([]interface{}) {
				res = append(res, v.(TypeDescriptions))
			}
		} else {
			res = make([]TypeDescriptions, 0)
		}
		f.ArgumentTypes = res
	}
	if data, ok := (*data)["arguments"]; ok {
		var res []Expression
		for _, v := range data.([]interface{}) {
			v := v.(map[string]interface{})
			res = append(res, ExpressionFactory(&v))
		}
		f.Arguments = res
	}
	if data, ok := (*data)["expression"]; ok {
		data := data.(map[string]interface{})
		res := ExpressionFactory(&data)
		res.Constructor(&data)
		f.Expression = res
	}
	if data, ok := (*data)["isConstant"]; ok {
		var res bool
		res = data.(bool)
		f.IsConstant = res
	}
	if data, ok := (*data)["isLValue"]; ok {
		var res bool
		res = data.(bool)
		f.IsLValue = res
	}
	if data, ok := (*data)["isPure"]; ok {
		var res bool
		res = data.(bool)
		f.IsPure = res
	}
	if data, ok := (*data)["kind"]; ok {
		f.Kind = FunctionCallKind(data.(string))
	}
	if data, ok := (*data)["lValueRequested"]; ok {
		var res bool
		res = data.(bool)
		f.LValueRequested = res
	}
	if data, ok := (*data)["nameLocations"]; ok {
		if data == nil {
			f.NameLocations = make([]string, 0)
		} else {
			var res []string
			for _, v := range data.([]interface{}) {
				res = append(res, v.(string))
			}
			f.NameLocations = res
		}
	}
	if data, ok := (*data)["names"]; ok {
		var res []string
		for _, v := range data.([]interface{}) {
			res = append(res, v.(string))
		}
		f.Names = res
	}
	if data, ok := (*data)["tryCall"]; ok {
		var res bool
		res = data.(bool)
		f.TryCall = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		f.TypeDescriptions = res
	}
}

func (f *FunctionCall) DescribeExpression() string {
	return fmt.Sprintf("This is a function call expression.")
}

// ----------------------------------------------------------------------------
// FunctionCallOptions node
type FunctionCallOptions struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	Expression       Expression         `json:"expression"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	Names            []string           `json:"names"`
	Options          []Expression       `json:"options"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (f *FunctionCallOptions) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    f.ArgumentTypes,
		"Expression":       f.Expression,
		"IsConstant":       f.IsConstant,
		"IsLValue":         f.IsLValue,
		"IsPure":           f.IsPure,
		"LValueRequested":  f.LValueRequested,
		"Names":            f.Names,
		"Options":          f.Options,
		"TypeDescriptions": f.TypeDescriptions,
	}
}

func (f *FunctionCallOptions) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		f.ArgumentTypes = res
	}
	if data, ok := (*data)["expression"]; ok {
		var res Expression
		data := data.(map[string]interface{})
		res.Constructor(&data)
		f.Expression = res
	}
	if data, ok := (*data)["isConstant"]; ok {
		var res bool
		res = data.(bool)
		f.IsConstant = res
	}
	if data, ok := (*data)["isLValue"]; ok {
		var res bool
		res = data.(bool)
		f.IsLValue = res
	}
	if data, ok := (*data)["isPure"]; ok {
		var res bool
		res = data.(bool)
		f.IsPure = res
	}
	if data, ok := (*data)["lValueRequested"]; ok {
		var res bool
		res = data.(bool)
		f.LValueRequested = res
	}
	if data, ok := (*data)["names"]; ok {
		var res []string
		for _, v := range data.([]interface{}) {
			res = append(res, v.(string))
		}
		f.Names = res
	}
	if data, ok := (*data)["options"]; ok {
		var res []Expression
		for _, v := range data.([]interface{}) {
			res = append(res, v.(Expression))
		}
		f.Options = res
	}
	if data, ok := (*data)["typeDescriptions"]; ok {
		f.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (f *FunctionCallOptions) DescribeExpression() string {
	return fmt.Sprintf("This is a function call options expression.")
}

// ----------------------------------------------------------------------------
// Identifier node
type Identifier struct {
	Common
	ArgumentTypes          []TypeDescriptions `json:"argumentTypes"`
	Name                   string             `json:"name"`
	OverloadedDeclarations []int              `json:"overloadedDeclarations"`
	ReferencedDeclaration  int                `json:"referencedDeclaration"`
	TypeDescriptions       TypeDescriptions   `json:"typeDescriptions"`
}

func (i *Identifier) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":          i.ArgumentTypes,
		"Name":                   i.Name,
		"OverloadedDeclarations": i.OverloadedDeclarations,
		"ReferencedDeclaration":  i.ReferencedDeclaration,
		"TypeDescriptions":       i.TypeDescriptions,
	}
}

func (i *Identifier) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"]; ok {
		var res []TypeDescriptions
		if data != nil {
			for _, v := range data.([]interface{}) {
				var r TypeDescriptions
				v := v.(map[string]interface{})
				r.Constructor(&v)
				res = append(res, r)
			}
		} else {
			res = make([]TypeDescriptions, 0)
		}
		i.ArgumentTypes = res
	}
	if data, ok := (*data)["name"]; ok {
		i.Name = data.(string)
	}
	if data, ok := (*data)["overloadedDeclarations"]; ok {
		var res []int
		for _, v := range data.([]interface{}) {
			r := (int)(v.(float64))
			res = append(res, r)
		}
		i.OverloadedDeclarations = res
	}
	if data, ok := (*data)["referencedDeclaration"]; ok {
		var res int
		if data != nil {
			res = (int)(data.(float64))
		} else {
			res = -1
		}
		i.ReferencedDeclaration = res
	}
	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		var res TypeDescriptions
		res.Constructor(&data)
		i.TypeDescriptions = res
	}
}

func (i *Identifier) DescribeExpression() string {
	return fmt.Sprintf("This is an identifier expression.")
}

// ----------------------------------------------------------------------------
// IndexAccess node
type IndexAccess struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	BaseExpression   Expression         `json:"baseExpression"`
	IndexExpression  Expression         `json:"indexExpression"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (i *IndexAccess) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    i.ArgumentTypes,
		"BaseExpression":   i.BaseExpression,
		"IndexExpression":  i.IndexExpression,
		"IsConstant":       i.IsConstant,
		"IsLValue":         i.IsLValue,
		"IsPure":           i.IsPure,
		"LValueRequested":  i.LValueRequested,
		"TypeDescriptions": i.TypeDescriptions,
	}
}

func (i *IndexAccess) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		i.ArgumentTypes = res
	}
	if data, ok := (*data)["BaseExpression"]; ok {
		var res Expression
		data := data.(map[string]interface{})
		res.Constructor(&data)
		i.BaseExpression = res
	}
	if data, ok := (*data)["IndexExpression"]; ok {
		var res Expression
		data := data.(map[string]interface{})
		res.Constructor(&data)
		i.IndexExpression = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		i.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		i.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		i.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		i.LValueRequested = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		i.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (i *IndexAccess) DescribeExpression() string {
	return fmt.Sprintf("This is an index access expression.")
}

// ----------------------------------------------------------------------------
// Literal node
type Literal struct {
	Common
	ArgumentTypes    []TypeDescriptions     `json:"argumentTypes"`
	HexValue         string                 `json:"hexValue"`
	IsConstant       bool                   `json:"isConstant"`
	IsLValue         bool                   `json:"isLValue"`
	IsPure           bool                   `json:"isPure"`
	LValueRequested  bool                   `json:"lValueRequested"`
	Kind             LiteralKind            `json:"kind"`
	Subdenomination  LiteralSubdenomination `json:"subdenomination"`
	TypeDescriptions TypeDescriptions       `json:"typeDescriptions"`
	Value            string                 `json:"value"`
}

func (l *Literal) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    l.ArgumentTypes,
		"HexValue":         l.HexValue,
		"IsConstant":       l.IsConstant,
		"IsLValue":         l.IsLValue,
		"IsPure":           l.IsPure,
		"LValueRequested":  l.LValueRequested,
		"Kind":             l.Kind,
		"Subdenomination":  l.Subdenomination,
		"TypeDescriptions": l.TypeDescriptions,
		"Value":            l.Value,
	}
}

func (l *Literal) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		l.ArgumentTypes = res
	}
	if data, ok := (*data)["HexValue"]; ok {
		l.HexValue = data.(string)
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		l.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		l.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		l.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		l.LValueRequested = res
	}
	if data, ok := (*data)["Kind"]; ok {
		l.Kind = LiteralKind(data.(string))
	}
	if data, ok := (*data)["Subdenomination"]; ok {
		l.Subdenomination = LiteralSubdenomination(data.(string))
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		l.TypeDescriptions = data.(TypeDescriptions)
	}
	if data, ok := (*data)["Value"]; ok {
		l.Value = data.(string)
	}
}

func (l *Literal) DescribeExpression() string {
	return fmt.Sprintf("This is a literal expression.")
}

// ----------------------------------------------------------------------------
// MemberAccess node
type MemberAccess struct {
	Common
	ArgumentTypes         []TypeDescriptions `json:"argumentTypes"`
	Expression            Expression         `json:"expression"`
	IsConstant            bool               `json:"isConstant"`
	IsLValue              bool               `json:"isLValue"`
	IsPure                bool               `json:"isPure"`
	LValueRequested       bool               `json:"lValueRequested"`
	MemberLocation        string             `json:"memberLocation"`
	MemberName            string             `json:"memberName"`
	ReferencedDeclaration int                `json:"referencedDeclaration"`
	TypeDescriptions      TypeDescriptions   `json:"typeDescriptions"`
}

func (m *MemberAccess) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":         m.ArgumentTypes,
		"Expression":            m.Expression,
		"IsConstant":            m.IsConstant,
		"IsLValue":              m.IsLValue,
		"IsPure":                m.IsPure,
		"LValueRequested":       m.LValueRequested,
		"MemberLocation":        m.MemberLocation,
		"MemberName":            m.MemberName,
		"ReferencedDeclaration": m.ReferencedDeclaration,
		"TypeDescriptions":      m.TypeDescriptions,
	}
}

func (m *MemberAccess) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		m.ArgumentTypes = res
	}
	if data, ok := (*data)["Expression"]; ok {
		var res Expression
		data := data.(map[string]interface{})
		res.Constructor(&data)
		m.Expression = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		m.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		m.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		m.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		m.LValueRequested = res
	}
	if data, ok := (*data)["MemberLocation"]; ok {
		m.MemberLocation = data.(string)
	}
	if data, ok := (*data)["MemberName"]; ok {
		m.MemberName = data.(string)
	}
	if data, ok := (*data)["ReferencedDeclaration"]; ok {
		var res int
		res = data.(int)
		m.ReferencedDeclaration = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		m.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (m *MemberAccess) DescribeExpression() string {
	return fmt.Sprintf("This is a member access expression.")
}

// ----------------------------------------------------------------------------
// NewExpression node
type NewExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
	TypeName         TypeName           `json:"typeName"`
}

func (n *NewExpression) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    n.ArgumentTypes,
		"IsConstant":       n.IsConstant,
		"IsLValue":         n.IsLValue,
		"IsPure":           n.IsPure,
		"LValueRequested":  n.LValueRequested,
		"TypeDescriptions": n.TypeDescriptions,
		"TypeName":         n.TypeName,
	}
}

func (n *NewExpression) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		n.ArgumentTypes = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		n.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		n.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		n.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		n.LValueRequested = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		n.TypeDescriptions = data.(TypeDescriptions)
	}
	if data, ok := (*data)["TypeName"]; ok {
		n.TypeName = data.(TypeName)
	}
}

func (n *NewExpression) DescribeExpression() string {
	return fmt.Sprintf("This is a new expression.")
}

// ----------------------------------------------------------------------------
// TupleExpression node
type TupleExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	Components       []Expression       `json:"components"`
	IsConstant       bool               `json:"isConstant"`
	IsInlineArray    bool               `json:"isInlineArray"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (t *TupleExpression) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ArgumentTypes":    t.ArgumentTypes,
		"Components":       t.Components,
		"IsConstant":       t.IsConstant,
		"IsInlineArray":    t.IsInlineArray,
		"IsLValue":         t.IsLValue,
		"IsPure":           t.IsPure,
		"LValueRequested":  t.LValueRequested,
		"TypeDescriptions": t.TypeDescriptions,
	}
}

func (t *TupleExpression) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		t.ArgumentTypes = res
	}
	if data, ok := (*data)["Components"]; ok {
		var res []Expression
		for _, v := range data.([]interface{}) {
			res = append(res, v.(Expression))
		}
		t.Components = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		t.IsConstant = res
	}
	if data, ok := (*data)["IsInlineArray"]; ok {
		var res bool
		res = data.(bool)
		t.IsInlineArray = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		t.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		t.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		t.LValueRequested = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		t.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (t *TupleExpression) DescribeExpression() string {
	return fmt.Sprintf("This is a tuple expression.")
}

// ----------------------------------------------------------------------------
// UnaryOperation node
type UnaryOperation struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"`
	Function         int                `json:"function"`
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	Operator         UnaryOperator      `json:"operator"`
	Prefix           bool               `json:"prefix"`
	SubExpression    Expression         `json:"subExpression"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"`
}

func (u *UnaryOperation) Attributes() map[string]interface{} {
	return map[string]interface{}{
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
	if data, ok := (*data)["ArgumentTypes"]; ok {
		var res []TypeDescriptions
		for _, v := range data.([]interface{}) {
			res = append(res, v.(TypeDescriptions))
		}
		u.ArgumentTypes = res
	}
	if data, ok := (*data)["Function"]; ok {
		var res int
		res = data.(int)
		u.Function = res
	}
	if data, ok := (*data)["IsConstant"]; ok {
		var res bool
		res = data.(bool)
		u.IsConstant = res
	}
	if data, ok := (*data)["IsLValue"]; ok {
		var res bool
		res = data.(bool)
		u.IsLValue = res
	}
	if data, ok := (*data)["IsPure"]; ok {
		var res bool
		res = data.(bool)
		u.IsPure = res
	}
	if data, ok := (*data)["LValueRequested"]; ok {
		var res bool
		res = data.(bool)
		u.LValueRequested = res
	}
	if data, ok := (*data)["Operator"]; ok {
		u.Operator = UnaryOperator(data.(string))
	}
	if data, ok := (*data)["Prefix"]; ok {
		var res bool
		res = data.(bool)
		u.Prefix = res
	}
	if data, ok := (*data)["SubExpression"]; ok {
		var res Expression
		data := data.(map[string]interface{})
		res.Constructor(&data)
		u.SubExpression = res
	}
	if data, ok := (*data)["TypeDescriptions"]; ok {
		u.TypeDescriptions = data.(TypeDescriptions)
	}
}

func (u *UnaryOperation) DescribeExpression() string {
	return fmt.Sprintf("This is a unary operation expression.")
}
