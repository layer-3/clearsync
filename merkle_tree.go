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
	"crypto/rand"
	"errors"
	"runtime"
	"sync"

	"github.com/txaty/gool"
)

const (
	// ModeProofGen is the proof generation configuration mode.
	ModeProofGen TypeConfigMode = iota
	// ModeTreeBuild is the tree building configuration mode.
	ModeTreeBuild
	// ModeProofGenAndTreeBuild is the proof generation and tree building configuration mode.
	ModeProofGenAndTreeBuild
	// Default hash result length using SHA256.
	defaultHashLen = 32
)

var wp *gool.Pool[argType, error]

// argType is used as the arguments for the handler functions when performing parallel computations.
// All the handler functions use this universal argument struct to eliminate interface conversion overhead.
// Each field in the struct may be used for different purpose in different handler functions,
// please refer to each handler function for details.
type argType struct {
	mt             *MerkleTree
	byteField1     [][]byte
	byteField2     [][]byte
	dataBlockField []DataBlock
	intField1      int
	intField2      int
	intField3      int
	intField4      int
	intField5      int
	uint32Field    uint32
}

// TypeConfigMode is the type in the Merkle Tree configuration indicating what operations are performed.
type TypeConfigMode int

// DataBlock is the interface of input data blocks to generate the Merkle Tree.
type DataBlock interface {
	Serialize() ([]byte, error)
}

// TypeHashFunc is the signature of the hash functions used for Merkle Tree generation.
type TypeHashFunc func([]byte) ([]byte, error)

// Config is the configuration of Merkle Tree.
type Config struct {
	// concatFunc is the function for concatenating two hashes.
	// If SortSiblingPairs in Config is true, then the sibling pairs are first sorted and then concatenated,
	// supporting the OpenZeppelin Merkle Tree protocol.
	// Otherwise, the sibling pairs are concatenated directly.
	concatFunc func([]byte, []byte) []byte
	// Customizable hash function used for tree generation.
	HashFunc TypeHashFunc
	// Number of goroutines run in parallel.
	// If RunInParallel is true and NumRoutine is set to 0, use number of CPU as the number of goroutines.
	NumRoutines int
	// Mode of the Merkle Tree generation.
	Mode TypeConfigMode
	// If RunInParallel is true, the generation runs in parallel, otherwise runs without parallelization.
	// This increase the performance for the calculation of large number of data blocks, e.g. over 10,000 blocks.
	RunInParallel bool
	// If true, generate a dummy node with random hash value.
	// Otherwise, then the odd node situation is handled by duplicating the previous node.
	NoDuplicates bool
	// SortSiblingPairs is the parameter for OpenZeppelin compatibility.
	// If set to `true`, the hashing sibling pairs are sorted.
	SortSiblingPairs bool
	// If true, the leaf nodes are NOT hashed before being added to the Merkle Tree.
	DisableLeafHashing bool
}

// MerkleTree implements the Merkle Tree structure.
type MerkleTree struct {
	*Config
	// leafMap is the map of the leaf hash to the index in the Tree slice.
	// It is only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
	leafMap sync.Map
	// nodes contains Merkle Tree's tree structure.
	// It is only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
	nodes [][][]byte
	// Root is the Merkle root hash.
	Root []byte
	// Leaves are Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation.
	Leaves [][]byte
	// Proofs are proofs to the data blocks generated during the tree building process.
	Proofs []*Proof
	// Depth is the Merkle Tree depth.
	Depth uint32
	// NumLeaves is the number of tree leaves, it is fixed when the tree is built.
	NumLeaves int
}

// Proof implements the Merkle Tree proof.
type Proof struct {
	Siblings [][]byte // sibling nodes to the Merkle Tree path of the data block.
	Path     uint32   // path variable indicating whether the neighbor is on the left or right.
}

