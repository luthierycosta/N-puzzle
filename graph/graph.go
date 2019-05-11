package graph

import (
    "fmt"
)

type Graph struct {
    Vertices []*Vertex
}

type Vertex struct {
    Value int
    Adjacents []*Vertex
}

func NewGraph(vertices ...int) *Graph {
    g := new(Graph)
    for i := range vertices {
        g.AddVertex(vertices[i])
    }
    return g
}

func NewVertex(v int) *Vertex {
    return &Vertex{Value: v}
}

func (g *Graph) AddVertex(v int) {
    g.Vertices = append(g.Vertices, NewVertex(v))
}

func (g *Graph) findVertex(v int) *Vertex {
    for _, vertex := range g.Vertices {
        if vertex.Value == v {
            return vertex
        }
    }
    return nil			// tratamento de erro 10/10
}

func (g *Graph) AddEdge(e1, e2 int) {
    v1, v2 := g.findVertex(e1), g.findVertex(e2)
    if v1 == nil || v2 == nil {
        fmt.Println("deu ruim man. nao tem esse vertice ai nao"); return
    } else {
        v1.Adjacents = append(v1.Adjacents, v2)
    }
}

func (g *Graph) ExistsEdge(e1, e2 int) bool {
    v1, v2 := g.findVertex(e1), g.findVertex(e2)
    for i := range v1.Adjacents {
        if v1.Adjacents[i] == v2 {
            return true
        }
    }
    return false
}

func (g *Graph) Size() int {
    return len(g.Vertices)
}

func (g *Graph) ShowVerts() {
    for _, vert := range g.Vertices {
        fmt.Println(vert.Value)
    }
}

