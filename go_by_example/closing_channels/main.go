package main

import "fmt"

func main() {
	// solution: https://gobyexample.com/closing-channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// receive and print each job
			// if there are no jobs, send data to done channel
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close jobs channel after all data is sent and print "all jobs sent"
	<-done
}
