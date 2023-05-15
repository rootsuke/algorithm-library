package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	testData := []int{5, 6, 4, 2, 1, 3}
	res := ShellSort(testData)

	expected := []int{1, 2, 3, 4, 5, 6}
	assert.Exactly(t, expected, res)
}
