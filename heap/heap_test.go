package heap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	data := []int{1, 5, 9, 9, 8, 2, 0, 3}
	h := New(data)
	assertHeapEqualSlice(t, h, data)

	data = []int{9, 2, 0, 1, 99, 2, 3, 5, 7, 0}
	for _, v := range data {
		h.Push(v)
	}
	assertHeapEqualSlice(t, h, data)
}

func assertHeapEqualSlice(t *testing.T, h *Heap, data []int) {
	assert.Len(t, *h, len(data))
	sort.Ints(data)
	for _, v := range data {
		assert.Equal(t, v, h.Pop())
	}
}
