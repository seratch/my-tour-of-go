package main

//var c, python, java bool

func echo(x *string) {
	println(x)
}

func main() {
	var i int
	i = 123
	println(i)

	var j, k int
	j = 234
	k = 345
	println(j, k)

	var foo string
	println(foo)
	echo(nil)
	var c, python, java = "ccc", true, false
	println(c, python, java)
}
