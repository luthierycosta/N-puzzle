package state

import (
	"fmt"
	"github.com/luthierycosta/N-puzzle/path"
)

// Pair representa uma casa (x,y) no tabuleiro do jogo.
type Pair struct {X, Y int}

// State representa um determinado estado no tabuleiro,
// com os valores das peças na matriz Board e a posição atual do 0 (espaço vazio) armazenada.
type State struct {
	Board     [][]int
	pos0      Pair
	Path      path.Path
}

// New State aloca e retorna um novo estado, cujo tabuleiro é recebido de argumento.
func New(board [][]int) State {
	state := State{}
	n := len(board)
	state.Board = make([][]int, n)
	for i := 0; i < n; i++ {
		state.Board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			state.Board[i][j] = board[i][j]
		}
	}
	state.pos0 = state.findPos(0)
	return state
}

// NewRandom aloca e retorna um novo estado de tabuleiro nxn
// preenchido aleatoriamente com valores de 0 a (n^2)-1.
func NewRandom(n int) State {
	state := State{}
	state.Board = make([][]int, n)
	for i := 0; i < n; i++ {
		state.Board[i] = make([]int, n)
	}
	state.shuffle(88)
	state.pos0 = state.findPos(0)
	return state
}

// makeCopy retorna uma cópia do estado s.
func (s State) makeCopy() State {
	copy := New(s.Board)
	copy.Path = make([]path.Direction, len(s.Path))
	for i := range copy.Path {
		copy.Path[i] = s.Path[i]
	}
	return copy
}

// ToString retorna a forma em string do tabuleiro em s.
func (s State) ToString() (res string) {
	res += "["
	for i := range s.Board {
		res += " ["
		for j := range s.Board[i] {
			res += fmt.Sprintf("%02d ", s.Board[i][j])
		}
		res += "]"
	}
	res += " ]"
	return
}

func (s State) Print() {
	for i := range s.Board {
		fmt.Printf("[ ")
		for j := range s.Board[i] {
			fmt.Printf("%02d ", s.Board[i][j])
		}
		fmt.Printf("]\n")
	}
}

// Encontra a posição (x,y) do bloco k no tabuleiro do estado a.
func (s State) findPos(k int) Pair {
    for i := range s.Board {
        for j := range s.Board[i] {
            if s.Board[i][j] == k {
                return Pair{j, i}
            }
        }
    }
    return Pair{-1,-1}      // alguma exceção que eu não sei implementar
}

// Equal checa se dois estados têm o mesmo tabuleiro.
func Equal(a, b State) bool {
	return a.ToString() == b.ToString()
}