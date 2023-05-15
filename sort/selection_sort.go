package sort

func SelectionSort(slice []int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		min_value_index := i

		for j := i + 1; j < n; j++ {
			if slice[j] < slice[min_value_index] {
				min_value_index = j
			}
		}

		slice[i], slice[min_value_index] = slice[min_value_index], slice[i]
	}

	return slice
}
