package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bkaraceylan/anidea/project"
)

func main() {
	//var config *conf.Config
	//config = conf.LoadConfig()

	//samples := sequence.ReadDNA("./data/H1N1_PB2.fasta", "FASTA", true)

	// fmt.Printf("%v \n", len(samples.Samples))
	// for _, sample := range samples.Samples {
	// 	fmt.Printf("%v \n", sample.Id)
	// }
	//samples = seq.Align(samples, "CW", config)
	//result := distance.PairDistConc(samples, "K80")
	//distance.PrettyPrintDist(result)
	//tr := tree.NJ(result)
	//tr.DumpTable()
	//tr.Print()
	//fmt.Printf("%s\n", tr.Newick())
	//distance.PrettyPrintDist(result)
	// project, _ := proj.ParseProject("./test.yaml")

	// project.RunProject()

	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("Please specify a project file.\n")
		os.Exit(0)
	}

	proj, err := project.ParseProject(flag.Arg(0))

	if err != nil {
		log.Fatal(err)
	}

	proj.RunProject()
}
