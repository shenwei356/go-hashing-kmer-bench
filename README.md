# Benchmark of hashing k-mers in Golang

## Hash functions

- ntHash
  - https://github.com/will-rowe/nthash
- xxhash v3
  - cespareV1 https://github.com/cespare/xxhash
  - cespareV2 https://github.com/cespare/xxhash/v2
  - zeebo https://github.com/zeebo/xxh3
- murmur
  - https://github.com/spaolacci/murmur3
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

BenchmarkNthash/1.00_KB-16                   123019        9477 ns/op
BenchmarkWyhashV3Zeebo/1.00_KB-16            10000         113580 ns/op
BenchmarkWyhashV4Orisano/1.00_KB-16          10000         117724 ns/op
BenchmarkXxhashV3Zeebo/1.00_KB-16            10000         118054 ns/op
BenchmarkXxhashV3CespareV1/1.00_KB-16        10000         125306 ns/op
BenchmarkWyhashV3Orisano/1.00_KB-16          9243          134073 ns/op
BenchmarkXxhashV3CespareV2/1.00_KB-16        8637          136508 ns/op
BenchmarkWyhashV3Dgryski/1.00_KB-16          8044          145037 ns/op
BenchmarkMurmur3/1.00_KB-16                  6536          183418 ns/op
BenchmarkWyhashV2Orisano/1.00_KB-16          4878          244684 ns/op

BenchmarkNthash/1.00_MB-16                   114           1.0230988e+07 ns/op
BenchmarkWyhashV3Zeebo/1.00_MB-16            10            1.05999001e+08 ns/op
BenchmarkXxhashV3Zeebo/1.00_MB-16            9             1.12517312e+08 ns/op
BenchmarkWyhashV4Orisano/1.00_MB-16          10            1.12686685e+08 ns/op
BenchmarkXxhashV3CespareV1/1.00_MB-16        9             1.17620747e+08 ns/op
BenchmarkWyhashV3Orisano/1.00_MB-16          8             1.26135382e+08 ns/op
BenchmarkXxhashV3CespareV2/1.00_MB-16        8             1.35829517e+08 ns/op
BenchmarkWyhashV3Dgryski/1.00_MB-16          8             1.40426644e+08 ns/op
BenchmarkMurmur3/1.00_MB-16                  6             1.79500025e+08 ns/op
BenchmarkWyhashV2Orisano/1.00_MB-16          5             2.37122505e+08 ns/op



```
