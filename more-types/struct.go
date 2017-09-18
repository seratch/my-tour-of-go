package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{X: 1, Y: 2})

	v := Vertex{
		X: 1,
		Y: 2,
	}
	v.Y = 10
	println(v.X, v.Y)

	p := &v
	(*p).X = 10
	// without the explicit dereference
	p.Y = 20
	println(v.X, v.Y)

	// literals
	{
		var (
			v1 = Vertex{}
			v2 = Vertex{X: 1}
			p  = &Vertex{}
		)
		fmt.Println(v1, v2, p)
	}

	// https://github.com/golang/go/issues/19885
	// more-types/struct.go:12:9: illegal types for operand: print
	// Vertex
	//println(Vertex{1, 2})
}
