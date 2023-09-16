package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)

	var result = 1
	for i := 0; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			targetStr := s[i:j]
			if targetStr == reverse(targetStr) {
				if result < len(targetStr) {
					result = len(targetStr)
				}
			}
		}
	}

	// // 偶数
	// for i := 2; i <= n; i += 2 {
	// 	if allMatchRune(s[:i], func(s string, i int) bool { return s[i] == s[i+1] }) {
	// 		if result < len(s[:i]) {
	// 			result = len(s[:i])
	// 		}
	// 	}
	// }

	// // 奇数
	// for i := 3; i <= n; i += 2 {
	// 	targetStr := s[:i]
	// 	if targetStr == reverse(targetStr) {
	// 		if result < len(targetStr) {
	// 			result = len(targetStr)
	// 		}
	// 	}
	// }

	fmt.Println(result)
}
