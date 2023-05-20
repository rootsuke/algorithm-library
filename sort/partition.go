package sort

func Partition(slice []int, p int, r int) int {
	x := slice[r]
	i := p - 1

	for j := p; j < r; j++ {
		if slice[j] <= x {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[r] = slice[r], slice[i+1]

	return i + 1
}
