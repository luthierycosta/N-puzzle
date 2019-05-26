package state

import (
	"math/rand"
)

type Pair struct {
	X int
	Y int
}

type State struct {
	gameBoard [][]int
	pos0      Pair
}

func NewState(n int) State {
	state := State{}
	state.gameBoard = make([][]int, n)
	for i := 0; i < n; i++ {
		state.gameBoard[i] = make([]int, n)
		for j := 0; j < n; j++ {
			state.gameBoard[i][j] = i*n + j + 1
			if j == n-1 && i == n-1 {
				state.gameBoard[i][j] = 0
			}
		}
	}
	state.pos0 = Pair{n - 1, n - 1}

	return state
}

func (s *State) shuffle(seed int64) {
	r := rand.New(rand.NewSource(seed))	// create new random source
	bSize := len(s.gameBoard)						// get board size
	perm := r.Perm(bSize * bSize)				// get random permutarion of size n^2

	// TODO:
	// cheque aqui que a paridade da permutacao eh segura

	// replace all blocks with the random ones from the permutation
	for i, randIndex := range perm {
		row := i / bSize
		col := i % bSize
		s.gameBoard[row][col] = randIndex
	}
}