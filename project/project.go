package project

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	dist "github.com/bkaraceylan/goophy/distance"
	seq "github.com/bkaraceylan/goophy/sequence"
	"github.com/bkaraceylan/goophy/tree"
	yaml "gopkg.in/yaml.v2"
)

type Input struct {
	Format  string `yaml: "format"`
	Aligned bool   `yaml: "aligned"`
	File    string `yaml: "file"`
}

type Distance struct {
	Method string `yaml: "method"`
	Output string `yaml: "output"`
}

type Project struct {
	Name      string       `yaml: "name"`
	Inputs    []Input      `yaml: "inputs"`
	Distances []Distance   `yaml: "distances"`
	Trees     []tree.Input `yaml: "trees"`
	Directory string
	Filepath  string
	Pools     []seq.DNAPool
	DistMats  []dist.DistMat
	//Trees     []tree.Tree
}

func ParseProject(file string) (Project, error) {
	var project Project
	projFile, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return project, err
	}

	err = yaml.Unmarshal(projFile, &project)

	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		return project, err
	}
	project.Filepath = filepath.Dir(file)

	return project, nil
}

func (project *Project) RunProject() {
	fmt.Printf("Beginning project %s\n", project.Name)
	currentTime := time.Now()
	dir := project.Name + "_" + currentTime.Format("01-02-2006")
	dir = filepath.Join(project.Filepath, dir)
	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		fmt.Printf("Error creating project directory %s\n", err)
		return
	}

	project.Directory = dir
	project.RunInput()
	project.RunDist()
	project.RunTree()

	fmt.Printf("Finished project %s\n", project.Name)
}

func (project *Project) RunInput() {
	for _, input := range project.Inputs {
		path := filepath.Join(project.Filepath, input.File)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Printf("Input file %s does not exist.\n", input.File)
			return
		}
		fmt.Printf("Parsing input file %s\n", input.File)
		pool := seq.ReadDNA(path, input.Format, input.Aligned)
		project.Pools = append(project.Pools, pool)
	}

}

func (project *Project) RunDist() {

	var path string
	for _, dista := range project.Distances {
		if dista.Output != "" {
			path = filepath.Join(project.Directory, "Distances")
			err := os.MkdirAll(path, os.ModePerm)

			if err != nil {
				log.Fatalf("Error creating distances directory %s\n", err)
			}
			break
		}
	}

	for _, pool := range project.Pools {
		for _, dista := range project.Distances {
			fmt.Printf("Calculating pairwise %s distances in %s pool.\n", dista.Method, pool.Name)
			result := dist.PairDistConc(pool, dista.Method)
			project.DistMats = append(project.DistMats, result)
			//dist.PrettyPrintDist(result)

			if dista.Output == "TXT" {
				name := dista.Method + "_" + pool.Name
				dist.PrettySaveDist(result, path, name)
			} else {
				dist.PrettyPrintDist(result)
			}

		}
	}
}

func (project *Project) RunTree() {
	var path string
	for _, dista := range project.Distances {
		if dista.Output != "" {
			path = filepath.Join(project.Directory, "Trees")
			err := os.MkdirAll(path, os.ModePerm)

			if err != nil {
				log.Fatalf("Error creating distances directory %s\n", err)
			}
			break
		}
	}

	for _, distmat := range project.DistMats {
		for _, phylo := range project.Trees {
			fmt.Printf("Calculating phylogenetic trees for %s distance matrix of %s using %s method.\n", distmat.Method, distmat.Alignment.Name, phylo.Method)
			tr := tree.TreeBuilder(phylo, distmat)

			if phylo.Output != "" {
				name := phylo.Method + "_" + distmat.Method + "_" + distmat.Alignment.Name
				tree.Save(tr, path, name, phylo.Output)
			} else {
				tr.Print()
			}
		}
	}
}
