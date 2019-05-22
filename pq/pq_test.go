package pq

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// sem usar a bliblioteca de asser
func TestPush(t *testing.T) {
	fila := PriorityQueue{}
	item := Item{2.4}
	fila.Push(item)
	if fila.Len() != 1 {
		t.Errorf("Deu erro no pq.Push()")
	}
}

// usando a biblioteca de assert
func TestPushMultiplos(t *testing.T) {
	assert := assert.New(t)

	fila := PriorityQueue{}
	fila.Push(Item{2.5})
	fila.Push(Item{2.6})
	fila.Push(Item{2.7})
	fila.Push(Item{2.8})

	assert.Equal(fila.Len(), 4)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)

	fila := PriorityQueue{}
	

	fila.Push(Item{0.3},Item{0.2}, Item{0.45},Item{0.7}, Item{0.4})

	a := fila.Pop()
	b := fila.Pop()
	fila.Push(Item{0.1}, Item{0.6})
	c := fila.Pop()

	assert.Equal(a, Item{0.2})
	assert.Equal(b, Item{0.3})
	assert.Equal(c, Item{0.1})
}

func TestPopOnEmpty(t *testing.T) {
	assert := assert.New(t)

	fila := PriorityQueue{}
	
	assert.Equal(fila.Pop(), Item{})
}