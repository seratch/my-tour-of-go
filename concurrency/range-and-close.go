package main

import (
	"fmt"
)

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	// Note: Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	// sender: ch <- x
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
