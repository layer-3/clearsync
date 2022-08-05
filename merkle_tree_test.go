package merkletree

import (
	"math/rand"
	"testing"
)

type mockDataBlock struct{}

func (t *mockDataBlock) Serialize() ([]byte, error) {
	dummyBytes := make([]byte, 100)
	_, err := rand.Read(dummyBytes)
	if err != nil {
		return nil, err
	}
	return dummyBytes, nil
}

func genTestDataBlocks(num int) []DataBlock {
	var blocks []DataBlock
	for i := 0; i < num; i++ {
		block := new(mockDataBlock)
		blocks = append(blocks, block)
	}
	return blocks
}

func TestMerkleTree_Build(t *testing.T) {
	type fields struct {
		Config *Config
		Root   *Node
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
