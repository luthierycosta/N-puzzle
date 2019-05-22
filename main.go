package main

import (
	// "fmt"
	"github.com/luthierycosta/N-puzzle/graph"
)

func main() {
	x := graph.Vertex{Value: 2}
	vertArr := []graph.Vertex{x, x}
	gra := graph.Graph{vertArr}

	gra.ShowVerts()

	// fmt.Println(gra)
}
