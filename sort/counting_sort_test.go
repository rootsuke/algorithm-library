package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	testData := []int{2, 5, 1, 3, 2, 3, 0}

	res := CountingSort(testData)
	expected := []int{0, 1, 2, 2, 3, 3, 5}
	assert.Exactly(t, expected, res)
}
