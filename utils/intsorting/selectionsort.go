package intsorting

func SelectionSort(slice []int) []int {
	sli := cpy(&slice)
	min := -1

	for i := 0; i < len(sli)-1; i++ {
		min = i

		for j := i + 1; j < len(sli); j++ {
			if sli[j] < sli[min] {
				min = j
			}
		}

		if i != min {
			intSwap(sli, i, min)
		}
	}

	return sli
}
