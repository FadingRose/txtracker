package symboltable

import (
	"strings"
	"txtracker/pkg/ast"
	"txtracker/pkg/logger"
)

type GlobalSymbolTable struct {
	Table map[string]Symbol
}

type Symbol struct {
	Namespace Namespace
	Type      SymbolType
	Arributes map[string]interface{}
}

type Namespace []string

func NewNamespace() Namespace {
	return make(Namespace, 0)
}

func (n Namespace) Push(name string) Namespace {
	if name == "" {
		// means a constructor or fallback function
		return n
	}
	return append(n, name)
}

func (n Namespace) Pop() Namespace {
	if len(n) == 0 {
		return n
	}
	return n[:len(n)-1]
}

func (n Namespace) String() string {
	return strings.Join(n, "::")
}

func (s Symbol) Name() string {
	return s.Namespace[len(s.Namespace)-1]
}

type SymbolType int

const (
	StateVariable SymbolType = iota
	Function
	Constructor
	Fallback
	Receive
	FreeFunction
)

func NewGlobalSymbolTable(root *ast.Common) *GlobalSymbolTable {
	gst := &GlobalSymbolTable{
		Table: make(map[string]Symbol),
	}

	contractDefs := _findContractDefinition(root)

	var symbols []*Symbol
	for _, contractDef := range contractDefs {
		symbols = append(symbols, _findGlobalSymbols(contractDef)...)
	}

	for _, symbol := range symbols {
		gst.InsertSymbol(*symbol)
	}

	return gst
}

func (gst *GlobalSymbolTable) InsertSymbol(symbol Symbol) {
	gst.Table[symbol.Namespace.String()] = symbol
}

func (gst *GlobalSymbolTable) LookupSymbol(symbolName string) Symbol {
	return gst.Table[symbolName]
}

// This function do NOT check namespace
func (gst *GlobalSymbolTable) IsExistWithIdentifierOnly(varname string) bool {
	for _, symbol := range gst.Table {
		if symbol.Name() == varname {
			return true
		}
	}
	return false
}

func _findContractDefinition(root *ast.Common) []*ast.Common {
	var res []*ast.Common
	if root.NodeType == string(SourceUnit) {
		childs := root.Children
		for _, child := range childs {
			if child.NodeType == string(ContractDefinition) {
				res = append(res, child)
			}
		}
	}
	return res
}

func _findGlobalSymbols(contractDef *ast.Common) []*Symbol {
	var res []*Symbol
	childs := contractDef.Children
	for _, child := range childs {
		if child.NodeType == string(VariableDeclaration) {
			res = append(res, &Symbol{
				Namespace: _findNamespace(child),
				Type:      StateVariable,
				Arributes: *child.ASTNode.Attributes(),
			})
		} else if child.NodeType == string(FunctionDefinition) {
			// whether it is a constructor
			attr := *child.ASTNode.Attributes()
			if data, ok := attr["Kind"].(ast.FunctionKind); ok {
				var functype SymbolType
				switch data {
				case ast.FunctionKind_Constructor:
					functype = Constructor
				case ast.FunctionKind_Fallback:
					functype = Fallback
				case ast.FunctionKind_Receive:
					functype = Receive
				case ast.FunctionKind_FreeFunction:
					functype = FreeFunction
				default:
					functype = Function
				}
				res = append(res, &Symbol{
					Namespace: _findNamespace(child),
					Type:      functype,
					Arributes: *child.ASTNode.Attributes(),
				})
			} else {
				logger.Warning.Println("")
			}

		}
	}
	return res

}

func _findNamespace(node *ast.Common) Namespace {
	var res Namespace

	contractName := _findRootContractName(node)
	if contractName != "" {
		res = append(res, contractName)
	} else {
		logger.Fatal.Println("Error: Contract name not found")
		panic("Error: Contract name not found")
	}

	if node.NodeType == string(VariableDeclaration) {
		attr := node.ASTNode.Attributes()
		if attr != nil {
			attr := *attr
			if name, ok := attr["Name"]; ok {
				res = append(res, name.(string))
			}
		}
	} else if node.NodeType == string(FunctionDefinition) {
		attr := node.ASTNode.Attributes()
		if attr != nil {
			attr := *attr
			if name, ok := attr["Name"]; ok {
				res = append(res, name.(string))
			}
		}
	}

	return res
}

func _findRootContractName(node *ast.Common) string {
	for node.NodeType != string(ContractDefinition) {
		node = node.Parent
	}
	attr := node.ASTNode.Attributes()
	if attr != nil {
		attr := *attr
		if name, ok := attr["Name"]; ok {
			return name.(string)
		}
	}
	return ""
}
