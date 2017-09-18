package main

func main() {
	defer println("after main function return")
	println("first of all")
	sample()

	println("counting")
	for i := 0; i < 10; i++ {
		// pushing the deferred functions onto a stack
		defer println(i)
		// IDEA detects possible resource leak here!
	}
	println("done")

	println("end of main")
}

func sample() {
	defer println("after sample function return")
	println("do something in sample")
}
