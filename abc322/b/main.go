package main

import "fmt"

func main() {
	var (
		N int
		M int
		S string
		T string
	)

	fmt.Scan(&N, &M, &S, &T)

	prefix := T[:N]
	suffix := T[M-N:]

	var result int
	if prefix == S && suffix == S {
		result = 0
	}
	if prefix == S && suffix != S {
		result = 1
	}
	if prefix != S && suffix == S {
		result = 2
	}
	if prefix != S && suffix != S {
		result = 3
	}

	fmt.Println(result)
}
