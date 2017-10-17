package main

import (
	"fmt"
	"strings"
)


//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
//
//The make function allocates a zeroed array and returns a slice that refers to that array:
//
//a := make([]int, 5)  // len(a)=5
//
//To specify a capacity, pass a third argument to make:
//
//b := make([]int, 0, 5) // len(b)=0, cap(b)=5
//
//b = b[:cap(b)] // len(b)=5, cap(b)=5
//b = b[1:]      // len(b)=4, cap(b)=4

func main() {
	a := make([]int, 5)
	printSlice_2("a", a)

	b := make([]int, 0, 5)
	printSlice_2("b", b)

	c := b[:2]
	printSlice_2("c", c)

	d := c[2:5]
	printSlice_2("d", d)



	// Slice of slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice_2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
