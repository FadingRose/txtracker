package models

// `Statement`:
// Block | Break | Continue | DoWhileStatement | EmitStatement | ExpressionStatement | ForStatement | IfStatement | InlineAssembly | PlaceholderStatement | Return | RevertStatement | TryStatement | UncheckedBlock | VariableDeclarationStatement | WhileStatement
import (
	"fmt"
	"txtracker/pkg/logger"
)

type Statement interface {
	DescribeStatement() string
	Constructor(*map[string]interface{})
}

type ForInitializationStatement interface {
	// VariableDeclarationStatement | ExpressionStatement implements this interface
	DescribeForInitializationStatement() string
	Constructor(*map[string]interface{})
}

// Factories
func ForInitializationStatementFactory(data *map[string]interface{}) ForInitializationStatement {
	common := &Common{
		ID:       (int)((*data)["id"].(float64)),
		NodeType: (*data)["nodeType"].(string),
		Src:      (*data)["src"].(string),
	}
	if data, ok := (*data)["nodeType"]; ok {
		switch data {
		case "VariableDeclarationStatement":
			return &VariableDeclarationStatement{Common: *common}
		case "ExpressionStatement":
			return &ExpressionStatement{Common: *common}
		default:
			logger.Fatal.Fatalf("Unknown type: %v", data)
			panic("Unknown type")
		}
	}
	return nil
}

func StatementFactory(data *map[string]interface{}) Statement {
	common := &Common{
		ID:       (int)((*data)["id"].(float64)),
		NodeType: (*data)["nodeType"].(string),
		Src:      (*data)["src"].(string),
	}
	if data, ok := (*data)["nodeType"]; ok {
		switch data {
		case "Block":
			return &Block{Common: *common}
		case "Break":
			return &Break{Common: *common}
		case "Continue":
			return &Continue{Common: *common}
		case "DoWhileStatement":
			return &DoWhileStatement{Common: *common}
		case "EmitStatement":
			return &EmitStatement{Common: *common}
		case "ExpressionStatement":
			return &ExpressionStatement{Common: *common}
		case "ForStatement":
			return &ForStatement{Common: *common}
		case "IfStatement":
			return &IfStatement{Common: *common}
		case "PlaceholderStatement":
			return &PlacehoderStatement{Common: *common}
		case "Return":
			return &Return{Common: *common}
		case "RevertStatement":
			return &RevertStatement{Common: *common}
		case "TryStatement":
			return &TryStatement{Common: *common}
		case "UncheckedBlock":
			return &UncheckedBlock{Common: *common}
		case "VariableDeclarationStatement":
			return &VariableDeclarationStatement{Common: *common}
		case "WhileStatement":
			return &WhileStatement{Common: *common}
		default:
			logger.Fatal.Fatalf("Unknown type: %v", data)
			panic("Unknown type")
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// Block node
type Block struct {
	Common
	Statements []Statement `json:"statements"`
}

func (b *Block) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Statements": b.Statements,
	}
}

func (b *Block) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["statements"]; ok {
		var res = []Statement{}
		for _, v := range value.([]interface{}) {
			v := v.(map[string]interface{})
			switch v["nodeType"] {
			case "Block":
				var r Block
				r.Constructor(&v)
				res = append(res, &r)
			case "Break":
				var r Break
				r.Constructor(&v)
				res = append(res, &r)
			case "Continue":
				var r Continue
				r.Constructor(&v)
				res = append(res, &r)
			case "DoWhileStatement":
				var r DoWhileStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "EmitStatement":
				var r EmitStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "ExpressionStatement":
				var r ExpressionStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "ForStatement":
				var r ForStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "IfStatement":
				var r IfStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "PlaceholderStatement":
				var r PlacehoderStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "Return":
				var r Return
				r.Constructor(&v)
				res = append(res, &r)
			case "RevertStatement":
				var r RevertStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "TryStatement":
				var r TryStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "UncheckedBlock":
				var r UncheckedBlock
				r.Constructor(&v)
				res = append(res, &r)
			case "VariableDeclarationStatement":
				var r VariableDeclarationStatement
				r.Constructor(&v)
				res = append(res, &r)
			case "WhileStatement":
				var r WhileStatement
				r.Constructor(&v)
				res = append(res, &r)
			default:
				logger.Fatal.Fatalf("Unknown type: %v", v["nodeType"])
				panic("Unknown type")
			}
		}
		b.Statements = res
	}
}

