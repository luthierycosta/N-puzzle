package path

// Direction enum de direcao
type Direction uint8

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