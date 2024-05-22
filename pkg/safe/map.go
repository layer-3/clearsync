package safe

import (
	"sync"
)

// Map is a thread safe generic map.
type Map[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

// NewMap creates a new thread safe map.
func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{
		m: make(map[K]V),
	}
}

// NewMapWithData creates a new thread safe map with initial data.
func NewMapWithData[K comparable, V any](data map[K]V) Map[K, V] {
	return Map[K, V]{m: data}
}

// Load returns the value stored in the map for a key,
// otherwise false if no value is present.
func (m *Map[K, V]) Load(k K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.m[k]
	return v, ok
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(k K, v V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if val, ok := m.m[k]; ok {
		return val, true
	}

	m.m[k] = v
	return v, false
}

// Delete removes a key from the map.
func (m *Map[K, V]) Delete(k K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, k)
}

// Range iterates over the map, calling the provided function
// for each key/value pair. If the function returns false,
// the iteration stops.
func (m *Map[K, V]) Range(f func(k K, v V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}

func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.m)
}

// UpdateInTx allows to update the map in a transactional way.
func (m *Map[K, V]) UpdateInTx(updateFunc func(map[K]V)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	updateFunc(m.m)
}