func (b *Block) DescribeStatement() string {
	return fmt.Sprintf("This is a block.")
}

// -----------------------------
// Break Statement Node
type Break struct {
	Common
}

func (b *Break) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (b *Break) Constructor(data *map[string]interface{}) {
}

func (b *Break) DescribeStatement() string {
	return fmt.Sprintf("This is a break statement.")
}

// ----------------------------------------------------------------------------
// Continue Statement Node
type Continue struct {
	Common
}

func (c *Continue) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (c *Continue) Constructor(data *map[string]interface{}) {
}

func (c *Continue) DescribeStatement() string {
	return fmt.Sprintf("This is a continue statement.")
}

// ----------------------------------------------------------------------------
// DoWhile Statement Node
type DoWhileStatement struct {
	Common
	Condition Expression
	Body      Statement
}

func (d *DoWhileStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Condition": d.Condition,
		"Body":      d.Body,
	}
}

func (d *DoWhileStatement) Constructor(data *map[string]interface{}) {
	if condition, ok := (*data)["condition"]; ok {
		var res Expression
		condition := condition.(map[string]interface{})
		res.Constructor(&condition)
		d.Condition = res
	}
	if body, ok := (*data)["body"]; ok {
		var res Statement
		body := body.(map[string]interface{})
		res.Constructor(&body)
		d.Body = res
	}
}

func (d *DoWhileStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a do-while statement.")
}

// ----------------------------------------------------------------------------
// Emit Statement Node
type EmitStatement struct {
	Common
	EventCall FunctionCall
}

func (e *EmitStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"EventCall": e.EventCall,
	}
}

func (e *EmitStatement) Constructor(data *map[string]interface{}) {
	if eventCall, ok := (*data)["eventCall"]; ok {
		var res FunctionCall
		eventCall := eventCall.(map[string]interface{})
		res.Constructor(&eventCall)
		e.EventCall = res
	}
}

func (e *EmitStatement) DescribeStatement() string {
	return fmt.Sprintf("This is an emit statement.")
}

// ----------------------------------------------------------------------------
// Expression Statement Node
type ExpressionStatement struct {
	Common
	Expression Expression `json:"expression"`
}

func (e *ExpressionStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Expression": e.Expression,
	}
}

func (e *ExpressionStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["expression"]; ok {
		value := value.(map[string]interface{})
		res := ExpressionFactory(&value)
		res.Constructor(&value)
		e.Expression = res
	} else {
		logger.Fatal.Fatal("ExpressionStatement Constructor: expression key not found in data")
	}
}

func (e *ExpressionStatement) DescribeStatement() string {
	return fmt.Sprintf("This is an expression statement.")
}

func (e *ExpressionStatement) DescribeForInitializationStatement() string {
	return fmt.Sprintf("This is an for initialization statement.")
}

// ----------------------------------------------------------------------------
// For Statement Node
type ForStatement struct {
	Common
	Body                     Statement                  `json:"body"`
	Condition                Expression                 `json:"condition"`                // Expression | nil
	InitializationExpression ForInitializationStatement `json:"initializationExpression"` // VariableDeclarationStatement | ExpressionStatement | nil
	IsSimpleCounterLoop      bool                       `json:"isSimpleCounterLoop"`      // true | false | nil
	LoopExpression           ExpressionStatement        `json:"loopExpression"`           // ExpressionStatement | nil
}

