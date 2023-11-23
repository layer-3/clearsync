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

import "errors"

var (
	// ErrInvalidNumOfDataBlocks is the error for an invalid number of data blocks.
	ErrInvalidNumOfDataBlocks = errors.New("the number of data blocks must be greater than 1")
	// ErrInvalidConfigMode is the error for an invalid configuration mode.
	ErrInvalidConfigMode = errors.New("invalid configuration mode")
	// ErrProofIsNil is the error for a nil proof.
	ErrProofIsNil = errors.New("proof is nil")
	// ErrDataBlockIsNil is the error for a nil data block.
	ErrDataBlockIsNil = errors.New("data block is nil")
	// ErrProofInvalidModeTreeNotBuilt is the error for an invalid mode in Proof() function.
	// Proof() function requires a built tree to generate the proof.
	ErrProofInvalidModeTreeNotBuilt = errors.New("merkle tree is not in built, could not generate proof by this method")
	// ErrProofInvalidDataBlock is the error for an invalid data block in Proof() function.
	ErrProofInvalidDataBlock = errors.New("data block is not a member of the merkle tree")
)
