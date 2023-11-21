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
	"bytes"
	"errors"
	"math/bits"
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
)

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

// DataBlock is the interface for input data blocks used to generate the Merkle Tree.
// Implementations of DataBlock should provide a serialization method
// that converts the data block into a byte slice for hashing purposes.
type DataBlock interface {
	// Serialize converts the data block into a byte slice.
	// It returns the serialized byte slice and an error, if any occurs during the serialization process.
	Serialize() ([]byte, error)
}

// workerArgs is used as the arguments for the worker functions when performing parallel computations.
// Each worker function has its own dedicated argument struct embedded within workerArgs,
// which eliminates the need for interface conversion overhead and provides clear separation of concerns.
type workerArgs struct {
	generateProofs   *workerArgsGenerateProofs
	updateProofs     *workerArgsUpdateProofs
	generateLeaves   *workerArgsGenerateLeaves
	computeTreeNodes *workerArgsComputeTreeNodes
}

// TypeConfigMode is the type in the Merkle Tree configuration indicating what operations are performed.
type TypeConfigMode int

// TypeHashFunc is the signature of the hash functions used for Merkle Tree generation.
type TypeHashFunc func([]byte) ([]byte, error)

type typeConcatHashFunc func([]byte, []byte) []byte

// Config is the configuration of Merkle Tree.
type Config struct {
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
	// SortSiblingPairs is the parameter for OpenZeppelin compatibility.
	// If set to `true`, the hashing sibling pairs are sorted.
	SortSiblingPairs bool
	// If true, the leaf nodes are NOT hashed before being added to the Merkle Tree.
	DisableLeafHashing bool
}

// MerkleTree implements the Merkle Tree data structure.
type MerkleTree struct {
	Config
	// leafMap maps the data (converted to string) of each leaf node to its index in the Tree slice.
	// It is only available when the configuration mode is set to ModeTreeBuild or ModeProofGenAndTreeBuild.
	leafMap map[string]int
	// leafMapMu is a mutex that protects concurrent access to the leafMap.
	leafMapMu sync.Mutex
	// wp is the worker pool used for parallel computation in the tree building process.
	wp *gool.Pool[workerArgs, error]
	// concatHashFunc is the function for concatenating two hashes.
	// If SortSiblingPairs in Config is true, then the sibling pairs are first sorted and then concatenated,
	// supporting the OpenZeppelin Merkle Tree protocol.
	// Otherwise, the sibling pairs are concatenated directly.
	concatHashFunc typeConcatHashFunc
	// nodes contains the Merkle Tree's internal node structure.
	// It is only available when the configuration mode is set to ModeTreeBuild or ModeProofGenAndTreeBuild.
	nodes [][][]byte
	// Root is the hash of the Merkle root node.
	Root []byte
	// Leaves are the hashes of the data blocks that form the Merkle Tree's leaves.
	// These hashes are used to generate the tree structure.
	// If the DisableLeafHashing configuration is set to true, the original data blocks are used as the leaves.
	Leaves [][]byte
	// Proofs are the proofs to the data blocks generated during the tree building process.
	Proofs []*Proof
	// Depth is the depth of the Merkle Tree.
	Depth int
	// NumLeaves is the number of leaves in the Merkle Tree.
	// This value is fixed once the tree is built.
	NumLeaves int
}

// Proof represents a Merkle Tree proof.
type Proof struct {
	Siblings [][]byte // Sibling nodes to the Merkle Tree path of the data block.
	Path     uint32   // Path variable indicating whether the neighbor is on the left or right.
}

