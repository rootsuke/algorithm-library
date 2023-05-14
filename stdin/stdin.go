package stdin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 文字列を1行取得
func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	return strings.TrimSpace(stringInput)
}

// 整数値1つ取得
func IntStdin() (int, error) {
	stringInput := StrStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

// デリミタで分割して文字列スライスを取得
func SplitStrStdin(delim string) []string {
	// 文字列で1行取得
	stringInput := StrStdin()

	// 空白で分割
	return splitWithoutEmpty(stringInput, delim)
}

// デリミタで分割して整数値スライスを取得
func SplitIntStdin(delim string) ([]int, error) {
	// 文字列スライスを取得
	stringSplited := SplitStrStdin(" ")

	var intSlices []int
	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err := strconv.Atoi(stringSplited[i])
		if err != nil {
			panic(err)
		}
		intSlices = append(intSlices, iparam)
	}

	return intSlices, nil
}

// 空白や空文字だけの値を除去したSplit()
func splitWithoutEmpty(stringTargeted string, delim string) []string {
	stringSplited := strings.Split(stringTargeted, delim)

	var stringReturned []string
	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return stringReturned
}