// New generates a new Merkle Tree with specified configuration.
func New(config *Config, blocks []DataBlock) (m *MerkleTree, err error) {
	if len(blocks) <= 1 {
		return nil, errors.New("the number of data blocks must be greater than 1")
	}
	if config == nil {
		config = new(Config)
	}
	m = &MerkleTree{Config: config, NumLeaves: len(blocks), Depth: calTreeDepth(len(blocks))}
	// Hash function initialization.
	if m.HashFunc == nil {
		if m.RunInParallel {
			m.HashFunc = defaultHashFuncParallel // Parallelized hash function must be concurrent safe.
		} else {
			m.HashFunc = defaultHashFunc
		}
	}
	// Hash concatenation function initialization.
	if m.concatFunc == nil {
		if m.SortSiblingPairs {
			m.concatFunc = concatSortHash
		} else {
			m.concatFunc = concatHash
		}
	}
	// Configuration for parallelization.
	if m.RunInParallel {
		// If NumRoutines is not set or invalid, set it to the number of CPU.
		if m.NumRoutines <= 0 {
			m.NumRoutines = runtime.NumCPU()
		}
		// Generic wait group initialization (for parallelized computation) and leaf generation.
		// Task channel capacity is passed as 0, so use the default value: 2 * numWorkers.
		wp = gool.NewPool[argType, error](m.NumRoutines, 0)
		defer wp.Close()
		if m.Leaves, err = m.leafGenParallel(blocks); err != nil {
			return nil, err
		}
	} else {
		if m.Leaves, err = m.leafGen(blocks); err != nil {
			return nil, err
		}
	}

	// Mode defined actions.
	// If the configuration mode is not set, then set it to ModeProofGen by default.
	if m.Mode == 0 {
		m.Mode = ModeProofGen
	}
	if m.Mode == ModeProofGen {
		err = m.proofGen()
		return
	}
	if m.Mode == ModeTreeBuild {
		err = m.treeBuild()
		return
	}
	if m.Mode == ModeProofGenAndTreeBuild {
		if err = m.treeBuild(); err != nil {
			return
		}
		m.initProofs()
		if m.RunInParallel {
			for i := 0; i < len(m.nodes); i++ {
				m.updateProofsParallel(m.nodes[i], len(m.nodes[i]), i)
			}
			return
		}
		for i := 0; i < len(m.nodes); i++ {
			m.updateProofs(m.nodes[i], len(m.nodes[i]), i)
		}
		return
	}

	return nil, errors.New("invalid configuration mode")
}

func concatHash(b1 []byte, b2 []byte) []byte {
	return append(b1, b2...)
}

func concatSortHash(b1 []byte, b2 []byte) []byte {
	if bytes.Compare(b1, b2) < 0 {
		return append(b1, b2...)
	}
	return append(b2, b1...)
}

// calTreeDepth calculates the tree depth.
// The tree depth is then used to declare the capacity of the proof slices.
func calTreeDepth(blockLen int) uint32 {
	depth := uint32(0)
	for blockLen > 1 {
		blockLen = (blockLen + 1) / 2
		depth++
	}
	return depth
}

func (m *MerkleTree) initProofs() {
	m.Proofs = make([]*Proof, m.NumLeaves)
	for i := 0; i < m.NumLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
}

func (m *MerkleTree) proofGen() (err error) {
	m.initProofs()
	buf := make([][]byte, m.NumLeaves)
	copy(buf, m.Leaves)
	var prevLen int
	if buf, prevLen, err = m.fixOdd(buf, m.NumLeaves); err != nil {
		return
	}
	if m.RunInParallel {
		buff := make([][]byte, prevLen>>1)
		m.updateProofsParallel(buf, m.NumLeaves, 0)
		numRoutines := m.NumRoutines
		for step := 1; step < int(m.Depth); step++ {
			if numRoutines > prevLen {
				numRoutines = prevLen
			}
			argList := make([]argType, numRoutines)
			for i := 0; i < numRoutines; i++ {
				argList[i] = argType{
					mt:         m,
					byteField1: buf,
					byteField2: buff,
					intField1:  i << 1, // starting index
					intField2:  prevLen,
					intField3:  numRoutines,
				}
			}
			errList := wp.Map(proofGenHandler, argList)
			for _, err = range errList {
				if err != nil {
					return
				}
			}
			buf, buff = buff, buf
			prevLen >>= 1
			if buf, prevLen, err = m.fixOdd(buf, prevLen); err != nil {
				return
			}
			m.updateProofsParallel(buf, prevLen, step)
		}
	} else {
		m.updateProofs(buf, m.NumLeaves, 0)
		for step := 1; step < int(m.Depth); step++ {
			for idx := 0; idx < prevLen; idx += 2 {
				buf[idx>>1], err = m.HashFunc(m.concatFunc(buf[idx], buf[idx+1]))
				if err != nil {
					return
				}
			}
			prevLen >>= 1
			if buf, prevLen, err = m.fixOdd(buf, prevLen); err != nil {
				return
			}
			m.updateProofs(buf, prevLen, step)
		}
	}

	m.Root, err = m.HashFunc(m.concatFunc(buf[0], buf[1]))
	return
}

// proofGenHandler generates the proofs in parallel.
func proofGenHandler(arg argType) error {
	var (
		hashFunc    = arg.mt.HashFunc
		concatFunc  = arg.mt.concatFunc
		buf1        = arg.byteField1
		buf2        = arg.byteField2
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := hashFunc(concatFunc(buf1[i], buf1[i+1]))
		if err != nil {
			return err
		}
		buf2[i>>1] = newHash
	}
	return nil
}

