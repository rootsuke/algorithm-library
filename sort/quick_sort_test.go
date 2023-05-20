package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	testData := []int{8, 5, 9, 2, 6, 3, 7, 1, 10, 4}

	QuickSort(testData, 0, len(testData)-1)
}
