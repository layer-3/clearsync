// MIT License
//
// Copyright (c) 2022 Tommy TIAN
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
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"
	"runtime"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	// Default hash result length using SHA256.
	defaultHashLen = 32
	// ModeProofGen is the proof generation configuration mode.
	ModeProofGen ModeType = iota
	// ModeTreeBuild is the tree building configuration mode.
	ModeTreeBuild
	// ModeProofGenAndTreeBuild is the proof generation and tree building configuration mode.
	ModeProofGenAndTreeBuild
)

// ModeType is the type in the Merkle Tree configuration indicating what operations are performed.
type ModeType int

// DataBlock is the interface of input data blocks to generate the Merkle Tree.
type DataBlock interface {
	Serialize() ([]byte, error)
}

// HashFuncType is the signature of the hash functions used for Merkle Tree generation.
type HashFuncType func([]byte) ([]byte, error)

// Config is the configuration of Merkle Tree.
type Config struct {
	// Customizable hash function used for tree generation.
	HashFunc HashFuncType
	// If true, the generation runs in parallel, otherwise runs without parallelization.
	// This increase the performance for the calculation of large number of data blocks, e.g. over 10,000 blocks.
	RunInParallel bool
	// Number of goroutines run in parallel.
	// If RunInParallel is true and NumRoutine is set to 0, use number of CPU as the number of goroutines.
	NumRoutines int
	// If true, generate a dummy node with random hash value.
	// Otherwise, then the odd node situation is handled by duplicating the previous node.
	NoDuplicates bool
	// Mode of the Merkle Tree generation.
	Mode ModeType
}

// MerkleTree implements the Merkle Tree structure
type MerkleTree struct {
	// Config is the Merkle Tree configuration
	*Config
	// Root is the Merkle root hash
	Root []byte
	// Leaves are Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation
	Leaves [][]byte
	// Proofs are proofs to the data blocks generated during the tree building process
	Proofs []*Proof
	// Depth is the Merkle Tree depth
	Depth uint32
	// tree is the Merkle Tree structure, only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	tree [][][]byte
	// leafMap is the map of the leaf hash to the index in the Tree slice,
	// only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	leafMap sync.Map
}

// Proof implements the Merkle Tree proof.
type Proof struct {
	Path     uint32   // path variable indicating whether the neighbor is on the left or right
	Siblings [][]byte // sibling nodes to the Merkle Tree path of the data block
}

// New generates a new Merkle Tree with specified configuration.
func New(config *Config, blocks []DataBlock) (m *MerkleTree, err error) {
	if len(blocks) <= 1 {
		return nil, errors.New("the number of data blocks must be greater than 1")
	}
	if config == nil {
		config = new(Config)
	}
	if config.HashFunc == nil {
		config.HashFunc = defaultHashFunc
	}
	// If the configuration mode is not set, then set it to ModeProofGen by default.
	if config.Mode == 0 {
		config.Mode = ModeProofGen
	}
	// If RunInParallel is true and NumRoutines is unset, then set NumRoutines to the number of CPU.
	if config.RunInParallel && config.NumRoutines == 0 {
		config.NumRoutines = runtime.NumCPU()
	}
	m = &MerkleTree{
		Config: config,
	}
	m.Depth = calTreeDepth(len(blocks))
	if m.Mode == ModeProofGen {
		if m.RunInParallel {
			m.Leaves, err = leafGenParal(blocks, m.HashFunc, m.NumRoutines)
			if err != nil {
				return
			}
			err = m.proofGenParal()
			return
		}
		m.Leaves, err = leafGen(blocks, m.HashFunc)
		if err != nil {
			return
		}
		err = m.proofGen()
		return
	}
	if m.Mode == ModeTreeBuild {
		if m.RunInParallel {
			m.Leaves, err = leafGenParal(blocks, m.HashFunc, m.NumRoutines)
			if err != nil {
				return
			}
			err = m.treeBuildParal()
			return
		}
		m.Leaves, err = leafGen(blocks, m.HashFunc)
		if err != nil {
			return
		}
		err = m.treeBuild()
		return
	}
	if m.Mode == ModeProofGenAndTreeBuild {
		if m.RunInParallel {
			m.Leaves, err = leafGenParal(blocks, m.HashFunc, m.NumRoutines)
			if err != nil {
				return
			}
			err = m.treeBuildParal()
			if err != nil {
				return
			}
			numLeaves := len(m.Leaves)
			m.Proofs = make([]*Proof, numLeaves)
			for i := 0; i < numLeaves; i++ {
				m.Proofs[i] = new(Proof)
			}
			for i := 0; i < len(m.tree); i++ {
				m.updateProofsParal(m.tree[i], len(m.tree[i]), i)
			}
		}
		m.Leaves, err = leafGen(blocks, m.HashFunc)
		if err != nil {
			return
		}
		err = m.treeBuild()
		if err != nil {
			return
		}
		numLeaves := len(m.Leaves)
		m.Proofs = make([]*Proof, numLeaves)
		for i := 0; i < numLeaves; i++ {
			m.Proofs[i] = new(Proof)
			m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
		}
		for i := 0; i < len(m.tree); i++ {
			m.updateProofs(m.tree[i], len(m.tree[i]), i)
		}
		return
	}
	return nil, errors.New("invalid configuration mode")
}

