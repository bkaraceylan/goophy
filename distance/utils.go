package distance

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

//PrettyPrintDist prints a distance matrix to the stdout in a nice table format.
func PrettyPrintDist(dmat DistMat) {
	var table *tablewriter.Table

	table = tablewriter.NewWriter(os.Stdout)

	table.SetAutoFormatHeaders(false)
	table.SetHeader(append([]string{""}, dmat.Ids...))

	for idx, v := range dmat.Matrix {
		var strrow []string
		strrow = append(strrow, dmat.Ids[idx])

		for _, v2 := range v {
			strrow = append(strrow, strconv.FormatFloat(v2, 'f', 6, 32))
		}

		table.Append(strrow)
	}

	table.Render()
}

//PrettySaveDist saves a distance matrix in the disk in a nice table format.
func PrettySaveDist(dmat DistMat, path string, name string) {
	var table *tablewriter.Table
	path = filepath.Join(path, name+".txt")
	fmt.Println(path)
	file, _ := os.Create(path)
	defer file.Close()
	table = tablewriter.NewWriter(file)

	table.SetAutoFormatHeaders(false)
	table.SetHeader(append([]string{""}, dmat.Ids...))

	for idx, v := range dmat.Matrix {
		var strrow []string
		strrow = append(strrow, dmat.Ids[idx])

		for _, v2 := range v {
			strrow = append(strrow, strconv.FormatFloat(v2, 'f', 6, 32))
		}

		table.Append(strrow)
	}

	table.Render()
}
