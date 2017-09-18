package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x1094ef8]

		goroutine 1 [running]:
		main.main()
			/Users/kazuhirosera/github/oss/a-tour-of-go/methods_interfaces/nil-interface-values.go:12 +0x38
	*/
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
