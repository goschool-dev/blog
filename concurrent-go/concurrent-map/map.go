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

// fatal error: concurrent map writes
