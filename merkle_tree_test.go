package merkletree

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
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
			name: "test_4",
			args: args{
				blocks: genTestDataBlocks(4),
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
					HashFunc:      defaultHashFunc,
					RunInParallel: true,
					NumRoutines:   4,
				},
			},
			wantErr: false,
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
					HashFunc: defaultHashFunc,
					Mode:     ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_4",
			args: args{
				blocks: genTestDataBlocks(4),
				config: &Config{
					HashFunc: defaultHashFunc,
					Mode:     ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_5",
			args: args{
				blocks: genTestDataBlocks(5),
				config: &Config{
					HashFunc: defaultHashFunc,
					Mode:     ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_8",
			args: args{
				blocks: genTestDataBlocks(8),
				config: &Config{
					HashFunc: defaultHashFunc,
					Mode:     ModeTreeBuild,
				},
			},
			wantErr: false,
		},
		{
			name: "test_build_tree_1000",
			args: args{
				blocks: genTestDataBlocks(1000),
				config: &Config{
					HashFunc: defaultHashFunc,
					Mode:     ModeTreeBuild,
				},
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
			m1, err := New(nil, tt.args.blocks)
			if err != nil {
				t.Errorf("test setup error %v", err)
				return
			}
			if !bytes.Equal(m.Root, m1.Root) {
				fmt.Println("m", m.Root)
				fmt.Println("m1", m1.Root)
				t.Errorf("tree generated is wrong")
				return
			}
		})
	}
}

func verifySetup(size int) (*MerkleTree, []DataBlock, error) {
	blocks := genTestDataBlocks(size)
	m, err := New(&Config{
		HashFunc: defaultHashFunc,
	}, blocks)
	if err != nil {
		return nil, nil, err
	}
	return m, blocks, nil
}

func verifySetupParallel(size int) (*MerkleTree, []DataBlock, error) {
	blocks := genTestDataBlocks(size)
	m, err := New(&Config{
		HashFunc:      defaultHashFunc,
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
			name:      "test_pseudo_random_10",
			setupFunc: verifySetup,
			blockSize: 10,
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

func TestVerify(t *testing.T) {
	m, blocks, _ := verifySetup(2)
	// hashFunc is nil
	got, err := Verify(blocks[0], m.Proofs[0], []byte{}, nil)
	if err != nil {
		t.Errorf("Verify() error = %v, wantErr %v", err, nil)
		return
	}
	if got {
		t.Errorf("Verify() got = %v, want %v", got, false)
	}
}

func BenchmarkMerkleTreeNew(b *testing.B) {
	config := &Config{
		HashFunc: defaultHashFunc,
	}
	for i := 0; i < b.N; i++ {
		_, err := New(config, genTestDataBlocks(benchSize))
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}

func BenchmarkMerkleTreeNewParallel(b *testing.B) {
	config := &Config{
		HashFunc:      defaultHashFunc,
		RunInParallel: true,
		NumRoutines:   runtime.NumCPU(),
	}
	for i := 0; i < b.N; i++ {
		_, err := New(config, genTestDataBlocks(benchSize))
		if err != nil {
			b.Errorf("Build() error = %v", err)
		}
	}
}

func TestMerkleTree_GenerateProof(t *testing.T) {
	tests := []struct {
		name    string
		blocks  []DataBlock
		wantErr bool
	}{
		{
			name:   "test_2",
			blocks: genTestDataBlocks(2),
		},
		{
			name:   "test_3",
			blocks: genTestDataBlocks(3),
		},
		{
			name:   "test_4",
			blocks: genTestDataBlocks(4),
		},
		{
			name:   "test_5",
			blocks: genTestDataBlocks(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m1, err := New(nil, tt.blocks)
			if err != nil {
				t.Errorf("m1 New() error = %v", err)
				return
			}
			m2, err := New(&Config{
				Mode: ModeTreeBuild,
			}, tt.blocks)
			if err != nil {
				t.Errorf("m2 New() error = %v", err)
				return
			}
			for idx, block := range tt.blocks {
				got, err := m2.GenerateProof(block)
				if (err != nil) != tt.wantErr {
					t.Errorf("GenerateProof() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, m1.Proofs[idx]) {
					t.Errorf("GenerateProof() %d got = %v, want %v", idx, got, m1.Proofs[idx])
					return
				}
			}
		})
	}
}
