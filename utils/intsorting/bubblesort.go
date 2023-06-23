package intsorting

func BubbleSort(slice []int) []int {
	sli := cpy(&slice)
	swapped := false

	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-1-i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return sli
}
