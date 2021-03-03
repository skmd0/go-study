package main

func main() {
	// solution: https://gobyexample.com/range-over-channels
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// range over channel data and print each result
}
