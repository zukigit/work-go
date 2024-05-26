package main

import (
	"fmt"
	"reflect"
)

func Index_string_int[T string | int](expect T) {
	fmt.Println(reflect.TypeOf(expect))
}

func Index_comparable[T comparable](expect T) {
	fmt.Println(reflect.TypeOf(expect))
}

func Index[T any](expect T) {
	fmt.Println(reflect.TypeOf(expect))
}

func main() {
	Index_string_int("")
	Index_string_int(1616161616161616188)
	Index(false)
}
