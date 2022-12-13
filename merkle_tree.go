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
	"errors"
	"math"
	"runtime"
	"sync"

	"github.com/txaty/gool"
)

const (
	// ModeProofGen is the proof generation configuration mode.
	ModeProofGen ModeType = iota
	// ModeTreeBuild is the tree building configuration mode.
	ModeTreeBuild
	// ModeProofGenAndTreeBuild is the proof generation and tree building configuration mode.
	ModeProofGenAndTreeBuild
	// Default hash result length using SHA256.
	defaultHashLen = 32
)

// ModeType is the type in the Merkle Tree configuration indicating what operations are performed.
type ModeType int

// DataBlock is the interface of input data blocks to generate the Merkle Tree.
type DataBlock interface {
	Serialize() ([]byte, error)
}

type concatHashFuncType func([]byte, []byte) []byte

// HashFuncType is the signature of the hash functions used for Merkle Tree generation.
type HashFuncType func([]byte) ([]byte, error)

// Config is the configuration of Merkle Tree.
type Config struct {
	// appendFunc is the function for concatenating two hashes.
	// If SortSiblingPairs in Config is true, then the sibling pairs are first sorted and then concatenated.
	concatHashFunc concatHashFuncType
	// Customizable hash function used for tree generation.
	HashFunc HashFuncType
	// Number of goroutines run in parallel.
	// If RunInParallel is true and NumRoutine is set to 0, use number of CPU as the number of goroutines.
	NumRoutines int
	// Mode of the Merkle Tree generation.
	Mode ModeType
	// If RunInParallel is true, the generation runs in parallel, otherwise runs without parallelization.
	// This increase the performance for the calculation of large number of data blocks, e.g. over 10,000 blocks.
	RunInParallel bool
	// If true, generate a dummy node with random hash value.
	// Otherwise, then the odd node situation is handled by duplicating the previous node.
	NoDuplicates bool
	// SortSiblingPairs is the parameter for OpenZeppelin compatibility.
	// If set to `true`, the hashing sibling pairs are sorted.
	SortSiblingPairs bool
}

// MerkleTree implements the Merkle Tree structure
type MerkleTree struct {
	*Config
	// leafMap is the map of the leaf hash to the index in the Tree slice,
	// only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	leafMap sync.Map
	// tree is the Merkle Tree structure, only available when config mode is ModeTreeBuild or ModeProofGenAndTreeBuild
	tree [][][]byte
	// Root is the Merkle root hash
	Root []byte
	// Leaves are Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation
	Leaves [][]byte
	// Proofs are proofs to the data blocks generated during the tree building process
	Proofs []*Proof
	// Depth is the Merkle Tree depth
	Depth uint32
}

// Proof implements the Merkle Tree proof.
type Proof struct {
	Siblings [][]byte // sibling nodes to the Merkle Tree path of the data block
	Path     uint32   // path variable indicating whether the neighbor is on the left or right
}

// argType is used as the arguments for the handler functions when performing parallel computations.
// All the handler functions use this universal argument struct to eliminate interface conversion overhead.
// Each field in the struct may be used for different purpose in different handler functions,
// please refer to the comments at each handler function for details.
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