// New generates a new Merkle Tree with the specified configuration and data blocks.
func New(config *Config, blocks []DataBlock) (m *MerkleTree, err error) {
	// Check if there are enough data blocks to build the tree.
	if len(blocks) <= 1 {
		return nil, ErrInvalidNumOfDataBlocks
	}

	// Initialize the configuration if it is not provided.
	if config == nil {
		config = new(Config)
	}

	// Create a MerkleTree with the provided configuration.
	m = &MerkleTree{
		Config:    *config,
		NumLeaves: len(blocks),
		Depth:     bits.Len(uint(len(blocks) - 1)),
	}

	// Initialize the hash function.
	if m.HashFunc == nil {
		if m.RunInParallel {
			// Use a concurrent safe hash function for parallel execution.
			m.HashFunc = DefaultHashFuncParallel
		} else {
			m.HashFunc = DefaultHashFunc
		}
	}

	// Hash concatenation function initialization.
	if m.concatHashFunc == nil {
		if m.SortSiblingPairs {
			m.concatHashFunc = concatSortHash
		} else {
			m.concatHashFunc = concatHash
		}
	}

	// Configure parallelization settings.
	if m.RunInParallel {
		// Set NumRoutines to the number of CPU cores if not specified or invalid.
		if m.NumRoutines <= 0 {
			m.NumRoutines = runtime.NumCPU()
		}
		// Initialize a wait group for parallel computation and generate leaves.
		// Task channel capacity is passed as 0, so use the default value: 2 * numWorkers.
		m.wp = gool.NewPool[workerArgs, error](m.NumRoutines, 0)
		defer m.wp.Close()
		if m.Leaves, err = m.generateLeavesInParallel(blocks); err != nil {
			return nil, err
		}
	} else {
		// Generate leaves without parallelization.
		if m.Leaves, err = m.generateLeaves(blocks); err != nil {
			return nil, err
		}
	}

	// Perform actions based on the configured mode.
	// Set the mode to ModeProofGen by default if not specified.
	if m.Mode == 0 {
		m.Mode = ModeProofGen
	}

	// Generate proofs in ModeProofGen.
	if m.Mode == ModeProofGen {
		err = m.generateProofs()
		return
	}
	// Initialize the leafMap for ModeTreeBuild and ModeProofGenAndTreeBuild.
	m.leafMap = make(map[string]int)

	// Build the tree in ModeTreeBuild.
	if m.Mode == ModeTreeBuild {
		err = m.buildTree()
		return
	}

	// Build the tree and generate proofs in ModeProofGenAndTreeBuild.
	if m.Mode == ModeProofGenAndTreeBuild {
		if err = m.buildTree(); err != nil {
			return
		}
		m.initProofs()
		if m.RunInParallel {
			for i := 0; i < len(m.nodes); i++ {
				m.updateProofsInParallel(m.nodes[i], len(m.nodes[i]), i)
			}
			return
		}
		for i := 0; i < len(m.nodes); i++ {
			m.updateProofs(m.nodes[i], len(m.nodes[i]), i)
		}
		return
	}

	// Return an error if the configuration mode is invalid.
	return nil, ErrInvalidConfigMode
}

// concatSortHash concatenates two byte slices, b1 and b2, in a sorted order.
func concatHash(b1 []byte, b2 []byte) []byte {
	result := make([]byte, len(b1)+len(b2))
	copy(result, b1)
	copy(result[len(b1):], b2)
	return result
}

// concatSortHash concatenates two byte slices, b1 and b2, in a sorted order.
// The function ensures that the smaller byte slice (in terms of lexicographic order)
// is placed before the larger one. This is used for compatibility with OpenZeppelin's
// Merkle Proof verification implementation.
func concatSortHash(b1 []byte, b2 []byte) []byte {
	if bytes.Compare(b1, b2) < 0 {
		return concatHash(b1, b2)
	}
	return concatHash(b2, b1)
}

// initProofs initializes the MerkleTree's Proofs with the appropriate size and depth.
func (m *MerkleTree) initProofs() {
	m.Proofs = make([]*Proof, m.NumLeaves)
	for i := 0; i < m.NumLeaves; i++ {
		m.Proofs[i] = new(Proof)
		m.Proofs[i].Siblings = make([][]byte, 0, m.Depth)
	}
}

// generateProofs constructs the Merkle Tree and generates the Merkle proofs for each leaf.
// It returns an error if there is an issue during the generation process.
func (m *MerkleTree) generateProofs() error {
	m.initProofs()
	buffer := make([][]byte, m.NumLeaves)
	copy(buffer, m.Leaves)
	var bufferLength int
	buffer, bufferLength = m.fixOddLength(buffer, m.NumLeaves)

	if m.RunInParallel {
		return m.generateProofsInParallel(buffer, bufferLength)
	}

	m.updateProofs(buffer, m.NumLeaves, 0)
	var err error
	for step := 1; step < m.Depth; step++ {
		for idx := 0; idx < bufferLength; idx += 2 {
			buffer[idx>>1], err = m.HashFunc(m.concatHashFunc(buffer[idx], buffer[idx+1]))
			if err != nil {
				return err
			}
		}
		bufferLength >>= 1
		buffer, bufferLength = m.fixOddLength(buffer, bufferLength)
		m.updateProofs(buffer, bufferLength, step)
	}

	m.Root, err = m.HashFunc(m.concatHashFunc(buffer[0], buffer[1]))
	return err
}

