package main

import (
	"fmt"
)

const inf = 1 << 61

func main() {
	var (
		m int
		s []string
	)

	fmt.Scan(&m)

	s = make([]string, 3)
	for i := 0; i < 3; i++ {
		var val string
		fmt.Scan(&val)
		s[i] = val
	}

	orders := [6][3]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	ans := inf
	for x := '0'; x <= '9'; x++ {
	nextOrder:
		for _, order := range orders {
			counter := 0
			for t := 0; t < 3*m; t++ {
				i := t % m
				if s[order[counter]][i] == byte(x) {
					counter++
					if counter == 3 {
						ans = min(ans, t)
						continue nextOrder
					}
				}
			}
		}
	}

	if ans == inf {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

func min(nums ...int) int {
	result := 1 << 61
	for _, num := range nums {
		if result > num {
			result = num
		}
	}
	return result
}
