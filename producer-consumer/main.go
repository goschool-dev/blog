package main

import (
	"log"
	"sync"
	"time"
)

func producer(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		c <- i
		log.Println("write:", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
	wg.Done()
}

func consumer(c chan int, wg *sync.WaitGroup) {
	for i := range c {
		log.Println("read:", i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)

	go consumer(ch, &wg)
	go producer(ch, &wg)

	// time.Sleep(10 * time.Second)
	wg.Wait()
}
