package state

// Calcula a função heurística de um estado a, que será usada na busca.
// Pega a soma da distâncias das posições dos blocos entre um estado e outro.
func (a *State) DistanceTo(b *State) (res int) {
    n := len(a.gameBoard)
    // for every block i in the puzzle, compare its position on both boards and take the difference
    for i := 1; i < n*n; i++ {
        res += difference(a.findPos(i), b.findPos(i))
    }
    return
}

func (s *State) findPos(k int) Pair {
    for i := range s.gameBoard {
        for j := range s.gameBoard[i] {
            if s.gameBoard[i][j] == k {
                return Pair{i, j}
            }
        }
    }
    return Pair{-1,-1}      // throw some error
}

func difference(a, b Pair) int {
    return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}