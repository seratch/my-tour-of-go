package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var prev float64 = 0.0
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		prev = z
		z = sqrt(z, x)
		if abs(z-prev) < 0.000000001 {
			return z, nil
		}
	}
	return z, nil
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

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
