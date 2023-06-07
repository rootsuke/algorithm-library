package string_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOddString(t *testing.T) {
	str := "abcde"

	reversed := reverse(str)

	assert.Equal(t, "edcba", reversed)
}

func TestReverseEvenString(t *testing.T) {
	str := "abcdef"

	reversed := reverse(str)

	assert.Equal(t, "fedcba", reversed)
}

func TestReverseSingleString(t *testing.T) {
	str := "a"

	reversed := reverse(str)

	assert.Equal(t, "a", reversed)
}

func TestReverseTwoString(t *testing.T) {
	str := "ab"

	reversed := reverse(str)

	assert.Equal(t, "ba", reversed)
}

func TestReverseBlankString(t *testing.T) {
	str := ""

	reversed := reverse(str)

	assert.Equal(t, "", reversed)
}
