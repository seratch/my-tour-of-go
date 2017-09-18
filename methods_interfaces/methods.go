package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v Vertex) Echo() {
	fmt.Println(v)
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scaled(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

type MyInteger int

func (i MyInteger) Echo() {
	println(i)
}

func main() {
	v := Vertex{3, 4}
	println("Initial", v.Abs())
	v.Echo()

	v.Scale(2.0)
	v.Echo()

	v2 := v.Scaled(3.0)
	v.Echo()
	v2.Echo()

	MyInteger(123).Echo()
}
