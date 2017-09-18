package main

func split(sum int) (x, y int) {
	x = sum / 3
	y = sum - x
	return
}

func double(i int) (j int) {
	j = i * 2
	return
}

func main() {
	a, b := split(3)
	println(a, b)
	println(double(12))

}
