package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	arr := make([]string, len(S))
	for i, r := range []rune(S) {
		arr[i] = string(r)
	}
	fmt.Println(strings.Join(arr, " "))
}
