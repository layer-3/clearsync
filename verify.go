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

import "bytes"

// Verify checks if the data block is valid using the Merkle Tree proof and the cached Merkle root hash.
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.Config)
}

// Verify checks if the data block is valid using the Merkle Tree proof and the provided Merkle root hash.
// It returns true if the data block is valid, false otherwise. An error is returned in case of any issues
// during the verification process.
func Verify(dataBlock DataBlock, proof *Proof, root []byte, config *Config) (bool, error) {
	// Validate input parameters.
	if dataBlock == nil {
		return false, ErrDataBlockIsNil
	}
	if proof == nil {
		return false, ErrProofIsNil
	}
	if config == nil {
		config = new(Config)
	}
	if config.HashFunc == nil {
		config.HashFunc = DefaultHashFunc
	}

	// Determine the concatenation function based on the configuration.
	concatFunc := concatHash
	if config.SortSiblingPairs {
		concatFunc = concatSortHash
	}

	// Convert the data block to a leaf.
	leaf, err := dataBlockToLeaf(dataBlock, config.HashFunc, config.DisableLeafHashing)
	if err != nil {
		return false, err
	}

	// Traverse the Merkle proof and compute the resulting hash.
	// Copy the slice so that the original leaf won't be modified.
	result := make([]byte, len(leaf))
	copy(result, leaf)
	path := proof.Path
	for _, sib := range proof.Siblings {
		if path&1 == 1 {
			result, err = config.HashFunc(concatFunc(result, sib))
		} else {
			result, err = config.HashFunc(concatFunc(sib, result))
		}
		if err != nil {
			return false, err
		}
		path >>= 1
	}
	return bytes.Equal(result, root), nil
}
