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
	"bytes"
	"crypto/sha256"
	"fmt"
	"sync/atomic"
	"testing"
)

func TestMerkleTreeNew_modeTreeBuild(t *testing.T) {
	var hashFuncCounter int
	type args struct {
		blocks []DataBlock
		config *Config
	}
	tests := []struct {
		name           string
		args           args
		checkingConfig *Config
		wantErr        bool
	}{
		{
			name: "test_build_tree_2",
			args: args{
				blocks: mockDataBlocks(2),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_3",
			args: args{
				blocks: mockDataBlocks(3),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_5",
			args: args{
				blocks: mockDataBlocks(5),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_8",
			args: args{
				blocks: mockDataBlocks(8),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_16",
			args: args{
				blocks: mockDataBlocks(16),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_32",
			args: args{
				blocks: mockDataBlocks(32),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_36",
			args: args{
				blocks: mockDataBlocks(36),
				config: &Config{
					Mode: ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_1000",
			args: args{
				blocks: mockDataBlocks(1000),
				config: &Config{
					Mode: ModeTreeBuild,
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
					Mode: ModeTreeBuild,
				},
			},
			wantErr: true,
		},
		{
			name: "test_hash_func_error_when_computing_root",
			args: args{
				blocks: mockDataBlocks(4),
				config: &Config{
					HashFunc: func(block []byte) ([]byte, error) {
						if hashFuncCounter == 6 {
							return nil, fmt.Errorf("hash func error")
						}
						hashFuncCounter++
						sha256Func := sha256.New()
						sha256Func.Write(block)
						return sha256Func.Sum(nil), nil
					},
					Mode: ModeTreeBuild,
				},
			},
			wantErr: true,
		},
		{
			name: "test_disable_leaf_hashing",
			args: args{
				blocks: mockDataBlocks(100),
				config: &Config{
					DisableLeafHashing: true,
					Mode:               ModeTreeBuild,
				},
			},
			checkingConfig: &Config{
				DisableLeafHashing: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.args.config, tt.args.blocks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			m1, err := New(tt.checkingConfig, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			if !tt.wantErr && !bytes.Equal(m.Root, m1.Root) && !tt.wantErr {
				fmt.Println("m", m.Root)
				fmt.Println("m1", m1.Root)
				t.Errorf("tree generated is wrong")
				return
			}
		})
	}
}

func TestMerkleTreeNew_modeTreeBuildParallel(t *testing.T) {
	var hashFuncCounter atomic.Uint32
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
				blocks: mockDataBlocks(2),
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
				blocks: mockDataBlocks(4),
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
				blocks: mockDataBlocks(5),
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
				blocks: mockDataBlocks(8),
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
				blocks: mockDataBlocks(8),
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
				blocks: mockDataBlocks(100),
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
		{
			name: "test_hash_func_error_when_computing_nodes_parallel",
			args: args{
				blocks: mockDataBlocks(4),
				config: &Config{
					HashFunc: func(block []byte) ([]byte, error) {
						if hashFuncCounter.Load() == 5 {
							return nil, fmt.Errorf("hash func error")
						}
						hashFuncCounter.Add(1)
						sha256Func := sha256.New()
						sha256Func.Write(block)
						return sha256Func.Sum(nil), nil
					},
				},
			},
			wantErr: true,
		},
		{
			name: "test_hash_func_error_when_computing_root_parallel",
			args: args{
				blocks: mockDataBlocks(4),
				config: &Config{
					HashFunc: func(block []byte) ([]byte, error) {
						if hashFuncCounter.Load() == 6 {
							return nil, fmt.Errorf("hash func error")
						}
						hashFuncCounter.Add(1)
						sha256Func := sha256.New()
						sha256Func.Write(block)
						return sha256Func.Sum(nil), nil
					},
					Mode:          ModeTreeBuild,
					RunInParallel: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashFuncCounter.Store(0)
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
			if !tt.wantErr && !bytes.Equal(m.Root, m1.Root) && !tt.wantErr {
				fmt.Println("m", m.Root)
				fmt.Println("m1", m1.Root)
				t.Errorf("tree generated is wrong")
				return
			}
		})
	}
}
