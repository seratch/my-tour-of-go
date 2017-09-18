package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // 8 bytes buffer to read at a time
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	/*
		n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
		b[:n] = "Hello, R"
		n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
		b[:n] = "eader!"
		n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
		b[:n] = ""
	*/
}
