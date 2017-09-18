package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	//methods_interfaces/interfaces-satisfied-implicitly.go:20:6: cannot use T literal (type T) as type I in assignment:
	//T does not implement I (missing M method)
	var i I = T{"hello"}
	i.M()
}
