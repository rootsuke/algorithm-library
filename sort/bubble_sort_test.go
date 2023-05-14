package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	testData := []int{5, 3, 2, 4, 1}
	res := BubbleSort(testData)

	expected := []int{1, 2, 3, 4, 5}
	assert.Exactly(t, expected, res)
}

func TestBubbleSort2(t *testing.T) {
	testData := []int{5, 3, 2, 4, 1}
	res := BubbleSort2(testData)

	expected := []int{1, 2, 3, 4, 5}
	assert.Exactly(t, expected, res)
}
