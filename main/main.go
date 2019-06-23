package main

import (
	"fmt"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"
	"github.com/luthierycosta/N-puzzle/pq"
)

// Os mapas de estados visitados associam um estado ao caminho para chegar até ele desde um estado inicial dado.
// Como não podemos colocar o tipo State como chave, a chave será forma em string do tabuleiro.
var visited [2]map[string]path.Path

func main() {
	// inicializa canal e mapas.
	ch := make(chan path.Path)
	visited[0], visited[1] = make(map[string]path.Path), make(map[string]path.Path)
	
	// cria os estados inicial e final
	begin := state.NewRandom(3)
	end := state.New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	})	
	// faz as buscas. 
	go search(begin, end, ch, 0)
	//go search(end, begin, ch, 1)

	result := <-ch
	fmt.Println(result)
}

// Busca que parte de um estado inicial a procura do estado final especificado.
// a função retorna o caminho encontrado no canal ch.
// i é o número da busca (0 ou 1), pois é necessário distinguir já que uma tem que olhar no vetor da outra.
func search(initial, target state.State, ch chan path.Path, i int) {
	// função de comparação de prioridade entre dois estados - checa se f(a) é menor que f(b).
	compareFunction := func(a, b state.State) bool {
		ag, ah := a.Heuristic(target)
		bg, bh := b.Heuristic(target)
		af, bf := ag+ah, bg+bh
		return	ah < bh || af == bf && ah < bh
	}
	pq := pq.New(compareFunction)
	pq.Push(initial)

	for pq.Len() != 0 {
		current, _ := pq.Top()
		fmt.Printf("TOPO: %v, G: %v, H2: %v\n", current.Board, len(current.Path), current.DistanceTo(target))
		pq.Pop()

		// Se o estado atual é o final, solução encontrada. Retorne o caminho associado ao atual.
		if state.Equal(current, target) {
			ch <- current.Path
		}
		// Se o estado atual já foi visitado pela outra busca, a solução é a junção dos caminhos das 2 buscas.
		if pathFound, existsKey := visited[1-i][current.ToString()]; existsKey {
			ch <- current.Path.Join(pathFound)
		}
		// Marca como visitado no próprio vetor.
		visited[i][current.ToString()] = current.Path
		// Para cada vizinho possível a partir do atual, insira-o na fila se ele ainda não fora visitado.
		for _, neighbor := range current.GetNeighbors() {
			if _, alreadyVisited := visited[i][neighbor.ToString()]; !alreadyVisited {
				pq.Push(neighbor)
			}
		}
	}
}