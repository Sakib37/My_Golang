package main

import "fmt"

type Vertex struct{ X, Y int }

func vertexFunc() {
	V := Vertex{1, 2}
	fmt.Println(V)
	fmt.Println(Vertex{1, 2})

	V.X = 20
	fmt.Println(V)

	// Pointer to struct
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference.
	P := &V
	P.X = 80
	fmt.Println(V)
	fmt.Println(*P)

	// Struct Literals
	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, p, v2, v3)

}
