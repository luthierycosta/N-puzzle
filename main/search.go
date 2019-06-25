package main

import (
	"fmt"
	"sync"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"
	"github.com/luthierycosta/N-puzzle/pq"
)

var mutex sync.Mutex

// Busca que parte de um estado inicial a procura do estado final especificado.
// a função retorna o caminho encontrado no canal ch.
// i é o número da busca (0 ou 1), pois é necessário distinguir já que uma tem que olhar no vetor da outra.
func search(initial, target state.State, ch chan path.Path, i int) {
	// função de comparação de prioridade entre dois estados - checa se f(a) é menor que f(b).
	compareFunction := func(a, b state.State) bool {
		ag, ah := a.Heuristic(target)
		bg, bh := b.Heuristic(target)
		af, bf := ag+ah, bg+bh
		return	ah < bh || ah == bh && af < bf
	}
	pq := pq.New(compareFunction)
	pq.Push(initial)

	for pq.Len() != 0 {
		current, _ := pq.Top()
		fmt.Printf("TOPO: %v, G: %v, H: %v, F: %v\n", current.ToString(), len(current.Path), current.DistanceTo(target), len(current.Path)+current.DistanceTo(target))
		pq.Pop()
		
		// Pega o lock, marca como visitado no próprio vetor e vê se ele já foi visitado pela outra.
		mutex.Lock()
		pathFound, visitedbyOther	:= visited[1-i][current.ToString()]
		visited[i][current.ToString()] = current.Path
		mutex.Unlock()

		// Se o estado atual é o final, solução encontrada. Retorne o caminho associado ao atual.
		if state.Equal(current, target) {
			ch <- current.Path
		}
		// Se o estado atual já foi visitado pela outra busca, a solução é a junção dos caminhos das 2 buscas.
		if visitedbyOther {
			if i == 0 {
				ch <- current.Path.Join(pathFound)
			} else {
				ch <- pathFound.Join(current.Path)
			}
		}
		// Para cada vizinho possível a partir do atual, insira-o na fila se ele ainda não fora visitado.
		for _, neighbor := range current.GetNeighbors() {
			if _, alreadyVisited := visited[i][neighbor.ToString()]; !alreadyVisited {
				pq.Push(neighbor)
			}
		}
	}
}