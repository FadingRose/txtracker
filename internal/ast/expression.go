package ast

// Expression: Assignment | BinaryOperation | Conditional | ElementaryTypeNameExpression | FunctionCall | FunctionCallOptions | Identifier | IndexAccess | IndexRangeAccess | Literal | MemberAccess | NewExpression | TupleExpression | UnaryOperation
type Expression = *Common

func (e *Common) RetrieveVarSymbols() []string {
	var symbols []string
	switch (*e).NodeType {
	case "Assignment":
		assign := (*e).ASTNode.(*Assignment)
		symbols = append(symbols, assign.LeftHandSide.RetrieveVarSymbols()...)
		symbols = append(symbols, assign.RightHandSide.RetrieveVarSymbols()...)
	case "BinaryOperation":
		binOp := (*e).ASTNode.(*BinaryOperation)
		symbols = append(symbols, binOp.LeftExpression.RetrieveVarSymbols()...)
		symbols = append(symbols, binOp.RightExpression.RetrieveVarSymbols()...)
	// case "Conditional":
	// 	cond := (*e).ASTNode.(*Conditional)
	// 	symbols = append(symbols, cond.TrueExpression.RetrieveVarSymbols()...)
	// 	symbols = append(symbols, cond.FalseExpression.RetrieveVarSymbols()...)
	case "ElementaryTypeNameExpression":
		// do nothing
	case "FunctionCall":
		funCall := (*e).ASTNode.(*FunctionCall)
		symbols = append(symbols, funCall.Expression.RetrieveVarSymbols()...)
		for _, arg := range funCall.Arguments {
			symbols = append(symbols, arg.RetrieveVarSymbols()...)
		}
	case "FunctionCallOptions":
		// do nothing
	case "Identifier":
		ident := (*e).ASTNode.(*Identifier)
		symbols = append(symbols, ident.Name)
	case "IndexAccess":
		index := (*e).ASTNode.(*IndexAccess)
		symbols = append(symbols, index.BaseExpression.RetrieveVarSymbols()...)
		if index.IndexExpression != nil {
			symbols = append(symbols, index.IndexExpression.RetrieveVarSymbols()...)
		}
	case "IndexRangeAccess":
		// do nothing
	case "Literal":
		// do nothing
	case "MemberAccess":
		member := (*e).ASTNode.(*MemberAccess)
		symbols = append(symbols, member.Expression.RetrieveVarSymbols()...)
	case "NewExpression":
		// do nothing
	case "TupleExpression":
		// do nothing
		// case "UnaryOperation":
		// 	unary := (*e).ASTNode.(*UnaryOperation)
		// 	symbols = append(symbols, unary.SubExpression.RetrieveVarSymbols()...)
	}
	return symbols
}

