package estado

type Pair struct {
	X int
	Y int
}

type State struct {
	gameBoard [][]int
	pos0      Pair
}

func NewStateFinal(n int) State {
	state := State{}
	state.gameBoard = make([][]int, n)
	for i := 0; i < n; i++ {
		state.gameBoard[i] = make([]int, n)
		for j := 0; j < n; j++ {
			state.gameBoard[i][j] = i*n + j + 1
			if j == n-1 && i == n-1 {
				state.gameBoard[i][j] = 0
			}
		}
	}
	state.pos0 = Pair{n - 1, n - 1}

	return state
}
