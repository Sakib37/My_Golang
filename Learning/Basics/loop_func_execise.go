package main

import (
	"fmt"
)

func SqrtLoop(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func loopExercise() {
	fmt.Println(SqrtLoop(2))
}
