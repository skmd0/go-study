package main

import (
	"fmt"
	"time"
)

func main() {
	// solution: https://gobyexample.com/tickers
	var ticker1 *time.Ticker
	// create a new ticker with 500 ms trigger time
	done := make(chan bool)

	// in a goroutine print returned time from a ticker until it stops

	time.Sleep(1600 * time.Millisecond)
	ticker1.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
