package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "a"
	fmt.Println(a)
	a[1] = "b"
	fmt.Println(a)
	// more-types/array.go:11:3: invalid array index 2 (out of bounds for 2-element array)
	// a[2] = "c"

	//for i := 0; i < 5; i++ {
	//	a[i] = fmt.Sprintf("no.%d", i)
	//	//panic: runtime error: index out of range
	//	//	goroutine 1 [running]:
	//	//	main.main()
	//	//	/Users/kazuhirosera/github/oss/a-tour-of-go/more-types/array.go:15 +0x549
	//}

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices, length, capacity
	s := primes[1:4]
	fmt.Println(s, len(s), cap(s)) // [3 5 7] 3 5

	s2 := primes[1:]
	fmt.Println(s2, len(s2), cap(s2)) // [3 5 7 11 13] 5 5

	s3 := primes[:3]
	fmt.Println(s3, len(s3), cap(s3)) // [2 3 5] 3 6

	// slices are like references to arrays
	s[0] = 10001
	fmt.Println(s)      // [10001 5 7]
	fmt.Println(primes) // [2 10001 5 7 11 13]

	q := []int{1, 2, 3, 4}
	fmt.Println(q)

	ss := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
	}
	fmt.Println(ss)

	{
		a1 := make([]int, 5)              // len 5, cap 5
		fmt.Println(a1, len(a1), cap(a1)) // [0 0 0 0 0] 5 5

		a2 := make([]int, 3, 5)           // len 3, cap 5
		fmt.Println(a2, len(a2), cap(a2)) // [0 0 0] 3 5
	}

	{
		aa := [][]string{[]string{"a", "b", "c"}, []string{"A", "B", "C"}}
		fmt.Println(aa)                       // [[a b c] [A B C]]
		fmt.Println(strings.Join(aa[0], ",")) // a,b,c

	}

}