// calTreeDepth calculates the tree depth,
// the tree depth is then used to declare the capacity of the proof slices.
func calTreeDepth(blockLen int) uint32 {
	log2BlockLen := math.Log2(float64(blockLen))
	// check if log2BlockLen is an integer
	if log2BlockLen != math.Trunc(log2BlockLen) {
		return uint32(log2BlockLen) + 1
	}
	return uint32(log2BlockLen)
}

func (m *MerkleTree) proofGen() (err error) {
	numLeaves := len(m.Leaves)
	m.Proofs = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
	buf := make([][]byte, numLeaves)
	copy(buf, m.Leaves)
	var prevLen int
	buf, prevLen, err = m.fixOdd(buf, numLeaves)
	if err != nil {
		return
	}
	m.updateProofs(buf, numLeaves, 0)
	for step := 1; step < int(m.Depth); step++ {
		for idx := 0; idx < prevLen; idx += 2 {
			buf[idx>>1], err = m.HashFunc(append(buf[idx], buf[idx+1]...))
			if err != nil {
				return
			}
		}
		prevLen >>= 1
		buf, prevLen, err = m.fixOdd(buf, prevLen)
		if err != nil {
			return
		}
		m.updateProofs(buf, prevLen, step)
	}
	m.Root, err = m.HashFunc(append(buf[0], buf[1]...))
	return
}

func (m *MerkleTree) proofGenParal() (err error) {
	numRoutines := m.NumRoutines
	numLeaves := len(m.Leaves)
	m.Proofs = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proofs[i] = new(Proof)
	}
	buf1 := make([][]byte, numLeaves)
	copy(buf1, m.Leaves)
	var prevLen int
	buf1, prevLen, err = m.fixOdd(buf1, numLeaves)
	if err != nil {
		return
	}
	buf2 := make([][]byte, prevLen>>1)
	m.updateProofsParal(buf1, numLeaves, 0)
	for step := 1; step < int(m.Depth); step++ {
		if err != nil {
			return
		}
		g := new(errgroup.Group)
		for i := 0; i < numRoutines && i < prevLen; i++ {
			idx := i << 1
			g.Go(func() error {
				for j := idx; j < prevLen; j += numRoutines << 1 {
					newHash, err := m.HashFunc(append(buf1[j], buf1[j+1]...))
					if err != nil {
						return err
					}
					buf2[j>>1] = newHash
				}
				return nil
			})
		}
		if err = g.Wait(); err != nil {
			return
		}
		buf1, buf2 = buf2, buf1
		prevLen >>= 1
		buf1, prevLen, err = m.fixOdd(buf1, prevLen)
		if err != nil {
			return
		}
		m.updateProofsParal(buf1, prevLen, step)
	}
	m.Root, err = m.HashFunc(append(buf1[0], buf1[1]...))
	return
}

// if the length of the buffer calculating the Merkle Tree is odd, then append a node to the buffer
// if AllowDuplicates is true, append a node by duplicating the previous node
// otherwise, append a node by random
func (m *MerkleTree) fixOdd(buf [][]byte, prevLen int) ([][]byte, int, error) {
	if prevLen&1 == 0 {
		return buf, prevLen, nil
	}
	var appendNode []byte
	if m.NoDuplicates {
		var err error
		appendNode, err = getDummyHash()
		if err != nil {
			return nil, 0, err
		}
	} else {
		appendNode = buf[prevLen-1]
	}
	prevLen++
	if len(buf) < prevLen {
		buf = append(buf, appendNode)
	} else {
		buf[prevLen-1] = appendNode
	}
	return buf, prevLen, nil
}

