package divide_and_conquer

func FindMax(slice []int, leftIndex int, rightIndex int) int {
	if leftIndex == rightIndex-1 {
		return slice[leftIndex]
	}

	middleIndex := (leftIndex + rightIndex) / 2

	leftMax := FindMax(slice, leftIndex, middleIndex)
	rightMax := FindMax(slice, middleIndex, rightIndex)

	if leftMax > rightMax {
		return leftMax
	} else {
		return rightMax
	}
}
