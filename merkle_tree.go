package merkletree

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	// Default hash result length using SHA256.
	defaultHashLen          = 32
	ModeProofGen   ModeType = iota
	ModeBuildTree
	ModeProofGenAndBuildTree
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
	NumRoutines int
	// If true, then the odd node situation is handled by duplicating the previous node.
	// Otherwise, generate a dummy node with random hash value.
	AllowDuplicates bool
	// Mode of the Merkle Tree generation.
	Mode ModeType
}

// MerkleTree implements the Merkle Tree structure
type MerkleTree struct {
	*Config                // Merkle Tree configuration
	Root    []byte         // Merkle root hash
	Leaves  [][]byte       // Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation
	Proofs  []*Proof       // proofs to the data blocks generated during the tree building process
	Depth   uint32         // the Merkle Tree depth
	Tree    [][][]byte     // the Merkle Tree, only available when config mode is ModeBuildTree or ModeProofGenAndBuildTree
	leafMap map[string]int // map of the leaf hash to the index in the Tree slice, only available when config mode is ModeBuildTree or ModeProofGenAndBuildTree
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
	m = &MerkleTree{
		Config: config,
	}
	m.Depth = calTreeDepth(len(blocks))
	if m.Mode == ModeBuildTree {
		if m.RunInParallel {
			panic("not implemented")
		}
		m.Leaves, err = leafGenParal(blocks, m.HashFunc, m.NumRoutines)
		if err != nil {
			return
		}
		err = m.buildTree()
		return
	}
	if m.Mode == ModeProofGenAndBuildTree {
		panic("not implemented")
	}
	// ModeProofGen by default
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
	var (
		step, prevLen int
	)
	buf := make([][]byte, numLeaves)
	copy(buf, m.Leaves)
	buf, prevLen, err = m.fixOdd(buf, numLeaves)
	if err != nil {
		return
	}
	m.updateProofs(buf, numLeaves, 0)
	for {
		for idx := 0; idx < prevLen; idx += 2 {
			buf[idx/2], err = m.HashFunc(append(buf[idx], buf[idx+1]...))
			if err != nil {
				return
			}
		}
		prevLen /= 2
		if prevLen == 1 {
			break
		}
		buf, prevLen, err = m.fixOdd(buf, prevLen)
		if err != nil {
			return
		}
		step++
		m.updateProofs(buf, prevLen, step)
	}
	m.Root = buf[0]
	return
}

func (m *MerkleTree) proofGenParal() (err error) {
	numRoutines := m.NumRoutines
	numLeaves := len(m.Leaves)
	m.Proofs = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proofs[i] = new(Proof)
	}
	var (
		step, prevLen int
	)
	buf1 := make([][]byte, numLeaves)
	copy(buf1, m.Leaves)
	buf1, prevLen, err = m.fixOdd(buf1, numLeaves)
	if err != nil {
		return
	}
	buf2 := make([][]byte, prevLen/2)
	m.updateProofsParal(buf1, numLeaves, 0)
	for {
		if err != nil {
			return
		}
		g := new(errgroup.Group)
		for i := 0; i < numRoutines && i < prevLen; i++ {
			idx := 2 * i
			g.Go(func() error {
				for j := idx; j < prevLen; j += 2 * numRoutines {
					newHash, err := m.HashFunc(append(buf1[j], buf1[j+1]...))
					if err != nil {
						return err
					}
					buf2[j/2] = newHash
				}
				return nil
			})
		}
		if err = g.Wait(); err != nil {
			return
		}
		buf1, buf2 = buf2, buf1
		prevLen /= 2
		if prevLen == 1 {
			break
		}
		buf1, prevLen, err = m.fixOdd(buf1, prevLen)
		if err != nil {
			return
		}
		step++
		m.updateProofsParal(buf1, prevLen, step)
	}
	m.Root = buf1[0]
	return
}

// if the length of the buffer calculating the Merkle Tree is odd, then append a node to the buffer
// if AllowDuplicates is true, append a node by duplicating the previous node
// otherwise, append a node by random
func (m *MerkleTree) fixOdd(buf [][]byte, prevLen int) ([][]byte, int, error) {
	if prevLen%2 == 1 {
		var appendNode []byte
		if m.AllowDuplicates {
			appendNode = buf[prevLen-1]
		} else {
			var err error
			appendNode, err = getDummyHash()
			if err != nil {
				return nil, 0, err
			}
		}
		if len(buf) <= prevLen+1 {
			buf = append(buf, appendNode)
		} else {
			buf[prevLen] = appendNode
		}
		prevLen++
	}
	return buf, prevLen, nil
}

func (m *MerkleTree) updateProofs(buf [][]byte, bufLen, step int) {
	if bufLen < 2 {
		return
	}
	batch := 1 << step
	for i := 0; i < bufLen; i += 2 {
		m.updatePairProof(buf, bufLen, i, batch, step)
	}
}

func (m *MerkleTree) updateProofsParal(buf [][]byte, bufLen, step int) {
	numRoutines := m.NumRoutines
	if bufLen < 2 {
		return
	}
	batch := 1 << step
	wg := new(sync.WaitGroup)
	for i := 0; i < numRoutines; i++ {
		idx := 2 * i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := idx; j < bufLen; j += 2 * numRoutines {
				m.updatePairProof(buf, bufLen, j, batch, step)
			}
		}()
	}
	wg.Wait()
}

func (m *MerkleTree) updatePairProof(buf [][]byte, bufLen, idx, batch, step int) {
	if bufLen < 2 {
		return
	}
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

func (m *MerkleTree) buildTree() (err error) {
	numLeaves := len(m.Leaves)
	m.leafMap = make(map[string]int)
	for i := 0; i < numLeaves; i++ {
		m.leafMap[string(m.Leaves[i])] = i
	}
	m.Tree = make([][][]byte, m.Depth+1)
	m.Tree[m.Depth] = make([][]byte, numLeaves)
	copy(m.Tree[m.Depth], m.Leaves)
	var prevLen int
	m.Tree[m.Depth], prevLen, err = m.fixOdd(m.Tree[m.Depth], numLeaves)
	for i := m.Depth; i > 0; i-- {
		m.Tree[i-1] = make([][]byte, prevLen/2)
		if err != nil {
			return
		}
		for j := 0; j < prevLen; j += 2 {
			appendHash := append(m.Tree[i][j], m.Tree[i][j+1]...)
			m.Tree[i-1][j/2], err = m.HashFunc(appendHash)
			if err != nil {
				return
			}
		}
		prevLen /= 2
		m.Tree[i-1], prevLen, err = m.fixOdd(m.Tree[i-1], prevLen)
	}
	m.Root = m.Tree[0][0]
	return
}

// Verify verifies the data block with the Merkle Tree proof
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.HashFunc)
}

// Verify verifies the data block with the Merkle Tree proof and Merkle root hash
func Verify(dataBlock DataBlock, proof *Proof, root []byte, hashFunc HashFuncType) (bool, error) {
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
