package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func main() {
	// create a worker goroutine and block main thread until it finishes
	// solution: https://gobyexample.com/channel-synchronization
	go worker()
}
