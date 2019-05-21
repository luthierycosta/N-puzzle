package pq

// No futuro substituiremos Item pela struct State importada de fora
type Item struct {
	F float64
}

func (i *Item) Priority() float64 {return i.F}

// Fila de prioridade que ordena os itens de forma crescente (menor valor primeiro).
type PriorityQueue []Item

// Retorna o tamanho atual da fila.
func (pq *PriorityQueue) Len() int {
    return len(*pq)
}

// Insere um elemento na fila de prioridade.
func (pq *PriorityQueue) Push(values ...Item) {
    for _, value := range values {
        // insere-o no fim da fila
        *pq = append(*pq, value)
        // percorre a fila pra posicionar o elemento na ordem certa
        for i := len(*pq)-1; i > 0 && value.Priority() < (*pq)[i-1].Priority(); i-- {
            (*pq)[i-1], (*pq)[i] = (*pq)[i], (*pq)[i-1]
        }
    }
}

// Retira o elemento com maior prioridade na fila
func (pq *PriorityQueue) Pop() (result Item) {
    if pq.Len() == 0 {
        result = Item{}                 // não tem como retornar nil
    } else {
        result = (*pq)[0]
        *pq = (*pq)[1:]
    }
    return
}