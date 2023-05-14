package sort

func BubbleSort(slice []int) []int {
	n := len(slice)

	// iが増える度にソート済み領域が増える
	// n-iがソート済み領域の最初のインデックス
	for i := 0; i < n; i++ {
		// 現在のインデックスより1つ前と比較するので1始まり
		// jがソート済み領域のインデックスに侵入しないようにする
		for j := 1; j < n-i; j++ {
			// 大小関係が逆になっていたら入れ替える
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
		}
	}

	return slice
}

func BubbleSort2(slice []int) []int {
	n := len(slice)
	// 値を入れ替えたか管理する変数
	// ループを起動させるために初回だけtrueにしておく
	swapped := true

	// 最後まで走査して入れ替えが発生しなければ終了
	for swapped {
		swapped = false

		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}

		// 最後まで走査する度に、最後の値がソート済みとなる
		// 不要な走査を減らすためにnをデクリメントして走査範囲を狭めていく
		n--
	}

	return slice
}
