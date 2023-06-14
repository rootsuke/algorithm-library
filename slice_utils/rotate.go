package slice_utils

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	res := [][]int{}

	for i := 0; i < n; i++ {
		row := make([]int, n)
		res = append(res, row)
	}

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			res[n-y-1][x] = matrix[x][y]
		}
	}

	return res
}
