package pq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPQ(t *testing.T) {
	pq := New()
	data := map[int]interface{}{
		10:  "test10",
		111: "test111",
		0:   "test0",
	}
	for i := range data {
		pq.Push(data[i], i)
	}
	assert.Len(t, *pq.m, len(data))
	for _, i := range []int{111, 10, 0} {
		v, score := pq.Pop()
		assert.Equal(t, data[i], v)
		assert.Equal(t, i, score)
	}
}
