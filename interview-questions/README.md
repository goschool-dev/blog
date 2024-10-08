# Golang Interview Questions

1. How to stop a goroutine?
2. Implement the SumOfSquares function which takes an integer, n and returns the sum of all squares between 1 and n. You’ll need to use select statements, goroutines, and channels. For example, entering 5 would return 55 because (1)^2 + (2)^2 + (3)^2 + (4)^2 + (5)^2 = 55.
3. Using two goroutines one that prints odd numbers and one even numbers, print 1-10 in sequence. 
4. Read and write Fibonacci series to channel.
5. What's wrong this this program.
```
package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))

	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
```

6. Dining Philosophers Problem
```
Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks. Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed. The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.
```

7. Checkpoint Synchronization
```
The checkpoint synchronization is a problem of synchronizing multiple tasks. Consider a workshop where several workers assembling details of some mechanism. When each of them completes his work, they put the details together. There is no store, so a worker who finished its part first must wait for others before starting another one. Putting details together is the checkpoint at which tasks synchronize themselves before going their paths apart.
```

8. Producer Consumer Problem
```
The problem describes two  processes, the producer and the consumer, who share a common, fixed-size buffer used as a queue. The producer's job is to generate data, put it into the buffer, and start again. At the same time, the consumer is consuming the data (i.e., removing it from the buffer), one piece at a time. The problem is to make sure that the producer won't try to add data into the buffer if it's full and that the consumer won't try to remove data from an empty buffer. The solution for the producer is to either go to sleep or discard data if the buffer is full. The next time the consumer removes an item from the buffer, it notifies the producer, who starts to fill the buffer again. In the same way, the consumer can go to sleep if it finds the buffer empty. The next time the producer puts data into the buffer, it wakes up the sleeping consumer.
```

9. Sleeping Barber Problem
```
The barber has one barber's chair in a cutting room and a waiting room containing a number of chairs in it. When the barber finishes cutting a customer's hair, he dismisses the customer and goes to the waiting room to see if there are others waiting. If there are, he brings one of them back to the chair and cuts their hair. If there are none, he returns to the chair and sleeps in it. Each customer, when they arrive, looks to see what the barber is doing. If the barber is sleeping, the customer wakes him up and sits in the cutting room chair. If the barber is cutting hair, the customer stays in the waiting room. If there is a free chair in the waiting room, the customer sits in it and waits their turn. If there is no free chair, the customer leaves.
```

10. Cigarette Smokers Problem
```
Assume a cigarette requires three ingredients to make and smoke: tobacco, paper, and matches. There are three smokers around a table, each of whom has an infinite supply of one of the three ingredients — one smoker has an infinite supply of tobacco, another has paper, and the third has matches. A fourth party, with an unlimited supply of everything, chooses at random a smoker, and put on the table the supplies needed for a cigarrette. The chosen smoker smokes, and the process should repeat indefinitely.
``` 

11. How do you handle HTTP Client server load balancing in Go?

12. How do you handle HTTP client server security in Go?

13. Launch multiple Goroutines and each goroutine adding values to a Channel.

14. Rate Limiter
```
Create a rate limiter that restricts the number of requests processed over a time window. Use a goroutine to check and manage the allowed rate of incoming requests.
```

15. Web Crawler
```
Write a simple web crawler that fetches URLs concurrently. Use goroutines for fetching and channels to communicate the results. Implement a way to limit the number of concurrent requests.
```

16. Bank Account
```
Implement a bank account with deposit and withdrawal methods. Use goroutines to perform these operations concurrently and ensure that the balance remains consistent.
```

17. Fan-Out, Fan-In
```
Create a fan-out, fan-in pattern where multiple goroutines process tasks concurrently, and their results are collected into a single channel. This can be applied to tasks like processing data in parallel.
```

18. File Downloader
```
Implement a file downloader that can download multiple files concurrently. Use goroutines for each download and channels to report progress or completion.
```

19.  Event Counter
```
Create a service that counts events coming from multiple sources concurrently. Use goroutines to process events from each source and aggregate the counts.
```

20. Image Processing Pipeline
```
Build an image processing pipeline where multiple goroutines handle different stages of processing (e.g., loading, filtering, saving). Use channels to pass images between stages.
```

21. Task Queue
```
Implement a task queue system where tasks are submitted by one or more producers and processed by multiple worker goroutines. Ensure proper synchronization and error handling.
```

22. Concurrency Limit
```
Write a program that limits the number of concurrent operations to a specified maximum. For example, manage a pool of worker goroutines that handle incoming tasks without exceeding the limit.
```

23. Blocking and Non-blocking Channels
```
Create a scenario where you demonstrate the difference between blocking and non-blocking channels. Implement a simple producer-consumer model and explore the effects of buffer sizes.
```

24. Timeouts and Contexts
```
Implement a function that fetches data from multiple URLs with timeouts. Use context.Context to manage the deadlines and cancel operations if they exceed the timeout.
```
25. Monitoring Goroutines
```
Develop a simple monitoring tool that tracks the number of running goroutines and displays it at regular intervals. This will help you understand goroutine lifecycles and resource usage.
```
26. Concurrent Fibonacci Calculation
```
Write a function that computes Fibonacci numbers concurrently. Use goroutines for each computation and a channel to collect results, demonstrating parallelism.
```
27. State Machine
```
Create a finite state machine that processes events concurrently. Use goroutines to handle transitions and ensure that the state remains consistent.
```
28. Stock Price Monitor
```
Implement a service that monitors stock prices from multiple sources concurrently. Use goroutines to fetch data and channels to aggregate updates.
```
29. Game State Synchronization
```
Create a simple multiplayer game backend where multiple players can perform actions concurrently. Use goroutines to handle player actions and synchronize game state.
```
30. Parallel Sorting
```
Implement a parallel sorting algorithm using goroutines. Divide the input data into chunks, sort them concurrently, and merge the results.
```
31. Deadlock Detection
```
Design a simple system where you intentionally create deadlocks, then implement a way to detect and resolve them. This will help you understand deadlock conditions in Go.
```
