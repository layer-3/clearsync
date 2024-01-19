package opendax_protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	t.Parallel()

	msg := Msg{
		Type:   Request,
		ReqID:  42,
		Method: "test",
		Args:   []any{"hello", "there"},
	}

	enc, err := msg.Encode()

	assert.NoError(t, err)
	assert.Equal(t, `[1,42,"test",["hello","there"]]`, string(enc))
}

func TestEncodingEvent(t *testing.T) {
	t.Parallel()

	msg := Msg{
		Type:   EventPrivate,
		ReqID:  42,
		Method: "test",
		Args:   []any{"hello", "there"},
	}

	enc, err := msg.Encode()

	assert.NoError(t, err)
	assert.Equal(t, `[4,"test",["hello","there"]]`, string(enc))
}