func (f *ForStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Body":                     f.Body,
		"Condition":                f.Condition,
		"InitializationExpression": f.InitializationExpression,
		"LoopExpression":           f.LoopExpression,
		"IsSimpleCounterLoop":      f.IsSimpleCounterLoop,
	}
}

func (f *ForStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["body"]; ok {
		value := value.(map[string]interface{})
		res := StatementFactory(&value)
		res.Constructor(&value)
		f.Body = res
	}
	if value, ok := (*data)["condition"]; ok {
		value := value.(map[string]interface{})
		res := ExpressionFactory(&value)
		res.Constructor(&value)
		f.Condition = res
	} else {
		f.Condition = nil
	}
	if value, ok := (*data)["initializationExpression"]; ok {
		value := value.(map[string]interface{})
		res := ForInitializationStatementFactory(&value)
		res.Constructor(&value)
		f.InitializationExpression = res
	} else {
		f.InitializationExpression = nil
	}
	if value, ok := (*data)["isSimpleCounterLoop"]; ok {
		var res bool
		res = value.(bool)
		f.IsSimpleCounterLoop = res
	} else {
		f.IsSimpleCounterLoop = false // default value
	}
	if value, ok := (*data)["loopExpression"]; ok {
		var res ExpressionStatement
		value := value.(map[string]interface{})
		res.Constructor(&value)
		f.LoopExpression = res
	}
}

func (f *ForStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a for statement.")
}

// ----------------------------------------------------------------------------
// IfStatement node
type IfStatement struct {
	Common
	Condition Expression `json:"condition"`
	FalseBody Statement  `json:"falseBody"`
	TrueBody  Statement  `json:"trueBody"`
}

func (i *IfStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Condition": i.Condition,
		"FalseBody": i.FalseBody,
		"TrueBody":  i.TrueBody,
	}
}

func (i *IfStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["condition"]; ok {
		value := value.(map[string]interface{})
		res := ExpressionFactory(&value)
		res.Constructor(&value)
		i.Condition = res
	}
	if value, ok := (*data)["falseBody"]; ok {
		var res Statement
		if value != nil {
			value := value.(map[string]interface{})
			res.Constructor(&value)
		} else {
			res = nil
		}
		i.FalseBody = res
	}
	if value, ok := (*data)["trueBody"].(map[string]interface{}); ok {
		res := StatementFactory(&value)
		res.Constructor(&value)
		i.TrueBody = res
	}
}

func (i *IfStatement) DescribeStatement() string {
	return fmt.Sprintf("This is an if statement.")
}

// ----------------------------------------------------------------------------
// PlacehoderStatement node
type PlacehoderStatement struct {
	Common
}

func (p *PlacehoderStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *PlacehoderStatement) Constructor(data *map[string]interface{}) {
}

func (p *PlacehoderStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a placeholder statement.")
}

// ----------------------------------------------------------------------------
// ReturnStatement node
type Return struct {
	Common
	Expression               Expression `json:"expression"`
	FunctionReturnParameters int        `json:"functionReturnParameters"`
}

func (r *Return) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Expression":               r.Expression,
		"FunctionReturnParameters": r.FunctionReturnParameters,
	}
}

func (r *Return) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["expression"]; ok {
		if value != nil {
			value := value.(map[string]interface{})
			res := ExpressionFactory(&value)
			res.Constructor(&value)
		} else {
			r.Expression = nil
		}
	}
	if value, ok := (*data)["functionReturnParameters"]; ok {
		var res int
		res = (int)(value.(float64))
		r.FunctionReturnParameters = res
	}
}

func (r *Return) DescribeStatement() string {
	return fmt.Sprintf("This is a return statement.")
}

// ----------------------------------------------------------------------------
// RevertStatement node
type RevertStatement struct {
	Common
	ErrorCall FunctionCall `json:"errorCall"`
}

func (r *RevertStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ErrorCall": r.ErrorCall,
	}
}

func (r *RevertStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["errorCall"]; ok {
		var res FunctionCall
		value := value.(map[string]interface{})
		res.Constructor(&value)
		r.ErrorCall = res
	}
}

