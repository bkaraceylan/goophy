package tree

import (
	"bytes"
	"fmt"
	"strings"
)

//Node is a structre for a single node
type Node struct {
	Label    string
	Length   float64
	Children []*Node
	Parent   *Node
}

//IsLeaf returns true if the node doesn't have any children.
func (node Node) IsLeaf() bool {
	if len(node.Children) == 0 {
		return true
	}

	return false
}

//SetLength sets the branch length of the node.
func (node *Node) SetLength(length float64) {
	node.Length = length
}

//Newick returns a newick representations of the node.
func (node *Node) Newick(newick *bytes.Buffer) {

	if len(node.Children) == 0 {
		newick.WriteString(node.Label)
	} else {
		newick.WriteString("(")
		for idx, child := range node.Children {
			child.Newick(newick)
			if idx < len(node.Children)-1 {
				newick.WriteString(",")
			}
		}

		newick.WriteString(")")
	}

	if node.Length > -1 {
		newick.WriteString(fmt.Sprintf(":%f", node.Length))
	}
}

const _cross string = " ├─"
const _corner string = " └─"
const _vertical string = " │ "
const _space string = "   "
const _nolabel string = "─┐"
const _foo string = " ┌─"

//PrintNode recursively prints nodes to stdout
func (node *Node) PrintNode(indent string, isLast bool, isFirst bool) {
	fmt.Printf("%s", indent)

	isRoot := (node.Label == "root")
	isInternal := strings.HasPrefix(node.Label, "int_")
	//isParentRoot := (node.Parent != nil)

	if isLast {
		if !isRoot {
			fmt.Printf("%s", _corner)
			indent += _space
		}
	} else {
		fmt.Printf("%s", _cross)
		indent += _vertical
	}

	if !isRoot {

		if isInternal {
			fmt.Printf("%s\n", _nolabel)
		} else {
			fmt.Printf("%s\n", node.Label)
		}
	}

	numChild := len(node.Children)
	for x := 0; x < numChild; x++ {
		child := node.Children[x]
		isLast := (x == (numChild - 1))
		isFirst := (x == 0)
		child.PrintNode(indent, isLast, isFirst)
	}

}
