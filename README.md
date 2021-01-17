# Benchmark of hashing k-mers in Golang

## Hash functions

- ntHash
  - https://github.com/will-rowe/nthash
- xxhash v3
  - cespareV1 https://github.com/cespare/xxhash
  - cespareV2 https://github.com/cespare/xxhash/v2
  - zeebo https://github.com/zeebo/xxh3
- murmur3
  - https://github.com/spaolacci/murmur3
  - https://github.com/twmb/murmur3
- wyhash v2
  - orisano https://github.com/orisano/wyhash/
- wyhash v3
  - dgryski https://github.com/dgryski/go-wyhash
  - orisano https://github.com/orisano/wyhash/v3
  - zeebo https://github.com/zeebo/wyhash
- wyhash v4
  - orisano https://github.com/orisano/wyhash/v4
  
## Result

```
$ go test . -bench=Benchmark*

goos: linux
goarch: amd64
pkg: github.com/shenwei356/hash-bench
cpu: AMD Ryzen 7 2700X Eight-Core Processor         

BenchmarkNthash/1.00_KB-16                   125276        9431 ns/op
BenchmarkWyhashV3Zeebo/1.00_KB-16            10000         115491 ns/op
BenchmarkWyhashV3Orisano/1.00_KB-16          10000         119180 ns/op
BenchmarkXxhashV3Zeebo/1.00_KB-16            10000         119337 ns/op
BenchmarkWyhashV4Orisano/1.00_KB-16          9868          121229 ns/op
BenchmarkXxhashV3CespareV1/1.00_KB-16        9506          125942 ns/op
BenchmarkXxhashV3CespareV2/1.00_KB-16        9098          134062 ns/op
BenchmarkWyhashV3Dgryski/1.00_KB-16          8604          134560 ns/op
BenchmarkMurmur3Twmb/1.00_KB-16              6706          179790 ns/op
BenchmarkMurmur3Spaolacci/1.00_KB-16         6616          182283 ns/op
BenchmarkWyhashV2Orisano/1.00_KB-16          4830          246171 ns/op

BenchmarkNthash/1.00_MB-16                   100           1.0124147e+07 ns/op
BenchmarkWyhashV3Zeebo/1.00_MB-16            10            1.089315e+08 ns/op
BenchmarkXxhashV3Zeebo/1.00_MB-16            10            1.10472245e+08 ns/op
BenchmarkWyhashV3Orisano/1.00_MB-16          9             1.12783217e+08 ns/op
BenchmarkWyhashV4Orisano/1.00_MB-16          9             1.14976674e+08 ns/op
BenchmarkXxhashV3CespareV1/1.00_MB-16        9             1.17990413e+08 ns/op
BenchmarkXxhashV3CespareV2/1.00_MB-16        8             1.26893102e+08 ns/op
BenchmarkWyhashV3Dgryski/1.00_MB-16          8             1.28696861e+08 ns/op
BenchmarkMurmur3Spaolacci/1.00_MB-16         6             1.78456772e+08 ns/op
BenchmarkMurmur3Twmb/1.00_MB-16              6             1.79073018e+08 ns/op
BenchmarkWyhashV2Orisano/1.00_MB-16          5             2.38264358e+08 ns/op



```
