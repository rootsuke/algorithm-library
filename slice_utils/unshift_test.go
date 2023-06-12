package slice_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnshift(t *testing.T) {
	slice := []int{1, 2, 3}
	value := 0

	slice = Unshift(slice, value)

	expected := []int{0, 1, 2, 3}
	assert.Exactly(t, expected, slice)
}
