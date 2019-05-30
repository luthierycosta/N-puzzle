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

func TestJoinPath(t *testing.T) {
	assert := assert.New(t)

	path1 := Path{}
	path1 = path1.MakeMove(Up)
	path1 = path1.MakeMove(Down)
	path1 = path1.MakeMove(Left)
	
	path2 := Path{}
	path2 = path2.MakeMove(Up)
	path2 = path2.MakeMove(Left)
	path2 = path2.MakeMove(Down)
	path2 = path2.MakeMove(Right)

	joinedPath := path1.Join(path2)

	assert.Equal(joinedPath, Path{
		Up, Down, Left, Left, Up, Right, Down })
}