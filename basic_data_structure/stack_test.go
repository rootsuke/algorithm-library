package basic_data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 1, stack.stack[0])
	assert.Equal(t, 2, stack.stack[1])

	a := stack.Pop()
	b := stack.Pop()

	// 後入れ先出しになっていることを確認
	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)
}
