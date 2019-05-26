package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestShuffle(t *testing.T) {
	assert := assert.New(t)

	state := NewState(3)

	state.shuffle(10)
	assert.Equal([][]int{{1, 2, 8}, {4, 6, 3}, {5, 7, 0}}, state.gameBoard)
}
