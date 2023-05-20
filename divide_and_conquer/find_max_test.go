package divide_and_conquer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMax(t *testing.T) {
	testData := []int{5, 6, 4, 2, 1, 3}
	leftIndex := 2
	rightIndex := len(testData)
	max := FindMax(testData, leftIndex, rightIndex)

	assert.Equal(t, 4, max)
}
