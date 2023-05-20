package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	testData := []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11}

	r := len(testData) - 1
	res := Partition(testData, 0, r)

	assert.Equal(t, 7, res)
}
