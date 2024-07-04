package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (n int, e error) {
	n = copy(b, "A")
	e = nil

	return n, e
}

func main() {
	reader.Validate(MyReader{})
}