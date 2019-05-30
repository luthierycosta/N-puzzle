package pq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func setup() PriorityQueue {
	pq := New(func(a,b Item)bool { return a.value <= b.value})	
	return pq
}

func TestPush(t *testing.T) {
	queue := setup()
	item := Item{2.4}
	queue.Push(item)
	assert.Equal(t, queue.Len(), 1, "O número de elementos na fila deve ser 1")
}

func TestTopEmpty(t *testing.T) {
	queue := setup()
	_, errorHandle := queue.Top()
	assert.EqualErrorf(t, errorHandle, "Pilha vazia", "Deve retornar uma mensagem de erro %s ao tentar desempilhar de uma fila vázia!", "Pilha vázia")
}

func TestTop(t *testing.T) {
	queue := setup()
	item := Item{2.4}
	queue.Push(item)
	top, errorHandle := queue.Top()
	assert.Equal(t, top, item, "O top deve retornar o elemento empilhado!")
	assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado!")
}

func TestTopMultiple(t *testing.T) {
	queue := setup()
	items := []Item{Item{2.4}, Item{3.4}, Item{0.1}, Item{6.4}, Item{4.7}, Item{7.4}, Item{10.4}}
	order := []int{2, 0, 1, 4, 3, 5, 6}
	queue.Push(items...)
	assert.Equal(t, queue.Len(), len(items), "O número de elementos na fila deve ser %d", len(items))

	for i := 0; i < len(items); i++ {
		Top, errorHandle := queue.Top()
		assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado ao recuperar o %dº elemento!", i)
		assert.Equal(t, Top, items[order[i]], "O %dº elemento da fila deve ser Item{%d}", i, items[order[i]])
		errorHandle = queue.Pop()
		assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado ao remover o %dº elemento!", i)
	}
	_, errorHandle := queue.Top()
	assert.EqualErrorf(t, errorHandle, "Pilha vazia", "A pilha deve estar vazia após desempilhar todos os elementos!")
}