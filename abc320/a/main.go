package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	result := int(math.Pow(a, b) + math.Pow(b, a))
	fmt.Println(result)
}
