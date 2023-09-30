package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./tests/sample-2.in"
	close    func() error
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		debug(dataPath)
	}

	sc.Split(bufio.ScanWords)
}

func debug(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	sc = bufio.NewScanner(f)
	close = func() error { return f.Close() }
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	if close != nil {
		defer close()
	}

	var (
		N int
		M int
		A []int
	)

	N = scanInt(sc)
	M = scanInt(sc)

	A = make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = scanInt(sc)
	}

	var counter int
	results := make([]int, N)
	for i := 0; i < N; i++ {
		elm := A[counter] - (i + 1)
		results[i] = elm
		if elm == 0 {
			counter++
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