func (r *RevertStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a revert statement.")
}

// ----------------------------------------------------------------------------
// TryStatement node
type TryStatement struct {
	Common
	Clauses      []TryCatchClause
	ExternalCall FunctionCall
}

func (t *TryStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Clauses":      t.Clauses,
		"ExternalCall": t.ExternalCall,
	}
}

func (t *TryStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["clauses"]; ok {
		var res []TryCatchClause
		for _, val := range value.([]map[string]interface{}) {
			var temp TryCatchClause
			temp.Constructor(&val)
			res = append(res, temp)
		}
		t.Clauses = res
	}
	if value, ok := (*data)["externalCall"]; ok {
		var res FunctionCall
		value := value.(map[string]interface{})
		res.Constructor(&value)
		t.ExternalCall = res
	}
}

func (t *TryStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a try statement.")
}

// ----------------------------------------------------------------------------
// UncheckedBlock node
type UncheckedBlock struct {
	Common
	Statements []Statement
}

func (u *UncheckedBlock) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Statements": u.Statements,
	}
}

func (u *UncheckedBlock) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["statements"]; ok {
		var res []Statement
		for _, val := range value.([]map[string]interface{}) {
			var temp Statement
			temp.Constructor(&val)
			res = append(res, temp)
		}
		u.Statements = res
	}
}

func (u *UncheckedBlock) DescribeStatement() string {
	return fmt.Sprintf("This is an unchecked block.")
}

// ----------------------------------------------------------------------------
// VariableDeclarationStatement node

type VariableDeclarationStatement struct {
	Common
	Assignments   []int                 `json:"assignments"`
	Declarrations []VariableDeclaration `json:"declarations"`
	InitialValue  Expression            `json:"initialValue"`
}

func (v *VariableDeclarationStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Assignments":   v.Assignments,
		"Declarrations": v.Declarrations,
		"InitialValue":  v.InitialValue,
	}
}

func (v *VariableDeclarationStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["assignments"]; ok {
		if value != nil {
			var res []int
			for _, val := range value.([]interface{}) {
				res = append(res, (int)(val.(float64)))
			}
			v.Assignments = res
		} else {
			v.Assignments = nil
		}
	}
	if value, ok := (*data)["declarations"]; ok {
		if value != nil {
			var res []VariableDeclaration
			for _, val := range value.([]interface{}) {
				var temp VariableDeclaration
				val := val.(map[string]interface{})
				temp.Constructor(&val)
				res = append(res, temp)
			}
			v.Declarrations = res
		} else {
			v.Declarrations = nil
		}
	}
	if value, ok := (*data)["initialValue"]; ok {
		if value != nil {
			value := value.(map[string]interface{})
			res := ExpressionFactory(&value)
			res.Constructor(&value)
			v.InitialValue = res
		} else {
			v.InitialValue = nil
		}
	}
}

func (v *VariableDeclarationStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a variable declaration statement.")
}

func (v *VariableDeclarationStatement) DescribeForInitializationStatement() string {
	return fmt.Sprintf("This is a for initialization statement.")
}

// ----------------------------------------------------------------------------
// WhileStatement node
type WhileStatement struct {
	Common
	Body      Statement  `json:"body"`
	Condition Expression `json:"condition"`
}

func (w *WhileStatement) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Body":      w.Body,
		"Condition": w.Condition,
	}
}

func (w *WhileStatement) Constructor(data *map[string]interface{}) {
	if value, ok := (*data)["body"]; ok {
		var res Statement
		value := value.(map[string]interface{})
		res.Constructor(&value)
		w.Body = res
	}
	if value, ok := (*data)["condition"]; ok {
		var res Expression
		value := value.(map[string]interface{})
		res.Constructor(&value)
		w.Condition = res
	}
}

func (w *WhileStatement) DescribeStatement() string {
	return fmt.Sprintf("This is a while statement.")
}
