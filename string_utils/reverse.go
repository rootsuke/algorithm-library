package string_utils

func reverse(str string) string {
	n := len(str)
	rs := []rune(str)
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-i-1] = rs[n-i-1], rs[i]
	}

	return string(rs)
}
