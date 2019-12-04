package tree

import "fmt"

type Tree struct {
	Name      string
	NodeTable map[string]*Node
	Root      *Node
}

type Node struct {
	Label       string
	Lenght      float64
	Children    []*Node
	ParentLabel string
}

func (tr *Tree) Add(label string, length float64, parentLabel string) {
	//fmt.Printf("add: label=%v length=%v parentLabel=%v\n", label, length, parentLabel)

	node := &Node{Label: label, Lenght: length, Children: []*Node{}, ParentLabel: parentLabel}

	parent, ok := tr.NodeTable[parentLabel]
	if !ok {
		fmt.Printf("add: parentLabel=%v: not found\n", parentLabel)
		return
	} else {
		parent.Children = append(parent.Children, node)
	}

	tr.NodeTable[label] = node
}

func (tr *Tree) GetChildren(label string) []*Node {
	node := tr.NodeTable[label]
	return node.Children
}

func (tr *Tree) IsRooted() bool {
	if len(tr.Root.Children) == 2 {
		return true
	} else {
		return false
	}
}

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

func PrintNode(node *Node, indent string, isLast bool, isFirst bool) {
	fmt.Printf("%s", indent)

	isRoot := (node.Label == "root")
	isParentRoot := (node.ParentLabel == "root")

	if isLast {
		if !isRoot {
			fmt.Printf("%s", _corner)
			indent += _space
		}
	} else {
		if isParentRoot && isFirst {
			fmt.Printf("%s", _foo)
		} else {
			fmt.Printf("%s", _cross)
		}

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

func (node Node) IsLeaf() bool {
	if len(node.Children) == 0 {
		return true
	} else {
		return false
	}
}

// func (node Node) IsRoot() bool {
// 	node.ParentLabel == "top"
// }

func CreateTree(name string) *Tree {
	tree := new(Tree)
	tree.NodeTable = map[string]*Node{}

	rootnode := new(Node)
	rootnode.Label = "root"

	tree.Root = rootnode
	tree.NodeTable["root"] = rootnode
	return tree
}
