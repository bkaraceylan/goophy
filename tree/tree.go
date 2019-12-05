package tree

import (
	"bytes"
	"fmt"
)

//Tree is a pylogenetic tree structure
type Tree struct {
	Name      string
	NodeTable map[string]*Node
	Root      *Node
}

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

//DumpTable prints the NodeTable to stdout (Ugly)
func (tr *Tree) DumpTable() {
	fmt.Print("Label			Parent				Length				Children\n")
	for _, node := range tr.NodeTable {
		label := node.Label
		length := node.Length
		parent := "*"
		if node.Parent != nil {
			parent = node.Parent.Label
		}

		children := ""

		for _, child := range node.Children {
			children += child.Label + ", "
		}
		isRoot := ""
		if node == tr.Root {
			isRoot = "*"
		}
		fmt.Printf("%s%s			%s				%v				%s\n", label, isRoot, parent, length, children)
	}
}

//AddNode adds a node to a tree
func (tr *Tree) AddNode(label string, length float64, parentLabel string) *Node {

	_, ok := tr.NodeTable[label]
	if ok {
		fmt.Printf("A node with the label=%s already exists!!\n", label)
		return nil
	}

	var parent *Node

	if parentLabel != "" {
		parent, ok = tr.NodeTable[parentLabel]

		if !ok {
			fmt.Print("Parent not found!")
			return nil
		}

	} else {
		parent = nil
	}

	node := &Node{Label: label, Length: length}

	if parent != nil {
		tr.AddChild(parent, node)
	}

	tr.NodeTable[label] = node

	return node
}

//AddChild adds a single child to a parent node
func (tr *Tree) AddChild(parent *Node, newchild *Node) {
	//If node already has parent, remove it from the parent's children
	if newchild.Parent != nil {
		tr.RemoveChild(newchild.Parent, newchild)
	}

	newchild.Parent = parent
	parent.Children = append(parent.Children, newchild)
}

//RemoveChild removes a single child from a parent node
func (tr *Tree) RemoveChild(parent *Node, child *Node) {
	var newchildren []*Node
	child.Parent = nil
	for _, node := range parent.Children {
		if node != child {
			newchildren = append(newchildren, node)
		}
	}

	parent.Children = newchildren
}

//GetChildren returns all children of a node
func (tr *Tree) GetChildren(label string) []*Node {
	node := tr.NodeTable[label]
	return node.Children
}

//GetNode returns a pointer to the specified node from the hashtable
func (tr *Tree) GetNode(label string) *Node {
	return tr.NodeTable[label]
}

//IsRooted returns true if the root node has two children.
func (tr *Tree) IsRooted() bool {
	if len(tr.Root.Children) == 2 {
		return true
	}

	return false
}

//Print method prints a phylogenetic tree to stdout. Currently ignodes branch lengths.
func (tr *Tree) Print() {
	if tr.Root == nil {
		fmt.Printf("show: root node not found\n")
		return
	}

	// for _, node := range tr.Root.Children {
	// 	PrintNode(node, "", true)
	// }
	tr.Root.PrintNode("", true, false)
}

//NewRoot sets the root of the tree to the specified node. Parents aren't removed from the hashtable.
func (tr *Tree) NewRoot(node *Node) {
	tr.Root = node
}

//AdoptChildren transfers all children of the transferring node to the receiving node.
func (tr *Tree) AdoptChildren(rnode *Node, tnode *Node) {
	rnode.Children = append(rnode.Children, tnode.Children...)
	tnode.Children = nil
	for _, child := range tnode.Children {
		child.Parent = rnode
	}
}

//RemoveNode removes a node from the hashtable and from the children list of the parent node.
func (tr *Tree) RemoveNode(node *Node) {

	if len(node.Children) > 0 {
		fmt.Println("Cannot remove a node with children")
		return
	}

	if node.Parent != nil {

		var tmp []*Node
		for _, child := range node.Parent.Children {
			if child != node {
				tmp = append(tmp, child)
			}
		}

		node.Parent.Children = tmp
	}

	delete(tr.NodeTable, node.Label)
}

//Newick returns a newick representations of a tree.
func (tr *Tree) Newick() string {
	var newick bytes.Buffer
	tr.Root.Newick(&newick)
	newick.WriteString(";")

	return newick.String()
}
