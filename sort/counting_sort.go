package sort

var MAX = 10_001

func CountingSort(slice []int) []int {
	C := make([]int, MAX)

	for i := 0; i < len(slice); i++ {
		C[slice[i]]++
	}

	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1]
	}

	B := make([]int, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		B[C[slice[i]]-1] = slice[i]
		C[slice[i]]--
	}

	return B
}
