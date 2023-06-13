package string_utils

func IsPalindrome(str string) bool {
	n := len(str)

	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}

	return true
}
