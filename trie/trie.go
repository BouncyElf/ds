package trie

type Trie struct {
	IsLeaf   bool
	children map[byte]*Trie
}

func New() *Trie {
	return newEmpty()
}

func newEmpty() *Trie {
	return &Trie{
		children: make(map[byte]*Trie),
	}
}

func (t *Trie) Insert(word string) {
	if t == nil {
		panic("trie: trying to insert value to a nil tree")
	}
	now := t
	for i := 0; i < len(word); i++ {
		if now.children[word[i]] == nil {
			now.children[word[i]] = newEmpty()
		}
		now = now.children[word[i]]
	}
	now.IsLeaf = true
}

func (t *Trie) Search(word string) bool {
	last := t.getLastNode([]byte(word))
	return last != nil && last.IsLeaf
}

func (t *Trie) HasPrefix(prefix string) bool {
	return t.getLastNode([]byte(prefix)) != nil
}

func (t *Trie) getLastNode(data []byte) *Trie {
	if t == nil {
		return nil
	}
	now := t
	for i := 0; i < len(data); i++ {
		if now.children[data[i]] == nil {
			return nil
		}
		now = now.children[data[i]]
	}
	return now
}
