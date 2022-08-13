# Go Merkle Tree

[![Go Reference](https://pkg.go.dev/badge/github.com/txaty/go-merkletree.svg)](https://pkg.go.dev/github.com/txaty/go-merkletree)
[![Go Report Card](https://goreportcard.com/badge/github.com/txaty/go-merkletree)](https://goreportcard.com/report/github.com/txaty/go-merkletree)
![Coverage](https://img.shields.io/badge/Coverage-87.6%25-brightgreen)

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

In our implementation, ```tree.Build()``` performs tree generation and the proof generation at the same time (time complexity: O(nlogn)), cbergoon/merkletree's ```tree.NewTree()``` only generates the tree. So we benchmark our tree building process with cbergoon/merkletree's tree build + get merkle path ```tree.GetMerklePath()``` for each data block as the proof generation test.

1,000 blocks:

<table>
<thead><tr><th>Linux (i7-9750H)</th><th>M1 Macbook Air</th></tr></thead>
<tbody>
<tr><td>

```bash
goos: linux
goarch: amd64
pkg: github.com/txaty/go-merkletree
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-12             	     685	   1761708 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-12     	     694	   1737775 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-12    	     684	   1750200 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12               	     174	   6866456 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12      	      46	  24669106 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-8            	    2581	    445515 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-8    	    2080	    561253 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-8   	    2588	    464377 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8              	    1026	   1169075 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8     	     200	   6017512 ns/op
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
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-12             	      40	  25252643 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-12     	      96	  13112520 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-12    	      44	  25051554 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12               	      12	  93765796 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12      	       2	 774985957 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-8            	     160	   6930908 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-8    	     205	   5795240 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-8   	     168	   6999481 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8              	      70	  15945774 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8     	       4	 280509167 ns/op
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
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-12             	       5	 209020151 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-12     	       7	 143142651 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-12    	       5	 217958259 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-12               	       1	1152525580 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-12      	       1	47569698074 ns/op
PASS
```

</td><td>

```bash
goos: darwin
goarch: arm64
pkg: github.com/txaty/go-merkletree
BenchmarkMerkleTreeProofGen
BenchmarkMerkleTreeProofGen-8            	      16	  71094940 ns/op
BenchmarkMerkleTreeProofGenParallel
BenchmarkMerkleTreeProofGenParallel-8    	      15	  70876911 ns/op
Benchmark_cbergoonMerkleTreeProofGen
Benchmark_cbergoonMerkleTreeProofGen-8   	      15	  71411361 ns/op
BenchmarkMerkleTreeVerify
BenchmarkMerkleTreeVerify-8              	       6	 192947236 ns/op
Benchmark_cbergoonMerkleTreeVerify
Benchmark_cbergoonMerkleTreeVerify-8     	       1	29294817917 ns/op
PASS
```

</td></tr>
</tbody></table>

(```63145758422 ns/op``` means each function execution takes 63145758422 nanoseconds (around 63.15 seconds, 10^9 ns = 1s))

In conclusion, with large sets of data blocks, our implementation is much faster than cbergoon/merkletree at both tree & proof generation and data block verification.
