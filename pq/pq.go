package pq

type EmptyStackError struct {
	msg string
}

func (e *EmptyStackError) Error() string {
    return e.msg
}
// Struct Item, que depois será substituída pela struct State
type Item struct {
	value float64
}

// Fila de prioridade, implementada em uma heap armazenada em array.
type PriorityQueue []Item

// Função de comparação entre 2 itens usada para ordenar a fila de prioridade, ainda não definida.
var compare func(Item,Item)bool

// Aloca e retorna uma nova fila de prioridade, cuja função compare() é recebida de parâmetro.
func New(customCompare func(Item,Item)bool) PriorityQueue {
	pq := new(PriorityQueue)
	compare = customCompare
	return *pq
}
// Retorna a quantidade de elementos na fila.
func (pq *PriorityQueue) Len() int {
    return len(*pq)
}

// Insere uma quantidade qualquer de itens na fila.
func (pq *PriorityQueue) Push(items ...Item) {
    for _, obj := range items {
		*pq = append((*pq), obj)

		for i := pq.Len()-1;
			i != 0 && !(compare((*pq)[pq.parent(i)], (*pq)[i]));
			i = pq.parent(i) {
			(*pq)[pq.parent(i)], (*pq)[i] = (*pq)[i], (*pq)[pq.parent(i)]
		}
    }
}

// Retorna o valor de maior prioridade na fila.
func (pq *PriorityQueue) Top() (result Item, errorHandle error) {
	errorHandle = nil
	if pq.Len() == 0 {
		errorHandle = &EmptyStackError{"Pilha vazia"}
	} else {
		result = (*pq)[0]
	}
	return
}

// Retira o elemento de maior prioridade na fila.
func (pq *PriorityQueue) Pop() (errorHandle error) {
	_, errorHandle = pq.Top()
	if errorHandle != nil { return }

	(*pq)[0], (*pq)[pq.Len()-1] = (*pq)[pq.Len()-1], (*pq)[0]
	*pq = (*pq)[:pq.Len()-1]
	pq.fixHeapQueue(0)
	return
}

func (pq *PriorityQueue) parent(i int) int {
	return (i-1)/2
}

func (pq *PriorityQueue) left(i int) int {
	return 2*i+1
}

func (pq *PriorityQueue) right(i int) int {
	return 2*i+2
}

func (pq *PriorityQueue) fixHeapQueue(i int) {
	l, r := pq.left(i), pq.right(i)
	small := i
	if l < pq.Len() && compare((*pq)[l], (*pq)[i]) {
		small = l		
	}
	if r < pq.Len() && compare((*pq)[r], (*pq)[small]) {
		small = r
	}
	if small != i {
		(*pq)[i], (*pq)[small] = (*pq)[small], (*pq)[i]
		pq.fixHeapQueue(small)
	}
}