package main

func swap(x, y string) (string, string) {
	//return (y, x)
	return y, x
}

func main() {
	b, a := swap("a", "b")
	println(a + ", " + b)
}
