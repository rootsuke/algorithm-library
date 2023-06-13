package string_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddLengthPalindrome(t *testing.T) {
	testData := "abcba"
	res := IsPalindrome(testData)
	assert.True(t, res)
}

func TestOddLengthNoPalindrome(t *testing.T) {
	testData := "abcda"
	res := IsPalindrome(testData)
	assert.False(t, res)
}

func TestEvenLengthPalindrome(t *testing.T) {
	testData := "abccba"
	res := IsPalindrome(testData)
	assert.True(t, res)
}

func TestEvenLengthNoPalindrome(t *testing.T) {
	testData := "abcdba"
	res := IsPalindrome(testData)
	assert.False(t, res)
}

func TestBlank(t *testing.T) {
	testData := ""
	res := IsPalindrome(testData)
	assert.True(t, res)
}
