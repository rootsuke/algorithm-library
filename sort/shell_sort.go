package sort

func ShellSort(slice []int) []int {
	n := len(slice)
	intervals := calcInterval(n)

	for i := len(intervals) - 1; i >= 0; i-- {
		insertionSort(slice, intervals[i])
	}

	return slice
}

func calcInterval(n int) []int {
	var h = []int{}

	for i := 1; i < n; i *= 3 + 1 {
		h = append(h, i)
	}

	return h
}

func insertionSort(slice []int, interval int) {
	n := len(slice)

	for i := interval; i < n; i++ {
		crr := slice[i]
		j := i - interval

		for j >= 0 && crr < slice[j] {
			slice[j+interval] = slice[j]
			j -= interval
		}

		slice[j+interval] = crr
	}
}
