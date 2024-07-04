package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = reader.r.Read(b)
	
	if err != nil {
		return n, err
	}
	
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])	
	}

	return n, nil
}

func rot13(x byte) byte {
    input := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
    output := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
    match := bytes.Index(input, []byte{x})
	
    if match == -1 {
        return x
    }
	
    return output[match]
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}