package main

import "fmt"

func main() {
	i := 10
	if i < 10 {
		panic("!!!")
	} else {
		println(i)
	}

	if j, k := i-5, i-3; j <= 5 && k > 5 {
		fmt.Printf("ok with %d, %d", j, k)
	} else {
		// can access j, k
		fmt.Printf("ng with %d, %d", j, k)
	}
}
