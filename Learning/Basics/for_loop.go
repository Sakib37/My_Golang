package main

import "fmt"

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//go for loop can be converted into a while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
