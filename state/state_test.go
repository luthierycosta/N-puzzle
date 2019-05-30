package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewState(t *testing.T) {
	assert := assert.New(t)

	size := 6
	state := NewState(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i == size-1 && j == size-1) {
				assert.Equal(state.gameBoard[i][j], 0)
			} else {
				assert.Equal(state.gameBoard[i][j], i * size + j + 1)
			}
		}
	}
}

// testa se pos0 é mesmo a posição do 0 no estado final
func TestPos0Final(t *testing.T) {
	assert := assert.New(t)

	for i := 3; i < 6; i++ {
		state := NewState(3)
		x := state.pos0.X
		y := state.pos0.Y
		assert.Equal(state.gameBoard[x][y], 0)
	}
}

func TestIsSolvable(t *testing.T) {
	assert := assert.New(t)

	valid1 :=  []int{ 1, 8, 2, 0, 4, 3, 7, 6, 5 }
	valid2 :=  []int{ 13, 2, 10, 3, 1, 12, 8, 4, 5, 0, 9, 6, 15, 14, 11, 7 }
	valid3 :=  []int{ 6, 13, 7, 10, 8, 9, 11, 0, 15, 2, 12, 5, 14, 3, 1, 4 }

	invalid1 := []int{ 3, 9, 1, 15, 14, 11, 4, 6, 13, 0, 10, 12, 2, 7, 8, 5 }

	assert.True(isSolvable(valid1, 3, Pair{0, 1}))
	assert.True(isSolvable(valid2, 4, Pair{1, 2}))
	assert.True(isSolvable(valid3, 4, Pair{3, 1}))
	assert.False(isSolvable(invalid1, 4, Pair{1, 2}))
}

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	state := NewState(3)

	state.shuffle(10)
	assert.Equal([][]int{{1, 2, 8}, {4, 6, 3}, {5, 7, 0}}, state.gameBoard)
}