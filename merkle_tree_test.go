package merkletree

import (
	"math/rand"
	"testing"
)

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

func TestMerkleTree_Build(t *testing.T) {
	type fields struct {
		Config *Config
		Root   []byte
		Leaves []*Node
		Proves []*Proof
	}
	type args struct {
		blocks []DataBlock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_0",
			fields: fields{
				Config: &Config{
					HashFunc: defaultHashFunc,
				},
				Root:   nil,
				Leaves: nil,
				Proves: nil,
			},
			args: args{
				blocks: genTestDataBlocks(0),
			},
			wantErr: false,
		},
		{
			name: "test_4",
			fields: fields{
				Config: &Config{
					HashFunc: defaultHashFunc,
				},
				Root:   nil,
				Leaves: nil,
				Proves: nil,
			},
			args: args{
				blocks: genTestDataBlocks(4),
			},
			wantErr: false,
		},
		{
			name: "test_8",
			fields: fields{
				Config: &Config{
					HashFunc:        defaultHashFunc,
					AllowDuplicates: true,
				},
				Root:   nil,
				Leaves: nil,
				Proves: nil,
			},
			args: args{
				blocks: genTestDataBlocks(8),
			},
			wantErr: false,
		},
		{
			name: "test_5",
			fields: fields{
				Config: &Config{
					HashFunc:        defaultHashFunc,
					AllowDuplicates: true,
				},
				Root:   nil,
				Leaves: nil,
				Proves: nil,
			},
			args: args{
				blocks: genTestDataBlocks(5),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MerkleTree{
				Config: tt.fields.Config,
				Root:   tt.fields.Root,
				Leaves: tt.fields.Leaves,
				Proves: tt.fields.Proves,
			}
			if err := m.Build(tt.args.blocks); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func verifySetup(size int) (*MerkleTree, []DataBlock, error) {
	blocks := genTestDataBlocks(size)
	m := NewMerkleTree(&Config{
		HashFunc:        defaultHashFunc,
		AllowDuplicates: true,
	})
	err := m.Build(blocks)
	if err != nil {
		return nil, nil, err
	}
	return m, blocks, nil
}

func TestMerkleTree_Verify(t *testing.T) {
	type fields struct {
		Config *Config
		Root   []byte
		Leaves []*Node
		Proves []*Proof
	}
	type args struct {
		dataBlock DataBlock
		proof     *Proof
	}
	tests := []struct {
		name       string
		setupFunc  func(int) (*MerkleTree, []DataBlock, error)
		blockSize  int
		blockIndex int
		want       bool
		wantErr    bool
	}{
		{
			name:       "test_2",
			setupFunc:  verifySetup,
			blockSize:  2,
			blockIndex: 0,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "test_3",
			setupFunc:  verifySetup,
			blockSize:  3,
			blockIndex: 2,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "test_pseudo_random_4",
			setupFunc:  verifySetup,
			blockSize:  4,
			blockIndex: 3,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "test_pseudo_random_5",
			setupFunc:  verifySetup,
			blockSize:  5,
			blockIndex: 4,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "test_pseudo_random_6",
			setupFunc:  verifySetup,
			blockSize:  6,
			blockIndex: 5,
			want:       true,
			wantErr:    false,
		},
		{
			name:       "test_pseudo_random_8",
			setupFunc:  verifySetup,
			blockSize:  8,
			blockIndex: 7,
			want:       true,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, blocks, err := tt.setupFunc(tt.blockSize)
			if err != nil {
				t.Errorf("setupFunc() error = %v", err)
				return
			}
			got, err := m.Verify(blocks[tt.blockIndex], m.Proves[tt.blockIndex])
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}