type BinaryOperation struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	CommonType       TypeDescriptions   `json:"commonType"`
	Function         int                `json:"function"` // int | null
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	LeftExpression   Expression         `json:"leftExpression"`
	Operator         Operator           `json:"operator"`
	RightExpression  Expression         `json:"rightExpression"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
}

func (b *BinaryOperation) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			b.ArgumentTypes = append(b.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["commonType"].(map[string]interface{}); ok {
		b.CommonType.Constructor(&data)
	}

	if data, ok := (*data)["function"].(int); ok {
		b.Function = data
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		b.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		b.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		b.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		b.LValueRequested = data
	}

	if data, ok := (*data)["leftExpression"].(map[string]interface{}); ok {
		b.LeftExpression = NodeFactory(data)
		b.LeftExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["operator"].(string); ok {
		b.Operator = Operator(data)
	}

	if data, ok := (*data)["rightExpression"].(map[string]interface{}); ok {
		b.RightExpression = NodeFactory(data)
		b.RightExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		b.TypeDescriptions.Constructor(&data)
	}
}

func (b *BinaryOperation) DescribeExpression() string {
	return "BinaryOperation"
}

type Identifier struct {
	Common
	ArgumentTypes          []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	Name                   string             `json:"name"`
	OverloadedDeclarations []int              `json:"overloadedDeclarations"`
	ReferencedDeclaration  int                `json:"referencedDeclaration"` // int | null
	TypeDescriptions       TypeDescriptions   `json:"typeDescriptions"`
}

func (i *Identifier) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"ArgumentTypes":          i.ArgumentTypes,
		"Name":                   i.Name,
		"OverloadedDeclarations": i.OverloadedDeclarations,
		"ReferencedDeclaration":  i.ReferencedDeclaration,
		"TypeDescriptions":       i.TypeDescriptions,
	}
}

func (i *Identifier) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			i.ArgumentTypes = append(i.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["name"].(string); ok {
		i.Name = data
	}

	if data, ok := (*data)["overloadedDeclarations"].([]interface{}); ok {
		for _, v := range data {
			i.OverloadedDeclarations = append(i.OverloadedDeclarations, int(v.(float64)))
		}
	}

	if data, ok := (*data)["referencedDeclaration"].(int); ok {
		i.ReferencedDeclaration = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		i.TypeDescriptions.Constructor(&data)
	}
}

func (i *Identifier) DescribeExpression() string {
	return "Identifier"
}

type Literal struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	HexValue         string             `json:"hexValue"`      // string | null
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	Kind             LiteralKind        `json:"kind"`
	LValueRequested  bool               `json:"lValueRequested"`
	Subdenomination  Subdenomination    `json:"subdenomination"`  // string | null
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
	Value            string             `json:"value"`
}

func (l *Literal) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
		"ArgumentTypes":    l.ArgumentTypes,
		"HexValue":         l.HexValue,
		"IsConstant":       l.IsConstant,
		"IsLValue":         l.IsLValue,
		"IsPure":           l.IsPure,
		"Kind":             l.Kind,
		"LValueRequested":  l.LValueRequested,
		"Subdenomination":  l.Subdenomination,
		"TypeDescriptions": l.TypeDescriptions,
		"Value":            l.Value,
	}
}

func (l *Literal) Constructor(data *map[string]interface{}) {
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			l.ArgumentTypes = append(l.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["hexValue"].(string); ok {
		l.HexValue = data
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		l.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		l.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		l.IsPure = data
	}

	if data, ok := (*data)["kind"].(string); ok {
		l.Kind = LiteralKind(data)
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		l.LValueRequested = data
	}

	if data, ok := (*data)["subdenomination"].(string); ok {
		l.Subdenomination = Subdenomination(data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		l.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*data)["value"].(string); ok {
		l.Value = data
	}
}

func (l *Literal) DescribeExpression() string {
	return "Literal"
}

type Assignment struct {
	Common
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	LeftHandSide     Expression         `json:"leftHandSide"`
	Operator         AssignmentOperator `json:"operator"`
	RightHandSide    Expression         `json:"rightHandSide"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
}

func (a *Assignment) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["isConstant"].(bool); ok {
		a.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		a.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		a.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		a.LValueRequested = data
	}

	if data, ok := (*data)["leftHandSide"].(map[string]interface{}); ok {
		a.LeftHandSide = NodeFactory(data)
		a.LeftHandSide.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["operator"].(string); ok {
		a.Operator = AssignmentOperator(data)
	}

	if data, ok := (*data)["rightHandSide"].(map[string]interface{}); ok {
		a.RightHandSide = NodeFactory(data)
		a.RightHandSide.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		a.TypeDescriptions.Constructor(&data)
	}
}

type MemberAccess struct {
	Common
	ArgumentTypes         []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	Expression            Expression         `json:"expression"`
	IsConstant            bool               `json:"isConstant"`
	IsLValue              bool               `json:"isLValue"`
	IsPure                bool               `json:"isPure"`
	LValueRequested       bool               `json:"lValueRequested"`
	MemberLocation        string             `json:"memberLocation"`
	MemberName            string             `json:"memberName"`
	ReferencedDeclaration int                `json:"referencedDeclaration"` // int | null
	TypeDescriptions      TypeDescriptions   `json:"typeDescriptions"`      // TypeDescriptions | null
}

