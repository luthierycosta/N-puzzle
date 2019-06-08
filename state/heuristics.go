package state

// DistanceTo calcula a função heurística de um estado a, que será usada na busca até um estado b.
// Pega a soma da distâncias das posições dos blocos entre um estado e outro.
func (a State) DistanceTo(b State) (res int) {
    n := len(a.gameBoard)
    // para cada bloco i no tabuleiro, obtenha sua posição nos dois estados e calcule a diferença entre eles
    for i := 1; i < n*n; i++ {
        res += difference(a.findPos(i), b.findPos(i))
    }
    return
}

// Calcula quantos passos distante a casa (xa,ya) está de (xb, yb).
func difference(a, b Pair) int {
    return abs(a.X-b.X) + abs(a.Y-b.Y)
}

// Função módulo, pois o math.Abs() só suporta floats.
func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}