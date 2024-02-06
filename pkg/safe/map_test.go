package safe

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	t.Parallel()

	m := NewMap[string, int]()
	assert.NotNil(t, m.m, "Should create a non-nil map")
}

func TestNewMapWithData(t *testing.T) {
	t.Parallel()

	initialData := map[string]int{"key": 1}
	m := NewMapWithData(initialData)
	assert.Equal(t, 1, m.m["key"], "Should initialize map with correct data")
}

func TestMap_Load(t *testing.T) {
	t.Parallel()

	m := NewMap[string, int]()
	m.Store("key", 2)
	v, ok := m.Load("key")
	assert.True(t, ok, "Load should find the key")
	assert.Equal(t, 2, v, "Should return the correct value")
}

func TestMap_Store(t *testing.T) {
	t.Parallel()

	m := NewMap[string, int]()
	m.Store("key", 2)
	assert.Equal(t, 2, m.m["key"], "Should correctly set the value")
}

func TestMap_Delete(t *testing.T) {
	t.Parallel()

	m := NewMap[string, int]()
	m.Store("key", 2)
	m.Delete("key")
	_, ok := m.Load("key")
	assert.False(t, ok, "Should remove the key")
}

func TestMap_Range(t *testing.T) {
	t.Parallel()

	t.Run("Should iterate over all elements", func(t *testing.T) {
		t.Parallel()

		m := NewMap[string, int]()
		m.Store("key1", 1)
		m.Store("key2", 2)

		var count int
		m.Range(func(k string, v int) bool {
			count++
			return true
		})

		assert.Equal(t, 2, count, "Should iterate over all elements")
	})

	t.Run("Must obey the stop condition", func(t *testing.T) {
		t.Parallel()

		m := NewMap[string, int]()
		m.Store("key1", 1)
		m.Store("key2", 2)

		var count int
		m.Range(func(k string, v int) bool {
			count++
			return false
		})

		assert.Equal(t, 1, count, "Should break iteration after the first function call")
	})
}

func TestMap_UpdateInTx(t *testing.T) {
	t.Parallel()

	m := NewMap[string, int]()
	m.Store("key1", 1)

	m.UpdateInTx(func(m map[string]int) {
		m["key1"] = 2
		m["key2"] = 3
	})

	assert.Equal(t, 2, m.m["key1"], "Should update key1 correctly")
	assert.Equal(t, 3, m.m["key2"], "Should add and update key2 correctly")
}

func TestMap_ConcurrentAccess(t *testing.T) {
	t.Parallel()

	m := NewMap[int, int]()
	var wg sync.WaitGroup

	// Perform concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			m.Store(val, val)
		}(i)
	}

	wg.Wait()

	// Verify concurrent writes
	for i := 0; i < 100; i++ {
		v, ok := m.Load(i)
		assert.True(t, ok, "Concurrent Store/Load should find key")
		assert.Equal(t, i, v, "Concurrent Store/Load should return the correct value for key")
	}
}
