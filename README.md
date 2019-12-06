# Goophy

Goophy is a simple phylogenetic analysis tool written in golang.


## Usage

```
./goophy ./data/test.yaml
```

Goophy takes a [project](./data/test.yaml) file in yaml format as input.

A project consists of three lists: inputs, distances and trees.

Currently only a limited set of formats and analyzes are supported.

* Inputs: Only DNA sequences in FASTA format.
* Distances: Jukes-Cantor and Kimura-2-parameter distances (w/o gamma distribution). Can output pretty tables in ascii to STDOUT or a file.
* Trees: Only Neighbor-Joining algorithm is supported. Can print to STDOUT in ascii or write to a file in NEWICK format.

### Building

Clone the repo, cd into the directory and compile with:
```
go build main.go
```

Or to run without building:

```
go run main.go
```
* **Burak Karaceyln** - *Initial work* - [bkaraceylan](https://github.com/bkaraceylan)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
