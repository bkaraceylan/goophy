package tree

import "fmt"

//Tree is the structure holding a pylogenetic tree
type Tree struct {
	Name      string
	NodeTable map[string]*Node
	Root      *Node
}

//Node is a single node on a tree
type Node struct {
	Label       string
	Lenght      float64
	Children    []*Node
	ParentLabel string
}

//Add is a method that adds a node to a tree
func (tr *Tree) Add(label string, length float64, parentlabel string) *Node {
	//fmt.Printf("add: label=%v length=%v parentLabel=%v\n", label, length, parentLabel)

	node := &Node{Label: label, Lenght: length, Children: []*Node{}, ParentLabel: parentlabel}

	parent, ok := tr.NodeTable[parentlabel]
	if !ok {
		return nil
	}

	parent.Children = append(parent.Children, node)

	tr.NodeTable[label] = node

	return node
}

//GetChildren returns all children of a node
func (tr *Tree) GetChildren(label string) []*Node {
	node := tr.NodeTable[label]
	return node.Children
}

//IsRooted returns true if the root node has two children.
func (tr *Tree) IsRooted() bool {
	if len(tr.Root.Children) == 2 {
		return true
	}

	return false
}

//Print method prints a phylogenetic tree to stdout.
func (tr *Tree) Print() {
	if tr.Root == nil {
		fmt.Printf("show: root node not found\n")
		return
	}

	// for _, node := range tr.Root.Children {
	// 	PrintNode(node, "", true)
	// }
	PrintNode(tr.Root, "", true, false)
}

const _cross string = " ├─"
const _corner string = " └─"
const _vertical string = " │ "
const _space string = "   "
const _root string = "─┐"
const _foo string = " ┌─"

//PrintNode recursively prints nodes to stdout
func PrintNode(node *Node, indent string, isLast bool, isFirst bool) {
	fmt.Printf("%s", indent)

	isRoot := (node.Label == "root")
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
		fmt.Printf("%s\n", node.Label)
	}

	numChild := len(node.Children)
	for x := 0; x < numChild; x++ {
		child := node.Children[x]
		isLast := (x == (numChild - 1))
		isFirst := (x == 0)
		PrintNode(child, indent, isLast, isFirst)
	}

}

//IsLeaf returns true if the node doesn't have any children
func (node Node) IsLeaf() bool {
	if len(node.Children) == 0 {
		return true
	}

	return false
}

// func (node Node) IsRoot() bool {
// 	node.ParentLabel == "top"
// }

//CreateTree initializes a tree.
func CreateTree(name string) *Tree {
	tree := new(Tree)
	tree.NodeTable = map[string]*Node{}

	rootnode := new(Node)
	rootnode.Label = "root"

	tree.Root = rootnode
	tree.NodeTable["root"] = rootnode
	return tree
}
