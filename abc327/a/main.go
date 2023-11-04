package main

import "fmt"

func main() {
	var (
		N int
		S string
	)

	fmt.Scan(&N, &S)

	runes := []rune(S)
	var result bool
	for i, r := range runes {
		if r == 'a' && i != N-1 {
			if runes[i+1] == 'b' {
				result = true
				break
			}
		}

		if r == 'a' && i != 0 {
			if runes[i-1] == 'b' {
				result = true
				break
			}
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
