package state

// Calcula a função heurística de um estado a, que será usada na busca.
// Pega a soma da distâncias das posições dos blocos entre um estado e outro.
func (a *State) DistanceTo(b *State) (res int) {
    n := len(a.gameBoard)
    // for every block i in the puzzle, compare its position on both boards and take the difference
    for i := 0; i < n*n-1; i++ {
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

func difference(a, b Pair) (res int) {
    // below code is equivalent to: res = math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
    // Abs only supports floats
    xdiff, ydiff := a.X-b.X, a.Y-b.Y
    if xdiff < 0 { res += -xdiff } else { res += xdiff }
    if ydiff < 0 { res += -ydiff } else { res += ydiff }
    return
}