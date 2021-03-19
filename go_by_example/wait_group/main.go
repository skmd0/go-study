package main

import (
	"fmt"
	"sync"
	"time"
)

// notify the WaitGroup that job is done
func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("worker", id, "starting...")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "finished")
}

func main() {
	// solution: https://gobyexample.com/waitgroups

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		// increase the WaitGroup size
		go worker(i, &wg)
	}
	// block main thread until work is done
}
