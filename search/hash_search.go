package search

// 0 ~ n-1
func primaryHash(key int, n int) int {
	return key % n
}

// 1 ~ n-1
func secondaryHash(key int, n int) int {
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
			hashTable[hash] = key
			return hash
		} else if hashTable[hash] == key {
			return hash
		} else {
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
			return hash
		} else if hashTable[hash] == 0 {
			return -1
		} else {
			i++
		}
	}

	return -1
}
