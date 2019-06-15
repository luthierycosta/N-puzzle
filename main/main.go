package main

import (
	"fmt"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"
	"github.com/luthierycosta/N-puzzle/pq"
)

// Mapa de estados visitados.
// Como não podemos colocar o tipo State como chave, a chave será forma em string do tabuleiro.
var visited [2]map[string]path.Path

func main() {
	ch := make(chan path.Path)
	visited[0], visited[1] = make(map[string]path.Path), make(map[string]path.Path)
	begin := state.NewRandom(3)
	end := state.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	})

	go search(begin, end, ch, 0)
	//go search(end, begin, ch, 1)

	result := <-ch
	fmt.Println(result)
}

// Função busca que navega de um estado de início até um estado final
func search(initial, target state.State, ch chan path.Path, i int) {

	compareFunction := func(a, b state.State) bool {
		if ah, bh := a.Heuristic(initial, target), b.Heuristic(initial, target); ah < bh {
			return true
		} else if ah == bh {
			return len(a.Path) < len(b.Path)
		} else {
			return false
		}
	}
	pq := pq.New(compareFunction)
	pq.Push(initial)

	for pq.Len() != 0 {
		current, _ := pq.Top()
		fmt.Println("TOPO:", current.Board, "H:", current.DistanceTo(initial), current.DistanceTo(target))
		pq.Pop()

		if state.Equal(current, target) {
			ch <- current.Path
		} else if pathFound, existsKey := visited[1-i][fmt.Sprint(current.Board)]; existsKey {
			ch <- current.Path.Join(pathFound)
		} else {
			var lastMove path.Direction
			if pq.Len() == 0 {
				lastMove = path.None
			} else {
				lastMove = current.Path[len(current.Path)-1]
			}
			pq.Push(current.GetNeighbors(lastMove)...)
			visited[i][fmt.Sprint(current.Board)] = current.Path
		}
	}
}