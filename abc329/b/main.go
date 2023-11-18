package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	var max int
	for i := 0; i < N; i++ {
		var v int
		fmt.Scan(&v)
		A[i] = v
		if max < v {
			max = v
		}
	}

	var result int
	for _, a := range A {
		if result < a && a != max {
			result = a
		}
	}

	fmt.Println(result)

	// var max, second int
	// for _, a := range A {
	// 	if max < a {
	// 		max = a
	// 		continue
	// 	}

	// 	if second < a && max != a {
	// 		second = a
	// 		continue
	// 	}
	// }

	// fmt.Println(second)
}
