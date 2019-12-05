package main

import (
	"github.com/bkaraceylan/anidea/distance"
	"github.com/bkaraceylan/anidea/sequence"
	"github.com/bkaraceylan/anidea/tree"
)

func main() {
	//var config *conf.Config
	//config = conf.LoadConfig()

	samples := sequence.ReadDNA("./H1N1_PB2.fasta", "FASTA", true)

	// fmt.Printf("%v \n", len(samples.Samples))
	// for _, sample := range samples.Samples {
	// 	fmt.Printf("%v \n", sample.Id)
	// }
	//samples = seq.Align(samples, "CW", config)
	result := distance.PairDistConc(samples, "K80")
	//distance.PrettyPrintDist(result)
	tr := tree.NJ(result)
	//tr.DumpTable()
	tr.Print()
	//fmt.Printf("%s\n", tr.Newick())
	//distance.PrettyPrintDist(result)
	// project, _ := proj.ParseProject("./test.yaml")

	// project.RunProject()
}
