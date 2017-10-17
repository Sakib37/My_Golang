package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primes_slice := [ ]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println(primes_slice)

	//A slice does not store any data, it just describes a section of an underlying array.
	//
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.
	//
	//Other slices that share the same underlying array will see those changes.
	fmt.Println("\n\nArrays as slice: \n")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("\n\nSlice literals: \n")


	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	z := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(z)


	// When slicing, you may omit the high or low bounds to use their defaults instead.
	//The default is zero for the low bound and the length of the slice for the high bound.
	p := []int{2, 3, 5, 7, 11, 13}

	p = p[1:4]
	fmt.Println(p)

	p = p[:2]
	fmt.Println(p)

	p = p[1:]
	fmt.Println(p)
}
