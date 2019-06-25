package state

// import (
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// 	"github.com/luthierycosta/N-puzzle/path"
// )

// func TestGetNeighbors(t *testing.T) {
// 	assert := assert.New(t)

// 	// GetNeighbors nao deve alterar o campo original
// 	stateA := NewRandom(3)
// 	stateACopy := stateA.makeCopy()
// 	assert.Equal(stateA, stateACopy)
// 	stateA.GetNeighbors(path.Down)
// 	assert.Equal(stateA, stateACopy)

// 	// quando o 0 esta no canto, so deve ter uma jogada possivel
// 	cantoInfDir := New([][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 0} })
// 	infDirNeigh := cantoInfDir.GetNeighbors(path.Down)
// 	assert.Equal(len(infDirNeigh), 1)
// 	assert.Equal(New([][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 0, 8} }), infDirNeigh[0])

// 	cantoSupEsq := New([][]int{
// 		{0, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 1} })
// 	supEsqNeigh := cantoSupEsq.GetNeighbors(path.Left)
// 	assert.Equal(len(supEsqNeigh), 1)
// 	assert.Equal(New([][]int{
// 		{4, 2, 3},
// 		{0, 5, 6},
// 		{7, 8, 1} }), supEsqNeigh[0])

// 	// se o 0 esta na borda, so devem ter 2 jogadas possiveis
// 	bordaSup := New([][]int{
// 		{1, 0, 3}, {4, 5, 6}, {7, 8, 2}})
// 	bordaSupNei1 := bordaSup.GetNeighbors(path.Right)
// 	bordaSupNei2 := bordaSup.GetNeighbors(path.Up)
	
// 	assert.Equal(2, len(bordaSupNei1))
// 	assert.Equal(2, len(bordaSupNei2))

// 	assert.Equal([]State{
// 		New([][]int {{1, 5, 3}, {4, 0, 6}, {7, 8, 2}}),
// 		New([][]int {{1, 3, 0}, {4, 5, 6}, {7, 8, 2}}) },
// 		bordaSupNei1 )
// 	assert.Equal([]State{
// 		New([][]int {{0, 1, 3}, {4, 5, 6}, {7, 8, 2}}),
// 		New([][]int {{1, 3, 0}, {4, 5, 6}, {7, 8, 2}}) },
// 		bordaSupNei2 )

// 	// se 0 nao esta na borda nem no canto, devem ter 3 jogadas possiveis
// 	meio := New([][]int {
// 		{1, 2, 3},
// 		{4, 0, 6},
// 		{7, 8, 5} })
// 	meioNeigh := meio.GetNeighbors(path.Right)

// 	assert.Equal([]State{
// 		New([][]int{{1, 0, 3}, {4, 2, 6}, {7, 8, 5}}),
// 		New([][]int{{1, 2, 3}, {4, 8, 6}, {7, 0, 5}}),
// 		New([][]int{{1, 2, 3}, {4, 6, 0}, {7, 8, 5}}) },
// 		meioNeigh)

// }