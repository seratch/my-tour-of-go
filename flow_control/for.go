package main

func main() {
	{
		sum := 0
		//for (i := 0; i < 10; i++) {
		for i := 0; i < 10; i++ {
			sum += i
		}
		println(sum)
	}

	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		println(sum)
	}

	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		println(sum)
	}

	{
		// IDEA detects infinite loop!
		//for {
		//}
	}
}