// fixOdd fixes the odd-length slice by appending a node to it.
// If NoDuplicates is true, append a node by duplicating the previous node.
// Otherwise, append a node by random.
func (m *MerkleTree) fixOdd(buf [][]byte, prevLen int) ([][]byte, int, error) {
	if prevLen&1 == 0 {
		return buf, prevLen, nil
	}
	var appendNode []byte
	if m.NoDuplicates {
		var err error
		if appendNode, err = dummyHash(); err != nil {
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
		m.updatePairProofs(buf, i, batch, step)
	}
}

// updateProofHandler updates the proofs in parallel.
func updateProofHandler(arg argType) error {
	var (
		mt          = arg.mt // The Merkle Tree instance
		buf         = arg.byteField1
		start       = arg.intField1
		batch       = arg.intField2
		step        = arg.intField3
		bufLen      = arg.intField4
		numRoutines = arg.intField5
	)
	for i := start; i < bufLen; i += numRoutines << 1 {
		mt.updatePairProofs(buf, i, batch, step)
	}
	// return the nil error to be compatible with the handler type
	return nil
}

func (m *MerkleTree) updateProofsParallel(buf [][]byte, bufLen, step int) {
	batch := 1 << step
	numRoutines := m.NumRoutines
	if numRoutines > bufLen {
		numRoutines = bufLen
	}
	argList := make([]argType, numRoutines)
	for i := 0; i < numRoutines; i++ {
		argList[i] = argType{
			mt:         m,
			byteField1: buf,
			intField1:  i << 1, // starting index
			intField2:  batch,
			intField3:  step,
			intField4:  bufLen,
			intField5:  numRoutines,
		}
	}
	wp.Map(updateProofHandler, argList)
}

func (m *MerkleTree) updatePairProofs(buf [][]byte, idx, batch, step int) {
	start := idx * batch
	end := min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Path += 1 << step
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buf[idx+1])
	}
	start += batch
	end = min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buf[idx])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dummyHash generates a dummy hash to make odd-length buffer even.
