package pq

import "container/heap"

type PriorityQueue struct {
	m *minHeap
}

func New() *PriorityQueue {
	m := minHeap(make([]*entry, 0, 100))
	return &PriorityQueue{
		m: &m,
	}
}

func (pq *PriorityQueue) Push(val interface{}, score int) {
	heap.Push(pq.m, &entry{
		val:   val,
		score: score,
	})
}

func (pq *PriorityQueue) Top() (val interface{}, score int) {
	if pq.IsEmpty() {
		return nil, 0
	}
	top := (*pq.m)[0]
	return top.val, top.score
}

func (pq *PriorityQueue) Pop() (val interface{}, score int) {
	v := heap.Pop(pq.m).(*entry)
	return v.val, v.score
}

func (pq *PriorityQueue) Len() int {
	return pq.m.Len()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

type entry struct {
	score int
	val   interface{}
}

type minHeap []*entry

func (m minHeap) Less(i, j int) bool { return m[i].score > m[j].score }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m minHeap) Len() int           { return len(m) }
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(*entry))
}
func (m *minHeap) Pop() interface{} {
	if m.Len() == 0 {
		return (*entry)(nil)
	}
	last := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return last
}
