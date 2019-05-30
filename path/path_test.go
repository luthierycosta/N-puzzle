package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeMove(t *testing.T) {
	assert := assert.New(t)

	path := Path{}
	path = path.MakeMove(Up)
	path = path.MakeMove(Up)
	path = path.MakeMove(Down)
	path = path.MakeMove(Down)
	path = path.MakeMove(Left)
	path = path.MakeMove(Right)
	path = path.MakeMove(Left)
	path = path.MakeMove(Right)

	assert.Equal(path, Path{
		Up, Up, Down, Down, Left, Right, Left, Right})
}