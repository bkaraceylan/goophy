package tree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func Save(tree *Tree, path string, name string, format string) {
	var contents string
	var ext string
	switch format {
	case "NEWICK":
		contents = tree.Newick()
		ext = ".newick"
	default:
		fmt.Printf("Unknown tree output format %s\n", format)
	}

	fullpath := filepath.Join(path, name+ext)
	err := ioutil.WriteFile(fullpath, []byte(contents), 0644)

	if err != nil {
		fmt.Printf("Error writing to file %s\n", fullpath)
	}
}
