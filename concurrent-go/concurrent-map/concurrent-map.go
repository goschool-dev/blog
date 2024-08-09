package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	mu   *sync.RWMutex
	data map[int]int
}

func NewMap() *Map {
	return &Map{
		mu:   &sync.RWMutex{},
		data: map[int]int{},
	}
}

func (m *Map) Set(k, v int) {
	m.mu.Lock()
	m.data[k] = v
	m.mu.Unlock()
}

func (m *Map) Get(k int) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[k]
	return v, ok
}

func (m *Map) Delete(k int) {
	m.mu.Lock()
	delete(m.data, k)
	m.mu.Unlock()
}

func (m *Map) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

func insert(m *Map, k, v int) {
	fmt.Println("setting key:", k)
	m.Set(k, v)
}

func remove(m *Map, k int) {
	fmt.Println("removing key:", k)
	m.Delete(k)
}

func main() {
	cm := NewMap()

	for i := 0; i < 10; i++ {
		go insert(cm, i, i)
		if i%2 == 0 {
			go remove(cm, i)
		}
	}

	time.Sleep(time.Second * 3) // wait for goroutines to finish
}
