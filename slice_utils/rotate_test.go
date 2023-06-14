package slice_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate1x1(t *testing.T) {
	testData := [][]int{{1}}

	res := rotate(testData)

	assert.Exactly(t, []int{1}, res[0])
}

func TestRotate2x2(t *testing.T) {
	testData := [][]int{{1, 2}, {3, 4}}

	res := rotate(testData)

	assert.Exactly(t, []int{2, 4}, res[0])
	assert.Exactly(t, []int{1, 3}, res[1])
}

func TestRotate3x3(t *testing.T) {
	testData := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	res := rotate(testData)

	assert.Exactly(t, []int{3, 6, 9}, res[0])
	assert.Exactly(t, []int{2, 5, 8}, res[1])
	assert.Exactly(t, []int{1, 4, 7}, res[2])
}

func TestRotate4x4(t *testing.T) {
	testData := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	res := rotate(testData)

	assert.Exactly(t, []int{4, 8, 12, 16}, res[0])
	assert.Exactly(t, []int{3, 7, 11, 15}, res[1])
	assert.Exactly(t, []int{2, 6, 10, 14}, res[2])
	assert.Exactly(t, []int{1, 5, 9, 13}, res[3])
}

func TestRotate5x5(t *testing.T) {
	testData := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	res := rotate(testData)

	assert.Exactly(t, []int{5, 10, 15, 20, 25}, res[0])
	assert.Exactly(t, []int{4, 9, 14, 19, 24}, res[1])
	assert.Exactly(t, []int{3, 8, 13, 18, 23}, res[2])
	assert.Exactly(t, []int{2, 7, 12, 17, 22}, res[3])
	assert.Exactly(t, []int{1, 6, 11, 16, 21}, res[4])
}
