package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = byte(65) // 'A' has ASCII code 65
	}
	return len(b), nil
}

func exercise() {
	reader.Validate(MyReader{})
}
