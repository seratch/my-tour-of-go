package main

import "fmt"

func fibonacci(ch, quitCh chan int) {
	x, y := 0, 1
	for {
		// A select blocks until one of its cases can run, then it executes that case.
		// It chooses one at random if multiple are ready.
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quitCh:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quitCh := make(chan int)

	// go routine with function literal
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quitCh <- 0
	}()
	fibonacci(ch, quitCh)
}
