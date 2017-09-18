package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Init(p []int, len int) {
	// p is a slice of array, no need to be a pointer
	for i := 0; i < len; i++ {
		p[i] = int(math.Pow(2.0, float64(i)))
	}
}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	var len = 8
	var pow2 []int = make([]int, len)
	Init(pow2, len)

	for i, v := range pow2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow2 {
		println(v)
	}
}
