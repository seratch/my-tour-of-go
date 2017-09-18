package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var prev float64 = 0.0
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		prev = z
		z = sqrt(z, x)
		if abs(z-prev) < 0.000000001 {
			return z
		}
	}
	return z
}

func sqrt(z float64, x float64) float64 {
	return z - (z*z-x)/(2*z)
}
func abs(z float64) float64 {
	if z < 0 {
		return z * -1
	} else {
		return z
	}
}

func main() {
	for i := 2; i < 100; i++ {
		println("---------------------------")
		fmt.Println(i)
		fmt.Println(math.Sqrt(float64(i)))
		fmt.Println(Sqrt(float64(i)))
	}
}
