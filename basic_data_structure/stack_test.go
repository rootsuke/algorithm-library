package basic_data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchTargetByBinarySearch(t *testing.T) {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)

	b := stack.Pop()
	a := stack.Pop()

	// 後入れ先出しになっていることを確認
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}
