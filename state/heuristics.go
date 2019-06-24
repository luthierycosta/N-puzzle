package state

// Heuristic calcula a função heurística de um estado a, que será usada na busca até um estado alvo.
func (a State) Heuristic(target State) (int,int) {
	return len(a.Path), a.DistanceTo(target)
}

// DistanceTo a -> b é o somatório das diferenças das posições dos blocos entre a e b.
func (a State) DistanceTo(b State) (res int) {
	n := len(a.Board)
	// para cada bloco i no tabuleiro, obtenha sua posição nos dois estados e calcule a diferença entre eles
	for i := 1; i < n*n; i++ {
		res += difference(a.findPos(i), b.findPos(i))
	}
	return
}

// Calcula quantos passos distante a casa a=(xa,ya) está de b=(xb, yb)
// em outras palavras, a distância de Manhattan entre os dois pontos.
func difference(a, b Pair) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

// Função módulo, pois o math.Abs() só suporta floats.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}