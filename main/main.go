package main

import (
	"fmt"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"
	"github.com/luthierycosta/N-puzzle/pq"
)

// Mapa de estados visitados.
// Como não podemos colocar o tipo State como chave, a chave será forma em string do tabuleiro.
var visited map[string]path.Path

func main() {
	ch := make(chan path.Path)
	visited = make(map[string]path.Path)
	begin := state.NewRandom(3)
	end := state.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	})

	go search(begin, end, ch)
	//go search(end, begin, ch)

	result := <-ch
	fmt.Println(result)
}

// Função busca que navega de um estado de início até um estado final
func search(initial, target state.State, ch chan path.Path) {

	compareFunction := func(a, b state.State) bool {
		return a.Heuristic(target) <= b.Heuristic(target)
	}
	pq := pq.New(compareFunction)
	pq.Push(initial)

	for pq.Len() != 0 {
		fmt.Println("FILA:",pq)
		current, _ := pq.Top()
		pq.Pop()

		if state.Equal(current, target) {
			ch <- current.Path
		} else if pathFound, existsKey := visited[fmt.Sprint(current.Board)]; existsKey {
			ch <- current.Path.Join(pathFound)
		} else {
			var lastMove path.Direction
			if pq.Len() == 0 {
				lastMove = path.None
			} else {
				lastMove = current.Path[len(current.Path)-1]
			}
			pq.Push(current.GetNeighbors(lastMove)...)
			visited[fmt.Sprint(current.Board)] = current.Path
		}
	}
}