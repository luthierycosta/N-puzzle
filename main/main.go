package main

import (
	"fmt"
	"time"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"
)

// Os mapas de estados visitados associam um estado ao caminho para chegar até ele desde um estado inicial dado.
// Como não podemos colocar o tipo State como chave, a chave será forma em string do tabuleiro.
var visited [2]map[string]path.Path

func main() {
	// cria relógio para contar o tempo de execução
	timer := time.Now()
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
	go search(end, begin, ch, 1)
	
	result := <-ch
	defer fmt.Printf("Tempo decorrido na busca: %s\n", time.Since(timer))
	print(begin, result)
}

