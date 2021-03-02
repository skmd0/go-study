package main

import (
	"time"
)

func fetchUpdates(ch chan string) {
	// external resource that needs 2 seconds to fetch results
	time.Sleep(2 * time.Second)
	ch <- "result successful"
}

func main() {
	// solution: https://gobyexample.com/timeouts
	ch := make(chan string, 2)

	go fetchUpdates(ch)
	// implement timeout pattern - print channel result or timeout msg (1 sec)

	go fetchUpdates(ch)
	// implement timeout pattern - print channel result or timeout msg (3 sec)
}
