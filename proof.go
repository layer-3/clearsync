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

// Proof represents a Merkle Tree proof.
type Proof struct {
	Siblings [][]byte // Sibling nodes to the Merkle Tree path of the data block.
	Path     uint32   // Path variable indicating whether the neighbor is on the left or right.
}

// Proof generates the Merkle proof for a data block using the previously generated Merkle Tree structure.
// This method is only available when the configuration mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
// In ModeProofGen, proofs for all the data blocks are already generated, and the Merkle Tree structure
// is not cached.
func (m *MerkleTree) Proof(dataBlock DataBlock) (*Proof, error) {
	if m.Mode != ModeTreeBuild && m.Mode != ModeProofGenAndTreeBuild {
		return nil, ErrProofInvalidModeTreeNotBuilt
	}

	// Convert the data block to a leaf.
	leaf, err := dataBlockToLeaf(dataBlock, m.HashFunc, m.DisableLeafHashing)
	if err != nil {
		return nil, err
	}

	// Retrieve the index of the leaf in the Merkle Tree.
	m.leafMapMu.Lock()
	idx, ok := m.leafMap[string(leaf)]
	m.leafMapMu.Unlock()
	if !ok {
		return nil, ErrProofInvalidDataBlock
	}

	// Compute the path and siblings for the proof.
	var (
		path     uint32
		siblings = make([][]byte, m.Depth)
	)
	for i := 0; i < m.Depth; i++ {
		if idx&1 == 1 {
			siblings[i] = m.nodes[i][idx-1]
		} else {
			path += 1 << i
			siblings[i] = m.nodes[i][idx+1]
		}
		idx >>= 1
	}
	return &Proof{
		Path:     path,
		Siblings: siblings,
	}, nil
}
