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

import "golang.org/x/sync/errgroup"

// treeBuild builds the Merkle Tree and stores all the nodes.
func (m *MerkleTree) treeBuild() (err error) {
	finishMap := make(chan struct{})
	go m.workerBuildLeafMap(finishMap)
	m.initNodes()
	for i := 0; i < m.Depth-1; i++ {
		m.nodes[i] = appendNodeIfOdd(m.nodes[i])
		numNodes := len(m.nodes[i])
		m.nodes[i+1] = make([][]byte, numNodes>>1)
		for j := 0; j < numNodes; j += 2 {
			if m.nodes[i+1][j>>1], err = m.HashFunc(
				m.concatHashFunc(m.nodes[i][j], m.nodes[i][j+1]),
			); err != nil {
				return
			}
		}
	}
	if m.Root, err = m.HashFunc(m.concatHashFunc(
		m.nodes[m.Depth-1][0], m.nodes[m.Depth-1][1],
	)); err != nil {
		return
	}
	<-finishMap
	return
}

// treeBuildParallel builds the Merkle Tree and stores all the nodes in parallel.
func (m *MerkleTree) treeBuildParallel() (err error) {
	finishMap := make(chan struct{})
	go m.workerBuildLeafMap(finishMap)
	m.initNodes()
	for i := 0; i < m.Depth-1; i++ {
		m.nodes[i] = appendNodeIfOdd(m.nodes[i])
		numNodes := len(m.nodes[i])
		m.nodes[i+1] = make([][]byte, numNodes>>1)
		numRoutines := min(m.NumRoutines, numNodes)
		eg := new(errgroup.Group)
		for startIdx := 0; startIdx < numRoutines; startIdx++ {
			startIdx := startIdx
			eg.Go(func() error {
				for j := startIdx << 1; j < numNodes; j += numRoutines << 1 {
					newHash, err := m.HashFunc(m.concatHashFunc(
						m.nodes[i][j], m.nodes[i][j+1],
					))
					if err != nil {
						return err
					}
					m.nodes[i+1][j>>1] = newHash
				}
				return nil
			})
		}
		if err = eg.Wait(); err != nil {
			return
		}
	}
	if m.Root, err = m.HashFunc(m.concatHashFunc(
		m.nodes[m.Depth-1][0], m.nodes[m.Depth-1][1],
	)); err != nil {
		return
	}
	<-finishMap
	return
}

func (m *MerkleTree) workerBuildLeafMap(finishChan chan struct{}) {
	m.leafMapMu.Lock()
	defer m.leafMapMu.Unlock()
	for i := 0; i < m.NumLeaves; i++ {
		m.leafMap[string(m.Leaves[i])] = i
	}
	finishChan <- struct{}{} // empty channel to serve as a wait group for map generation
}

func (m *MerkleTree) initNodes() {
	m.nodes = make([][][]byte, m.Depth)
	m.nodes[0] = make([][]byte, m.NumLeaves)
	copy(m.nodes[0], m.Leaves)
}

func appendNodeIfOdd(buffer [][]byte) [][]byte {
	if len(buffer)&1 == 0 {
		return buffer
	}
	appendNode := buffer[len(buffer)-1]
	buffer = append(buffer, appendNode)
	return buffer
}
