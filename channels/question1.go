package main

import (
	"fmt"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

func main() {

	startTime := time.Now()

	ch1 := worker()
	ch2 := worker()

	_, _ = <-ch1, <-ch2

	// _, _,  = <-worker(), <-worker()

	fmt.Println(time.Since(startTime).Seconds())

}
