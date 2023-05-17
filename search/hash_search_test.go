package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAndSearch(t *testing.T) {
	hashTable := [11]int{}

	key := 14
	insertRes := Insert(hashTable[:], key)

	assert.Equal(t, 3, insertRes)

	searchRes := Search(hashTable[:], key)
	assert.Equal(t, 3, searchRes)
	assert.Equal(t, 14, hashTable[searchRes])
}

func TestHashCollision(t *testing.T) {
	hashTable := [11]int{}

	key := 25
	res := Insert(hashTable[:], key)
	assert.Equal(t, 3, res)

	collisionKey := 14
	collisionRes := Insert(hashTable[:], collisionKey)
	assert.Equal(t, 8, collisionRes)

	collisionSearchRes := Search(hashTable[:], collisionKey)
	assert.Equal(t, 8, collisionSearchRes)
	assert.Equal(t, 14, hashTable[collisionSearchRes])
}

func TestKeyNotFound(t *testing.T) {
	hashTable := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	key := 12

	res := Search(hashTable[:], key)

	assert.Equal(t, -1, res)
}

func TestHashTableIsFull(t *testing.T) {
	hashTable := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	key := 12

	res := Insert(hashTable[:], key)

	assert.Equal(t, -1, res)
}
