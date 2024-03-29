package main

import "fmt"

type VertexMAP struct {
	Lat, Long float64
}

//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
//
//The make function returns a map of the given type, initialized and ready for use.
var m map[string]VertexMAP

func mapFunc() {
	m = make(map[string]VertexMAP)
	//m := map[string]VertexMAP{}
	fmt.Println(m)
	m["Bell Labs"] = VertexMAP{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m2 = map[string]VertexMAP{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)

	// Mutating Maps
	//Delete an element:
	//delete(m, key)
	//Test that a key is present with a two-value assignment:
	//elem, ok = m[key]
	//If key is in m, ok is true. If not, ok is false.
	m3 := make(map[string]int)
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
