package sequence

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

//ReadDNA reads a sequence file from the disk.
func ReadDNA(path string, format string, aligned bool) DNAPool {
	var dna DNAPool
	switch format {
	case "FASTA":
		dna = Fasta(path)
	}

	if aligned {
		dna.Aligned = true
	}

	return dna
}

//SpaceMap removes whitespaces from a string.
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

//Fasta parses fasta formatted files.
func Fasta(path string) DNAPool {

	var samples []DNA

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequence string
	var id string
	for scanner.Scan() {
		newsample, _ := regexp.MatchString(`^>`, scanner.Text())
		newline, _ := regexp.MatchString(`^[ ACGTNacgtn-]+$`, scanner.Text())

		if newsample {
			if len(sequence) > 0 {

				samples = append(samples, DNA{id, strings.ToUpper(SpaceMap(sequence))})
				sequence = ""
				id = ""
			}

			id = scanner.Text()[1:]
		} else if newline {
			sequence += scanner.Text()
		} else {
			log.Fatal("Invalid fasta file.")
		}

	}
	//Scanner EOF, iast sample
	samples = append(samples, DNA{id, strings.ToUpper(SpaceMap(sequence))})

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file.")
	}

	if len(samples) < 0 {
		log.Fatal("No sequence data in the file.")
	}

	pool := DNAPool{}
	pool.Name = file.Name()
	pool.File, _ = filepath.Abs(filepath.Dir(file.Name()))
	pool.Samples = samples
	pool.Format = "fasta"

	return pool

}
