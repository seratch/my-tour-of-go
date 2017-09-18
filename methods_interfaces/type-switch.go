package main

import "fmt"

func echoForType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	echoForType(21)
	echoForType("hello")
	echoForType(true)
	/*
		Twice 21 is 42
		"hello" is 5 bytes long
		I don't know about type bool!
	*/
}
