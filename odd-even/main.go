// Using two goroutines one that prints odd numbers and one even numbers, print 1-10 in sequence.
package main

import (
	"fmt"
	"sync"
)

func odd(n int, ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= n; i += 2 {
		<-ch1
		fmt.Println(i)
		if i != n {
			ch2 <- struct{}{}
		}
	}

	close(ch1)
}

func even(n int, ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= n; i += 2 {
		<-ch2
		fmt.Println("		", i)
		if i != n {
			ch1 <- struct{}{}
		}
	}

	close(ch2)
}

func main() {
	n := 28

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)
	go even(n, ch1, ch2, &wg)
	go odd(n, ch1, ch2, &wg)

	ch1 <- struct{}{}
	wg.Wait()
}
