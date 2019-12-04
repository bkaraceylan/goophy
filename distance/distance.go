package distance

import (
	"log"
	"math"
	"sync"

	seq "github.com/bkaraceylan/anidea/sequence"
)

type DistMat struct {
	Ids    []string
	Matrix [][]float64
	Method string
}

func PDist(dna1 seq.DNA, dna2 seq.DNA) float64 {
	seq1 := dna1.Seq
	seq2 := dna2.Seq

	if len(seq1) != len(seq2) {
		log.Fatalf("Varying sequence lengths %v != %v \n", len(seq1), len(seq2))
	}

	numnuc := 0
	numdiff := 0

	for x := 0; x < len(seq1); x++ {
		if !seq.IsNuc(rune(seq1[x])) || !seq.IsNuc(rune(seq2[x])) {
			continue
		}

		if rune(seq1[x]) != rune(seq2[x]) {
			numdiff++
		}

		numnuc++
	}

	return (float64(numdiff) / float64(numnuc))
}

func JCDist(dna1 seq.DNA, dna2 seq.DNA) float64 {
	p_dist := PDist(dna1, dna2)
	pow := math.Log(float64(1 - (1.3333 * p_dist)))
	dist := -0.75 * pow

	return dist
}

func K80Dist(dna1 seq.DNA, dna2 seq.DNA) float64 {
	len, _, ts, tv := seq.ComputeTrans(dna1, dna2)

	P := float64(ts) / float64(len)
	Q := float64(tv) / float64(len)

	a1 := 1 - 2*P - Q
	a2 := 1 - 2*Q

	dist := -0.5 * math.Log(a1*math.Sqrt(a2))

	return dist
}

func PairDist(pool seq.DNAPool, method string) DistMat {
	var result [][]float64
	dmat := DistMat{}

	for _, dna1 := range pool.Samples {
		var row []float64

		for _, dna2 := range pool.Samples {
			var dist float64
			switch method {
			case "P":
				dist = PDist(dna1, dna2)
			case "JC":
				dist = JCDist(dna1, dna2)
			case "K80":
				dist = K80Dist(dna1, dna2)
			}

			row = append(row, dist)
		}

		result = append(result, row)
	}

	for _, v := range pool.Samples {
		dmat.Ids = append(dmat.Ids, v.Id)
	}

	dmat.Matrix = result

	return dmat
}

type seqJob struct {
	addr  *DistMat
	dna1  seq.DNA
	dna2  seq.DNA
	idx1  int
	idx2  int
	model string
}

func PairDistConc(pool seq.DNAPool, model string) DistMat {
	dmat := DistMat{}
	dmat.Matrix = make([][]float64, len(pool.Samples))
	chann := make(chan seqJob, 100)

	go func() {
		defer close(chann)
		for idx1, dna1 := range pool.Samples {
			dmat.Matrix[idx1] = make([]float64, len(pool.Samples))

			for idx2, dna2 := range pool.Samples {
				chann <- seqJob{&dmat, dna1, dna2, idx1, idx2, model}
			}
		}
	}()

	go func() {
		for _, v := range pool.Samples {
			dmat.Ids = append(dmat.Ids, v.Id)
		}
	}()

	var wg sync.WaitGroup

	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go PairDistWorker(chann, &wg)
	}

	wg.Wait()
	return dmat
}

func PairDistWorker(seqjob <-chan seqJob, wg *sync.WaitGroup) {
	var dist float64
	defer wg.Done()

	for j := range seqjob {
		if j.idx1 == j.idx2 {
			j.addr.Matrix[j.idx1][j.idx2] = 0
		} else {
			switch j.model {
			case "P":
				dist = PDist(j.dna1, j.dna2)
			case "JC":
				dist = JCDist(j.dna1, j.dna2)
			case "K80":
				dist = K80Dist(j.dna1, j.dna2)
			}

			if dist < 0 || dist == 0 {
				j.addr.Matrix[j.idx1][j.idx2] = 0
			} else {
				j.addr.Matrix[j.idx1][j.idx2] = dist
			}
		}
	}
}
