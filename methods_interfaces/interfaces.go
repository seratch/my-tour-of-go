package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	var a Abser
	v := Vertex{1, 2}
	a = &v
	fmt.Println(a.Abs())

	f := MyFloat(2.3)
	a = f
	fmt.Println(a.Abs())
}
