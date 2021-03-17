package algorithm

func qsort(values []int, l, r int) {
	if r - l < 2 {
		return
	}
	
	x, i, j := values[l], l, r - 1
	for i < j {
		for i < j && values[j] >= x {
			j--
		}
		if i < j {
			values[i] = values[j]
			i++
		}
		for i < j && values[i] <= x {
			i++
		}
		if i < j {
			values[j] = values[i]
			j--
		}
	}
	values[i] = x
	qsort(values, l, i-1)
	qsort(values, i+1, r)
} 


func QuickSort(values []int) {
	qsort(values, 0, len(values))
}