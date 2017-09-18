package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		println("macOS")
	case "linux":
		println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 17:
		println("Good afternoon!")
	default:
		println("Good evening!")
	}
}