func (m *MerkleTree) updateProofs(buf [][]byte, bufLen, step int) {
	batch := 1 << step
	for i := 0; i < bufLen; i += 2 {
		m.updatePairProof(buf, i, batch, step)
	}
}

func (m *MerkleTree) updateProofsParal(buf [][]byte, bufLen, step int) {
	numRoutines := m.NumRoutines
	batch := 1 << step
	wg := new(sync.WaitGroup)
	for i := 0; i < numRoutines; i++ {
		idx := i << 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := idx; j < bufLen; j += numRoutines << 1 {
				m.updatePairProof(buf, j, batch, step)
			}
		}()
	}
	wg.Wait()
}

func (m *MerkleTree) updatePairProof(buf [][]byte, idx, batch, step int) {
	start := idx * batch
	end := start + batch
	if end > len(m.Proofs) {
		end = len(m.Proofs)
	}
	for j := start; j < end; j++ {
		m.Proofs[j].Path += 1 << step
		m.Proofs[j].Siblings = append(m.Proofs[j].Siblings, buf[idx+1])
	}
	start = (idx + 1) * batch
	end = start + batch
	if end > len(m.Proofs) {
		end = len(m.Proofs)
	}
	for j := start; j < end; j++ {
		m.Proofs[j].Siblings = append(m.Proofs[j].Siblings, buf[idx])
	}
}

