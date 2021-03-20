package main

import (
	"fmt"
	"sync"
)

func main() {
	// solution: https://gobyexample.com/atomic-counters
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// fix this by using sync/atomic
				ops += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