// workerArgsGenerateProofs contains the parameters required for workerGenerateProofs.
type workerArgsGenerateProofs struct {
	hashFunc       TypeHashFunc
	concatHashFunc typeConcatHashFunc
	buffer         [][]byte
	tempBuffer     [][]byte
	startIdx       int
	bufferLength   int
	numRoutines    int
}

// workerGenerateProofs is the worker function that generates Merkle proofs in parallel.
// It processes a portion of the buffer based on the provided worker arguments.
func workerGenerateProofs(args workerArgs) error {
	chosenArgs := args.generateProofs
	var (
		hashFunc     = chosenArgs.hashFunc
		concatFunc   = chosenArgs.concatHashFunc
		buffer       = chosenArgs.buffer
		tempBuffer   = chosenArgs.tempBuffer
		startIdx     = chosenArgs.startIdx
		bufferLength = chosenArgs.bufferLength
		numRoutines  = chosenArgs.numRoutines
	)
	for i := startIdx; i < bufferLength; i += numRoutines << 1 {
		newHash, err := hashFunc(concatFunc(buffer[i], buffer[i+1]))
		if err != nil {
			return err
		}
		tempBuffer[i>>1] = newHash
	}
	return nil
}

// generateProofsInParallel generates proofs concurrently for the MerkleTree.
func (m *MerkleTree) generateProofsInParallel(buffer [][]byte, bufferLength int) (err error) {
	tempBuffer := make([][]byte, bufferLength>>1)
	m.updateProofsInParallel(buffer, m.NumLeaves, 0)
	numRoutines := m.NumRoutines
	for step := 1; step < m.Depth; step++ {
		// Limit the number of workers to the previous level length.
		if numRoutines > bufferLength {
			numRoutines = bufferLength
		}

		// Create the list of arguments for the worker pool.
		argList := make([]workerArgs, numRoutines)
		for i := 0; i < numRoutines; i++ {
			argList[i] = workerArgs{
				generateProofs: &workerArgsGenerateProofs{
					hashFunc:       m.HashFunc,
					concatHashFunc: m.concatHashFunc,
					buffer:         buffer,
					tempBuffer:     tempBuffer,
					startIdx:       i << 1,
					bufferLength:   bufferLength,
					numRoutines:    numRoutines,
				},
			}
		}

		// Execute proof generation concurrently using the worker pool.
		errList := m.wp.Map(workerGenerateProofs, argList)
		for _, err = range errList {
			if err != nil {
				return
			}
		}

		// Swap the buffers for the next iteration.
		buffer, tempBuffer = tempBuffer, buffer
		bufferLength >>= 1

		// Fix the buffer if it has an odd number of elements.
		buffer, bufferLength = m.fixOddLength(buffer, bufferLength)

		// Update the proofs with the new buffer.
		m.updateProofsInParallel(buffer, bufferLength, step)
	}

	// Compute the root hash of the Merkle tree.
	m.Root, err = m.HashFunc(m.concatHashFunc(buffer[0], buffer[1]))
	return
}

// fixOddLength adjusts the buffer for odd-length slices by appending a node.
func (m *MerkleTree) fixOddLength(buffer [][]byte, bufferLength int) ([][]byte, int) {
	// If the buffer length is even, no adjustment is needed.
	if bufferLength&1 == 0 {
		return buffer, bufferLength
	}

	// Determine the node to append.
	appendNode := buffer[bufferLength-1]
	bufferLength++

	// Append the node to the buffer, either by extending the buffer or updating an existing entry.
	if len(buffer) < bufferLength {
		buffer = append(buffer, appendNode)
	} else {
		buffer[bufferLength-1] = appendNode
	}

	return buffer, bufferLength
}

func (m *MerkleTree) updateProofs(buffer [][]byte, bufferLength, step int) {
	batch := 1 << step
	for i := 0; i < bufferLength; i += 2 {
		m.updateProofPairs(buffer, i, batch, step)
	}
}

// workerArgsUpdateProofs contains arguments for the workerUpdateProofs function.
type workerArgsUpdateProofs struct {
	tree         *MerkleTree
	buffer       [][]byte
	startIdx     int
	batch        int
	step         int
	bufferLength int
	numRoutines  int
}

