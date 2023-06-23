package intsorting

func cpy(sli *[]int) []int {
	// cpy := make([]int, len(sli))
	cpy := append([]int{}, *sli...)
	return cpy
}

func intSwap(sli []int, a int, b int) {
	sli[a], sli[b] = sli[b], sli[a]
}
