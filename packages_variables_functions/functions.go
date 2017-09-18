package main

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func main() {
	println(add(12, 34))
	println(add2(12, 34))
}
