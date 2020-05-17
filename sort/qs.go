package sort

func QuickSort(data []int, l, r int) {
	if l >= r {
		return
	}
	temp, i, j := data[l], l, r
	for i < j {
		for i < j && data[j] >= temp {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
			i++
		}
		for i < j && data[i] <= temp {
			i++
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
			j--
		}
	}
	QuickSort(data, i+1, r)
	QuickSort(data, l, i-1)
}
