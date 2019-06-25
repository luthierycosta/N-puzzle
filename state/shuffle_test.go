package state

import (
	"github.com/stretchr/testify/assert"
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
	
	s := NewRandom(3)
	perm := getPermutation(s)
	t.Log(perm)
	t.Log(isSolvable(perm, len(s.Board), s.findPos(0).X))	
}

func TestPos0Shuffle(t *testing.T) {
	assert := assert.New(t)
	
	for i := 3; i < 6; i++ {
		state := NewRandom(3)
		x,y := state.pos0.X, state.pos0.Y
		assert.Equal(state.Board[y][x], 0)
	}
	state := NewRandom(6)
	for i := 0; i < 36; i++ {
		assert.NotEqual(state.findPos(i), Pair{-1,-1})
	}
}

func TestIsSolvable(t *testing.T) {
	assert := assert.New(t)

	valid1 :=  []int{ 1, 8, 2, 0, 4, 3, 7, 6, 5 }
	valid2 :=  []int{ 13, 2, 10, 3, 1, 12, 8, 4, 5, 0, 9, 6, 15, 14, 11, 7 }
	valid3 :=  []int{ 6, 13, 7, 10, 8, 9, 11, 0, 15, 2, 12, 5, 14, 3, 1, 4 }

	invalid1 := []int{ 3, 9, 1, 15, 14, 11, 4, 6, 13, 0, 10, 12, 2, 7, 8, 5 }

	assert.True(isSolvable(valid1, 3, 1))
	assert.True(isSolvable(valid2, 4, 2))
	assert.True(isSolvable(valid3, 4, 1))
	assert.False(isSolvable(invalid1, 4, 2))
}