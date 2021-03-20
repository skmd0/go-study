package main

import (
	"fmt"
	"time"
)

func main() {
	// solutiion: https://gobyexample.com/rate-limiting
	basicRateLimiting()
	//burstRateLimiting()
}

func basicRateLimiting() {
	const reqNum = 20
	requests := make(chan int, reqNum)
	for i := 1; i <= reqNum; i++ {
		requests <- i
	}
	close(requests)

	// create a basic rate limiter with 500 ms delay, print each request
}

func burstRateLimiting() {
	// no need to change any of this code, just a complete example of a burst rate limiting
	const reqNum = 20
	const reqLimit = 3
	burstLimiter := make(chan time.Time, reqLimit)
	// fill up with 3 values just for simulating a burst of request
	for i := 0; i < reqLimit; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			for i := 0; i < reqLimit; i++ {
				burstLimiter <- t
			}
		}
	}()

	burstRequests := make(chan int, reqNum)
	for i := 1; i <= reqNum; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request", req, time.Now())
	}
}
