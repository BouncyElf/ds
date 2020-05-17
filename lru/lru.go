package lru

import "container/list"

type Cache struct {
	capacity int

	l *list.List
	m map[string]*list.Element
}

type entry struct {
	key string
	val interface{}
}

func New(capacity int) *Cache {
	if capacity < 0 {
		capacity = 0
	}
	return &Cache{
		capacity: capacity,

		l: list.New(),
		m: make(map[string]*list.Element, capacity),
	}
}

func (c *Cache) Set(key string, val interface{}) {
	if c == nil {
		return
	}
	c.checkNil()
	if c.Len() >= c.capacity {
		c.RemoveOldest()
	}
	if e, ok := c.m[key]; ok {
		c.l.MoveToFront(e)
		e.Value.(*entry).val = val
	} else {
		e := c.l.PushFront(&entry{
			key: key,
			val: val,
		})
		c.m[key] = e
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if c.checkNil() {
		return nil, false
	}
	if e, ok := c.m[key]; ok {
		c.l.MoveToFront(e)
		return e.Value.(*entry).val, true
	}
	return nil, false
}

func (c *Cache) Remove(key string) (interface{}, bool) {
	if c.checkNil() {
		return nil, false
	}
	if e, ok := c.m[key]; ok {
		v := c.l.Remove(e).(*entry).val
		delete(c.m, key)
		return v, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	if c.checkNil() {
		return
	}
	k := c.l.Remove(c.l.Back()).(*entry).key
	delete(c.m, k)
}

func (c *Cache) Len() int {
	if c == nil {
		return 0
	}
	return len(c.m)
}

func (c *Cache) Clear() {
	if c.checkNil() {
		return
	}
	c.l = list.New()
	c.m = make(map[string]*list.Element, c.capacity)
}

func (c *Cache) checkNil() bool {
	if c == nil {
		return true
	}
	if c.m == nil || c.l == nil {
		c.l = list.New()
		c.m = make(map[string]*list.Element, c.capacity)
		return true
	}
	return false
}
