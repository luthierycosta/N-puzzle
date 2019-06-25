package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewState(t *testing.T) {
	assert := assert.New(t)

	state := New([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9,10,11,12},
		{13,14,15,0},
	})
	size := len(state.Board)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i == size-1 && j == size-1) {
				assert.Equal(state.Board[i][j], 0)
			} else {
				assert.Equal(state.Board[i][j], i * size + j + 1)
			}
		}
	}
}

// testa se pos0 é mesmo a posição do 0 no estado final
func TestPos0(t *testing.T) {
	assert := assert.New(t)

	final := New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	})
	assert.Equal(final.findPos(0), Pair{2,2})

	final.Board[2][2] = 10
	assert.Equal(final.findPos(0), Pair{-1,-1})
}

func TestMakeCopy(t *testing.T) {
	assert := assert.New(t)

	for i := 3; i < 10; i++ {
		state := NewRandom(i)
		copy := state.makeCopy()
		assert.Equal(state, copy)
		
		state.Board[0][0] = 100
		assert.NotEqual(state, copy)
	}
}

func TestToString(t *testing.T) {
	assert := assert.New(t)

	for i := 3; i <= 6; i++ {
		s1 := NewRandom(i)
		s2 := NewRandom(i)
		t.Logf("%v\n%v\n", s1.ToString(), s2.ToString())
	}

	s1 := New([][]int{
		{6, 13, 7, 10},
		{8, 9, 11, 0},
		{15, 2, 12, 5},
		{14, 3, 1, 4},
	})
	s2 := s1.makeCopy()
	assert.Equal(s1.ToString(), s2.ToString())
	assert.True(Equal(s1, s2))
}