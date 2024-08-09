# Concurrent Map in Go
In this blog we will try to understand the issues with concurrent access to a map and design a concurrent map in Go to tackle those issues.

Let us understand the problem first. Have a look at the below program and understand what it is trying to do.

```
package main

import (
	"fmt"
	"time"
)

func insert(m map[int]int, k, v int) {
	fmt.Println("setting key:", k)
	m[k] = v
}

func remove(m map[int]int, k int) {
	time.Sleep(time.Second * 2)
	fmt.Println("removing key:", k)
	delete(m, k)
}

func main() {
	m := map[int]int{}

	for i := 0; i < 10; i++ {
		go insert(m, i, i)
		if i%2 == 0 {
			go remove(m, i)
		}
	}

	time.Sleep(time.Second * 3) // wait for goroutines to finish
}
```

You can see the program spawns multiple goroutine calls to insert and remove a key from a map. On running this program you would encounter runtime panics due to concurrent map access. The go runtime will display an error __concurrent map writes__. This shows that Go maps are not safe for concurrent use, which means you can not access map from multiple goroutines withour proper synchronization.

Lets design a concurrent map that safely allows concurrent reads and writes. It should provide basic map methods like get(), set(), delete(), len().
We will use the built int sync package that provies the synchronization primitives like __sync.Mutex__. We can define our concurrent map struct like this:

```
type Map struct {
	mu   *sync.RWMutex
	data map[int]int 
}
```

Here we have two fields:

* __data__: map that stores keys and values.
* __mu__: sync.RWMutex{} that allows multiple readers or one writer, ensuring that read operations can occur concurrently, while write operations are mutually exclusive.

Now lets define map methods for this struct.

```
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
```

Here we have defined four methods on our structure:

* __Set(k, v int)__: Acquires a write lock, sets or updates the key-value pair, and then releases the lock.
* __Get(k int)__: Acquires a read lock, retrieves the value for the given key, and releases the lock.
* __Delete(k int)__: Acquires a write lock, removes the key from the map, and releases the lock.
* __Len()__: Acquires a read lock and returns the number of elements in the map.

Now if you update the same program with our concurrent map implementation and run it, you won't encounter any runtime panics for concurrent access. This implementation provides safe concurrent access to the map, ensuring that multiple goroutines can read from the map simultaneously, but write operations are performed exclusively.

You can find the full code [here](https://github.com/goschool-dev/blog/blob/master/concurrent-go/concurrent-map/concurrent-map.go).
