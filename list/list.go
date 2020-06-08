package list

// DoubleList 是双向链表
type DoubleList struct {
	head, tail *DoubleListNode
	length     int
}

type DoubleListNode struct {
	l          *DoubleList
	Value      interface{}
	prev, next *DoubleListNode
}

func (n *DoubleListNode) Next() *DoubleListNode {
	if n == nil || n.l == nil || n.next == n.l.head || n.next == n.l.tail {
		return nil
	}
	return n.next
}

func (n *DoubleListNode) Prev() *DoubleListNode {
	if n == nil || n.l == nil || n.prev == n.l.head || n.prev == n.l.tail {
		return nil
	}
	return n.prev
}

func NewDoubleList() *DoubleList {
	head, tail := new(DoubleListNode), new(DoubleListNode)
	head.next = tail
	tail.prev = head
	return &DoubleList{
		head: head,
		tail: tail,
	}
}

func (d *DoubleList) Append(val interface{}) {
	node := &DoubleListNode{
		Value: val,
		prev:  d.tail.prev,
		next:  d.tail,
		l:     d,
	}
	d.tail.prev.next = node
	d.tail.prev = node
	d.length++
}

func (d *DoubleList) Prepend(val interface{}) {
	node := &DoubleListNode{
		Value: val,
		prev:  d.head,
		next:  d.head.next,
		l:     d,
	}
	d.head.next.prev = node
	d.head.next = node
	d.length++
}

func (d *DoubleList) Head() *DoubleListNode {
	if d.length == 0 {
		return nil
	}
	return d.head.next
}

func (d *DoubleList) Tail() *DoubleListNode {
	if d.length == 0 {
		return nil
	}
	return d.tail.prev
}

func (d *DoubleList) Remove(node *DoubleListNode) interface{} {
	if node == nil {
		return nil
	}
	if node == d.head || node == d.tail || node.l != d {
		return node.Value
	}
	for now := d.head.next; now != d.tail; now = now.next {
		if now == node {
			now.prev.next = now.next
			now.next.prev = now.prev
			d.length--
			break
		}
	}
	return node.Value
}

func (d *DoubleList) Slice() []interface{} {
	res := make([]interface{}, 0, d.length)
	for now := d.head.next; now != d.tail; now = now.next {
		res = append(res, now.Value)
	}
	return res
}

func (d *DoubleList) Len() int {
	return d.length
}
