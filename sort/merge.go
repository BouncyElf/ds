package sort

func MergeSort(data []int, l, r int) {
	if r-l < 2 {
		return
	}
	mid := (l + r) / 2
	MergeSort(data, mid, r)
	MergeSort(data, l, mid)
	left, right := make([]int, mid-l), make([]int, r-mid)
	copy(left, data[l:mid])
	copy(right, data[mid:r])
	for now, i, j := l, 0, 0; now < r; now++ {
		if i >= len(left) {
			data[now] = right[j]
			j++
		} else if j >= len(right) {
			data[now] = left[i]
			i++
		} else if left[i] < right[j] {
			data[now] = left[i]
			i++
		} else {
			data[now] = right[j]
			j++
		}
	}
}
