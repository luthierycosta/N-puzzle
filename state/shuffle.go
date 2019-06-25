package state

import "math/rand"

func (s *State) shuffle(seed int64) {
	r := rand.New(rand.NewSource(seed)) // create new random source
	bSize := len(s.Board)           // get board size
	perm := r.Perm(bSize * bSize)       // get random permutarion of size n^2

	row0 := -1
	for i, val := range perm {
		if val == 0 {
			row0 = bSize - (i/bSize)
		}
	}

	// while isnt solvable, get new permutation
	for !isSolvable(perm, bSize, row0) {
		newSeed := r.Int63()
		r.Seed(newSeed)
		perm = r.Perm(bSize * bSize)
		for i, val := range perm {
			if val == 0 {
				row0 = bSize - (i/bSize)
			}
		}
	}

	// replace all blocks with the random ones from the permutation
	for i, randIndex := range perm {
		row := i / bSize
		col := i % bSize
		s.Board[row][col] = randIndex
	}
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

	return bSize%2 == 1 && inversion%2 == 0 || bSize%2 == 0 && row0%2+inversion%2 == 1

	// logic taken from https://www.geeksforgeeks.org/check-instance-15-puzzle-solvable/
}
