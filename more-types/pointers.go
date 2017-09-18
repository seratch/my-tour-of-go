package main

func main() {
	i, j := 42, 2701 /* 2701 = 37 * 73 */

	p := &i
	println(p)  // 0xc420041f60
	println(*p) // 42
	*p = 21
	println(*p, i, j) // 21 21 2701

	p = &j
	println(p) // 0xc420041f58
	*p = *p / 37
	println(*p, i, j) // 73 21 73
}
