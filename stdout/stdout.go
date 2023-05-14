package stdout

import (
	"fmt"
	"strings"
)

func PrintSlice[T any](slice []T) {
	res := fmt.Sprintf("%v", slice)
	res = strings.Trim(res, "[]")
	fmt.Println(res)
}
