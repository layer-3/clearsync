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
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

const benchSize = 10000

type mockDataBlock struct {
	data []byte
}

func (t *mockDataBlock) Serialize() ([]byte, error) {
	return t.data, nil
}

func genTestDataBlocks(num int) []DataBlock {
	var blocks []DataBlock
	for i := 0; i < num; i++ {
		block := &mockDataBlock{
			data: make([]byte, 100),
		}
		_, err := rand.Read(block.data)
		if err != nil {
			panic(err)
		}
		blocks = append(blocks, block)
	}
	return blocks
}

func TestMerkleTreeNew_proofGen(t *testing.T) {
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_0",
			args: args{
				blocks: genTestDataBlocks(0),
			},
			wantErr: true,
		},
		{
			name: "test_1",
			args: args{
				blocks: genTestDataBlocks(1),
			},
			wantErr: true,
		},
		{
			name: "test_2",
			args: args{
				blocks: genTestDataBlocks(2),
			},
			wantErr: false,
		},
		{
			name: "test_8",
			args: args{
				blocks: genTestDataBlocks(8),
			},
			wantErr: false,
		},
		{
			name: "test_5",
			args: args{
				blocks: genTestDataBlocks(5),
			},
			wantErr: false,
		},
		{
			name: "test_1000",
			args: args{
				blocks: genTestDataBlocks(1000),
			},
			wantErr: false,
		},
		{
			name: "test_100_parallel",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
				},
			},
			wantErr: false,
		},
		{
			name: "test_10_32_parallel",
			args: args{
				blocks: genTestDataBlocks(10),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   32,
				},
			},
			wantErr: false,
		},
		{
			name: "test_100_parallel_no_specify_num_of_routines",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					RunInParallel: true,
				},
			},
			wantErr: false,
		},
		{
			name: "test_100_parallel_random",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					NoDuplicates:  true,
					RunInParallel: true,
					NumRoutines:   4,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_hash_func_error_parallel",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
					RunInParallel: true,
				},
			},
			wantErr: true,
		},
		{
			name: "bad_mode",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					Mode: 5,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := New(tt.args.config, tt.args.blocks); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMerkleTreeNew_buildTree(t *testing.T) {
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_build_tree_2",
			args: args{
				blocks: genTestDataBlocks(2),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_5",
			args: args{
				blocks: genTestDataBlocks(5),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_8",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_1000",
			args: args{
				blocks: genTestDataBlocks(1000),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
					Mode: ModeTreeBuild,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			m1, err := New(nil, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			if !bytes.Equal(m.Root, m1.Root) && !tt.wantErr {
				fmt.Println("m", m.Root)
				fmt.Println("m1", m1.Root)
				t.Errorf("tree generated is wrong")
				return
			}
		})
	}
}

func TestMerkleTreeNew_treeBuildParallel(t *testing.T) {
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_build_tree_parallel_2",
			args: args{
				blocks: genTestDataBlocks(2),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_parallel_4",
			args: args{
				blocks: genTestDataBlocks(4),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_parallel_5",
			args: args{
				blocks: genTestDataBlocks(5),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_parallel_8",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_parallel_8_32",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   32,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error_parallel",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
					RunInParallel: true,
					Mode:          ModeTreeBuild,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			m1, err := New(nil, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			if !bytes.Equal(m.Root, m1.Root) && !tt.wantErr {
				fmt.Println("m", m.Root)
				fmt.Println("m1", m1.Root)
				t.Errorf("tree generated is wrong")
				return
			}
		})
	}
}

func TestMerkleTreeNew_proofGenAndTreeBuild(t *testing.T) {
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_build_tree_proof_2",
			args: args{
				blocks: genTestDataBlocks(2),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_4",
			args: args{
				blocks: genTestDataBlocks(4),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_5",
			args: args{
				blocks: genTestDataBlocks(5),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_8",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_9",
			args: args{
				blocks: genTestDataBlocks(9),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: true,
		},
		{
			name: "test_tree_build_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func(block []byte) ([]byte, error) {
						if len(block) == 64 {
							return nil, fmt.Errorf("hash func error")
						}
						sha256Func := sha256.New()
						sha256Func.Write(block)
						return sha256Func.Sum(nil), nil
					},
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			m1, err := New(nil, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			for i := 0; i < len(tt.args.blocks); i++ {
				if !reflect.DeepEqual(m.Proofs[i], m1.Proofs[i]) {
					t.Errorf("proofs generated are wrong for block %d", i)
					return
				}
			}
		})
	}
}

func TestMerkleTreeNew_proofGenAndTreeBuildParallel(t *testing.T) {
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_build_tree_proof_parallel_2",
			args: args{
				blocks: genTestDataBlocks(2),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_parallel_4",
			args: args{
				blocks: genTestDataBlocks(4),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_parallel_5",
			args: args{
				blocks: genTestDataBlocks(5),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_parallel_8",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
					Mode:          ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, fmt.Errorf("hash func error")
					},
					Mode:          ModeProofGenAndTreeBuild,
					RunInParallel: true,
				},
			},
			wantErr: true,
		},
		{
			name: "test_tree_build_hash_func_error",
			args: args{
				blocks: genTestDataBlocks(100),
				config: &Config{
					HashFunc: func(block []byte) ([]byte, error) {
						if len(block) == 64 {
							return nil, fmt.Errorf("hash func error")
						}
						sha256Func := sha256.New()
						sha256Func.Write(block)
						return sha256Func.Sum(nil), nil
					},
					Mode:          ModeProofGenAndTreeBuild,
					RunInParallel: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			m1, err := New(nil, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			for i := 0; i < len(tt.args.blocks); i++ {
				if !reflect.DeepEqual(m.Proofs[i], m1.Proofs[i]) {
					t.Errorf("proofs generated are wrong for block %d", i)
					return
				}
			}
		})
	}
}

func verifySetup(size int) (*MerkleTree, []DataBlock, error) {
	blocks := genTestDataBlocks(size)
	m, err := New(nil, blocks)
	if err != nil {
		return nil, nil, err
	}
	return m, blocks, nil
}

func verifySetupParallel(size int) (*MerkleTree, []DataBlock, error) {
	blocks := genTestDataBlocks(size)
	m, err := New(&Config{
		RunInParallel: true,
		NumRoutines:   4,
	}, blocks)
	if err != nil {
		return nil, nil, err
	}
	return m, blocks, nil
}

func TestMerkleTree_Verify(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(int) (*MerkleTree, []DataBlock, error)
		blockSize int
		want      bool
		wantErr   bool
	}{
		{
			name:      "test_2",
			setupFunc: verifySetup,
			blockSize: 2,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_3",
			setupFunc: verifySetup,
			blockSize: 3,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_4",
			setupFunc: verifySetup,
			blockSize: 4,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_5",
			setupFunc: verifySetup,
			blockSize: 5,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_6",
			setupFunc: verifySetup,
			blockSize: 6,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_8",
			setupFunc: verifySetup,
			blockSize: 8,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_9",
			setupFunc: verifySetup,
			blockSize: 9,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_1001",
			setupFunc: verifySetup,
			blockSize: 1001,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_64_parallel",
			setupFunc: verifySetupParallel,
			blockSize: 64,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_pseudo_random_1001_parallel",
			setupFunc: verifySetupParallel,
			blockSize: 1001,
			want:      true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, blocks, err := tt.setupFunc(tt.blockSize)
			if err != nil {
				t.Errorf("setupFunc() error = %v", err)
				return
			}
			for i := 0; i < tt.blockSize; i++ {
				got, err := m.Verify(blocks[i], m.Proofs[i])
				if (err != nil) != tt.wantErr {
					t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Verify() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMerkleTree_GenerateProof(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	tests := []struct {
		name        string
		config      *Config
		mock        func()
		blocks      []DataBlock
		proofBlocks []DataBlock
		wantErr     bool
	}{
		{
			name:   "test_2",
			config: &Config{Mode: ModeTreeBuild},
			blocks: genTestDataBlocks(2),
		},
		{
			name:   "test_4",
			config: &Config{Mode: ModeTreeBuild},
			blocks: genTestDataBlocks(4),
		},
		{
			name:   "test_5",
			config: &Config{Mode: ModeTreeBuild},
			blocks: genTestDataBlocks(5),
		},
		{
			name:    "test_wrong_mode",
			config:  &Config{Mode: ModeProofGen},
			blocks:  genTestDataBlocks(5),
			wantErr: true,
		},
		{
			name:   "test_wrong_blocks",
			config: &Config{Mode: ModeTreeBuild},
			blocks: genTestDataBlocks(5),
			proofBlocks: []DataBlock{
				&mockDataBlock{
					[]byte("test_wrong_blocks"),
				},
			},
			wantErr: true,
		},
		{
			name:   "test_data_block_serialize_error",
			config: &Config{Mode: ModeTreeBuild},
			mock: func() {
				patches.ApplyMethod(reflect.TypeOf(&mockDataBlock{}), "Serialize",
					func(*mockDataBlock) ([]byte, error) {
						return nil, errors.New("data block serialize error")
					})
			},
			blocks:  genTestDataBlocks(5),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1, err := New(nil, tt.blocks)
			if err != nil {
				t.Errorf("m1 New() error = %v", err)
				return
			}
			m2, err := New(tt.config, tt.blocks)
			if err != nil {
				t.Errorf("m2 New() error = %v", err)
				return
			}
			if tt.proofBlocks == nil {
				tt.proofBlocks = tt.blocks
			}
			if tt.mock != nil {
				tt.mock()
			}
			defer patches.Reset()
			for idx, block := range tt.proofBlocks {
				got, err := m2.GenerateProof(block)
				if (err != nil) != tt.wantErr {
					t.Errorf("GenerateProof() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantErr {
					return
				}
				if !reflect.DeepEqual(got, m1.Proofs[idx]) && !tt.wantErr {
					t.Errorf("GenerateProof() %d got = %v, want %v", idx, got, m1.Proofs[idx])
					return
				}
			}
		})
	}
}

func testHashFunc(data []byte) ([]byte, error) {
	sha256Func := sha256.New()
	sha256Func.Write(data)
	return sha256Func.Sum(nil), nil
}

func TestMerkleTree_proofGen(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	type args struct {
		config *Config
		blocks []DataBlock
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "test_fix_odd_err",
			args: args{
				config: &Config{
					NoDuplicates: true,
				},
				blocks: genTestDataBlocks(5),
			},
			mock: func() {
				patches.ApplyFunc(getDummyHash,
					func() ([]byte, error) {
						return nil, errors.New("test_get_dummy_hash_err")
					})
			},
			wantErr: true,
		},
		{
			name: "test_hash_func_err",
			args: args{
				config: &Config{
					HashFunc: testHashFunc,
				},
				blocks: genTestDataBlocks(5),
			},
			mock: func() {
				patches.ApplyFunc(testHashFunc,
					func([]byte) ([]byte, error) {
						return nil, errors.New("test_hash_func_err")
					})
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if err != nil {
				t.Errorf("New() error = %v", err)
				return
			}
			if tt.mock != nil {
				tt.mock()
			}
			defer patches.Reset()
			if err := m.proofGen(); (err != nil) != tt.wantErr {
				t.Errorf("proofGen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	blocks := genTestDataBlocks(5)
	m, err := New(nil, blocks)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	type args struct {
		dataBlock DataBlock
		proof     *Proof
		root      []byte
		hashFunc  HashFuncType
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    bool
		wantErr bool
	}{
		{
			name: "test_ok",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc:  m.HashFunc,
			},
			want: true,
		},
		{
			name: "test_wrong_root",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      []byte("test_wrong_root"),
				hashFunc:  m.HashFunc,
			},
			want: false,
		},
		{
			name: "test_wrong_hash_func",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc:  func([]byte) ([]byte, error) { return []byte("test_wrong_hash_hash"), nil },
			},
			want: false,
		},
		{
			name: "test_proof_nil",
			args: args{
				dataBlock: blocks[0],
				proof:     nil,
				root:      m.Root,
				hashFunc:  m.HashFunc,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "test_data_block_nil",
			args: args{
				dataBlock: nil,
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc:  m.HashFunc,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "test_hash_func_nil",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc:  nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test_hash_func_err",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc: func([]byte) ([]byte, error) {
					return nil, errors.New("test_hash_func_err")
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "data_block_serialize_err",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				hashFunc:  m.HashFunc,
			},
			mock: func() {
				patches.ApplyMethod(reflect.TypeOf(&mockDataBlock{}), "Serialize",
					func(m *mockDataBlock) ([]byte, error) {
						return nil, errors.New("test_data_block_serialize_err")
					})
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			defer patches.Reset()
			got, err := Verify(tt.args.dataBlock, tt.args.proof, tt.args.root, tt.args.hashFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_proofGenHandler(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	type args struct {
		argInterface interface{}
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "test_hash_func_err",
			args: args{
				argInterface: &proofGenArgs{
					hashFunc: func([]byte) ([]byte, error) {
						return nil, errors.New("test_hash_func_err")
					},
					buf1:        [][]byte{[]byte("test_buf1"), []byte("test_buf1")},
					buf2:        [][]byte{[]byte("test_buf2")},
					prevLen:     2,
					numRoutines: 2,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}
			defer patches.Reset()
			if err := proofGenHandler(tt.args.argInterface); (err != nil) != tt.wantErr {
				t.Errorf("proofGenHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkMerkleTreeNew(b *testing.B) {
	testCases := genTestDataBlocks(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(nil, testCases)
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNewParallel(b *testing.B) {
	config := &Config{
		RunInParallel: true,
	}
	testCases := genTestDataBlocks(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeBuild(b *testing.B) {
	testCases := genTestDataBlocks(benchSize)
	config := &Config{
		Mode: ModeTreeBuild,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeBuildParallel(b *testing.B) {
	config := &Config{
		Mode:          ModeTreeBuild,
		RunInParallel: true,
	}
	testCases := genTestDataBlocks(benchSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := New(config, testCases)
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}
