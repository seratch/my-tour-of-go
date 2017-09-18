package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// toString
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
	String() string
}
*/
type MyStringer fmt.Stringer

func main() {

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	//methods_interfaces/stringer.go:32:15: invalid type assertion: a.(fmt.Stringer) (non-interface type Person on left)
	//stringer := a.(fmt.Stringer)
	//methods_interfaces/stringer.go:34:16: invalid type assertion: a.(MyStringer) (non-interface type Person on left)
	//stringer2 := a.(MyStringer)

	var i interface{} = a
	stringer := i.(fmt.Stringer)
	// stringer:  (0x120a000,0xc420041ee0)
	println("stringer: ", stringer)
	stringer = stringer.(MyStringer)
	// stringer:  (0x120a000,0xc420041ee0)
	println("stringer: ", stringer)

	fmt.Println(a, z)
	// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}
