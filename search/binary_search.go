package search

func BinarySearch(slice []int, target int) int {
	n := len(slice)
	left := 0
	right := n

	for left < right {
		mid := (right + left) / 2

		if slice[mid] == target {
			return mid
		} else if target < slice[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}
