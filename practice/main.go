package main

import "fmt"

func interactiveSorting() {
	var n, q int
	var s string

	fmt.Scan(&n, &q)

	for i := 0; i < n; i++ {
		s += string(rune('A' + i))
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {
			fmt.Printf("? %c %c\n", s[i], s[j])
			var ans string
			fmt.Scan(&ans)
			if ans == ">" {
				s = s[:i] + s[j:j+1] + s[i:i+1] + s[j+1:]
			} else {
				s = s[:j] + s[i:i+1] + s[j:j+1] + s[i+1:]
			}
		}
	}

	fmt.Printf("! %s\n", s)
}

func main() {
	interactiveSorting()
}
