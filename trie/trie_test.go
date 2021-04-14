package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newEmpty(t *testing.T) {
	tests := []struct {
		name string
		want *Trie
	}{
		{
			name: "happy path",
			want: &Trie{children: map[byte]*Trie{}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, newEmpty(), tt.name)
	}
}

func TestTrie_Insert(t *testing.T) {
	tests := []struct {
		name string
		word string
	}{
		{
			name: "happy path(empty)",
		},
		{
			name: "happy path(word)",
			word: "this is***123ASLDKNgood",
		},
	}
	for _, tt := range tests {
		assert.NotPanics(t, func() {
			New().Insert(tt.word)
		}, tt.name)
	}

	assert.Panics(t, func() {
		nilTree := (*Trie)(nil)
		nilTree.Insert("")
	})

	assert.Panics(t, func() {
		nilTree := (*Trie)(nil)
		nilTree.Insert("panic")
	})
}

func TestTrie_Search(t *testing.T) {
	tests := []struct {
		name string
		tree *Trie
		word string
		want bool
	}{
		{
			name: "happy path(empty)",
		},
		{
			name: "happy path(found)",
			tree: func() *Trie {
				tree := New()
				tree.Insert("test")
				return tree
			}(),
			word: "test",
			want: true,
		},
		{
			name: "not found",
			tree: newEmpty(),
			word: "test",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.tree.Search(tt.word), tt.name)
	}
}

func TestTrie_HasPrefix(t *testing.T) {
	tests := []struct {
		name string
		tree *Trie
		word string
		want bool
	}{
		{
			name: "happy path(empty)",
		},
		{
			name: "happy path(fully match)",
			tree: func() *Trie {
				tree := New()
				tree.Insert("test")
				return tree
			}(),
			word: "test",
			want: true,
		},
		{
			name: "happy path(prefix match)",
			tree: func() *Trie {
				tree := New()
				tree.Insert("test")
				return tree
			}(),
			word: "tes",
			want: true,
		},
		{
			name: "not found",
			tree: newEmpty(),
			word: "test",
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.tree.HasPrefix(tt.word), tt.name)
	}
}

func TestTrie_getLastNode(t *testing.T) {
	tests := []struct {
		name string
		tree *Trie
		data []byte
		want *Trie
	}{
		{
			name: "happy path(nil)",
		},
		{
			name: "happy path(found)",
			tree: func() *Trie {
				t := New()
				t.Insert("123")
				return t
			}(),
			data: []byte("123"),
			want: func() *Trie {
				t := newEmpty()
				t.IsLeaf = true
				return t
			}(),
		},
		{
			name: "not found",
			tree: func() *Trie {
				t := New()
				t.Insert("123")
				return t
			}(),
			data: []byte("abc"),
			want: nil,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.tree.getLastNode(tt.data), tt.name)
	}
}
