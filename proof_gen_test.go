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
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/txaty/go-merkletree/mock"
)

func TestMerkleTreeNew_modeProofGen(t *testing.T) {
	dummyDataBlocks := []DataBlock{
		&mock.DataBlock{
			Data: []byte("dummy_data_0"),
		},
		&mock.DataBlock{
			Data: []byte("dummy_data_1"),
		},
		&mock.DataBlock{
			Data: []byte("dummy_data_2"),
		},
		&mock.DataBlock{
			Data: []byte("dummy_data_3"),
		},
		&mock.DataBlock{
			Data: []byte("dummy_data_4"),
		},
	}
	dummyRootSizeTwo, err := hex.DecodeString("30c87249cdfa43ed48a8e6a7747e05312933eaf2cbca38974e101649d3bae339")
	if err != nil {
		t.Fatal(err)
	}
	dummyRootSizeThree, err := hex.DecodeString("1437da40a42b98b5fd7a76493375a687f0b2418071a009d1459f24e61db48c70")
	if err != nil {
		t.Fatal(err)
	}
	dummyRootSizeFour, err := hex.DecodeString("bfef6f8177eaf488386177d31a24e0481781811388ac1c8e291d564dc3a79fb9")
	if err != nil {
		t.Fatal(err)
	}
	dummyRootSizeFive, err := hex.DecodeString("5e9ba12514859d30d6e79a6f2d6e9aa71a6b011ae35e6c07f487c6d570993d71")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantRoot []byte
	}{
		{
			name: "test_0",
			args: args{
				blocks: mockDataBlocks(0),
			},
			wantErr: true,
		},
		{
			name: "test_1",
			args: args{
				blocks: mockDataBlocks(1),
			},
			wantErr: true,
		},
		{
			name: "test_2",
			args: args{
				blocks: []DataBlock{dummyDataBlocks[0], dummyDataBlocks[1]},
			},
			wantErr:  false,
			wantRoot: dummyRootSizeTwo,
		},
		{
			name: "test_3",
			args: args{
				blocks: []DataBlock{dummyDataBlocks[0], dummyDataBlocks[1], dummyDataBlocks[2]},
			},
			wantErr:  false,
			wantRoot: dummyRootSizeThree,
		},
		{
			name: "test_4",
			args: args{
				blocks: []DataBlock{dummyDataBlocks[0], dummyDataBlocks[1], dummyDataBlocks[2], dummyDataBlocks[3]},
			},
			wantErr:  false,
			wantRoot: dummyRootSizeFour,
		},
		{
			name: "test_5",
			args: args{
				blocks: dummyDataBlocks,
			},
			wantErr:  false,
			wantRoot: dummyRootSizeFive,
		},
		{
			name: "test_7",
			args: args{
				blocks: mockDataBlocks(7),
			},
			wantErr: false,
		},
		{
			name: "test_8",
			args: args{
				blocks: mockDataBlocks(8),
			},
			wantErr: false,
		},
		{
			name: "test_5",
			args: args{
				blocks: mockDataBlocks(5),
			},
			wantErr: false,
		},
		{
			name: "test_1000",
			args: args{
				blocks: mockDataBlocks(1000),
			},
			wantErr: false,
		},
		{
			name: "test_100_parallel_4",
			args: args{
				blocks: mockDataBlocks(100),
				config: &Config{
					RunInParallel: true,
					NumRoutines:   4,
				},
			},
			wantErr: false,
		},
		{
			name: "test_10_parallel_32",
			args: args{
				blocks: mockDataBlocks(100),
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
				blocks: mockDataBlocks(100),
				config: &Config{
					RunInParallel: true,
				},
			},
			wantErr: false,
		},
		{
			name: "test_8_sorted",
			args: args{
				blocks: mockDataBlocks(8),
				config: &Config{
					SortSiblingPairs: true,
				},
			},
			wantErr: false,
		},
		{
			name: "test_hash_func_error",
			args: args{
				blocks: mockDataBlocks(100),
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
				blocks: mockDataBlocks(100),
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
			name: "test_100_disable_leaf_hashing",
			args: args{
				blocks: mockDataBlocks(100),
				config: &Config{
					DisableLeafHashing: true,
				},
			},
			wantErr: false,
		},
		{
			name: "test_100_disable_leaf_hashing_parallel_4",
			args: args{
				blocks: mockDataBlocks(100),
				config: &Config{
					DisableLeafHashing: true,
					RunInParallel:      true,
					NumRoutines:        4,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_mode",
			args: args{
				blocks: mockDataBlocks(100),
				config: &Config{
					Mode: 5,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if mt == nil {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantRoot == nil {
				for idx, block := range tt.args.blocks {
					ok, err := mt.Verify(block, mt.Proofs[idx])
					if err != nil {
						t.Errorf("proof verification error, idx %d, err %v", idx, err)
						return
					}
					if !ok {
						t.Errorf("proof verification failed, idx %d", idx)
						return
					}
				}
			} else {
				if !bytes.Equal(mt.Root, tt.wantRoot) {
					t.Errorf("root mismatch, got %x, want %x", mt.Root, tt.wantRoot)
					return
				}
			}
		})
	}
}

func mockHashFunc(data []byte) ([]byte, error) {
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
			name: "test_hash_func_err",
			args: args{
				config: &Config{
					HashFunc: mockHashFunc,
				},
				blocks: mockDataBlocks(5),
			},
			mock: func() {
				patches.ApplyFunc(mockHashFunc,
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

func TestMerkleTree_proofGenParallel(t *testing.T) {
	var hashFuncCounter atomic.Uint32
	type args struct {
		config *Config
		blocks []DataBlock
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_goroutine_err",
			args: args{
				config: &Config{
					HashFunc: func(data []byte) ([]byte, error) {
						if hashFuncCounter.Load() == 9 {
							return nil, errors.New("test_goroutine_err")
						}
						hashFuncCounter.Add(1)
						return mockHashFunc(data)
					},
					RunInParallel: true,
				},
				blocks: mockDataBlocks(4),
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
			if err := m.proofGenParallel(); (err != nil) != tt.wantErr {
				t.Errorf("proofGenParallel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
