# Worker pools in Go

A worker pool is a great way to handle concurrent processing efficiently. The concept is beneficial in various scenarios where tasks need to be handled in parallel, but you want to control the level of concurrency and resource usage.Let's implement a worker pool in Go and understand how it is useful in the context of database operations.

Let us first define the scenario:
You need to process a large number of records from a database, where each record requires some computational work or update operations. This scenario involves performing multiple concurrent database operations, such as batch processing records or executing parallel queries.
Here two things are significant:

* __Task__: A task represents a record or a batch of records to be processed. We can define it as:

```
type Task struct {
	ID int
}
```

* __Worker__: A fixed number of worker goroutines handle these tasks concurrently. We can define it as:

```
type WorkerPool struct {
	count int         // The number of worker goroutines.
	tasks chan Task   // Channel to communicate tasks between the main goroutine and worker goroutines.
	wg    sync.WaitGroup  // Waitgroup to wait for all workers to finish.
}
```

Let's define some methods on WorkerPool struct.
* __Start()__ initializes the worker pool and starts processing tasks.

```
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}
```

* __worker()__ processes tasks from channel.

```
func (wp *WorkerPool) worker(workerID int) {
	defer wp.wg.Done()
	for task := range wp.tasks {
		wp.processTask(workerID, task)
	}
}
```

* __processTask()__ does some task processing.

```
func (wp *WorkerPool) processTask(workerID int, task Task) {
	fmt.Printf("Worker %d processing task %d\n", workerID, task.ID)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Worker %d completed task %d\n", workerID, task.ID)
}
```

* __stop()__ closes the task channel and wait for workers to finish working.

```
func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
}
```

Now let us use our worker pool. Here is a complete program.

```
package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type WorkerPool struct {
	count int
	tasks chan Task
	wg    sync.WaitGroup
}

func NewWorkerPool(count int) *WorkerPool {
	return &WorkerPool{
		tasks: make(chan Task),
		count: count,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.count; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(workerID int) {
	defer wp.wg.Done()
	for task := range wp.tasks {
		wp.processTask(workerID, task)
	}
}

func (wp *WorkerPool) processTask(workerID int, task Task) {
	fmt.Printf("Worker %d processing task %d\n", workerID, task.ID)
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Printf("Worker %d completed task %d\n", workerID, task.ID)
}

func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
}

func main() {
	// Create a worker pool with 3 workers.
	workerPool := NewWorkerPool(3)

	// Start the worker pool.
	workerPool.Start()

	// Enqueue tasks.
	for i := 1; i <= 10; i++ {
		workerPool.tasks <- Task{ID: i}
	}

	// Gracefully shut down the worker pool.
	workerPool.Stop()
}
```

The main function creates a worker pool with 3 workers and submits 10 tasks. It waits for a while to simulate task processing abd stops the worker pool and waits for all workers to complete. This approach ensures that all tasks are processed concurrently by a fixed number of workers and that the pool shuts down gracefully once all tasks are completed.
