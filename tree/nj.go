package tree

import (
	"fmt"

	dist "github.com/bkaraceylan/anidea/distance"
)

func NJ(distmat dist.DistMat) *Tree {
	var intnodes []*Node
	tree := CreateTree("test")
	for _, lbl := range distmat.Ids {
		tree.Add(lbl, 0, "root")
	}

	for i := len(distmat.Matrix); i > 2; i-- {
		_, minqcoord, u1, u2 := minq(distmat)

		v1 := 0.5*distmat.Matrix[minqcoord[0]][minqcoord[1]] + 0.5*(u1-u2)
		v2 := 0.5*distmat.Matrix[minqcoord[0]][minqcoord[1]] + 0.5*(u2-u1)

		l1 := distmat.Ids[minqcoord[0]]
		l2 := distmat.Ids[minqcoord[1]]
		//fmt.Printf("%v - %v\n", l1, l2)
		//fmt.Printf("%v\n", v1+v2)
		intlbl := fmt.Sprintf("int_%d", len(intnodes))

		node := tree.Add(intlbl, 0, "root")
		intnodes = append(intnodes, node)
		tree.Add(l1, v1, intlbl)
		tree.Add(l2, v2, intlbl)

		recalMatrix(&distmat, minqcoord[0], minqcoord[1], intlbl)
		dist.PrettyPrintDist(distmat)
	}

	return tree
}

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

	distmat.Ids = RemoveLabel(distmat.Ids, i)
	distmat.Ids = RemoveLabel(distmat.Ids, j)

	distmat.Matrix = RemoveRow(distmat.Matrix, i)
	distmat.Matrix = RemoveRow(distmat.Matrix, j)
	for k := 0; k < len(distmat.Matrix); k++ {
		distmat.Matrix[k] = RemoveCol(distmat.Matrix[k], i)
		distmat.Matrix[k] = RemoveCol(distmat.Matrix[k], j)
	}
}

func RemoveRow(s [][]float64, index int) [][]float64 {
	return append(s[:index], s[index+1:]...)
}

func RemoveCol(s []float64, index int) []float64 {
	return append(s[:index], s[index+1:]...)
}

func RemoveLabel(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//Calculates the qmatrix end returns the minimum q value, coords of the q value, u1 and u2 values for the minimumq
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
	distmat.Matrix = qmat
	dist.PrettyPrintDist(distmat)
	return minq, minqcoord, us1, us2
}

func sumArray(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum = sum + num
	}

	return sum
}

func minArray(nums []float64) float64 {
	min := 0.0
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
