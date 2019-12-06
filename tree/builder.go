package tree

import (
	"fmt"

	"github.com/bkaraceylan/anidea/distance"
)

type Input struct {
	Method string `yaml: "method"`
	Output string `yaml: "output"`
}

func TreeBuilder(input Input, distmat distance.DistMat) *Tree {
	var tree *Tree
	switch input.Method {
	case "NJ":
		tree = NJ(distmat)
	default:
		fmt.Printf("Unknown tree building algorithm: %s\n", input.Method)
		return nil
	}

	return tree
}
