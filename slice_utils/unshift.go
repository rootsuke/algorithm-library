package slice_utils

func Unshift(slice []int, value int) []int {
	if len(slice) == 0 {
		slice = append(slice, value)
	} else if len(slice) == 1 {
		slice[1] = slice[0]
		slice[0] = value
	} else {
		slice, slice[0] = append(slice[:1], slice...), value
	}

	return slice
}
