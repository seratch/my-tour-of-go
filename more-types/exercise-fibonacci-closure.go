package main

import "fmt"

type Fibonacci struct {
	previous, current int
}

func (fib *Fibonacci) IsInitial() bool {
	return fib.previous == 0 && fib.current == 0
}

func (fib *Fibonacci) Next() int {
	switch {
	case fib.IsInitial():
		fib.current = 1
	default:
		result := fib.previous + fib.current
		fib.previous = fib.current
		fib.current = result
	}
	return fib.current
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := Fibonacci{}
	return func() int {
		return fib.Next()
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	/*
		1
		1
		2
		3
		5
		8
		13
		21
		34
		55
	*/
}
