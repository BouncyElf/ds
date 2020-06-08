package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleList(t *testing.T) {
	l, expect := NewDoubleList(), make([]interface{}, 0, 20)
	for i := range make([]struct{}, 10) {
		l.Append(i)
		l.Prepend(i * 3)
		expect = append(expect, i)
		expect = append([]interface{}{i * 3}, expect...)
	}
	assert.Equal(t, 20, l.Len())
	assert.Equal(t, expect, l.Slice())
}

func TestDoubleList_Iterator(t *testing.T) {
	l := NewDoubleList()
	for i := range make([]struct{}, 10) {
		l.Append(i)
	}
	for now, i := l.Head(), 0; now != nil; now, i = now.Next(), i+1 {
		assert.Equal(t, i, now.Value.(int))
	}
	for now, i := l.Tail(), 9; now != nil; now, i = now.Prev(), i-1 {
		assert.Equal(t, i, now.Value.(int))
	}
}

func TestDoubleList_Remove(t *testing.T) {
	l := NewDoubleList()
	for range make([]struct{}, 10) {
		l.Remove(l.Head())
		l.Remove(l.Tail())
	}
	l.Remove(new(DoubleListNode))
	assert.Equal(t, 0, l.Len())
	for i := range make([]struct{}, 10) {
		l.Append(i)
	}
	l.Remove(new(DoubleListNode))
	assert.Equal(t, 10, l.Len())
	for range make([]struct{}, 10) {
		l.Remove(l.Head())
		l.Remove(l.Tail())
	}
	l.Remove(new(DoubleListNode))
	assert.Equal(t, 0, l.Len())
	assert.Equal(t, l.head, l.tail.prev)
	assert.Equal(t, l.tail, l.head.next)
}
