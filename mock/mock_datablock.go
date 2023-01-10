package mock

type DataBlock struct {
	Data []byte
}

func (t *DataBlock) Serialize() ([]byte, error) {
	return t.Data, nil
}
