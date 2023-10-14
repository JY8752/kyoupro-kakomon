package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	result := true
	for i := 1; i < N; i++ {
		if A[i-1] != A[i] {
			result = false
			break
		}
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
