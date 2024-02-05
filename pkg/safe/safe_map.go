package safe

import "sync"

// Thread safe generic map. Uses a RWMutex to allow concurrent reads and
// exclusive writes.
type Map[K comparable, V any] struct {
	m *sync.Map
}

// New creates a new thread safe map.
func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		m: &sync.Map{},
	}
}

func NewMapWithData[K comparable, V any](data map[K]V) Map[K, V] {
	m := &sync.Map{}

	for k, v := range data {
		m.Store(k, v)
	}

	return Map[K, V]{
		m: m,
	}
}

// Load returns the value stored in the map for a key, or false if no value is
// present.
func (m *Map[K, V]) Load(k K) (V, bool) {
	v, ok := m.m.Load(k)
	return v.(V), ok
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(k K, v V) {
	m.m.Store(k, v)
}

// Delete removes a key from the map.
func (m *Map[K, V]) Delete(k K) {
	m.Delete(k)
}

// Range iterates over the map, calling the provided function for each key/value
// pair. If the function returns false, the iteration stops.
func (m *Map[K, V]) Range(f func(k K, v V) bool) {
	m.Range(f)
}

// UpdateInTx allows to update the map in a transactional way.
func (m *Map[K, V]) UpdateInTx(updateFunc func(map[K]V)) {
	m.UpdateInTx(updateFunc)
}
