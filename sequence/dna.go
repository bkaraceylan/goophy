package sequence

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"unicode"

	conf "github.com/bkaraceylan/anidea/config"
)

type DNA struct {
	Id  string
	Seq string
}

type DNAPool struct {
	Name        string
	Samples     []DNA
	File        string
	Format      string
	Aligned     bool
	AlignMethod string
}

//Returns total number of bases
func (dna DNA) NumBases() int {
	return len(dna.Seq)
}

//Returns total number of purines
func (dna DNA) NumPurines() int {
	purines := 0

	for _, base := range dna.Seq {
		if isPurine(base) {
			purines += 1
		}
	}

	return purines
}

//Returns total number of pyrimidines
func (dna DNA) NumPyrimidines() int {
	pyrmidine := 0

	for _, base := range dna.Seq {
		if isPyrimidine(base) {
			pyrmidine += 1
		}
	}

	return pyrmidine
}

//Returns true if not one of ATCG
func IsNuc(base rune) bool {
	bases := "ATGC"
	return strings.ContainsRune(bases, unicode.ToUpper(base))
}

//Returns true if the base is purine
func isPurine(base rune) bool {
	if base == 'A' || base == 'G' {
		return true
	} else {
		return false
	}
}

//Returns true if the base is pyrimidine
func isPyrimidine(base rune) bool {
	if base == 'C' || base == 'T' {
		return true
	} else {
		return false
	}
}

//Returns true if the change between bases is a transition
func IsTransition(base1 rune, base2 rune) bool {
	if (isPurine(base1) && isPurine(base2)) || (isPyrimidine(base1) && isPyrimidine(base2)) {
		return true
	} else {
		return false
	}
}

//Returns true if the change between bases is a transversion
func IsTransversion(base1 rune, base2 rune) bool {
	if (isPurine(base1) && isPyrimidine(base2)) || (isPyrimidine(base1) && isPurine(base2)) {
		return true
	} else {
		return false
	}
}

//Returns number of nucleotides, number of differences, number of transitions and number of transversions
func ComputeTrans(dna1 DNA, dna2 DNA) (int, int, int, int) {
	seq1 := dna1.Seq
	seq2 := dna2.Seq

	if len(seq1) != len(seq2) {
		log.Fatalf("Varying sequence lengths %v != %v \n", len(seq1), len(seq2))
	}

	numnuc := 0
	ts := 0
	numdiff := 0

	for x := 0; x < len(seq1); x++ {
		if !IsNuc(rune(seq1[x])) || !IsNuc(rune(seq2[x])) {
			continue
		}

		if rune(seq1[x]) != rune(seq2[x]) {
			numdiff++

			if IsTransition(rune(seq1[x]), rune(seq2[x])) {
				ts++
			}
		}

		numnuc++
	}

	tv := numdiff - ts

	return numnuc, numdiff, ts, tv
}

//Given a DNAPool returns an aligned DNAPool
func Align(pool DNAPool, method string, conf *conf.Config) DNAPool {
	var aligned DNAPool
	switch method {
	case "CW":
		clustalw(pool.Name, conf.ClustalW)
		aligned = Fasta(pool.Name)
		aligned.Aligned = true
		aligned.AlignMethod = "CW"
	}

	return aligned
}

//Executes clustalw on a file
func clustalw(file string, exe string) bool {
	infile := fmt.Sprintf("-infile=%s", file)

	cw := exec.Command(exe, infile, "-seqnos=ON -gapopen=2 -gapext=0.5 -output=FASTA")
	_, err := cw.Output()
	if err != nil {
		panic(err)
	}

	return true
}
