package printer

import (
	"fmt"
	"txtracker/pkg/common/models"
)

type Printer struct {
	Print (*models.Common)
}

type ASTPrinter struct {
	Root *models.Common
}

func NewASTPrinter(root *models.Common) *ASTPrinter {
	return &ASTPrinter{Root: root}
}

func (a *ASTPrinter) PrintAST() {
	fmt.Print("Printing AST\n")

	printAST(a.Root, 0)
}

func printAST(node *models.Common, spaces int) {
	if node == nil {
		return
	}
	printNode(node)
	childs := node.Children
	for _, child := range childs {
		printAST(child, len(child.Parent.NodeType)+spaces+1)
	}
	fmt.Printf("\n")

}

func printNode(node *models.Common) {
	if node == nil {
		return
	}
	fmt.Printf(" <%s>--> ", node.NodeType)
}
