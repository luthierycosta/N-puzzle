package pq

import "testing"

func TestPush(t *testing.T) {
	fila := PriorityQueue{}
	item := Item{2.4}
	fila.Push(item)
	if fila.Len() != 1 {
		t.Errorf("Deu erro no pq.Push()")
	}
}