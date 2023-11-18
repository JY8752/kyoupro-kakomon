package main

import "fmt"

func main() {
	var (
		N int
		S string
	)

	fmt.Scan(&N, &S)

	m := make(map[string]int, N)
	var (
		counter int
		current string
	)

	runes := []rune(S)
	for i, r := range runes {
		counter++
		current = string(r)
		if i == len(runes)-1 || current != string(runes[i+1]) {
			if m[current] < counter {
				m[current] = counter
			}
			counter = 0
		}
	}

	var result int
	for _, v := range m {
		result += v
	}

	fmt.Println(result)
}
