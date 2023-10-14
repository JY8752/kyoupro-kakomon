package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	// if N%6 == 0 || N%2 == 0 || N%3 == 0 || N == 1 {
	// 	fmt.Println("Yes")
	// } else {
	// 	fmt.Println("No")
	// }
	var resutl bool
loop:
	for i := 0; i < math.MaxInt; i++ {
		a := int(math.Pow(2, float64(i)))
		if a > N {
			break loop
		}
		for j := 0; j < math.MaxInt; j++ {
			b := int(math.Pow(3, float64(j)))
			num := a * b
			if num > N {
				continue loop
			}
			if num == N {
				resutl = true
				break loop
			}
		}
	}

	if resutl {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
