package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic
	fmt.Println(f)
	/*
		panic: interface conversion: interface {} is string, not float64

		goroutine 1 [running]:
		main.main()
			/Users/kazuhirosera/github/oss/a-tour-of-go/methods_interfaces/type-assertions.go:17 +0x67d
	*/
}