// New generates a new Merkle Tree with specified configuration.
func New(config *Config, blocks []DataBlock) (m *MerkleTree, err error) {
	if len(blocks) <= 1 {
		return nil, errors.New("the number of data blocks must be greater than 1")
	}
	if config == nil {
		config = new(Config)
	}
	m = &MerkleTree{Config: config}
	// hash function initialization
	if m.HashFunc == nil {
		if m.RunInParallel {
			m.HashFunc = defaultHashFuncParal // parallelized hash function must be concurrent safe
		} else {
			m.HashFunc = defaultHashFunc
		}
	}
	// If the configuration mode is not set, then set it to ModeProofGen by default.
	if m.Mode == 0 {
		m.Mode = ModeProofGen
	}
	// If RunInParallel is true and NumRoutines is unset, then set NumRoutines to the number of CPU.
	if m.RunInParallel && m.NumRoutines == 0 {
		m.NumRoutines = runtime.NumCPU()
	}
	// hash concatenation function initialization
	if m.SortSiblingPairs {
		m.concatHashFunc = concatSortHash
	} else {
		m.concatHashFunc = concatHash
	}
	m.Depth = calTreeDepth(len(blocks))
	// generic wait group initialization (for parallelized computation) and leaf generation
	var wp *gool.Pool[argType, error]
	if m.RunInParallel {
		// task channel capacity is passed as 0, so use the default value: 2 * numWorkers
		wp = gool.NewPool[argType, error](m.NumRoutines, 0)
		defer wp.Close()
		m.Leaves, err = m.leafGenParal(blocks, wp)
		if err != nil {
			return
		}
	} else {
		m.Leaves, err = m.leafGen(blocks)
		if err != nil {
			return
		}
	}

	if m.Mode == ModeProofGen {
		if m.RunInParallel {
			err = m.proofGenParal(wp)
			return
		}
		err = m.proofGen()
		return
	}
	if m.Mode == ModeTreeBuild {
		if m.RunInParallel {
			err = m.treeBuild(wp)
			return
		}
		err = m.treeBuild(nil)
		return
	}
	if m.Mode == ModeProofGenAndTreeBuild {
		if m.RunInParallel {
			err = m.treeBuild(wp)
			if err != nil {
				return
			}
			m.initProofs()
			for i := 0; i < len(m.tree); i++ {
				m.updateProofsParal(m.tree[i], len(m.tree[i]), i, wp)
			}
			return
		}
		err = m.treeBuild(nil)
		if err != nil {
			return
		}
		m.initProofs()
		for i := 0; i < len(m.tree); i++ {
			m.updateProofs(m.tree[i], len(m.tree[i]), i)
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

func (m *MerkleTree) initProofs() {
	numLeaves := len(m.Leaves)
	m.Proofs = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
}

func (m *MerkleTree) proofGen() (err error) {
	numLeaves := len(m.Leaves)
	m.initProofs()
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

			buf[idx>>1], err = m.HashFunc(m.Config.concatHashFunc(buf[idx], buf[idx+1]))
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

	m.Root, err = m.HashFunc(m.Config.concatHashFunc(buf[0], buf[1]))
	return
}

// proofGenHandler generates the proofs in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: buf1
//	byteField2: buf2
//	intField1: start
//	intField2: prevLen
//	intField3: numRoutines
//
// return:
//
//	error
func proofGenHandler(arg argType) error {
	var (
		hashFunc    = arg.mt.HashFunc
		buf1        = arg.byteField1
		buf2        = arg.byteField2
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := hashFunc(append(buf1[i], buf1[i+1]...))
		if err != nil {
			return err
		}
		buf2[i>>1] = newHash
	}
	return nil
}

func (m *MerkleTree) proofGenParal(wp *gool.Pool[argType, error]) (err error) {
	m.initProofs()
	numLeaves := len(m.Leaves)
	buf1 := make([][]byte, numLeaves)
	copy(buf1, m.Leaves)
	var prevLen int
	buf1, prevLen, err = m.fixOdd(buf1, numLeaves)
	if err != nil {
		return
	}
	buf2 := make([][]byte, prevLen>>1)
	m.updateProofsParal(buf1, numLeaves, 0, wp)
	numRoutines := m.NumRoutines
	for step := 1; step < int(m.Depth); step++ {
		if numRoutines > prevLen {
			numRoutines = prevLen
		}
		argList := make([]argType, numRoutines)
		for i := 0; i < numRoutines; i++ {
			argList[i] = argType{
				mt:         m,
				byteField1: buf1,
				byteField2: buf2,
				intField1:  i << 1, // starting index
				intField2:  prevLen,
				intField3:  numRoutines,
			}
		}
		errList := wp.Map(proofGenHandler, argList)
		for _, err = range errList {
			if err != nil {
				return err
			}
		}
		buf1, buf2 = buf2, buf1
		prevLen >>= 1
		buf1, prevLen, err = m.fixOdd(buf1, prevLen)
		if err != nil {
			return
		}
		m.updateProofsParal(buf1, prevLen, step, wp)
	}

	m.Root, err = m.HashFunc(m.Config.concatHashFunc(buf1[0], buf1[1]))
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

// updateProofHandler updates the proofs in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: buf
//	intField1: start
//	intField2: batch
//	intField3: step
//	intField4: bufLen
//	intField5: numRoutines
//
// return:
//
//	nothing (nil)
func updateProofHandler(arg argType) error {
	var (
		mt          = arg.mt
		buf         = arg.byteField1
		start       = arg.intField1
		batch       = arg.intField2
		step        = arg.intField3
		bufLen      = arg.intField4
		numRoutines = arg.intField5
	)
	for i := start; i < bufLen; i += numRoutines << 1 {
		mt.updatePairProof(buf, i, batch, step)
	}
	// return the nil error to be compatible with the handler type
	return nil
}

func (m *MerkleTree) updateProofsParal(buf [][]byte, bufLen, step int, wp *gool.Pool[argType, error]) {
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

func (m *MerkleTree) updatePairProof(buf [][]byte, idx, batch, step int) {
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

// generate a dummy hash to make odd-length buffer even
func getDummyHash() ([]byte, error) {
	dummyBytes := make([]byte, defaultHashLen)
	_, err := rand.Read(dummyBytes)
	if err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

func (m *MerkleTree) leafGen(blocks []DataBlock) ([][]byte, error) {
	var (
		lenLeaves = len(blocks)
		leaves    = make([][]byte, lenLeaves)
	)
	for i := 0; i < lenLeaves; i++ {
		data, err := blocks[i].Serialize()
		if err != nil {
			return nil, err
		}
		var hash []byte
		if hash, err = m.HashFunc(data); err != nil {
			return nil, err
		}
		leaves[i] = hash
	}
	return leaves, nil
}

// leafGenHandler generates the leaves in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	byteField1: leaves
//	dataBlockField: blocks
//	intField1: start
//	intField2: lenLeaves
//	intField3: numRoutines
//
// return:
//
//	error
func leafGenHandler(arg argType) error {
	var (
		hashFunc    = arg.mt.HashFunc
		blocks      = arg.dataBlockField
		leaves      = arg.byteField1
		start       = arg.intField1
		lenLeaves   = arg.intField2
		numRoutines = arg.intField3
	)
	for i := start; i < lenLeaves; i += numRoutines {
		data, err := blocks[i].Serialize()
		if err != nil {
			return err
		}
		var hash []byte
		if hash, err = hashFunc(data); err != nil {
			return err
		}
		leaves[i] = hash
	}
	return nil
}

func (m *MerkleTree) leafGenParal(blocks []DataBlock, wp *gool.Pool[argType, error]) ([][]byte, error) {
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
			mt:             m,
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

func (m *MerkleTree) treeBuild(wp *gool.Pool[argType, error]) (err error) {
	numLeaves := len(m.Leaves)
	finishMap := make(chan struct{})
	go func() {
		for i := 0; i < numLeaves; i++ {
			m.leafMap.Store(string(m.Leaves[i]), i)
		}
		finishMap <- struct{}{}
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
		if m.RunInParallel {
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
			for _, err = range errList {
				if err != nil {
					return
				}
			}
		} else {
			for j := 0; j < prevLen; j += 2 {
				m.tree[i+1][j>>1], err = m.HashFunc(
					m.Config.concatHashFunc(m.tree[i][j], m.tree[i][j+1]),
				)
				if err != nil {
					return
				}
			}
		}
		m.tree[i+1], prevLen, err = m.fixOdd(m.tree[i+1], len(m.tree[i+1]))
		if err != nil {
			return
		}
	}

	m.Root, err = m.HashFunc(m.Config.concatHashFunc(m.tree[m.Depth-1][0], m.tree[m.Depth-1][1]))
	if err != nil {
		return
	}
	<-finishMap
	return
}

// treeBuildHandler builds the tree in parallel.
// arg fields:
//
//	mt: the Merkle Tree instance
//	intField1: start
//	intField2: prevLen
//	intField3: numRoutines
//	uint32Field: depth
//
// return:
//
//	error
func treeBuildHandler(arg argType) error {
	var (
		mt          = arg.mt
		start       = arg.intField1
		prevLen     = arg.intField2
		numRoutines = arg.intField3
		depth       = arg.uint32Field
	)
	for i := start; i < prevLen; i += numRoutines << 1 {
		newHash, err := mt.HashFunc(append(mt.tree[depth][i], mt.tree[depth][i+1]...))
		if err != nil {
			return err
		}
		mt.tree[depth+1][i>>1] = newHash
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

	if config.concatHashFunc == nil {
		config.concatHashFunc = concatHash
	}

	var (
		data, err = dataBlock.Serialize()
		hash      []byte
	)

	if err != nil {
		return false, err
	}
	hash, err = config.HashFunc(data)
	if err != nil {
		return false, err
	}
	path := proof.Path
	for _, n := range proof.Siblings {
		if path&1 == 1 {
			hash, err = config.HashFunc(config.concatHashFunc(hash, n))
		} else {
			hash, err = config.HashFunc(config.concatHashFunc(n, hash))
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
	var blockHash []byte
	if blockHash, err = m.HashFunc(blockByte); err != nil {
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
