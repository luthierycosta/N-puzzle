package pq

// No futuro substituiremos Item pela struct State importada de fora
type Item struct {
	F float64
}

func (i *Item) Priority() float64 {return i.F}

// Fila de prioridade que ordena os itens de forma crescente (menor valor primeiro).
type PriorityQueue []Item

// Aloca e retorna uma nova fila de prioridade com os itens especificados (opcionais).
func New(elements ...Item) *PriorityQueue {
    pq := new(PriorityQueue)
    for i := range elements {
        pq.Push(elements[i])
    }
    return pq
}

// Retorna o tamanho atual da fila.
func (pq *PriorityQueue) Len() int {
    return len(*pq)
}

// Insere um elemento na fila de prioridade.
func (pq *PriorityQueue) Push(value Item) {
    // insere-o no fim da fila
    *pq = append(*pq, value)
    // percorre a fila pra posicionar o elemento na ordem certa
    for i := len(*pq)-1; i > 0 && value.Priority() < (*pq)[i-1].Priority(); i-- {
        (*pq)[i-1], (*pq)[i] = (*pq)[i], (*pq)[i-1]
    }
}

// Retira o elemento com maior prioridade na fila
func (pq *PriorityQueue) Pop() (result Item) {
    result = (*pq)[0]
    *pq = (*pq)[1:]
    return
}