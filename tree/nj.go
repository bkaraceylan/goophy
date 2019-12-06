package tree

import (
	"fmt"

	dist "github.com/bkaraceylan/goophy/distance"
)

//NJ creates an evolutionary tree from a distance matrix using neighbor-joining algorithm.
func NJ(distmat dist.DistMat) *Tree {
	var intnodes []*Node

	tree := CreateTree("test")
	for _, lbl := range distmat.Ids {
		tree.AddNode(lbl, -1, "root")
	}

	for i := len(distmat.Matrix); i > 2; i-- {
		_, minqcoord, u1, u2 := minq(distmat)

		v1 := 0.5*distmat.Matrix[minqcoord[0]][minqcoord[1]] + 0.5*(u1-u2)
		v2 := 0.5*distmat.Matrix[minqcoord[0]][minqcoord[1]] + 0.5*(u2-u1)

		l1 := distmat.Ids[minqcoord[0]]
		l2 := distmat.Ids[minqcoord[1]]
		intlbl := fmt.Sprintf("int_%d", len(intnodes))

		node := tree.AddNode(intlbl, -1, "")
		intnodes = append(intnodes, node)

		node2 := tree.GetNode(l1)
		node2.SetLength(v1)
		node3 := tree.GetNode(l2)
		node3.SetLength(v2)

		tree.AddChild(node, node2)
		tree.AddChild(node, node3)

		recalMatrix(&distmat, minqcoord[0], minqcoord[1], intlbl)
	}

	l1 := distmat.Ids[0]
	l2 := distmat.Ids[1]
	v1 := sumArray(distmat.Matrix[0]) / 2
	v2 := sumArray(distmat.Matrix[1]) / 2
	node2 := tree.GetNode(l1)
	node2.SetLength(v1)
	node3 := tree.GetNode(l2)
	node3.SetLength(v2)

	root := tree.GetNode("root")
	tree.AdoptChildren(root, node2)
	tree.RemoveNode(node2)
	tree.AddChild(root, node3)

	return tree
}

//recalMatrix recalculates the distance matrix.
func recalMatrix(distmat *dist.DistMat, i int, j int, lbl string) {
	distmat.Ids = append(distmat.Ids, lbl)
	arr := make([]float64, len(distmat.Matrix))
	distmat.Matrix = append(distmat.Matrix, arr)

	for k := 0; k < len(distmat.Matrix); k++ {
		distmat.Matrix[k] = append(distmat.Matrix[k], 0.0)
	}

	dist1 := distmat.Matrix[i][j]

	for k := 0; k < len(distmat.Matrix); k++ {
		if k == len(distmat.Matrix)-1 {
			continue
		}
		dist2 := distmat.Matrix[i][k]
		dist3 := distmat.Matrix[j][k]
		newdist := (dist2 + dist3 - dist1) / 2
		distmat.Matrix[len(distmat.Matrix)-1][k] = newdist
		distmat.Matrix[k][len(distmat.Matrix)-1] = newdist
	}

	distmat.Ids = removeLabel(distmat.Ids, i)
	distmat.Ids = removeLabel(distmat.Ids, j)

	distmat.Matrix = removeRow(distmat.Matrix, i)
	distmat.Matrix = removeRow(distmat.Matrix, j)
	for k := 0; k < len(distmat.Matrix); k++ {
		distmat.Matrix[k] = removeCol(distmat.Matrix[k], i)
		distmat.Matrix[k] = removeCol(distmat.Matrix[k], j)
	}
}

func removeRow(s [][]float64, index int) [][]float64 {
	return append(s[:index], s[index+1:]...)
}

func removeCol(s []float64, index int) []float64 {
	return append(s[:index], s[index+1:]...)
}

func removeLabel(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//minq calculates the qmatrix end returns the minimum q value, coords of the q value, u1 and u2 values for the minimumq
func minq(distmat dist.DistMat) (float64, [2]int, float64, float64) {
	matrix := distmat.Matrix
	qmat := make([][]float64, len(matrix))

	var minq float64
	var us1 float64
	var us2 float64
	var minqcoord [2]int
	for i := 0; i < len(qmat); i++ {

		qmat[i] = make([]float64, len(matrix))

		// j < len(qmat) old
		for j := 0; j < i; j++ {
			u1 := sumArray(matrix[i]) / float64((len(qmat) - 2))
			u2 := sumArray(matrix[j]) / float64((len(qmat) - 2))
			q := matrix[i][j] - u1 - u2
			qmat[i][j] = q

			if q < minq {
				minq = q
				us1 = u1
				us2 = u2
				minqcoord[0] = i
				minqcoord[1] = j
			}
		}
	}

	return minq, minqcoord, us1, us2
}

//sumArray returns the sum of values in an array.
func sumArray(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

//minArray returns the minimum value in an array.
func minArray(nums []float64) float64 {
	min := 0.0
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
