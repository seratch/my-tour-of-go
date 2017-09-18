package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//fmt.Println(<-ch)
	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()
			/Users/kazuhirosera/github/oss/a-tour-of-go/concurrency/bufferred-channels.go:11 +0x272

		Process finished with exit code 0
	*/
}
