# Implement a simple LRU (Least Recently Used) cache that:
* Stores a fixed number of items.
* Evict the least recently used item when the cache is full.
* Provide methods to Set and Get key-value pairs.

When we discuss about cache two terms that comes into mind is in-memory and key-value pairs. One way to implement LRU cache in Go is using maps as it can store key-value pairs in memory. Other way to implement it is using doubly linked list, we can easily remove the least recently used element from the tail and add the most recently at the top. In this blog we will see both the ways, let's start with maps.

We can define the structure like this:
```
type LRUCache struct {
	capacity int         // capacity of cache
	cache    map[int]int // stores key:value
	index    map[int]int //	stores the index of the key in usage array
	usage    []int       // stores the recently used keys, latest at the end
}
```

* Capacity is the maximum number of items cache can store.
* Cache is our key-value store.
* Usage is an array that stores the recently used keys, appends the most recent one at the last.
* Index is a map that stores the cache key and its index in the array for fast lookups.

Now, we can write a function that builds our LRUCache.
```
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]int, capacity),
		usage:    make([]int, 0, capacity),
        index:    make(map[int]int),
	}
}
```

Now let us write the Set() function that inserts a key-value into our cache.
```
func (l *LRUCache) Set(k int, v int) {
	_, ok := l.cache[k]
	if ok { // if element already exists
		l.cache[k] = v // update value
		l.moveToTop(k) 
		return
	}

	if len(l.cache) >= l.capacity { // if cache is fully filled
		lru := l.usage[0] // oldest element
		l.usage = l.usage[1:] // remove oldest from usage
		delete(l.cache, lru) // remove from cache
		delete(l.index, lru) // remove from index
	}

	l.cache[k] = v
	l.usage = append(l.usage, k)
	l.index[k] = len(l.usage) - 1
}
```

We can write a private function that moves the most recently used element to the top of the array.
```
func (l *LRUCache) moveToTop(k int) {
	position := l.index[k]
	if position == len(l.usage)-1 { 
        // already at top do nothing
		return
	}

    // delete the element
	l.usage = append(l.usage[:position], l.usage[position+1:]...)

    // append it to the end
	l.usage = append(l.usage, k)

    // update the indexes
	for i, v := range l.usage {
		l.index[v] = i
	}
}
```

Now lets complete our cache implementation by adding Get() function that retrieves the value associated with the key if it exists and moves it to the front of the index slice.
```
func (l *LRUCache) Get(k int) (int, bool) {
	v, ok := l.cache[k]
	if ok {
		l.moveToTop(k)
		return v, true
	}
	return 0, false
}
```

This is how we can implement LRU cache in Go using maps. You can find the full code [here](https://github.com/goschool-dev/blog/blob/master/problem-solving/least-recently-used-cache/lru-maps.go).

This way of implementing LRU involves maintaining order of keys based on recent usage and shifting keys in the slice, which can be less efficient for large caches.
Let's implement it again using doubly linked list. Here we will use the "contianer/list" package which implements a doubly linked list. 

We can define our LRU structure like this:
```
type LRUCache struct {
	capacity int // maximum capacity of the cache
	cache    map[int]*list.Element // stores key-value
	order    *list.List // maintains order of usage
}
```

Lets write a function that build our cache.
```
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}
```

Now, lets specify Set() and Get() methods.
```
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

func (l *LRUCache) Get(key int) (int, bool) {
	if e, ok := l.cache[key]; ok {
		l.order.MoveToFront(e)
		return e.Value.(*kv).val, true
	}
	return 0, false
}
```

In this way we can implement Least recently used cache in Golang. The doubly linked list approach is generally preferred in practice due to its efficiency, especially for larger caches where performance is crucial. However, for simpler or smaller use cases, the map and slice approach might be sufficient and more straightforward.

You can find the full code [here](https://github.com/goschool-dev/blog/blob/master/problem-solving/least-recently-used-cache/lru-dll.go).
