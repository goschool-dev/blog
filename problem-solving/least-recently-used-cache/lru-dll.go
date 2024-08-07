package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	order    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

type kv struct {
	key int
	val int
}

func (l *LRUCache) Get(key int) (int, bool) {
	if e, ok := l.cache[key]; ok {
		l.order.MoveToFront(e)
		return e.Value.(*kv).val, true
	}
	return 0, false
}

func (l *LRUCache) Put(key int, value int) {
	if elem, ok := l.cache[key]; ok {
		elem.Value.(*kv).val = value // update existing
		l.order.MoveToFront(elem)
		return
	}

	if l.order.Len() == l.capacity {
		// Remove the oldest item from the cache
		oldest := l.order.Back()
		if oldest != nil {
			l.order.Remove(oldest)
			delete(l.cache, oldest.Value.(*kv).key)
		}
	}
	// Add the new kv
	newElem := l.order.PushFront(&kv{key: key, val: value})
	l.cache[key] = newElem
}

func main() {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
