package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	testData := []int{8, 5, 9, 2, 6, 3, 7, 1, 10, 4}

	left := 0
	right := len(testData)
	MergeSort(testData, left, right)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Exactly(t, expected, testData)
}
