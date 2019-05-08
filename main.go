package main

import (
	// "fmt"
	. "./graph"
)

func main() {
	x := Vertex{Value: 2}
	vertArr := []Vertex{x, x}
	gra := Graph{vertArr}

	gra.ShowVerts()

	// fmt.Println(gra)
}