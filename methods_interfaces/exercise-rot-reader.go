package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// https://en.wikipedia.org/wiki/ROT13
type rot13Reader struct {
	r io.Reader
}

var from = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var to = []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")

func rot13(b byte) byte {
	//// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
	//func IndexByte(s []byte, c byte) int // ../runtime/asm_$GOARCH.s
	idx := bytes.IndexByte(from, b)
	if idx == -1 {
		return b
	} else {
		return to[idx]
	}
}

func (rot rot13Reader) Read(barr []byte) (int, error) {
	n, e := rot.r.Read(barr)
	for i := 0; i < n; i++ {
		barr[i] = rot13(barr[i])
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
var Rot13Map = make(map[byte]byte)

func InitIfNeeded() {
	if len(Rot13Map) == 0 {
		input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		toBe := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
		for idx, b := range input {
			Rot13Map[b] = toBe[idx]
		}
	}
}

func (rot rot13Reader) Read(barr []byte) (int, error) {
	InitIfNeeded()
	n, e := rot.r.Read(barr)
	for i := 0; i < n; i++ {
		to := Rot13Map[barr[i]]
		if to != 0 {
			barr[i] = to
		}
	}
	return n, e
}
*/
