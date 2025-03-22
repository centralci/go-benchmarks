package maps

import (
	"sync"
	"testing"
)

// RWMutexMap is a map protected by a read-write mutex
type RWMutexMap struct {
	mu    sync.RWMutex
	items map[int]int
}

// NewRWMutexMap creates a new RWMutexMap
func NewRWMutexMap() *RWMutexMap {
	return &RWMutexMap{
		items: make(map[int]int),
	}
}

// Load gets a value from the map
func (m *RWMutexMap) Load(key int) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.items[key]
	return val, ok
}

// Store sets a value in the map
func (m *RWMutexMap) Store(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

// Delete removes a key from the map
func (m *RWMutexMap) Delete(key int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}

// Prefill populates the map with initial data
func prefillRWMutexMap(m *RWMutexMap, count int) {
	for i := 0; i < count; i++ {
		m.Store(i, i)
	}
}

// Prefill populates the sync.Map with initial data
func prefillSyncMap(m *sync.Map, count int) {
	for i := 0; i < count; i++ {
		m.Store(i, i)
	}
}

// BenchmarkRWMutexMapReadHeavy tests RWMutexMap with 90% reads and 10% writes
func BenchmarkRWMutexMapReadHeavy(b *testing.B) {
	m := NewRWMutexMap()
	prefillRWMutexMap(m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			if i%10 == 0 {
				// 10% writes
				m.Store(key, i)
			} else {
				// 90% reads
				m.Load(key)
			}
			i++
		}
	})
}

// BenchmarkSyncMapReadHeavy tests sync.Map with 90% reads and 10% writes
func BenchmarkSyncMapReadHeavy(b *testing.B) {
	var m sync.Map
	prefillSyncMap(&m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			if i%10 == 0 {
				// 10% writes
				m.Store(key, i)
			} else {
				// 90% reads
				m.Load(key)
			}
			i++
		}
	})
}

// BenchmarkRWMutexMapWriteHeavy tests RWMutexMap with 50% reads and 50% writes
func BenchmarkRWMutexMapWriteHeavy(b *testing.B) {
	m := NewRWMutexMap()
	prefillRWMutexMap(m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			if i%2 == 0 {
				// 50% writes
				m.Store(key, i)
			} else {
				// 50% reads
				m.Load(key)
			}
			i++
		}
	})
}

// BenchmarkSyncMapWriteHeavy tests sync.Map with 50% reads and 50% writes
func BenchmarkSyncMapWriteHeavy(b *testing.B) {
	var m sync.Map
	prefillSyncMap(&m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			if i%2 == 0 {
				// 50% writes
				m.Store(key, i)
			} else {
				// 50% reads
				m.Load(key)
			}
			i++
		}
	})
}

// BenchmarkRWMutexMapMixedOps tests RWMutexMap with mixed operations
func BenchmarkRWMutexMapMixedOps(b *testing.B) {
	m := NewRWMutexMap()
	prefillRWMutexMap(m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			op := i % 10

			if op < 7 {
				// 70% reads
				m.Load(key)
			} else if op < 9 {
				// 20% writes
				m.Store(key, i)
			} else {
				// 10% deletes
				m.Delete(key)
			}
			i++
		}
	})
}

// BenchmarkSyncMapMixedOps tests sync.Map with mixed operations
func BenchmarkSyncMapMixedOps(b *testing.B) {
	var m sync.Map
	prefillSyncMap(&m, 1000)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := i % 1000
			op := i % 10

			if op < 7 {
				// 70% reads
				m.Load(key)
			} else if op < 9 {
				// 20% writes
				m.Store(key, i)
			} else {
				// 10% deletes
				m.Delete(key)
			}
			i++
		}
	})
}