func (m *MemberAccess) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			m.ArgumentTypes = append(m.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["expression"].(map[string]interface{}); ok {
		m.Expression = NodeFactory(data)
		m.Expression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		m.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		m.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		m.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		m.LValueRequested = data
	}

	if data, ok := (*data)["memberLocation"].(string); ok {
		m.MemberLocation = data
	}

	if data, ok := (*data)["memberName"].(string); ok {
		m.MemberName = data
	}

	if data, ok := (*data)["referencedDeclaration"].(int); ok {
		m.ReferencedDeclaration = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		m.TypeDescriptions.Constructor(&data)
	}
}

type IndexAccess struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	BaseExpression   Expression         `json:"baseExpression"`
	IndexExpression  Expression         `json:"indexExpression"` // Expression | null
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
}

func (i *IndexAccess) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			i.ArgumentTypes = append(i.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["baseExpression"].(map[string]interface{}); ok {
		i.BaseExpression = NodeFactory(data)
		i.BaseExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["indexExpression"].(map[string]interface{}); ok {
		i.IndexExpression = NodeFactory(data)
		i.IndexExpression.ASTNode.Constructor(&data)
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		i.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		i.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		i.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		i.LValueRequested = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		i.TypeDescriptions = TypeDescriptions{}
		i.TypeDescriptions.Constructor(&data)
	}
}

type ElementaryTypeNameExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
	TypeName         ElementaryTypeName `json:"typeName"`
}

func (e *ElementaryTypeNameExpression) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			e.ArgumentTypes = append(e.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		e.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		e.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		e.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		e.LValueRequested = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		e.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*data)["typeName"].(map[string]interface{}); ok {
		e.TypeName = ElementaryTypeName{}
		e.TypeName.Constructor(&data)
	}
}

type TupleExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	Components       []Expression       `json:"components"`
	IsConstant       bool               `json:"isConstant"`
	IsInlineArray    bool               `json:"isInlineArray"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
}

func (t *TupleExpression) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			t.ArgumentTypes = append(t.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["components"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			expr := NodeFactory(v)
			expr.ASTNode.Constructor(&v)
			t.Components = append(t.Components, expr)
		}
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		t.IsConstant = data
	}

	if data, ok := (*data)["isInlineArray"].(bool); ok {
		t.IsInlineArray = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		t.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		t.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		t.LValueRequested = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		t.TypeDescriptions.Constructor(&data)
	}
}

type NewExpression struct {
	Common
	ArgumentTypes    []TypeDescriptions `json:"argumentTypes"` // TypeDescriptions[] | null
	IsConstant       bool               `json:"isConstant"`
	IsLValue         bool               `json:"isLValue"`
	IsPure           bool               `json:"isPure"`
	LValueRequested  bool               `json:"lValueRequested"`
	TypeDescriptions TypeDescriptions   `json:"typeDescriptions"` // TypeDescriptions | null
	TypeName         TypeName           `json:"typeName"`
}

func (n *NewExpression) Attributes() *map[string]interface{} {
	return &map[string]interface{}{
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
	if data, ok := (*data)["argumentTypes"].([]interface{}); ok {
		for _, v := range data {
			v := v.(map[string]interface{})
			td := TypeDescriptions{}
			td.Constructor(&v)
			n.ArgumentTypes = append(n.ArgumentTypes, td)
		}
	}

	if data, ok := (*data)["isConstant"].(bool); ok {
		n.IsConstant = data
	}

	if data, ok := (*data)["isLValue"].(bool); ok {
		n.IsLValue = data
	}

	if data, ok := (*data)["isPure"].(bool); ok {
		n.IsPure = data
	}

	if data, ok := (*data)["lValueRequested"].(bool); ok {
		n.LValueRequested = data
	}

	if data, ok := (*data)["typeDescriptions"].(map[string]interface{}); ok {
		n.TypeDescriptions.Constructor(&data)
	}

	if data, ok := (*data)["typeName"].(map[string]interface{}); ok {
		n.TypeName = NodeFactory(data)
		n.TypeName.ASTNode.Constructor(&data)
	}
}
