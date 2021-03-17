package algorithm

func bubblesort(values []int) {
	for i := len(values)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}


func BubbleSort(values []int) {
	bubblesort(values)
}