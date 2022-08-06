package merkletree

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
)

const (
	// default hash result length, using SHA256
	defaultHashLen = 32
)

// DataBlock is the interface of input data blocks to generate the Merkle Tree
type DataBlock interface {
	Serialize() ([]byte, error)
}

// Config is the configuration of Merkle Tree
type Config struct {
	// customizable hash function used for tree generation
	HashFunc func([]byte) ([]byte, error)
	// if true, the generation runs in parallel
	RunInParallel bool
	// number of goroutines run in parallel
	NumRoutines int
	// if true, then the odd node situation is handled by duplicating the previous node
	// otherwise, generate a dummy node with random hash value
	AllowDuplicates bool
}

// MerkleTree implements the Merkle Tree structure
type MerkleTree struct {
	*Config          // Merkle Tree configuration
	Root    []byte   // Merkle root hash
	Leaves  []*Node  // Merkle Tree leaves, i.e. the hashes of the data blocks for tree generation
	Proves  []*Proof // proves to the data blocks generated during the tree building process
}

// Node implements the Merkle Tree node
type Node struct {
	Hash []byte
}

// Proof implements the Merkle Tree proof
type Proof struct {
	Path      uint16   // path variable indicating whether the neighbor is on the left or right
	Neighbors [][]byte // neighbor nodes near the path
}

// NewMerkleTree generates a new Merkle Tree with specified configuration
func NewMerkleTree(config *Config) *MerkleTree {
	if config.HashFunc == nil {
		config.HashFunc = defaultHashFunc
	}
	return &MerkleTree{
		Config: config,
	}
}

// Build builds up the Merkle Tree and generates the proves
func (m *MerkleTree) Build(blocks []DataBlock) (err error) {
	if len(blocks) <= 1 {
		return nil
	}
	if m.RunInParallel {
		m.Leaves, err = generateLeavesParallel()
		if err != nil {
			return err
		}
		// m.Root, err = m.buildTreeParallel()
		panic("not implemented")
	} else {
		m.Leaves, err = generateLeaves(blocks, m.HashFunc)
		if err != nil {
			return err
		}
		m.Root, err = m.buildTree()
	}
	return err
}

func (m *MerkleTree) buildTree() (root []byte, err error) {
	numLeaves := len(m.Leaves)
	m.Proves = make([]*Proof, numLeaves)
	for i := 0; i < numLeaves; i++ {
		m.Proves[i] = new(Proof)
	}
	var (
		step    = 1
		prevLen int
	)
	buf := make([]*Node, numLeaves)
	copy(buf, m.Leaves)
	buf, prevLen, err = m.fixOdd(buf, numLeaves)
	if err != nil {
		return nil, err
	}
	m.assignProves(buf, numLeaves, 0)
	for {
		buf, prevLen, err = m.fixOdd(buf, prevLen)
		if err != nil {
			return nil, err
		}
		for idx := 0; idx < prevLen; idx += 2 {
			appendHash := append(buf[idx].Hash, buf[idx+1].Hash...)
			buf[idx/2].Hash, err = m.HashFunc(appendHash)
		}
		prevLen /= 2
		if prevLen == 1 {
			break
		} else {
			buf, prevLen, err = m.fixOdd(buf, prevLen)
			if err != nil {
				return nil, err
			}
		}
		m.assignProves(buf, prevLen, step)
		step++
	}
	root = buf[0].Hash
	m.Root = root
	return
}

// if the length of the buffer calculating the Merkle Tree is odd, then append a node to the buffer
// if AllowDuplicates is true, append a node by duplicating the previous node
// otherwise, append a node by random
func (m *MerkleTree) fixOdd(buf []*Node, prevLen int) ([]*Node, int, error) {
	if prevLen%2 == 1 {
		var appendNode *Node
		if m.AllowDuplicates {
			appendNode = buf[prevLen-1]
		} else {
			dummyHash, err := getDummyHash()
			if err != nil {
				return nil, 0, err
			}
			appendNode = &Node{
				Hash: dummyHash,
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

func (m *MerkleTree) assignProves(buf []*Node, bufLen, step int) {
	if bufLen < 2 {
		return
	}
	batch := 1 << step
	for i := 0; i < bufLen; i += 2 {
		start := i * batch
		end := start + batch
		if end > len(m.Proves) {
			end = len(m.Proves)
		}
		for j := start; j < end; j++ {
			m.Proves[j].Path += 1 << step
			m.Proves[j].Neighbors = append(m.Proves[j].Neighbors, buf[i+1].Hash)
		}
		start = (i + 1) * batch
		end = start + batch
		if end > len(m.Proves) {
			end = len(m.Proves)
		}
		for j := start; j < end; j++ {
			m.Proves[j].Neighbors = append(m.Proves[j].Neighbors, buf[i].Hash)
		}
	}
}

// func (m *MerkleTree) buildTreeParallel() ([]byte, error) {
// 	panic("not implemented")
// }

// generate a dummy
func getDummyHash() ([]byte, error) {
	dummyBytes := make([]byte, defaultHashLen)
	_, err := rand.Read(dummyBytes)
	if err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

// default hash function using SHA256
func defaultHashFunc(data []byte) ([]byte, error) {
	sha256Func := sha256.New()
	sha256Func.Write(data)
	return sha256Func.Sum(nil), nil
}

func generateLeaves(blocks []DataBlock, hashFunc func([]byte) ([]byte, error)) ([]*Node, error) {
	var (
		lenLeaves = len(blocks)
		leaves    = make([]*Node, lenLeaves)
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
		leaves[i] = &Node{Hash: hash}
	}
	return leaves, nil
}

func generateLeavesParallel() ([]*Node, error) {
	panic("not implemented")
}

// Reset resets the Merkle Tree
func (m *MerkleTree) Reset() {
	m.Leaves = nil
	m.Root = nil
	m.Proves = nil
}

// Verify verifies the data block with the Merkle Tree proof
func (m *MerkleTree) Verify(dataBlock DataBlock, proof *Proof) (bool, error) {
	return Verify(dataBlock, proof, m.Root, m.HashFunc)
}

// Verify verifies the data block with the Merkle Tree proof and Merkle root hash
func Verify(dataBlock DataBlock, proof *Proof, root []byte,
	hashFunc func([]byte) ([]byte, error)) (bool, error) {
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
	hash, err = defaultHashFunc(data)
	if err != nil {
		return false, err
	}
	for _, n := range proof.Neighbors {
		dir := proof.Path & 1
		if dir == 1 {
			hash, err = defaultHashFunc(append(hash, n...))
		} else {
			hash, err = defaultHashFunc(append(n, hash...))
		}
		if err != nil {
			return false, err
		}
		proof.Path >>= 1
	}
	return bytes.Equal(hash, root), nil
}
