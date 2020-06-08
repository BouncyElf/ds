package us

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionSet_Find(t *testing.T) {
	u := New(10)
	for i := range make([]struct{}, 10) {
		assert.Equal(t, i, u.Find(i))
	}
	assert.Equal(t, -1, u.Find(-1))
	assert.Equal(t, -1, u.Find(10))
}

func TestUnionSet_Union(t *testing.T) {
	u := New(10)
	expect := u.data
	u.Union(10, 10)
	assert.Equal(t, expect, u.data)
	u.Union(10, -1)
	assert.Equal(t, expect, u.data)
	u.Union(5, 3)
	u.Union(1, 5)
	assert.Equal(t, 1, u.Find(1))
	assert.Equal(t, 1, u.Find(3))
	assert.Equal(t, 1, u.Find(5))
}