// workerUpdateProofs is the worker function that updates Merkle proofs in parallel.
func workerUpdateProofs(args workerArgs) error {
	chosenArgs := args.updateProofs
	var (
		tree         = chosenArgs.tree
		buffer       = chosenArgs.buffer
		startIdx     = chosenArgs.startIdx
		batch        = chosenArgs.batch
		step         = chosenArgs.step
		bufferLength = chosenArgs.bufferLength
		numRoutines  = chosenArgs.numRoutines
	)
	for i := startIdx; i < bufferLength; i += numRoutines << 1 {
		tree.updateProofPairs(buffer, i, batch, step)
	}
	// return the nil error to be compatible with the worker type
	return nil
}

// updateProofsInParallel updates proofs concurrently for the Merkle Tree.
func (m *MerkleTree) updateProofsInParallel(buffer [][]byte, bufferLength, step int) {
	batch := 1 << step
	numRoutines := m.NumRoutines
	if numRoutines > bufferLength {
		numRoutines = bufferLength
	}
	argList := make([]workerArgs, numRoutines)
	for i := 0; i < numRoutines; i++ {
		argList[i] = workerArgs{
			updateProofs: &workerArgsUpdateProofs{
				tree:         m,
				buffer:       buffer,
				startIdx:     i << 1,
				batch:        batch,
				step:         step,
				bufferLength: bufferLength,
				numRoutines:  numRoutines,
			},
		}
	}
	m.wp.Map(workerUpdateProofs, argList)
}

// updateProofPairs updates the proofs in the Merkle Tree in pairs.
func (m *MerkleTree) updateProofPairs(buffer [][]byte, idx, batch, step int) {
	start := idx * batch
	end := min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Path += 1 << step
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buffer[idx+1])
	}
	start += batch
	end = min(start+batch, len(m.Proofs))
	for i := start; i < end; i++ {
		m.Proofs[i].Siblings = append(m.Proofs[i].Siblings, buffer[idx])
	}
}

