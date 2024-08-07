package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int         // capacity of cache
	cache    map[int]int // stores key:value
	index    map[int]int //	stores the index of the key in usage array
	usage    []int       // stores the recently used keys, latest at the end
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]int, capacity),
		index:    make(map[int]int),
		usage:    make([]int, 0, capacity),
	}
}

func (l *LRUCache) Put(k int, v int) {
	_, ok := l.cache[k]
	if ok {
		l.cache[k] = v
		l.moveToTop(k)
		return
	}

	if len(l.cache) >= l.capacity {
		lru := l.usage[0]
		l.usage = l.usage[1:]
		delete(l.cache, lru)
		delete(l.index, lru)
	}
	l.cache[k] = v
	l.usage = append(l.usage, k)
	l.index[k] = len(l.usage) - 1
}

func (l *LRUCache) Get(k int) int {
	v, ok := l.cache[k]
	if ok {
		l.moveToTop(k)
		return v
	}
	return -1
}

func (l *LRUCache) moveToTop(k int) {
	position := l.index[k]
	if position == len(l.usage)-1 {
		return
	}
	l.usage = append(l.usage[:position], l.usage[position+1:]...)
	l.usage = append(l.usage, k)
	for i, v := range l.usage {
		l.index[v] = i
	}
}

func main() {
	lru := NewLRUCache(2)
	lru.Put(1, 10)
	lru.Put(2, 20)
	fmt.Println(lru.Get(1)) // returns 10
	lru.Put(3, 30)
	lru.Put(4, 40)          // removes 2
	fmt.Println(lru.Get(2)) // returns -1
	fmt.Println(lru.Get(3)) // returns 30
	fmt.Println(lru.Get(4)) // returns 40
}
