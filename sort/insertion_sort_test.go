package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	testData := []int{3, 2, 1}
	res := InsertionSort(testData)

	expected := []int{1, 2, 3}
	assert.Exactly(t, expected, res)
}
