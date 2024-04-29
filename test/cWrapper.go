package test

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: ./Cfiles/test.o
#include "headers/test.h"
*/
import "C"

func callTest() int {
	return int(C.add(67, 2))
}
