package main

func main() {
	// solution: https://gobyexample.com/non-blocking-channel-operations
	messages := make(chan string)

	// create a non-blocking channel pattern:
	// 1 - receive from channel
	// 2 - send to channel
	// 3 - print "no action"
}
