package merkletree

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	defaultHashLen = 32
)

type DataBlock interface {
	Serialize() ([]byte, error)
}

type Config struct {
	HashFunc        func([]byte) ([]byte, error)
	RunInParallel   bool
	NumRoutines     int
	AllowDuplicates bool
}

type MerkleTree struct {
	*Config
	Root   *Node
	Leaves []*Node
	Proves []*Proof
}

type Node struct {
	Hash []byte
}

type Proof struct {
	PathWay   uint16
	Neighbors [][]byte
}

func (m *MerkleTree) NewMerkleTree(config Config) *MerkleTree {
	if config.HashFunc == nil {
		config.HashFunc = defaultHashFunc
	}
	return &MerkleTree{
		Config: &config,
	}
}

func (m *MerkleTree) Build(blocks []DataBlock) (err error) {
	if len(blocks) == 0 {
		return nil
	}
	if m.RunInParallel {
		m.Leaves, err = generateLeavesParallel(blocks, m.HashFunc)
		if err != nil {
			return err
		}
		m.Root, err = m.buildTreeParallel()
	} else {
		m.Leaves, err = generateLeaves(blocks, m.HashFunc)
		if err != nil {
			return err
		}
		m.Root, err = m.buildTree()
	}
	return err
}

func (m *MerkleTree) buildTree() (root *Node, err error) {
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
	for prevLen > 1 {
		buf, prevLen, err = m.fixOdd(buf, prevLen)
		if err != nil {
			return nil, err
		}
		for idx := 0; idx < prevLen; idx++ {
			fmt.Println(hex.EncodeToString(buf[idx].Hash[:]))
		}
		fmt.Println()
		for idx := 0; idx < prevLen; idx += 2 {
			var appendHash []byte
			appendHash = append(buf[idx].Hash, buf[idx+1].Hash...)
			buf[idx/2].Hash, err = m.HashFunc(appendHash)
		}
		prevLen /= 2
		m.assignProves(buf, prevLen, step)
		step++
	}
	root = buf[0]
	m.Root = root
	for _, p := range m.Proves {
		fmt.Print(p.PathWay)
		fmt.Print(" ")
		for _, n := range p.Neighbors {
			fmt.Print(hex.EncodeToString(n))
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println(m.Root)
	return
}

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
			m.Proves[j].PathWay += 1 << step
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

func (m *MerkleTree) buildTreeParallel() (*Node, error) {
	panic("not implemented")
}

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

func generateLeavesParallel(blocks []DataBlock, hashFunc func([]byte) ([]byte, error)) ([]*Node, error) {
	panic("not implemented")
}
