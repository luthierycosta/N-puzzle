package state

import "time"

// Pair representa uma casa (x,y) no tabuleiro do jogo.
type Pair struct {X, Y int}

// State representa um determinado estado no tabuleiro,
// com os valores das peças na matriz gameBoard e a posição atual do 0 (espaço vazio) armazenada.
type State struct {
	gameBoard [][]int
	pos0      Pair
}

// New State aloca e retorna um novo estado, cujo tabuleiro é recebido de argumento.
func New(board [][]int) State {
	state := State{}
	n := len(board)
	state.gameBoard = make([][]int, n)
	for i := 0; i < n; i++ {
		state.gameBoard[i] = make([]int, n)
		for j := 0; j < n; j++ {
			state.gameBoard[i][j] = board[i][j]
		}
	}
	state.pos0 = state.findPos(0)
	return state
}

// NewRandom aloca e retorna um novo estado de tabuleiro nxn
// preenchido aleatoriamente com valores de 0 a (n^2)-1.
func NewRandom(n int) State {
	state := State{}
	state.gameBoard = make([][]int, n)
	for i := 0; i < n; i++ {
		state.gameBoard[i] = make([]int, n)
	}
	state.shuffle(time.Now().UnixNano())
	state.pos0 = state.findPos(0)
	return state
}

func (s State) makeCopy() State {
	return New(s.gameBoard)
}

// Encontra a posição (x,y) do bloco k no tabuleiro do estado a.
func (a State) findPos(k int) Pair {
    for i := range a.gameBoard {
        for j := range a.gameBoard[i] {
            if a.gameBoard[i][j] == k {
                return Pair{j, i}
            }
        }
    }
    return Pair{-1,-1}      // alguma exceção que eu não sei implementar
}