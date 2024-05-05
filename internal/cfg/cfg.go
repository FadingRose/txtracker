package cfg

import (
	AST "txtracker/internal/ast"
	"txtracker/internal/logger"
	ST "txtracker/internal/symbol_table"
)

func NewCFG(root *AST.Common, symbolTable *ST.GlobalSymbolTable) *CFG {
	cfg := &CFG{
		symbolTable: symbolTable,
		Visitor:     NewVisitor(),
	}
	cfg.EntryPoints = cfg._constructEntryFuncs(root)

	logger.Info.Println("CFG constructed")
	return cfg
}

func (cfg *CFG) _constructEntryFuncs(root *AST.Common) []*Function {
	contractDefs := cfg._findContractDefinition(root)
	var entryFuncs []*Function
	for _, contractDef := range contractDefs {
		entryFuncs = append(entryFuncs, cfg._findEntryFunc(contractDef)...)
	}

	return entryFuncs
}

func (cfg *CFG) _findContractDefinition(root *AST.Common) []*AST.Common {
	var contractDefs []*AST.Common

	for _, node := range root.Children {
		if node.NodeType == "ContractDefinition" {
			contractDefs = append(contractDefs, node)
		}
	}

	return contractDefs
}

func (cfg *CFG) _findEntryFunc(contractDef *AST.Common) []*Function {
	var entryFuncs []*Function
	contractName := contractDef.ASTNode.(*AST.ContractDefinition).Name
	cfg.Visitor.EnterNamespace(contractName)
	for _, node := range contractDef.Children {

		if node.NodeType == "FunctionDefinition" {
			funcDef := node.ASTNode.(*AST.FunctionDefinition)

			if funcDef.IsPublic() || funcDef.IsExternal() {
				if !funcDef.IsImplemented() {
					continue
				}
				cfg.Visitor.EnterNamespace(funcDef.Name)
				// BREAKPOINT usage:: funcDef.Name == "configurationCrowdsale"
				entryFuncs = append(entryFuncs, &Function{
					Name:       contractName + "::" + funcDef.Name,
					Block:      cfg._constructFuncLevelBlock(funcDef),
					Parameters: cfg._findFuncLevelParameters(funcDef),
					SrcID:      node.ID,
				})
				cfg.Visitor.ExitNamespace()
			}

		}

	}
	cfg.Visitor.ExitNamespace()

	return entryFuncs
}

func (cfg *CFG) _findFuncLevelParameters(funcDef *AST.FunctionDefinition) []*ST.Symbol {
	var parameters []*ST.Symbol
	for _, param := range funcDef.Parameters.Parameters {
		parameters = append(parameters, &ST.Symbol{
			Namespace:  *cfg.Visitor.CurrentNamespace,
			Identifier: param.Name,
			Type:       ST.LocalVariable,
		})
	}
	return parameters
}

func (cfg *CFG) _constructFuncLevelBlock(funcDef *AST.FunctionDefinition) *Block {
	block := &Block{
		ID:        funcDef.ID,
		Namespace: *cfg.Visitor.CurrentNamespace,
	}

	for _, stmt := range funcDef.Body.Statements {
		block.Statements = append(block.Statements, cfg._constructStatement(stmt))
	}

	return block
}
