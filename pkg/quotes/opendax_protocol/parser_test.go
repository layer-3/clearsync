package opendax_protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRaw(t *testing.T) {
	t.Parallel()

	t.Run("Successful test", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,42,"ping",[]]`))
		assert.NoError(t, err)
		assert.Equal(t,
			&Msg{
				Type:   Request,
				ReqID:  42,
				Method: "ping",
				Args:   []any{},
			}, msg)
		msg, err = ParseRaw([]byte(`[2,42,"pong",[]]`))
		assert.NoError(t, err)
		assert.Equal(t,
			&Msg{
				Type:   Response,
				ReqID:  42,
				Method: "pong",
				Args:   []any{},
			}, msg)
		msg, err = ParseRaw([]byte(`[3,"temperature",[28.7]]`))
		assert.NoError(t, err)
		assert.Equal(t,
			&Msg{
				Type:   EventPublic,
				ReqID:  0,
				Method: "temperature",
				Args:   []any{28.7},
			}, msg)
	})

	t.Run("Invalid message length", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,42,"ping"]`))
		assert.EqualError(t, err, "message must contain 4 elements")
		assert.Nil(t, msg)
	})

	t.Run("Bad JSON", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,42,"ping",[]`))
		assert.EqualError(t, err, "could not parse message: unexpected end of JSON input")
		assert.Nil(t, msg)
	})

	t.Run("Invalid types", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[5,42,"ping",[]]`))
		assert.EqualError(t, err, "message type must be 1, 2, 3 or 4")
		assert.Nil(t, msg)

		msg, err = ParseRaw([]byte(`[5,"ping",[]]`))
		assert.EqualError(t, err, "message type must be 1, 2, 3 or 4")
		assert.Nil(t, msg)

		msg, err = ParseRaw([]byte(`[1.1,42,"pong",[]]`))
		assert.EqualError(t, err, "failed to parse type: invalid type")
		assert.Nil(t, msg)

		msg, err = ParseRaw([]byte(`["1",42,"pong",[]]`))
		assert.NoError(t, err)
		assert.NotNil(t, msg)
	})

	t.Run("Invalid request ID", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,"42","ping",[]]`))
		assert.NoError(t, err)
		assert.NotNil(t, msg)

		msg, err = ParseRaw([]byte(`[1,42.1,"ping",[]]`))
		assert.EqualError(t, err, "failed to parse request ID: invalid type")
		assert.Nil(t, msg)
	})

	t.Run("Invalid method", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,42,true,[]]`))
		assert.EqualError(t, err, "failed to parse method: invalid type")
		assert.Nil(t, msg)
	})

	t.Run("Invalid arguments", func(t *testing.T) {
		t.Parallel()

		msg, err := ParseRaw([]byte(`[1,42,"ping",true]`))
		assert.EqualError(t, err, "failed to parse arguments: invalid type")
		assert.Nil(t, msg)

		msg, err = ParseRaw([]byte(`[1,42,"ping","hello"]`))
		assert.EqualError(t, err, "failed to parse arguments: invalid type")
		assert.Nil(t, msg)
	})
}
