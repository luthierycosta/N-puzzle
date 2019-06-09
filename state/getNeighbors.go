package state

import (
	// "fmt"
	"github.com/luthierycosta/N-puzzle/path"
)

// GetNeighbors gets the neighbors of a state s,
// that is, which new states can this one generate by making a move in the board.
// Dir represents the last movement made in the board until current s' state
// and its purpose is to prevent generating looping neighbors.
func (s State) GetNeighbors(dir path.Direction) []State {

	// start off without any neighbors
	neighbors := []State{}
	
	// initially, all 4 movement options are available
	up, down 		:= true, true
	left, right := true, true
	
	// cant move back
	if (dir == path.Up) { down = false } else
	if (dir == path.Down) { up = false } else
	if (dir == path.Left) { right = false } else
	if (dir == path.Right) { left = false }

	// if 0 is at the left border of the board, cant move left
	if (s.pos0.X == 0) { left = false }
	// if 0 is at the right border of the board, cant move right
	if (s.pos0.X == len(s.Board)-1) { right = false }
	// if 0 is at the upper border of the board, cant move up
	if (s.pos0.Y == 0) { up = false }
	// if 0 is at the lower border of the board, cant move down
	if (s.pos0.Y == len(s.Board)-1) { down = false }

	// now make the neighbors
	if (up) {
		upMovement := s.makeCopy()										// make copy of current state
		upMovement.move(path.Up)											// move up
		neighbors = append(neighbors, upMovement)			// add to neighbors array
	}
	if (down) {
		downMovement := s.makeCopy()									// make copy of current state
		downMovement.move(path.Down)									// move down
		neighbors = append(neighbors, downMovement)		// add to neighbors array
	}
	if (left) {
		leftMovement := s.makeCopy()									// make copy of current state
		leftMovement.move(path.Left)									// move left
		neighbors = append(neighbors, leftMovement)		// add to neighbors array
	}
	if (right) {
		rightMovement := s.makeCopy()									// make copy of current state
		rightMovement.move(path.Right)								// move right
		neighbors = append(neighbors, rightMovement)	// add to neighbors array
	}

	return neighbors
}

func (s *State) move(dir path.Direction) {
	x0, y0 := s.pos0.X, s.pos0.Y
	if (dir == path.Up) {
		s.Board[y0][x0], s.Board[y0-1][x0] = s.Board[y0-1][x0], s.Board[y0][x0]
		s.pos0.Y = y0-1
	}	else if (dir == path.Down) {
		s.Board[y0][x0], s.Board[y0+1][x0] = s.Board[y0+1][x0], s.Board[y0][x0]
		s.pos0.Y = y0+1
	}	else if (dir == path.Left) {
		s.Board[y0][x0], s.Board[y0][x0-1] = s.Board[y0][x0-1], s.Board[y0][x0]
		s.pos0.X = x0-1
	}	else if (dir == path.Right) {
		s.Board[y0][x0], s.Board[y0][x0+1] = s.Board[y0][x0+1], s.Board[y0][x0]
		s.pos0.X = x0+1
	}
	s.Path = append(s.Path, dir)
}