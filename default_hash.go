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

import "crypto/sha256"

// sha256Digest is the reusable digest for DefaultHashFunc.
// It is used to avoid creating a new hash digest for every call to DefaultHashFunc.
var sha256Digest = sha256.New()

// DefaultHashFunc is the default hash function used when no user-specified hash function is provided.
// It implements the SHA256 hash function and reuses sha256Digest to reduce memory allocations.
func DefaultHashFunc(data []byte) ([]byte, error) {
	defer sha256Digest.Reset()
	sha256Digest.Write(data)
	return sha256Digest.Sum(make([]byte, 0, sha256Digest.Size())), nil
}

// DefaultHashFuncParallel is the default hash function used by parallel algorithms when no user-specified
// hash function is provided. It implements the SHA256 hash function and creates a new hash digest for
// each call, ensuring that it is safe for concurrent use.
func DefaultHashFuncParallel(data []byte) ([]byte, error) {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(make([]byte, 0, digest.Size())), nil
}
