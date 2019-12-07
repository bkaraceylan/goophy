# Goophy

Goophy is a simple phylogenetic analysis tool written in golang.


## Usage

```
./goophy ./data/test.yaml
```

Goophy takes a [project](./data/test.yaml) file in yaml format as input, iteratively performs analyzes specified in the project and produces outputs.

A project consists of three lists: inputs, distances and trees.

All distance analyses in the project are performed on all inputs in the project, and all tree building methods in the project are performed using the results of all distance analyses.

A new folder with the project name and current date as suffix is created in the same directory as the project file. If outputs are specified in the project file, results of distance analyses are saved into the "Distance" folder and trees are saved into the "Trees" folder.

Currently only a limited set of formats and analyses are supported.

* **Inputs**: Only DNA sequences in FASTA format.
* **Distances**: Jukes-Cantor and Kimura-2-parameter distances (w/o gamma distribution). Can output pretty tables in ascii to STDOUT or a file.
* **Trees**: Only Neighbor-Joining algorithm is supported. Can print to STDOUT in ascii or write to a file in NEWICK format.

### Building

Clone the repo, cd into the directory and compile with:
```
go build main.go
```

Or to run without building:

```
go run main.go
```
## TODO

* Replace quick fixes with real fixes.
* Tidy up the code.
* Redo things in smarter ways.
* Add more analyzes.

## Authors

* **Burak Karaceylan** - [bkaraceylan](https://github.com/bkaraceylan)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
