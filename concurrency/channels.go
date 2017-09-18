package main

import "fmt"

func calculateSumAndSendToChannel(intArr []int, ch chan int) {
	sum := 0
	for _, v := range intArr {
		sum += v
	}
	ch <- sum // send calculateSumAndSendToChannel to ch
}

func main() {
	intArr := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int)
	go calculateSumAndSendToChannel(intArr[:len(intArr)/2], ch)
	go calculateSumAndSendToChannel(intArr[len(intArr)/2:], ch)
	x, y := <-ch, <-ch // receive from ch

	fmt.Println(x, y, x+y) // -5 17 12
}
