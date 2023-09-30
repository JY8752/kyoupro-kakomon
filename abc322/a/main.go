package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		N int
		S string
	)

	fmt.Scan(&N, &S)

	i := strings.Index(S, "ABC")

	var result int
	if i == -1 {
		result = i
	} else {
		result = i + 1
	}

	fmt.Println(result)
}
