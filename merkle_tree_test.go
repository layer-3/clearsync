// MIT License
//
// Copyright (c) 2023 Tommy TIAN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package merkletree

import (
	"bytes"
	crand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/txaty/go-merkletree/mock"
)

func mockDataBlocks(num int) []DataBlock {
	blocks := make([]DataBlock, num)
	for i := 0; i < num; i++ {
		byteLen := rand.Intn(1 << 15)
		block := &mock.DataBlock{
			Data: make([]byte, byteLen),
		}
		if _, err := crand.Read(block.Data); err != nil {
			panic(err)
		}
		blocks[i] = block
	}
	return blocks
}

func mockDataBlocksFixedSize(num int) []DataBlock {
	blocks := make([]DataBlock, num)
	for i := 0; i < num; i++ {
		block := &mock.DataBlock{
			Data: make([]byte, 128),
		}
		if _, err := crand.Read(block.Data); err != nil {
			panic(err)
		}
		blocks[i] = block
	}
	return blocks
}

func FuzzMerkleTreeNew(f *testing.F) {
	f.Add(10, 0, 4, false, false, false)
	f.Add(128, 1, 3, true, true, true)
	f.Add(256, 2, 2, false, false, true)
	f.Add(512, 0, 1, true, true, false)
	f.Fuzz(func(t *testing.T,
		numBlocks, modeInt, numRoutines int,
		runInParallel, sortSiblingPairs, disableLeafHashing bool,
	) {
		// 0 < numBlocks < 262144
		if numBlocks < 0 {
			numBlocks = -numBlocks
		}
		numBlocks %= 262144
		numBlocks++
		dataBlocks := mockDataBlocks(numBlocks)

		// 0 <= modeInt < 3
		if modeInt < 0 {
			modeInt = -modeInt
		}
		modeInt %= 3
		mode := TypeConfigMode(modeInt)

		// 0 <= numRoutines < 1024
		if numRoutines < 0 {
			numRoutines = -numRoutines
		}
		numRoutines %= 1024

		config := &Config{
			NumRoutines:        numRoutines,
			Mode:               mode,
			RunInParallel:      runInParallel,
			SortSiblingPairs:   sortSiblingPairs,
			DisableLeafHashing: disableLeafHashing,
		}
		mt, err := New(config, dataBlocks)
		if err != nil {
			return
		}
		if mt == nil {
			return
		}

		if mode == ModeProofGen || mode == ModeProofGenAndTreeBuild {
			for idx, block := range dataBlocks {
				ok, err := mt.Verify(block, mt.Proofs[idx])
				if err != nil {
					t.Errorf("proof verification error, idx %d, err %v", idx, err)
					return
				}
				if !ok {
					t.Errorf("proof verification failed, idx %d", idx)
					return
				}
			}
		} else {
			compareTree, err := New(&Config{
				SortSiblingPairs:   sortSiblingPairs,
				DisableLeafHashing: disableLeafHashing,
			}, dataBlocks)
			if err != nil {
				return
			}
			if !bytes.Equal(mt.Root, compareTree.Root) {
				t.Errorf("tree generated is wrong")
				return
			}
			for idx, block := range dataBlocks {
				proof, err := mt.Proof(block)
				if err != nil {
					return
				}
				ok, err := mt.Verify(block, proof)
				if err != nil {
					t.Errorf("proof verification error, idx %d, err %v", idx, err)
					return
				}
				if !ok {
					t.Errorf("proof verification failed, idx %d", idx)
					return
				}
				ok, err = compareTree.Verify(block, proof)
				if err != nil {
					t.Errorf("proof verification error, idx %d, err %v", idx, err)
					return
				}
				if !ok {
					t.Errorf("proof verification failed, idx %d", idx)
					return
				}
			}
		}
	})
}

const benchSize = 65536

func BenchmarkMerkleTreeNew_modeProofGen(b *testing.B) {
	testCases := mockDataBlocksFixedSize(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(nil, testCases)
		if err != nil {
			b.Errorf("New() proof gen error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNew_modeProofGenParallel(b *testing.B) {
	config := &Config{
		RunInParallel: true,
	}
	testCases := mockDataBlocksFixedSize(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("New() proof gen parallel error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNew_modeTreeBuild(b *testing.B) {
	testCases := mockDataBlocksFixedSize(benchSize)
	config := &Config{
		Mode: ModeTreeBuild,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("New() tree build error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNew_modeTreeBuildParallel(b *testing.B) {
	config := &Config{
		Mode:          ModeTreeBuild,
		RunInParallel: true,
	}
	testCases := mockDataBlocksFixedSize(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("New() tree build parallel error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNew_modeProofGenAndTreeBuild(b *testing.B) {
	config := &Config{
		Mode: ModeProofGenAndTreeBuild,
	}
	testCases := mockDataBlocksFixedSize(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("New() proof gen and tree build error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNew_modeProofGenAndTreeBuildParallel(b *testing.B) {
	config := &Config{
		Mode:          ModeProofGenAndTreeBuild,
		RunInParallel: true,
	}
	testCases := mockDataBlocksFixedSize(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("New() proof gen and tree build parallel error = %v", err)
		}
	}
}
