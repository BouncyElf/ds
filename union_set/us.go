package us

type UnionSet struct {
	data []int
}

func New(capacity int) *UnionSet {
	us := &UnionSet{
		data: make([]int, capacity),
	}
	for i := range us.data {
		us.data[i] = i
	}
	return us
}

func (us *UnionSet) Find(i int) int {
	if i >= len(us.data) || i < 0 {
		return -1
	}
	for ; us.data[i] != i; i = us.data[i] {
	}
	return i
}

func (us *UnionSet) Union(a, b int) {
	if a >= len(us.data) || b >= len(us.data) || a < 0 || b < 0 {
		return
	}
	if a < b {
		a, b = b, a
	}
	us.data[us.Find(a)] = us.Find(b)
}
