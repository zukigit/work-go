package main

import (
	"fmt"
	"math"
)

type ErrNegSqrt float64

func (e ErrNegSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegSqrt(x)
	}
	return math.Sqrt(x), nil
}

func ErrorExercise() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
