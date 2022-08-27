package merkletree

import (
	"crypto/sha256"
	"errors"
	"hash"
	"reflect"
	"testing"
)

type mockHash struct {
}

func NewMock() hash.Hash {
	return &mockHash{}
}

func (m *mockHash) Reset() {
}

func (m *mockHash) BlockSize() int {
	return 0
}

func (m *mockHash) Size() int {
	return 0
}

func (m *mockHash) Sum(in []byte) []byte {
	return in
}

func (m *mockHash) Write([]byte) (nn int, err error) {
	return 0, errors.New("mockHash.Write error")
}

func Test_defaultHashFunc(t *testing.T) {
	sha256Digest = NewMock()
	defer func() {
		sha256Digest = sha256.New()
	}()
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test_write_err",
			args: args{
				data: []byte{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultHashFunc(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultHashFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultHashFunc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
