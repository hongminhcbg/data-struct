// mapv2 switch pointer to support heavy read and interval write
package mapv2

import "sync"

type M struct {
	m  map[int]int
	mu sync.Mutex
}

func NewMap() *M {
	return &M{m: make(map[int]int)}
}

func (m *M) Get(key int) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.m[key]
}

func (m *M) Load(new map[int]int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m = new
}
