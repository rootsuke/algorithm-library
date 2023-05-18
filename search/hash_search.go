package search

// 1次ハッシュ
func primaryHash(key int, n int) int {
	return key % n
}

// 2次ハッシュ
func secondaryHash(key int, n int) int {
	// 再計算したハッシュ値が同じにならないように
	// 1以上かつハッシュテーブルの長さ未満の値を返す
	return 1 + (key % (n - 1))
}

func calcHash(key int, i int, n int) int {
	return (primaryHash(key, n) + i*secondaryHash(key, n)) % n
}

func Insert(hashTable []int, key int) int {
	n := len(hashTable)
	i := 0

	for i < n {
		hash := calcHash(key, i, n)
		if hashTable[hash] == 0 {
			// 格納しそのハッシュ値を返す
			hashTable[hash] = key
			return hash
		} else if hashTable[hash] == key {
			// すでに同じ値が格納されていればそのハッシュ値を返す
			return hash
		} else {
			// 衝突した場合はハッシュ値を再計算する
			i++
		}
	}

	return -1
}

func Search(hashTable []int, key int) int {
	n := len(hashTable)
	i := 0
	for i < n {
		hash := calcHash(key, i, n)
		if hashTable[hash] == key {
			// 見つかったらハッシュ値を返す
			return hash
		} else if hashTable[hash] == 0 {
			// 空だった場合は格納されていないと判定し-1を返す
			return -1
		} else {
			// 違う値が格納されていた場合はハッシュ値を再計算する
			i++
		}
	}

	return -1
}