func dummyHash() ([]byte, error) {
	dummyBytes := make([]byte, defaultHashLen)
	if _, err := rand.Read(dummyBytes); err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

func (m *MerkleTree) leafGen(blocks []DataBlock) ([][]byte, error) {
	var (
		leaves = make([][]byte, m.NumLeaves)
		err    error
	)
	for i := 0; i < m.NumLeaves; i++ {
		if leaves[i], err = leafFromBlock(blocks[i], m.Config); err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

func leafFromBlock(block DataBlock, config *Config) ([]byte, error) {
	blockBytes, err := block.Serialize()
	if err != nil {
		return nil, err
	}
	if config.DisableLeafHashing {
		// copy the value so that the original byte slice is not modified
		leaf := make([]byte, len(blockBytes))
		copy(leaf, blockBytes)
		return leaf, nil
	}
	return config.HashFunc(blockBytes)
}

// leafGenHandler generates the leaves in parallel.
func leafGenHandler(arg argType) error {
	var (
		blocks      = arg.dataBlockField
		leaves      = arg.byteField1
		start       = arg.intField1
		lenLeaves   = arg.intField2
		numRoutines = arg.intField3
	)
	var err error
	for i := start; i < lenLeaves; i += numRoutines {
		if leaves[i], err = leafFromBlock(blocks[i], arg.mt.Config); err != nil {
			return err
		}
	}
	return nil
}

func (m *MerkleTree) leafGenParallel(blocks []DataBlock) ([][]byte, error) {
	var (
		lenLeaves   = len(blocks)
		leaves      = make([][]byte, lenLeaves)
		numRoutines = m.NumRoutines
	)
	if numRoutines > lenLeaves {
		numRoutines = lenLeaves
	}
	argList := make([]argType, numRoutines)
	for i := 0; i < numRoutines; i++ {
		argList[i] = argType{
			mt:             m, // The Merkle Tree instance
			dataBlockField: blocks,
			byteField1:     leaves,
			intField1:      i, // starting index
			intField2:      lenLeaves,
			intField3:      numRoutines,
		}
	}
	errList := wp.Map(leafGenHandler, argList)
	for _, err := range errList {
		if err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

func (m *MerkleTree) treeBuild() (err error) {
	finishMap := make(chan struct{})
	go func() {
		for i := 0; i < m.NumLeaves; i++ {
			m.leafMap.Store(string(m.Leaves[i]), i)
		}
		finishMap <- struct{}{} // empty channel to serve as a wait group for map generation
	}()
	m.nodes = make([][][]byte, m.Depth)
	m.nodes[0] = make([][]byte, m.NumLeaves)
	copy(m.nodes[0], m.Leaves)
	var prevLen int
	if m.nodes[0], prevLen, err = m.fixOdd(m.nodes[0], m.NumLeaves); err != nil {
		return
	}
	if m.RunInParallel {
		if err := m.computeTreeNodeParallel(prevLen); err != nil {
			return err
		}
	}
	for i := uint32(0); i < m.Depth-1; i++ {
		m.nodes[i+1] = make([][]byte, prevLen>>1)
		for j := 0; j < prevLen; j += 2 {
			if m.nodes[i+1][j>>1], err = m.HashFunc(
				m.concatFunc(m.nodes[i][j], m.nodes[i][j+1]),
			); err != nil {
				return
			}
		}
		if m.nodes[i+1], prevLen, err = m.fixOdd(m.nodes[i+1], len(m.nodes[i+1])); err != nil {
			return
		}
	}
	if m.Root, err = m.HashFunc(m.concatFunc(
		m.nodes[m.Depth-1][0], m.nodes[m.Depth-1][1],
	)); err != nil {
		return
	}
	<-finishMap
	return
}

func (m *MerkleTree) computeTreeNodeParallel(prevLen int) error {
	for i := uint32(0); i < m.Depth-1; i++ {
		m.nodes[i+1] = make([][]byte, prevLen>>1)
		numRoutines := m.NumRoutines
		if numRoutines > prevLen {
			numRoutines = prevLen
		}
		argList := make([]argType, numRoutines)
		for j := 0; j < numRoutines; j++ {
			argList[j] = argType{
				mt:          m,
				intField1:   j << 1, // starting index
				intField2:   prevLen,
				intField3:   m.NumRoutines,
				uint32Field: i, // tree depth
			}
		}
		errList := wp.Map(treeBuildHandler, argList)
		for _, err := range errList {
			if err != nil {
				return err
			}
		}
		var err error
		if m.nodes[i+1], prevLen, err = m.fixOdd(m.nodes[i+1], len(m.nodes[i+1])); err != nil {
			return err
		}
	}
	return nil
}

// treeBuildHandler builds the tree in parallel.
func treeBuildHandler(arg argType) error {
	var (
		mt          = arg.mt // the Merkle Tree instance
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
		depth       = arg.uint32Field
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := mt.HashFunc(mt.concatFunc(mt.nodes[depth][i], mt.nodes[depth][i+1]))
		if err != nil {
			return err
		}
		mt.nodes[depth+1][i>>1] = newHash
	}
	return nil
}

// Verify verifies the data block with the Merkle Tree proof
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.Config)
}

// Verify verifies the data block with the Merkle Tree proof and Merkle root hash
func Verify(dataBlock DataBlock, proof *Proof, root []byte, config *Config) (bool, error) {
	if dataBlock == nil {
		return false, errors.New("data block is nil")
	}
	if proof == nil {
		return false, errors.New("proof is nil")
	}
	if config == nil {
		config = new(Config)
	}
	if config.HashFunc == nil {
		config.HashFunc = defaultHashFunc
	}
	if config.concatFunc == nil {
		config.concatFunc = concatHash
	}
	leaf, err := leafFromBlock(dataBlock, config)
	if err != nil {
		return false, err
	}
	// Copy the slice so that the original leaf won't be modified.
	result := make([]byte, len(leaf))
	copy(result, leaf)
	path := proof.Path
	for _, sib := range proof.Siblings {
		if path&1 == 1 {
			if result, err = config.HashFunc(config.concatFunc(result, sib)); err != nil {
				return false, err
			}
		} else {
			if result, err = config.HashFunc(config.concatFunc(sib, result)); err != nil {
				return false, err
			}
		}
		path >>= 1
	}
	return bytes.Equal(result, root), nil
}

// Proof generates the Merkle proof for a data block with the Merkle Tree structure generated beforehand.
// The method is only available when the configuration mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
// In ModeProofGen, proofs for all the data blocks are already generated, and the Merkle Tree structure is not cached.
func (m *MerkleTree) Proof(dataBlock DataBlock) (*Proof, error) {
	if m.Mode != ModeTreeBuild && m.Mode != ModeProofGenAndTreeBuild {
		return nil, errors.New("merkle Tree is not in built, could not generate proof by this method")
	}
	leaf, err := leafFromBlock(dataBlock, m.Config)
	if err != nil {
		return nil, err
	}
	val, ok := m.leafMap.Load(string(leaf))
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