// generateLeaves generates the leaves slice from the data blocks.
func (m *MerkleTree) generateLeaves(blocks []DataBlock) ([][]byte, error) {
	var (
		leaves = make([][]byte, m.NumLeaves)
		err    error
	)
	for i := 0; i < m.NumLeaves; i++ {
		if leaves[i], err = dataBlockToLeaf(blocks[i], &m.Config); err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

// dataBlockToLeaf generates the leaf from the data block.
// If the leaf hashing is disabled, the data block is returned as the leaf.
func dataBlockToLeaf(block DataBlock, config *Config) ([]byte, error) {
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

// workerArgsGenerateLeaves contains arguments for the workerGenerateLeaves function.
type workerArgsGenerateLeaves struct {
	config      *Config
	dataBlocks  []DataBlock
	leaves      [][]byte
	startIdx    int
	lenLeaves   int
	numRoutines int
}

// workerGenerateLeaves is the worker function that generates Merkle leaves in parallel.
func workerGenerateLeaves(args workerArgs) error {
	chosenArgs := args.generateLeaves
	var (
		config      = chosenArgs.config
		blocks      = chosenArgs.dataBlocks
		leaves      = chosenArgs.leaves
		start       = chosenArgs.startIdx
		lenLeaves   = chosenArgs.lenLeaves
		numRoutines = chosenArgs.numRoutines
	)
	var err error
	for i := start; i < lenLeaves; i += numRoutines {
		if leaves[i], err = dataBlockToLeaf(blocks[i], config); err != nil {
			return err
		}
	}
	return nil
}

// generateLeavesInParallel generates the leaves slice from the data blocks in parallel.
func (m *MerkleTree) generateLeavesInParallel(blocks []DataBlock) ([][]byte, error) {
	var (
		lenLeaves   = len(blocks)
		leaves      = make([][]byte, lenLeaves)
		numRoutines = m.NumRoutines
	)
	if numRoutines > lenLeaves {
		numRoutines = lenLeaves
	}
	argList := make([]workerArgs, numRoutines)
	for i := 0; i < numRoutines; i++ {
		argList[i] = workerArgs{
			generateLeaves: &workerArgsGenerateLeaves{
				config:      &m.Config,
				dataBlocks:  blocks,
				leaves:      leaves,
				startIdx:    i,
				lenLeaves:   lenLeaves,
				numRoutines: numRoutines,
			},
		}
	}
	errList := m.wp.Map(workerGenerateLeaves, argList)
	for _, err := range errList {
		if err != nil {
			return nil, err
		}
	}
	return leaves, nil
}

// buildTree builds the Merkle Tree.
func (m *MerkleTree) buildTree() (err error) {
	finishMap := make(chan struct{})
	go func() {
		m.leafMapMu.Lock()
		defer m.leafMapMu.Unlock()
		for i := 0; i < m.NumLeaves; i++ {
			m.leafMap[string(m.Leaves[i])] = i
		}
		finishMap <- struct{}{} // empty channel to serve as a wait group for map generation
	}()
	m.nodes = make([][][]byte, m.Depth)
	m.nodes[0] = make([][]byte, m.NumLeaves)
	copy(m.nodes[0], m.Leaves)
	var bufferLength int
	m.nodes[0], bufferLength = m.fixOddLength(m.nodes[0], m.NumLeaves)
	if m.RunInParallel {
		if err := m.computeTreeNodesInParallel(bufferLength); err != nil {
			return err
		}
	}
	for i := 0; i < m.Depth-1; i++ {
		m.nodes[i+1] = make([][]byte, bufferLength>>1)
		for j := 0; j < bufferLength; j += 2 {
			if m.nodes[i+1][j>>1], err = m.HashFunc(
				m.concatHashFunc(m.nodes[i][j], m.nodes[i][j+1]),
			); err != nil {
				return
			}
		}
		m.nodes[i+1], bufferLength = m.fixOddLength(m.nodes[i+1], len(m.nodes[i+1]))
	}
	if m.Root, err = m.HashFunc(m.concatHashFunc(
		m.nodes[m.Depth-1][0], m.nodes[m.Depth-1][1],
	)); err != nil {
		return
	}
	<-finishMap
	return
}

// workerArgsComputeTreeNodes contains arguments for the workerComputeTreeNodes function.
type workerArgsComputeTreeNodes struct {
	tree         *MerkleTree
	startIdx     int
	bufferLength int
	numRoutines  int
	depth        int
}

// workerBuildTree is the worker function that builds the Merkle tree in parallel.
func workerBuildTree(args workerArgs) error {
	chosenArgs := args.computeTreeNodes
	var (
		tree         = chosenArgs.tree
		start        = chosenArgs.startIdx
		bufferLength = chosenArgs.bufferLength
		numRoutines  = chosenArgs.numRoutines
		depth        = chosenArgs.depth
	)
	for i := start; i < bufferLength; i += numRoutines << 1 {
		newHash, err := tree.HashFunc(tree.concatHashFunc(
			tree.nodes[depth][i], tree.nodes[depth][i+1],
		))
		if err != nil {
			return err
		}
		tree.nodes[depth+1][i>>1] = newHash
	}
	return nil
}

// computeTreeNodesInParallel computes the tree nodes in parallel.
func (m *MerkleTree) computeTreeNodesInParallel(bufferLength int) error {
	for i := 0; i < m.Depth-1; i++ {
		m.nodes[i+1] = make([][]byte, bufferLength>>1)
		numRoutines := m.NumRoutines
		if numRoutines > bufferLength {
			numRoutines = bufferLength
		}
		argList := make([]workerArgs, numRoutines)
		for j := 0; j < numRoutines; j++ {
			argList[j] = workerArgs{
				computeTreeNodes: &workerArgsComputeTreeNodes{
					tree:         m,
					startIdx:     j << 1,
					bufferLength: bufferLength,
					numRoutines:  m.NumRoutines,
					depth:        i,
				},
			}
		}
		errList := m.wp.Map(workerBuildTree, argList)
		for _, err := range errList {
			if err != nil {
				return err
			}
		}
		m.nodes[i+1], bufferLength = m.fixOddLength(m.nodes[i+1], len(m.nodes[i+1]))
	}
	return nil
}

// Verify checks if the data block is valid using the Merkle Tree proof and the cached Merkle root hash.
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, &m.Config)
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
	leaf, err := dataBlockToLeaf(dataBlock, config)
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

// Proof generates the Merkle proof for a data block using the previously generated Merkle Tree structure.
// This method is only available when the configuration mode is ModeTreeBuild or ModeProofGenAndTreeBuild.
// In ModeProofGen, proofs for all the data blocks are already generated, and the Merkle Tree structure
// is not cached.
func (m *MerkleTree) Proof(dataBlock DataBlock) (*Proof, error) {
	if m.Mode != ModeTreeBuild && m.Mode != ModeProofGenAndTreeBuild {
		return nil, ErrProofInvalidModeTreeNotBuilt
	}

	// Convert the data block to a leaf.
	leaf, err := dataBlockToLeaf(dataBlock, &m.Config)
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
