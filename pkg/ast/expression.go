package ast

// Expression: Assignment | BinaryOperation | Conditional | ElementaryTypeNameExpression | FunctionCall | FunctionCallOptions | Identifier | IndexAccess | IndexRangeAccess | Literal | MemberAccess | NewExpression | TupleExpression | UnaryOperation
type Expression interface {
	DescribeExpression() string
	Constructor(*map[string]interface{})
}
