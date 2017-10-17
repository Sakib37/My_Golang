package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordmap := make(map[string]int)
    for _, value := range strings.Fields(s){
    	wordmap[value]++
	}
	return wordmap
}

func main() {
	fmt.Println(WordCount(" This This is A String"))
}
