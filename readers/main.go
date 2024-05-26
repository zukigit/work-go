package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader(strings.ToUpper("abcdefghijkmnlopqrstuvwxyz"))

	b := make([]byte, 26)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	fmt.Println("------")
	exercise()

	fmt.Println("------")
	exerciseRoot13()
}
