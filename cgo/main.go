package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: ./Cfiles/test.o
#include "headers/test.h"
*/
import "C"
import "fmt"

func main() {
	result := C.add(1, 2)
	fmt.Println(result)
}
