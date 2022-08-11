# Go Merkle Tree

[![Go Reference](https://pkg.go.dev/badge/github.com/tommytim0515/go-merkletree.svg)](https://pkg.go.dev/github.com/tommytim0515/go-merkletree)
[![Go Report Card](https://goreportcard.com/badge/github.com/tommytim0515/go-merkletree)](https://goreportcard.com/report/github.com/tommytim0515/go-merkletree)
![Coverage](https://img.shields.io/badge/Coverage-81.5%25-brightgreen)

High performance Merkle Tree Computation in Go (supports parallelization). 

## Installation

```bash
go get -u github.com/tommytim0515/go-merkletree
```


## Example

```go
package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	mt "github.com/tommytim0515/go-merkletree"
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
	// create a simple configuration for Merkle Tree generation
	config := &mt.Config{
		HashFunc: hashFunc, // if nil, use SHA256 by default
		// if true, handle odd-number-node situation by duplicating the last node
		AllowDuplicates: true,
	}
	tree := mt.NewMerkleTree(config)

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

	// build the Merkle Tree
	err := tree.Build(blocks)
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
