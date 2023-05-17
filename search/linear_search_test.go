package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTargetByLinearSearch(t *testing.T) {
	testData := []int{0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 5, 6, 6, 6, 8, 8, 9, 9, 9}
	target := 6

	idx := LinearSearch(testData, target)

	assert.Equal(t, 12, idx)
}

func TestTargetNotFoundByLinearSearch(t *testing.T) {
	testData := []int{0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 5, 6, 6, 6, 8, 8, 9, 9, 9}
	target := 4

	idx := LinearSearch(testData, target)

	assert.Equal(t, -1, idx)
}
