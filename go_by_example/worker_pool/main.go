package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// accept tasks from jobs until it's closed
	j := 1
	fmt.Println("worker", id, "started job", j)
	time.Sleep(time.Second)
	j = j * 2
	fmt.Println("worker", id, "finished job", j)
	// send result to results channel
}

func main() {
	// solution: https://gobyexample.com/worker-pools
	const numTasks = 5
	const numWorkers = 3
	jobs := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// start all workers
	worker(1, jobs, results)

	// feed tasks to jobs channel and close it after you are done
	// block main thread with results channel receiver
}
