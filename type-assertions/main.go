package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	var i interface{} = student{
		name: "zuki",
		age:  23,
	}

	s, ok := i.(student) //it's really cool
	fmt.Println(s, ok)

	str, ok := i.(string)
	fmt.Println(str, ok)

	switch i.(type) {
	case student:
		fmt.Println("It's student object!")
	default:
		fmt.Println("Other object!")
	}
}
