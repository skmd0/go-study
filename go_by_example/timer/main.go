package main

import (
	"fmt"
	"time"
)

func main() {
	// solution: https://gobyexample.com/timers
	var timer1 *time.Timer
	// create a 2 second timer on timer1
	<-timer1.C
	fmt.Println("Timer 1 fired")

	var timer2 *time.Timer
	// create a 1 second timer
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	var isStopped bool
	// stop the timer before it triggers and print message if the timer is stopped
	if isStopped {
		fmt.Println("Timer 2 stopped")
	}
}
