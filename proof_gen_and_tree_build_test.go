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
	"crypto/sha256"
	"fmt"
	"reflect"
	"testing"
)

func TestMerkleTreeNew_modeProofGenAndTreeBuild(t *testing.T) {
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
				blocks: mockDataBlocks(2),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_4",
			args: args{
				blocks: mockDataBlocks(4),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_5",
			args: args{
				blocks: mockDataBlocks(5),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_8",
			args: args{
				blocks: mockDataBlocks(8),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_proof_9",
			args: args{
				blocks: mockDataBlocks(9),
				config: &Config{
					Mode: ModeProofGenAndTreeBuild,
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
					Mode: ModeProofGenAndTreeBuild,
				},
			},
			wantErr: true,
		},
		{
			name: "test_tree_build_hash_func_error",
			args: args{
				blocks: mockDataBlocks(100),
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

func TestMerkleTreeNew_modeProofGenAndTreeBuildParallel(t *testing.T) {
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
				blocks: mockDataBlocks(2),
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
				blocks: mockDataBlocks(4),
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
				blocks: mockDataBlocks(5),
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
				blocks: mockDataBlocks(8),
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
				blocks: mockDataBlocks(100),
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
				blocks: mockDataBlocks(100),
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
