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

func isSolvable(perm []int, bSize int, row0 int) bool {
	inversion := 0
	bSize2 := bSize * bSize
	for i := 0; i < bSize2; i++ { // calculando o numero de inversoes
		for j := i; j < bSize2; j++ {
			if perm[i] > perm[j] && perm[j] != 0 {
				inversion++
			}
		}
	}

	return (bSize%2 == 1 && inversion%2 == 0) || (row0%2+inversion%2 == 1)

	// logic taken from https://www.geeksforgeeks.org/check-instance-15-puzzle-solvable/
}

func (s *State) shuffle(seed int64) {
	r := rand.New(rand.NewSource(seed)) // create new random source
	bSize := len(s.gameBoard)           // get board size
	perm := r.Perm(bSize * bSize)       // get random permutarion of size n^2

	row0 := -1
	for i, val := range perm {
		if val == 0 {
			row0 = i
		}
	}

	// while isnt solvable, get new permutation
	for !isSolvable(perm, bSize, row0) {
		newSeed := r.Int63()
		r.Seed(newSeed)
		perm = r.Perm(bSize * bSize)
		for i, val := range perm {
			if val == 0 {
				row0 = i
			}
		}
	}

	// replace all blocks with the random ones from the permutation
	for i, randIndex := range perm {
		row := i / bSize
		col := i % bSize
		s.gameBoard[row][col] = randIndex
	}
}

func (s State) makeCopy() State {
	copy := State{[][]int{}, s.pos0}
	for _, arr := range s.gameBoard {
		line := []int{}
		for _, val := range arr {
			line = append(line, val)
		}
		copy.gameBoard = append(copy.gameBoard, line)
	}
	return copy
}