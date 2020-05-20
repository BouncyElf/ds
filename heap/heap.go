package heap

type Heap []int

const (
	defaultCap = 100
)

func down(data []int, i int) {
	for now := i; now*2+1 < len(data); {
		flag := now*2 + 1
		if now*2+2 < len(data) && data[now*2+2] < data[now*2+1] {
			flag = now*2 + 2
		}
		if data[flag] < data[now] {
			data[now], data[flag] = data[flag], data[now]
			now = flag
		} else {
			break
		}
	}
}

func up(data []int, i int) {
	for now := i; (now-1)/2 >= 0 && data[(now-1)/2] > data[now]; now = (now - 1) / 2 {
		data[now], data[(now-1)/2] = data[(now-1)/2], data[now]
	}
}

func New(datas ...[]int) *Heap {
	length := 0
	for i := range datas {
		length += len(datas[i])
	}
	capacity := defaultCap
	if length != 0 {
		capacity = length
	}
	res := make([]int, 0, capacity)
	for i := range datas {
		res = append(res, datas[i]...)
	}
	h := (*Heap)(&res)
	h.init()
	return h
}

func (h *Heap) init() {
	for now := h.Len()/2 - 1; now >= 0; now-- {
		down(*h, now)
	}
}

func (h *Heap) Push(val int) {
	*h = append(*h, val)
	up(*h, h.Len()-1)
}

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) Pop() int {
	if h.Len() == 0 {
		return 0
	}
	(*h)[0], (*h)[h.Len()-1] = (*h)[h.Len()-1], (*h)[0]
	res := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	down(*h, 0)
	return res
}

func (h Heap) Top() int {
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}
