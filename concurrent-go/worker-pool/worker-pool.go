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
