package sort

func QuickSort(slice []int, p int, r int) {
	// 要素数が1になったら終了
	if p+1 >= r {
		return
	}

	res := Partition(slice, p, r)
	QuickSort(slice, p, res-1)
	QuickSort(slice, res+1, r)
}
