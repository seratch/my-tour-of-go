package main

import "golang.org/x/tour/reader"

type MyReader struct{}

var AByte = byte('A')

// Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(bytes []byte) (int, error) {
	for idx := range bytes {
		bytes[idx] = AByte
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
