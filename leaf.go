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

import "golang.org/x/sync/errgroup"

// computeLeafNodes compute the leaf nodes from the data blocks.
func (m *MerkleTree) computeLeafNodes(blocks []DataBlock) ([][]byte, error) {
	var (
		leaves             = make([][]byte, m.NumLeaves)
		hashFunc           = m.HashFunc
		disableLeafHashing = m.DisableLeafHashing
		err                error
	)
	for i := 0; i < m.NumLeaves; i++ {
		if leaves[i], err = dataBlockToLeaf(blocks[i], hashFunc, disableLeafHashing); err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

// computeLeafNodesParallel compute the leaf nodes from the data blocks in parallel.
func (m *MerkleTree) computeLeafNodesParallel(blocks []DataBlock) ([][]byte, error) {
	var (
		lenLeaves          = len(blocks)
		leaves             = make([][]byte, lenLeaves)
		numRoutines        = m.NumRoutines
		hashFunc           = m.HashFunc
		disableLeafHashing = m.DisableLeafHashing
		eg                 = new(errgroup.Group)
	)
	numRoutines = min(numRoutines, lenLeaves)
	for startIdx := 0; startIdx < numRoutines; startIdx++ {
		startIdx := startIdx
		eg.Go(func() error {
			var err error
			for i := startIdx; i < lenLeaves; i += numRoutines {
				if leaves[i], err = dataBlockToLeaf(blocks[i], hashFunc, disableLeafHashing); err != nil {
					return err
				}
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return leaves, nil
}

// dataBlockToLeaf generates the leaf from the data block.
// If the leaf hashing is disabled, the data block is returned as the leaf.
func dataBlockToLeaf(block DataBlock, hashFunc TypeHashFunc, disableLeafHashing bool) ([]byte, error) {
	blockBytes, err := block.Serialize()
	if err != nil {
		return nil, err
	}
	if disableLeafHashing {
		// copy the value so that the original byte slice is not modified
		leaf := make([]byte, len(blockBytes))
		copy(leaf, blockBytes)
		return leaf, nil
	}
	return hashFunc(blockBytes)
}
