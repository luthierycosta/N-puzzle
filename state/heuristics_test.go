package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	assert := assert.New(t)

	s1 := New([][]int{
		{2, 8, 3},
		{1, 6, 4},
		{7, 0, 5},
	})
	s2 := New([][]int{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	})
	s3 := s1.makeCopy()

	assert.Equal(s1.DistanceTo(s3), 0, "Se os estados são iguais, a distância tem que ser 0")
	assert.Equal(s1.DistanceTo(s2), 5, "Exemplo do Li lá de IIA")
	assert.Equal(s1.DistanceTo(s2), s2.DistanceTo(s1), "A distância é comutativa")
}