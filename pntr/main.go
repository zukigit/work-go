package main

import "fmt"

func doubleNumber(number *int) {
	*number = *number * 2
}

func main() {
	fmt.Println("Pointer test")
	num := 12
	doubleNumber(&num)
	fmt.Println("After doubleNumber(), num:", num)
}
