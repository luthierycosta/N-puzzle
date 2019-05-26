package state

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// testa se pos0 é mesmo a posição do 0 no estado final
func TestPos0Final(t *testing.T) {
	assert := assert.New(t)

	for i := 3; i < 6; i++ {
		state := NewStateFinal(3)
		x := state.pos0.X
		y := state.pos0.Y
		assert.Equal(state.gameBoard[x][y], 0)
	}
}

func TestParity(t *testing.T) {
	// assert := assert.New(t)

	// state := NewStateFinal(4)

	// assert.Equal(state.getParity(), 0)
}
