package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTargetByBinarySearch(t *testing.T) {
	testData := []int{0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 5, 6, 6, 6, 8, 8, 9, 9, 9}
	target := 6

	idx := BinarySearch(testData, target)

	assert.Equal(t, 13, idx)
}

func TestTargetNotFoundByBinarySearch(t *testing.T) {
	testData := []int{0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 5, 6, 6, 6, 8, 8, 9, 9, 9}
	target := 4

	idx := BinarySearch(testData, target)

	assert.Equal(t, -1, idx)
}
