package stdout

import (
	"fmt"
	"strings"
)

func PrintIntSlice(slice []int) {
	fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
}
