package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//In if statement parentheses ( ) are not mandatory but the braces { } are required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	//Check this answer for clarification about "Sprint"
	//https://stackoverflow.com/questions/45203052/why-would-you-use-fmt-sprint/45203310#45203310
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//Like for, the if statement can start with a short statement to execute before the condition
	// Here 'v < lim' is the condition for if statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Pow(x, n, lim float64) float64 {
	//Variables declared inside an if short statement are also available inside any of the else blocks.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func conditionFunc() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//Variables declared inside an if short statement are also available inside any of the else blocks.
	fmt.Println(
		Pow(3, 2, 10),
		Pow(3, 3, 20),
	)
}
