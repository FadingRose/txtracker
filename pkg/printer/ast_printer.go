package printer

import (
	"fmt"
	"strings"
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

	printerHelper(a.Root, 0, "")
}

// printerHelper is a recursive helper function that prints the tree structure.
// It takes a node, the current depth, and the prefix string for indentation.
func printerHelper(node *models.Common, depth int, prefix string) {
	if node == nil {
		return
	}

	// Calculate indentation based on depth
	indent := strings.Repeat(" ", depth*4) // 4 spaces per depth level
	if depth > 0 {
		fmt.Println(prefix + "-> " + node.NodeType)
	} else {
		fmt.Println(node.NodeType)
	}

	// If the node has children, recursively call this function for each child
	newPrefix := indent + strings.Repeat(" ", len(node.NodeType)+3)
	for _, child := range node.Children {
		printerHelper(child, depth+1, newPrefix)
	}
}
