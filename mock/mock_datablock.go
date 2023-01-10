// Package mock provides a mock implementation of the DataBlock interface.
package mock

// DataBlock is a mock implementation of the DataBlock interface.
type DataBlock struct {
	Data []byte
}

// Serialize returns the serialized form of the DataBlock.
func (t *DataBlock) Serialize() ([]byte, error) {
	return t.Data, nil
}
