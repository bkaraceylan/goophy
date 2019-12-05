package main

import (
	"github.com/bkaraceylan/anidea/distance"
	"github.com/bkaraceylan/anidea/sequence"
	"github.com/bkaraceylan/anidea/tree"
)

func main() {
	//var config *conf.Config
	//config = conf.LoadConfig()

	samples := sequence.ReadDNA("./woodmouse.fasta", "FASTA", true)

	// fmt.Printf("%v \n", len(samples.Samples))
	// for _, sample := range samples.Samples {
	// 	fmt.Printf("%v \n", sample.Id)
	// }
	//samples = seq.Align(samples, "CW", config)
	result := distance.PairDistConc(samples, "K80")
	//distance.PrettyPrintDist(result)
	tree.NJ(result)
	//dist.PrettyPrintDist(qmat)
	// project, _ := proj.ParseProject("./test.yaml")

	// project.RunProject()

	//tree := tree.CreateTree("test")
	//tree.Add("lbltest", 2.5, "root")
	//tree.Add("lbltest2", 2.5, "root")
	//tree.Add("lbltest3", 2.5, "lbltest2")
	//tree.Add("lbltest4", 2.5, "lbltest2")
	//tree.Add("lbltest5", 2.5, "root")
	//tree.Add("lbltest6", 2.5, "lbltest5")
	//tree.Add("lbltest7", 2.5, "lbltest5")
	//tree.Add("lbltest8", 2.5, "root")

	//fmt.Printf("Is rooted: %v \n", tree.IsRooted())

	//node := tree.NodeTable["lbltest"]
	//fmt.Printf("%v \n", node.IsLeaf())
	//fmt.Printf("%v \n", node.ParentLabel)
	//fmt.Printf("%v \n", tree.GetChildren("root"))

	//tr.Print()

}
