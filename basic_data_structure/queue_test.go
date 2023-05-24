package basic_data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := Queue{}
	assert.True(t, queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)

	assert.False(t, queue.IsEmpty())
	assert.Equal(t, 1, queue.queue[0])
	assert.Equal(t, 2, queue.queue[1])

	a := queue.Dequeue()
	b := queue.Dequeue()
	assert.True(t, queue.IsEmpty())
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}
