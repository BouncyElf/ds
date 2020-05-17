package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	type expect struct {
		val   interface{}
		exist bool
	}
	cases := []struct {
		name     string
		capacity int
		entries  []*entry
		getKeys  []string
		expects  []expect
	}{
		{
			name: "happy case(no test case)",
		},
		{
			name:     "happy case(within capacity)",
			capacity: 3,
			entries: []*entry{
				&entry{
					key: "k1",
					val: "v1",
				},
			},
			getKeys: []string{"k1"},
			expects: []expect{
				expect{
					val:   "v1",
					exist: true,
				},
			},
		},
		{
			name:     "happy case(beyond capacity)",
			capacity: 1,
			entries: []*entry{
				&entry{
					key: "k1",
					val: "v1",
				},
				&entry{
					key: "k2",
					val: "v2",
				},
			},
			getKeys: []string{"k1", "k2"},
			expects: []expect{
				expect{
					val:   nil,
					exist: false,
				},
				expect{
					val:   "v2",
					exist: true,
				},
			},
		},
	}
	for _, c := range cases {
		cache := New(c.capacity)
		for _, e := range c.entries {
			cache.Set(e.key, e.val)
		}
		assert.Equal(t, len(c.getKeys), len(c.expects),
			"expects's length should be same with getKeys")
		for i, key := range c.getKeys {
			res, ok := cache.Get(key)
			assert.Equal(t, c.expects[i].val, res)
			assert.Equal(t, c.expects[i].exist, ok)
		}
	}
}
