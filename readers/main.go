package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!zzaa")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	test := strings.NewReader("z")
	byte_arr := make([]byte, 1)
	number, err := test.Read(byte_arr)

	fmt.Printf("n = %v err = %v b = %v\n", number, err, byte_arr)
	fmt.Printf("b[:n] = %q\n", b[:number])
}
