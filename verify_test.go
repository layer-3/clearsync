// MIT License
//
// # Copyright (c) 2023 Tommy TIAN
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
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/txaty/go-merkletree/mock"
)

func setupTestVerify(size int) (*MerkleTree, []DataBlock) {
	blocks := mockDataBlocks(size)
	m, err := New(nil, blocks)
	if err != nil {
		panic(err)
	}
	return m, blocks
}

func setupTestVerifyParallel(size int) (*MerkleTree, []DataBlock) {
	blocks := mockDataBlocks(size)
	m, err := New(&Config{
		RunInParallel: true,
		NumRoutines:   1,
	}, blocks)
	if err != nil {
		panic(err)
	}
	return m, blocks
}

func TestMerkleTreeVerify(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(int) (*MerkleTree, []DataBlock)
		blockSize int
		want      bool
		wantErr   bool
	}{
		{
			name:      "test_2",
			setupFunc: setupTestVerify,
			blockSize: 2,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_3",
			setupFunc: setupTestVerify,
			blockSize: 3,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_4",
			setupFunc: setupTestVerify,
			blockSize: 4,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_5",
			setupFunc: setupTestVerify,
			blockSize: 5,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_6",
			setupFunc: setupTestVerify,
			blockSize: 6,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_8",
			setupFunc: setupTestVerify,
			blockSize: 8,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_9",
			setupFunc: setupTestVerify,
			blockSize: 9,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_1001",
			setupFunc: setupTestVerify,
			blockSize: 1001,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_2_parallel",
			setupFunc: setupTestVerifyParallel,
			blockSize: 2,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_4_parallel",
			setupFunc: setupTestVerifyParallel,
			blockSize: 4,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_64_parallel",
			setupFunc: setupTestVerifyParallel,
			blockSize: 64,
			want:      true,
			wantErr:   false,
		},
		{
			name:      "test_1001_parallel",
			setupFunc: setupTestVerifyParallel,
			blockSize: 1001,
			want:      true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, blocks := tt.setupFunc(tt.blockSize)
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

func TestVerify(t *testing.T) {
	blocks := mockDataBlocks(5)
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
		config    *Config
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
				config: &Config{
					HashFunc: m.HashFunc,
				},
			},
			want: true,
		},
		{
			name: "test_config_nil",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
			},
			want: true,
		},
		{
			name: "test_wrong_root",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      []byte("test_wrong_root"),
				config: &Config{
					HashFunc: m.HashFunc,
				},
			},
			want: false,
		},
		{
			name: "test_wrong_hash_func",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) { return []byte("test_wrong_hash_hash"), nil },
				},
			},
			want: false,
		},
		{
			name: "test_proof_nil",
			args: args{
				dataBlock: blocks[0],
				proof:     nil,
				root:      m.Root,
				config: &Config{
					HashFunc: m.HashFunc,
				},
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
				config: &Config{
					HashFunc: m.HashFunc,
				},
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
				config: &Config{
					HashFunc: nil,
				},
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
				config: &Config{
					HashFunc: func([]byte) ([]byte, error) {
						return nil, errors.New("test_hash_func_err")
					},
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "test_disable_leaf_hashing_and_hash_func_err",
			args: args{
				dataBlock: blocks[0],
				proof:     m.Proofs[0],
				root:      m.Root,
				config: &Config{
					DisableLeafHashing: true,
					HashFunc: func([]byte) ([]byte, error) {
						return nil, errors.New("test_hash_func_err")
					},
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
				config: &Config{
					HashFunc: m.HashFunc,
				},
			},
			mock: func() {
				patches.ApplyMethod(reflect.TypeOf(&mock.DataBlock{}), "Serialize",
					func(m *mock.DataBlock) ([]byte, error) {
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
			got, err := Verify(tt.args.dataBlock, tt.args.proof, tt.args.root, tt.args.config)
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
