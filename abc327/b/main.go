package main

import (
	"fmt"
	"math"
)

func main() {
	var B int
	fmt.Scan(&B)

	// min := B/2
	// max := B
	// for {
	// 	num := int64(math.Pow(float64(min), float64(min)))

	// }
	min := 1
	max := math.MaxInt64
	for min < max {
		center := (min + max) / 2
		r := int(math.Pow(float64(center), float64(center)))
		if r == B {
			fmt.Println(center)
			break
		}
		if r < B {
			min = center
		}
		if r > B {
			max = center
		}
	}

	// for i := 1; ; i++ {
	// 	r := int64(math.Pow(float64(i), float64(i)))
	// 	if r == B {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// 	if r > B {
	// 		fmt.Println(-1)
	// 		break
	// 	}
	// }
}
