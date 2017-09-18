package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(*v)
}

func main() {
	v := Vertex{1, 2}
	// methods_interfaces/functions.go:14:11: cannot use v (type Vertex) as type *Vertex in argument to ScaleFunc
	// ScaleFunc(v, 2.0)
	ScaleFunc(&v, 2.0)

	p := &Vertex{2, 3}
	ScaleFunc(p, 3.0)
}
