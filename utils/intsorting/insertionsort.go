package intsorting

func InsertionSort(slice []int) []int {
	sli := cpy(&slice)

	var tmp int
	for i := 1; i < len(sli); i++ {
		tmp = sli[i]
		j := i - 1

		for ; j >= 0 && sli[j] > tmp; j-- {
			sli[j+1] = sli[j]
		}

		sli[j+1] = tmp
	}

	return sli
}
