package state

import (
	"testing"
)

func getPermutation(s State) []int {
	perm := []int{}
	for i := range s.Board {
		for j := range s.Board[i] {
			perm = append(perm, s.Board[i][j])
		}
	}
	return perm
}

func TestNewRandom(t *testing.T) {
	
	s := New([][]int{
		{2, 8, 3},
		{1, 6, 4},
		{7, 0, 5},
	})
	perm := getPermutation(s)
	t.Log(perm)
	t.Log(isSolvable(perm, len(s.Board), s.findPos(0).X))	
}