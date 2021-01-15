package bench

import (
	"testing"

	cespareV1 "github.com/cespare/xxhash"
	cespareV2 "github.com/cespare/xxhash/v2"
	dgryski "github.com/dgryski/go-wyhash"
	orisanoV2 "github.com/orisano/wyhash"
	orisanoV3 "github.com/orisano/wyhash/v3"
	orisanoV4 "github.com/orisano/wyhash/v4"
	"github.com/shenwei356/bio/seq"
	"github.com/shenwei356/util/bytesize"
	"github.com/spaolacci/murmur3"
	"github.com/will-rowe/nthash"
	zeebo "github.com/zeebo/wyhash"
	"github.com/zeebo/xxh3"
)

const k = 31

func BenchmarkXxhashV3CespareV1(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, cespareV1.Sum64(_seq.Seq))
				}
			}
		})
	}
}

func BenchmarkXxhashV3CespareV2(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, cespareV2.Sum64(_seq.Seq))
				}
			}
		})
	}
}
func BenchmarkXxhashV3Zeebo(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, xxh3.Hash(_seq.Seq))
				}
			}
		})
	}
}

func BenchmarkMurmur3(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)
			hash := murmur3.New64()

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					hash.Reset()
					hash.Write(_seq.Seq)
					codes = append(codes, hash.Sum64())
				}
			}
		})
	}
}

func BenchmarkWyhashV3Dgryski(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, dgryski.Hash(_seq.Seq, 1))
				}
			}
		})
	}
}

func BenchmarkWyhashV3Zeebo(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, zeebo.Hash(_seq.Seq, 1))
				}
			}
		})
	}
}

func BenchmarkWyhashV2Orisano(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, orisanoV2.Sum64(1, _seq.Seq))
				}
			}
		})
	}
}

func BenchmarkWyhashV3Orisano(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, orisanoV3.Sum64(1, _seq.Seq))
				}
			}
		})
	}
}

func BenchmarkWyhashV4Orisano(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var slider func() (*seq.Seq, bool)
			var _seq *seq.Seq
			var _ok bool
			codes := make([]uint64, 0, size)

			for j := 0; j < b.N; j++ {
				slider = benchSeqs[i].Slider(k, 1, false, false)
				codes = codes[:0]
				for {
					_seq, _ok = slider()
					if !_ok {
						break
					}

					codes = append(codes, orisanoV4.Sum64(1, _seq.Seq))
				}
			}
		})
	}
}

func BenchmarkNthash(b *testing.B) {
	for i := range benchSeqs {
		size := len(benchSeqs[i].Seq)

		b.Run(bytesize.ByteSize(size).String(), func(b *testing.B) {
			var _code uint64
			var _ok bool
			codes := make([]uint64, 0, size)
			var err error
			var hasher *nthash.NTHi

			for j := 0; j < b.N; j++ {

				hasher, err = nthash.NewHasher(&benchSeqs[i].Seq, uint(k))
				if err != nil {
					b.Errorf("fail to create ntHash iterator. seq length: %d", size)
				}

				codes = codes[:0]
				for {
					_code, _ok = hasher.Next(false)
					if !_ok {
						break
					}

					codes = append(codes, _code)
				}
			}
		})
	}
}
