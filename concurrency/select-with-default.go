package main

import (
	"fmt"
	"time"
)

func main() {
	// receive-only type <-chan Time
	var tick <-chan time.Time = time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case t := <-tick:
			fmt.Printf("tick. %v\n", t)
		case b := <-boom:
			fmt.Printf("BOOM! %v\n", b)
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
