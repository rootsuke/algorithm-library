package basic_data_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	dll := NewLinkedList()

	dll.Insert(5)
	assert.Equal(t, 5, dll.head.next.key)
	current := dll.head.next
	assert.Equal(t, dll.head, current.next)
	assert.Equal(t, 5, dll.head.prev.key)

	dll.Insert(2)
	assert.Equal(t, 2, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 5, current.next.key)
	assert.Equal(t, 2, current.next.prev.key)

	dll.Insert(3)
	assert.Equal(t, 3, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 2, current.next.key)
	assert.Equal(t, 3, current.next.prev.key)

	dll.Insert(1)
	assert.Equal(t, 1, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 3, current.next.key)
	assert.Equal(t, 1, current.next.prev.key)

	dll.DeleteKey(3)
	assert.Equal(t, 1, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 2, current.next.key)
	assert.Equal(t, 1, current.next.prev.key)

	dll.Insert(6)
	assert.Equal(t, 6, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 1, current.next.key)
	assert.Equal(t, 6, current.next.prev.key)

	dll.DeleteKey(5)
	assert.Equal(t, 2, dll.head.prev.key)

	dll.DeleteFirst()
	assert.Equal(t, 1, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, 2, current.next.key)
	assert.Equal(t, 1, current.next.prev.key)

	dll.DeleteLast()
	assert.Equal(t, 1, dll.head.next.key)
	current = dll.head.next
	assert.Equal(t, dll.head, current.next)
	assert.Equal(t, 1, current.next.prev.key)
}
