package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsIterator(t *testing.T) {
	t.Parallel()

	t.Run("no errors", func(t *testing.T) {
		t.Parallel()

		args := []interface{}{
			float64(1),
			"123",
			float64(3),
			"145",
			float64(3),
			"145",
			"56.89",
			"hello",
			float64(0.65),
			true,
			12345.565656,
			[]interface{}{1, 2, "hello"},
		}
		it := NewArgIterator(args)

		uintv := it.NextUint64()
		assert.NoError(t, it.Err())
		assert.Equal(t, uint64(1), uintv)

		uintv = it.NextUint64()
		assert.NoError(t, it.Err())
		assert.Equal(t, uint64(123), uintv)

		int64v := it.NextInt64()
		assert.NoError(t, it.Err())
		assert.Equal(t, int64(3), int64v)

		int64v = it.NextInt64()
		assert.NoError(t, it.Err())
		assert.Equal(t, int64(145), int64v)

		intv := it.NextInt()
		assert.NoError(t, it.Err())
		assert.Equal(t, int(3), intv)

		intv = it.NextInt()
		assert.NoError(t, it.Err())
		assert.Equal(t, int(145), intv)

		floatv := it.NextFloat64()
		assert.NoError(t, it.Err())
		assert.Equal(t, float64(56.89), floatv)

		stringv := it.NextString()
		assert.NoError(t, it.Err())
		assert.Equal(t, "hello", stringv)

		floatv = it.NextFloat64()
		assert.NoError(t, it.Err())
		assert.Equal(t, float64(0.65), floatv)

		boolv := it.NextBool()
		assert.NoError(t, it.Err())
		assert.Equal(t, true, boolv)

		stringv = it.NextString()
		assert.NoError(t, it.Err())
		assert.Equal(t, "12345.565656", stringv)

		slicev := it.NextSlice()
		assert.NoError(t, it.Err())
		assert.Equal(t, slicev, []interface{}{1, 2, "hello"})

		it.NextUint64()
		assert.Equal(t, it.Err(), ErrIterationDone)
	})

	t.Run("no errors", func(t *testing.T) {
		t.Parallel()

		args := []interface{}{
			"123",
			"56.89",
			"hello",
		}
		it := NewArgIterator(args)

		var res []string
		for it.More() {
			res = append(res, it.NextString())
		}

		assert.NoError(t, it.Err())
		assert.Equal(t, []string{"123", "56.89", "hello"}, res)
	})
}