// generate a dummy hash to make odd-length buffer even
func getDummyHash() ([]byte, error) {
	dummyBytes := make([]byte, defaultHashLen)
	_, err := rand.Read(dummyBytes)
	if err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

// defaultHashFunc is used when no user hash function is specified.
// It implements SHA256 hash function.
func defaultHashFunc(data []byte) ([]byte, error) {
	sha256Func := sha256.New()
	sha256Func.Write(data)
	return sha256Func.Sum(nil), nil
}

func leafGen(blocks []DataBlock, hashFunc HashFuncType) ([][]byte, error) {
	var (
		lenLeaves = len(blocks)
		leaves    = make([][]byte, lenLeaves)
	)
	for i := 0; i < lenLeaves; i++ {
		data, err := blocks[i].Serialize()
		if err != nil {
			return nil, err
		}
		hash, err := hashFunc(data)
		if err != nil {
			return nil, err
		}
		leaves[i] = hash
	}
	return leaves, nil
}

func leafGenParal(blocks []DataBlock, hashFunc HashFuncType, numRoutines int) ([][]byte, error) {
	var (
		lenLeaves = len(blocks)
		leaves    = make([][]byte, lenLeaves)
	)
	g := new(errgroup.Group)
	for i := 0; i < numRoutines; i++ {
		idx := i
		g.Go(func() error {
			for j := idx; j < lenLeaves; j += numRoutines {
				data, err := blocks[j].Serialize()
				if err != nil {
					return err
				}
				var hash []byte
				hash, err = hashFunc(data)
				if err != nil {
					return err
				}
				leaves[j] = hash
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return leaves, nil
}

func (m *MerkleTree) treeBuild() (err error) {
	numLeaves := len(m.Leaves)
	m.leafMap = sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numLeaves; i++ {
			m.leafMap.Store(string(m.Leaves[i]), i)
		}
	}()
	m.tree = make([][][]byte, m.Depth)
	m.tree[0] = make([][]byte, numLeaves)
	copy(m.tree[0], m.Leaves)
	var prevLen int
	m.tree[0], prevLen, err = m.fixOdd(m.tree[0], numLeaves)
	if err != nil {
		return
	}
	for i := uint32(0); i < m.Depth-1; i++ {
		m.tree[i+1] = make([][]byte, prevLen>>1)
		for j := 0; j < prevLen; j += 2 {
			m.tree[i+1][j>>1], err = m.HashFunc(append(m.tree[i][j], m.tree[i][j+1]...))
			if err != nil {
				return
			}
		}
		m.tree[i+1], prevLen, err = m.fixOdd(m.tree[i+1], len(m.tree[i+1]))
		if err != nil {
			return
		}
	}
	m.Root, err = m.HashFunc(append(m.tree[m.Depth-1][0], m.tree[m.Depth-1][1]...))
	if err != nil {
		return
	}
	wg.Wait()
	return
}

func (m *MerkleTree) treeBuildParal() (err error) {
	numRoutines := m.NumRoutines
	numLeaves := len(m.Leaves)
	m.leafMap = sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numLeaves; i++ {
			m.leafMap.Store(string(m.Leaves[i]), i)
		}
	}()
	m.tree = make([][][]byte, m.Depth)
	m.tree[0] = make([][]byte, numLeaves)
	copy(m.tree[0], m.Leaves)
	var prevLen int
	m.tree[0], prevLen, err = m.fixOdd(m.tree[0], numLeaves)
	if err != nil {
		return
	}
	for i := uint32(0); i < m.Depth-1; i++ {
		m.tree[i+1] = make([][]byte, prevLen>>1)
		g := new(errgroup.Group)
		for j := 0; j < numRoutines && j < prevLen; j++ {
			idx := j << 1
			g.Go(func() error {
				for k := idx; k < prevLen; k += numRoutines << 1 {
					newHash, err := m.HashFunc(append(m.tree[i][k], m.tree[i][k+1]...))
					if err != nil {
						return err
					}
					m.tree[i+1][k>>1] = newHash
				}
				return nil
			})
		}
		if err = g.Wait(); err != nil {
			return
		}
		m.tree[i+1], prevLen, err = m.fixOdd(m.tree[i+1], len(m.tree[i+1]))
		if err != nil {
			return
		}
	}
	m.Root, err = m.HashFunc(append(m.tree[m.Depth-1][0], m.tree[m.Depth-1][1]...))
	if err != nil {
		return
	}
	wg.Wait()
	return
}

// Verify verifies the data block with the Merkle Tree proof
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.HashFunc)
}

// Verify verifies the data block with the Merkle Tree proof and Merkle root hash
func Verify(dataBlock DataBlock, proof *Proof, root []byte, hashFunc HashFuncType) (bool, error) {
	if dataBlock == nil {
		return false, errors.New("data block is nil")
	}
	if proof == nil {
		return false, errors.New("proof is nil")
	}
	if hashFunc == nil {
		hashFunc = defaultHashFunc
	}
	var (
		data, err = dataBlock.Serialize()
		hash      []byte
	)
	if err != nil {
		return false, err
	}
	hash, err = hashFunc(data)
	if err != nil {
		return false, err
	}
	path := proof.Path
	for _, n := range proof.Siblings {
		if path&1 == 1 {
			hash, err = hashFunc(append(hash, n...))
		} else {
			hash, err = hashFunc(append(n, hash...))
		}
		if err != nil {
			return false, err
		}
		path >>= 1
	}
	return bytes.Equal(hash, root), nil
}

// GenerateProof generates the Merkle proof for a data block with the Merkle Tree structure generated beforehand.
// The method is only available when the configuration mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
// In ModeProofGen, proofs for all the data blocks are already generated, and the Merkle Tree structure is not cached.
func (m *MerkleTree) GenerateProof(dataBlock DataBlock) (*Proof, error) {
	if m.Mode != ModeTreeBuild && m.Mode != ModeProofGenAndTreeBuild {
		return nil, errors.New("merkle Tree is not in built, could not generate proof by this method")
	}
	blockByte, err := dataBlock.Serialize()
	if err != nil {
		return nil, err
	}
	blockHash, err := m.HashFunc(blockByte)
	if err != nil {
		return nil, err
	}
	val, ok := m.leafMap.Load(string(blockHash))
	if !ok {
		return nil, errors.New("data block is not a member of the Merkle Tree")
	}
	var (
		idx      = val.(int)
		path     uint32
		siblings = make([][]byte, m.Depth)
	)
	for i := uint32(0); i < m.Depth; i++ {
		if idx&1 == 1 {
			siblings[i] = m.tree[i][idx-1]
		} else {
			path += 1 << i
			siblings[i] = m.tree[i][idx+1]
		}
		idx >>= 1
	}
	return &Proof{
		Path:     path,
		Siblings: siblings,
	}, nil
}
