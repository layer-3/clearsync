# Go Merkle Tree

[![Go Reference](https://pkg.go.dev/badge/github.com/txaty/go-merkletree.svg)](https://pkg.go.dev/github.com/txaty/go-merkletree)
[![Go Report Card](https://goreportcard.com/badge/github.com/txaty/go-merkletree)](https://goreportcard.com/report/github.com/txaty/go-merkletree)
![Coverage](https://img.shields.io/badge/Coverage-87.8%25-brightgreen)

High performance Merkle Tree Computation in Go (supports parallelization).

## Installation

```bash
go get -u github.com/txaty/go-merkletree
```

## Example

```go
package main

import (
    "crypto/rand"
    "crypto/sha256"
    "fmt"

    mt "github.com/txaty/go-merkletree"
)

// first define a data structure with Serialize method to be used as data block
type testData struct {
    data []byte
}

func (t *testData) Serialize() ([]byte, error) {
    return t.data, nil
}

// define a hash function in this format
func hashFunc(data []byte) ([]byte, error) {
    sha256Func := sha256.New()
    sha256Func.Write(data)
    return sha256Func.Sum(nil), nil
}

func main() {
    // generate dummy data blocks
    var blocks []mt.DataBlock
    for i := 0; i < 1000; i++ {
        block := &testData{
            data: make([]byte, 100),
        }
        _, err := rand.Read(block.data)
        handleError(err)
        blocks = append(blocks, block)
    }

    // create a simple configuration for Merkle Tree generation
    config := &mt.Config{
        HashFunc: hashFunc, // if nil, use SHA256 by default
        // if true, handle odd-number-node situation by duplicating the last node
        AllowDuplicates: true,
    }
    // build a new Merkle Tree
    tree, err := mt.New(config, blocks)
    handleError(err)
    // get the root hash of the Merkle Tree
    rootHash := tree.Root
    // get proves
    proofs := tree.Proofs
    // verify the proofs
    for i := 0; i < len(proofs); i++ {
        ok, err := tree.Verify(blocks[i], proofs[i])
        handleError(err)
        fmt.Println(ok)
    }
    // or you can also do this
    for i := 0; i < len(blocks); i++ {
        // if hashFunc is nil, use SHA256 by default
        ok, err := mt.Verify(blocks[i], proofs[i], rootHash, hashFunc)
        handleError(err)
        fmt.Println(ok)
    }
}

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}
```

## Benchmark

Benchmark with [cbergoon/merkletree](https://github.com/cbergoon/merkletree) (in [bench branch](https://github.com/cbergoon/merkletree)).

In our implementation, ```tree.Build()``` performs tree generation and the proof generation at the same time (time complexity: O(nlogn)), cbergoon/merkletree's ```tree.NewTree()``` only generates the tree. So we benchmark our tree building process with cbergoon/merkletree's tree build + get merkle path ```tree.GetMerklePath()``` for each data block.

1000 blocks:

<table>
<thead><tr><th>Linux (i7-9750H)</th><th>M1 Macbook Air</th></tr></thead>
<tbody>
<tr><td>

```bash
goos: linux
goarch: amd64
pkg: github.com/txaty/go-merkletree
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-12                       523       2221038 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-12               678       1758174 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-12              164       7193082 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12                      176       6787151 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12              48      24503759 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-8                     1926        621450 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-8             1980        597595 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-8             416       2873425 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8                    1024       1162340 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8            198       6064883 ns/op
PASS
```

</td></tr>
</tbody></table>

10,000 blocks:

<table>
<thead><tr><th>Linux (i7-9750H)</th><th>M1 Macbook Air</th></tr></thead>
<tbody>
<tr><td>

```bash
goos: linux
goarch: amd64
pkg: github.com/txaty/go-merkletree
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-12                        44      26247088 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-12                88      13200423 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-12                2     522912836 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12                       12      92832728 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12               2     775982655 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-8                      150       7583059 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-8              193       6213593 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-8               5     231274467 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8                      72      16243839 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8              4     282454323 ns/op
PASS
```

</td></tr>
</tbody></table>

100,000 blocks

<table>
<thead><tr><th>Linux (i7-9750H)</th><th>M1 Macbook Air</th></tr></thead>
<tbody>
<tr><td>

```bash
goos: linux
goarch: amd64
pkg: github.com/txaty/go-merkletree
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-12                         4     314272598 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-12                 7     144025900 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-12                1    59839840747 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12                        1    1128593176 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12               1    63145758422 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeBuild
BenchmarkMerkleTreeBuild-8                       12      99413837 ns/op
BenchmarkMerkleTreeBuildParallel
BenchmarkMerkleTreeBuildParallel-8               14      77042113 ns/op
Benchmark_cbergoonMerkleTreeBuild
Benchmark_cbergoonMerkleTreeBuild-8               1    29609023292 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8                       6     193811917 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8              1    30393054541 ns/op
PASS
```

</td></tr>
</tbody></table>

(```63145758422 ns/op``` means each function execution takes 63145758422 nanoseconds (around 63.15 seconds, 10^9 ns = 1s))

In conclusion, with large sets of data blocks, our implementation is much faster than cbergoon/merkletree at both tree & proof generation and data block verification.
