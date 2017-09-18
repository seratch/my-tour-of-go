package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	/*
		(<nil>, *main.T)
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10952d4]

		goroutine 1 [running]:
		main.(*T).M(0x0)
			/Users/kazuhirosera/github/oss/a-tour-of-go/methods_interfaces/interfaces-values-nil.go:18 +0x34
		main.main()
			/Users/kazuhirosera/github/oss/a-tour-of-go/methods_interfaces/interfaces-values-nil.go:27 +0x47
		exit status 2
	*/
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t // nil
	describe(i)
	i.M() // panic: runtime error: invalid memory address or nil pointer dereference

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
