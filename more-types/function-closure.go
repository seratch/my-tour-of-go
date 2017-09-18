package main

import "fmt"

func adder() func(int) int {
	sum := 0
	// this function is bound to the local variables which exist in the scope when defining the function
	// in this case, the `sum` is referred by the function
	return func(x int) int {
		fmt.Println("x: ", x, "sum: ", sum)
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
