package sort

import (
	"fmt"
	"math"
)

func MergeSort(slice []int, l int, r int) {
	fmt.Printf("slice[l:r] %v : %p \n", slice[l:r], slice)
	// 要素数が1になったら分割終了
	if l+1 >= r {
		return
	}

	mid := (l + r) / 2
	MergeSort(slice, l, mid)
	MergeSort(slice, mid, r)
	merge(slice, l, mid, r)
}

func merge(slice []int, l int, mid int, r int) []int {
	fmt.Println("merge start!")
	fmt.Printf("slice: %v : %p \n", slice, slice)
	tmpLSlice := []int{}
	tmpRSlice := []int{}
	fmt.Printf("tmpLSlice0: %v len=%v cap=%v (%p) \n", tmpLSlice, len(slice), cap(slice), &tmpLSlice)
	fmt.Printf("tmpRSlice0: %v len=%v cap=%v (%p) \n", tmpRSlice, len(slice), cap(slice), &tmpRSlice)

	tmpLSlice = append(tmpLSlice, slice[l:mid]...)
	tmpRSlice = append(tmpRSlice, slice[mid:r]...)
	fmt.Printf("tmpLSlice1: %v len=%v cap=%v (%p) \n", tmpLSlice, len(slice), cap(slice), &tmpLSlice)
	fmt.Printf("tmpRSlice1: %v len=%v cap=%v (%p) \n", tmpRSlice, len(slice), cap(slice), &tmpRSlice)

	tmpLSlice = append(tmpLSlice, math.MaxInt64)
	tmpRSlice = append(tmpRSlice, math.MaxInt64)

	leftIndex := 0
	rightIndex := 0

	fmt.Printf("tmpLSlice2: %v len=%v cap=%v (%p) \n", tmpLSlice, len(slice), cap(slice), &tmpLSlice)
	fmt.Printf("tmpRSlice2: %v len=%v cap=%v (%p) \n", tmpRSlice, len(slice), cap(slice), &tmpRSlice)

	for i := l; i < r; i++ {
		if tmpLSlice[leftIndex] <= tmpRSlice[rightIndex] {
			slice[i] = tmpLSlice[leftIndex]
			leftIndex++
		} else {
			slice[i] = tmpRSlice[rightIndex]
			rightIndex++
		}
	}

	return slice
}
