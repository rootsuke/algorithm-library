package divide_and_conquer

// index以降のsliceの要素を使って整数mを作れる場合はtrueを返す
func ExhaustiveSearch(slice []int, index int, m int) bool {
	if m == 0 {
		return true
	}
	if index >= len(slice) {
		return false
	}

	res := ExhaustiveSearch(slice, index+1, m) || ExhaustiveSearch(slice, index+1, m-slice[index])

	return res
}
