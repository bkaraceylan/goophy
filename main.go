package main

import (
	tree "github.com/bkaraceylan/anidea/tree"
)

func main() {
	//var config *conf.Config
	//config = conf.LoadConfig()

	//samples := seq.ReadDNA("./H1N1_PB2.fasta", "FASTA", true)

	// fmt.Printf("%v \n", len(samples.Samples))
	// for _, sample := range samples.Samples {
	// 	fmt.Printf("%v \n", sample.Id)
	// }
	//samples = seq.Align(samples, "CW", config)
	//result := dist.PairDistConc(samples, "K80")
	//dist.PrettyPrintDist(result)

	// project, _ := proj.ParseProject("./test.yaml")

	// project.RunProject()

	tree := tree.CreateTree("test")
	tree.Add("lbltest", 2.5, "root")
	tree.Add("lbltest2", 2.5, "root")
	tree.Add("lbltest3", 2.5, "lbltest2")
	tree.Add("lbltest4", 2.5, "lbltest2")
	// tree.Add("lbltest5", 2.5, "root")
	// tree.Add("lbltest6", 2.5, "lbltest5")
	// tree.Add("lbltest7", 2.5, "lbltest5")
	// tree.Add("lbltest8", 2.5, "root")

	//fmt.Printf("Is rooted: %v \n", tree.IsRooted())

	//node := tree.NodeTable["lbltest"]
	//fmt.Printf("%v \n", node.IsLeaf())
	//fmt.Printf("%v \n", node.ParentLabel)
	//fmt.Printf("%v \n", tree.GetChildren("root"))

	tree.Print()

}

//ACTGACTAGCTAGCTAACTG
//GCATCGTAGCTAGCTACGAT
//* ****          ****
//+ ----          ----
