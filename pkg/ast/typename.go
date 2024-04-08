package ast

// TypeName: ArrayTypeName | ElementaryTypeName | FunctionTypeName | Mapping | UserDefinedTypeName
type TypeName interface {
	DescribeTypeName() string
	Constructor(*map[string]interface{})
}
