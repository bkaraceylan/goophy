package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	dist "github.com/bkaraceylan/anidea/distance"
	seq "github.com/bkaraceylan/anidea/sequence"
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
	Name      string     `yaml: "name"`
	Inputs    []Input    `yaml: "inputs"`
	Distances []Distance `yaml: "distances"`
	Directory string
	Pools     []seq.DNAPool
	DistMats  []dist.DistMat
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

	return project, nil
}

func (project *Project) RunProject() {
	fmt.Printf("Beginning project %s\n", project.Name)
	currentTime := time.Now()
	dir := project.Name + "_" + currentTime.Format("01-02-2006")
	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		fmt.Printf("Error creating project directory %s\n", err)
	}

	project.Directory = dir
	project.RunInput()
	project.RunDist()

	fmt.Printf("Finished project %s\n", project.Name)
}

func (project *Project) RunInput() {
	for _, input := range project.Inputs {
		fmt.Printf("Parsing input file %s\n", input.File)
		pool := seq.ReadDNA(input.File, input.Format, input.Aligned)
		project.Pools = append(project.Pools, pool)
	}

}

func (project *Project) RunDist() {
	path := filepath.Join(project.Directory, "Distances")
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		fmt.Printf("Error creating distances directory %s\n", err)
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
			}

		}
	}
}
