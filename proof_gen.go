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

// Package merkletree implements a high-performance Merkle Tree in Go.
// It supports parallel execution for enhanced performance and
// offers compatibility with OpenZeppelin through sorted sibling pairs.
package merkletree

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

// generateProofs constructs the Merkle Tree and generates the Merkle proofs for each leaf.
// It returns an error if there is an issue during the generation process.
func (m *MerkleTree) generateProofs() (err error) {
	m.initProofs()
	buffer, bufferSize := initBuffer(m.Leaves)
	for step := 0; step < m.Depth; step++ {
		bufferSize = fixOddNumOfNodes(buffer, bufferSize, step)
		m.updateProofs(buffer, bufferSize, step)
		for idx := 0; idx < bufferSize; idx += 2 {
			leftIdx := idx << step
			rightIdx := min(leftIdx+(1<<step), len(buffer)-1)
			buffer[leftIdx], err = m.HashFunc(m.concatHashFunc(buffer[leftIdx], buffer[rightIdx]))
			if err != nil {
				return
			}
		}
		bufferSize >>= 1
	}
	m.Root = buffer[0]
	return
}

// generateProofsParallel generates proofs concurrently for the MerkleTree.
func (m *MerkleTree) generateProofsParallel() (err error) {
	m.initProofs()
	buffer, bufferSize := initBuffer(m.Leaves)
	numRoutines := m.NumRoutines
	for step := 0; step < m.Depth; step++ {
		// Limit the number of workers to the previous level length.
		numRoutines = min(numRoutines, bufferSize)
		bufferSize = fixOddNumOfNodes(buffer, bufferSize, step)
		m.updateProofsParallel(buffer, bufferSize, step)
		eg := new(errgroup.Group)
		for startIdx := 0; startIdx < numRoutines; startIdx++ {
			startIdx := startIdx << 1
			eg.Go(func() error {
				var err error
				for i := startIdx; i < bufferSize; i += numRoutines << 1 {
					leftIdx := i << step
					rightIdx := min(leftIdx+(1<<step), len(buffer)-1)
					buffer[leftIdx], err = m.HashFunc(m.concatHashFunc(buffer[leftIdx], buffer[rightIdx]))
					if err != nil {
						return err
					}
				}
				return nil
			})
		}
		if err = eg.Wait(); err != nil {
			return
		}
		bufferSize >>= 1
	}
	m.Root = buffer[0]
	return
}

// initProofs initializes the MerkleTree's Proofs with the appropriate size and depth.
// This is to reduce overhead of slice resizing during the generation process.
func (m *MerkleTree) initProofs() {
	m.Proofs = make([]*Proof, m.NumLeaves)
	for i := 0; i < m.NumLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
}

// initBuffer initializes the buffer with the leaves and returns the buffer size.
// If the number of leaves is odd, the buffer size is increased by 1.
func initBuffer(leaves [][]byte) ([][]byte, int) {
	var (
		numLeaves = len(leaves)
		buffer    [][]byte
	)
	// If the number of leaves is odd, make initial buffer size even by adding 1.
	if numLeaves&1 == 1 {
		buffer = make([][]byte, numLeaves+1)
	} else {
		buffer = make([][]byte, numLeaves)
	}
	copy(buffer, leaves)
	return buffer, numLeaves
}

// fixOddNumOfNodes adjusts the buffer size if it has an odd number of nodes.
// It appends the last node to the buffer if the buffer length is odd.
func fixOddNumOfNodes(buffer [][]byte, bufferSize, step int) int {
	// If the buffer length is even, no adjustment is needed.
	if bufferSize&1 == 0 {
		return bufferSize
	}
	// Determine the node to append.
	appendNodeIndex := (bufferSize - 1) << step
	// The appended node will be put at the end of the buffer.
	buffer[len(buffer)-1] = buffer[appendNodeIndex]
	bufferSize++
	return bufferSize
}

// updateProofs updates the proofs for all the leaves while constructing the Merkle Tree.
func (m *MerkleTree) updateProofs(buffer [][]byte, bufferSize, step int) {
	batch := 1 << step
	for i := 0; i < bufferSize; i += 2 {
		updateProofInTwoBatches(m.Proofs, buffer, i, batch, step)
	}
}

// updateProofsParallel updates the proofs for all the leaves while constructing the Merkle Tree in parallel.
func (m *MerkleTree) updateProofsParallel(buffer [][]byte, bufferLength, step int) {
	var (
		batch = 1 << step
		wg    sync.WaitGroup
	)
	numRoutines := min(m.NumRoutines, bufferLength)
	wg.Add(numRoutines)
	for startIdx := 0; startIdx < numRoutines; startIdx++ {
		go func(startIdx int) {
			defer wg.Done()
			for i := startIdx; i < bufferLength; i += numRoutines << 1 {
				updateProofInTwoBatches(m.Proofs, buffer, i, batch, step)
			}
		}(startIdx << 1)
	}
	wg.Wait()
}

// updateProofInTwoBatches updates the path and the siblings of the proof in two batches.
func updateProofInTwoBatches(proofs []*Proof, buffer [][]byte, idx, batch, step int) {
	start := idx * batch
	end := min(start+batch, len(proofs))
	siblingNodeIdx := min((idx+1)<<step, len(buffer)-1)
	for i := start; i < end; i++ {
		proofs[i].Path += 1 << step
		proofs[i].Siblings = append(proofs[i].Siblings, buffer[siblingNodeIdx])
	}
	start += batch
	end = min(start+batch, len(proofs))
	siblingNodeIdx = min(idx<<step, len(buffer)-1)
	for i := start; i < end; i++ {
		proofs[i].Siblings = append(proofs[i].Siblings, buffer[siblingNodeIdx])
	}
}
