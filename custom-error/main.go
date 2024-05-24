package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s at %v", e.What, e.When)
}

func (e MyError) String() string {
	return fmt.Sprintf("What: %s, When: %v",
		e.What, e.When)
}

func main() {
	error := MyError{
		time.Now(),
		"Get fucked up",
	}
	fmt.Println(error)
	fmt.Println("Exercise-----")
	ErrorExercise()
}
