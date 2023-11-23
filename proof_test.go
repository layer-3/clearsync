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
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/txaty/go-merkletree/mock"
)

func TestMerkleTree_Proof(t *testing.T) {
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
			blocks: mockDataBlocks(2),
		},
		{
			name:   "test_4",
			config: &Config{Mode: ModeTreeBuild},
			blocks: mockDataBlocks(4),
		},
		{
			name:   "test_5",
			config: &Config{Mode: ModeTreeBuild},
			blocks: mockDataBlocks(5),
		},
		{
			name:    "test_wrong_mode",
			config:  &Config{Mode: ModeProofGen},
			blocks:  mockDataBlocks(5),
			wantErr: true,
		},
		{
			name:   "test_wrong_blocks",
			config: &Config{Mode: ModeTreeBuild},
			blocks: mockDataBlocks(5),
			proofBlocks: []DataBlock{
				&mock.DataBlock{
					Data: []byte("test_wrong_blocks"),
				},
			},
			wantErr: true,
		},
		{
			name:   "test_data_block_serialize_error",
			config: &Config{Mode: ModeTreeBuild},
			mock: func() {
				patches.ApplyMethod(reflect.TypeOf(&mock.DataBlock{}), "Serialize",
					func(*mock.DataBlock) ([]byte, error) {
						return nil, errors.New("data block serialize error")
					})
			},
			blocks:  mockDataBlocks(5),
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
				got, err := m2.Proof(block)
				if (err != nil) != tt.wantErr {
					t.Errorf("Proof() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if tt.wantErr {
					return
				}
				if !reflect.DeepEqual(got, m1.Proofs[idx]) && !tt.wantErr {
					t.Errorf("Proof() %d got = %v, want %v", idx, got, m1.Proofs[idx])
					return
				}
			}
		})
	}
}
