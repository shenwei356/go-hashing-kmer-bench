package bench

import (
	"math/rand"

	"github.com/shenwei356/bio/seq"
)

var benchSeqs []*seq.Seq

var bit2base = [4]byte{'A', 'C', 'G', 'T'}

func init() {
	rand.Seed(11)

	sizes := []int{1 << 10, 1 << 20} // ,10 << 20}
	benchSeqs = make([]*seq.Seq, len(sizes))
	var err error
	for i, size := range sizes {
		sequence := make([]byte, size)

		// fmt.Printf("generating pseudo DNA with lenght of %s ...\n", bytesize.ByteSize(size))
		for j := 0; j < size; j++ {
			sequence[j] = bit2base[rand.Intn(4)]
		}
		benchSeqs[i], err = seq.NewSeq(seq.DNA, sequence)
		if err != nil {
			panic("should not happen")
		}
	}
	// fmt.Printf("%d DNA sequences generated\n", len(sizes))
}
