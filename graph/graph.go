package graph

import (
	"fmt"
)

type Vertex struct {
	Value int
}

type Graph struct {
	Verts []Vertex
}

func (g *Graph) ShowVerts() {
	for _, vert := range g.Verts {
		fmt.Println(vert.Value)
	}
}