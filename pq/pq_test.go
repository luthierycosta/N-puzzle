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