package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"github.com/luthierycosta/N-puzzle/path"
)

func TestNewState(t *testing.T) {
	assert := assert.New(t)

	state := New([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9,10,11,12},
		{13,14,15,0},
	})
	size := len(state.gameBoard)

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
func TestPos0(t *testing.T) {
	assert := assert.New(t)

	final := New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	})
	assert.Equal(final.findPos(0), Pair{2,2})

	final.gameBoard[2][2] = 10
	assert.Equal(final.findPos(0), Pair{-1,-1})
}

func TestPos0Shuffle(t *testing.T) {
	assert := assert.New(t)
	
	for i := 3; i < 6; i++ {
		state := NewRandom(3)
		x,y := state.pos0.X, state.pos0.Y
		assert.Equal(state.gameBoard[y][x], 0)
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

func TestMakeCopy(t *testing.T) {
	//assert := assert.New(t)

	for i := 3; i < 10; i++ {
		state := NewRandom(i)
		copy := state.makeCopy()
		//assert.Equal(state, copy)
		fmt.Println("state:",state)
		fmt.Println("copy:",copy)
		state.gameBoard[0][0] = 10
		//assert.NotEqual(state, copy)
		fmt.Println("new state:",state)
		fmt.Println("new copy:",copy)
	}
}

func TestGetNeighbors(t *testing.T) {
	assert := assert.New(t)

	// getNeighbors nao deve alterar o campo original
	stateA := NewRandom(3)
	stateACopy := stateA.makeCopy()
	assert.Equal(stateA, stateACopy)
	stateA.getNeighbors(path.Down)
	assert.Equal(stateA, stateACopy)

	// quando o 0 esta no canto, so deve ter uma jogada possivel
	cantoInfDir := New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0} })
	infDirNeigh := cantoInfDir.getNeighbors(path.Down)
	assert.Equal(len(infDirNeigh), 1)
	assert.Equal(New([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8} }), infDirNeigh[0])

	cantoSupEsq := New([][]int{
		{0, 2, 3},
		{4, 5, 6},
		{7, 8, 1} })
	supEsqNeigh := cantoSupEsq.getNeighbors(path.Left)
	assert.Equal(len(supEsqNeigh), 1)
	assert.Equal(New([][]int{
		{4, 2, 3},
		{0, 5, 6},
		{7, 8, 1} }), supEsqNeigh[0])

	// se o 0 esta na borda, so devem ter 2 jogadas possiveis
	bordaSup := New([][]int{
		{1, 0, 3}, {4, 5, 6}, {7, 8, 2}})
	bordaSupNei1 := bordaSup.getNeighbors(path.Right)
	bordaSupNei2 := bordaSup.getNeighbors(path.Up)
	
	assert.Equal(2, len(bordaSupNei1))
	assert.Equal(2, len(bordaSupNei2))

	assert.Equal([]State{
		New([][]int {{1, 5, 3}, {4, 0, 6}, {7, 8, 2}}),
		New([][]int {{1, 3, 0}, {4, 5, 6}, {7, 8, 2}}) },
		bordaSupNei1 )
	assert.Equal([]State{
		New([][]int {{0, 1, 3}, {4, 5, 6}, {7, 8, 2}}),
		New([][]int {{1, 3, 0}, {4, 5, 6}, {7, 8, 2}}) },
		bordaSupNei2 )

	// se 0 nao esta na borda nem no canto, devem ter 3 jogadas possiveis
	meio := New([][]int {
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 5} })
	meioNeigh := meio.getNeighbors(path.Right)

	assert.Equal([]State{
		New([][]int{{1, 0, 3}, {4, 2, 6}, {7, 8, 5}}),
		New([][]int{{1, 2, 3}, {4, 8, 6}, {7, 0, 5}}),
		New([][]int{{1, 2, 3}, {4, 6, 0}, {7, 8, 5}}) },
		meioNeigh)

}