package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// input1 int value N
	scan(scanner)
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("input value must be integer. err: %s\n", err.Error())
	}

	// input2 sequece of int An
	a := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		scan(scanner)
		ii, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("input value must be sequence of integers. err: %s\n", err.Error())
		}
		a[i] = ii
		sum += ii
	}

	// array B after operation
	b := make([]int, len(a))
	p := sum / n

	for i := range b {
		b[i] = p
	}

	for i := 0; i < sum%n; i++ {
		b[n-1-i]++
	}

	sort.Ints(a)

	var answer int
	for i := 0; i < n; i++ {
		answer += int(math.Abs(float64(a[i] - b[i])))
	}

	fmt.Println(answer / 2)
}
