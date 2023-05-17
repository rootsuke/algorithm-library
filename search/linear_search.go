package search

func LinearSearch(slice []int, target int) int {
	n := len(slice)

	// 番兵
	slice = append(slice, target)

	i := 0

	for slice[i] != target {
		i++
	}

	if i == n {
		return -1
	} else {
		return i
	}
}
