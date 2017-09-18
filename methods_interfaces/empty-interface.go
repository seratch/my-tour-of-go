package main

import "fmt"

func main() {
	// empty interface
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	/*
		(<nil>, <nil>)
		(42, int)
		(hello, string)
	*/
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
