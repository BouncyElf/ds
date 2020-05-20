package heap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	data := []int{1, 5, 9, 9, 8, 2, 0, 3}
	h := New(data)
	assert.Equal(t, 0, h.Top())
	assertHeapEqualSlice(t, h, data)

	data = []int{9, 2, 0, 1, 99, 2, 3, 5, 7, 0}
	h = New()
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

func TestPanic(t *testing.T) {
	assert.NotPanics(t, func() { down(nil, 10) })
	assert.NotPanics(t, func() { down(nil, 0) })
	assert.NotPanics(t, func() { down(nil, -5) })
	assert.NotPanics(t, func() { down([]int{1, 2, 3}, 10) })
	assert.NotPanics(t, func() { down([]int{1, 2, 3}, 0) })
	assert.NotPanics(t, func() { down([]int{1, 2, 3}, -5) })

	assert.NotPanics(t, func() { up(nil, 0) })
	assert.NotPanics(t, func() { up(nil, -5) })
	assert.NotPanics(t, func() { up(nil, 10) })
	assert.NotPanics(t, func() { up([]int{1, 2, 3}, 10) })
	assert.NotPanics(t, func() { up([]int{1, 2, 3}, -5) })
	assert.NotPanics(t, func() { up([]int{1, 2, 3}, 0) })

	assert.NotPanics(t, func() { getNilHeap().Top() })
	assert.NotPanics(t, func() { getNilHeap().Push(0) })
	assert.NotPanics(t, func() { getNilHeap().Pop() })
	assert.NotPanics(t, func() { getNilHeap().Len() })
	assert.NotPanics(t, func() { getNilHeap().init() })
}

func getNilHeap() *Heap {
	var data []int
	data = nil
	return (*Heap)(&data)
}
