package main

import (
    "fmt"
    "github.com/luthierycosta/N-puzzle/pq"
)

func main() {
    fila := pq.New()
    fila.Push(pq.Item{0.3})
    fila.Push(pq.Item{0.2})
    fila.Push(pq.Item{0.45})
    fila.Push(pq.Item{0.7})
    fila.Push(pq.Item{0.4})

    fmt.Println(fila)
}