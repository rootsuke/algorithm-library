package divide_and_conquer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExhaustiveSearch(t *testing.T) {
	testData := []int{1, 5, 7, 10, 21}
	var res bool

	res = ExhaustiveSearch(testData, 0, 2)
	assert.False(t, res)

	res = ExhaustiveSearch(testData, 0, 17)
	assert.True(t, res)

	res = ExhaustiveSearch(testData, 0, -1)
	assert.False(t, res)
}

func TestExhaustiveSearchIncludeMinus(t *testing.T) {
	testData := []int{1, -5, 7, -10, 21}
	var res bool

	res = ExhaustiveSearch(testData, 0, -1)
	assert.False(t, res)

	res = ExhaustiveSearch(testData, 0, 17)
	assert.True(t, res)

	res = ExhaustiveSearch(testData, 0, -14)
	assert.True(t, res)
}
