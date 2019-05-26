package pq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func setup() (result PriorityQueue) {
	result = PriorityQueue{}
	result.compare = func (a, b Item) bool {
		return a.compare(&b);
	};
	return;
}

func TestPush(t *testing.T) {
	queue := setup();
	item := Item{2.4};
	queue.push(item);
	assert.Equal(t, queue.len(), 1, "O número de elementos na fila deve ser 1");
}

func TestTopEmpty(t *testing.T) {
	queue := setup();
	_, errorHandle := queue.top();
	assert.EqualErrorf(t, errorHandle, "Pilha vázia", "Deve retornar uma mensagem de erro %s ao tentar desempilhar de uma fila vázia!", "Pilha vázia");
}

func TestTop(t *testing.T) {
	queue := setup();
	item := Item{2.4};
	queue.push(item);
	top, errorHandle := queue.top();
	assert.Equal(t, top, item, "O top deve retornar o elemento empilhado!");
	assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado!");
}

func TestTopMultiple(t *testing.T) {
	queue := setup();
	items := []Item{Item{2.4}, Item{3.4}, Item{0.1}, Item{6.4}, Item{7.4}, Item{10.4}};
	order := []int{2, 0, 1, 3, 4, 5};
	queue.push(items...);
	assert.Equal(t, queue.len(), len(items), "O número de elementos na fila deve ser %d", len(items));

	for i := 0; i < len(items); i++ {
		top, errorHandle := queue.top();
		assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado ao recuperar o %dº elemento!", i);
		assert.Equal(t, top, items[order[i]], "O %dº elemento da fila deve ser Item{%d}", i, items[order[i]]);
		errorHandle = queue.pop();
		assert.Equal(t, errorHandle, nil, "Nenhum erro deve ser retornado ao remover o %dº elemento!", i);
	}
	_, errorHandle := queue.top();
	assert.EqualErrorf(t, errorHandle, "Pilha vázia", "A pilha deve estar vázia após desempilhar todos os elementos!");
}