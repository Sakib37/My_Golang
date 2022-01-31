package main

import "fmt"

func appendSlice() {
	var s []int
	if s == nil {
		fmt.Println("nil!")
	}
	printSlice_3(s)

	test_slice := []int{}
	printSlice_3(test_slice)

	// append works on nil slices.
	//The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	s = append(s, 0)
	printSlice_3(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice_3(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice_3(s)
}

func printSlice_3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
