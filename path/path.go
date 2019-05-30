package path

// Direction enum de direcao
type Direction uint8

func (d Direction) invert() Direction {
	if (d < 2) {
		return d + 2
	}
	return d - 2
}

// valores possiveis de Direction
const (
	Up Direction = 0
	Right Direction = 1
	Down Direction = 2
	Left Direction = 3
)

// Path representacao de um caminho
type Path []Direction

// MakeMove funcao de fazer um movimento
func (p *Path) MakeMove(dir Direction) Path {
	return append(*p, dir)
}

func (p Path) Join(pathToJoin Path) Path {
	pLen := len(pathToJoin)
	for left, right := 0, pLen-1; left < right; left, right = left+1, right-1 {
		pathToJoin[left], pathToJoin[right] = pathToJoin[right].invert(), pathToJoin[left].invert()
	}

	return append(p, pathToJoin...)
}